package goodbye

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/desaias/argocd-demo/config"
	"github.com/desaias/argocd-demo/internal/middleware"
	"go.uber.org/zap"
)

type GoodbyeService struct {
	env    config.Config
	logger *zap.Logger
}

func NewGoodByeService(env config.Config, logger *zap.Logger) *GoodbyeService {
	app := GoodbyeService{
		env:    env,
		logger: logger,
	}

	return &app
}

func (app *GoodbyeService) Run() error {
	mux := http.NewServeMux()
	mux.Handle("GET /", middleware.LoggingMiddleware2(app.logger, http.HandlerFunc(goodbyeHandler)))

	app.logger.Sugar().Infof("Starting server on :%d", app.env.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", app.env.Port), mux)
}

type GoodbyeMessage struct {
	Text string `json:"text"`
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	response := GoodbyeMessage{Text: "Goodbye!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
