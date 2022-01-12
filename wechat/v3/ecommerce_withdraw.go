package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// 二级商户余额提现API
// 注意：本接口会提交一些敏感信息，需调用 client.V3EncryptText() 进行加密
// Code = 0 is success
// 适用对象：电商平台
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_2.shtml
func (c *ClientV3) V3EcommerceFundWithdraw(ctx context.Context, bm gopay.BodyMap) (*EcommerceFundWithdrawRsp, error) {
	if err := bm.CheckEmptyError(
		"sub_mchid", "out_request_no", "amount"); err != nil {
		return nil, err
	}

	authorization, err := c.authorization(MethodPost, v3EcommerceFundWithdraw, bm)
	if err != nil {
		return nil, err
	}

	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceFundWithdraw, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &EcommerceFundWithdrawRsp{Code: Success, SignInfo: si}
	wxResp.Response = new(EcommerceFundWithdraw)
	if err = json.Unmarshal(bs, wxResp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}

	if res.StatusCode != http.StatusOK {
		wxResp.Code = res.StatusCode
		wxResp.Error = string(bs)
		return wxResp, nil
	}
	return wxResp, c.verifySyncSign(si)
}

// 电商平台提现API
// Code = 0 is success
// 适用对象：电商平台
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_5.shtml
func (c *ClientV3) V3MerchantFundWithdraw(ctx context.Context, bm gopay.BodyMap) (*MerchantFundWithdrawRsp, error) {
	if err := bm.CheckEmptyError(
		"out_request_no", "amount", "account_type"); err != nil {
		return nil, err
	}

	authorization, err := c.authorization(MethodPost, v3MerchantFundWithdraw, bm)
	if err != nil {
		return nil, err
	}

	res, si, bs, err := c.doProdPost(ctx, bm, v3MerchantFundWithdraw, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &MerchantFundWithdrawRsp{Code: Success, SignInfo: si}
	wxResp.Response = new(MerchantFundWithdraw)
	if err = json.Unmarshal(bs, wxResp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}

	if res.StatusCode != http.StatusOK {
		wxResp.Code = res.StatusCode
		wxResp.Error = string(bs)
		return wxResp, nil
	}
	return wxResp, c.verifySyncSign(si)
}

// 电商平台查询提现状态API
// Code = 0 is success
// 适用对象：电商平台
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_6.shtml
func (c *ClientV3) V3MerchantFundWithdrawQuery(ctx context.Context, orderNoType EcommerceWithdrawOrderNoType, orderNo string) (*MerchantFundWithdrawQueryRsp, error) {
	var uri string
	switch orderNoType {
	case EcommerceWithdrawWithdrawId:
		uri = fmt.Sprintf(v3MerchantFundWithdrawQueryByWithdrawId, orderNo)
	case EcommerceWithdrawOutRequestNo:
		uri = fmt.Sprintf(v3MerchantFundWithdrawQueryByOutRequestNo, orderNo)
	default:
		return nil, errors.New("unsupported order number type")
	}

	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &MerchantFundWithdrawQueryRsp{Code: Success, SignInfo: si}
	wxResp.Response = new(MerchantFundWithdrawQuery)
	if err = json.Unmarshal(bs, wxResp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}

	if res.StatusCode != http.StatusOK {
		wxResp.Code = res.StatusCode
		wxResp.Error = string(bs)
		return wxResp, nil
	}
	return wxResp, c.verifySyncSign(si)
}
