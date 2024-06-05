package media

import (
	"testing"
)

func TestNewSegment(t *testing.T) {
	var initialSegment Segment
	initialSegment.StartSegment = 0
	initialSegment.EndSegment = 100

	initialSegment.AddNewSegment(25)

	cur := &initialSegment
	idx := 0
	expected := []int{0, 24, 25, 100}
	for cur != nil {
		if cur.StartSegment != int64(expected[idx]) && cur.EndSegment != int64(expected[idx+1]) {
			t.FailNow()
		}
		idx += 2
		cur = cur.NextSegment
	}
}
