// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !jsoniter
// +build !jsoniter

package json

import jsontime "github.com/ryanker/gin_v1/jsontime"

var (
	json = ConfigWithCustomTimeFormat
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)
