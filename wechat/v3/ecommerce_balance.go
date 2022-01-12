package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// 查询二级商户账户日终余额API
// Code = 0 is success
// 适用对象：电商平台
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_2.shtml
func (c *ClientV3) V3EcommerceFundEnddayBalanceQuery(ctx context.Context, subMchid string, bm gopay.BodyMap) (*EcommerceFundEnddayBalanceQueryRsp, error) {
	if err := bm.CheckEmptyError("date"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(v3EcommerceFundEnddayBalanceQuery, subMchid)
	uri = uri + "?" + bm.EncodeURLParams()

	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &EcommerceFundEnddayBalanceQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceFundEnddayBalanceQuery)
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
