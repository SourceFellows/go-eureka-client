package eureka

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (c *Client) SendHeartbeat(appId, instanceId string) error {
	values := []string{"apps", appId, instanceId}
	path := strings.Join(values, "/")
	resp, err := c.Put(path, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusNotFound:
		return newError(ErrCodeInstanceNotFound,
			"Instance resource not found when sending heartbeat", 0)
	}

	logrus.WithField("status", resp.StatusCode).
		WithField("body", string(resp.Body)).
		Trace("heartbeat sent")

	return nil
}
