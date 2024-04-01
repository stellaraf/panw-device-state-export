package panw_test

import (
	"testing"

	"github.com/stellaraf/panw-device-state-export/internal/panw"
	"github.com/stellaraf/panw-device-state-export/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Panorama(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		pra, err := panw.NewPanorama(test.Env.PanoramaHost, test.Env.PanoramaUsername, test.Env.PanoramaPassword)
		require.NoError(t, err)
		check, err := pra.Check()
		require.NoError(t, err)
		assert.True(t, check)
	})
	t.Run("devices", func(t *testing.T) {
		pra, err := panw.NewPanorama(test.Env.PanoramaHost, test.Env.PanoramaUsername, test.Env.PanoramaPassword)
		require.NoError(t, err)
		devices, err := pra.ConnectedDevices()
		require.NoError(t, err)
		assert.True(t, len(devices) > 0)
	})
}
