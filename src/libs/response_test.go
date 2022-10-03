package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		assert.Equal(t, "OK", getStatus(200), "result must be OK")
	})

	t.Run("Status Created", func(t *testing.T) {
		assert.Equal(t, "Created", getStatus(201), "result must be Created")
	})

	t.Run("Status Accepted", func(t *testing.T) {
		assert.Equal(t, "Accepted", getStatus(202), "result must be Accepted")
	})

	t.Run("Status Not Modified", func(t *testing.T) {
		assert.Equal(t, "Not Modified", getStatus(304), "result must be Not Modified")
	})

	t.Run("Status Bad Request", func(t *testing.T) {
		assert.Equal(t, "Bad Request", getStatus(400), "result must be Bad Request")
	})

	t.Run("Status Unauthorized", func(t *testing.T) {
		assert.Equal(t, "Unauthorized", getStatus(401), "result must be Unauthorized")
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		assert.Equal(t, "Forbidden", getStatus(403), "result must be Forbidden")
	})

	t.Run("Status Not Found", func(t *testing.T) {
		assert.Equal(t, "Not Found", getStatus(404), "result must be Not Found")
	})

	t.Run("Status Unsupported Media Type", func(t *testing.T) {
		assert.Equal(t, "Unsupported Media Type", getStatus(415), "result must be Unsupported Media Type")
	})

	t.Run("Status Internal Server Error", func(t *testing.T) {
		assert.Equal(t, "Internal Server Error", getStatus(500), "result must be Internal Server Error")
	})

	t.Run("Status Bad Gateway", func(t *testing.T) {
		assert.Equal(t, "Bad Gateway", getStatus(502), "result must be Unsupported Bad Gateway")
	})

	t.Run("Status Code Undefined", func(t *testing.T) {
		assert.Equal(t, "Status Code Undefined", getStatus(666), "result must be Status Code Undefined")
	})
}

func TestGetResponse(t *testing.T) {
	t.Run("Get Response OK", func(t *testing.T) {
		code := 200
		isError := false
		data := "dummy"
		result := &Response{
			Code:    code,
			Status:  getStatus(code),
			IsError: isError,
			Data:    data,
		}
		assert.Equal(t, result, GetResponse(data, code, isError), "Unexpected response standard")
	})

	t.Run("Get Response Error", func(t *testing.T) {
		code := 500
		isError := true
		data := "error description"
		result := &Response{
			Code:        code,
			Status:      getStatus(code),
			IsError:     isError,
			Description: data,
		}
		assert.Equal(t, result, GetResponse(data, code, isError), "Unexpected response standard")
	})
}
