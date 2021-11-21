package practice4

import "time"

/*
1.参考 Hystrix 实现一个滑动窗口计数器。
 */

// LimitQueue 队列
var LimitQueue map[string][]int64

// CurrentLimiting 时间滑动窗口限流法
// queueName 滑动窗口名
// count 滑动窗口队列长度
// timeWindow 滑动窗口的时间长度
func CurrentLimiting(queueName string, count uint, timeWindow int64) bool {
	// 获取当前时间时间戳
	currTime := time.Now().Unix()

	// 如果限制队列是 nil 就初始化
	if LimitQueue == nil {
		LimitQueue = make(map[string][]int64)
	}
	// 如果取不到 queueName 对应的滑动窗口，就初始化 queueName 的滑动窗口
	if _, ok := LimitQueue[queueName]; !ok {
		LimitQueue[queueName] = make([]int64, 0)
	}

	// 如果 queueName 滑动窗口未被存满，继续存放当前时间戳，然后返回 true 表示滑动窗口允许接口进去
	if uint(len(LimitQueue[queueName])) < count {
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
		return true
	}

	// 如果队列满了，取出最早访问的时间
	earlyTime := LimitQueue[queueName][0]
	// 如果两时间差值小于等于滑动窗口的时间长度，说明新进来的请求应该被限流应该被抛弃，队列最先进来的请求还没有过期
	if currTime-earlyTime <= timeWindow {
		return false
	} else {
		// 说明最早期的访问应该过期了，抛弃队列头部的请求
		LimitQueue[queueName] = LimitQueue[queueName][1:]
		// 加上最新请求
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
	}
	return true
}
