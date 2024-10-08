package main

import (
    "fmt"
    "time"
)

func main() {
    // 番茄工作法：25分钟学习，5分钟休息
    workDuration := 25 * time.Minute
    breakDuration := 5 * time.Minute

    for {
        fmt.Println("开始学习25分钟！")
        timer(workDuration)

        fmt.Println("休息5分钟！")
        timer(breakDuration)
    }
}

func timer(duration time.Duration) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    endTime := time.Now().Add(duration)
    for range time.Tick(1 * time.Second) {
    remaining := time.Until(endTime)
    if remaining <= 0 {
        fmt.Println("时间到！")
        return
    }
    fmt.Printf("\r剩余时间: %v", remaining.Round(time.Second))
}
 {
	for range time.Tick(1 * time.Second) {
		remaining := time.Until(endTime)
		if remaining <= 0 {
		    fmt.Println("时间到！")
		    return
		}
		fmt.Printf("\r剩余时间: %v", remaining.Round(time.Second))
	    }
    }
}
