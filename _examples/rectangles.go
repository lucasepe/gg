// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import "github.com/lucasepe/gg"

func main() {
	dc := gg.NewContext(1000, 200)
	dc.SetHexColor("fafafa")
	dc.Clear()

	x0, y0 := 10.0, 10.0
	w, h := 240.0, 120.0

	//dc.DrawRectangle(x0, y0, w, h)
	dc.DrawRoundedRectangle(x0, y0, w, h, 0, 0, 0, 0)
	dc.SetRGB255(255, 0, 0)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	x0 = x0 + w + 40

	dc.DrawRoundedRectangle(x0, y0, w, h, 30, 30, 30, 30)
	dc.SetRGB255(255, 0, 0)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	x0 = x0 + w + 40

	dc.DrawRoundedRectangle(x0, y0, w, h, 20, 30.0, 40, 50)
	dc.SetRGB255(255, 0, 0)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	dc.SavePNG("rectangles.png")
}
