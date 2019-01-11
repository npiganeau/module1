// Copyright 2019 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package main

import (
	"github.com/npiganeau/module1/subpackg"
	"github.com/npiganeau/module2"
)

func main() {
	module2.MyFunc(subpackg.SPType{
		SomeField: "My field",
	})
}
