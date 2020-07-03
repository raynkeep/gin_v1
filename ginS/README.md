# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/ryanker/gin_v1"
	"github.com/ryanker/gin_v1/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
