package listnames

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/wutaosamuel/goscript/listnames/command"
	"github.com/wutaosamuel/goscript/listnames/config"
	j "github.com/wutaosamuel/goscript/listnames/job"
)

// Main func
func Main() {
	var result, selectedFiles []string
	var cmd = command.NewCommand()
	var cfg = config.NewConfig()

	// sort files
	cfg = cmd.Execute()
	result, err := Execute(cfg.Files, cfg.Pick, cfg.Reverse, cfg.Operation)
	if err != nil {
		fmt.Println(err)
		cmd.UsageExit()
	}

	// check if files have been selection
	selectedFiles = SelectFiles(result, cfg.Pick)

	// print out result
	// 	- no selection
	//	- selection
	fmt.Println("List file names: ")
	fmt.Println()
	if len(selectedFiles) == 0 {
		ConsoleString(result)
	}
	if len(selectedFiles) != 0 {
		SelectConsoleString(result, selectedFiles)
	}
}

// Execute get result
func Execute(files []string, pick []int, reverse bool, operation config.OpCode) ([]string, error) {
	var job = j.NewJobRead(files, pick, reverse)

	if operation == config.ErrorOp {
		return []string{}, errors.New("Job Execute: Unknown Error")
	}
	if operation == config.NameOp || operation == config.DefaultOp{
		return job.SortName(), nil
	}
	if operation == config.TimeOp {
		return job.SortMTime(), nil
	}
	if operation == config.SizeOp {
		return job.SortSize(), nil
	}
	if operation == config.ExtensionOp {
		return job.SortExtension(), nil
	}

	return []string{}, errors.New("Job Execute: Unknown Error")
}

// SelectFiles select files
// 	- start from 0
func SelectFiles(files []string, pick []int) []string {
	var isStart bool
	var pLocation int
	var result []string
	// do not nothing, when no files or selection
	if len(files) == 0 || len(pick) == 0 {
		return files
	}

	isStart = false
	pLocation = 0
	sort.Ints(pick)
	for k, v := range files {
		if pLocation < len(pick) {
			if k == pick[pLocation] {
				isStart = !isStart
				pLocation++
			}
		}
		if isStart {
			result = append(result, v)
		}
	}

	return result
}

// ConsoleString print out console string
func ConsoleString(messages []string) {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 4, 0, '\t', 0)

	for k, v := range messages {
		n := strconv.Itoa(k)
		fmt.Fprintln(writer, n+"\t"+v)
	}
	writer.Flush()
}

// SelectConsoleString print out selected files
func SelectConsoleString(files []string, selectFiles []string) {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 4, 0, '\t', 0)

	for k, v := range files {
		for _, s := range selectFiles {
			if v == s {
				n := strconv.Itoa(k)
				fmt.Fprintln(writer, n+"\t"+v)
				break
			}
		}
	}

	writer.Flush()
}
