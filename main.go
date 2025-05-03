package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aryansharma2k4/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}


func main(){
	godotenv.Load()
	portString := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("Port is not provided in the env file or the .env file is missing")
	}else{
		fmt.Println("Port: ",portString)
	}
	if dbUrl == "" {
		log.Fatal("DB URL is not provided in the env file or the .env file is missing")
	}else{
		fmt.Println("DB URL: ",dbUrl)
	}

	conn, err := sql.Open("postgres",dbUrl)

	if err != nil {
		log.Fatal("Can't connect to the database")
	}


	apiCfg := apiConfig{
		 DB: database.New(conn),
	}

	router := chi.NewRouter()

	log.Printf("Server starting on port %v", portString)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","DELETE","POST","OPTIONS","PUT"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users",apiCfg.handlerGetUser)
	router.Mount("/v1", v1Router )


	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}
	err = srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
} 
