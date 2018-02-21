package main
 
import (
    "fmt"
    "log"
    "net/rpc/jsonrpc"
)
 
type Params struct {
    Width, Height int
}
 
func main() {
    rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        log.Fatal(err)
    }
    ret := 0
    err = rpc.Call("Rectangle.Area", Params{50, 100}, &ret)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(ret)
    err = rpc.Call("Rectangle.Perimeter", Params{50, 100}, &ret)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(ret)
}
