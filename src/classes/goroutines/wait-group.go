package goroutines

import (
	"sync"
	"time"

	"github.com/maurodesouza/learning-go/src/utils"
)

const TOTAL_THREADS = 4

func LessonWaitGroup() {
	utils.PrintHeader("WAIT GROUP")

	var waitGroup sync.WaitGroup

	waitGroup.Add(TOTAL_THREADS)

	for index := range TOTAL_THREADS {
		go func() {
			for j := 0; j < TOTAL_THREADS; j++ {
				println("Hello from goroutine", index+1, "| iteration (", j+1, ")")
				time.Sleep(time.Millisecond * 500)
			}
			defer waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
