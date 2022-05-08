package auths

import (
	"backend/configs"
	"context"

	"github.com/coreos/go-oidc"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func VerifyToken(token string) (*oidc.IDToken, map[string]interface{}, error) {
	idToken, err := getOauth2Verifier().Verify(context.Background(), token)
	if err != nil {
		return idToken, nil, err
	}
	claim := map[string]interface{}{}

	if err := idToken.Claims(&claim); err != nil {
		logrus.Println(err.Error())
		return nil, nil, err
	}
	return idToken, claim, nil
}

func getOauth2Config() *oauth2.Config {
	provider, err := oidc.NewProvider(context.Background(), configs.Configs.AuthIssuer)
	if err != nil {
		panic(err)
	}
	oauth2Config := oauth2.Config{
		ClientID:     configs.Configs.AuthClientID,
		ClientSecret: configs.Configs.AuthClientSecret,
		RedirectURL:  configs.Configs.AuthRedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}
	return &oauth2Config
}

func getOauth2Verifier() *oidc.IDTokenVerifier {
	provider, err := oidc.NewProvider(context.Background(), configs.Configs.AuthIssuer)
	if err != nil {
		panic(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: configs.Configs.AuthClientID,
	}

	return provider.Verifier(oidcConfig)
}

func ClaimsToUser(claims map[string]interface{}) User {

	user := User{}
	user.FirstName = claims["given_name"].(string)
	user.LastName = claims["family_name"].(string)
	user.FullName = claims["name"].(string)
	emailInterface := claims["email"]

	if emailInterface != nil {
		user.Email = emailInterface.(string)
	}

	realmAccess := claims["realm_access"]
	if realmAccess != nil {
		roles := realmAccess.(map[string]interface{})["roles"]
		if roles != nil {
			for _, v := range roles.([]interface{}) {
				user.Roles = append(user.Roles, v.(string))
			}
		}
	}

	return user
}
