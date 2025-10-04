package docgen

import (
	"fmt"

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
		OutputFile:    outputFile,
		RoleDirectory: roleDirectory,
		TemplateFile:  templateFile,
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
	return nil
}
