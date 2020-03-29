package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"gateway/db"
	"gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type BaseController struct {
}

func (b *BaseController) getUUID(_prefix string) string {
	return fmt.Sprintf("%s-%s", _prefix, strings.Replace(uuid.NewV4().String(), "-", "", -1))
}

func (b *BaseController) bind(c *gin.Context, obj interface{}) error {
	var err error
	contentType := c.Request.Header.Get("Content-Type")
	switch contentType {
	case "":
		err = c.BindJSON(obj)
	case "application/json":
		err = c.BindJSON(obj)
	case "application/json;charset=UTF-8":
		err = c.BindJSON(obj)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(obj, binding.Form)
	default:
		err = errors.New("please input Content-Type.")
	}
	return err
}

// 格式化返回结果
// param[in] c http上下文 /code 返回的状态码 /msg  返回的消息 /data 返回的内容
func (b *BaseController) jsonRet(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, model.Resp{Code: code, Msg: msg, Data: data})
}

func (b *BaseController) genPageArgs(ctx *gin.Context, out *db.PageQueryArgs) error {
	args, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err != nil {
		return err
	}
	if out.Page, err = strconv.ParseInt(args.Get("page"), 10, 64); err != nil {
		return err
	}
	if out.PageSize, err = strconv.ParseInt(args.Get("page_size"), 10, 64); err != nil {
		return err
	}
	tmp := args.Get("query_items")
	if err = json.Unmarshal([]byte(tmp), &out.QueryItems); err != nil {
		return err
	}
	return nil
}

func (b *BaseController) pageUtil(ctx *gin.Context, query *db.PageQuery, data interface{}) {
	var err error
	var args db.PageQueryArgs
	if err = b.genPageArgs(ctx, &args); err != nil {
		b.jsonRet(ctx, 0, "faild", err.Error())
		return
	}

	var total int64
	if total, err = query.GetPageData(&args, data); err != nil {
		b.jsonRet(ctx, 0, "faild", err.Error())
		return
	}
	b.jsonRet(ctx, 1, "success", model.PageData{
		Total:    total,
		Page:     args.Page,
		PageSize: args.PageSize,
		Results:  data,
	})
}
