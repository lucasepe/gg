// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import "github.com/lucasepe/gg"

func main() {
	dc := gg.NewContext(500, 500)
	dc.DrawQuad(20, 100, 140, 200, 400, 320, 80, 360)
	dc.SetHexColor("ff0000")
	dc.Fill()
	dc.SavePNG("quad.png")
}
