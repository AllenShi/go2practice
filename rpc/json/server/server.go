package main
 
import (
    "net"
    "log"
    "net/rpc"
    "net/rpc/jsonrpc"
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
 
func chkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
 
func main() {
    rect := new(Rectangle)
    rpc.Register(rect)
    tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
    chkError(err)
    tcplisten, err := net.ListenTCP("tcp", tcpaddr)
    chkError(err)
    for {
        conn, err3 := tcplisten.Accept()
        if err3 != nil {
            continue
        }
        go jsonrpc.ServeConn(conn)
    }
}
