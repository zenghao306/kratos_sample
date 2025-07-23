package device

import (
	"github.com/gin-gonic/gin"
	device_v1 "kratos_sample/api/device/v1"
	user_v1 "kratos_sample/api/user/v1"
	"kratos_sample/app/gateway/internal/router/ginplus"
)

// Close 关闭 RedisSubscriber
func (r *DeviceHandler) Close() {
	if r.redisSubscriber != nil {
		r.redisSubscriber.Close()
	}
}

func (r *DeviceHandler) ListDevice(c *gin.Context) {
	var params device_v1.ListDeviceRequest
	err := ginplus.ParseJSON(c, &params)
	if err != nil {
		ginplus.ResponseError(c, err)
		return
	}

	// 调用Kratos[device.service] rpc服务
	reply, err := r.deviceClient.ListDevice(c.Request.Context(), &params)
	if err != nil {
		ginplus.ResponseError(c, err)
		return
	}

	//c.JSON(200, reply)
	pr := ginplus.PaginationResult{
		Page:     params.Page,
		PageSize: params.PageSize,
		Total:    uint32(reply.Total),
	}
	ginplus.ResponsePagePost(c, reply.Data, &pr)
}

func (r *DeviceHandler) AllList(c *gin.Context) {
	var (
		result = make(map[string]interface{})
	)
	u := ginplus.GetUserID(c) //用户ID

	// 调用Kratos[user.service] rpc服务 获取用户分组
	res, err := r.userClient.GetUserGroupId(c, &user_v1.GetUserGroupRequest{UserId: u.User.UserId})
	if err != nil {
		//c.JSON(500, gin.H{"error": "test"})
		ginplus.ResponseError(c, err)
		return
	}

	// 2.根据group_id查询分组信息
	resG, err := r.deviceClient.DeviceGroupAllList(c, &device_v1.IDRequest{Id: res.GroupId})
	if err != nil {
		ginplus.ResponseError(c, err)
		return
	}
	if resG != nil {
		result["result"] = resG.List
	} else {
		result["result"] = []interface{}{}
	}
	ginplus.ResJsonForGeneral(c, ginplus.StatusOK, "success", result)
}
