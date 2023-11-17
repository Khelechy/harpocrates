package main

import (
	"github.com/khelechy/harpocrates/utils"
)

func main() {
	utils.Set("passworda", "this is my passworda")

	utils.Get("password")
	utils.Get("passworda")
}
