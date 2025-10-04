/*
Copyright © 2025 Navid Yaghoobi <navidys@fedoraproject.org>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

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
	RunE:  runRoot,
}

var (
	outputDirectory       = "." //nolint:gochecknoglobals
	roleDirectory         = ""  //nolint:gochecknoglobals
	errEmptyRoleDirectory = errors.New("empty role directory name")
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runRoot(cmd *cobra.Command, args []string) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	// check if debug mode
	runDebug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return fmt.Errorf("failed to get debug flag: %w", err)
	}

	if runDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// check if version display
	displayVersion, err := cmd.Flags().GetBool("version")
	if err != nil {
		return fmt.Errorf("failed to get version flag: %w", err)
	}

	if displayVersion {
		log.Info().Msgf("%s v%s", appName, appVersion)

		return nil
	}

	// role directory
	if roleDirectory == "" {
		return errEmptyRoleDirectory
	}

	return nil
}

func init() { //nolint:gochecknoinits
	rootCmd.Flags().StringVarP(&outputDirectory, "output", "o", outputDirectory, "Output directory")
	rootCmd.Flags().StringVarP(
		&roleDirectory,
		"role",
		"r",
		roleDirectory,
		"Specifies the directory path to the Ansible role.",
	)
	rootCmd.Flags().BoolP("version", "v", false, "Display version and exit")
	rootCmd.Flags().BoolP("debug", "d", false, "Run in debug mode")
}
