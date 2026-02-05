package oauth

import (
	"bufio"
	"net/http"
	"path/filepath"

	"encoding/json"

	"net/url"

	"golang.org/x/oauth2"

	"log"
	"os"

	"github.com/pkg/browser"

	"context"
	"fmt"
)

// getTokenFromWeb requests a token from the web, then returns the token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	// Create a channel to receive the auth code
	codeCh := make(chan string)

	port := getPortFromUrl(config.RedirectURL)
	// 1. Start a temporary local server
	server := &http.Server{Addr: ":" + port}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the code from the URL query string
		code := r.URL.Query().Get("code")
		if code != "" {
			fmt.Fprintf(w, "Authorization successful! You can close this window.")
			codeCh <- code // Send the code to our channel
			return
		}
		fmt.Fprintf(w, "Authorization failed.")
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// 2. Generate the URL and open it in the browser
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser: \n%v\n", authURL)
	fmt.Println("\n--> Press 'Enter' to open the browser automatically...")

	// Wait for user to hit Enter
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Now open the URL
	err := browser.OpenURL(authURL)
	if err != nil {
		fmt.Printf("Failed to open browser: %v. Please copy the link manually.\n", err)
	}

	// Note: You might need to manually click the link or use a library
	// like 'github.com/pkg/browser' to open it automatically.

	// 3. Wait for the code to come in from the web server
	authCode := <-codeCh

	// 4. Shut down the temporary server
	server.Shutdown(context.Background())

	// 5. Exchange the code for a token
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
		return nil, err
	}
	return tok, nil
}

// tokenFromFile retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// saveToken saves a token to a file path, creating required dir if not exist.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)

	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return fmt.Errorf("unable to create directory: %v", err)
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

func getPortFromUrl(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Port()
}
