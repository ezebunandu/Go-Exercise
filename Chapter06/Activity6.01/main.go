package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
}

func (d *directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}

func (d *directDeposit) report() {
	fmt.Printf("%v", d)
}
func main() {
	testDeposit := directDeposit{lastName: "", firstName: "Abe", bankName: "XYZ Inc.", routingNumber: 17, accountNumber: 1809}
	err := testDeposit.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	testDeposit.validateRoutingNumber()
	testDeposit.report()
}
