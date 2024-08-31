package utilstest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// TestRunner : This is test runner
func TestRunner(method string, url string, body io.Reader) (*http.Request, error) {
	req := httptest.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// ResponseReader : This is response reader
func ResponseReader(resp *http.Response) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("CreateOtpEnpoint() error = %w", err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
