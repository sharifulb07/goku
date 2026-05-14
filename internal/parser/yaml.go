package parser

import (
	"os"

	"github.com/sharifulb07/goku/internal/models"
	"gopkg.in/yaml.v3"
)

func ReadYAML(path string) (*models.Config, error){

	data, err:=os.ReadFile(path)
	if err !=nil{
		return nil, err
	}

	var cfg models.Config

	err=yaml.Unmarshal(data, &cfg)

	if err!=nil{
		return nil, err 
	}

	return &cfg, nil 
}