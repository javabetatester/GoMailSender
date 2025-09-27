package main

import (
	"GoMailSender/internal/domain/campaign"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

type Calculadora interface {
	Calcula(valor1, valor2 int) int
	Printa() (int, int)
}

type soma struct {
	valor1 int
	valor2 int
}

func (s *soma) Calcula(valor1, valor2 int) int {
	s.valor1 = valor1
	s.valor2 = valor2
	return valor1 + valor2

}

func (s *soma) Printa() (int, int) {
	return s.valor1, s.valor2
}

func main() {

	campaign := campaign.Campaign{
		Name:      "sexos",
		ID:        fmt.Sprintf("%s", xid.New().String()),
		Content:   "DADASDKASKDASKDKASDK",
		CreatedOn: time.Now(),
		Contacts:  []campaign.Contact{{Email: "test@test.com"}}}

	validate := validator.New()

	err := validate.Struct(campaign)

	if err == nil {
		fmt.Println("Nenhum erro de input")
	} else {
		validationErors := err.(validator.ValidationErrors)
		for _, v := range validationErors {
			fmt.Println(v.Error())
		}
	}
}
