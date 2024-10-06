package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// SessionLoad loads and saves the session request on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
