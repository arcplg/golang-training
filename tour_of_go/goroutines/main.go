package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Chạy đồng thời: Goroutines thực thi song song nhưng không đồng bộ.
// Chúng sẽ không chờ đợi lẫn nhau mà có thể tiếp tục khi có tài nguyên sẵn có.
func main() {
	go say("world 1")
	say("hello")

	go say("world 2")
	say("thien")
}
