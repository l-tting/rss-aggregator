package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load(".env")
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("Port not found in environment")
	}

	// Set up router instance
	router := chi.NewRouter()

	// CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*","http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, 
	}))


	v1router := chi.NewRouter()

	v1router.Get("/healthz",handlerReadiness)

	v1router.Get("/err",handlerErr)

	router.Mount("/v1",v1router)


	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Chi + CORS!"))
	})

	// Start the server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}
	log.Printf("Starting server on port %v", portstring)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
