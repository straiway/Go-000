package slide_win

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSlideWinCounter(t *testing.T) {
	begin := time.Now()
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Microsecond)
		fmt.Printf(
			"result:%t, i: %d, reqCount: %d, time: %s, timeFromBegin: %d\n",
			Check(), i, reqCount, time.Now().Format("04:05.000"), time.Since(begin).Milliseconds(),
		)
	}
}
