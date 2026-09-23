package main

import (
	"fmt"
	"os"

	"github.com/DaliaReds/gator/internal/config"
)

type state struct {
	config config.Config
}

func newState() (state, error) {
	cfg, err := config.Read()
	if err != nil {
		return state{}, err
	}

	return state{cfg}, nil
}

func main() {
	st, err := newState()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read configuration: %s\n", err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "login":
		if len(os.Args) < 3 {
			fmt.Println("Username required")
			os.Exit(1)
		}

		if err := st.config.SetUser(os.Args[2]); err != nil {
			fmt.Fprintf(os.Stderr, "Could not update configuration: %s\n", err.Error())
			os.Exit(2)
		}
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
