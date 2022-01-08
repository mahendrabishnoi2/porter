package porter

import (
	"fmt"

	"github.com/pkg/errors"
)

// HelloOptions Define flags and arguments for `porter hello`.
type HelloOptions struct {
	Name string
}

// Validate the options passed to `porter hello`.
func (o HelloOptions) Validate() error {
	if o.Name == "" {
		return errors.New("--name is required")
	}
	return nil
}

// Hello contains the implementation for `porter hello`.
func (p *Porter) Hello(opts HelloOptions) error {
	fmt.Printf("Hello %s!\n", opts.Name)
	return nil
}
