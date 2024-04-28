package trans

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func Int32ToPointer(a int32) *int32 {
	return &a
}

// TimeToTimestamppb time.Time -> timestamppb.Timestamp
func TimeToTimestamppb(tm *time.Time) *timestamppb.Timestamp {
	if tm != nil {
		return timestamppb.New(*tm)
	}
	return nil
}
