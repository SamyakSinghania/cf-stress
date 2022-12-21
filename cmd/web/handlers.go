package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (app *application) homeHandler(c echo.Context) error {
	zap.S().Info("Executing homeHandler")
	bs, err := ioutil.ReadFile("ui/html/home.html")
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
	return c.HTMLBlob(http.StatusOK, bs)
}

func (app *application) GetStressTestHandler(c echo.Context) error {
	zap.S().Info("Executing GetStressTestHandler")
	//contestIDString := c.Param("contestID")
	////problemIndex := c.Param("problemIndex")
	//_, err := strconv.Atoi(contestIDString)
	//if err != nil {
	//	return c.String(http.StatusBadRequest,
	//		http.StatusText(http.StatusBadRequest))
	//}
	//return c.String(http.StatusOK, fmt.Sprintf("You are in GetStressTestHandler, contestID: %d,"+
	//	" problemIndex: %s", contestID, problemIndex))
	bs, err := ioutil.ReadFile("ui/html/testing.html")
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
	return c.HTMLBlob(http.StatusOK, bs)
}
func (app *application) PostStressTestHandler(c echo.Context) error {
	zap.S().Info("Executing PostStressTestHandler")
	contestIDString := c.Param("contestID")
	problemIndex := c.Param("problemIndex")
	contestID, err := strconv.Atoi(contestIDString)
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest))
	}
	submissionIDString := c.FormValue("submissionID")
	submissionID, err := strconv.Atoi(submissionIDString)
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest))
	}
	app.mutex.Lock()
	//defer app.mutex.Unlock()
	app.Counter++
	app.mutex.Unlock()

	return c.String(http.StatusOK, fmt.Sprintf("You are in PostStressTestHandler, contestID: %d, "+
		"problemIndex: %s with submission: %d. Your ticket number is %d ", contestID, problemIndex, submissionID, app.Counter))
}

func (app *application) PricingHandler(c echo.Context) error {
	zap.S().Info("Executing PricingHandler")
	bs, err := ioutil.ReadFile("ui/html/pricing.html")
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
	return c.HTMLBlob(http.StatusOK, bs)
	//return c.String(http.StatusOK, "You are in PricingHandler")
}

func (app *application) GetStatusHandler(c echo.Context) error {
	zap.S().Info("Executing GetStatusHandler")
	bs, err := ioutil.ReadFile("ui/html/status.html")
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
	return c.HTMLBlob(http.StatusOK, bs)
	//return c.String(http.StatusOK, "You are in GetStatusHandler")
}

func (app *application) PostStatusHandler(c echo.Context) error {
	zap.S().Info("Executing PostStatusHandler")
	bs, err := ioutil.ReadFile("ui/html/status.html")
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
	return c.HTMLBlob(http.StatusOK, bs)
	//return c.String(http.StatusOK, "You are in PostStatusHandler")
}

func (app *application) GetMailingList(c echo.Context) error {
	zap.S().Info("Executing GetMailingList")
	bs, err := ioutil.ReadFile("ui/html/mailing.html")
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
	return c.HTMLBlob(http.StatusOK, bs)
	//return c.String(http.StatusOK, "You are in GetMailingList")
}

func (app *application) PostMailingList(c echo.Context) error {
	zap.S().Info("Executing PostMailingList")
	bs, err := ioutil.ReadFile("ui/html/mailing.html")
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
	return c.HTMLBlob(http.StatusOK, bs)
	//return c.String(http.StatusOK, "You are in PostMailingList")
}

func (app *application) TicketStatus(c echo.Context) error {
	zap.S().Info("Executing TicketStatus")
	return c.String(http.StatusOK, "You are in TicketStatus")
}
