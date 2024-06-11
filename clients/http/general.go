//
// Copyright (C) 2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"github.com/agile-edgex/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/agile-edgex/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/agile-edgex/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/agile-edgex/go-mod-core-contracts/v3/dtos/common"
	"github.com/agile-edgex/go-mod-core-contracts/v3/errors"
)

type generalClient struct {
	baseUrl      string
	authInjector interfaces.AuthenticationInjector
}

func NewGeneralClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.GeneralClient {
	return &generalClient{
		baseUrl:      baseUrl,
		authInjector: authInjector,
	}
}

func (g *generalClient) FetchConfiguration(ctx context.Context) (res dtoCommon.ConfigResponse, err errors.EdgeX) {
	err = utils.GetRequest(ctx, &res, g.baseUrl, common.ApiConfigRoute, nil, g.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return res, nil
}
