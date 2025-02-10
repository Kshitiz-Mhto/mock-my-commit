package hooks

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/spf13/cobra"
)

const (
	hookContent = `#!/bin/sh
exec mock-my-commit run-hook "$1" || exit 1`
)

var (
	local           bool
	global          bool
	hookPath        string
	InstallHooksCmd = &cobra.Command{
		Use:   "install",
		Short: "Subcommand to install hooks for commit messages either local or global level.",
		Run:   runInstallHooksSetupCmd,
	}
)

func runInstallHooksSetupCmd(cmd *cobra.Command, args []string) {
	if local && global {
		utility.Error("❌ Cannot specify both --local and --global flags.")
		os.Exit(1)
	}

	switch {
	case local:
		SetupLocalHooks(local)
	case global:
		SetupGlobalHooks(global)
	default:
		utility.Error("❌ Please specify either --local or --global.")
		os.Exit(1)
	}
}

func init() {
	InstallHooksCmd.Flags().BoolVarP(&local, "local", "", false, "Indicates the  hook setup/installation locally, specific to a repository's .git repository.")
	InstallHooksCmd.Flags().BoolVarP(&global, "global", "", false, "Indicates the hook setup/installation globally, specific to parent .git repository.")
}

func SetupLocalHooks(local bool) {
	//get the local .git directory.
	gitDir, err := exec.Command("git", "rev-parse", "--git-dir").Output()
	if err != nil {
		utility.Error("❌ Not a git repository or unable to determine .git directory: %v", err)
		os.Exit(1)
	}
	hookPath = filepath.Join(strings.TrimSpace(string(gitDir)), "hooks")

	hookFile := filepath.Join(hookPath, "commit-msg")
	if err := os.WriteFile(hookFile, []byte(hookContent), 0755); err != nil {
		utility.Error("❌ Failed to install hooks: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(hookFile, 0755); err != nil {
		utility.Error("❌ Failed to set executable permissions on commit hook: %v", err)
		os.Exit(1)
	}

	utility.Success("✅ Hooks installed succesfully!!")
}

func SetupGlobalHooks(global bool) {

	home, err := os.UserHomeDir()
	if err != nil {
		utility.Error("❌ Unable to determine home directory: %v", err)
		os.Exit(1)
	}
	hookPath = filepath.Join(home, ".mock-my-commit-hooks")
	if err := os.MkdirAll(hookPath, 0755); err != nil {
		utility.Error("❌ Failed to create hooks directory: %v", err)
		os.Exit(1)
	}
	// Set the global Git hooks path
	if err := exec.Command("git", "config", "--global", "core.hooksPath", hookPath).Run(); err != nil {
		utility.Error("❌ Failed to configure global hooks path: %v", err)
		os.Exit(1)
	}

	hookFile := filepath.Join(hookPath, "commit-msg")
	if err := os.WriteFile(hookFile, []byte(hookContent), 0755); err != nil {
		utility.Error("❌ Failed to install hooks: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(hookFile, 0755); err != nil {
		utility.Error("❌ Failed to set executable permissions on commit hook: %v", err)
		os.Exit(1)
	}

	utility.Success("✅ Hooks installed succesfully!!")
}
