package handle

import (
	"app/models"
	"app/utils"
	"net/http"
	"strings"
	"time"
)

type homeResponse struct {
	Uri      string `json:"uri"`
	DateTime string `json:"date"`
	Time     int64  `json:"time"`
}

func Home(w http.ResponseWriter, r *http.Request) error {
	// 获取请求地址
	uri := strings.TrimLeft(r.RequestURI, "/")
	if uri == "" {
		return responseSuccess(w, homeResponse{
			Uri:      uri,
			DateTime: utils.GetCurrentDateTime(),
			Time:     time.Now().Unix(),
		})
	}

	// 拿到ID查询数据库
	url, err := models.FindUrlByShortId(uri)
	if err != nil {
		return err
	}

	// 执行跳转
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}
