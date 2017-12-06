package yunpian

import (
	"context"
	"net"
	"net/http"
	"time"
)

// Config - SDK 配置项，可以将允许用户配置的参数放入这里
type Config struct {
	UseSSL     *bool
	HTTPClient *http.Client
	APIKey     *string
	Context    context.Context

	userHost  string
	signHost  string
	tplHost   string
	smsHost   string
	voiceHost string
	flowHost  string
}

func (c *Config) WithAPIKey(key string) *Config {
	c.APIKey = &key
	return c
}

func (c *Config) WithUseSSL(use bool) *Config {
	c.UseSSL = &use
	return c
}

func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

func (c *Config) WithContext(ctx context.Context) *Config {
	c.Context = ctx
	return c
}

// MergeIn merges the passed in configs into the existing config object.
func (c *Config) MergeIn(cfgs ...*Config) {
	for _, other := range cfgs {
		mergeInConfig(c, other)
	}
}

func mergeInConfig(dst *Config, other *Config) {
	if other == nil {
		return
	}

	if other.APIKey != nil {
		dst.APIKey = other.APIKey
	}
	if other.UseSSL != nil {
		dst.UseSSL = other.UseSSL
	}
	if other.HTTPClient != nil {
		dst.HTTPClient = other.HTTPClient
	}
	if other.Context != nil {
		dst.Context = other.Context
	}
}

// DefaultDevConfig returns a default dev client config pointer
func DefaultDevConfig() *Config {
	cfg := &Config{
		userHost:  "test-api.yunpian.com",
		signHost:  "test-api.yunpian.com",
		tplHost:   "test-api.yunpian.com",
		smsHost:   "test-api.yunpian.com",
		voiceHost: "test-api.yunpian.com",
		flowHost:  "test-api.yunpian.com",
	}
	return cfg.WithUseSSL(true).WithHTTPClient(defaultHTTPClient())
}

// DefaultConfig returns a default client config pointer
func DefaultConfig() *Config {
	cfg := &Config{
		userHost:  "sms.yunpian.com",
		signHost:  "sms.yunpian.com",
		tplHost:   "sms.yunpian.com",
		smsHost:   "sms.yunpian.com",
		voiceHost: "voice.yunpian.com",
		flowHost:  "flow.yunpian.com",
	}
	return cfg.WithUseSSL(true).WithHTTPClient(defaultHTTPClient())
}

func defaultHTTPClient() *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:        100,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	return &http.Client{Transport: transport}
}
