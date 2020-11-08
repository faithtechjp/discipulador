package model

type Person struct {
	ID          string `json:"id,omitempty"`
	Firstname   string `json:"nome,omitempty"`
	Contactinfo `json:"contato,omitempty"`
	Birthdate   string `json:"aniversario,omitempty"`
}

type Contactinfo struct {
	Address string `json:"endereco,omitempty"`
	City    string `json:"cidade,omitempty"`
	Zipcode string `json:"cep,omitempty"`
	Phone   string `json:"telefone,omitempty"`
	Email   string `json:"email,omitempty"`
}
