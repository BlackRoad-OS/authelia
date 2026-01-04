package authentication

import (
	"github.com/BlackRoad-OS/authelia/v4/internal/configuration/schema"
)

const (
	testLDAPURL  = "ldap://127.0.0.1:389"
	testLDAPSURL = "ldaps://127.0.0.1:389"
)

var (
	testLDAPAddress  = MustParseAddress(testLDAPURL)
	testLDAPSAddress = MustParseAddress(testLDAPSURL)
)

func MustParseAddress(input string) *schema.AddressLDAP {
	address, err := schema.NewAddress(input)
	if err != nil {
		panic(err)
	}

	return &schema.AddressLDAP{Address: *address}
}
