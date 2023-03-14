package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/aiomni/aish/ai"
	"github.com/aiomni/aish/config"
	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/liamg/tml"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type AskResponse struct {
	Code    int8   `json:"code"`
	Command string `json:"command,omitempty"`
	Note    string `json:"note,omitempty"`
	Warning string `json:"warning,omitempty"`
}

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask AI what you want to do",
	Long:  `Ask AI what you want to do`,
	Run: func(cmd *cobra.Command, args []string) {

		if config.GetAPIKey() == "" {
			tml.Printf("<red>Error: </red>OpenAI API Key NOT setting, use <green>`aish config set`</green> set API Key.\n")
			os.Exit(1)
		}

		if config.GetOrganizationID() == "" {
			tml.Printf("<red>Error: </red>OpenAI Organization ID NOT setting, use <green>`aish config set`</green> set Organization ID.\n")
			os.Exit(1)
		}

		content := strings.Join(args, "")

		if len(args) == 0 {
			prompt := promptui.Prompt{
				Label: "Please Input What you want todo",
				Validate: func(input string) error {

					if len(input) == 0 {
						return errors.New("invalid input")
					}
					return nil
				},
			}

			content, _ = prompt.Run()
		}

		s := spinner.New(spinner.CharSets[38], 100*time.Millisecond)
		s.Suffix = " Searching from chatGPT...\n"
		s.Start()
		res, err := ai.AskChatGPT(content)
		s.Stop()

		if err != nil {
			color.Red(err.Error())
			return
		}

		var recommend AskResponse
		err = json.Unmarshal([]byte(res), &recommend)
		if err != nil || recommend.Code != 1 || len(recommend.Command) == 0 {
			color.Red("Unable to find the command you are looking for, please modify and try again.")
			return
		}

		colorCommand := color.New(color.FgHiMagenta, color.BgBlack).Sprint(recommend.Command)
		fmt.Printf("Found the command for you: " + colorCommand + "\n")
		if len(recommend.Note) > 0 {
			color.Cyan(recommend.Note)
		}
		if len(recommend.Warning) > 0 {
			color.Yellow(recommend.Warning)
		}

		// TODO: Find Other
		options := []string{"Execute", "Copy", "Abort"}

		prompt := promptui.Select{
			Label: "Select What you want to do",
			Items: options,
			// Templates: templates,
			Size: 4,
			// Searcher:  searcher,
		}

		opt, _, err := prompt.Run()
		if err != nil {
			color.Red(err.Error())
			return
		}

		switch opt {
		case 0:
			execute(recommend.Command)
		case 1:
			err := clipboard.WriteAll(recommend.Command)
			if err != nil {
				color.Red("复制失败，请手动复制！")
				os.Exit(1)
			}
			color.Green("复制成功！")
		case 2:
			os.Exit(1)
		}
	},
}

func execute(command string) {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}
