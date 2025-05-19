package main

import (
	"fmt"
	"time"
)

// timeDemo 时间对象的年月日时分秒
func timeDemo() {
	now := time.Now()

	year := now.Year()
	month := now.Month()
	day := now.Day()

	fmt.Printf("%d-%02d-%02d", year, month, day)

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf(" %02d:%02d:%02d", hour, minute, second)
}

// timezoneDemo 时区示例
func timezoneDemo() {

	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Printf("Beijing Time: %v\n", beijing)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Printf("LoadLocation error: %v\n", err)
		return
	}
	fmt.Printf("New York Time: %v\n", newYork)

	// 加载上海所在的时区
	shanghai, _ := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	fmt.Printf("Shanghai Time: %v\n", shanghai)

	// 加载东京所在的时区
	tokyo, _ := time.LoadLocation("Asia/Tokyo") // UTC+09:00
	fmt.Printf("Tokyo Time: %v\n", tokyo)

	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	timeLocal := time.Date(2025, 05, 19, 16, 28, 0, 0, time.Local)
	fmt.Printf("Local Time: %v\n", timeLocal)

	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Printf("UTC Time: %v\n", timeInUTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	fmt.Printf("sameTimeInBeijing: %v\n", sameTimeInBeijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)
	fmt.Printf("sameTimeInNewYork: %v\n", sameTimeInNewYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Printf("timesAreEqual: %v\n", timesAreEqual) // true

	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Printf("timesAreEqual: %v\n", timesAreEqual) // true

	// 在日常编码过程中使用时间对象的时候一定要注意其时区信息，由于time.LoadLocation依赖系统的时区数据库
	// 在不太确定程序运行环境的情况下建议先定义时区偏移量然后使用time.FixedZone的方式指定时区

}

// timestampDemo 时间戳
func timestampDemo() {
	now := time.Now()

	timestamp := now.Unix() // 秒级时间戳
	fmt.Printf("timestamp: %d\n", timestamp)
	milli := now.UnixMilli() // 毫妙时间戳
	fmt.Printf("milli: %d\n", milli)
	micro := now.UnixMicro() // 微秒时间戳
	fmt.Printf("micro: %d\n", micro)
	nano := now.UnixNano() // 纳秒时间戳
	fmt.Printf("nano: %d\n", nano)
}

// timestamp2Time 将时间戳转换为时间对象
func timestamp2Time() {

	// 获取北京时间所在的东八区的时间对象
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Printf("Beijing Time: %v\n", beijing)

	// 北京时间 2022-02-22 22:22:22.000000022 +0800 CST
	t := time.Date(2022, 02, 22, 22, 22, 22, 22, beijing)
	fmt.Printf("t: %v\n", t)

	var (
		sec   = t.Unix()
		milli = t.UnixMilli()
		micro = t.UnixMicro()
	)
	fmt.Printf("sec: %d, milli: %d, micro: %d\n", sec, milli, micro)

	// 将秒级时间戳转为时间对象（第二个参数为不足1秒的纳秒数）
	timeObj := time.Unix(sec, 22)
	fmt.Printf("timeObj: %v\n", timeObj) // timeObj: 2022-02-22 22:22:22.000000022 +0800 CST
	timeObj = time.UnixMilli(milli)
	fmt.Printf("timeObj: %v\n", timeObj) // timeObj: 2022-02-22 22:22:22 +0800 CST
	timeObj = time.UnixMicro(micro)
	fmt.Printf("timeObj: %v\n", timeObj) // timeObj: 2022-02-22 22:22:22 +0800 CST

}

// formatDemo 时间格式化
func formatDemo() {
	now := time.Now()
	fmt.Printf("now: %v\n", now)

	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan")) // 2025-05-19 18:45:21.214 Mon May
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) // 2025-05-19 06:45:21.214 PM Mon May

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2025/05/19 18:45:21.214
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2025/05/19 18:45:21.214

	// 只格式化时分秒不分
	fmt.Println(now.Format("15:04:05")) // 18:46:40

	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02")) // 2025.05.19
}

// 解析字符串格式的时间
// parseDemo1 指定时区解析时间
func parseDemo1() {

	// 在没有时区指示符的情况下，time.Parse 返回 UTC 时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2025/05/19 18:46:40")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2025-05-19 18:46:40 +0000 UTC

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj, err = time.Parse(time.RFC3339, "2025-05-19T18:46:40+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2025-05-19 18:46:40 +0800 CST

}

// parseDemo2 解析时间
func parseDemo2() {
	now := time.Now()

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2025-05-19 18:46:40", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)          // 2025-05-19 18:46:40 +0800 CST
	fmt.Println(timeObj.Sub(now)) // -25m35.86684s
}

func main() {

	//timeDemo()

	//timezoneDemo()

	//timestampDemo()

	//timestamp2Time()

	//formatDemo()

	//parseDemo1()

	parseDemo2()
}
