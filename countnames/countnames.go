package countnames

import (
	"os"
	"fmt"
	"strconv"
	"text/tabwriter"

	"./command"
	"./config"
	j "./job"
)

// Main func
func Main() {
	var cmd = command.NewCommand()
	var cfg = config.NewConfig()

	cfg = cmd.Execute()
	result := Execute(cfg.Files, cfg.Char)

	ConsoleString(cfg.Files, result)
}

// Execute count filenames
func Execute(files []string, char string) []string {
	if len(files) == 0 {
		return []string{}
	}

	job := j.NewJob()
	job.SetJob(files, char)
	if char == "" {
		return job.CountDefault()
	}
	if char != "" {
		return job.CountMatch()
	}

	return []string{}
}

// ConsoleString print out console string
func ConsoleString(files []string, result []string) {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 8, 0, '\t', 0)
	// check if no files
	if len(files) == 0 && len(result) == 0 {
		fmt.Println("no files")
		return
	}
	// check if files not equal to result
	if len(files) != len(result) {
		fmt.Println("unknown error: files not equal to result")
	}

	// display -> n filename.txt 8.txt
	for k, v := range files {
		n := strconv.Itoa(k)
		fmt.Fprintln(writer, n+".\t"+v+"\t"+result[k])
	}
	writer.Flush()
}
