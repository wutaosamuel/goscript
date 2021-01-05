package renames

import (
	"./common"
	"./command"
	j "./job"
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
	}
	if cfg.OpCode == common.DeleteOp {
		job.Delete()
	}
	if cfg.OpCode == common.RenameOp {
		job.Rename()
	}
	if cfg.OpCode == common.ListOp {
		job.List()
	}
	if cfg.OpCode == common.CountOp {
		job.Count()
	}
}
