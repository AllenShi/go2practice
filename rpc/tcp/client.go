package main

import (
  "net/rpc"
  "fmt"
  "log"
)

type Params struct {
  Width, Height int;
}

func main() {
  rpc, err := rpc.Dial("tcp", "127.0.0.1:8080")
  if err != nil {
    log.Fatal(err)
  }

  ret := 0
  err = rpc.Call("Rectangle.Area", Params{50, 100}, &ret)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(err)
  
  err = rpc.Call("Rectangle.Perimeter", Params{50, 100}, &ret)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ret)
}
