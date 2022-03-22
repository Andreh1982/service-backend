package models

type Buyer struct {
	Id         int    `json:"id"`
	Nome       string `json:"nome"`
	Cpf        int    `json:"cpf"`
	Contatocel string `json:"contatocel"`
}

type Product struct {
	Id         int    `json:"id"`
	Nome       string `json:"nome"`
	Descricao  string `json:"descricao"`
	Quantidade int    `json:"quantidade"`
	Seller     int    `json:"seller"`
}

type Seller struct {
	Id          int    `json:"id"`
	Nome        string `json:"nome"`
	Cpf         int    `json:"cpf"`
	Contatocel  string `json:"contatocel"`
	Contatomail string `json:"contatomail"`
	Websitelink string `json:"websitelink"`
}
