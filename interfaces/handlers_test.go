package interfaces

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestHealthcheckGetHandler(t *testing.T) {
    // Set up Gin router
    router := gin.Default()
    router.GET("/health", HealthcheckGetHandler())

    // Create a test request
    req, _ := http.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "healthy")
}

func TestMessageGetHandler(t *testing.T) {
    // Set up Gin router
    router := gin.Default()
    router.GET("/greetings", MessageGetHandler())

    // Create a test request
    req, _ := http.NewRequest("GET", "/greetings", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Hello")
}
