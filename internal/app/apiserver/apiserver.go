package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// Сконфигуренная инстанция структуры APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Функция запуска HTTP сервера, которая подключается к БД и прочем,
// Возвращает ошибку, если занят порт или бд не получилось подключить и тп.
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Функция конфигурации логгера, возвращает ошибку, если мы указываем неверный
// Уровень логгирования и тп.
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// Функция описывающая роутер
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	type request struct {
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello)")
	}
}
