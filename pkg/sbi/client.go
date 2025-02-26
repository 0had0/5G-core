package sbi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/0had0/5G-core/pkg/common/errors"
	"github.com/0had0/5G-core/pkg/common/logger"
	"github.com/0had0/5G-core/pkg/common/metrics"
	"go.uber.org/zap"
)

// Client handles HTTP requests to 5G Core Network Functions
type Client struct {
	httpClient *http.Client
	serviceName string
}

// NewClient creates a new SBI client
func NewClient(serviceName string, timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		serviceName: serviceName,
	}
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, url string, target interface{}) error {
	return c.doRequest(ctx, http.MethodGet, url, nil, target)
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, url string, body interface{}, target interface{}) error {
	return c.doRequest(ctx, http.MethodPost, url, body, target)
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, url string, body interface{}, target interface{}) error {
	return c.doRequest(ctx, http.MethodPut, url, body, target)
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, url string) error {
	return c.doRequest(ctx, http.MethodDelete, url, nil, nil)
}

// doRequest performs the HTTP request
func (c *Client) doRequest(ctx context.Context, method, url string, body, target interface{}) error {
	startTime := time.Now()
	
	// Create request
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return errors.NewInternalError("Failed to marshal request body", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}
	
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return errors.NewInternalError("Failed to create request", err)
	}
	
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		metrics.RequestCounter.WithLabelValues(c.serviceName, method, "error").Inc()
		return errors.NewInternalError(fmt.Sprintf("Failed to execute request to %s", url), err)
	}
	defer resp.Body.Close()
	
	// Record metrics
	duration := time.Since(startTime).Seconds()
	metrics.RequestDuration.WithLabelValues(c.serviceName, method).Observe(duration)
	metrics.RequestCounter.WithLabelValues(c.serviceName, method, fmt.Sprintf("%d", resp.StatusCode)).Inc()
	
	// Log the request
	logger.Debug("SBI request",
		zap.String("method", method),
		zap.String("url", url),
		zap.Int("status", resp.StatusCode),
		zap.Float64("duration", duration),
	)
	
	// Check for error status codes
	if resp.StatusCode >= 400 {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			// If we can't decode the error response, just return a generic error
			return c.mapStatusCodeToError(resp.StatusCode, url)
		}
		
		// Return a more specific error based on the response
		errorMsg, _ := errorResponse["message"].(string)
		if errorMsg == "" {
			errorMsg = fmt.Sprintf("Request failed with status %d", resp.StatusCode)
		}
		
		return c.mapStatusCodeToError(resp.StatusCode, errorMsg)
	}
	
	// Decode the response if a target was provided
	if target != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
			return errors.NewInternalError("Failed to decode response", err)
		}
	}
	
	return nil
}

// mapStatusCodeToError maps HTTP status codes to appropriate error types
func (c *Client) mapStatusCodeToError(statusCode int, message string) error {
	switch statusCode {
	case http.StatusBadRequest:
		return errors.NewBadRequestError(message, nil)
	case http.StatusUnauthorized:
		return errors.NewUnauthorizedError(message, nil)
	case http.StatusForbidden:
		return errors.NewForbiddenError(message, nil)
	case http.StatusNotFound:
		return errors.NewNotFoundError(message, nil)
	case http.StatusConflict:
		return errors.NewConflictError(message, nil)
	case http.StatusRequestTimeout, http.StatusGatewayTimeout:
		return errors.NewTimeoutError(message, nil)
	default:
		return errors.NewInternalError(message, nil)
	}
}
