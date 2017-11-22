package main

import (
	"os"
	"time"

	"fmt"

	"strconv"

	"github.com/mariomang/logsys"
)

func main() {
	fmt.Println(time.Now().String())
	file, err := os.OpenFile("stdout", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	file.Sync()
	//file.Close()

	fmt.Println("Create file success")
	time.Sleep(time.Second * 5)

	//file, err = os.OpenFile("stdout", os.O_RDWR|os.O_APPEND, os.ModePerm)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	logsys.Init(file, logsys.WARN, false)
	s := time.Now()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 2)
		fmt.Printf("times :%d\n", i)
		logsys.Debug("HelloWorld" + strconv.Itoa(i))
		logsys.Warn("HelloWorld" + strconv.Itoa(i))
		logsys.Error("HelloWorld" + strconv.Itoa(i))
		logsys.Info("HelloWorld" + strconv.Itoa(i))
	}

	e := time.Now()
	fmt.Println(e.Sub(s))

	time.Sleep(time.Second * 5)
}
