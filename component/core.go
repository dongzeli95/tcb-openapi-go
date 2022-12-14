package component

import (
	"github.com/dongzeli95/tcb-openapi-go/context"
	"github.com/dongzeli95/tcb-openapi-go/http"
	"github.com/dongzeli95/tcb-openapi-go/sts"
	"github.com/dongzeli95/tcb-openapi-go/util"
	"net/url"
)

type Core struct {
	ctx    *context.Context
	client *http.Client
}

func NewCore(ctx *context.Context, sts *sts.Sts) *Core {
	return &Core{ctx, http.NewHttpClient(ctx, sts)}
}

//GET
func (c *Core) HttpGetJson(path string, params url.Values, headers ...map[string]string) (body string, err error) {
	//附加数据
	body, err = c.client.HttpGetJson(path, params, headers...)
	if err != nil {
		return
	}
	return
}

//POST
func (c *Core) HttpPostJson(path string, request interface{}, headers ...map[string]string) (body string, err error) {
	//附加数据
	body, err = c.client.HttpPostJson(path, util.JsonEncode(request), headers...)
	if err != nil {
		return
	}
	return
}

//PATCH
func (c *Core) HttpPatchJson(path string, request interface{}, headers ...map[string]string) (body string, err error) {
	//附加数据
	body, err = c.client.HttpPatchJson(path, util.JsonEncode(request), headers...)
	if err != nil {
		return
	}
	return
}

//DELETE
func (c *Core) HttpDeleteJson(path string, request interface{}, headers ...map[string]string) (body string, err error) {
	//附加数据
	body, err = c.client.HttpDeleteJson(path, util.JsonEncode(request), headers...)
	if err != nil {
		return
	}
	return
}
