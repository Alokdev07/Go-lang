package main

import "fmt"

type payment struct {
	gateway paymenter
}
type paymenter interface{
	pay(amount float32)
} // this interface should solve the problem where pay method was there and its parameter is the same we can directly pass in the payment gateway

func (payment payment) make_payment(amount float32) {
	// razorpay_payment := razorpay{}
	// razorpay_payment.pay(amount)
	// stripe := stripe{}
	// stripe.make_stripe_payment(int(amount))

	payment.gateway.pay(amount)
}

type razorpay struct {}
type stripe struct{}

func (stripe stripe) pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}

func (razor razorpay) pay(amount float32) {
	// logic to implement payment
	fmt.Println("making payment using razorpay",amount)
}

type paypal struct{}

func (paypal paypal) pay(amount float32){
	fmt.Println("payment through paypal " , amount)
}


func main() {
	// razorpay := razorpay{}
	// stripe := stripe{}
	paypal := paypal{}
	payment := payment{
		gateway: paypal,
	}
	payment.make_payment(100)
}