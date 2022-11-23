package cmd

import (
	"fmt"
	"github.com/lifangjunone/go-micro/conf"
	"github.com/spf13/cobra"
	"os"
)

var (
	printVersion bool
	configType   string
	configFile   string
	configEtcd   string
)

var RootCmd = &cobra.Command{
	Use:   "go-micro",
	Short: "GO 微服务框架",
	Long:  "GO 微服务框架，实现了Grpc和HTTP协议",
	RunE: func(cmd *cobra.Command, args []string) error {
		if printVersion {
			fmt.Println(conf.Version)
			return nil
		}
		return cmd.Help()
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&configType, "config-type", "t", "file", "The service config type [file/env/etcd")
	RootCmd.PersistentFlags().StringVarP(&configFile, "config-file", "f", "etc/config.toml", "The service config from file")
	RootCmd.PersistentFlags().StringVarP(&configEtcd, "config-etcd", "e", "127.0.1:2379", "The service config from etcd")
	RootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "v", false, "If print version")
}
