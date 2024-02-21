package main

import (
	"fmt"

	validation "github.com/beriloqueiroz/go-libs-validation"
)

func main() {
	fmt.Println(validation.Cpf("123654878"))
}
