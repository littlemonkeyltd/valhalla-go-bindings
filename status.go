package valhalla

import (
	"bytes"
	"encoding/json"
)

type StatusResponse struct {
	Version             string `json:"version"`
	TileSetLastModified int    `json:"tileset_last_modified"`
}

func (c *Client) Status() (StatusResponse, error) {

	b := []byte{}
	response, err := c.request("GET", "status", bytes.NewReader(b))
	if err != nil {
		return StatusResponse{}, err
	}

	result := StatusResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return StatusResponse{}, err
	}

	return result, nil
}
