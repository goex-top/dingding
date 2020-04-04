package dingding

import (
	"testing"
	"time"
)

var (
	token   = ""
	ding    = NewDing("")
	keyword = "aaa\n"
)

func TestDing_SendMarkdown(t *testing.T) {
	t.Log(ding.SendMarkdown(Markdown{Title: "markdown测试", Content: keyword + "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n"}))
}

func TestDing_SendMessage(t *testing.T) {
	t.Log(ding.SendMessage(Message{
		Content:  keyword + "测试",
		AtPerson: nil,
		AtAll:    false,
	}))
}

func TestDing_SendLink(t *testing.T) {
	t.Log(ding.SendLink(Link{Title: "link测试", Content: keyword + "测试", ContentURL: "https://www.baidu.com"}))
}

func TestDingQueue(t *testing.T) {
	ding := &DingQueue{Title: "queue测试", Interval: 1, AccessToken: token}
	ding.Init()

	go ding.Start()

	ding.Push(keyword + "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	ding.Push(keyword + "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	ding.Push(keyword + "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")

	ding.Push(keyword + "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	ding.Push(keyword + "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	time.Sleep(10 * time.Second)
}
