package utility

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gookit/color"
)

const (
	CLI_BINARY_PATH = "bin/mock-my-commit"
)

// OpenInBrowser attempts to open the specified URL in the default browser.
func OpenInBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		return fmt.Errorf("failed to open URL: %w", err)
	}
	return nil
}

// GetBuildDate is the function to get build stats of binary
func GetBuildDate() string {
	filePath := CLI_BINARY_PATH
	info, err := os.Stat(filePath)
	if err != nil {
		return ""
	}
	return info.ModTime().Format(time.RFC1123)
}

// Error is the function to handler all error in the Cli
func Error(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", color.Red.Sprintf("Error"), fmt.Sprintf(msg, args...))
}

// Info is the function to handler all info messages in the Cli
func Info(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", color.Blue.Sprintf("Info"), fmt.Sprintf(msg, args...))
}

// Warning is the function to handler all warnings in the Cli
func Warning(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", color.Yellow.Sprintf("Warning"), fmt.Sprintf(msg, args...))
}

// Success is the fucntion to handler all the success in Cli
func Success(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", color.Green.Sprintf("sucess"), fmt.Sprintf(msg, args...))
}

// Green is the function to convert str to green in console
func Green(value string) string {
	newColor := color.FgGreen.Render
	return newColor(value)
}

// Yellow is the function to convert str to yellow in console
func Yellow(value string) string {
	newColor := color.New(color.FgYellow).Render
	return newColor(value)
}

// Red is the function to convert str to red in console
func Red(value string) string {
	newColor := color.New(color.FgRed).Render
	return newColor(value)
}
