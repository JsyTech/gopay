package alipay

import (
	"fmt"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf(`{"code": "%s","msg": "%s","sub_code": "%s","sub_msg": "%s"}`, e.Code, e.Msg, e.SubCode, e.SubMsg)
}

func (e *ErrorResponse) CodeSucceed() bool {
	return e.Code == "10000"
}

type CloudSaleErrorResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	Sign    string `json:"sign,omitempty"`
}

func (e *CloudSaleErrorResponse) Error() string {
	return fmt.Sprintf(`{"code": "%s","msg": "%s","sub_code": "%s","sub_msg": "%s"}`, e.Code, e.Msg, e.SubCode, e.SubMsg)
}

func (e *CloudSaleErrorResponse) CodeSucceed() bool {
	return e.Code == 10000
}
