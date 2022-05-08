package configs

import (
	"sync"
)

var (
	Configs *config
	once    sync.Once
)

type config struct {
	AuthRealm        string
	AuthClientID     string
	AuthClientSecret string
	AuthState        string
	AuthIssuer       string
	AuthRedirectURL  string
	AuthEndpoint     string
	Cookie           string
}

func Set() {

	once.Do(func() {
		Configs = &config{}
		Configs.AuthClientID = "myschool"
		Configs.AuthClientSecret = "ty4qlzVZw3qwlzFY6Imn5rOkZclE9ZrZ"
		Configs.AuthIssuer = "http://localhost:8081/realms/myschool"
		Configs.AuthRealm = "myschool"
		Configs.AuthState = "prilax"
		Configs.AuthRedirectURL = "http://localhost:8000/auth/callback"
		Configs.AuthEndpoint = "http://localhost:8081/realms/myschool/protocol/openid-connect/auth"
		Configs.Cookie = "access-token"
	})

}
