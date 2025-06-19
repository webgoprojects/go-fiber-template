package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	app := fiber.New()
	RegisterHealthRoutes(app)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Check response body
	buf := make([]byte, resp.ContentLength)
	resp.Body.Read(buf)
	resp.Body.Close()
	assert.JSONEq(t, `{"status":"ok"}`, string(buf))
} 