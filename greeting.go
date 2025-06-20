package greeting2

import (
	"time"
)

func Do() string {
	hour := time.Now().Hour()
	switch {
	case hour >= 4 && hour < 10:
		return "おはよう"
	case hour >= 10 && hour < 17:
		return "こんにちは"
	default:
		return "こんばんは"
	}
}
