package main

import (
	"flag"
	"fmt"
	"os"

	p "github.com/warehouse-13/camo/password"
)

func main() {
	var cost int
	flag.IntVar(&cost, "cost", p.DefaultCost, "The cost weight, range of 4-31 (default: 10)")
	flag.Parse()

	if !hasPipedStdin() {
		fmt.Println("usage: `echo -n foo | camo`")
		os.Exit(1)
	}

	password := p.NewPasswordGenerator(os.Stdin, os.Stdout, cost)
	if err := password.Generate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func hasPipedStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
