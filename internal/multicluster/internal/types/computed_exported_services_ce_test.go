// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent
// +build !consulent

package types

import (
	"github.com/hashicorp/consul/internal/resource/resourcetest"
	multiclusterv1alpha1 "github.com/hashicorp/consul/proto-public/pbmulticluster/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComputedExportedServicesValidations(t *testing.T) {
	type testcase struct {
		Resource *pbresource.Resource
	}
	run := func(t *testing.T, tc testcase) {
		err := MutateComputedExportedServices(tc.Resource)
		require.NoError(t, err)
		err = ValidateComputedExportedServices(tc.Resource)
		require.NoError(t, err)
	}

	cases := map[string]testcase{
		"computed exported services with peer": {
			Resource: resourcetest.Resource(multiclusterv1alpha1.ComputedExportedServicesType, ComputedExportedServicesName).
				WithData(t, validComputedExportedServicesWithPeer()).
				Build(),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			run(t, tc)
		})
	}
}
