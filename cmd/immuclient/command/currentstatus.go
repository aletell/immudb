/*
Copyright 2019-2020 vChain, Inc.

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

package immuclient

import (
	"fmt"

	c "github.com/codenotary/immudb/cmd/helper"
	"github.com/spf13/cobra"
)

func (cl *commandline) currentRoot(cmd *cobra.Command) {
	ccmd := &cobra.Command{
		Use:               "current",
		Short:             "Return the last merkle tree root and index stored locally",
		Aliases:           []string{"crt"},
		PersistentPreRunE: cl.connect,
		PersistentPostRun: cl.disconnect,
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := cl.immucl.CurrentRoot(args)
			if err != nil {
				c.QuitToStdErr(err)
			}
			fmt.Println(resp)
			return nil
		},
		Args: cobra.ExactArgs(0),
	}
	cmd.AddCommand(ccmd)
}
