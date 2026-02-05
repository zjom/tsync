package oauth

var scopes = []string{}

func AddScope(s ...string) {
	scopes = append(scopes, s...)
}
