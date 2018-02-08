package main

import (
	_ "ffmpeg-web/routers"

	_ "github.com/astaxie/beego"
)

func main() {
	// beego.Run()
	testCmd()
}
