package password

// Inspired by and adapted from https://github.com/bitnami/bcrypt-cli/blob/master/cmd.go

import (
	"fmt"
	"io"
	"io/ioutil"

	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost = 10
)

type PasswordGenerator struct {
	cost     int
	inReader io.Reader
}

func NewPasswordGenerator(in io.Reader, cost int) PasswordGenerator {
	return PasswordGenerator{
		cost:     cost,
		inReader: in,
	}
}

func (g PasswordGenerator) Generate() ([]byte, error) {
	if g.cost < 4 || g.cost > 31 {
		return nil, fmt.Errorf("cost %d is outside allowed range (4,31)", g.cost)
	}

	data, err := ioutil.ReadAll(g.inReader)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword(data, g.cost)
	if err != nil {
		return nil, fmt.Errorf("error producing bcrypt hash: %v", err)
	}

	return hash, nil
}
