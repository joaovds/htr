package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joaovds/htr"
)

func main() {
	result, err := htr.LoadConfig("test.htr.yaml")
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(os.Stdout).Encode(result)
}
