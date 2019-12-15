# short url

使用go编写的短网址项目

## 接口列表

### 1. 创建短网址

#### 1.1 请求地址

需要POST请求

```
/create
```
#### 1.2 请求参数

|参数名称|类型|是否必填|说明|
|:------|:--------|:---|:---|
|app_key|string|Y|应用秘钥|
|url|string|Y|需要生成的网址|

#### 1.3 响应信息

|参数名称|类型|说明|
|:------|:--------|:---|
|code|int|响应状态**200为正常**|
|msg|string|响应提示信息|
|data|object|响应内容|
| └ data.url|string|生成的网址|
| └ data.short_url|string|短网址|
| └ data.short_id|string|短网址ID|
| └ data.created_at|string|创建时间|

### 2. 返回短网址