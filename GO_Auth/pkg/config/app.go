package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var conf Config

func Configure() {
	conf.GoogleLoginConfig = oauth2.Config{
		ClientID:     "Google_Client_Id",
		ClientSecret: "Client_Secret",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://127.0.0.1:5000/login/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}

func GetConf() Config {
	return conf
}
