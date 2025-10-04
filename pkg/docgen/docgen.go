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
	outputFilename, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, fmt.Errorf("failed to get output flag: %w", err)
	}

	log.Debug().Msgf("role directory path: %s", roleDirectory)
	log.Debug().Msgf("output filename: %s", outputFilename)

	docg := DocumentGenerator{
		OutputFilename: outputFilename,
		RoleDirectory:  roleDirectory,
	}

	return &docg, nil
}

// Generate generates roles documents.
func (d *DocumentGenerator) Generate() error {
	return nil
}
