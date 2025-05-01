package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not provided in the env file or the .env file is missing")
	}else{
		fmt.Println("Port: ",portString)
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

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
} 