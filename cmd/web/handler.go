package main

import (
	"net/http"

	"github.com/gocisse/ecom-go-stripe/internal/models"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key

	if err := app.renderTemplate(w, r, "terminal", &templateData{
		StringMap: stringMap,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
		panic(err)

	}

}

func (app *application) succeededPayment(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
	}

	var (
		cardholder      = r.Form.Get("cardholder_name")
		email           = r.Form.Get("cardholder_email")
		paymentintent   = r.Form.Get("payment_intent")
		paymentmethod   = r.Form.Get("payment_method")
		paymentamount   = r.Form.Get("payment_amount")
		paymentcurrency = r.Form.Get("payment_currency")
	)
	data := make(map[string]interface{})

	data["cardholder"] = cardholder
	data["email"] = email
	data["pi"] = paymentintent
	data["pm"] = paymentmethod
	data["pa"] = paymentamount
	data["pc"] = paymentcurrency

	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
		return

	}
}

// Display a page to buy one widgets
func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key

	widget := models.Widget{
		ID:             1,
		Name:           "Custom Widget",
		Description:    "A very nice widget",
		InventoryLevel: 10,
		Price:          1000,
	}
	data := make(map[string]interface{})
	data["widget"] = widget

	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		StringMap: stringMap,
		Data:      data,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
		return

	}
}
