package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GenerateOutputFileName(inputFile, outputfmt string) (string, error) {

	ext := filepath.Ext(inputFile)
	base := filepath.Base(inputFile)

	base=strings.TrimSuffix(base, ext)


	switch outputfmt {
	case "yaml":
		return base +".yaml", nil 
	case "json":
		return base+".json", nil 
	
	default:
		return "", fmt.Errorf("unsupported output format. ")
		
	}

}