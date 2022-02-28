package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// 查询支持个人业务的银行列表API
//	Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_2.shtml
func (c *ClientV3) V3BanksPersonalBanks(ctx context.Context, bm gopay.BodyMap) (*PersonalBankingRsp, error) {
	uri := v3BanksPersonalBanking + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PersonalBankingRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PersonalBanking)
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

// 查询支持对公业务的银行列表API
//	Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_3.shtml
func (c *ClientV3) V3BanksCorporateBanks(ctx context.Context, bm gopay.BodyMap) (*CorporateBankingRsp, error) {
	uri := v3BanksCorporateBanking + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &CorporateBankingRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(CorporateBanking)
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

// 查询支行列表
//	Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_6.shtml
func (c *ClientV3) V3BanksBranches(ctx context.Context, bankAliasCode string, bm gopay.BodyMap) (*BanksBranchesRsp, error) {
	uri := fmt.Sprintf(v3BanksBranches, bankAliasCode) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &BanksBranchesRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BanksBranches)
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
