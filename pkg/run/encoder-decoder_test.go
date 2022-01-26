package run

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/kubetrail/base58/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestDecoderFlag(t *testing.T) {
	cmd := cobra.Command{}

	cmd.Flags().BoolP(flags.Decode, "d", false, "decode input")
	_ = viper.BindPFlag(flags.Decode, cmd.Flags().Lookup(flags.Decode))

	// check default
	if viper.GetBool(flags.Decode) {
		t.Fatal("expected default decode value to be false, found true")
	}

	// set to true
	_ = cmd.Flags().Set(flags.Decode, "true")
	if !viper.GetBool(flags.Decode) {
		t.Fatal("expected set decode value to be true, found false")
	}

	// set to false
	_ = cmd.Flags().Set(flags.Decode, "false")
	if viper.GetBool(flags.Decode) {
		t.Fatal("expected set decode value to be false, found true")
	}
}

func TestEncoder(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().BoolP(flags.Decode, "d", false, "decode input")

	bb := new(bytes.Buffer)
	bw := bufio.NewWriter(bb)
	cmd.SetOut(bw)

	in := "hello"
	out := "Cn8eVZg"
	if err := EncoderDecoder(cmd, []string{in}); err != nil {
		t.Fatal(err)
	}

	if err := bw.Flush(); err != nil {
		t.Fatal(err)
	}

	if bb.String() != out {
		t.Fatal("expected ", out, " got ", bb.String())
	}
}

func TestDecoder(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().BoolP(flags.Decode, "d", false, "decode input")
	_ = cmd.Flags().Set(flags.Decode, "true")

	bb := new(bytes.Buffer)
	bw := bufio.NewWriter(bb)
	cmd.SetOut(bw)

	out := "hello"
	in := "Cn8eVZg"
	if err := EncoderDecoder(cmd, []string{in}); err != nil {
		t.Fatal(err)
	}

	if err := bw.Flush(); err != nil {
		t.Fatal(err)
	}

	if bb.String() != out {
		t.Fatal("expected ", out, " got ", bb.String())
	}
}
