package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"

	p "github.com/warehouse-13/camo/password"
	"github.com/warehouse-13/camo/secret"
)

func main() {
	var (
		cost         int
		encode       bool
		createSecret bool
	)

	flag.IntVar(&cost, "cost", p.DefaultCost, "The cost weight, range of 4-31")
	flag.BoolVar(&encode, "encode", false, "Base64 encode the resulting hash (default false)")
	flag.BoolVar(&createSecret, "secret", false, "Add the resulting value to a k8s secret and print that out (default false)")
	flag.Parse()

	if !hasPipedStdin() {
		fmt.Println("usage: `echo -n foo | camo`")

		os.Exit(1)
	}

	pwg := p.NewPasswordGenerator(os.Stdin, cost)

	hash, err := pwg.Generate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result string

	switch {
	case createSecret:
		s, err := secret.Create(hash)
		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		result = string(s)
	case encode:
		result = b64.StdEncoding.EncodeToString(hash)
	default:
		result = string(hash)
	}

	if _, err := fmt.Fprintf(os.Stdout, "%s\n", result); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	os.Exit(0)
}

func hasPipedStdin() bool {
	// TODO: handle this err
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	return false
}
