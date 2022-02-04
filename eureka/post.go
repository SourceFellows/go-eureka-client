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
		return nil
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("could not register instance. response %v", string(response.Body))
	}

	return nil
}
