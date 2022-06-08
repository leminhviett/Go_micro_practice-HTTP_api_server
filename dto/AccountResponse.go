package dto

type AccountResponseDTO struct {
	AccountId string `json:"account_id"`
	AccountName string `json:"account_name"`
	Ammount float64 `json:"balance"`
}