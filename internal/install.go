package internal

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-getter"
)

func InstallPackage(ctx context.Context, global bool) error {

	err := getter.Get("", "https://github.com/SilasMeyer4/Lumadi/releases/download/v0.2.0/Lumadi-0.1.0-win64.tar.gz")

	if err != nil {
		fmt.Println("Error installing package:", err)
		return err
	}

	return nil
}
