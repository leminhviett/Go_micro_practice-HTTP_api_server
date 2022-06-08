package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/dto"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/service"
	"github.com/leminhviett/Go_micro_practice-utils_lib/errs"
)

type AccountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) AccountHandler {
	return AccountHandler{service: service}
}

func (handler AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accs, err := handler.service.FindAll()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, accs)
}

func (handler AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id := args["id"]

	acc, err := handler.service.FindById(id)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusAccepted, acc)
}

func (handler AccountHandler) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	var accReq dto.AccountRequestDTO

	err := json.NewDecoder(r.Body).Decode(&accReq)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewValidationError("Error while decoding input body"))
		return
	}

	acc, err2 := handler.service.InsertNewAccount(&accReq)

	if err2 != nil {
		writeResponse(w, http.StatusInternalServerError, err2.AsMessage())
		return
	}
	writeResponse(w, http.StatusAccepted, acc)
}