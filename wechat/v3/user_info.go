package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// 获取用户Unionid
// 文档： https://pay.weixin.qq.com/wiki/doc/wxfacepay/develop/android/payscore.html#%E5%AF%B9%E6%8E%A5%E5%88%B7%E8%84%B8%E6%94%AF%E4%BB%98%E5%88%86
func (c *ClientV3) V3FacemchUsers(ctx context.Context, faceSid string, bm gopay.BodyMap) (*FacemchUsersRsp, error) {
	uri := fmt.Sprintf(v3FacemchUsers, faceSid) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &FacemchUsersRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FacemchUsers)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
