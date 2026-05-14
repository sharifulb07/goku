package parser

import (
	"encoding/json"
	"os"

	"github.com/sharifulb07/goku/internal/models"
)

func ReadJSON(path string) (*models.Config, error){

	data, err:=os.ReadFile(path)

	if err !=nil{
		return nil, err 
	}

	var cfg models.Config

	err=json.Unmarshal(data, &cfg )

	if err !=nil{
		return nil, err 
	}

	return &cfg, nil 
}