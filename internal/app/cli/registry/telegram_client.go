package registry

import (
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/proxy"
)

func newTelegramRestyClient(tg *config.TgConfig) (*resty.Client, error) {
	if tg == nil {
		return nil, wrap.Errorf("telegram config is nil")
	}

	proxyRaw := strings.TrimSpace(tg.Socks5Proxy)
	if proxyRaw == "" {
		c := resty.New()
		c.SetTimeout(tg.Timeout)
		return c, nil
	}

	u, err := url.Parse(proxyRaw)
	if err != nil {
		return nil, wrap.Errorf("invalid TG_SOCKS5_PROXY URL: %w", err)
	}

	scheme := strings.ToLower(u.Scheme)
	if scheme != "socks5" && scheme != "socks5h" {
		return nil, wrap.Errorf("TG_SOCKS5_PROXY must use socks5 or socks5h scheme, got %q", u.Scheme)
	}
	if u.Host == "" {
		return nil, wrap.Errorf("TG_SOCKS5_PROXY: host is empty")
	}

	var auth *proxy.Auth
	if u.User != nil {
		pw, _ := u.User.Password()
		auth = &proxy.Auth{
			User:     u.User.Username(),
			Password: pw,
		}
	}

	baseDialer := &net.Dialer{
		Timeout:   tg.Timeout,
		KeepAlive: 30 * time.Second,
	}
	socksDialer, err := proxy.SOCKS5("tcp", u.Host, auth, baseDialer)
	if err != nil {
		return nil, wrap.Errorf("socks5 proxy setup: %w", err)
	}

	contextDialer, ok := socksDialer.(proxy.ContextDialer)
	if !ok {
		return nil, wrap.Errorf("socks5 dialer: no context support")
	}

	defaultTransport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return nil, wrap.Errorf("http.DefaultTransport is not *http.Transport")
	}
	transport := defaultTransport.Clone()
	transport.Proxy = nil
	transport.DialContext = contextDialer.DialContext

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   tg.Timeout,
	}

	c := resty.NewWithClient(httpClient)
	return c, nil
}
