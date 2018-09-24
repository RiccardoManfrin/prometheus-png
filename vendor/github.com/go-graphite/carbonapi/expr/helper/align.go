package helper

import (
	"math"
	"time"
)

// GetBuckets returns amount buckets for timeSeries (defined with startTime, stopTime and step (bucket) size.
func GetBuckets(start, stop, bucketSize int64) int64 {
	return int64(math.Ceil(float64(stop-start) / float64(bucketSize)))
}

// AlignStartToInterval aligns start of serie to interval
func AlignStartToInterval(start, stop, bucketSize int64) int64 {
	for _, v := range []int64{86400, 3600, 60} {
		if bucketSize >= v {
			start -= start % v
			break
		}
	}

	return start
}

// AlignToBucketSize aligns start and stop of serie to specified bucket (step) size
func AlignToBucketSize(start, stop, bucketSize int64) (int64, int64) {
	start = time.Unix(start, 0).Truncate(time.Duration(bucketSize) * time.Second).Unix()
	newStop := time.Unix(stop, 0).Truncate(time.Duration(bucketSize) * time.Second).Unix()

	// check if a partial bucket is needed
	if stop != newStop {
		newStop += bucketSize
	}

	return start, newStop
}
