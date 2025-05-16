package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PSS2134/go_restapi/router"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	fmt.Println("Welcome to REST API Series!")
	r := router.Router()
	handler := c.Handler(r)
	fmt.Println("Starting server on the port 8000...")
	//whatsapp.SendTemplateMessage(919199329417)
	//invoice.GetInvoice()
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", handler))
	fmt.Println("Server listening on the port 8000...")

}
