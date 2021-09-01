package main

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/server"
	"github.com/spf13/cobra"
	"gopkg.pigeonligh.com/easygo/cli/flags"
	"gopkg.pigeonligh.com/easygo/elog"
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
	_ = flags.ObjectVar(cmd.Flags(), config, nil)

	if err := cmd.Execute(); err != nil {
		elog.Fatal(err)
	}
}
