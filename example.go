package main

import (
	"fmt"

	"github.com/franela/goreq"
)

func main() {
	res, _ := goreq.Request{Uri: "http://www.google.com"}.Do()
	fmt.Println(res.StatusCode)
}
