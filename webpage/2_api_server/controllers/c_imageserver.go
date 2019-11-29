package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	f "github.com/paulxiong/cervical/webpage/2_api_server/functions"
	m "github.com/paulxiong/cervical/webpage/2_api_server/models"
)

// checkReferer 防盗链功能
func checkReferer(c *gin.Context, rootURLpath string) {
	if m.SystemCfg == nil || m.SystemCfg.RefererEn != 2 {
		return
	}

	referers := m.SystemCfg.Referers2
	ref := c.Request.Referer()
	// 非法请求
	if ref == "" || len(referers) < 1 {
		c.Request.URL.Path = m.SystemCfg.Referer401URL
		return
	}
	//检查referer是否合法, 合法的不做任何操作直接返回
	for _, _ref := range referers {
		if strings.HasPrefix(ref, _ref) {
			// 图片不存在
			ret, err := f.PathExists(c.Request.URL.Path)
			if ret == false || err != nil {
				c.Request.URL.Path = m.SystemCfg.Referer404URL
			}
			return
		}
	}

	c.Request.URL.Path = m.SystemCfg.Referer401URL
	return
}

// ImageAPI 图片服务器API
// @Summary 图片服务器API（图片在线缩放、旋转、裁剪等）
// @Description 图片URL，详细参数请参考 https://github.com/pierrre/imageserver
// @tags API1 文件（不需要认证）
// @Accept  json
// @Produce json
// @Param width query string false "width, 图片宽度，不传值表示原始尺寸，只传width不传height表示按照width等比例缩放"
// @Param height query string false "height, 图片高度，不传值表示原始尺寸，只传height不传width表示按照height等比例缩放"
// @Success 200 {string} json "{"ping": "pong",	"status": 200}"
// @Router /imgs [get]
func ImageAPI(c *gin.Context) {
	imgexpires := 1
	if m.SystemCfg != nil && m.SystemCfg.ImgExpires > 0 {
		imgexpires = m.SystemCfg.ImgExpires
	}

	cfg := f.ImageServerSettings{
		MaxWidth:     4096,
		MaxHeight:    4096,
		ImgDir:       ".",
		Cachedir:     "cache",
		MemCacheSize: int64(128 * (1 << 20)), //128M
		HTTPExpires:  time.Duration(imgexpires) * time.Hour,
	}

	checkReferer(c, "/imgs")
	h := http.StripPrefix("/imgs", f.ImageServer(&cfg))
	h.ServeHTTP(c.Writer, c.Request)
}
