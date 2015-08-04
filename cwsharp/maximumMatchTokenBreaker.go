// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

//正向匹配的分词器(4个规则)
type maximumMatchTokenBreaker struct {
	dawg                   *dawg
	whiteSpaceTokenBreaker //anonymouse field
}

type chunkFilter interface {
	Apply(chunks []chunk, cou int) int
}

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

var filters = []func([]chunk, int) int{lawlFunc, svwlFunc, lawlFunc}

func (this *maximumMatchTokenBreaker) Next() Token {
	baseOffset := this.reader.cursor
	code := this.reader.Peek()
	if code == 0 {
		return token_empty
	}
	node, _ := this.dawg.root.Next(code)
	if node == nil || !node.HasChilds() {
		this.reader.Seek(baseOffset)
		return this.whiteSpaceTokenBreaker.Next()
	}
	firstOfNodes := this.MatchedNodes(baseOffset)
	if len(firstOfNodes) == 0 {
		this.reader.Seek(baseOffset)
		return this.whiteSpaceTokenBreaker.Next()
	}
	maxLength := 0
	chunks := make([]chunk, 0)
	for i := len(firstOfNodes) - 1; i >= 0; i-- {
		offset1 := baseOffset + int(firstOfNodes[i].depth) + 1
		secondOfNodes := this.MatchedNodes(offset1)
		if len(secondOfNodes) > 0 {
			for j := len(secondOfNodes) - 1; j >= 0; j-- {
				offset2 := offset1 + int(secondOfNodes[j].depth) + 1
				thirdOfNodes := this.MatchedNodes(offset2)
				if len(thirdOfNodes) > 0 {
					for k := len(thirdOfNodes) - 1; k >= 0; k-- {
						offset3 := offset2 + int(thirdOfNodes[k].depth) + 1
						length := offset3 - baseOffset
						if length >= maxLength {
							maxLength = length
							chunk := chunk{length, []wordPoint{
								wordPoint{baseOffset, offset1 - baseOffset, int(firstOfNodes[i].frequency)},
								wordPoint{offset1, offset2 - offset1, int(secondOfNodes[j].frequency)},
								wordPoint{offset2, offset3 - offset2, int(thirdOfNodes[k].frequency)},
							}}
							chunks = append(chunks, chunk)
							//fmt.Println(chunk.ToString(this.reader))
						}
					}
				} else {
					length := offset2 - baseOffset
					if length > maxLength {
						maxLength = length
						chunk := chunk{length, []wordPoint{
							wordPoint{baseOffset, offset1 - baseOffset, int(firstOfNodes[i].frequency)},
							wordPoint{offset1, offset2 - offset1, int(secondOfNodes[j].frequency)},
						}}
						chunks = append(chunks, chunk)
					}
				}
			}
		} else {
			length := offset1 - baseOffset
			if length > maxLength {
				maxLength = length
				chunk := chunk{length, []wordPoint{
					wordPoint{baseOffset, offset1 - baseOffset, int(firstOfNodes[i].frequency)},
				}}
				chunks = append(chunks, chunk)
			}
		}
	}
	if len(chunks) > 1 {
		count := len(chunks)
		for _, filter := range filters {
			count = filter(chunks, count)
			if count == 1 {
				break
			}
		}
	}
	bestChunk := chunks[0]
	this.reader.Seek(bestChunk.WordPoints[0].Offset)
	stringLength := bestChunk.WordPoints[0].Length
	return NewToken(string(this.reader.ReadCount(stringLength)), TokenType_CJK)
}

func (this *maximumMatchTokenBreaker) MatchedNodes(offset int) []*dawgNode {
	nodes := make([]*dawgNode, 0)
	this.reader.Seek(offset)
	node := this.dawg.root
	for code := this.reader.Read(); code != 0; code = this.reader.Read() {
		node, _ = node.Next(code)
		if node == nil {
			break
		}
		if node.eow {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
