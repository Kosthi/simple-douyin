// Code generated by hertz generator.

package relation

import (
	"context"
	"strconv"

	client "simple-douyin/api/biz/client"
	mw "simple-douyin/api/biz/middleware"
	relation "simple-douyin/api/biz/model/relation"

	apiLog "github.com/prometheus/common/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	apiLog.Info("RelationAction")
	var err error
	var req relation.RelationActionRequest
	apiLog.Info("RelationAction2")
	resp := new(relation.RelationActionResponse)
	apiLog.Info("RelationAction3")
	err = c.BindAndValidate(&req)
	apiLog.Info("RelationAction4")
	if err != nil {
		apiLog.Info("RelationAction4-e")
		apiLog.Error("RelationAction", "err", err.Error())
		resp := new(relation.RelationActionResponse)
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	apiLog.Info("RelationAction", "req", req)
	loggedClaims, exist := c.Get("JWT_PAYLOAD")
	if !exist {
		apiLog.Error("RelationAction", "err", "Unauthorized")
		resp.StatusCode = 57006
		// if resp.StatusMsg == nil {
		// 	resp.StatusMsg = new(string)
		// }
		// *resp.StatusMsg = "Unauthorized"
		str := "Unauthorized"
		resp.StatusMsg = &str
		c.JSON(consts.StatusOK, resp)
		return
	}
	userID := int64(loggedClaims.(jwt.MapClaims)[mw.JwtMiddleware.IdentityKey].(float64))
	apiLog.Info("RelationAction", "userID", userID)
	apiLog.Info("RelationAction", "req.ToUserID", req.ToUserID)
	if userID == req.ToUserID {
		apiLog.Error("RelationAction", "err", "can't follow yourself")
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "can't follow yourself"
		c.JSON(consts.StatusOK, resp)
		return
	}
	if req.ToUserID <= 0 {
		apiLog.Error("RelationAction", "err", "invalid to_user_id : ", req.ToUserID)
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "invalid to_user_id"
		c.JSON(consts.StatusOK, resp)
		return
	}
	req.Token = strconv.FormatInt(userID, 10)
	err = client.RelationAction(ctx, &req, resp)
	if err != nil {
		apiLog.Error("RelationAction", "err", err.Error())
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	apiLog.Info("RelationActio7")
	c.JSON(consts.StatusOK, resp)
}

// RelationFollowList .
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationFollowListRequest
	resp := new(relation.RelationFollowListResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	loggedClaims, exist := c.Get("JWT_PAYLOAD")
	if !exist {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "Unauthorized"
		c.JSON(consts.StatusOK, resp)
		return
	}
	userID := int64(loggedClaims.(jwt.MapClaims)[mw.JwtMiddleware.IdentityKey].(float64))
	req.Token = strconv.FormatInt(userID, 10)
	err = client.RelationFollowList(ctx, &req, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationFollowerListRequest
	resp := new(relation.RelationFollowerListResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	loggedClaims, exist := c.Get("JWT_PAYLOAD")
	if !exist {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "Unauthorized"
		c.JSON(consts.StatusOK, resp)
		return
	}
	userID := int64(loggedClaims.(jwt.MapClaims)[mw.JwtMiddleware.IdentityKey].(float64))
	req.Token = strconv.FormatInt(userID, 10)
	err = client.RelationFollowerList(ctx, &req, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// RelationFriendList
// @router /douyin/relation/friend/list/

func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationFriendListRequest
	resp := new(relation.RelationFriendListResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	loggedClaims, exist := c.Get("JWT_PAYLOAD")
	if !exist {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "Unauthorized"
		c.JSON(consts.StatusOK, resp)
		return
	}
	userID := int64(loggedClaims.(jwt.MapClaims)[mw.JwtMiddleware.IdentityKey].(float64))
	req.Token = strconv.FormatInt(userID, 10)
	err = client.RelationFriendList(ctx, &req, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
