package main

import (
  "fmt"
  "net"
  "io"
  "log"
  "bufio"
)

func ListenAndServe(address string) {
  listener, err := net.Listen("tcp", address)
  if err != nil {
    log.Fatal(fmt.Sprintf("listen err: %v", err))
  }
  defer listener.Close()
  log.Println(fmt.Sprintf("bind: %s, start listening...", address))

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatal(fmt.Sprintf("accept err: %v", err))
    }
    go Handle(conn)
  }
}

func Handle(conn net.Conn) {
  reader := bufio.NewReader(conn)
  for {
      msg, err := reader.ReadString('\n')
      if err != nil {
        if err != io.EOF {
          log.Println("connection close")
        } else {
          log.Println(err)
        }
        return
      }
      bytes := []byte(msg)
      conn.Write(bytes)
  }
}

func main() {
  ListenAndServe(":8000")
}
