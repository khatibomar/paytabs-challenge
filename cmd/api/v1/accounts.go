package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/khatibomar/paytabs-challenge/internal/customerrors"
	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
)

func (app *application) createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name    string  `json:"name"`
		Balance float64 `json:"balance,omitempty"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account := &datastructure.Account{
		Guid:    uuid.New().String(),
		Name:    input.Name,
		Balance: input.Balance,
	}

	if err := account.ValidateAccount(); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.store.Add(account)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/accounts/%s", account.Guid))

	err = app.writeJSON(w, http.StatusCreated, envelope{"account": account}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	account, err := app.store.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrAccountDoesNotExist):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts := app.store.All()
	err := app.writeJSON(w, http.StatusOK, envelope{"accounts": accounts}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) depositAccountsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID     string  `json:"id"`
		Amount float64 `json:"amount"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account, err := app.store.Get(input.ID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrAccountDoesNotExist):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = account.Deposit(input.Amount)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) withdrawAccountsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID     string  `json:"id"`
		Amount float64 `json:"amount"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	account, err := app.store.Get(input.ID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrAccountDoesNotExist):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = account.Withdraw(input.Amount)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) transferTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FromID string  `json:"from"`
		ToID   string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fromAccount, err := app.store.Get(input.FromID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrAccountDoesNotExist):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	toAccount, err := app.store.Get(input.FromID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrAccountDoesNotExist):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = fromAccount.Transfer(toAccount, input.Amount)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": fromAccount}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
