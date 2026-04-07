// Copyright 2021 oncilla
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yeller

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	"image/png"

	"github.com/nfnt/resize"
)

var oldmanSmiles image.Image = func() image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(rawOldManSmiles))
	m, err := png.Decode(reader)
	if err != nil {
		panic(err)
	}
	return m
}()

// SmileAt creates an image with Abe Simpson smiling at the target.
func SmileAt(target image.Image) image.Image {
	bounds := oldmanSmiles.Bounds()

	smiled := image.NewRGBA(bounds)
	draw.Draw(smiled, bounds, oldmanSmiles, image.Point{}, draw.Src)

	at := scaleDownSmiles(target)
	draw.Draw(smiled, at.Bounds(), at, image.Point{}, draw.Over)

	return smiled
}

func scaleDownSmiles(target image.Image) image.Image {
	s := target.Bounds().Size()
	// Scale proportionally to template size, increased by 75%
	width := 138.0 * 1.75
	height := float64(s.Y) * (width / float64(s.X))
	return resize.Resize(uint(width), uint(height), target, resize.Lanczos3)
}
