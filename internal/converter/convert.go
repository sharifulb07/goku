package converter

import (
	"encoding/json"

	"github.com/sharifulb07/goku/internal/models"
	"gopkg.in/yaml.v3"
)

// tojson

func ToJSON(cfg *models.Config)([]byte, error){

	return json.MarshalIndent(cfg, "", " ")

}



// toyaml

func ToYAML(cfg *models.Config) ([]byte, error){

	return yaml.Marshal(cfg)


}

