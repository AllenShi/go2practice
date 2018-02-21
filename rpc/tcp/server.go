package main

import (
  "net"
  "log"
  "net/rpc"
)

type Params struct {
  Width, Height int;
}

type Rectangle struct {}

func (r *Rectangle) Area(p Params, result *int) error {
  *result = p.Width * p.Height
  return nil
}

func (r *Rectangle) Perimeter(p Params, result *int) error {
  *result = (p.Width + p.Height) * 2
   return nil
}

func checkError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  rect := new(Rectangle)
  rpc.Register(rect)

  tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
  checkError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr) 
  checkError(err)

  for {
    conn, err := listener.Accept() 
    if err != nil {
      continue
    }

    go rpc.ServeConn(conn)
  }
  
}
