package config

import (
	"github.com/chirzul/gorent/src/databases/orm"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple backend for vehicle rental app with golang",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
