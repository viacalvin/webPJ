package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/douglasmakey/oauth2-example/handlers"
)

func main() {
    server := &http.Server{
        Addr: fmt.Sprintf(":8000"),
        Handler: handlers.New(),
    }

    log.Printf("HTTP 백엔드 서버 %q", server.Addr)
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        log.Printf("%v", err)
    } else {
        log.Println("서버 종료!")
    }
}
