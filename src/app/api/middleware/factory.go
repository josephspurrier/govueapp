package middleware

import (
	"net/http"

	"app/api"
	"app/api/middleware/cors"
	"app/api/middleware/jwt"
	"app/api/middleware/logrequest"
)

// Factory will return the http.Handler wrapped in middleware.
func Factory(h http.Handler, l logrequest.ILog, whitelist []string,
	webtoken api.IToken, ctx api.IContext) http.Handler {
	// JWT validation.
	token := jwt.New(whitelist, webtoken, ctx)
	h = token.Handler(h)

	// CORS for the endpoints.
	h = cors.Handler(h)

	// Log every request.
	lr := logrequest.New()
	lr.SetLog(l)
	h = lr.Handler(h)

	return h
}
