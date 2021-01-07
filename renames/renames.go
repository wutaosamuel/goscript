package renames

import (
	"fmt"

	"github.com/wutaosamuel/goscript/renames/common"
	"github.com/wutaosamuel/goscript/renames/command"
	j "github.com/wutaosamuel/goscript/renames/job"
)

// Main func
func Main() {
	var cmd = command.NewCommand()
	var cfg = cmd.Execute()
	var job = j.ReadConfig(cfg)

	if cfg.OpCode == common.DefaultOp {
		job.List()
	}
	if cfg.OpCode == common.AddOp {
		job.Add()
		fmt.Println("Add done")
	}
	if cfg.OpCode == common.DeleteOp {
		job.Delete()
		fmt.Println("Delete done")
	}
	if cfg.OpCode == common.RenameOp {
		job.Rename()
		fmt.Println("Rename done")
	}
	if cfg.OpCode == common.ListOp {
		job.List()
	}
	if cfg.OpCode == common.CountOp {
		job.Count()
	}
}
