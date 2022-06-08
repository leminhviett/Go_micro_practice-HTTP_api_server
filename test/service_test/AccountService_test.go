package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/domain"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/dto"
	"github.com/leminhviett/Go_micro_practice-HTTP_api_server/service"
	mock_domain "github.com/leminhviett/Go_micro_practice-HTTP_api_server/test/mocks/domain"
)

var mockAccountRepo *mock_domain.MockAccountRepo
var accountService service.AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockAccountRepo = mock_domain.NewMockAccountRepo(ctrl)
	accountService = service.NewDefAccountService(mockAccountRepo)

	return func() {
		defer ctrl.Finish()
	}
}

func Test_insert_new_account_succ_case(t *testing.T) {
	tear_down := setup(t)
	defer tear_down()

	accountReq := dto.AccountRequestDTO{AccountName: "vietle", Ammount:123}
	accountDomain := domain.Account{AccountName: accountReq.AccountName, Ammount: accountReq.Ammount}
	returnedReq := domain.Account{AccountName: accountReq.AccountName, Ammount: accountReq.Ammount}
	returnedReq.AccountId = "123"

	mockAccountRepo.EXPECT().InsertNewAccount(&accountDomain).Return(&returnedReq, nil)

	//Act
	dtoAccResponse, err := accountService.InsertNewAccount(&accountReq)

	//asert
	if dtoAccResponse.AccountId != returnedReq.AccountId {
		t.Error("fail account ID assertion")
	}
	if err != nil {
		t.Error("fail error assertion")
	}
}