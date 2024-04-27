package time

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeToTimestamppb time.Time -> timestamppb.Timestamp
func TimeToTimestamppb(tm *time.Time) *timestamppb.Timestamp {
	if tm != nil {
		return timestamppb.New(*tm)
	}
	return nil
}
