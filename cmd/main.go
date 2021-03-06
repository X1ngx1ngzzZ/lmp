package main

import (
	"lmp_ui/routers"

	"github.com/gin-gonic/gin"
)

func Run() {
	//initDB()

	initWeb()
}

func initWeb() {
	r := gin.Default()
	r.LoadHTMLFiles("../static/mulSelect1.html")
	routers.RegisterRouter(r)
	r.Run()
}

func main() {
	Run()
}
