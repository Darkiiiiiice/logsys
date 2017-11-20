package main

import (
	"os"
	"time"

	"fmt"

	"strconv"

	"github.com/mariomang/logsys"
)

func main() {
	file, err := os.Create("stdout")
	if err != nil {
		fmt.Println(err.Error())
	}
	logsys.Init(file, logsys.WARN, false)
	s := time.Now()
	for i := 0; i < 10; i++ {
		logsys.Debug("HelloWorld" + strconv.Itoa(i))
		logsys.Warn("HelloWorld" + strconv.Itoa(i))
		logsys.Error("HelloWorld" + strconv.Itoa(i))
		logsys.Info("HelloWorld" + strconv.Itoa(i))
	}

	e := time.Now()
	fmt.Println(e.Sub(s))
}
