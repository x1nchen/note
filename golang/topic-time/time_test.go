package topic_time

import (
	"testing"
	"time"
)

func TestTimeZone(t *testing.T) {

	now := time.Now()
	// year, mon, day := now.UTC().Date()
	// hour, min, sec := now.UTC().Clock()
	zone, offset := now.Zone()
	t.Log(zone, offset)

	// fmt.Println(now.In(time.UTC).Add(time.Hour*time.Duration(0)).Format("2006/01/02 15:04:05"))
	//
	// //date := time.Date(year, mon, day, hour, min, sec, 0, time.FixedZone("UTC", 1000))
	// //fmt.Println(date)
}
