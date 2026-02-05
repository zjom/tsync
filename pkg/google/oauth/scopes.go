package oauth

var scopes = []string{}

func RegisterScope(s ...string) {
	scopes = append(scopes, s...)
}
