package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"./app"
	"./base/configuration"
	"github.com/spf13/cobra"
)

var (
	appName      = "inventory"
	majorVersion = "0"
	minorVersion = "1.beta"
	revision     = ""

	rootDir      = "."
	secureConfig = false

	rootCmd *cobra.Command
)

func init() {
	var appName string
	var config configuration.Config

	// init root command
	rootCmd = &cobra.Command{
		Use:   appName,
		Short: "Inventory",
		Long:  "A testing inventory",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var dir string
			if appName != filepath.Base(os.Args[0]) {
				dir = filepath.Dir(appName)
				appName = strings.TrimSuffix(appName, filepath.Ext(appName))
			}

			config, err = configuration.New(appName, rootCmd, dir)
			if err != nil {
				log.Println(err)
			}
			// override by command line
			config.Update()

			app.Start(config)
		},
	}

	rootCmd.PersistentFlags().StringVar(&rootDir, "root-dir", ".", "App root directory")
	rootCmd.PersistentFlags().StringVar(&appName, "config-file", filepath.Base(os.Args[0]), "Set json config file to use")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print program version and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s v.%s.%s-%s\n", appName, majorVersion, minorVersion, revision)
			os.Exit(0)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "config",
		Short: "Show available config",
		Run: func(cmd *cobra.Command, args []string) {
			config, _ := configuration.New("", rootCmd)
			config.Dump(os.Stderr)
			os.Exit(0)
		},
	})

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
