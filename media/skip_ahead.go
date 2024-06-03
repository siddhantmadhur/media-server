package media

import (
	"errors"
	"fmt"
	"time"
)

func (f *Ffmpeg) SkipTo(playbackTime int64) error {
	f.SegmentBuffer.Lock.Lock()
	defer f.SegmentBuffer.Lock.Unlock()

	cur := f.SegmentBuffer
	segment := playbackTime / 2
	if doesSegmentExist(f, segment) {
		return nil
	}

	if int64(f.LastSegment)+5 > segment {
		for !doesSegmentExist(f, segment) {
		}
		return nil
	}

	for cur != nil {
		fmt.Printf("Skipping ahead...\n")
		if cur.EndSegment > segment {
			// Add to linked list
			f.Stop()
			var newSegment Segment
			newSegment.EndSegment = cur.EndSegment
			cur.EndSegment = int64(f.LastSegment)
			newSegment.StartSegment = cur.EndSegment
			newSegment.PrevSegment = cur
			newSegment.NextSegment = cur.NextSegment
			cur.NextSegment = &newSegment

			f.CurrentPlaybackSecond = playbackTime
			f.StopPlaybackSecond = newSegment.EndSegment * 2
			fmt.Printf("starting at %d, ending at %d", f.CurrentPlaybackSecond, f.StopPlaybackSecond)
			f.Start()
			time.Sleep(time.Second)

			return nil

		}

		cur = cur.NextSegment
	}

	return errors.New("Could not skip ahead")

}
