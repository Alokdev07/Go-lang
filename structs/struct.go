package main

import (
	"fmt"
	"time"
)

// user struct

type user struct{
	name string
	email string
}

// -> order struct

type order struct{
	_id string
	amount float64
	status string
	created_at time.Time // nano second precision
	user user
}

//receiver type
func (order *order) change_status(status string){
	order.status = status
}

//custom getter in go 
func (order order) getAmount() float64{
	return order.amount
}

func newOrder(_id string,amount float64,status string) *order {
	/*
	 if someone used it like class
	 -> so here is the constructor
	 -> initial setup goes here
	*/
	order := order{
		_id: _id,
		amount: amount,
		status: status,
	}
	return &order
}

func main(){
	first_order := order{
		_id: "first",
		amount: 100,
		status: "delivered",
		created_at: time.Now(),
	}
	second_order := order{
		_id : "second",
		amount : 200,
		status : "received",
		created_at : time.Now(),
	}
	fmt.Printf("%+v\n", first_order)
	fmt.Println("_id ->" , first_order._id)
	fmt.Printf("%+v\n", second_order)
	fmt.Println("_id ->" , second_order._id)
	first_order.change_status("received")
	fmt.Println(first_order)
	fmt.Println(first_order.getAmount())

	third_order := newOrder("third",300,"cancelled")
	fmt.Println(third_order)

	// -> short hand notation

	role := struct{
		role string
		isAccess bool
	} {
		role: "admin",
		isAccess: true,
	}
	fmt.Println(role)

	// embed struct

	user := user{
		name: "Alok",
		email: "alokbhuyan163@gmail.com",
	}

	receipt := order{
		_id: "full",
		amount: 500,
		status: "confirmed",
		created_at: time.Now(),
		user: user,
	}

	receipt.user.name = "Rajaram"
	fmt.Println(receipt)
}