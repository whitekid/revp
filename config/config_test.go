package config

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestMultiViper(t *testing.T) {
	cmd1 := &cobra.Command{}
	cmd2 := &cobra.Command{}

	Client.Init("revp", cmd1.Flags())
	Server.Init("revps", cmd2.Flags())
}

func TestEnvConfig(t *testing.T) {
	type args struct {
		config func() string
	}
	tests := [...]struct {
		name string
		args args
	}{
		{keySecret, args{Secret}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envName := envReplacer.Replace("RP_" + strings.ToUpper(tt.name))

			require.NotEqual(t, "", tt.args.config())
			require.Equal(t, os.Getenv(envName), tt.args.config())
		})
	}
}
