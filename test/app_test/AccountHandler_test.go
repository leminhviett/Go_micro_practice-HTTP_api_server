package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/app"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/dto"
	mock_service "github.com/leminhviett/Go_micro_practice-HTTP_api_server/test/mocks/service"
	"github.com/leminhviett/Go_micro_practice-utils_lib/errs"
)

var router *mux.Router
var ch app.AccountHandler
var mockAccountService *mock_service.MockAccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockAccountService = mock_service.NewMockAccountService(ctrl)
	ch = app.NewAccountHandler(mockAccountService)
	router = mux.NewRouter()
	router.HandleFunc("/accounts", ch.GetAllAccounts)

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

//Test GetAllAccounts
func Test_successful_return_accounts_info(t *testing.T) {
	//Arrange
	tear_down := setup(t)
	defer tear_down()

	accounts := []*dto.AccountResponseDTO {
		{"1", "VietLe", 123},
	}
	mockAccountService.EXPECT().FindAll().Return(accounts, nil)

	request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_fail_return_error(t *testing.T) {
	tear_down := setup(t)
	defer tear_down()

	mockAccountService.EXPECT().FindAll().Return(nil, errs.NewInternalError("Internal Error"))

	request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}