package main

import (
	"fmt"

	"github.com/CodyMcCarty/aggreGATOR/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg)

	err = cfg.SetUser("Cody")
	if err != nil {
		fmt.Println(err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
