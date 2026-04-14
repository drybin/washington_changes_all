package registry

import (
	"testing"
	"time"

	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/stretchr/testify/require"
)

func Test_newTelegramRestyClient(t *testing.T) {
	t.Parallel()

	t.Run("nil config", func(t *testing.T) {
		t.Parallel()
		_, err := newTelegramRestyClient(nil)
		require.Error(t, err)
	})

	t.Run("no proxy", func(t *testing.T) {
		t.Parallel()
		c, err := newTelegramRestyClient(&config.TgConfig{Timeout: 5 * time.Second})
		require.NoError(t, err)
		require.NotNil(t, c)
	})

	t.Run("bad scheme", func(t *testing.T) {
		t.Parallel()
		_, err := newTelegramRestyClient(&config.TgConfig{
			Timeout:     5 * time.Second,
			Socks5Proxy: "http://127.0.0.1:8080",
		})
		require.Error(t, err)
	})

	t.Run("socks5", func(t *testing.T) {
		t.Parallel()
		c, err := newTelegramRestyClient(&config.TgConfig{
			Timeout:     5 * time.Second,
			Socks5Proxy: "socks5://127.0.0.1:1080",
		})
		require.NoError(t, err)
		require.NotNil(t, c)
	})

	t.Run("socks5h", func(t *testing.T) {
		t.Parallel()
		c, err := newTelegramRestyClient(&config.TgConfig{
			Timeout:     5 * time.Second,
			Socks5Proxy: "socks5h://127.0.0.1:1080",
		})
		require.NoError(t, err)
		require.NotNil(t, c)
	})
}
