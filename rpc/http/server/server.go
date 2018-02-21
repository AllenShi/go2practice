package main
 
import (
    "net/rpc"
    "net/http"
    "log"
)
 
type Params struct {
    Width, Height int
}
 
type Rectangle struct{}
 
func (r *Rectangle) Area(p Params, result *int) error {
    *result = p.Width * p.Height
    return nil
}
 
func (r *Rectangle) Perimeter(p Params, result *int) error {
    *result = (p.Width + p.Height) * 2
    return nil
}
 
func main() {
    rect := new(Rectangle)
    rpc.Register(rect)
    rpc.HandleHTTP()
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
