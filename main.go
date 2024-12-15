// TODO:
// * Serve all of this
// * Dockerize
// * Account for PORT 80 / environment PORT variable somehow.
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Methods", "GET")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func HttpIncrement(counter *SafeCounter) func(c *gin.Context) {
	return func(c *gin.Context) {
		counter.Increment()
		c.String(http.StatusOK, counter.StringCount())
	}
}

func HttpDecrement(counter *SafeCounter) func(c *gin.Context) {
	return func(c *gin.Context) {
		counter.Decrement()
		c.String(http.StatusOK, counter.StringCount())
	}
}

func HttpGetCount(counter *SafeCounter) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, counter.StringCount())
	}
}

func main() {
	myCounter := NewSafeCounter()
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(CorsMiddleware)

	// r.Use(func() gin.HandlerFunc {
	// 	secureMiddleware := secure.New(secure.Options{
	// 		FrameDeny:             true,                              // Prevent clickjacking
	// 		ContentTypeNosniff:    true,                              // Prevent MIME-type sniffing
	// 		BrowserXssFilter:      true,                              // Enable XSS filter
	// 		// ContentSecurityPolicy: "default-src 'self'",              // Restrict resource loading
	// 		// ReferrerPolicy:        "strict-origin-when-cross-origin", // Referrer policy

	//    // STS or HSTS refers to things regarding HTTPS. As we're using nginx for https, I don't think we need this.
	// 		// SSLRedirect:           true,                              // Enforce HTTPS
	// 		// STSSeconds:            31536000,                          // 1 year HSTS
	// 		// STSIncludeSubdomains:  true,
	// 		// STSPreload:            true,
	// 	})
	// 	return func(c *gin.Context) {
	// 		if err := secureMiddleware.Process(c.Writer, c.Request); err != nil {
	// 			c.AbortWithError(http.StatusInternalServerError, err)
	// 			return
	// 		}
	// 		c.Next()
	// 	}
	// }())

	r.GET("/", HttpGetCount(myCounter))
	r.GET("/inc", HttpIncrement(myCounter))
	r.GET("/dec", HttpDecrement(myCounter))
	r.Run()
}
