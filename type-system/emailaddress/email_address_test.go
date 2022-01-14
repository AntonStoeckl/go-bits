package emailaddress_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"go-bits/type-system/emailaddress"
)

func TestEmailAddress(t *testing.T) {
	Convey(`Pointer to a concrete type - not nil`, t, func() {
		emailAddress := emailaddress.NewUnverifiedEmailAddressAsType("anton@example.com")
		_, err := emailaddress.NewPerson(emailAddress)

		Convey(`OK - no error`, func() {
			So(err, ShouldBeNil)
		})
	})

	Convey(`Interface with an underlying pointer type - not nil `, t, func() {
		emailAddress := emailaddress.NewUnverifiedEmailAddressAsInterface("anton@example.com")
		_, err := emailaddress.NewPerson(emailAddress)

		Convey(`OK - no error`, func() {
			So(err, ShouldBeNil)
		})
	})

	Convey(`Nil pointer`, t, func() {
		var emailAddress *emailaddress.UnverifiedEmailAddress
		_, err := emailaddress.NewPerson(emailAddress)

		Convey(`OK - error`, func() {
			So(err, ShouldBeError)
		})
	})

	Convey(`Nil interface`, t, func() {
		var emailAddress emailaddress.EmailAddress
		_, err := emailaddress.NewPerson(emailAddress)

		Convey(`OK - error`, func() {
			So(err, ShouldBeError)
		})
	})

	Convey(`Interface with an underlying pointer type - NIL!!!`, t, func() {
		var emailAddress emailaddress.EmailAddress
		var unverifiedEmailAddress *emailaddress.UnverifiedEmailAddress
		emailAddress = unverifiedEmailAddress

		Convey(`Trouble - test does not fail because the interface itself is not nil`, func() {
			So(emailAddress != nil, ShouldBeTrue)

			Convey(`We need to type-cast and check that the underlying is not nil to detect that we got some nil`, func() {
				concreteEmailAddress, ok := emailAddress.(*emailaddress.UnverifiedEmailAddress)
				So(ok, ShouldBeTrue)
				So(concreteEmailAddress != nil, ShouldBeTrue)
			})
		})
	})
}
