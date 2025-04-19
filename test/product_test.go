package test

import (
	"encoding/json"
	"fiber_rest_api/database"
	"fiber_rest_api/route"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()

	db := database.GetConnection()

	api := app.Group("/api")

	route.ProductRoute(api, db)

	return app
}

func TestGetProduct(t *testing.T) {
	app := setupApp()

	request := httptest.NewRequest(http.MethodGet, "/api/products/1", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	expectedResponseBody := map[string]any{
		"status":  "success",
		"message": "Product retrieved successfully",
		"data": map[string]any{
			"id":       1,
			"name":     "Book",
			"quantity": 6,
		},
	}
	expectedJSON, err := json.Marshal(expectedResponseBody)

	assert.Nil(t, err)
	assert.Equal(t, string(expectedJSON), string(bytes))
	t.Log("✅")
}

func TestGetProducts(t *testing.T) {
	app := setupApp()

	request := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	expectedResponseBody := map[string]any{
		"status":  "success",
		"message": "Products retrieved successfully",
		"data": []map[string]any{
			{
				"id":       1,
				"name":     "Book",
				"quantity": 6,
			},
		},
	}
	expectedJSON, err := json.Marshal(expectedResponseBody)

	assert.Nil(t, err)
	assert.Equal(t, string(expectedJSON), string(bytes))
	t.Log("✅")
}

func TestStoreProduct(t *testing.T) {
	app := setupApp()

	body := strings.NewReader(`{"name":"Pencil","quantity":3}`)
	request := httptest.NewRequest(http.MethodPost, "/api/products", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	expectedResponseBody := map[string]any{
		"status":  "success",
		"message": "Product created successfully",
		"data": map[string]any{
			"id":       2,
			"name":     "Pencil",
			"quantity": 3,
		},
	}
	expectedJSON, err := json.Marshal(expectedResponseBody)

	assert.Nil(t, err)
	assert.Equal(t, string(expectedJSON), string(bytes))
	t.Log("✅")
}

func TestUpdateProduct(t *testing.T) {
	app := setupApp()

	body := strings.NewReader(`{"name":"Book","quantity":5}`)
	request := httptest.NewRequest(http.MethodPatch, "/api/products/1", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	expectedResponseBody := map[string]any{
		"status":  "success",
		"message": "Product updated successfully",
		"data": map[string]any{
			"id":       1,
			"name":     "Book",
			"quantity": 5,
		},
	}
	expectedJSON, err := json.Marshal(expectedResponseBody)

	assert.Nil(t, err)
	assert.Equal(t, string(expectedJSON), string(bytes))
	t.Log("✅")
}

func TestDeleteProduct(t *testing.T) {
	app := setupApp()

	request := httptest.NewRequest(http.MethodDelete, "/api/products/1", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	expectedResponseBody := map[string]any{
		"status":  "success",
		"message": "Product deleted successfully",
	}
	expectedJSON, err := json.Marshal(expectedResponseBody)

	assert.Nil(t, err)
	assert.Equal(t, string(expectedJSON), string(bytes))
	t.Log("✅")
}
