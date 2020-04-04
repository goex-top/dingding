package dingding

import (
	"container/list"
	"sync"
	"time"
)

// DingQueue 用queue 方式发送消息
// 会发送 markdown 类型消息
type DingQueue struct {
	AccessToken string
	ding        Ding
	Interval    uint       // 发送间隔s，最小为1
	Limit       uint       // 每次发送消息限制，0 无限制，到达时间则发送队列所有消息，大于1则到时间发送最大Limit数量的消息
	Title       string     // 摘要
	messages    *list.List // 消息队列
	lock        sync.Mutex
}

// NewQueue 创建一个队列
func NewQueue(token, title string, interval, limit uint) *DingQueue {
	dingQueue := &DingQueue{
		AccessToken: token,
		Title:       title,
		Interval:    interval,
		Limit:       limit,
	}
	dingQueue.Init()
	return dingQueue
}

// Init 初始化 DingQueue
func (ding *DingQueue) Init() {
	ding.ding.AccessToken = ding.AccessToken
	ding.messages = list.New()
	if ding.Interval == 0 {
		ding.Interval = 1
	}
}

// Push push 消息到队列
func (ding *DingQueue) Push(message string) {
	defer ding.lock.Unlock()
	ding.lock.Lock()
	ding.messages.PushBack(SimpleMessage{Title: ding.Title, Content: message})
}

// PushWithTitle push 消息到队列
func (ding *DingQueue) PushWithTitle(title, message string) {
	defer ding.lock.Unlock()
	ding.lock.Lock()
	if title == "" {
		title = ding.Title
	}

	ding.messages.PushBack(SimpleMessage{Title: title, Content: message})
}

// PushMessage push 消息到队列
func (ding *DingQueue) PushMessage(m SimpleMessage) {
	defer ding.lock.Unlock()
	ding.lock.Lock()
	ding.messages.PushBack(m)
}

// Start 开始工作
func (ding *DingQueue) Start() {
	sendQueueMessage(ding)
	timer := time.NewTicker(time.Second * time.Duration(ding.Interval))
	for {
		select {
		case <-timer.C:
			sendQueueMessage(ding)
		}
	}
}

func sendQueueMessage(ding *DingQueue) {
	defer ding.lock.Unlock()
	ding.lock.Lock()
	title := ding.Title
	msg := ""
	if ding.Limit == 0 {
		for {
			m := ding.messages.Front()
			if m == nil {
				break
			}
			ding.messages.Remove(m)
			switch m.Value.(type) {
			case SimpleMessage:
				v := m.Value.(SimpleMessage)
				msg += v.Content + "\n\n"

			case string:
				msg += m.Value.(string) + "\n\n"
			}

		}
	} else {
	label:
		for i := uint(0); i < ding.Limit; i++ {
			for {
				m := ding.messages.Front()

				if m == nil {
					break label
				}
				ding.messages.Remove(m)
				switch m.Value.(type) {
				case SimpleMessage:
					v := m.Value.(SimpleMessage)
					msg += v.Content + "\n\n"
				case string:
					msg += m.Value.(string) + "\n\n"
				}
			}
		}
	}

	if msg != "" {
		go func() {
			r := ding.ding.Send(Markdown{Title: title, Content: msg})
			if !r.Success {
				//log.Println("err:" + r.ErrMsg)
				NewDing(ding.ding.AccessToken).Send("消息太长，请通过其他途径查看，比如邮件")
			}
		}()
	}
}
