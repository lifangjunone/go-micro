package cmd

import (
	"context"
	"github.com/lifangjunone/go-micro/conf"
	"github.com/spf13/cobra"
	"io/ioutil"
	"time"
)

var (
	createTableFilePath string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "server init",
	Long:  "To init server config",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := loadServerConfig(configType); err != nil {
			return err
		}
		err := createTables()
		if err != nil {
			return err
		}
		return nil
	},
}

func createTables() error {
	db, err := conf.GetConfig().MySQL.GetDB()
	if err != nil {
		return err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	// 读取SQL文件
	sqlFile, err := ioutil.ReadFile(createTableFilePath)
	if err != nil {
		return err
	}
	// 自行SQL
	_, err = db.ExecContext(ctx, string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}

func init() {
	initCmd.PersistentFlags().StringVarP(&createTableFilePath, "sql-file-path", "s", "schema/tables.sql", "The sql file path")
	RootCmd.AddCommand(initCmd)
}
