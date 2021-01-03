package slide_win

import (
	"container/list"
	"fmt"
	"time"
)

// 配置参数可调整
const (
	// 每个时间片的大小
	sliceSize = 1 * time.Millisecond

	// 窗口大小
	winSize = 1 * time.Second

	// 窗口内的请求数限制
	ReqLimit = 5000
)

// 当前窗口内的总请求数
var reqCount uint

// 窗口内的时间片
var winTimeSlices *list.List

type sliceData struct {
	time  time.Time
	count uint
}

func init() {
	winTimeSlices = list.New()
}

func Check() bool {
	now := time.Now()

	// 从窗口左侧去掉过期的时间片
	for {
		slice := winTimeSlices.Back()
		if slice == nil {
			break
		}

		sliceValue := slice.Value.(sliceData)
		if sliceValue.time.Add(sliceSize + winSize).Before(now) {
			winTimeSlices.Remove(slice)
			reqCount -= sliceValue.count
			fmt.Printf("remove slice. time: %s, count: %d\n", sliceValue.time.Format("04:05.000"), sliceValue.count)
		} else {
			break
		}
	}

	// 检查请求数是否超限
	if reqCount >= ReqLimit {
		return false
	}

	// 若当前时间仍在第一个时间片内，则累计到第一个时间片
	slice := winTimeSlices.Front()
	if slice != nil {
		sliceValue := slice.Value.(sliceData)
		if sliceValue.time.Add(sliceSize).After(now) {
			sliceValue.count++
			slice.Value = sliceValue
			reqCount++
			fmt.Printf("modify first slice. time: %s, count: %d\n", sliceValue.time.Format("04:05.000"), sliceValue.count)
			return true
		}
	}

	// 当前时间不在第一个时间片内，创建一个新的时间片并加入到窗口中
	winTimeSlices.PushFront(sliceData{
		time:  now,
		count: 1,
	})
	reqCount++
	fmt.Printf("push new slice. time: %s\n", now.Format("04:05.000"))

	return true
}
