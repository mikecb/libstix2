// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package stix

type indicatorType struct {
	MessageType  string   `json:"type,omitempty"`
	ID           string   `json:"id,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	Title        string   `json:"title,omitempty"`
	Descriptions []string `json:"descriptions,omitempty"`
}

func newIndicator() indicatorType {
	var obj indicatorType
	obj.MessageType = "indicator"
	obj.ID = createID("indicator")
	obj.CreatedAt = getCurrentTime()
	return obj
}

func (i *indicatorType) SetTitle(t string) {
	i.Title = t
}
