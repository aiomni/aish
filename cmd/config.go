package cmd

import (
	"errors"

	"github.com/aiomni/aish/config"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set your custom config",
	Long:  `Set your custom config, include OpenAI API Key, HTTP Proxy, and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := []string{"API Key", "Organization ID", "Proxy Domain"}
		prompt := promptui.Select{
			Label: "Select What you want to set",
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
		case 0: // 设置 API Key
			setAPIKey()
			return
		case 1: // 设置 Organization ID
			setOrgID()
			return
		case 2:
			setProxyDomain()
			return
		}

	},
}

func setAPIKey() {
	prompt := promptui.Prompt{
		Label: "Please Input your custom OpenAI API Key",
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("invalid input")
			}
			return nil
		},
	}

	content, _ := prompt.Run()

	err := config.SetAPIKey(content)
	if err != nil {
		color.Red("Set API Key Error: %s", err.Error())
		return
	}
	color.Green("Set API Key Success.")
}

func setOrgID() {
	prompt := promptui.Prompt{
		Label: "Please Input your custom OpenAI Organization ID",
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("invalid input")
			}
			return nil
		},
	}

	content, _ := prompt.Run()

	err := config.SetOrganizationID(content)
	if err != nil {
		color.Red("Set Organization ID Error: %s", err.Error())
		return
	}
	color.Green("Set Organization ID Success.")
}

func setProxyDomain() {
	prompt := promptui.Prompt{
		Label: "Please Input your custom Proxy Domain",
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("invalid input")
			}
			return nil
		},
	}

	content, _ := prompt.Run()

	err := config.SetProxyDomain(content)
	if err != nil {
		color.Red("Set Organization ID Error: %s", err.Error())
		return
	}
	color.Green("Set Organization ID Success.")
}
