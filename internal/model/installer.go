package model

type InstallerData struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	CPF          string `json:"cpf"`
	CNPJ         string `json:"cnpj"`
	CompanyName  string `json:"company_name"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Complement   string `json:"complement"`
	CEP          string `json:"cep"`
	BirthDate    string `json:"birth_date"`
	Reference    string `json:"reference"`
}
