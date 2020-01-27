// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package opcua

// BrowsedNode represents the details of a browsed OPC-UA node.
type BrowsedNode struct {
	NodeID      string
	Type        string
	Description string
}

// Browser represents the OPC-UA Server Nodes browser.
type Browser interface {
	// Browse availlable Nodes for a given URI.
	Browse(string, string) ([]BrowsedNode, error)
}
