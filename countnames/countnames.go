package countnames

import (
	"fmt"

	"./command"
	"./config"
)

// Main func
func Main() {
	var cmd = command.NewCommand()
	var cfg = config.NewConfig()

	cfg = cmd.Execute()
	fmt.Println(cfg)
}