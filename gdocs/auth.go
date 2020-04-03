package gdocs

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// GetTokenFromWeb TODO
func GetTokenFromWeb(config *oauth2.Config, tokenViper *viper.Viper, tkFilePath string) error {
	err := launchAuth(config)
	if err != nil {
		return err
	}

	err = startServer(tokenViper)
	if err != nil && tokenViper.Get("code") == nil {
		return err
	}

	authCode, ok := (tokenViper.Get("code")).(string)
	if !ok {
		return errors.New("Failed to cast auth code to string")
	}

	token, err := exchangeCodeWithToken(config, authCode)
	if err != nil {
		return err
	}

	tokenViper.Set("token", *token)
	err = tokenViper.WriteConfigAs(tkFilePath)
	if err != nil {
		return err
	}

	return nil
}

// GetClient retrieves a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config, tok *oauth2.Token) *http.Client {
	return config.Client(context.Background(), tok)
}

// ConvertToToken TODO
func ConvertToToken(tkValues interface{}) (*oauth2.Token, error) {
	jsonTK, err := json.Marshal(tkValues)
	if err != nil {
		return nil, err
	}

	token := &oauth2.Token{}
	return token, json.Unmarshal(jsonTK, token)
}

// launchAuth requests a token from the web, then returns the retrieved token.
func launchAuth(config *oauth2.Config) error {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	err := openBrowser(authURL)
	if err != nil {
		return err
	}

	return nil
}

// exchangeCodeWithToken TODO
func exchangeCodeWithToken(config *oauth2.Config, authCode string) (*oauth2.Token, error) {
	return config.Exchange(oauth2.NoContext, authCode)
}
