package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	app_http "product-img-service/internal/http"
	"product-img-service/internal/service/files"
	"syscall"
	"time"

	muxHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	logger := log.New(os.Stdout, "product-images", log.LstdFlags)

	files := files.New()

	handlers := app_http.New(logger, files)
	router.Use(muxHandler.CORS(muxHandler.AllowedOrigins([]string{"*"})))
	router.Path("/images/{id}").Methods("POST").HandlerFunc(handlers.FileUploader)
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	server := http.Server{
		Addr:         ":3000",
		Handler:      router,
		IdleTimeout:  time.Second * 120,
		ReadTimeout:  time.Second * 4,
		WriteTimeout: time.Second * 4,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	server.Shutdown(ctx)
}
