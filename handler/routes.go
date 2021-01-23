package handler

import (
	"net/http"

	"SendGrid-with-Node.js/service/ping"
	"SendGrid-with-Node.js/service/sendgrid"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {

	ping := ping.NewEndpoint()
	sendgrid := sendgrid.NewEndpoint()

	r.transaction = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/ping",
			Endpoint:    ping.PingGetEndpoint,
		},
		{
			Name:        "Ping Pong : GET Prams",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/ping/:name",
			Endpoint:    ping.PingGetParamsEndpoint,
		},
		{
			Name:        "Ping Pong : POST Prams+Body",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodPost,
			Pattern:     "/ping/:name",
			Endpoint:    ping.PingPostParamsAndBodyEndpoint,
		},
		{
			Name:        "SendGrid : POST Body",
			Description: "SendGrid : Send mail",
			Method:      http.MethodPost,
			Pattern:     "/send",
			Endpoint:    sendgrid.SendGridPost,
		},
	}

	ro := gin.New()

	//ro.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"POST", "GET"},
	//	AllowHeaders:     []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//}))

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
