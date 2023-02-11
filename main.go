package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello dousheng")
	r := gin.Default()
	initRouter(r)
	r.Run()
}
