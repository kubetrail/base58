package run

import (
	"fmt"
	"io"
	"strings"

	"github.com/kubetrail/base58/pkg/flags"
	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// EncoderDecoder encodes input to base58 and decodes base58 input
func EncoderDecoder(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Decode, cmd.Flags().Lookup(flags.Decode))
	decode := viper.GetBool(flags.Decode)

	if decode {
		var input string
		if len(args) > 0 {
			if len(args) > 1 {
				return fmt.Errorf("please provide just one argument")
			}
			input = args[0]
		} else {
			b, err := io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
			input = strings.Trim(string(b), "\n")
		}
		out, err := base58.Decode(input)
		if err != nil {
			return fmt.Errorf("failed to decode input: %w", err)
		}

		if _, err := cmd.OutOrStdout().Write(out); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
	} else {
		var input []byte
		if len(args) > 0 {
			input = []byte(strings.Join(args, " "))
		} else {
			var err error
			input, err = io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
		}

		out := base58.Encode(input)

		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s", out); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
	}

	return nil
}
