// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the corresponding structures to the
// "Diagnostics" part of the LSP specification.

package protocol

// PublishDiagnosticsParams defines the parameters to the
// `textDocument/publishDiagnostics` method.
type PublishDiagnosticsParams struct {
	// URI for which diagnostic information is reported.
	URI DocumentURI `json:"uri"`

	// Diagnostics contains an array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}
