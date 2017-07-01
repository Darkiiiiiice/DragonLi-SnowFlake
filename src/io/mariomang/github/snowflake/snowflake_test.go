package snowflake

import (
	"fmt"
	"testing"
	"time"
)

func TestNewSnowFlake(t *testing.T) {
	fmt.Println("Testing...")
}

func BenchmarkNewSnowFlake(b *testing.B) {
	snow := NewSnowFlake(1, 1)
	for i := 0; i < b.N; i++ {
		snow.GetID()
	}
	s := time.Now()
	fmt.Println(s)
	fmt.Println(s.UnixNano())
	fmt.Println(s.UnixNano() / 1e6)
	e := s.Add(time.Duration(time.Second))
	fmt.Println(e)
	fmt.Println(e.UnixNano())
	fmt.Println(e.UnixNano() / 1e6)

}
