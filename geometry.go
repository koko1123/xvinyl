package main

const MaxInt int = 2147483648

func (from Window) distanceScore(to Window, xdir int, ydir int) int {
	// If `to` is mostly in the correct direction, return the square distance between `from` and `to`
	// Otherwise, return MaxInt
	ydiff := to.Ymid - from.Ymid
	xdiff := to.Xmid - from.Xmid
	// TODO: scale directions by aspect ratio
	dot := xdir*xdiff + ydir*ydiff
	other := xdir*ydiff + -ydir*xdiff
	if other < 0 {
		other = -other
	}

	if dot > other {
		return ydiff*ydiff + xdiff*xdiff
	} else {
		return MaxInt
	}
}


func (w Window) getNextBy (windows *[]Window, xdir int, ydir int) (Window) {
	minSeen := MaxInt
	var ret Window = w
	for _, other := range *windows {
		f := w.distanceScore(other, xdir, ydir)
		if f < minSeen {
			minSeen = f
			ret = other
		}
	}
	return ret
}