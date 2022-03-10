package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// 创建支付分订单API
// Code = 0 is success
// 适用对象：服务商
// 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter3_1.shtml
func (c *ClientV3) V3ScorePartnerCreate(ctx context.Context, bm gopay.BodyMap) (wxRsp *ScorePartnerOrderCreateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScorePartnerOrderCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScorePartnerOrderCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScorePartnerOrderCreateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScorePartnerOrderCreate)
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

// 查询支付分订单API
// Code = 0 is success
// 适用对象：服务商
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter3_2.shtml
func (c *ClientV3) V3ScorePartnerQuery(ctx context.Context, orderNoType OrderNoType, serviceId, subMchid, orderNo string) (wxRsp *ScorePartnerOrderQueryRsp, err error) {
	var uri string
	switch orderNoType {
	case OutTradeNo:
		uri = v3ScorePartnerOrderQuery + "?service_id=" + serviceId + "&sub_mchid=" + subMchid + "&out_order_no=" + orderNo
	case QueryId:
		uri = v3ScorePartnerOrderQuery + "?service_id=" + serviceId + "&sub_mchid=" + subMchid + "&query_id=" + orderNo
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

	wxRsp = &ScorePartnerOrderQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScorePartnerOrderQuery)
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

// 取消支付分订单
// Code = 0 is success
// 适用对象：服务商
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter3_3.shtml
func (c *ClientV3) V3ScorePartnerOrderCancel(ctx context.Context, outOrderNo, serviceId, subMchid, reason string) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ScorePartnerOrderCancel, outOrderNo)
	bm := make(gopay.BodyMap)
	bm.Set("service_id", serviceId)
	bm.Set("sub_mchid", subMchid)
	bm.Set("reason", reason)

	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 修改订单金额
// Code = 0 is success
// 适用对象：服务商
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter3_4.shtml
func (c *ClientV3) V3ScorePartnerOrderModify(ctx context.Context, outOrderNo string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ScorePartnerOrderModify, outOrderNo)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 完结支付分订单
// Code = 0 is success
// 适用对象：服务商
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter3_5.shtml
func (c *ClientV3) V3ScorePartnerOrderComplete(ctx context.Context, outOrderNo string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ScorePartnerOrderComplete, outOrderNo)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户发起催收扣款
// Code = 0 is success
// 适用对象：服务商
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter3_6.shtml
func (c *ClientV3) V3ScorePartnerOrderPay(ctx context.Context, outOrderNo, serviceId, subMchid string) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ScorePartnerOrderPay, outOrderNo)
	bm := make(gopay.BodyMap)
	bm.Set("service_id", serviceId)
	bm.Set("sub_mchid", subMchid)

	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 同步订单信息
// Code = 0 is success
// 适用对象：服务商
// 商户文档: https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter3_7.shtml
func (c *ClientV3) V3ScorePartnerOrderSync(ctx context.Context, outOrderNo string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ScorePartnerOrderSync, outOrderNo)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 收付通子商户申请绑定支付分服务
// Code = 0 is success
// 适用对象：服务商
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter9_1.shtml
func (c *ClientV3) V3ScorePartnerBindService(ctx context.Context, bm gopay.BodyMap) (wxRsp *ScorePartnerBindServiceRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScorePartnerBindService, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScorePartnerBindService, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScorePartnerBindServiceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScorePartnerBindService)
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

// 查询收付通子商户服务绑定结果
// Code = 0 is success
// 适用对象：服务商
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/payscore_partner/chapter9_2.shtml
func (c *ClientV3) V3ScorePartnerBindServiceQuery(ctx context.Context, outApplyNo string) (wxRsp *ScorePartnerBindServiceQueryRsp, err error) {
	url := fmt.Sprintf(v3ScorePartnerBindServiceQuery, outApplyNo)
	authorization, err := c.authorization(MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, v3ScorePartnerBindServiceQuery, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScorePartnerBindServiceQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScorePartnerBindService)
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
