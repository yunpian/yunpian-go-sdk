package yunpian

import (
	"errors"
)

// SMS is used to manipulate the sms API
type SMS struct {
	c *Client
}

// SMS is used to return a handle to the SMS APIs
func (c *Client) SMS() *SMS {
	return &SMS{c}
}

// SingleSendRequest - 单条短信发送参数
type SingleSendRequest struct {
	Mobile      string `schema:"mobile"`
	Text        string `schema:"text"`
	Extend      string `schema:"extend"`
	UID         string `schema:"uid"`
	CallbackURL string `schema:"callback_url"`
	Register    bool   `schema:"register"`
}

// Verify used to check the correctness of the request parameters
func (req *SingleSendRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if len(req.Text) == 0 {
		return errors.New("Miss param: text")
	}

	return nil
}

// SingleSendResponse - 单条短信发送响应
type SingleSendResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"msg"`
	Count   int     `json:"count"`
	Fee     float64 `json:"fee"`
	Unit    string  `json:"unit"`
	Mobile  string  `json:"mobile"`
	SID     int64   `json:"sid"`
}

// IsSuccess used to determine whether to send successfully
func (resp *SingleSendResponse) IsSuccess() bool {
	return resp.Code == 0
}

// SingleSend - 单条短信发送接口
func (sms *SMS) SingleSend(input *SingleSendRequest) (*SingleSendResponse, error) {
	if input == nil {
		input = &SingleSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", "/v2/sms/single_send.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := sms.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := sms.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SingleSendResponse
	if err = sms.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BatchSendRequest - 批量发送请求参数
type BatchSendRequest struct {
	Mobile      string `schema:"mobile"`
	Text        string `schema:"text"`
	Extend      string `schema:"extend"`
	CallbackURL string `schema:"callback_url"`
}

// Verify used to check the correctness of the request parameters
func (req *BatchSendRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if len(req.Text) == 0 {
		return errors.New("Miss param: text")
	}

	return nil
}

// BatchSendResponse - 批量发送响应结构
type BatchSendResponse struct {
	TotalCount int                  `json:"total_count"`
	TotalFee   string               `json:"total_fee"`
	Unit       string               `json:"unit"`
	Data       []SingleSendResponse `json:"data"`
}

// BatchSend - 批量发送接口
func (sms *SMS) BatchSend(input *BatchSendRequest) (*BatchSendResponse, error) {
	if input == nil {
		input = &BatchSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", "/v2/sms/batch_send.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := sms.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := sms.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result BatchSendResponse
	if err = sms.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
