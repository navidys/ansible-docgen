package docgen

import (
	"fmt"
	"os"
	"text/template"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewDocumentGenerator creates a new document generator.
func NewDocumentGenerator(cmd *cobra.Command, args []string) (*DocumentGenerator, error) {
	// role directory path
	roleDirectory, err := cmd.Flags().GetString("role")
	if err != nil {
		return nil, fmt.Errorf("failed to get role flag: %w", err)
	}

	// output filename
	outputFile, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, fmt.Errorf("failed to get output flag: %w", err)
	}

	// template filename
	templateFile, err := cmd.Flags().GetString("template")
	if err != nil {
		return nil, fmt.Errorf("failed to get template flag: %w", err)
	}

	log.Debug().Msgf("role directory: %s", roleDirectory)
	log.Debug().Msgf("output file: %s", outputFile)
	log.Debug().Msgf("template file: %s", templateFile)

	docg := DocumentGenerator{
		outputFile:    outputFile,
		roleDirectory: roleDirectory,
		templateFile:  templateFile,
	}

	return &docg, nil
}

// Generate generates roles documents.
func (d *DocumentGenerator) Generate() error {
	err := d.parseMetaDirectory()
	if err != nil {
		return err
	}

	err = d.write()
	if err != nil {
		return err
	}

	return nil
}

func (d *DocumentGenerator) write() error {
	tmpl, err := template.ParseFiles(d.templateFile)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	output, err := os.Create(d.outputFile)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}

	defer func() {
		err := output.Close()
		if err != nil {
			log.Error().Msgf("output file close: %s", err.Error())
		}
	}()

	err = tmpl.Execute(output, d)
	if err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}
