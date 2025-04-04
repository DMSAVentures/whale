package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var serverAddr = "http://localhost:8080"

func TestEchoEndpoint(t *testing.T) {
	t.Run("GET /echo successfully echoes matrix", func(t *testing.T) {
		filePath := "../matrix.csv"
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", filepath.Base(filePath))
		assert.NoError(t, err)

		file, err := os.Open(filePath)
		assert.NoError(t, err)
		defer file.Close()

		_, err = io.Copy(part, file)
		assert.NoError(t, err)
		writer.Close()

		req, err := http.NewRequest("GET", serverAddr+"/echo", body)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)

		assert.Equal(t, "1,2,3\n4,5,6\n7,8,9\n\n", string(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /echo responds with 400 when csv is invalid", func(t *testing.T) {

		req, err := http.NewRequest("GET", serverAddr+"/echo", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "multipart/form-data")

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func createMultipartRequest(t *testing.T, method, url, filePath string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	assert.NoError(t, err)

	file, err := os.Open(filePath)
	assert.NoError(t, err)
	t.Cleanup(func() { file.Close() })

	_, err = io.Copy(part, file)
	assert.NoError(t, err)
	writer.Close()

	req, err := http.NewRequest(method, url, body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

func TestMatrixOperations(t *testing.T) {
	client := &http.Client{}
	filePath := "../matrix.csv"

	t.Run("GET /invert successfully inverts numeric matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/invert", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "1,4,7\n2,5,8\n3,6,9\n\n", string(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /invert responds with 400 on missing file", func(t *testing.T) {
		req, err := http.NewRequest("GET", serverAddr+"/invert", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "multipart/form-data")

		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("GET /sum successfully sums numeric matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/sum", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "45\n", string(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /multiply successfully multiplies numeric matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/multiply", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "362880\n", string(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /flatten successfully flattens numeric matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/flatten", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "1,2,3,4,5,6,7,8,9\n", string(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	// Additional edge case
	t.Run("GET /multiply returns overflow error on large numeric matrix", func(t *testing.T) {
		overflowPath := "../overflowMatrix.csv" // matrix with large values
		req := createMultipartRequest(t, "GET", serverAddr+"/multiply", overflowPath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)

		assert.Equal(t, "failed to process request: unsupported operation\n", string(respBody))
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("GET /sum returns overflow error on large numeric matrix", func(t *testing.T) {
		overflowPath := "../overflowMatrix.csv" // matrix with large values
		req := createMultipartRequest(t, "GET", serverAddr+"/sum", overflowPath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "failed to process request: unsupported operation\n", string(respBody))
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}

func TestStringMatrixOperations(t *testing.T) {
	client := &http.Client{}
	filePath := "../stringMatrix.csv"

	t.Run("GET /invert successfully inverts string matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/invert", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "a,e,h\nb,f,i\nc,g,j\n\n", string(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /invert responds with 400 on missing file", func(t *testing.T) {
		req, err := http.NewRequest("GET", serverAddr+"/invert", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "multipart/form-data")

		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("GET /sum returns error on string matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/sum", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "failed to process request: unsupported operation\n", string(respBody))
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("GET /multiply returns error on string matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/multiply", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "failed to process request: unsupported operation\n", string(respBody))
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("GET /flatten successfully flattens string matrix", func(t *testing.T) {
		req := createMultipartRequest(t, "GET", serverAddr+"/flatten", filePath)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "a,b,c,e,f,g,h,i,j\n", string(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
