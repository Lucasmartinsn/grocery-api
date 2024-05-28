package Server

import (
	"log"

	"github.com/Lucasmartinsn/grocery-api/Server/Routers"
	"github.com/gin-gonic/gin"
)

type Serve struct {
	port  string
	serve *gin.Engine
	origemIP string
}

func NewServer() Serve {
	return Serve{
		port:  "5000",
		serve: gin.Default(),
		origemIP: "0.0.0.0", 
	}
}

func (s *Serve) Run() {
	router := Routers.ConfigRoutes(s.serve)

	log.Print("server is running:", s.port)
	log.Fatal(router.Run(s.origemIP + ":" + s.port))
}