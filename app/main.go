package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/EnsurityTechnologies/apiconfig"
)

func main() {
	key := os.Getenv("API_CONFIG_KEY")
	if key == "" {
		if len(os.Args) != 4 {
			fmt.Printf("Invalid Arguments")
			return
		}
		key = os.Args[3]
	} else {
		if len(os.Args) < 3 || len(os.Args) > 4{
			fmt.Printf("Invalid Arguments")
			return
		}
		if len(os.Args) == 4 {
			key = os.Args[3]
		}
	}

	configData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to open input file %v\n", err)
		return
	}
	err = apiconfig.CreateAPIConfig(os.Args[2], key, configData)

	if err != nil {
		fmt.Printf("Failed to create API config %v\n", err)
		return
	}
	fmt.Printf("API Config created successfully")

}
