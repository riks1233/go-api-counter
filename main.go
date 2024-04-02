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
	r.GET("/", HttpGetCount(myCounter))
	r.GET("/increment", HttpIncrement(myCounter))
	r.GET("/decrement", HttpDecrement(myCounter))
	r.Run()
}
