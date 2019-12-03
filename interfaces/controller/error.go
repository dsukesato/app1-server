package controller

import "log"

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
