package main

import (
	"github.com/pigeonligh/easygo/cli/flags"
	"github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/ffxiv-todo/pkg/server"
	"github.com/spf13/cobra"
)

func main() {
	config := &server.Config{
		Debug: false,
		Port:  80,
		Data:  "ffxiv-datamining-cn",
	}

	cmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			_ = server.RunServer(config)
		},
	}
	_ = flags.ObjectVar(cmd.Flags(), config, "")

	if err := cmd.Execute(); err != nil {
		elog.Fatal(err)
	}
}
