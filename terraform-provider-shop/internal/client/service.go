package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
	ServiceID string  `json:"serviceId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}

func (c *Client) CreateService(ctx context.Context, req Service) (*Service, error) {
	body, _ := json.Marshal(req)

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/services", c.Endpoint),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create service, status: %d", resp.StatusCode)
	}

	var service Service
	if err := json.NewDecoder(resp.Body).Decode(&service); err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) GetServiceByID(ctx context.Context, id string) (*Service, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/services/%s", c.Endpoint, id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get service, status: %d", resp.StatusCode)
	}

	var service Service
	if err := json.NewDecoder(resp.Body).Decode(&service); err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) UpdateService(ctx context.Context, id string, req Service) error {
	body, _ := json.Marshal(req)

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s/services/%s", c.Endpoint, id),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update service, status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) DeleteService(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/services/%s", c.Endpoint, id),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete service, status: %d", resp.StatusCode)
	}

	return nil
}
