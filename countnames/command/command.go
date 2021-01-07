package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"goscript/countnames/common"
	"goscript/countnames/config"
)

// Command command object
type Command struct {
	rootCmd *cobra.Command // root command is main application command from cobra library
}

// NewCommand create a new command object
func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:     "countnames",
		Version: common.Version(),
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	return &Command{
		rootCmd: rootCmd,
	}
}

// Execute execute main rootCmd
func (c *Command) Execute() *config.Config {
	// setup rootCmd flags
	var (
		inputFlag []string
		charFlag  string
	)

	// setup root cmd
	c.rootCmd.PersistentFlags().StringSliceVarP(&inputFlag, "input", "i", make([]string, 0), "input files")
	c.rootCmd.PersistentFlags().StringVarP(&charFlag, "char", "c", "", "Characters for match or unmatch, empty displays total number of filenames")
	// required flag
	c.rootCmd.MarkPersistentFlagRequired("input")

	// setup subcommand and rootCmd
	c.rootCmd.Execute()

	// check input from cli
	inputFiles, err := c.CheckInput(inputFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// setup config
	config := config.NewConfig()
	config.Files = inputFiles
	config.Char = charFlag

	return config
}

// UsageExit printout the usage of the root command the exit
func (c *Command) UsageExit() {
	fmt.Print(c.rootCmd.UsageString())
	os.Exit(1)
}

// SetArgs set args, only using testing
func (c *Command) SetArgs(args []string) {
	c.rootCmd.SetArgs(args)
}
