// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build go all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := integration.ProgramTestOptions{}
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-grafana/sdk",
		},
	})

	return baseGo
}
