package yunpian

import (
	"errors"
)

// Sign is used to manipulate the sign API
type Sign struct {
	c *Client
}

// Sign is used to return a handle to the sign APIs
func (c *Client) Sign() *Sign {
	return &Sign{c}
}

// SignAddRequest - 添加签名请求参数
type SignAddRequest struct {
	Sign         string `schema:"sign,omitempty"`
	Notify       bool   `schema:"notify,omitempty"`
	ApplyVIP     bool   `schema:"apply_vip,omitempty"`
	OnlyGlobal   bool   `schema:"is_only_global,omitempty"`
	IndustryType string `schema:"industry_type,omitempty"`
	LicenseURL   string `schema:"license_url,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *SignAddRequest) Verify() error {
	if len(req.Sign) == 0 {
		return errors.New("Miss param: sign")
	}
	return nil
}

// SignAddEntry is
type SignAddEntry struct {
	ApplyState   string `json:"apply_state"`
	Sign         string `json:"sign"`
	ApplyVIP     bool   `json:"is_apply_vip"`
	OnlyGlobal   bool   `json:"is_only_global"`
	IndustryType string `json:"industry_type"`
}

// SignAddResponse - 添加签名响应
type SignAddResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	Sign    []SignAddEntry `json:"sign"`
}

// Add - 添加签名接口
func (sign *Sign) Add(input *SignAddRequest) (*SignAddResponse, error) {
	if input == nil {
		input = &SignAddRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sign.c.newRequest("POST", sign.c.config.signHost, "/v2/sign/add.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := sign.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := sign.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SignAddResponse
	if err = sign.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SignGetRequest - 获取签名请求参数
type SignGetRequest struct {
	Sign     string `schema:"sign,omitempty"`
	PageNum  int    `schema:"page_num,omitempty"`
	PageSize int    `schema:"page_size,omitempty"`
}

// SignGetEntry is
type SignGetEntry struct {
	Chan         string `json:"chan"`
	CheckStatus  string `json:"check_status"`
	Enabled      bool   `json:"enabled"`
	Extend       string `json:"extend"`
	IndustryType string `json:"industry_type"`
	OnlyGlobal   bool   `json:"only_global"`
	Remark       string `json:"remark"`
	Sign         string `json:"sign"`
	VIP          bool   `json:"vip"`
}

// SignGetResponse - 获取签名响应
type SignGetResponse struct {
	Code  int            `json:"code"`
	Total int            `json:"total"`
	Sign  []SignGetEntry `json:"sign"`
}

// Get - 获取签名接口
func (sign *Sign) Get(input *SignGetRequest) (*SignGetResponse, error) {
	if input == nil {
		input = &SignGetRequest{}
	}

	r := sign.c.newRequest("POST", sign.c.config.signHost, "/v2/sign/get.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := sign.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := sign.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SignGetResponse
	if err = sign.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SignUpdateRequest - 修改签名请求参数
type SignUpdateRequest struct {
	OldSign      string `schema:"old_sign,omitempty"`
	Sign         string `schema:"sign,omitempty"`
	Notify       bool   `schema:"notify,omitempty"`
	ApplyVIP     bool   `schema:"apply_vip,omitempty"`
	OnlyGlobal   bool   `schema:"is_only_global,omitempty"`
	IndustryType string `schema:"industry_type,omitempty"`
	LicenseURL   string `schema:"license_url,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *SignUpdateRequest) Verify() error {
	if len(req.OldSign) == 0 {
		return errors.New("Miss param: old_sign")
	}
	if len(req.Sign) == 0 {
		return errors.New("Miss param: sign")
	}
	return nil
}

// SignUpdateEntry is
type SignUpdateEntry struct {
	ApplyState   string `json:"apply_state"`
	Sign         string `json:"sign"`
	ApplyVIP     bool   `json:"is_apply_vip"`
	OnlyGlobal   bool   `json:"is_only_global"`
	IndustryType string `json:"industry_type"`
}

// SignUpdateResponse - 修改签名响应
type SignUpdateResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"msg"`
	Sign    SignUpdateEntry `json:"sign"`
}

// Update - 修改签名接口
func (sign *Sign) Update(input *SignUpdateRequest) (*SignUpdateResponse, error) {
	if input == nil {
		input = &SignUpdateRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sign.c.newRequest("POST", sign.c.config.signHost, "/v2/sign/update.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := sign.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := sign.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SignUpdateResponse
	if err = sign.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
