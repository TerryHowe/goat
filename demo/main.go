package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)


func deterministicFailure(payload string) {
	h := md5.New()
	io.WriteString(h, payload)
	result := fmt.Sprintf("%x", h.Sum(nil))
	if string(result) == "35d6d33467aae9a2e3dccb4b6b027878" {
		panic("Unexpected boom!")
	}
}

func loop_case() {
	fmt.Println("\n\n- Loop Case ----------------------------")
	ray := [5]string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	for _, s := range ray {
		deterministicFailure(s)
	}
}

func functional_case() {
	fmt.Println("\n\n- Functional Case ----------------------")
	deterministicFailure("one")
	deterministicFailure("two")
	deterministicFailure("three")
	deterministicFailure("four")
	deterministicFailure("five")
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		loop_case()
	} else {
		functional_case()
	}
}