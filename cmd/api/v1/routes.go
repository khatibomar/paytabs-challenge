package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/accounts", app.listAccountsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/accounts", app.createAccountHandler)
	router.HandlerFunc(http.MethodGet, "/v1/accounts/:id", app.showAccountHandler)

	router.HandlerFunc(http.MethodPost, "/v1/deposit", app.depositAccountsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/withdraw", app.withdrawAccountsHandler)

	return router
}
