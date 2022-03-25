package alipay

import (
	"context"
	"github.com/go-pay/gopay"
)

// CloudSaleApiPay 云售卖单次代扣
func (a *Client) CloudSaleApiPay(ctx context.Context, bm gopay.BodyMap) (aliRsp CloudSaleApiPayRsp, err error) {
	if err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject", "user_id"); err != nil {
		return
	}
	_, err = a.doAndParse(ctx, bm, &aliRsp, "cloudsale_api_pay", true)
	return
}

// CloudsaleApiPayCancel 云售卖取消代扣
func (a *Client) CloudsaleApiPayCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp CloudsaleApiPayCancelRsp, err error) {
	if err = bm.CheckEmptyError("out_trade_no"); err != nil {
		return
	}
	_, err = a.doAndParse(ctx, bm, &aliRsp, "cloudsale_api_pay_cancel", true)
	return
}

// CloudsaleApiRefund 云售卖退款
func (a *Client) CloudsaleApiRefund(ctx context.Context, bm gopay.BodyMap) (aliRsp CloudsaleApiRefundRsp, err error) {
	if err = bm.CheckEmptyError("out_trade_no", "refund_amount"); err != nil {
		return
	}
	_, err = a.doAndParse(ctx, bm, &aliRsp, "cloudsale_api_refund", true)
	return
}

// CloudsaleApiPayQuery 云售卖查询订单信息
func (a *Client) CloudsaleApiPayQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp CloudsaleApiPayQueryRsp, err error) {
	if err = bm.CheckEmptyError("out_trade_no"); err != nil {
		return
	}
	_, err = a.doAndParse(ctx, bm, &aliRsp, "cloudsale_api_pay_query", true)
	return
}
