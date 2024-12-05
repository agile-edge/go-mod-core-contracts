//
// Copyright (c) 2018 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package logger

import (
	"testing"

	"github.com/agile-edge/go-mod-core-contracts/v3/models"
)

func TestLogLevel(t *testing.T) {
	expectedLogLevel := models.DebugLog
	lc := NewClient("testService", expectedLogLevel)
	lc.Debugf("key is %s", "edge")
}
