package model

type Person struct {
	ID          string `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Contactinfo `json:"contactinfo,omitempty"`
}

type Contactinfo struct {
	City    string `json:"city,omitempty"`
	Zipcode string `json:"Zipcode,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
}
