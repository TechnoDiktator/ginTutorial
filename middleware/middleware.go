package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized",
			"message": "Token not found"})
		return
	}

	// Perform authentication logic here
	// For example, check if the user is logged in or has a valid token

	// If authentication fails, abort the request with an error
	// c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})

	// If authentication succeeds, continue to the next middleware or handler
	c.Next()
}


func AddHeader(c *gin.Context) {
	// Add a custom header to the response
	c.Writer.Header().Set("Key", "MyCustomValue")
	c.Next()
}


// lets write the middleware function in a different way
// func Authenticate() gin.HandlerFunc {

// 	//we can write custom logic here before the authenticate middleware
// 	// for example we can log the request
// 	log.Println("Request received at: ", time.Now())

// 	return func(c *gin.Context) {
// 		if !(c.Request.Header.Get("Token") == "auth") {
// 			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized",
// 				"message": "Token not found"})
// 			return
// 		}

// 		// Perform authentication logic here
// 		// For example, check if the user is logged in or has a valid token

// 		// If authentication fails, abort the request with an error
// 		// c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})

// 		// If authentication succeeds, continue to the next middleware or handler
// 		c.Next()
// 	}
// }
