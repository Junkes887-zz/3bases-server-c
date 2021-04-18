package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Junkes887/3bases-server-c/controller"
	"github.com/Junkes887/3bases-server-c/database"
	"github.com/Junkes887/3bases-server-c/repository"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/olivere/elastic/v7"
	"github.com/rs/cors"
)

func injects(client *elastic.Client, context context.Context) controller.Client {
	data := repository.Client{
		DB:  client,
		CTX: context,
	}

	controller := controller.Client{
		DB:  client,
		REP: data,
	}

	return controller
}

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")

	client := database.CreateConnection()
	controller := injects(client, context.Background())

	router := httprouter.New()
	router.GET("/", controller.FindAll)
	router.GET("/:cpf", controller.Find)
	router.POST("/", controller.Save)
	router.PUT("/:id", controller.Upadate)
	router.DELETE("/:id", controller.Delete)

	c := cors.AllowAll()
	handlerCors := c.Handler(router)

	fmt.Println("Listem " + PORT + ".....")
	http.ListenAndServe(":"+PORT, handlerCors)
}
