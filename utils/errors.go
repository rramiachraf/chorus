package utils

import "log"

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
