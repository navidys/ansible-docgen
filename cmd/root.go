/*
Copyright © 2025 Navid Yaghoobi <navidys@fedoraproject.org>
*/
package cmd

import (
	"os"
	"time"

	"github.com/navidys/ansible-docgen/pkg/docgen"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	appName    = "ansible-docgen"
	appVersion = "0.1.0-dev"
)

var rootCmd = &cobra.Command{ //nolint:exhaustruct,gochecknoglobals
	Use:   appName,
	Short: "Auto documentation for Ansible roles",
	Long:  `Auto documentation for Ansible roles`,
	Run:   runRoot,
}

var (
	outputFilename = "README.md" //nolint:gochecknoglobals
	roleDirectory  = ""          //nolint:gochecknoglobals
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runRoot(cmd *cobra.Command, args []string) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	// check if debug mode
	runDebug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal().Msgf("failed to get debug flag: %s", err.Error())
	}

	if runDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// check if version display
	displayVersion, err := cmd.Flags().GetBool("version")
	if err != nil {
		log.Fatal().Msgf("failed to get version flag: %s", err.Error())
	}

	if displayVersion {
		log.Info().Msgf("%s v%s", appName, appVersion)

		os.Exit(0)
	}

	doc, err := docgen.NewDocumentGenerator(cmd, args)
	if err != nil {
		log.Error().Msgf("failed to create document generator: %s", err.Error())

		os.Exit(1)
	}

	err = doc.Generate()
	if err != nil {
		log.Error().Msgf("failed to generate document: %s", err.Error())

		os.Exit(1)
	}
}

func init() { //nolint:gochecknoinits
	rootCmd.Flags().StringVarP(
		&outputFilename,
		"output",
		"o",
		outputFilename,
		"output readme file name",
	)
	rootCmd.Flags().StringVarP(
		&roleDirectory,
		"role",
		"r",
		roleDirectory,
		"Ansible role directory path",
	)
	rootCmd.Flags().BoolP("version", "v", false, "display version and exit")
	rootCmd.Flags().BoolP("debug", "d", false, "run in debug mode")
}
