package model

type WithdrawMoneyRequest struct {
	ID           string  `json:"id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	CPF          string  `json:"cpf" binding:"required"`
	CNPJ         string  `json:"cnpj" binding:"required"`
	Key          string  `json:"key" binding:"required"`
	BankName     string  `json:"bank_name" binding:"required"`
	Value        float64 `json:"value" binding:"required"`
	Request_Date string  `json:"request_date" binding:"required"`
}
