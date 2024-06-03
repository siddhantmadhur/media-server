package media

import "sync"

type Segment struct {
	StartSegment int64
	EndSegment   int64
	NextSegment  *Segment
	PrevSegment  *Segment
	Lock         *sync.Mutex
}
