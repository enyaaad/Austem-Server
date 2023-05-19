package server

import (
	H "AustemServer/handlers"
	"AustemServer/package/auth"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer  *http.Server
	authUseCase auth.UseCase
}

func main() {

	r := gin.New()
	r.Use(H.CORSMiddleware())

	err := r.Run(":8080")
	if err != nil {
		return
	}

	fmt.Println("Starting Server at port :8080")

}

//func NewApp() *App {
//
//	authUseCase := any
//	return &App{
//		authUseCase: authUseCase,
//	}
//}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())

	api := router.Group("/")
	H.SignInEndpoint(api, a.authUseCase)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve : %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
