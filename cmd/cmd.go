package cmd

import (
	"fmt"

	"github.com/jerwheaton/jwtutil/pkg/jwt"
	"github.com/spf13/cobra"
)

const (
	_maxJWTSize = 8 * 1024 * 1024
)

var (
	// Only decode for now
	rootCmd = &cobra.Command{
		Use:   "jwtutil",
		Short: "JWT Utilities",
		Long:  `A collection of utilities for working with JSON web tokens.`,
	}

	decodeCmd = &cobra.Command{
		Use:   "decode",
		Short: "Decode a JWT",
		Long:  `Decode a base 64 encoded JWT with a header, payload, and signature`,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := cmd.InOrStdin()
			inputBytes := make([]byte, _maxJWTSize)
			n, err := r.Read(inputBytes)
			if err != nil {
				return fmt.Errorf("error reading from stdin: %w", err)
			}

			if n == 0 {
				return fmt.Errorf("empty input")
			}

			decoded, err := jwt.Decode(inputBytes)
			if err != nil {
				return fmt.Errorf("error decoding jwt: %w", err)
			}

			fmt.Print(decoded)
			return nil
		},
	}
)

// RunCommand returns a cli command for running the application.
func RunCommand() *cobra.Command {
	rootCmd.AddCommand(decodeCmd)
	return rootCmd
}
