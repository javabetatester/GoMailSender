package main

import (
	"GoMailSender/internal/domain/campaign"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

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
