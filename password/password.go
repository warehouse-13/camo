package password

// Inspired by and adapted from https://github.com/bitnami/bcrypt-cli/blob/master/cmd.go

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"io/ioutil"

	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost = 10
)

type PasswordGenerator struct {
	cost      int
	outWriter io.Writer
	inReader  io.Reader
}

func NewPasswordGenerator(in io.Reader, out io.Writer, cost int) PasswordGenerator {
	return PasswordGenerator{
		cost:      cost,
		inReader:  in,
		outWriter: out,
	}
}

func (g PasswordGenerator) Generate() error {
	if g.cost < 4 || g.cost > 31 {
		return fmt.Errorf("cost %d is outside allowed range (4,31)", g.cost)
	}

	data, err := ioutil.ReadAll(g.inReader)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword(data, g.cost)
	if err != nil {
		return fmt.Errorf("error producing bcrypt hash: %v", err)
	}

	enc := b64.StdEncoding.EncodeToString(hash)
	_, err = fmt.Fprintf(g.outWriter, "%s\n", enc)

	return err
}
