package concurrency 

import (
    "testing"
    "sync"
    "time"
    "fmt"
)

func TestSyncGroupFunc(t *testing.T) {
  var wg sync.WaitGroup
  
  wg.Add(2)
  go sleepFun(1, &wg)
  go sleepFun(3, &wg)
  wg.Wait()
  fmt.Println("Main goroutine exit")
}

func sleepFun(sec time.Duration, wg *sync.WaitGroup) {
  defer wg.Done()
  time.Sleep(sec * time.Second)
  fmt.Println("Child goroutine exit")
}
