// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.11

package http2

func traceHasWroteHeaderField(trace *clientTrace) bool {
	return trace != nil && trace.WroteHeaderField != nil
}

func traceWroteHeaderField(trace *clientTrace, k, v string) {
	if trace != nil && trace.WroteHeaderField != nil {
		trace.WroteHeaderField(k, []string{v})
	}
}
