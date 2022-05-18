package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	newid := uuid.New()
	fmt.Println(newid)
}
