// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "cmd/internal/src"

// hardwiredZero convert constant zero value to hardwired zero register.
func hardwiredZero(f *Func) {
	if !f.Config.hasZeroReg {
		return
	}

	// we should make sure only one zero value in a function, this is required by regalloc.
	var zero *Value
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op == OpRISCV64MOVDconst && v.AuxInt == 0 {
				if zero == nil {
					zero = f.Entry.NewValue0(src.NoXPos, OpHardwiredZero, f.Config.Types.Int)
				}
				v.copyOf(zero)
			}
		}
	}
}
