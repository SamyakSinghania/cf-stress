package main

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"log"
	"sync"
)

type application struct {
	mutex   sync.Mutex
	Counter int
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(logger)
	zap.ReplaceGlobals(logger)

	app := new(application)

	e := echo.New()
	e.GET("/", app.homeHandler)
	e.GET("/test/:contestID/:problemIndex/", app.GetStressTestHandler)
	e.POST("/test/:contestID/:problemIndex/", app.PostStressTestHandler)

	e.GET("/pricing", app.PricingHandler)

	e.GET("/status", app.GetStatusHandler)
	e.POST("/status", app.PostStatusHandler)

	e.GET("/subscribe", app.GetMailingList)
	e.POST("/subscribe", app.PostMailingList)

	e.GET("/status/:ticketID", app.TicketStatus)

	zap.S().Info("Starting server")
	e.Logger.Fatal(e.Start(":4000"))
}
