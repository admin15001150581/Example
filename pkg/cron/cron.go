package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

var myCron *cron.Cron

func Cron(){
	myCron = cron.New(cron.WithSeconds())//实例化Cron  精确到秒

	defer myCron.Stop().Done()

	_, err :=myCron.AddFunc(fmt.Sprintf("*/%v * * * * ?",3),aaa) //3秒
	if err != nil {
		fmt.Printf(err.Error())
	}
	// start
	myCron.Start()
}

func aaa(){
	fmt.Printf("aaa/")
}