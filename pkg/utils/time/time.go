package time

import (
	"fmt"
	"time"
)

// DiffForHumans 格式化时间
func DiffForHumans(gt *time.Time) string {
	if gt == nil {
		return ""
	}

	n := time.Now().Unix()
	t := gt.Unix()

	var ys int64 = 31536000
	var ds int64 = 86400
	var hs int64 = 3600
	var ms int64 = 60
	var ss int64 = 1

	var rs string

	d := n - t
	switch {
	case d > ys:
		rs = fmt.Sprintf("%d 年前", int(d/ys))
	case d > ds:
		rs = fmt.Sprintf("%d 天前", int(d/ds))
	case d > hs:
		rs = fmt.Sprintf("%d 小时前", int(d/hs))
	case d > ms:
		rs = fmt.Sprintf("%d 分钟前", int(d/ms))
	case d > ss:
		rs = fmt.Sprintf("%d 秒前", int(d/ss))
	default:
		rs = "刚刚"
	}

	return rs
}

func ToDateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:01:05")
}
