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
	Mobile      string `schema:"mobile,omitempty"`
	Text        string `schema:"text,omitempty"`
	Extend      string `schema:"extend,omitempty"`
	UID         string `schema:"uid,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
	Register    bool   `schema:"register,omitempty"`
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
	Mobile      string `schema:"mobile,omitempty"`
	Text        string `schema:"text,omitempty"`
	Extend      string `schema:"extend,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
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

// MultiSendRequest - 批量个性化发送请求参数
type MultiSendRequest struct {
	Mobile      string `schema:"mobile,omitempty"`
	Text        string `schema:"text,omitempty"`
	Extend      string `schema:"extend,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *MultiSendRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if len(req.Text) == 0 {
		return errors.New("Miss param: text")
	}

	return nil
}

// MultiSendResponse - 批量个性化发送响应
type MultiSendResponse struct {
	TotalCount int                  `json:"total_count"`
	TotalFee   string               `json:"total_fee"`
	Unit       string               `json:"unit"`
	Data       []SingleSendResponse `json:"data"`
}

// MultiSend - 批量个性化发送接口
func (sms *SMS) MultiSend(input *MultiSendRequest) (*MultiSendResponse, error) {
	if input == nil {
		input = &MultiSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", "/v2/sms/multi_send.json")
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

	var result MultiSendResponse
	if err = sms.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TPLSingleSendRequest - 指定模板单发请求参数
type TPLSingleSendRequest struct {
	Mobile   string `schema:"mobile,omitempty"`
	TPLID    int64  `schema:"tpl_id,omitempty"`
	TPLValue string `schema:"tpl_value,omitempty"`
	Extend   string `schema:"extend,omitempty"`
	UID      string `schema:"uid,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLSingleSendRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if req.TPLID == 0 {
		return errors.New("Miss param: tpl_id")
	}
	if len(req.TPLValue) == 0 {
		return errors.New("Miss param: tpl_value")
	}

	return nil
}

// TPLSingleSend - 指定模板单发接口
func (sms *SMS) TPLSingleSend(input *TPLSingleSendRequest) (*SingleSendResponse, error) {
	if input == nil {
		input = &TPLSingleSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", "/v2/sms/tpl_single_send.json")
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

// TPLBatchSendRequest - 指定模板群发请求参数
type TPLBatchSendRequest struct {
	Mobile   string `schema:"mobile,omitempty"`
	TPLID    int64  `schema:"tpl_id,omitempty"`
	TPLValue string `schema:"tpl_value,omitempty"`
	Extend   string `schema:"extend,omitempty"`
	UID      string `schema:"uid,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLBatchSendRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if req.TPLID == 0 {
		return errors.New("Miss param: tpl_id")
	}
	if len(req.TPLValue) == 0 {
		return errors.New("Miss param: tpl_value")
	}

	return nil
}

// TPLBatchSend - 指定模板群发
func (sms *SMS) TPLBatchSend(input *TPLBatchSendRequest) (*BatchSendResponse, error) {
	if input == nil {
		input = &TPLBatchSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", "/v2/sms/tpl_batch_send.json")
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

// PullStatusRequest - 获取状态报告请求参数
type PullStatusRequest struct {
	PageSize int `schema:"page_size,omitempty"`
}

// PullStatusResponse - 获取状态报告响应
type PullStatusResponse struct {
	SID             int64  `json:"sid"`
	UID             string `json:"uid"`
	UserReceiveTime string `json:"user_receive_time"`
	ErrorMessage    string `json:"error_msg"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
}

// PullStatus - 获取状态报告接口
func (sms *SMS) PullStatus(input *PullStatusRequest) ([]*PullStatusResponse, error) {
	if input == nil {
		input = &PullStatusRequest{}
	}

	r := sms.c.newRequest("POST", "/v2/sms/pull_status.json")
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

	var result []*PullStatusResponse
	if err = sms.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// PullReplyRequest - 获取回复短信请求参数
type PullReplyRequest struct {
	PageSize int `schema:"page_size,omitempty"`
}

// PullReplyResponse - 获取回复短信响应
type PullReplyResponse struct {
	ID         string `json:"id"`
	Mobile     string `json:"mobile"`
	Text       string `json:"text"`
	ReplyTime  string `json:"reply_time"`
	Extend     string `json:"extend"`
	BaseExtend string `json:"base_extend"`
	Sign       string `json:"_sign"`
}

// PullReply - 获取回复短信接口
func (sms *SMS) PullReply(input *PullReplyRequest) ([]*PullReplyResponse, error) {
	if input == nil {
		input = &PullReplyRequest{}
	}

	r := sms.c.newRequest("POST", "/v2/sms/pull_reply.json")
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

	var result []*PullReplyResponse
	if err = sms.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetRecordRequest - 查短信发送记录请求参数
type GetRecordRequest struct {
	Mobile    string `schema:"mobile,omitempty"`
	StartTime string `schema:"start_time,omitempty"`
	EndTime   string `schema:"end_time,omitempty"`
	PageNum   int    `schema:"page_num,omitempty"`
	PageSize  int    `schema:"page_size,omitempty"`
	Type      string `schema:"type,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *GetRecordRequest) Verify() error {
	if len(req.StartTime) == 0 {
		return errors.New("Miss param: start_time")
	}
	if len(req.EndTime) == 0 {
		return errors.New("Miss param: end_time")
	}
	return nil
}

// GetRecordResponse - 查短信发送记录响应
type GetRecordResponse struct {
	SID             string  `json:"sid"`
	Mobile          string  `json:"mobile"`
	SendTime        string  `json:"send_time"`
	Text            string  `json:"text"`
	SendStatus      string  `json:"send_status"`
	ReportStatus    string  `json:"report_status"`
	Fee             float64 `json:"fee"`
	UserReceiveTime string  `json:"user_receive_time"`
	ErrorMessage    string  `json:"error_msg"`
	UID             string  `json:"uid"`
}

// GetRecord - 查短信发送记录接口
func (sms *SMS) GetRecord(input *GetRecordRequest) ([]*GetRecordResponse, error) {
	if input == nil {
		input = &GetRecordRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", "/v2/sms/get_record.json")
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

	var result []*GetRecordResponse
	if err = sms.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetTotalFeeRequest - 日账单导出请求参数
type GetTotalFeeRequest struct {
	Date string `schema:"date"`
}

// Verify used to check the correctness of the request parameters
func (req *GetTotalFeeRequest) Verify() error {
	if len(req.Date) == 0 {
		return errors.New("Miss param: date")
	}
	return nil
}

// GetTotalFeeResponse - 日账单导出响应
type GetTotalFeeResponse struct {
	Count        int    `json:"totalCount"`
	Fee          string `json:"totalFee"`
	SuccessCount int    `json:"totalSuccessCount"`
	FailedCount  int    `json:"totalFailedCount"`
	UnknownCount int    `json:"totalUnknownCount"`
}

// GetTotalFee - 日账单导出接口
func (sms *SMS) GetTotalFee(input *GetTotalFeeRequest) (*GetTotalFeeResponse, error) {
	if input == nil {
		input = &GetTotalFeeRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", "/v2/sms/get_total_fee.json")
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

	var result GetTotalFeeResponse
	if err = sms.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
