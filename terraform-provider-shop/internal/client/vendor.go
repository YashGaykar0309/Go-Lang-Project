package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Vendor struct {
	VendorID     string `json:"vendorId"`
	Name         string `json:"Name"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
	Address      string `json:"address"`
	Contact      string `json:"contact"`
}

func (c *Client) CreateVendor(ctx context.Context, req Vendor) (*Vendor, error) {
	body, _ := json.Marshal(req)

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/vendors", c.Endpoint),
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
		return nil, fmt.Errorf("failed to create vendor, status: %d", resp.StatusCode)
	}

	var vendor Vendor
	if err := json.NewDecoder(resp.Body).Decode(&vendor); err != nil {
		return nil, err
	}

	return &vendor, nil
}

func (c *Client) GetVendorByID(ctx context.Context, id string) (*Vendor, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/vendors/%s", c.Endpoint, id),
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
		return nil, fmt.Errorf("failed to get vendor, status: %d", resp.StatusCode)
	}

	var vendor Vendor
	if err := json.NewDecoder(resp.Body).Decode(&vendor); err != nil {
		return nil, err
	}

	return &vendor, nil
}

func (c *Client) UpdateVendor(ctx context.Context, id string, req Vendor) error {
	body, _ := json.Marshal(req)

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s/vendors/%s", c.Endpoint, id),
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
		return fmt.Errorf("failed to update vendor, status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) DeleteVendor(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/vendors/%s", c.Endpoint, id),
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
		return fmt.Errorf("failed to delete vendor, status: %d", resp.StatusCode)
	}

	return nil
}
