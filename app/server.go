package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

)


type Server struct {

}

func newServer() *Server{
	return  &Server{}
}


func (s *Server) Run(){
	r := gin.New()
	r.Use(gin.Logger())
	app := r.Group("/sa")
	{
		app.GET("/check", GetHealthCheck)
	}
	err := r.Run(":"+"8888")
	if err != nil{
		panic(err)
	}
}

func NewCobraCommand() *cobra.Command{
	server := newServer()
	cmd := &cobra.Command{
		Use: "Taiwei",
		Long: `The taiwei is a simple demo. This command have simple http `,
		Run: func(cmd *cobra.Command, args []string) {
			server.Run()
		},
	}
	return cmd
}

func GetHealthCheck(c *gin.Context){
	fmt.Println("request")
}
