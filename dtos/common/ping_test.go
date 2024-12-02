//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	"github.com/agile-edge/go-mod-core-contracts/v3/common"
)

func TestNewPingResponse(t *testing.T) {
	serviceName := uuid.NewString()
	target := NewPingResponse(serviceName)

	assert.Equal(t, common.ApiVersion, target.ApiVersion)
	_, err := time.Parse(time.UnixDate, target.Timestamp)
	assert.NoError(t, err)
	assert.Equal(t, serviceName, target.ServiceName)
}
