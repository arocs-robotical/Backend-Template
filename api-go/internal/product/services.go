package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name           string   `json:"name"`
	Varian         string   `json:"varian"`
	InStock        int      `json:"in_stock"`
	StockMinimum   int      `json:"stock_minimum"`
	StockDetailIn  []string `json:"stock_detail_in"`
	StockDetailOut []string `json:"stock_detail_out"`
}

type ProductService struct {
	PocketbaseURL string
}

func NewProductService(url string) *ProductService {
	return &ProductService{PocketbaseURL: url}
}

// Get
func (s *ProductService) FetchProducts() ([]map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/collections/products/records", s.PocketbaseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products from Pocketbase: %w", err)
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

func (s *ProductService) FetchProductByID(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/collections/products/records/%s", s.PocketbaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product by ID from Pocketbase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var product map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return product, nil
}

// Create
func (s *ProductService) CreateProduct(product Product) (map[string]interface{}, error) {
	body, err := json.Marshal(product)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/api/collections/products/records", s.PocketbaseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
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
func (s *ProductService) UpdateProduct(id string, product Product) (map[string]interface{}, error) {
	body, err := json.Marshal(product)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/collections/products/records/%s", s.PocketbaseURL, id), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create update request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
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
func (s *ProductService) DeleteProduct(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/collections/products/records/%s", s.PocketbaseURL, id), nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}
