package model

import (
	"fmt"
	"time"
)

func Init() {

	defer fmt.Println("ended")

	for i := 1; i <= 10; i++ {

		go func(count int) {
			for i := 0; i < count; i++ {
				str := fmt.Sprintf("%d", i)
				fmt.Println(str)
			}
		}(i)
		time.Sleep(time.Second * 1)

	}
}
