package run

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Kshitiz-Mhto/mock-my-commit/cli/setup"
	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/enescakir/emoji"
	"github.com/gage-technologies/mistral-go"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
)

const (
	PATTERN                      = "^(feat|fix|docs|style|refactor|test|feature|chore|fixes|ci|perf): .+"
	PROMPT_STRUCTURE             = "You are a grumpy and frustrated senior developer. Roast bad git commit messages brutally in short only one line only. Use sarcastic passive-aggressive manner with different opening word and use appropriate emojis."
	PROMPT_CONTENT               = "You are a grumpy senior developer. Roast  meaningless, unclear, or gibberish commit messages brutally in short only one line only. Use sarcastic passive-aggressive manner with different unique word and use appropriate emojis."
	PROMPT_CHECK_QUALITY         = "Evaluate the git commit message. If it clearly describes a git commit message, respond with 'YES'. Otherwise, respond with 'NO'."
	COMMIT_MSG_LENTH_MIN         = 20
	COMMIT_MSG_LENTH_MAX         = 70
	MISTRAL_lARGE_MODEL_VERSION  = "mistral-large-latest"
	MISTRAL_MEDIUM_MODEL_VERSION = "mistral-medium-latest"
	MISTRAL_SMALL_MODEL_VERSION  = "mistral-small-latest"
	SYS_ROLE                     = "system"
	USER_ROLE                    = "user"
)

var (
	ExecuteHookCmd = &cobra.Command{
		Use:     "run-hook",
		Aliases: []string{"exec-hook"},
		Short:   "Subcommand that runs the commit hook to process the commit message.",
		Run:     runHookExecutionCmd,
	}
	commitMsgRegex = regexp.MustCompile(PATTERN)
)

func runHookExecutionCmd(cmd *cobra.Command, args []string) {
	RunHook()
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
		roast := GenerateRoast(msg, PROMPT_STRUCTURE)
		fmt.Println(roast)
		os.Exit(1)
	}

	// validate the meaningless commit message
	if ShouldBlockContent(msg) {
		roast := GenerateRoast(msg, PROMPT_CONTENT)
		fmt.Println(roast)
		os.Exit(1)
	}

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
	if len(msg) < COMMIT_MSG_LENTH_MIN || len(msg) > COMMIT_MSG_LENTH_MAX {
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
	key, err := os.ReadFile(setup.ConfigFile)
	if err != nil {
		utility.Error("❌ Error reading API key: %v", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(key))
}

func IsMessageMeaningful(msg string) bool {
	apiKey := GetAPIKey()
	client := mistral.NewMistralClientDefault(apiKey)

	chatRes, err := client.Chat(
		MISTRAL_lARGE_MODEL_VERSION,
		[]mistral.ChatMessage{
			{
				Role:    SYS_ROLE,
				Content: PROMPT_CHECK_QUALITY,
			},
			{
				Role:    USER_ROLE,
				Content: fmt.Sprintf("Commit message: %s'", msg),
			},
		},
		nil,
	)

	if err != nil {
		utility.Error("❌ Error checking message quality: %v", err)
		return false
	}

	if len(chatRes.Choices) == 0 {
		return false
	}
	response := strings.ToUpper(strings.TrimSpace(chatRes.Choices[0].Message.Content))
	return strings.HasPrefix(response, "YES")
}

func GenerateRoast(msg, prompt string) string {
	apiKey := GetAPIKey()

	client := mistral.NewMistralClientDefault(apiKey)
	chatRes, err := client.Chat(
		MISTRAL_lARGE_MODEL_VERSION, []mistral.ChatMessage{
			{
				Role:    SYS_ROLE,
				Content: prompt,
			},
			{
				Role:    USER_ROLE,
				Content: fmt.Sprintf("Commit message: %s'", msg),
			},
		}, &mistral.ChatRequestParams{
			Temperature: 1,
			TopP:        1,
			MaxTokens:   200,
			SafePrompt:  false,
		},
	)

	if err != nil {
		return getFallbackRoast()
	}

	if len(chatRes.Choices) > 0 {
		return fmt.Sprintf("%s %s", emoji.FirstPlaceMedal, chatRes.Choices[0].Message.Content)
	}

	return getFallbackRoast()
}

func getFallbackRoast() string {

	fallbackRoastArray := []string{
		"Oh wow, another 'shit'. I'm sure this one won't educate anyone. 🤡",
		"Commit messages like these make us question our civilization. 🏺",
		"Oh good, another commit titled ‘Oops’—confidence inspiring. 👏",
		"‘WIP’—which stands for ‘Will It Pass’? 🤞",
		"This commit is so ugly even Git refuses to merge it. 🤢",
		"If your code were a novel, this commit would be the plot hole. 📖",
		"‘Commit message for clarity’—and now no one understands it. 🔍",
		"Your commit history is a crime scene, and this one is the murder weapon. 🔪",
		"Ah, another commit where we pretend everything is fine. 🔥🐶",
		"‘Fixed typo’—in the commit message, or the code? 🤔",
		"This commit looks like it was written under duress. 😰",
		"I see you’ve embraced the ‘Commit first, think later’ methodology. 🎭",
		"The best part of this commit is that it can be reverted. 🔄",
		"Your Git history reads like a bad soap opera. 🎭",
		"Ah, the legendary commit message. Until next time. ⏳",
		"If your commits had a theme song, it would be a sad violin. 🎻",
		"This commit deserves its own horror movie. 🎬",
	}

	rand.Seed(uint64(time.Now().UnixNano()))

	return string(emoji.FirstPlaceMedal) + " " + fallbackRoastArray[rand.Intn(len(fallbackRoastArray))]
}
