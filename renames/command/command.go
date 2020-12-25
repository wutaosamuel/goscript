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
		Use:     "listnames",
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
		inputFlag     []string
		outputDirFlag string
		reverseFlag   bool
		pickFlag      []int
	)
	var operation config.OpCode

	// check it has been execute
	if c.isExecute {
		fmt.Println("Pls, check it has been executed")
		os.Exit(0)
	}

	// setup root cmd
	c.rootCmd.PersistentFlags().StringSliceVarP(&inputFlag, "input", "i", make([]string, 0), "input directory or files")
	c.rootCmd.PersistentFlags().StringVarP(&outputDirFlag, "output", "o", "", "The directory of outputing copied files")
	c.rootCmd.PersistentFlags().BoolVarP(&reverseFlag, "reverse", "r", false, "Reverse orders")
	c.rootCmd.PersistentFlags().IntSliceVarP(&pickFlag, "pick", "p", make([]int, 0), "select files in a range, start from 0")
	// required flag
	c.rootCmd.MarkPersistentFlagRequired("input")

	// setup subcommand cmd
	operation = config.DefaultOp
	c.setupNameOp(&operation)
	c.setupTimeOp(&operation)
	c.setupSizeOp(&operation)
	c.setupExtensionOp(&operation)
	c.rootCmd.Execute()
	c.isExecute = true

	// check input from cli
	inputFiles, err := c.CheckInput(inputFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	c.checkIntSlice(pickFlag)
	if outputDirFlag != "" {
		err = c.CheckOutputDir(outputDirFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	// setup config
	config := config.NewConfig()
	config.Files = inputFiles
	config.OutputDir = outputDirFlag
	config.Reverse = reverseFlag
	config.Pick = pickFlag
	config.Operation = operation

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

// setupNameOp setup name command
func (c *Command) setupNameOp(op *config.OpCode) {
	var nameCmd = &cobra.Command{
		Use:   "name",
		Short: "Name sorts filenames by their name, default operation",
		Long: `Name sorts filenames by their name characters.
Default operation of program, if no specify of the subcommand`,
		Run: func(cmd *cobra.Command, arg []string) {
			if runError(op) {
				*op = config.NameOp
			}
		},
	}

	c.rootCmd.AddCommand(nameCmd)
}

// setupTimeOp setup time command
func (c *Command) setupTimeOp(op *config.OpCode) {
	var timeCmd = &cobra.Command{
		Use:   "time",
		Short: "Time sorts filenames by their created or added time. [time-1] -> name+1",
		Long: `Time sorts filenames by their created or added time,
Default order: the ordest file is at the first one. [time-1] -> name+1`,
		Run: func(cmd *cobra.Command, arg []string) {
			if runError(op) {
				*op = config.TimeOp
			}
		},
	}

	c.rootCmd.AddCommand(timeCmd)
}

// setupSizeOp setup size command
func (c *Command) setupSizeOp(op *config.OpCode) {
	var sizeCmd = &cobra.Command{
		Use:   "size",
		Short: "Size sorts filenames by their size. [size+1] -> name+1.",
		Long: `Time sorts filenames by their created or added time,
Default order: the smallest size file is at the first one [size+1] -> name+1`,
		Run: func(cmd *cobra.Command, arg []string) {
			if runError(op) {
				*op = config.SizeOp
			}
		},
	}

	c.rootCmd.AddCommand(sizeCmd)
}

// setupExtensionOp setup extension
func (c *Command) setupExtensionOp(op *config.OpCode) {
	var extensionCmd = &cobra.Command{
		Use:   "extension",
		Short: "Extension sorts filename by their extension",
		Long: `Extension sorts filename by their extension,
Default order: is similiar to sorted by name`,
		Run: func(cmd *cobra.Command, arg []string) {
			if runError(op) {
				*op = config.ExtensionOp
			}
		},
	}

	c.rootCmd.AddCommand(extensionCmd)
}

func runError(op *config.OpCode) bool {
	if *op != config.DefaultOp {
		*op = config.ErrorOp
		return false
	}

	return true
}
