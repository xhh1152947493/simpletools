package data

import "time"

// Now 带偏移的当前时间，仅在游戏服使用，基础库可以直接使用time.Now。线上禁止设置偏移时间
func Now() time.Time {
	now := time.Now()
	offsetTs := GTimeOffsetTs.Load()
	return now.Add(time.Duration(offsetTs * int64(time.Second)))
}

func SetGTimeOffset(offset int) {
	GTimeOffsetTs.Store(int64(offset))
}

func Seconds() int64 {
	return Now().Unix()
}
