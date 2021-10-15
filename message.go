package dingding

// 详情参见 https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq

const (
	MsgTypeText     = "text"     //text 类型
	MsgTypeLink     = "link"     //link 类型
	MsgTypeMarkdown = "markdown" //markdown 类型
)

// Message 普通消息
type Message struct {
	Content  string   `validate:"required"`
	AtPerson []string `json:"atUserIds"`
	AtAll    bool     `json:"isAtAll"`
}

// Link 链接消息
type Link struct {
	Content    string   `json:"text" validate:"required"`       // 要发送的消息， 必填
	Title      string   `json:"title" validate:"required"`      // 标题， 必填
	ContentURL string   `json:"messageUrl" validate:"required"` // 点击消息跳转的URL 必填
	PictureURL string   `json:"picUrl"`                         // 图片 url
	AtPerson   []string `json:"atUserIds"`
	AtAll      bool     `json:"isAtAll"`
}

// Markdown markdown 类型
type Markdown struct {
	Content  string   `json:"text" validate:"required"`  // 要发送的消息， 必填
	Title    string   `json:"title" validate:"required"` // 标题， 必填
	AtPerson []string `json:"atUserIds"`
	AtAll    bool     `json:"isAtAll"`
}

// SimpleMessage push message
type SimpleMessage struct {
	Content string
	Title   string
}
