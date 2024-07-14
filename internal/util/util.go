package util

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CloseChannel(ch chan error) {
	if _, ok := <-ch; ok {
		close(ch)
	}
}
