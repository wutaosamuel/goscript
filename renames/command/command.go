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
	rootCmd *cobra.Command // root command is main application command from cobra library
}

// NewCommand create a new command object
func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:     "listnames",
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
		inputFlag     []string
		outputDirFlag string
		beginFlag     bool
		pickFlag      []int
		reverseFlag   bool
		listFlag      int
	)
	operation := config.DefaultOp
	config := config.NewConfig()

	// setup root cmd
	c.rootCmd.PersistentFlags().StringSliceVarP(&inputFlag, "input", "i", make([]string, 0), "input directory or files")
	c.rootCmd.PersistentFlags().StringVarP(&outputDirFlag, "output", "o", "", "The directory of outputing copied files")
	c.rootCmd.PersistentFlags().BoolVarP(&beginFlag, "begin", "b", false, "add/delete/rename, at beginning characters of filename, default false: at end of filename")
	c.rootCmd.PersistentFlags().IntSliceVarP(&pickFlag, "pick", "p", make([]int, 0), "select files in a range, start from 0")
	c.rootCmd.PersistentFlags().BoolVarP(&reverseFlag, "reverse", "r", false, "Reverse files orders")
	c.rootCmd.PersistentFlags().IntVarP(&listFlag, "list", "l", 0, "use list files order/number, 0: default; 1: by name; 2: by time; 3: by size; 4: by extension")
	// required flag
	c.rootCmd.MarkPersistentFlagRequired("input")

	// setup subcommand cmd
	name, time, size, extension := false, false, false, false
	c.setupAddOp(&operation,config.Add)
	c.setupDeleteOp(&operation, config.Delete)
	c.setupRenameOp(&operation, config.Rename)
	c.setupListOp(&operation,&name, &time, &size, &extension)
	c.setupCountOp(&operation, config.Count)
	c.rootCmd.Execute()

	// check input from cli
	inputFiles, err := c.CheckInput(inputFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if outputDirFlag != "" {
		if err := c.CheckOutputDir(outputDirFlag); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
	c.checkIntSlice(pickFlag)
	c.checkList(listFlag)

	// setup config
	config.Files = inputFiles
	config.OutputDir = outputDirFlag
	config.Begin = beginFlag
	config.Pick = pickFlag
	config.Reverse = reverseFlag
	config.SetListOperation(listFlag)

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

// setupAddOp setup add command
func (c *Command) setupAddOp(op *config.OpCode, add *config.Add) {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add number or characters to filenames",
		Long: `Add number or characters to filename,
Default: auto accumulate by add number from 0.`,
		Run: func(cmd *cobra.Command, arg []string) {
			c.runError(op)
			*op = config.AddOp
		},
	}

	// set flags
	addCmd.Flags().StringVarP(&add.Char, "char", "c", "", "characters add into filename")
	addCmd.Flags().IntVarP(&add.Number, "number", "n", 0, "auto accumulate number by add 1 and start with n. if you want static number, pls use -c, --char")

	c.rootCmd.AddCommand(addCmd)
}

// setupDeleteOp setup Delete command
func (c *Command) setupDeleteOp(op *config.OpCode, delete *config.Delete) {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete characters of filename",
		Long: `Delete characters of filename. It not allow to process, if it has same filenames
Default: delete a character of finename`,
		Run: func(cmd *cobra.Command, arg []string) {
			c.runError(op)
			*op = config.DeleteOp
		},
	}

	// set flags
	deleteCmd.Flags().IntVarP(&delete.Number, "number", "n", 0, "delete number of characters of filename")

	c.rootCmd.AddCommand(deleteCmd)
}

// setupRenameOp setup rename command
func (c *Command) setupRenameOp(op *config.OpCode, rename *config.Rename) {
	var renameCmd = &cobra.Command{
		Use:   "rename",
		Short: "Rename filename",
		Long: `Rename filename to specify chars, auto accumulate number and extension,
Default: do nothing`,
		Run: func(cmd *cobra.Command, arg []string) {
			c.runError(op)
			*op = config.RenameOp
		},
	}

	// set flags
	renameCmd.Flags().StringVarP(&rename.Char, "char", "c", "", "rename characters. if no specify n/number, the accumulate number start with 0")
	renameCmd.Flags().IntVarP(&rename.Number, "number", "n", 0, "accumulate number, by add 1 and start with n")
	renameCmd.Flags().StringVarP(&rename.Extension, "extension", "e", "", "characters to change file extension")

	c.rootCmd.AddCommand(renameCmd)
}

// setupListOp setup extension
func (c *Command) setupListOp(op *config.OpCode,
	name, time, size, extension *bool) {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List ordered files",
		Long: `List files by name, time, size and extension.
-l, --list can use this file orders.
Default: list by default`,
		Run: func(cmd *cobra.Command, arg []string) {
			c.runError(op)
			*op = config.ListOp
		},
	}

	// set flags
	listCmd.Flags().BoolVarP(name, "name", "n", false, "sort by name")
	listCmd.Flags().BoolVarP(time, "time", "t", false, "sort by time")
	listCmd.Flags().BoolVarP(size, "size", "s", false, "sort by size")
	listCmd.Flags().BoolVarP(extension, "extension", "e", false, "sort by list")

	c.rootCmd.AddCommand(listCmd)
}

// setupCountOp setup extension
func (c *Command) setupCountOp(op *config.OpCode, count *config.Count) {
	var countCmd = &cobra.Command{
		Use:   "count",
		Short: "Count total number or unmatched characters of filename",
		Long: `Count total number or unmatched characters of filename,
Default: display total number of filename`,
		Run: func(cmd *cobra.Command, arg []string) {
			c.runError(op)
			*op = config.ListOp
		},
	}

	// set flags
	countCmd.Flags().StringVarP(&count.Char, "char", "c", "", "count filename")

	c.rootCmd.AddCommand(countCmd)
}

// runError running multiple command
func (c *Command) runError(op *config.OpCode) {
	if *op != config.DefaultOp {
		fmt.Println("Not allow multiple subcommands")
		c.UsageExit()
	}
}
