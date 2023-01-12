package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shuklaritvik06/GoProjects/GO_AUTH/pkg/config"
)

type RedirectURL struct {
	URL string
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConf()
	url := conf.GoogleLoginConfig.AuthCodeURL("state")
	res, _ := json.Marshal(RedirectURL{
		URL: url,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConf().GoogleLoginConfig
	state := r.FormValue("state")
	if state != "state" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	if code == "" {
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User has denied Permission.."))
		}
	} else {
		token, err := conf.Exchange(context.Background(), code)
		if err != nil {
			return
		}
		cookie := http.Cookie{
			Name:    "oauth",
			Value:   token.AccessToken,
			Expires: time.Now().AddDate(0, 0, 2),
			// HttpOnly: true,
			Path: "/",
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "http://127.0.0.1:3000", http.StatusTemporaryRedirect)
	}
}
