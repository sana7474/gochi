package main

import (
	"log"

	"gochi/core"
)

func main() {
	res := core.SumProp(22, "М", "Т")
	log.Println(res)
}
