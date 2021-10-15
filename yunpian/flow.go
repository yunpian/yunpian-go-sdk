package yunpian

import (
	"errors"
)

// Flow is used to manipulate the flow API
type Flow struct {
	c *Client
}

// Flow is used to return a handle to the flow APIs
func (c *Client) Flow() *Flow {
	return &Flow{c}
}

// FlowGetPackageRequest - 获取流量包请求参数
type FlowGetPackageRequest struct {
	Carrier string `schema:"carrier,omitempty"`
}

// FlowPackage - 流量包
type FlowPackage struct {
	SN           int64   `json:"sn"`
	CarrierPrice float64 `json:"carrier_price"`
	Discount     float64 `json:"discount"`
	Capacity     int     `json:"capacity"`
	Carrier      int     `json:"carrier"`
	Name         string  `json:"name"`
}

// FlowGetPackageResponse - 获取流量包响应
type FlowGetPackageResponse struct {
	V1Response
	Package []FlowPackage `json:"flow_package"`
}

// GetPackageV1 - 获取流量包V1
func (f *Flow) GetPackageV1(carrier string) (*FlowGetPackageResponse, error) {
	input := &FlowGetPackageRequest{Carrier: carrier}

	r := f.c.newRequest("POST", f.c.config.flowHost, "/v1/flow/get_package.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := f.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := f.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FlowGetPackageResponse
	if err = f.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// FlowRechargeRequest - 充值流量请求参数
type FlowRechargeRequest struct {
	Mobile      string `schema:"mobile,omitempty"`
	SN          int64  `schema:"sn,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *FlowRechargeRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if req.SN == 0 {
		return errors.New("Miss param: sn")
	}
	return nil
}

// FlowRechargeResponse - 充值流量响应
type FlowRechargeResponse struct {
	V1Response
	Result struct {
		Count int     `json:"count"`
		Fee   float64 `json:"fee"`
		SID   string  `json:"sid"`
	}
}

// RechargeV1 - 流量充值v1接口
func (f *Flow) RechargeV1(input *FlowRechargeRequest) (*FlowRechargeResponse, error) {
	if input == nil {
		input = &FlowRechargeRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := f.c.newRequest("POST", f.c.config.flowHost, "/v1/flow/recharge.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := f.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := f.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FlowRechargeResponse
	if err = f.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// FlowStatus - 流量充值状态
type FlowStatus struct {
	SID             string `json:"sid"`
	UserReceiveTime string `json:"user_receive_time"`
	ErrorMessage    string `json:"error_msg"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	SN              int64  `json:"sn"`
}

// FlowPullStatusResponse - 流量充值状态响应
type FlowPullStatusResponse struct {
	V1Response
	Status []FlowStatus `json:"flow_status"`
}

// PullStatusV1 - 流量充值状态查询接口v1
func (f *Flow) PullStatusV1(pageSize int) (*FlowPullStatusResponse, error) {
	input := &struct {
		PageSize int `schema:"page_size,omitempty"`
	}{
		PageSize: pageSize,
	}

	r := f.c.newRequest("POST", f.c.config.flowHost, "/v1/flow/pull_status.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := f.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := f.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FlowPullStatusResponse
	if err = f.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
