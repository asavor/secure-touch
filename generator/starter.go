package generator

import (
	"encoding/json"
	"log"
)

type starter struct {
	DeviceId      string `json:"device_id"`
	ClientVersion string `json:"clientVersion"`
	DeviceType    string `json:"deviceType"`
	AuthToken     string `json:"authToken"`
}

func (c *Client) GenerateStarterPayload() (string, error) {

	payload := &starter{
		DeviceId:      c.DeviceID,
		ClientVersion: c.ClientVersion,
		DeviceType:    c.DeviceType,
		AuthToken:     "",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("could not marshal json")
		return "", err
	}

	return string(jsonPayload), nil
}
