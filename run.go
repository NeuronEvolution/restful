package rest

import (
	"github.com/NeuronFramework/log"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func Run(initHandler func() (http.Handler, error)) {
	log.Init()
	logger := zap.L().Named("Run")

	// todo prometheus

	defer func() {
		if err := recover(); err != nil {
			logger.Error("recover", zap.Any("error", err))
		}
	}()

	h, err := initHandler()
	if err != nil {
		logger.Error("initHandler", zap.Error(err))
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Listen", zap.String("port", port))

	err = http.ListenAndServe(":"+port,
		Logging(Recovery(cors.AllowAll().Handler(h))))
	if err != nil {
		logger.Error("ListenAndServe", zap.Error(err))
	}
}
