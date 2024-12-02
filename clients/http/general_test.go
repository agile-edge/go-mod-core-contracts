//
// Copyright (C) 2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/agile-edge/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/agile-edge/go-mod-core-contracts/v3/dtos/common"
)

func Test_generalClient_FetchConfiguration(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiConfigRoute, dtoCommon.ConfigResponse{})
	defer ts.Close()

	client := NewGeneralClient(ts.URL, NewNullAuthenticationInjector())
	res, err := client.FetchConfiguration(context.Background())
	require.NoError(t, err)
	require.IsType(t, dtoCommon.ConfigResponse{}, res)
}
