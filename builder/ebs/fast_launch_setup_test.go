// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebs

import "testing"

func TestFastLaunchPrepare(t *testing.T) {
	tests := []struct {
		name      string
		config    FastLaunchConfig
		expectErr bool
	}{
		{
			"OK - empty config",
			FastLaunchConfig{},
			false,
		},
		{
			"OK - all specified, with template id",
			FastLaunchConfig{
				UseFastLaunch:         true,
				LaunchTemplateID:      "id",
				LaunchTemplateVersion: 2,
				MaxParallelLaunches:   10,
				TargetResourceCount:   20,
			},
			false,
		},
		{
			"OK - all specified, with template name",
			FastLaunchConfig{
				UseFastLaunch:         true,
				LaunchTemplateName:    "name",
				LaunchTemplateVersion: 2,
				MaxParallelLaunches:   10,
				TargetResourceCount:   20,
			},
			false,
		},
		{
			"Error - max parallel launches < 6",
			FastLaunchConfig{
				MaxParallelLaunches: 3,
			},
			true,
		},
		{
			"Error - target resource count < 0",
			FastLaunchConfig{
				TargetResourceCount: -1,
			},
			true,
		},
		{
			"Error - launch template ID & name specified",
			FastLaunchConfig{
				LaunchTemplateID:   "id",
				LaunchTemplateName: "name",
			},
			true,
		},
		{
			"Error - launch template version without name/id",
			FastLaunchConfig{
				LaunchTemplateVersion: 2,
			},
			true,
		},
		{
			"Error - launch template version < 0",
			FastLaunchConfig{
				LaunchTemplateVersion: -1,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.config.Prepare()
			if (len(errs) != 0) != tt.expectErr {
				t.Errorf("error mismatch, expected %t, got %d errors", tt.expectErr, len(errs))
			}

			for _, err := range errs {
				t.Logf("got error %q", err)
			}
		})
	}
}
