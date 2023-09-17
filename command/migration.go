package command

import (
	"fmt"
	"mobile-ecommerce/config"
	"os"
	"strings"
	"time"

	migrateV4 "github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const timeLayoutString = "20060102150405"

func MigrationCommand(cfg config.DBConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use: "migrate",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "up",
		Short: "migrate up",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrateV4.New(cfg.MigrateFolder, cfg.Source)
			if err != nil {
				logrus.Fatalf("get error when create migrate instance: %v", err)
			}
			logrus.Info("starting migrate up...")
			err = m.Up()
			if err != nil {
				logrus.Fatalf("get error when migrate up: %v", err)
			}
			logrus.Info("migrate up successfully")
		},
	}, &cobra.Command{
		Use:   "create",
		Short: "create migration files",
		Run: func(cmd *cobra.Command, args []string) {
			folder := strings.ReplaceAll(cfg.MigrateFolder, "file://", "")
			time := time.Now().Format(timeLayoutString)
			name := strings.Join(args, "-")

			upFile := fmt.Sprintf("%s/%s_%s.up.sql", folder, time, name)
			downFile := fmt.Sprintf("%s/%s_%s.down.sql", folder, time, name)

			if _, err := os.Create(upFile); err != nil {
				logrus.Fatalf("get error when create up file: %v", err)
			}
			if _, err := os.Create(downFile); err != nil {
				logrus.Fatalf("get error when create down file: %v", err)
			}

			logrus.Info("create migration files successfully")
		},
	})

	return cmd
}
