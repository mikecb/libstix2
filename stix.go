// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package stix

import (
	"time"

	"github.com/pborman/uuid"
)

// Version = "0.1"
const (
	TIME_RFC_3339       = "2006-01-02T15:04:05Z07:00"
	TIME_RFC_3339_MICRO = "2006-01-02T15:04:05.999999Z07:00"
)

func createID(t string) string {
	// TODO Add check to validate input value
	id := t + "--" + uuid.New()
	return id
}

func getCurrentTime() string {
	t := time.Now().UTC().Format(TIME_RFC_3339)
	return t
}

//PackageType is a generic quantum of information.
type PackageType struct {
	MessageType string          `json:"type,omitempty"`
	ID          string          `json:"id,omitempty"`
	CreatedAt   string          `json:"created_at,omitempty"`
	Indicators  []indicatorType `json:"indicators,omitempty"`
}

//NewPackage creates a new package.
func NewPackage() PackageType {
	var obj PackageType
	obj.MessageType = "package"
	obj.ID = createID("package")
	obj.CreatedAt = getCurrentTime()
	return obj
}

//NewIndicatorPackage creates a new indicator package.
func (p *PackageType) NewIndicatorPackage() *indicatorType {
	i := newIndicator()
	pos := p.addIndicator(i)
	return &p.Indicators[pos]
}

func (p *PackageType) addIndicator(c indicatorType) int {
	if p.Indicators == nil {
		a := make([]indicatorType, 0)
		p.Indicators = a
	}
	pos := len(p.Indicators)
	p.Indicators = append(p.Indicators, c)
	return pos
}
