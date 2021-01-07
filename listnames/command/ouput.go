package command

import (
	"errors"

	"github.com/wutaosamuel/goscript/listnames/utils"
)

// CheckOutputDir check output dir
func (c *Command) CheckOutputDir(dir string) error {
	isDir, err := utils.IsDir(dir)
	if err != nil {
		return err
	}
	if isDir {
		return nil
	}
	if !isDir {
		c.UsageExit()
	}

	return errors.New("Unknown error")
}
