package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/aiomni/aish/config"
	"github.com/aiomni/aish/utils"
	"github.com/fatih/color"
	"github.com/liamg/tml"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var configOptions = []string{"API Key", "Organization ID", "Proxy Domain"}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and Get your custom config",
	Long:  `Set and Get your custom config, include OpenAI API Key, HTTP Proxy, and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		var option string
		if len(args) > 0 {
			option = args[0]
		} else {
			prompt := promptui.Select{
				Label: "Select What you want to do",
				Items: []string{"Set Config", "Get Config", "Reset Config"},
				// Templates: templates,
				Size: 4,
				// Searcher:  searcher,
			}

			opt, _, err := prompt.Run()
			if err != nil {
				tml.Printf("<red>Something went wrong!</red>\n")
				os.Exit(0)
			}

			switch opt {
			case 0:
				option = "set"
			case 1:
				option = "get"
			case 2:
				option = "reset"
			default:
				os.Exit(1)
			}
		}

		switch option {
		case "set":
			setConfig()
		case "get":
			getConfig()
		case "reset":
			resetConfig()
		default:
			tml.Printf("<red>Unknown Command!</red>\n")
			os.Exit(1)
		}
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set your custom config",
	Long:  `Set your custom config, include OpenAI API Key, HTTP Proxy, and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		setConfig()
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your custom config",
	Long:  `Get your custom config, include OpenAI API Key, HTTP Proxy, and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		getConfig()
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset your custom config",
	Long:  `Reset your custom config, include OpenAI API Key, HTTP Proxy, and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		resetConfig()
	},
}

func setConfig() {
	prompt := promptui.Select{
		Label: "Select What you want to set",
		Items: configOptions,
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

	if !utils.ValidateApiKey(content) {
		color.Red("API Key NOT Validate: %s", content)
		return
	}

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

func getConfig() {
	prompt := promptui.Select{
		Label: "Select What you want to get",
		Items: configOptions,
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
		printResult("API Key", config.GetAPIKey())
	case 1:
		printResult("Organization ID", config.GetOrganizationID())
	case 2:
		printResult("Proxy Domain", config.GetProxyDomain())
	}
}

func printResult(name, content string) {
	if content == "" {
		tml.Printf(fmt.Sprintf("<red>There is no %s set.</red>\n", name))
	} else {
		tml.Printf(fmt.Sprintf("%s is: <green>%s</green>\n", name, content))
	}
}

func resetConfig() {
	prompt := promptui.Select{
		Label: "Select What you want to reset",
		Items: configOptions,
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
		config.SetAPIKey("")
		tml.Printf("<green>Reset API Key Config Success.</green>")
	case 1: // 设置 Organization ID
		config.SetOrganizationID("")
		tml.Printf("<green>Reset Organization ID Config Success.</green>")
	case 2:
		config.SetProxyDomain("")
		tml.Printf("<green>Reset Proxy Domain Config Success.</green>")
	}
}
