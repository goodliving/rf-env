package main

import (
    "fmt"
    "net/http"
    "goodliving/rf-env/routers"
    "goodliving/rf-env/pkg/setting"
)

func main() {
    router := routers.InitRouter()
    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}