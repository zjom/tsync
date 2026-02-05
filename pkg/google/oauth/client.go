package oauth

import (
	"context"
	"log"
	"net/http"
	"os"
	"path"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GetClient retrieves a token, saves the token, then returns the generated client.
func GetClient(optFuncs ...ClientOpt) (*http.Client, error) {
	opts := defaultClientOpts()
	for _, optFn := range optFuncs {
		optFn(opts)
	}
	config, err := getGoogleOAuthOptsFromFile(opts.credentialsFilePath, scopes...)
	if err != nil {
		return nil, err
	}

	tok, err := tokenFromFile(opts.cachedTokenFilePath)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		if err := saveToken(opts.cachedTokenFilePath, tok); err != nil {
			log.Printf("Unable to cache oauth token: %v\n", err)
		}
	}
	return config.Client(context.Background(), tok), nil
}

type clientOpts struct {
	credentialsFilePath string
	cachedTokenFilePath string
	ctx                 context.Context
}

func defaultClientOpts() *clientOpts {
	co := &clientOpts{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		co.cachedTokenFilePath = ".gts-token"
	} else {
		co.cachedTokenFilePath = path.Join(homeDir, ".local", "state", "gts", ".gts-token")
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		co.credentialsFilePath = ".gts-credentials"
	} else {
		co.credentialsFilePath = path.Join(configDir, "gts", ".gts-credentials")
	}

	co.ctx = context.Background()

	return co
}

type ClientOpt func(*clientOpts)

func WithCredentialsFilePath(path string) func(*clientOpts) {
	return func(co *clientOpts) {
		co.credentialsFilePath = path
	}
}

func WithCachedTokenFilePath(path string) func(*clientOpts) {
	return func(co *clientOpts) {
		co.cachedTokenFilePath = path
	}
}

func WithContext(ctx context.Context) func(*clientOpts) {
	return func(co *clientOpts) {
		co.ctx = ctx
	}
}

func getGoogleOAuthOptsFromFile(path string, scopes ...string) (*oauth2.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		// log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		return nil, err
	}

	return config, nil
}
