package collectors

import (
	"firebrick-exporter/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CollectViaHTTP(device *config.DeviceConfig, endpoint string) ([]byte, error) {
	// create Client with 3 Second Timeout
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	response, err := httpClient.Get(fmt.Sprintf("http://%s:%s@%s%s", device.Username, device.Password, device.Address, endpoint))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
