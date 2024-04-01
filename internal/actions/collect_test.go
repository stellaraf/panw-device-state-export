package actions_test

import (
	"testing"

	"github.com/stellaraf/panw-device-state-export/internal/actions"
	"github.com/stellaraf/panw-device-state-export/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CollectDeviceStateExports(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		outDir := t.TempDir()
		results, err := actions.CollectDeviceStateExports(
			false,
			outDir,
			test.Env.HostnamePattern,
			test.Env.PanoramaHost,
			test.Env.PanoramaUsername,
			test.Env.PanoramaPassword,
			test.Env.FirewallUsername,
			test.Env.FirewallPassword,
		)
		require.NoError(t, err)
		assert.NotEqual(t, 0, len(results))
	})
}
