package dingding

import (
	"testing"
	"time"
)

var (
	token   = "8684e74f0d02da648dbf576ae91db7ac632714b00f91c8076fbd4e69b1eb10e6"
	ding    = NewDing("8684e74f0d02da648dbf576ae91db7ac632714b00f91c8076fbd4e69b1eb10e6")
	keyword = "multi-acc\n"
)

func TestDing_SendMarkdown(t *testing.T) {
	t.Log(ding.SendMarkdown(Markdown{Title: "markdown测试", Content: "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n" + keyword}))
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
	ding := &DingQueue{Title: "queue测试", Interval: 1, Limit: 1, AccessToken: token}
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
