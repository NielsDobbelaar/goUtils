package assert

import (
	"log"
)

func Assert(condition bool, msg string) {
	if !condition {
		log.Fatal(msg)
	}
}

func AssertWithError(condition bool, msg string, err error) {
	if !condition {
		log.Fatal(msg + err.Error())
	}
}
