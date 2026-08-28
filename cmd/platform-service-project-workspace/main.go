package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openmcp-project/controller-utils/pkg/fips"
	"github.com/openmcp-project/platform-service-project-workspace/cmd/platform-service-project-workspace/app"
)

func main() {
	cmd := app.NewPlatformServiceProjectWorkspaceCommand()

	fips.Verify(context.Background())

	if err := cmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
