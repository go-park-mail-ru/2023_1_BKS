package user

import (
	"config"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"google.golang.org/grpc"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"

	authmiddlevare "pkg/middleware"
	serverGrpc "user/delivery/grpc"
	v2 "user/delivery/http/v2"
	"user/usecase"
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

func AsyncRunGrpc(server *grpc.Server, lis net.Listener, cfg config.Config) error {
	go func() {
		err := server.Serve(lis)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
			os.Exit(1)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	server.GracefulStop()

	return nil
}

func CreateMiddleware(v JWSValidator, swagger *openapi3.T) ([]echo.MiddlewareFunc, error) {

	validator := oapimiddleware.OapiRequestValidatorWithOptions(swagger,
		&oapimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticator(v),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}

func Run(cfg config.Config) {

	ctx := context.Background()

	swagger, err := v2.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки спецификации swagger\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	command, query := usecase.NewUsecase(ctx, cfg)

	serverHandler := v2.CreateHttpServer(command, query)

	grpcHandler := serverGrpc.CreateGrpcServer(command, query)

	server := grpc.NewServer()

	e := echo.New()

	fa, err := NewAuthenticator()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}

	mw, err := authmiddlevare.CreateMiddleware(fa, swagger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки сервера grpc\n: %s", err)
		os.Exit(1)
	}
	e.Use(middleware.Logger())
	e.Use(mw...)

	e.Use(middleware.Logger())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))

	v2.RegisterHandlers(e, &serverHandler)

	serverGrpc.RegisterUserServer(server, &grpcHandler)

	errs := make(chan error, 2)
	go func() {
		errs <- AsyncRunHTTP(e, cfg)
	}()

	go func() {
		errs <- AsyncRunGrpc(server, lis, cfg)
	}()
	err = <-errs

	log.Warn("Terminating aplication:", err)
}
