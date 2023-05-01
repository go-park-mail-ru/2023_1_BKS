package post

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"post/config"
	"post/usecase"
	"syscall"
	"time"

	authmiddlevare "pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	v2 "post/delivery/http/v2"
)

func AsyncRunHTTP(e *echo.Echo, cfg config.Config) error {
	go func() {
		err := e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port))
		if err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	timeout := time.Duration(10)
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return e.Shutdown(ctx)
}

// func AsyncRunGrpc(server *grpc.Server, lis net.Listener, cfg config.Config) error {
// 	go func() {
// 		err := server.Serve(lis)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
// 			os.Exit(1)
// 		}
// 	}()

// 	interrupt := make(chan os.Signal, 1)
// 	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
// 	<-interrupt

// 	server.GracefulStop()

// 	return nil
// }

func Run(cfg config.Config) {

	ctx := context.Background()

	swagger, err := v2.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки спецификации swagger\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	command, query := usecase.NewUsecase(ctx, cfg)

	serverHandler := v2.CreateHttpServer(command, query)

	e := echo.New()

	fa, err := authmiddlevare.NewInstanceAuthenticator()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	mw, err := authmiddlevare.CreateMiddleware(fa, swagger, "AppUniqFrontend", "AppUniqUser")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}
	e.Use(middleware.Logger())
	e.Use(mw...)

	v2.RegisterHandlers(e, &serverHandler)
	errs := make(chan error, 1)
	go func() {
		errs <- AsyncRunHTTP(e, cfg)
	}()
	log.Warn("Terminating aplication:", err)
}
