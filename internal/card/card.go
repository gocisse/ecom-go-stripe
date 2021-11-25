package card

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

// card type will hold information to use stripe
type Card struct {
	Secret   string
	Key      string
	Currency string
}

// Transaction will hold our transactions with stripe 
type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFourDigit       string
	BankReturnCode      string
}

// call our paymentIntent function 
func (c *Card)Charge(currency string, amount int)(*stripe.PaymentIntent, string, error){
	return c.CreatePaymentIntent(currency, amount)
}

//Create a function to charge credit card 
func (c *Card)CreatePaymentIntent(currency string, amount int)(*stripe.PaymentIntent, string, error){
	//access your stripe secret key 
	stripe.Key = c.Secret

	//create a payment Intent 
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	// params.AddMetadata("key", "value")

	// lets get our payment intent with a var "pi"
	pi, err :=  paymentintent.New(params)
	if err != nil {
		var msg = ""
		stripeErr, ok := err.(*stripe.Error)
		if ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err

	}

	return pi, "", err
}

//lets write our card error message 
func cardErrorMessage(code stripe.ErrorCode) string{
	var msg = ""
	switch code{
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your card has expired"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "Incorrect CVC code"
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect zip/postal code"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "The amount is too Large to charge your card"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "The amount is too Small to charge your card"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Insufficient balance"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Your postal code is invalid"
	default:
		msg = "Your card was declined"
	}

	return msg
}