package main

import (
	_ "github.com/gin-gonic/gin"
)

func main() {

	r := engine()

	r.Run(":3000")
}
