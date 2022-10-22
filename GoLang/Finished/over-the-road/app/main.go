package main

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println("Connecting to DB")

	db := sqlx.Open("", "")

	strings.Compare("a", "b")

}
