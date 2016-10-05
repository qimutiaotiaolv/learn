package util

import (
	"fmt"
	"math/rand"
	"time"
)

// /**
//  *生成 start <= rendom < end
//  */
// func RandomInteger(start, end int) int {
// 	randomer := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	randomer.Intn(n)
// }

func NewSecurity(count int) string {
	randomer := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result = ""
	for i := 0; i != count; i++ {
		num := randomer.Intn(11)
		result = fmt.Sprintf("%s%d", result, num)
	}
	return result
}
