// Code generated by hertz generator.

package account

import (
	"context"
	"fmt"

	account "user/setting/biz/model/sys/account"

	"github.com/cloudwego/hertz/pkg/app"
)

// LoginMethod .
// @router /login [POST]
func LoginMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req account.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	fmt.Println(req.Username)
	fmt.Println(req.Password)
	resp := new(account.LoginResp)

	c.JSON(200, resp)
}

// LogoutMethod .
// @router /logout [POST]
func LogoutMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req account.LogoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(account.LogoutResp)

	c.JSON(200, resp)
}
