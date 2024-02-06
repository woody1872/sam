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
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/woody1872/sam/checksum"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the checksum of a file",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		hashToVerify := args[0]
		fileToHash := args[1]

		f, err := os.Open(fileToHash)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		h, err := checksum.NewHashAlgorithm(cmd.Flag(algorithmFlag).Value.String())
		if _, err := io.Copy(h, f); err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		fileChecksum := hex.EncodeToString(h.Sum(nil))
		if hashToVerify != fileChecksum {
			err = fmt.Errorf("the checksum does not match: \"%s\" != \"%s\"", hashToVerify, fileChecksum)
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("success: the checksum is verified")
	},
}

func init() {
	checksumCmd.AddCommand(verifyCmd)
}
