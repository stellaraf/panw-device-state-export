package panw_test

import (
	"testing"

	"github.com/stellaraf/panw-device-state-export/internal/panw"
	"github.com/stellaraf/panw-device-state-export/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFirewall_Check(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		device := panw.Device{
			IPAddress:    Env.FirewallIP,
			Hostname:     "stcorp-fw01.phx01",
			Platform:     "PA-VM",
			SerialNumber: "007251000122716",
			Version:      "11.0.3-h5",
		}
		fw, err := panw.NewFirewall(device, test.Env.FirewallUsername, test.Env.FirewallPassword)
		require.NoError(t, err)
		ok, err := fw.Check()
		require.NoError(t, err)
		assert.True(t, ok)
	})
}

func TestFirewall_ExportDeviceState(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		device := panw.Device{
			IPAddress:    test.Env.FirewallIP,
			Hostname:     "stcorp-fw01.phx01",
			Platform:     "PA-VM",
			SerialNumber: "007251000122716",
			Version:      "11.0.3-h5",
		}
		fw, err := panw.NewFirewall(device, test.Env.FirewallUsername, test.Env.FirewallPassword)
		require.NoError(t, err)
		outDir := "/Users/mdl/Desktop"
		outFile, err := fw.ExportDeviceState(outDir)
		require.NoError(t, err)
		assert.NotEqual(t, "", outFile)
	})
}
