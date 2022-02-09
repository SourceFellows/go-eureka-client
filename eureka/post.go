package eureka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) RegisterInstance(appId string, instanceInfo *InstanceInfo) error {
	values := []string{"apps", appId}
	path := strings.Join(values, "/")
	instance := &Instance{
		Instance: instanceInfo,
	}
	body, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	response, err := c.Post(path, body)
	if err != nil {
		return fmt.Errorf("could not register instance. %v", err)
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated && response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not register instance. response (%d) %v", response.StatusCode, string(response.Body))
	}

	return nil
}
