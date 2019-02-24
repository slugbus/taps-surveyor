//    Copyright 2019 The SlugBus++ Authors.
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package cmd

import (
	"os"
	"time"

	"github.com/slugbus/taps-surveyor/internal/surveyor"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taps-surveyor",
	Short: "Survery the UCSC TAPS API via the command line",
	RunE:  surveyor.Main,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().DurationP("interval", "i", 3*time.Second, "how often to ping the TAPS server")
	rootCmd.Flags().DurationP("duration", "d", 30*time.Second, "how long to ping the TAPS server for")
	rootCmd.Flags().Uint64P("number", "n", 10, "how many times to ping the TAPS server, this flag takes precedence over the duration flag")
	rootCmd.Flags().StringP("server", "s", "http://bts.ucsc.edu:8081/location/get", "specify a custom server to ping")
}
