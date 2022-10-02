package main

import (
	"net/http"

	"mygin"
)

func main() {
	r := mygin.Default()
	r.GET("/", func(c *mygin.Context) {
		c.String(http.StatusOK, "Hello mygin\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *mygin.Context) {
		names := []string{"mygin"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
