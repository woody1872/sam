/*
Copyright © 2024 Sam Wood <samwooddev@gmail.com>

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

	"github.com/spf13/cobra"
	"github.com/woody1872/sam/checksum"
)

var algorithmFlag = "algorithm"
var algorithm string

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:     "checksum",
	Aliases: []string{"sum"},
	Short:   "Checksum related tools",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		alg := cmd.Flag(algorithmFlag).Value.String()
		if _, err := checksum.NewHashAlgorithm(alg); err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.PersistentFlags().StringVarP(&algorithm, algorithmFlag, "a", "sha256", "change the hash algorithm")
}
