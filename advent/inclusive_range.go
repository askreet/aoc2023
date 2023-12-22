package advent

type InclusiveRange struct {
	Start int
	End   int
}

func IRange(start, end int) InclusiveRange {
	return InclusiveRange{Start: start, End: end}
}

func (ir *InclusiveRange) Size() int {
	return ir.End - ir.Start + 1
}

type SparseRangeSet struct {
	Ranges []InclusiveRange
}

func (srs *SparseRangeSet) Insert(r InclusiveRange) {
	for idx, _ := range srs.Ranges {
		//           x----r---x
		//  x--srs--x
		if srs.Ranges[idx].End < r.Start ||
			// x----r----x
			//            x--srs--x
			srs.Ranges[idx].Start > r.End {
			continue
		}

		srs.Ranges[idx].Start = min(srs.Ranges[idx].Start, r.Start)
		srs.Ranges[idx].End = max(srs.Ranges[idx].End, r.End)

		replay := srs.Ranges[idx+1:]
		srs.Ranges = srs.Ranges[0 : idx+1]
		for _, rng := range replay {
			srs.Insert(rng)
		}

		return
	}

	srs.Ranges = append(srs.Ranges, r)
}

func (srs *SparseRangeSet) Total() int {
	// There's an edge case where an Insert combines three or more Ranges, but it isn't relevant
	// for my initial use case.
	sum := 0
	for i := range srs.Ranges {
		sum += srs.Ranges[i].Size()
	}

	return sum
}
