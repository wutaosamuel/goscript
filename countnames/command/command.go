package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"../common"
	"../config"
)

// Command command object
type Command struct {
	rootCmd   *cobra.Command // root command is main application command from cobra library
	isExecute bool           // does root command has been exectue()
}

// NewCommand create a new command object
func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:     "countnames",
		Version: common.Version(),
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	return &Command{
		rootCmd:   rootCmd,
		isExecute: false,
	}
}

// GetIsExecute get state of root command has been executed
func (c *Command) GetIsExecute() bool {
	return c.isExecute
}

// Execute execute main rootCmd
func (c *Command) Execute() *config.Config {
	// setup rootCmd flags
	var (
		inputFlag []string
		charFlag string
		matchFlag bool
	)

	// check if it has bee execute
	if c.isExecute {
		fmt.Println("Pls, check it has been executed")
		os.Exit(0)
	}

	// setup root cmd
	c.rootCmd.PersistentFlags().StringSliceVarP(&inputFlag, "input", "i", make([]string, 0), "input files")
	c.rootCmd.PersistentFlags().StringVarP(&charFlag, "char", "c", "", "Characters for match or unmatch, empty displays total number of filenames")
	c.rootCmd.PersistentFlags().BoolVarP(&matchFlag, "match", "m", false, "match or unmatch for characters")
	// required flag
	c.rootCmd.MarkPersistentFlagRequired("input")

	// setup subcommand and rootCmd
	c.rootCmd.Execute()

	// setup config
	config := config.NewConfig()

	return config
}

// UsageExit printout the usage of the root command the exit
func (c *Command) UsageExit() {
	if !c.isExecute {
		fmt.Println("The rootCmd has not been executed")
		os.Exit(0)
	}

	fmt.Print(c.rootCmd.UsageString())
	os.Exit(1)
}

// SetArgs set args, only using testing
func (c *Command) SetArgs(args []string) {
	c.rootCmd.SetArgs(args)
}
