/*
Copyright © 2020 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const (
	defaultMaxLenght int = 4096
)

var (
	cfgFile       string
	report        bool
	maxLenght     int
	notNullOutput bool
	quiet         bool
	versionFlag   bool
	nonZeroExit   bool
)

var (
	version string = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                   "cron-sentry [flags] command [arg ...]",
	Short:                 "Wraps commands and reports those that fail to Sentry",
	DisableFlagsInUseLine: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: rootCmdRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "/etc/cron-sentry.conf", "config file")

	rootCmd.PersistentFlags().StringP("dsn", "d", "", "Sentry DSN (env SENTRY_DSN)")
	viper.BindEnv("dsn", "SENTRY_DSN")

	rootCmd.PersistentFlags().BoolVarP(&report, "report-all", "a", false, "report to Sentry even if the task has succeeded (env SENTRY_REPORT_ALL)")
	viper.BindEnv("report-all", "SENTRY_REPORT_ALL")

	rootCmd.PersistentFlags().IntVarP(&maxLenght, "string-max-length", "M", defaultMaxLenght, "the maximum characters of a string that should be sent to Sentry (env STRING_MAX_LENGTH)")
	viper.BindEnv("STRING_MAX_LENGTH")

	rootCmd.PersistentFlags().BoolVarP(&notNullOutput, "not-null-output", "n", false, "not send to sentry if exit-code {0,1} and stdout/stderr is null (env SENTRY_NOT_NULL_OUTPUT)")
	viper.BindEnv("not-null-output", "SENTRY_NOT_NULL_OUTPUT")

	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress all command output (env SENTRY_QUIET)")
	viper.BindEnv("quiet", "SENTRY_QUIET")

	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "show program's version number and exit")

	rootCmd.PersistentFlags().BoolVarP(&nonZeroExit, "non-zero-exit", "z", false, "not send to Sentry if exit-code 0 (env SENTRY_NON_ZERO_EXIT)")
	viper.BindEnv("non-zero-exit", "SENTRY_NON_ZERO_EXIT")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("env")
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadInConfig()
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	if versionFlag {
		fmt.Println("cron-sentry", version)
		os.Exit(0)
	}
	dsn := viper.GetString("dsn")
	if dsn == "" {
		fmt.Println("dns not set")
		os.Exit(1)
	}
	fmt.Printf("DSN: %s\n", dsn)
	fmt.Printf("Run: %s\n", strings.Join(args, " "))
}
