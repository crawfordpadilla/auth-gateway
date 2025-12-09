package authgateway

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func createTLSConfig(keyPath, certPath string) (*tls.Config, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	if keyPath != "" && certPath != "" {
		_, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, err
		}

		tlsConfig.Certificates = []tls.Certificate{*tls.Certificate{Key: keyPath, Cert: certPath}}
	}

	return tlsConfig, nil
}

func createRequest(method, url string, headers map[string]string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if body != nil {
		req.Body = http.NoBody
	}

	return req, nil
}

func createTimeoutRequest(method, url string, headers map[string]string, body interface{}, timeout time.Duration) (*http.Request, error) {
	req, err := createRequest(method, url, headers, body)
	if err != nil {
		return nil, err
	}

	req.Timeout = timeout

	return req, nil
}

func httpGetRequest(url string, headers map[string]string, timeout time.Duration) (*http.Response, error) {
	req, err := createTimeoutRequest("GET", url, headers, nil, timeout)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func httpPostRequest(url string, headers map[string]string, body interface{}, timeout time.Duration) (*http.Response, error) {
	req, err := createTimeoutRequest("POST", url, headers, body, timeout)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func httpPutRequest(url string, headers map[string]string, body interface{}, timeout time.Duration) (*http.Response, error) {
	req, err := createTimeoutRequest("PUT", url, headers, body, timeout)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func httpDeleteRequest(url string, headers map[string]string, timeout time.Duration) (*http.Response, error) {
	req, err := createTimeoutRequest("DELETE", url, headers, nil, timeout)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func httpPatchRequest(url string, headers map[string]string, body interface{}, timeout time.Duration) (*http.Response, error) {
	req, err := createTimeoutRequest("PATCH", url, headers, body, timeout)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}