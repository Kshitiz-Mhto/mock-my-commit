/*
Copyright © 2025 Kshitiz Mhto <kshitizmhto101@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cli

import (
	"fmt"
	"os"

	"github.com/Kshitiz-Mhto/mock-my-commit/cli/hooks"
	"github.com/Kshitiz-Mhto/mock-my-commit/cli/run"
	"github.com/Kshitiz-Mhto/mock-my-commit/cli/setup"
	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	version bool
	rootCmd = &cobra.Command{
		Use:   "mock-my-commit",
		Short: "An AI powered tool that validates your commit message and generate short one liner sarcastic review.",
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				versionCMD.Run(cmd, args)
			} else {
				fmt.Print(utility.Yellow(logo))
				cmd.Help()
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCMD)
	rootCmd.AddCommand(setup.SetupCMD)
	rootCmd.AddCommand(hooks.InstallHooksCmd)
	rootCmd.AddCommand(run.ExecuteHookCmd)

	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Provides version of the tool.")
}
