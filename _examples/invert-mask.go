// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import "github.com/lucasepe/gg"

func main() {
	dc := gg.NewContext(1024, 1024)
	dc.DrawCircle(512, 512, 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, 0, 1024, 1024)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}
