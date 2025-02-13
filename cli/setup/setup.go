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
	"runtime"
	"strings"

	"github.com/Kshitiz-Mhto/mock-my-commit/pkg/env"
	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	ConfigFile string
	SetupCMD   = &cobra.Command{
		Use:     "setup",
		Aliases: []string{"set"},
		Short:   "Subcommand to setup the API key for the API Client. A propmt will be provided to enter the API key.",
		Example: "mock-my-commit setup",
		Run:     runSetupForTool,
	}
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		utility.Error("❌ Failed to get home directory: %v", err)
	}

	switch runtime.GOOS {
	case env.LINUX_OS:
		ConfigFile = filepath.Join(homeDir, env.APIKEY_STORAGE)
	case env.WINDOWS_OS:
		appDataDir := filepath.Join(homeDir, env.APP_DATA_DIR, env.ROAMING_DIR)
		ConfigFile = filepath.Join(appDataDir, env.APIKEY_STORAGE)

		if err := os.MkdirAll(appDataDir, 0755); err != nil {
			utility.Error("\n❌ Failed to create AppData directory: %v", err)
			os.Exit(1)
		}
	default:
		clickableURL := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", env.GITHUB_REPO_ISSUE_LIST, env.GITHUB_REPO_ISSUE_LIST)

		utility.Error("Unsupported platform/OS. please raise a feature request in our repository [%s]. Thank you!", clickableURL)

		if err := utility.OpenInBrowser(env.GITHUB_REPO_ISSUE_LIST); err != nil {
			utility.Info("Please manually visit: %s", env.GITHUB_REPO_ISSUE_LIST)
		}
	}
}

func runSetupForTool(cmd *cobra.Command, args []string) {
	switch runtime.GOOS {
	case env.LINUX_OS, env.MAC_OS:
		SetupAPIKeyForLinuxAndUnix()
	case env.WINDOWS_OS:
		SetupAPIKeyForWindows()
	default:
		clickableURL := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", env.GITHUB_REPO_ISSUE_LIST, env.GITHUB_REPO_ISSUE_LIST)

		utility.Error("Unsupported platform/OS. please raise a feature request in our repository [%s]. Thank you!", clickableURL)

		if err := utility.OpenInBrowser(env.GITHUB_REPO_ISSUE_LIST); err != nil {
			utility.Info("Please manually visit: %s", env.GITHUB_REPO_ISSUE_LIST)
		}
	}
}

func SetupAPIKeyForLinuxAndUnix() {
	fmt.Print("Enter API key: ")
	key, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		utility.Error("\n❌ Failed to read API key: %s", err)
		os.Exit(1)
	}

	apiKey := strings.TrimSpace(string(key))
	if apiKey == "" {
		utility.Error("❌ API key cannot be empty.")
		os.Exit(1)
	}

	if len(apiKey) != 32 {
		utility.Error("❌ Invalid key format for Mistral API")
		os.Exit(1)
	}

	fmt.Println()
	utility.Info(string(emoji.FileFolder)+"Key saved to %s", ConfigFile)
	err = os.WriteFile(ConfigFile, []byte(key), 0600)
	if err != nil {
		utility.Error("❌ Error saving API key: %s", err)
		os.Exit(1)
	}

	utility.Success("🔑 API key saved successfully!!")
}

func SetupAPIKeyForWindows() {
	fmt.Print("Enter API key: ")
	key, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		utility.Error("\n❌ Failed to read API key: %s", err)
		os.Exit(1)
	}

	apiKey := strings.TrimSpace(string(key))
	if apiKey == "" {
		utility.Error("❌ API key cannot be empty.")
		os.Exit(1)
	}

	if len(apiKey) != 32 {
		utility.Error("❌ Invalid key format for Mistral API")
		os.Exit(1)
	}

	fmt.Println()
	utility.Info(string(emoji.FileFolder)+"Key saved to %s", ConfigFile)
	err = os.WriteFile(ConfigFile, []byte(key), 0600)
	if err != nil {
		utility.Error("❌ Error saving API key: %s", err)
		os.Exit(1)
	}

	utility.Success("🔑 API key saved successfully!!")
}
