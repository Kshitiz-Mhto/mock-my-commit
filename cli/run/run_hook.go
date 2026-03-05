package run

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/Kshitiz-Mhto/mock-my-commit/cli/setup"
	"github.com/Kshitiz-Mhto/mock-my-commit/pkg/env"
	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
	"google.golang.org/api/option"
)

var (
	ExecuteHookCmd = &cobra.Command{
		Use:     "run-hook",
		Aliases: []string{"exec-hook"},
		Short:   "Subcommand that runs the commit hook to process the commit message.",
		Run:     runHookExecutionCmd,
	}
	commitMsgRegex  = regexp.MustCompile(`^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([a-zA-Z0-9_\-]+\))?: [a-z].+`)
	blockedMessages = map[string]bool{
		"update":  true,
		"fix":     true,
		"changes": true,
		"misc":    true,
		"test":    true,
	}
)

func runHookExecutionCmd(cmd *cobra.Command, args []string) {
	switch runtime.GOOS {
	case env.LINUX_OS, env.MAC_OS, env.WINDOWS_OS:
		RunHook()
	default:
		clickableURL := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", env.GITHUB_REPO_ISSUE_LIST, env.GITHUB_REPO_ISSUE_LIST)

		utility.Error("Unsupported platform/OS. please raise a feature request in our repository [%s]. Thank you!", clickableURL)

		if err := utility.OpenInBrowser(env.GITHUB_REPO_ISSUE_LIST); err != nil {
			utility.Info("Please manually visit: %s", env.GITHUB_REPO_ISSUE_LIST)
		}
	}
}

func RunHook() {
	// get commit message
	msg, err := GetCommitMessage()
	if err != nil {
		utility.Error("❌ %s", err)
		os.Exit(1)
	}

	// validate the commit message.
	if ShouldBlockCommit(msg) {
		roast := GenerateRoast(msg, env.PROMPT_STRUCTURE)
		fmt.Println(roast)
		os.Exit(1)
	}

	// // validate the meaningless commit message
	// if ShouldBlockContent(msg) {
	// 	roast := GenerateRoast(msg, env.PROMPT_CONTENT)
	// 	fmt.Println(roast)
	// 	os.Exit(1)
	// }

	os.Exit(0)
}

func GetCommitMessage() (string, error) {
	if len(os.Args) < 3 {
		return "", fmt.Errorf("commit message file not provided")
	}

	filename := os.Args[2]
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read commit message file: %w", err)
	}
	return string(content), nil
}

func ShouldBlockCommit(msg string) bool {
	if len(msg) < env.COMMIT_MSG_LENTH_MIN || len(msg) > env.COMMIT_MSG_LENTH_MAX {
		return true
	}

	if !commitMsgRegex.MatchString(msg) {
		return true
	}

	return false
}

func ShouldBlockContent(msg string) bool {
	return !IsMessageMeaningful(msg)
}

func GetAPIKey() string {
	_ = godotenv.Load()
	if key := os.Getenv(env.GEMINI_API_KEY_ENV); key != "" {
		return key
	}

	key, err := os.ReadFile(setup.ConfigFile)
	if err != nil {
		utility.Error("❌ Error reading API key: %v", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(key))
}

func IsMessageMeaningful(msg string) bool {
	apiKey := GetAPIKey()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		utility.Error("❌ Error creating Gemini client: %v", err)
		return false
	}
	defer client.Close()

	model := client.GenerativeModel(env.GEMINI_MODEL)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(env.PROMPT_CHECK_QUALITY)},
	}

	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("Commit message: %s", msg)))
	if err != nil {
		utility.Error("❌ Error checking message quality: %v", err)
		return false
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return false
	}

	part := resp.Candidates[0].Content.Parts[0]
	response := strings.ToUpper(strings.TrimSpace(fmt.Sprintf("%v", part)))
	return strings.HasPrefix(response, "YES")
}

func GenerateRoast(msg, prompt string) string {
	apiKey := GetAPIKey()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return getFallbackRoast()
	}
	defer client.Close()

	model := client.GenerativeModel(env.GEMINI_MODEL)
	model.SetTemperature(1.0)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(fmt.Sprintf("%s The response must be strictly under %d words.", prompt, env.ROAST_WORD_LIMIT))},
	}

	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("Commit message: %s", msg)))
	if err != nil {
		fmt.Println(err.Error())
		return getFallbackRoast()
	}

	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil && len(resp.Candidates[0].Content.Parts) > 0 {
		return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	}

	return getFallbackRoast()
}

func getFallbackRoast() string {

	fallbackRoastArray := []string{
		"Oh wow, another 'shit'. I'm sure this one won't educate anyone. 🤡",
		"Commit messages like these make us question our civilization. 🏺",
		"Oh good, another commit titled ‘Oops’—confidence inspiring. 👏",
		"This commit is so ugly even Git refuses to merge it. 🤢",
		"If your code were a novel, this commit would be the plot hole. 📖",
		"Your commit history is a crime scene, and this one is the murder weapon. 🔪",
		"Ah, another commit where we pretend everything is fine. 🔥🐶",
		"This commit looks like it was written under duress. 😰",
		"I see you’ve embraced the ‘Commit first, think later’ methodology. 🎭",
		"The best part of this commit is that it can be reverted. 🔄",
		"Your Git history reads like a bad soap opera. 🎭",
		"Ah, the legendary commit message. Until next time. ⏳",
		"If your commits had a theme song, it would be a sad violin. 🎻",
		"This commit deserves its own horror movie. 🎬",
	}

	rand.Seed(uint64(time.Now().UnixNano()))

	return fallbackRoastArray[rand.Intn(len(fallbackRoastArray))]
}
