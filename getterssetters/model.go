package main

type User struct {
	Name      string
	Email     string
	Addresses []Address
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zipcode      string
}
