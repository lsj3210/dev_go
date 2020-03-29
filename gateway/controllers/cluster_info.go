package controller

import (
	"gateway/db"
	"gateway/model"
	"gateway/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type ClusterInfoController struct {
	BaseController
}

// @Tags cluster
// @Summary 注册或更新集群信息
// @Description 注册或更新集群信息,注册时id请传0""
// @Accept  json
// @Produce  json
// @Param	body  body	model.MsgCluster	true	"用户信息"
// @Success 200 {object} model.Resp
// @Router /v1/cluster/save [post]
func (c *ClusterInfoController) NewOrModify(ctx *gin.Context) {
	var msg model.MsgCluster
	c.bind(ctx, &msg)
	tmp := db.ClusterInfo{
		Name:        msg.Name,
		Description: msg.Description,
	}
	if msg.ID == 0 {
		//新增
		if err := tmp.New(); err != nil {
			c.jsonRet(ctx, 0, "faild", err.Error())
			return
		}
		msg.ID = tmp.ID
		c.jsonRet(ctx, 1, "success", msg)
	} else {
		//更新
		tmp.ID = msg.ID
		tmp.UpdatedAt = time.Now()
		if err := tmp.Modify(); err != nil {
			c.jsonRet(ctx, 0, "faild", err.Error())
			return
		}
		c.jsonRet(ctx, 1, "success", nil)
	}
}

// @Tags cluster
// @Summary 查询集群详情
// @Description 查询集群详情
// @Accept  json
// @Produce  json
// @Param	id  path	int64	true	"id"
// @Success 200 {object} model.Resp
// @Router /v1/cluster/query/{id} [get]
func (c *ClusterInfoController) Query(ctx *gin.Context) {
	id := utils.StringToInt64(ctx.Param("id"))
	var obj db.ClusterInfo
	if err := obj.QueryByID(&id); err != nil {
		c.jsonRet(ctx, 0, "faild", err.Error())
		return
	}
	c.jsonRet(ctx, 1, "success", obj)
}

// @Tags cluster
// @Summary 删除集群
// @Description 删除集群接口
// @Accept  json
// @Produce  json
// @Param	id  path	int64	true	"id"
// @Success 200 {object} model.Resp
// @Router /v1/cluster/delete/{id} [delete]
func (c *ClusterInfoController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	obj := db.ClusterInfo{ID: utils.StringToInt64(id)}
	if err := obj.Del(); err != nil {
		c.jsonRet(ctx, 0, "faild", err.Error())
		return
	}
	c.jsonRet(ctx, 1, "success", nil)
}
