package flowout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FlowOut struct {
	Name            string   `json:"name"`
	CustomerName    string   `json:"customer_name"`
	CustomerContact string   `json:"customer_contact"`
	CustomerAddress string   `json:"customer_address"`
	DeadlineSent    string   `json:"deadline_sent"`
	Status          string   `json:"status"`
	ProductOut      []string `json:"product_out"`
}

type FlowOutService struct {
	PocketbaseURL string
}

func NewFlowOutService(url string) *FlowOutService {
	return &FlowOutService{PocketbaseURL: url}
}

// Get
func (s *FlowOutService) FetchFlowOut() ([]map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/collections/flow_out/records", s.PocketbaseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flow_out from Pocketbase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	flow_out, ok := result["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	var formattedFlowOut []map[string]interface{}
	for _, item := range flow_out {
		flowout, ok := item.(map[string]interface{})
		if ok {
			formattedFlowOut = append(formattedFlowOut, flowout)
		}
	}

	return formattedFlowOut, nil
}

func (s *FlowOutService) FetchFlowOutByID(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/collections/flow_out/records/%s", s.PocketbaseURL, id)
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
func (s *FlowOutService) CreateFlowOut(flow_out FlowOut) (map[string]interface{}, error) {
	body, err := json.Marshal(flow_out)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/api/collections/flow_out/records", s.PocketbaseURL), "application/json", bytes.NewBuffer(body))
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
func (s *FlowOutService) UpdateFlowOut(id string, flow_out FlowOut) (map[string]interface{}, error) {
	body, err := json.Marshal(flow_out)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal flow_in: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/collections/flow_out/records/%s", s.PocketbaseURL, id), bytes.NewBuffer(body))
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
func (s *FlowOutService) DeleteFlowOut(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/collections/flow_out/records/%s", s.PocketbaseURL, id), nil)
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
