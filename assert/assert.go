package assert

import (
	"fmt"
	"log"
	"log/slog"
)

type ErrorContext struct {
	Name  string
	Value any
}

func AssertNoFatal(msg string, err error) {
	if err != nil {
		log.Println(msg, ":")
		slog.Error(err.Error())
		fmt.Println("")
	}
}

func LogContext(context ...ErrorContext) {
	if context != nil {
		log.Println("Current context:")
		for _, c := range context {
			log.Println("\t", c.Name, ": ", c.Value)
		}
		fmt.Println("")
	}
}

func Assert(condition bool, msg string) {
	if !condition {
		log.Fatal(msg)
	}
}

func AssertWithError(condition bool, msg string, err error) {
	if !condition {
		log.Fatal(msg, ": \n", err.Error())
	}
}

func AssertWithErrorAndContext(condition bool, msg string, err error, context ...ErrorContext) {
	if !condition {
		log.Println(msg)
		if context != nil {
			log.Println("context:")
			for _, c := range context {
				log.Println("\t", c.Name, ": ", c.Value)
			}
		}
		log.Fatal("error: ", err.Error())
	}
}
