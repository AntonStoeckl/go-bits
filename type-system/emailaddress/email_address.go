package emailaddress

import (
	"errors"
)

/***** EmailAddress interface *****/

// EmailAddress is an interface for multiple implementations of EmailAddresses
type EmailAddress interface {
	String() string
}

/***** UnverifiedEmailAddress type *****/

// UnverifiedEmailAddress is an EmailAddress that has not been verified yet
type UnverifiedEmailAddress struct {
	value string
}

func NewUnverifiedEmailAddressAsType(input string) *UnverifiedEmailAddress {
	return &UnverifiedEmailAddress{
		value: input,
	}
}

func NewUnverifiedEmailAddressAsInterface(input string) EmailAddress {
	return &UnverifiedEmailAddress{
		value: input,
	}
}

func (e *UnverifiedEmailAddress) String() string {
	return e.value
}

/***** VerifiedEmailAddress type *****/

// VerifiedEmailAddress is an EmailAddress that has been verified
type VerifiedEmailAddress struct {
	value string
}

func (e *VerifiedEmailAddress) String() string {
	return e.value
}

// Verify turns an UnverifiedEmailAddress into a VerifiedEmailAddress
func Verify(unverifiedEmailAddress *UnverifiedEmailAddress) *VerifiedEmailAddress {
	return &VerifiedEmailAddress{
		value: unverifiedEmailAddress.value,
	}
}

/***/

type Person struct {
	emailAddress EmailAddress
}

func NewPerson(emailAddress EmailAddress) (*Person, error) {
	if emailAddress == nil {
		return nil, errors.New("got nil emailAddress")
	}

	person := &Person{
		emailAddress: emailAddress,
	}

	return person, nil
}

func (p *Person) VerifyEmailAddress() {
	switch actual := p.emailAddress.(type) {
	case *UnverifiedEmailAddress:
		p.emailAddress = Verify(actual)
	case *VerifiedEmailAddress:
		// nothing to do
	case nil:
		panic("emailAddress is nil!")
	}
}
