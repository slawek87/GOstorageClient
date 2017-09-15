package main

import (
	"github.com/slawek87/GOstorageClient/conf"
	"fmt"
)

func main() {
    fmt.Println(conf.Settings.GetSettings("HOST"))
}
