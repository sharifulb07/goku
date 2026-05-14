package validator

import (
	"errors"
	"strings"
)

func ValidInput(ext, outputFmt string) error {

	outputFmt = strings.ToLower(outputFmt)
	ext = strings.ToLower(ext)

	switch ext {
	case ".json":
		if outputFmt == "json" {
			return errors.New("output format must differ from input format")
		}

	case ".yaml", ".yml":
		if outputFmt == "yaml" {
			return errors.New("output format must differ from input format")
		}
	}

	if outputFmt != "json" && outputFmt != "yaml" {
		return errors.New("supported output formats: json, yaml")
	}

	return nil
}