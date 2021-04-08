package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ParimAPIURL = "https://api.parin.io"
)

type ParinClient struct {
	client *http.Client
	apiKey string
}

func NewParinClient(apiKey string, httpClient *http.Client) *ParinClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &ParinClient{
		client: httpClient,
		apiKey: apiKey,
	}
}

type ParinSensorData struct {
	Temperature float64   `json:"temperature,string"`
	Timestamp   time.Time `json:"timestamp"`
}

func (c *ParinClient) ListSensorData(sensorID string) ([]ParinSensorData, error) {
	//Create request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/sensor/%v", ParimAPIURL, sensorID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-api-key", c.apiKey)

	//Do request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Check response
	err = c.checkResponse(resp)
	if err != nil {
		return nil, err
	}

	//Unmarshal return
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	v := []ParinSensorData{}
	err = json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (c *ParinClient) checkResponse(r *http.Response) error {
	if code := r.StatusCode; code >= 200 && code <= 299 {
		return nil
	}

	type jsonAPIError struct {
		Message string `json:"message"`
	}
	var apiErr jsonAPIError
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &apiErr)
	if err != nil {
		return err
	}

	return &HTTPError{Message: apiErr.Message, StatusCode: r.StatusCode}
}

// HTTPError is an error message that wraps the returned message by the API with the StatusCode of the http response
type HTTPError struct {
	Message    string
	StatusCode int
}

// HTTPError Error formats the HTTPError in a string format
func (m *HTTPError) Error() string {
	return fmt.Sprintf("%v (Status Code %v)", m.Message, m.StatusCode)
}
