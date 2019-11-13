package uuid2

import (
	"fmt"
	"testing"
)

/**********************************************
** @Des: snowflake_test
** @Author: zhangxueyuan 
** @Date:   2019-11-13 16:27:16
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-11-13 16:27:16
***********************************************/

func TestGenerate(t *testing.T) {

	worker, err := NewSnowflake(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64)
	count := 10000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := worker.Generate()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++  {
		id := <- ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			t.Error("ID is not unique!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All", count, "snowflake ID Get successed!")
}


func BenchmarkGenerate(b *testing.B) {

	worker, err := NewSnowflake(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		worker.Generate()
	}
	b.StopTimer()
}