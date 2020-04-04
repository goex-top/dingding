# 钉钉机器人

```go
package main

import (
	"github.com/beaquant/dingding"
)

func main()  {
    var (
        token   = "xxx"
        keyword = "\nkeywords"
        ding    = dingding.NewDing("xxx")
    )
    
    ding.SendMarkdown(Markdown{Title: "markdown测试", Content: "#### 杭州天气\n" +
        "> 9度，西北风1级，空气良89，相对温度73%\n\n" +
        "> ![screenshot](http://image.jpg)\n" +
        "> ###### 10点20分发布 [天气](http://www.thinkpage.cn/)" + keyword}))
}
```