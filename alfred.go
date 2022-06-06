package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type WorkflowOptions struct {
	Profile string `env:"CHROME_PROFILE"`
	Limit   int    `env:"LIMIT"`
}

func runWithAlfred(wf *aw.Workflow) {
	opts := WorkflowOptions{}
	cfg := aw.NewConfig()
	if err := cfg.To(&opts); err != nil {
		panic(err)
	}

	args := wf.Args()
	logrus.Infof("Args len: %d %v", len(args), args)
	var targetTime time.Time
	if len(args) == 0 {
		targetTime = time.Now()
		item := wf.NewItem("当前时间")
		//timeFormated := now.Format("2006-01-02 15:04:05")
		timeFormated := targetTime.Format("2006-01-02 15:04:05")
		item.Subtitle(timeFormated)
		item.Arg(timeFormated)
		item.Valid(true)
		item1 := wf.NewItem("10位 精确到秒")
		secondFormat := strconv.FormatInt(targetTime.Unix(), 10)
		item1.Subtitle(secondFormat)
		item1.Arg(secondFormat)
		item1.Valid(true)
		item2 := wf.NewItem("13位 精确到毫秒")
		microSecondFormat := strconv.FormatInt(targetTime.UnixMilli(), 10)

		item2.Subtitle(microSecondFormat)
		item2.Arg(microSecondFormat)
		item2.Valid(true)
	} else {
		arg := strings.TrimSpace(args[0])
		//判定纯数字
		match, err2 := regexp.Match(`^[1-9]\d*$`, []byte(arg))
		if err2 != nil {
			panic(err2)
		}
		if match {
			number, _ := strconv.ParseInt(arg, 10, 64)
			var title string
			if len(arg) == 10 {
				targetTime = time.Unix(number, 0)
				title = "10位 精确到秒"
			} else if len(arg) == 13 {
				targetTime = time.UnixMilli(number)
				title = "13位 精确到毫秒"
			} else {
				//todo
			}
			item := wf.NewItem(title)
			item.Subtitle(targetTime.Format("2006-01-02 15:04:05"))
			item.Arg(targetTime.Format("2006-01-02 15:04:05"))
			item.Valid(true)
		} else {
			match, err2 = regexp.Match(`^\d{4}-\d{2}-\d{2}\s{1}\d{2}:\d{2}:\d{2}$`, []byte(arg))
			if match {
				targetTime, _ = time.ParseInLocation("2006-01-02 15:04:05", arg, time.Local)
				item1 := wf.NewItem("10位 精确到秒")
				secondFormat := strconv.FormatInt(targetTime.Unix(), 10)
				item1.Subtitle(secondFormat)
				item1.Arg(secondFormat)
				item1.Valid(true)
				item2 := wf.NewItem("13位 精确到毫秒")
				microSecondFormat := strconv.FormatInt(targetTime.UnixMilli(), 10)
				item2.Subtitle(microSecondFormat)
				item2.Arg(microSecondFormat)
				item2.Valid(true)
			}

		}
		//atoi, err := strconv.Atoi(args[0])
		//if err != nil {
		//	logrus.Warn("input error")
		//	item := wf.NewItem("error")
		//	item.Arg("data error")
		//	item.Valid(true)
		//} else {
		//	logrus.Info("convert success")
		//	targetTime := time.UnixMilli(int64(atoi * 1000))
		//	item := wf.NewItem("10位 精确到秒")
		//	item.Subtitle(targetTime.Format("2006-01-02 15:04:05"))
		//	item.Arg(targetTime.Format("2006-01-02 15:04:05"))
		//	item.Valid(true)
		//}
	}
	wf.SendFeedback()
}
