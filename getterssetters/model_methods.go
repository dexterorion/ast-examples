package main

func (s *User) SetName(name string) {
	s.Name = name
}

func (s *User) SetEmail(email string) {
	s.Email = email
}

func (s *User) SetAddresses(addresses []Address) {
	s.Addresses = addresses
}

func (s *Address) SetAddressLine1(addressline1 string) {
	s.AddressLine1 = addressline1
}

func (s *Address) SetAddressLine2(addressline2 string) {
	s.AddressLine2 = addressline2
}

func (s *Address) SetCity(city string) {
	s.City = city
}

func (s *Address) SetState(state string) {
	s.State = state
}

func (s *Address) SetZipcode(zipcode string) {
	s.Zipcode = zipcode
}

