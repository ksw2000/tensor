package dense

import "gorgonia.org/shapes"

// if dims = 2 and axis -1 it returns the last dimension. In this case 1
func resolveAxis(axis int, dims int) int {
	res := axis % dims
	if (res < 0 && dims > 0) || (res > 0 && dims < 0) {
		return res + dims
	}

	return res
}

func elimInnermostOutermost(a, b shapes.Shape) shapes.Shape {
	a2 := a.Clone()
	return append(a2[:len(a)-1], b[1:]...)
}

func largestShape(shps ...shapes.Shape) shapes.Shape {
	var maxShape shapes.Shape
	for _, s := range shps {
		if s.TotalSize() > maxShape.TotalSize() {
			maxShape = s
		}
	}
	return maxShape
}
