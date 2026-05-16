package parser

import (
	"encoding/json"
	"os"

	"github.com/sharifulb07/goku/internal/models"
)

// internal/parser/parser.go
func ReadResourceJSON(path string) (*models.Resource, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var res models.Resource
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}