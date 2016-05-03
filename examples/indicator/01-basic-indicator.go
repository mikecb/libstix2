// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/mikecb/stix"
)

func main() {
	// Create TAXII Message called tm
	sm := stix.NewPackage()
	i := sm.NewIndicatorPackage()

	i.SetTitle("Malware C2 Indicator 2016")

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
