package main

import (
	"fmt"
	"github.com/yinggaozhen/GoLearn/toy/cron"
	"strconv"
	"time"
)

type MyJob struct {
	cron.Job
}
func (c MyJob) Run() {
	fmt.Println("hello MyJob  --- " + strconv.Itoa(int(time.Now().Unix())))
}

type MyJob2 struct {
	cron.Job
}
func (c MyJob2) Run() {
	fmt.Println("hello MyJob2 --- " + strconv.Itoa(int(time.Now().Unix())))
}

func main() {
	c := cron.New()
	_ = c.AddJob("*/2 * * * * ?", MyJob{})
	_ = c.AddJob("*/5 * * * * ?", MyJob2{})
	c.Start()

	select{}
}

