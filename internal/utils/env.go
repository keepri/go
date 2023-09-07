package utils

import (
	"log"
	"os"
)

func Env(name string, is_required bool) string {
	v := os.Getenv(name)
	if v == "" && is_required {
		log.Fatalf("%v environment variable not set\n", name)
	}
	return v
}
