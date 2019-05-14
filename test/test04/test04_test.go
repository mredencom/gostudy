package test04
//go test -v -run="none" -bench=. -benchtime="3s" -benchmen
//-run="none" 表示测试没有其他干扰
//-bench=. "."的意思包所有的方法.  "BenchmarkSprintf"指定方法
//-benchtime 表示3秒内时间
//-benchmen 使用内存的使用
import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", number)
	}
}

func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
