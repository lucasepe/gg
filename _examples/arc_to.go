// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import "github.com/lucasepe/gg"

func main() {
	dc := gg.NewContext(200, 200)
	dc.SetHexColor("fafafa")
	dc.Clear()

	dc.MoveTo(40, 40)
	dc.LineTo(100, 40)
	dc.ArcTo(150, 40, 150, 90, 50)
	dc.LineTo(150, 120)

	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)
	dc.Stroke()

	dc.SavePNG("arc-to.png")
}
