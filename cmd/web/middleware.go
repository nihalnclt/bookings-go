package main

import "net/http"

// SessionLoad loads and saves the session request on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
