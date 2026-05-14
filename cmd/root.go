package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sharifulb07/goku/internal/converter"
	"github.com/sharifulb07/goku/internal/parser"
	"github.com/sharifulb07/goku/internal/utils"
	"github.com/sharifulb07/goku/internal/validator"
	"github.com/sharifulb07/goku/internal/writer"

	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFmt  string // ← for json/yaml
	outputFile string // ← custom filename
)

var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "JSON ↔ YAML converter CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate
		err := validator.ValidInput(filepath.Ext(inputFile), outputFmt)
		if err != nil {
			return err
		}

		var data []byte

		switch filepath.Ext(inputFile) {
		case ".json":
			cfg, err := parser.ReadJSON(inputFile)
			if err != nil {
				return err
			}
			data, err = converter.ToYAML(cfg)
			if err != nil {
				return err
			}

		case ".yaml", ".yml":
			cfg, err := parser.ReadYAML(inputFile)
			if err != nil {
				return err
			}
			data, err = converter.ToJSON(cfg)
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("unsupported input format")
		}

		// Generate output filename if not provided
		if outputFile == "" {
			outputFile, err = utils.GenerateOutputFileName(inputFile, outputFmt)
			if err != nil {
				return err
			}
		}

		if err := writer.WriteFile(outputFile, data); err != nil {
			return err
		}

		fmt.Printf("✅ Success! Generated: %s\n", outputFile)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	rootCmd.Flags().StringVarP(&outputFmt, "format", "o", "", "Output format (json or yaml)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "f", "", "Custom output file name (optional)")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("format")
}