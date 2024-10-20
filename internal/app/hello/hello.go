package hello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/desaias/argocd-demo/config"
	"github.com/desaias/argocd-demo/internal/middleware"
	"go.uber.org/zap"
)

type HelloService struct {
	env    config.Config
	logger *zap.Logger
}

func NewHelloService(env config.Config, logger *zap.Logger) *HelloService {
	return &HelloService{
		env:    env,
		logger: logger,
	}
}

func (app *HelloService) Run() error {
	mux := http.NewServeMux()
	mux.Handle("GET /", middleware.LoggingMiddleware2(app.logger, http.HandlerFunc(helloHandler)))

	app.logger.Sugar().Infof("Starting server on :%d", app.env.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", app.env.Port), mux)
}

type Message struct {
	Text string `json:"text"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Text: "Hello"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
