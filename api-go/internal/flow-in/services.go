package flowin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FlowIn struct {
	Name             string   `json:"name"`
	VendorName       string   `json:"vendor_name"`
	ContactVendor    string   `json:"contact_vendor"`
	ScheduledArrived string   `json:"scheduled_arrived"`
	Status           string   `json:"status"`
	ProductIn        []string `json:"product_in"`
}

type FlowInService struct {
	PocketbaseURL string
}

func NewFlowInService(url string) *FlowInService {
	return &FlowInService{PocketbaseURL: url}
}

// Get
func (s *FlowInService) FetchFlowIn() ([]map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/collections/flow_in/records", s.PocketbaseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flow_in from Pocketbase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	products, ok := result["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	var formattedProducts []map[string]interface{}
	for _, item := range products {
		product, ok := item.(map[string]interface{})
		if ok {
			formattedProducts = append(formattedProducts, product)
		}
	}

	return formattedProducts, nil
}

func (s *FlowInService) FetchFlowInByID(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/collections/flow_in/records/%s", s.PocketbaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flow_in by ID from Pocketbase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// Create
func (s *FlowInService) CreateFlowIn(flow_in FlowIn) (map[string]interface{}, error) {
	body, err := json.Marshal(flow_in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/api/collections/flow_in/records", s.PocketbaseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create flow_in: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// Update
func (s *FlowInService) UpdateFlowIn(id string, flow_in FlowIn) (map[string]interface{}, error) {
	body, err := json.Marshal(flow_in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal flow_in: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/collections/flow_in/records/%s", s.PocketbaseURL, id), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create update request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update flow_in: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// Delete
func (s *FlowInService) DeleteFlowIn(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/collections/flow_in/records/%s", s.PocketbaseURL, id), nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete flow_in: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}
