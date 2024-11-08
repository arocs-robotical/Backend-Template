package robotallocation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RobotAllocation struct {
	Location_x   int    `json:"location_x"`
	Location_y   int    `json:"location_y"`
	Location_z   int    `json:"location_z"`
	Product_take string `json:"product_take"`
	Product_put  string `json:"product_put"`
	Status       string `json:"status"`
}

type RobotAllocationService struct {
	PocketbaseURL string
}

func NewRobotAllocationService(url string) *RobotAllocationService {
	return &RobotAllocationService{PocketbaseURL: url}
}

// Get
func (s *RobotAllocationService) FetchRobotAllocation() ([]map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/collections/robot_allocation/records", s.PocketbaseURL))
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

	robots, ok := result["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	var formattedRobots []map[string]interface{}
	for _, item := range robots {
		robot, ok := item.(map[string]interface{})
		if ok {
			formattedRobots = append(formattedRobots, robot)
		}
	}

	return formattedRobots, nil
}

func (s *RobotAllocationService) FetchRobotAllocationByID(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/collections/robot_allocation/records/%s", s.PocketbaseURL, id)
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
func (s *RobotAllocationService) CreateRobotAllocation(robot_allocation RobotAllocation) (map[string]interface{}, error) {
	body, err := json.Marshal(robot_allocation)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/api/collections/robot_allocation/records", s.PocketbaseURL), "application/json", bytes.NewBuffer(body))
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
func (s *RobotAllocationService) UpdateRobotAllocation(id string, robot_allocation RobotAllocation) (map[string]interface{}, error) {
	body, err := json.Marshal(robot_allocation)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal flow_in: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/collections/robot_allocation/records/%s", s.PocketbaseURL, id), bytes.NewBuffer(body))
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
func (s *RobotAllocationService) DeleteRobotAllocation(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/collections/robot_allocation/records/%s", s.PocketbaseURL, id), nil)
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
