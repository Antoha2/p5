package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	HTTPClient HttpClient
}

type ResponseData struct {
	Message string `json:"message"`
}

func (c *Client) FetchData(url string) (*ResponseData, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestFetchData(t *testing.T) {
	mockClient := new(MockHttpClient)

	responseData := ResponseData{Message: "Hello, World!"}
	responseBody, _ := json.Marshal(responseData)
	response := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(responseBody)),
		Header:     make(http.Header),
	}

	mockClient.On("Do", mock.Anything).Return(response, nil)

	client := &Client{
		HTTPClient: mockClient,
	}

	data, err := client.FetchData("http://127.0.0.1")

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "Hello, World!", data.Message)

	mockClient.AssertExpectations(t)
}
