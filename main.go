package main

import (
	"context"
	"errors"
	http2 "github.com/Arvin619/url-shortener/adapter/delivery/http"
	"github.com/Arvin619/url-shortener/adapter/delivery/http/middleware"
	"github.com/Arvin619/url-shortener/adapter/repository"
	"github.com/Arvin619/url-shortener/application/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	repo := repository.NewInMemRepo()

	addShortUrlUseCase := usecase.NewAddShortUrlUseCase(repo)
	takeOutOriginalUrlUseCase := usecase.NewTakeOutOriginalUrlUseCase(repo)

	controller := http2.NewUrlShortenerController(takeOutOriginalUrlUseCase, addShortUrlUseCase)

	// TODO 加上 Swagger
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())
	v1 := router.Group("api/v1")
	shortUrl := v1.Group("short")
	shortUrl.GET(":shortId", controller.TakeOutOriginalUrl)
	shortUrl.POST("", controller.AddShortUrl)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s\n", err)
	}
	log.Println("Server exiting")
}
