package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	pya "github.com/SilasMeyer4/pya/internal"

	"github.com/urfave/cli/v3"
)

func main() {
	app := InitAppData()
	app.Version = "0.1.0"

	fmt.Printf("app: %v\n", app.AppConfiguration.PersistorVersion)

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Installs a package",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "global",
						Aliases: []string{"g"},
						Usage:   "Install the package globally",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("Installing package...")
					pya.InstallPackage(ctx, cmd.Bool("global"))
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Removes a package",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("Removing package...")
					return nil
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Displays the version of the application",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(app.Version)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func InitAppData() App {
	app := App{}

	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Error getting user config dir:", err)
	}

	dir = filepath.Join(dir, "pya")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal("Error creating config directory:", err)
		}
	}

	if _, err := os.Stat(filepath.Join(dir, "config.json")); os.IsNotExist(err) {
		file, err := os.Create(filepath.Join(dir, "config.json"))
		if err != nil {
			log.Fatal("Error creating config file:", err)
		}
		CreateInitialConfig(file)
		file.Close()
	} else {
		file, err := os.OpenFile(filepath.Join(dir, "config.json"), os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal("Error opening config file:", err)
		}

		config := ReadConfig(file)
		app.AppConfiguration = config
		file.Close()
	}
	return app
}

type App struct {
	Version          string
	AppConfiguration AppConfiguration
}
