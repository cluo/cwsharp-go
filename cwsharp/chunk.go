// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"math"
)

type wordPoint struct {
	Offset, Length, Freq int
}

type chunk struct {
	Length     int
	WordPoints []wordPoint
}

func (this *chunk) Print(sr *stringReader) string {
	//{0}_{1}_{2}_{3}
	var sb string
	for _, word := range this.WordPoints {
		runes := sr.s[word.Offset : word.Length+word.Offset]
		sb = sb + string(runes) + "_"
	}
	return sb
}

func (this *chunk) Get(index int) wordPoint {
	return this.WordPoints[index]
}

func (this *chunk) WordAverageLength() float64 {
	return float64(this.Length) / float64(len(this.WordPoints))
}

func (this *chunk) Variance() float64 {
	averageLength := this.WordAverageLength()
	sum := float64(0)
	for _, wordPoint := range this.WordPoints {
		sum = sum + (math.Pow(float64(wordPoint.Length)-averageLength, 2))
	}
	return math.Sqrt(sum / float64(len(this.WordPoints)))
}

func (this *chunk) Degree() float64 {
	sum := float64(0)
	for _, wordPoint := range this.WordPoints {
		sum = sum + math.Log10(float64(wordPoint.Freq))
	}
	return sum
}
