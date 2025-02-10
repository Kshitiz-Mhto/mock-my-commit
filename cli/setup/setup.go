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

package setup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const (
	APIKEY_STORAGE = ".mock-my-commitrc"
	HOME           = "HOME"
)

var (
	ConfigFile = filepath.Join(os.Getenv(HOME), APIKEY_STORAGE)
	SetupCMD   = &cobra.Command{
		Use:     "setup",
		Aliases: []string{"set"},
		Short:   "Subcommand to setup the API key for the API Client. A propmt will be provided to enter the API key.",
		Example: "mock-my-commit setup",
		Run:     runSetupForTool,
	}
)

func runSetupForTool(cmd *cobra.Command, args []string) {
	SetupAPIKey()
}

func SetupAPIKey() {
	fmt.Print("Enter API key: ")
	key, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		utility.Error("\n❌ Failed to read API key: %s", err)
	}

	apiKey := strings.TrimSpace(string(key))
	if apiKey == "" {
		utility.Error("❌ API key cannot be empty.")
		os.Exit(1)
	}

	utility.Info(string(emoji.FileFolder)+"Key saved to %s", ConfigFile)
	err = os.WriteFile(ConfigFile, []byte(key), 0600)
	if err != nil {
		utility.Error("❌ Error saving API key: %s", err)
		os.Exit(1)
	}

	utility.Success("🔑 API key saved successfully!!")
}
