package command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wutaosamuel/goscript/renames/utils"
)

// CheckInput check the input
func (c *Command) CheckInput(input []string) ([]string, error) {
	var dirs []string
	var fileNames []string

	// check all files and dirs exist
	// classify the dir and files
	if len(input) == 0 || input == nil {
		os.Exit(1)
		return input, nil
	}
	var notExistFile []string
	for _, f := range input {
		isExist, err := utils.IsExist(f)
		if err != nil {
			return nil, err
		}
		if !isExist {
			notExistFile = append(notExistFile, f)
			continue
		}

		isDir, err := utils.IsDir(f)
		if err != nil {
			return nil, err
		}
		if isDir {
			dirs = append(dirs, f)
			continue
		}

		fileNames = append(fileNames, f)
	}
	// if it has not existed files, not allow to run program
	if len(notExistFile) != 0 {
		for _, v := range notExistFile {
			fmt.Println(v, " Not exist")
		}
		c.UsageExit()
	}

	// not allow have files and dirs, only one type of input
	if len(fileNames) != 0 && len(dirs) != 0 {
		fmt.Println("Not allow process file and dir at the same time")
		c.UsageExit()
	}

	// get file
	if len(fileNames) > 0 {
		return fileNames, nil
	}

	// only dir
	// not allow multiple directory
	// TODO: support multiple directory
	if len(dirs) > 1 {
		fmt.Println("Not support process more than 1 directory")
		c.UsageExit()
	}
	// read file from directory
	if len(dirs) == 1 {
		files, err := ioutil.ReadDir(dirs[0])
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			fileNames = append(fileNames, filepath.Join(dirs[0], f.Name()))
		}

		return fileNames, nil
	}

	return nil, errors.New("Unknown error")
}

// checkIntSlice not negative number
func (c *Command) checkIntSlice(ints []int) {
	for _, v := range ints {
		if v < 0 {
			fmt.Println("The range should be positive")
			c.UsageExit()
		}
	}
}

// checkList check list no more than 5
func (c *Command) checkList(list int) {
	if list < 0 {
		fmt.Println("Pls, enter a positive integer")
		c.UsageExit()
	}
	if list > 4 {
		fmt.Println("Have only 5 operations(0~5)")
		c.UsageExit()
	}
}
