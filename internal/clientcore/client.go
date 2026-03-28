package clientcore

import (
	"bytes"
	"context"
	"encoding/json"
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

type Config struct {
	BaseURL     string
	Lang        string
	HTTPClient  *http.Client
	MinInterval time.Duration
}

type Client struct {
	apiKey      string
	baseURL     string
	lang        string
	httpClient  *http.Client
	minInterval time.Duration

	mu          sync.Mutex
	nextRequest time.Time
}

type APIError struct {
	HTTPStatus int
	Code       int
	Message    string
	Body       string
}

type envelope struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

func New(apiKey string, cfg Config) (*Client, error) {
	baseURL := strings.TrimRight(strings.TrimSpace(cfg.BaseURL), "/")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	lang := strings.TrimSpace(cfg.Lang)
	if lang == "" {
		lang = DefaultLanguage
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}

	minInterval := cfg.MinInterval
	if minInterval <= 0 {
		minInterval = defaultMinInterval
	}

	return &Client{
		apiKey:      apiKey,
		baseURL:     baseURL,
		lang:        lang,
		httpClient:  httpClient,
		minInterval: minInterval,
	}, nil
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

func (c *Client) Do(ctx context.Context, path string, requestBody any, out any) error {
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
