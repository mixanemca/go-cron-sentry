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

	"github.com/mixanemca/go-cron-sentry/app"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const (
	defaultMaxLenght int = 4096
)

var (
	cfgFile     string
	versionFlag bool
)

var (
	version string = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                   "go-cron-sentry [flags] -- command [arg ...]",
	Example:               "  go-cron-sentry --quite -- aptly-mirror-update --repo debian --distribution buster",
	Short:                 "Wraps commands and reports those that fail to Sentry",
	DisableFlagsInUseLine: true,
	Run:                   rootCmdRun,
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

	rootCmd.PersistentFlags().BoolP("report-all", "a", false, "report to Sentry even if the task has succeeded (env SENTRY_REPORT_ALL)")
	viper.BindEnv("report-all", "SENTRY_REPORT_ALL")

	rootCmd.PersistentFlags().IntP("string-max-length", "M", defaultMaxLenght, "the maximum characters of a string that should be sent to Sentry (env STRING_MAX_LENGTH)")
	viper.BindEnv("STRING_MAX_LENGTH")

	rootCmd.PersistentFlags().BoolP("not-null-output", "n", false, "not send to Sentry if exit-code {0,1} and stdout/stderr is null (env SENTRY_NOT_NULL_OUTPUT)")
	viper.BindEnv("not-null-output", "SENTRY_NOT_NULL_OUTPUT")

	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "suppress all command output (env SENTRY_QUIET)")
	// viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	viper.BindEnv("quiet", "SENTRY_QUIET")

	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "show program's version number and exit")

	rootCmd.PersistentFlags().BoolP("non-zero-exit", "z", false, "not send to Sentry if exit-code 0 (env SENTRY_NON_ZERO_EXIT)")
	viper.BindEnv("non-zero-exit", "SENTRY_NON_ZERO_EXIT")

	rootCmd.PersistentFlags().StringP("runner", "r", "local", "cron task runner {local} (env SENTRY_RUNNER)")
	viper.BindEnv("runner", "SENTRY_RUNNER")

	// Bind all persistent flags to Viper
	viper.BindPFlags(rootCmd.PersistentFlags())
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
	if len(args) < 1 {
		fmt.Printf("ERROR: you must be specify the command\n\n")
		cmd.Help()
		os.Exit(1)
	}

	dsn := viper.GetString("dsn")
	if dsn == "" {
		fmt.Printf("ERROR: you must be specify the DSN\n\n")
		cmd.Help()
		os.Exit(1)
	}

	a, err := app.New(
		app.WithDSN(dsn),
		app.WithTask(args[0], args[1:]...),
		app.WithRunnerBackend(viper.GetString("runner")),
		app.WithQuiet(viper.GetBool("quiet")),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	out, err := a.Runner().Run(args[0], args[1:]...)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("DSN: %s\n", dsn)
	fmt.Printf("Run: %s\n", strings.Join(args, " "))
	fmt.Print(out)
}
