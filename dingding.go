package dingding

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// DingAPIURL api 地址
	DingAPIURL = `https://oapi.dingtalk.com/robot/send?access_token=`
)

// Result 发送结果
// Success true 成功，否则失败
// ErrMsg 错误信息，如果是钉钉接口错误，会返回钉钉的错误信息，否则返回内部err信息
// ErrCode 钉钉返回的错误码
type Result struct {
	Success bool
	ErrMsg  string `json:"errmsg"`  // ErrMsg 错误信息
	ErrCode int    `json:"errcode"` // 错误码
}

// Group 钉钉组
type Group struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

// Ding 钉钉消息发送实体
type Ding struct {
	AccessToken string // token
}

// NewDing new 一个没有队列的ding
func NewDing(token string) *Ding {
	return &Ding{AccessToken: token}
}

// SendMessage 发送普通文本消息
func (ding Ding) SendMessage(message Message) Result {
	return ding.Send(message)
}

// SendLink 发送link类型消息
func (ding Ding) SendLink(message Link) Result {
	return ding.Send(message)
}

// SendMarkdown 发送markdown格式消息
func (ding Ding) SendMarkdown(message Markdown) Result {
	return ding.Send(message)
}

// Send 发送消息
func (ding Ding) Send(message interface{}) Result {
	if ding.AccessToken == "" {
		return Result{ErrMsg: "access token is required"}
	}

	// 检查必填项目
	if err := validator.New().Struct(message); err != nil {
		return Result{ErrMsg: "field valid error: " + err.Error()}
	}

	var paramsMap map[string]interface{}

	switch message.(type) {
	case *Message:
		paramsMap = convertMessage(*message.(*Message))
	case Message:
		paramsMap = convertMessage(message.(Message))
	case Link:
		paramsMap = convertLink(message.(Link))
	case *Link:
		paramsMap = convertLink(*message.(*Link))
	case Markdown:
		paramsMap = convertMarkdown(message.(Markdown))
	case *Markdown:
		paramsMap = convertMarkdown(*message.(*Markdown))
	default:
		return Result{ErrMsg: "not support message type"}
	}

	buf, err := json.Marshal(paramsMap)
	if err != nil {
		return Result{ErrMsg: "marshal message error:" + err.Error()}
	}

	return postMessage(DingAPIURL+ding.AccessToken, string(buf))
}

func convertMessage(m Message) map[string]interface{} {
	var paramsMap = make(map[string]interface{})
	paramsMap["msgtype"] = MsgTypeText
	paramsMap[MsgTypeText] = map[string]string{"content": m.Content}
	paramsMap["at"] = map[string]interface{}{"atMobiles": m.AtPerson, "isAtAll": m.AtAll}
	return paramsMap
}

func convertLink(m Link) map[string]interface{} {
	var paramsMap = make(map[string]interface{})
	paramsMap["msgtype"] = MsgTypeLink
	paramsMap[MsgTypeLink] = map[string]string{MsgTypeText: m.Content, "title": m.Title, "picUrl": m.PictureURL, "messageUrl": m.ContentURL}
	if m.AtAll {
		paramsMap["at"] = map[string]interface{}{"isAtAll": true, "atUserIds": nil, "atMobiles": nil}
	}
	return paramsMap
}

func convertMarkdown(m Markdown) map[string]interface{} {
	var paramsMap = make(map[string]interface{})
	paramsMap["msgtype"] = MsgTypeMarkdown
	paramsMap[MsgTypeMarkdown] = map[string]string{MsgTypeText: m.Content, "title": m.Title}
	if m.AtAll {
		paramsMap["at"] = map[string]interface{}{"isAtAll": true, "atUserIds": nil, "atMobiles": nil}
	}
	return paramsMap
}

func postMessage(url string, message string) Result {
	var result Result

	resp, err := http.Post(url, "application/json", strings.NewReader(message))
	if err != nil {
		result.ErrMsg = "post data to api error:" + err.Error()
		return result
	}

	//log.Println("message:", message)

	defer resp.Body.Close()
	var content []byte
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		result.ErrMsg = "read http response body error:" + err.Error()
		return result
	}

	//log.Println("response result:", string(content))
	if err = json.Unmarshal(content, &result); err != nil {
		result.ErrMsg = "unmarshal http response body error:" + err.Error()
		return result
	}

	if result.ErrCode == 0 {
		result.Success = true
	}

	return result
}
