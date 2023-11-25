package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WillfJunior/api-go/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/clients", strings.NewReader(`{"name": "John", "age": 25}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, controllers.CreateClient(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Additional assertions as needed...
}
