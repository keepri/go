package env

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type env struct {
	Name string
	Val  string
}

func Get(name string, is_required bool) *env {
	name_upper := strings.ToUpper(name)
	val := os.Getenv(name_upper)
	if val == "" && is_required {
		log.Fatalf("%v environment variable not set\n", name_upper)
	}
	return &env{
		Name: name_upper,
		Val:  val,
	}
}

func (e *env) Atoi() int {
	val, err := strconv.Atoi(e.Val)
	if err != nil {
		log.Fatalf("%v environment variable failed atoi\n", e.Name)
	}
	return val
}
