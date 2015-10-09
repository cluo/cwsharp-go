package mmseg

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
