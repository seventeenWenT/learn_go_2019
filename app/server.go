package app

import (
	h "basic/cobra/utils"

	e "basic/cobra/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
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
		app.GET("/info",GetInfo)
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

type ServerInfo struct {
	Hostname string `json:"hostname"`
	LocalIp string  `json:"localIp"`
}


func GetInfo(c *gin.Context){
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", host)
	}

	rep := &ServerInfo{
		Hostname: host,
		LocalIp: "",
	}
	h.JSONR(c,http.StatusOK,e.SuccessMsg,rep)
}