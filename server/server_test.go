package server

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"github.com/whitekid/revp/config"
)

func TestMain(m *testing.M) {
	cmd := &cobra.Command{}
	config.Init("revp", cmd.Flags())

	cmd = &cobra.Command{}
	config.Client.Init("revp", cmd.Flags())

	cmd = &cobra.Command{}
	config.Server.Init("revps", cmd.Flags())

	os.Exit(m.Run())
}

func TestAllocatePort(t *testing.T) {
	portRange := config.Server.PortRange()

	for i := 0; i < 100; {
		i++
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ln, err := allocatePort()
			require.NoError(t, err)
			defer ln.Close()

			port := ln.Addr().(*net.TCPAddr).Port

			require.NotEqual(t, -1, port)
			require.GreaterOrEqual(t, port, portRange[0])
			require.LessOrEqual(t, port, portRange[1])
		})
	}
}
