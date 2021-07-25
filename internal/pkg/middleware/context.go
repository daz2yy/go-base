package middleware

import (
	"github.com/gin-gonic/gin"
)

// UsernameKey defines the key in gin context which represents the owner of the secret.
// const UsernameKey = "username"
// Defines common log fields.
const (
	KeyRequestID string = "requestID"
	KeyUsername  string = "username"
	UsernameKey         = "username"
)

// Context is a middleware that injects common prefix fields to gin.Context.
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KeyRequestID, c.GetString(XRequestIDKey))
		c.Set(KeyUsername, c.GetString(UsernameKey))
		c.Next()
	}
}
