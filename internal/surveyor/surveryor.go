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

package surveyor

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/slugbus/taps/v2"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// Main is the entry point for the surveyor
func Main(cmd *cobra.Command, args []string) error {
	type flags struct {
		interval time.Duration
		duration time.Duration
		url      string
		number   uint64
	}
	var (
		err error
		f   flags
	)

	// Parse the flags
	if f.interval, err = cmd.Flags().GetDuration("interval"); err != nil {
		return err
	}
	if f.duration, err = cmd.Flags().GetDuration("duration"); err != nil {
		return err
	}
	if f.number, err = cmd.Flags().GetUint64("number"); err != nil {
		return err
	}
	if f.url, err = cmd.Flags().GetString("server"); err != nil {
		return err
	}

	// Set the custom url for taps to hit
	s := taps.NewSource(f.url)

	// If a number was specified, ping the server n times.
	if cmd.Flags().Changed("number") {
		pingNTimes(s, f.number, f.interval)
		return nil
	}

	// Otherwise do the duration approach.
	pingDuration(s, f.duration, f.interval)

	return nil
}

func pingDuration(s taps.Source, duration, interval time.Duration) {
	fmt.Println("[")
	// Keep track of the number of times we've pinged.
	count := uint64(0)

	// Run this part of the code in a goroutine.
	go func() {
		buses, err := s.Query()
		if err != nil {
			logrus.Error(err)
		} else {
			// Print the comment
			fmt.Printf("\t// [%v]: Query %04d\n", time.Now().Format(time.UnixDate), count+1)

			// Marshal the json.
			bytes, err := json.MarshalIndent(buses, "", "    ")
			if err != nil {
				logrus.Error(err)
			}
			fmt.Printf("\t%s,\n", string(bytes))
			count++
		}

		for range time.Tick(interval) {
			// Query
			buses, err := s.Query()
			if err != nil {
				logrus.Error(err)
				continue
			}

			// Print the comment
			fmt.Printf("\t// [%v]: Query %04d\n", time.Now().Format(time.UnixDate), count+1)

			// Marshal the json.
			bytes, err := json.MarshalIndent(buses, "", "    ")
			if err != nil {
				logrus.Error(err)
			}
			fmt.Printf("\t%s,\n", string(bytes))
			count++
		}
	}()

	// Sleep this go routine for duration seconds
	time.Sleep(duration)
	fmt.Println("]")
}

func pingNTimes(s taps.Source, n uint64, interval time.Duration) {
	fmt.Println("[")
	// Keep track of the number of times we've pinged.
	count := uint64(0)
	for range time.Tick(interval) {
		// Query
		buses, err := s.Query()
		if err != nil {
			logrus.Error(err)
			continue
		}

		// Print the comment
		fmt.Printf("\t// [%v]: Query %04d\n", time.Now().Format(time.UnixDate), count+1)

		// Marshal the json.
		bytes, err := json.MarshalIndent(buses, "", "    ")
		if err != nil {
			logrus.Error(err)
		}

		// Print the json.
		if count+1 == n {
			fmt.Printf("\t%s\n", string(bytes))
		} else {
			fmt.Printf("\t%s,\n", string(bytes))
		}

		// Increment
		// and Check exit condition.
		count++
		if count >= n {
			break
		}
	}

	fmt.Println("]")
}
