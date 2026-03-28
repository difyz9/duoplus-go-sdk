package duoplus

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	DefaultBaseURL     = "https://openapi.duoplus.cn"
	DefaultIntlBaseURL = "https://openapi.duoplus.net"
	DefaultLanguage    = "zh"
	defaultTimeout     = 30 * time.Second
	defaultMinInterval = time.Second
)

type Option func(*Client)

type Client struct {
	apiKey      string
	baseURL     string
	lang        string
	httpClient  *http.Client
	minInterval time.Duration

	CloudPhones          *CloudPhoneService
	CloudNumbers         *CloudNumberService
	Groups               *GroupService
	Proxies              *ProxyService
	SubscriptionStartups *SubscriptionStartupService
	Apps                 *AppService
	CloudDisk            *CloudDiskService
	Automation           *AutomationService

	mu          sync.Mutex
	nextRequest time.Time
}

type APIError struct {
	HTTPStatus int
	Code       int
	Message    string
	Body       string
}

func (e *APIError) Error() string {
	parts := make([]string, 0, 3)
	if e.HTTPStatus > 0 {
		parts = append(parts, fmt.Sprintf("http_status=%d", e.HTTPStatus))
	}
	if e.Code > 0 {
		parts = append(parts, fmt.Sprintf("code=%d", e.Code))
	}
	if e.Message != "" {
		parts = append(parts, e.Message)
	}
	if len(parts) == 0 && e.Body != "" {
		parts = append(parts, e.Body)
	}
	if len(parts) == 0 {
		return "duoplus api error"
	}
	return "duoplus api error: " + strings.Join(parts, ", ")
}

type envelope struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("duoplus api key is required")
	}

	client := &Client{
		apiKey:      apiKey,
		baseURL:     DefaultBaseURL,
		lang:        DefaultLanguage,
		httpClient:  &http.Client{Timeout: defaultTimeout},
		minInterval: defaultMinInterval,
	}

	for _, opt := range opts {
		opt(client)
	}

	client.initServices()

	return client, nil
}

func (c *Client) initServices() {
	c.CloudPhones = &CloudPhoneService{client: c}
	c.CloudNumbers = &CloudNumberService{client: c}
	c.Groups = &GroupService{client: c}
	c.Proxies = &ProxyService{client: c}
	c.SubscriptionStartups = &SubscriptionStartupService{client: c}
	c.Apps = &AppService{client: c}
	c.CloudDisk = &CloudDiskService{client: c}
	c.Automation = &AutomationService{client: c}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		if strings.TrimSpace(baseURL) != "" {
			c.baseURL = strings.TrimRight(baseURL, "/")
		}
	}
}

func WithLanguage(lang string) Option {
	return func(c *Client) {
		if strings.TrimSpace(lang) != "" {
			c.lang = lang
		}
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

func WithMinInterval(interval time.Duration) Option {
	return func(c *Client) {
		if interval > 0 {
			c.minInterval = interval
		}
	}
}

func (c *Client) do(ctx context.Context, path string, requestBody any, out any) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.waitTurn(ctx); err != nil {
		return err
	}

	payload := []byte("{}")
	if requestBody != nil {
		var err error
		payload, err = json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Lang", c.lang)
	req.Header.Set("DuoPlus-API-Key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	var apiResp envelope
	if err := json.Unmarshal(body, &apiResp); err != nil {
		if resp.StatusCode >= http.StatusBadRequest {
			return &APIError{HTTPStatus: resp.StatusCode, Body: string(body)}
		}
		return fmt.Errorf("decode response: %w", err)
	}

	if resp.StatusCode >= http.StatusBadRequest || apiResp.Code != http.StatusOK {
		return &APIError{
			HTTPStatus: resp.StatusCode,
			Code:       apiResp.Code,
			Message:    apiResp.Message,
			Body:       string(body),
		}
	}

	if out == nil || len(apiResp.Data) == 0 || string(apiResp.Data) == "null" {
		return nil
	}

	if err := json.Unmarshal(apiResp.Data, out); err != nil {
		return fmt.Errorf("decode data: %w", err)
	}

	return nil
}

func (c *Client) waitTurn(ctx context.Context) error {
	c.mu.Lock()
	now := time.Now()
	when := c.nextRequest
	if when.Before(now) {
		when = now
	}
	c.nextRequest = when.Add(c.minInterval)
	c.mu.Unlock()

	wait := time.Until(when)
	if wait <= 0 {
		return nil
	}

	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
