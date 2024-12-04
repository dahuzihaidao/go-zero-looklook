package tool

import (
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = r.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + r.Intn(scope))
	}
	return string(result)
}
