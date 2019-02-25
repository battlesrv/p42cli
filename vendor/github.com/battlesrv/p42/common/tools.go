package common

import (
	"crypto/sha256"
	"fmt"

	"github.com/urfave/cli"
)

// CheckFlags ..
func CheckFlags(c *cli.Context, f ...string) error {
	for _, flag := range f {
		if !c.IsSet(flag) {
			cli.ShowSubcommandHelp(c)
			return fmt.Errorf("flag -%s is required", flag)
		}
	}
	return nil
}

// Sha256Sum ..
func Sha256Sum(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
