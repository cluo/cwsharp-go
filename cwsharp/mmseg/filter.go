package mmseg

var lawlFunc = func(chunks []chunk, count int) int {
	maxLength := float64(0)
	nextCount := 0
	for i := 0; i < count; i++ {
		chunk := chunks[i]
		averageLength := chunk.WordAverageLength()
		if averageLength > maxLength {
			maxLength = averageLength
			nextCount = 0
			chunks[nextCount] = chunk
			nextCount++
		} else if averageLength == maxLength {
			chunks[nextCount] = chunk
			nextCount++
		}
	}
	return nextCount
}

var svwlFunc = func(chunks []chunk, count int) int {
	minVariance := float64(2 ^ 32)
	nextCount := 0
	for i := 0; i < count; i++ {
		chunk := chunks[i]
		variance := chunk.Variance()
		if variance < minVariance {
			minVariance = variance
			nextCount = 0
			chunks[nextCount] = chunk
			nextCount++
		} else if variance == minVariance {
			chunks[nextCount] = chunk
			nextCount++
		}
	}
	return nextCount
}

var lsdmfocwFunc = func(chunks []chunk, count int) int {
	maxDegree := float64(0)
	nextCount := 0
	for i := 0; i < count; i++ {
		chunk := chunks[i]
		degree := chunk.Degree()
		if degree > maxDegree {
			nextCount = 0
			maxDegree = degree
			chunks[nextCount] = chunk
			nextCount++
		} else if degree == maxDegree {
			chunks[nextCount] = chunk
			nextCount++
		}
	}
	return nextCount
}
