/// to create a "personal" type struct with fields "authorize" (boolean type), "contact", and "emailid"
//  print "emailid" if "authorize" is true

package main

import (
	"fmt"
)

type personal struct {
	authorize bool
	contact   string
	emailid   string
}

func main() {
	m := personal{authorize: false, contact: "abc", emailid: "a@b.com"}
	if m.authorize {
		fmt.Println(m.emailid)
	}

	n := personal{authorize: true, contact: "nnnn", emailid: "n@n.com"}
	if n.authorize {
		fmt.Println(n.emailid)
	}

}
