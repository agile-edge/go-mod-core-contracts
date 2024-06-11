//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/agile-edgex/go-mod-core-contracts/v3/dtos/common"
	"github.com/agile-edgex/go-mod-core-contracts/v3/errors"
)

type GeneralClient interface {
	// FetchConfiguration obtains configuration information from the target service.
	FetchConfiguration(ctx context.Context) (common.ConfigResponse, errors.EdgeX)
}
