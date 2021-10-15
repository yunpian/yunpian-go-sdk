package yunpian

import "errors"

// SMSSendRequest - 智能匹配模板发送请求参数(V1)
type SMSSendRequest struct {
	Mobile      string `schema:"mobile,omitempty"`
	Text        string `schema:"text,omitempty"`
	Extend      string `schema:"extend,omitempty"`
	UID         string `schema:"uid,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *SMSSendRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if len(req.Text) == 0 {
		return errors.New("Miss param: text")
	}

	return nil
}

// SMSSendV1Response - 智能匹配模板发送响应(V1)
type SMSSendV1Response struct {
	V1Response
	Result struct {
		Count int     `json:"count"`
		Fee   float64 `json:"fee"`
		SID   int64   `json:"sid"`
	} `json:"result"`
}

// SendV1 - 智能匹配模板发送接口(V1)
func (sms *SMS) SendV1(input *SingleSendRequest) (*SMSSendV1Response, error) {
	if input == nil {
		input = &SingleSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/send.json")
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

	var result SMSSendV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SMSPullStatusV1Response - 获取状态报告响应v1
type SMSPullStatusV1Response struct {
	V1Response
	Status []PullStatusResponse `json:"sms_status"`
}

// PullStatusV1 - 获取状态报告接口(V1)
func (sms *SMS) PullStatusV1(pageSize int) (*SMSPullStatusV1Response, error) {
	input := &PullStatusRequest{PageSize: pageSize}
	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/pull_status.json")
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

	var result SMSPullStatusV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SMSReply -
type SMSReply struct {
	Mobile     string `json:"mobile"`
	Time       string `json:"reply_time"`
	Text       string `json:"text"`
	Extend     string `json:"extend"`
	BaseExtend string `json:"base_extend"`
}

// SMSPullReplyV1Response - 获取回复短信响应(V1)
type SMSPullReplyV1Response struct {
	V1Response
	Reply []SMSReply `json:"sms_reply"`
}

// PullReplyV1 - 获取回复短信接口(V1)
func (sms *SMS) PullReplyV1(pageSize int) (*SMSPullReplyV1Response, error) {
	input := &PullReplyRequest{PageSize: pageSize}
	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/pull_reply.json")
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

	var result SMSPullReplyV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SMSGetReplyV1Request - 查回复的短信请求参数(V1)
type SMSGetReplyV1Request struct {
	StartTime string `schema:"start_time,omitempty"`
	EndTime   string `schema:"end_time,omitempty"`
	PageNum   int    `schema:"page_num,omitempty"`
	PageSize  int    `schema:"page_size,omitempty"`
	Mobile    string `schema:"mobile,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *SMSGetReplyV1Request) Verify() error {
	if len(req.StartTime) == 0 {
		return errors.New("Miss param: start_time")
	}
	if len(req.EndTime) == 0 {
		return errors.New("Miss param: end_time")
	}
	if req.PageNum == 0 {
		return errors.New("Miss param: page_num")
	}
	if req.PageSize == 0 {
		return errors.New("Miss param: page_size")
	}
	return nil
}

// SMSGetReplyV1Response - 查回复的短信响应(V1)
type SMSGetReplyV1Response struct {
	V1Response
	Reply []struct {
		Mobile string `json:"mobile"`
		Text   string `json:"text"`
		Time   string `json:"reply_time"`
	} `json:"sms_reply"`
}

// GetReplyV1 - 查回复的短信接口(V1)
func (sms *SMS) GetReplyV1(input *SMSGetReplyV1Request) (*SMSGetReplyV1Response, error) {
	if input == nil {
		input = &SMSGetReplyV1Request{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/get_reply.json")
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

	var result SMSGetReplyV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SMSGetBlackWorldV1Response - 查屏蔽词响应(V1)
type SMSGetBlackWorldV1Response struct {
	V1Response
	Result struct {
		BlackWorld string `json:"black_word"`
	} `json:"result"`
}

// GetBlackWorldV1 - 查屏蔽词接口(V1)
func (sms *SMS) GetBlackWorldV1(text string) (*SMSGetBlackWorldV1Response, error) {
	input := struct {
		Text string `schema:"text,omitempty"`
	}{
		Text: text,
	}

	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/get_black_word.json")
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

	var result SMSGetBlackWorldV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SMSMultiSendV1Request - 批量个性化发送请求参数(V1)
type SMSMultiSendV1Request struct {
	Mobile      string `schema:"mobile,omitempty"`
	Text        string `schema:"text,omitempty"`
	Extend      string `schema:"extend,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
	UID         string `schema:"uid,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *SMSMultiSendV1Request) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if len(req.Text) == 0 {
		return errors.New("Miss param: text")
	}

	return nil
}

// SMSMutilSendV1Response - 批量个性化发送响应(V1)
type SMSMutilSendV1Response struct {
	V1Response
	Result struct {
		Count int     `json:"count"`
		Fee   float64 `json:"fee"`
		SID   int64   `json:"sid"`
	} `json:"result"`
}

// MutilSendV1 - 批量个性化发送接口(V1)
func (sms *SMS) MutilSendV1(input *SMSMultiSendV1Request) ([]*SMSMutilSendV1Response, error) {
	if input == nil {
		input = &SMSMultiSendV1Request{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/multi_send.json")
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

	var result []*SMSMutilSendV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SMSGetRecordV1Request - 查短信发送记录请求参数(V1)
type SMSGetRecordV1Request struct {
	Mobile    string `schema:"mobile,omitempty"`
	StartTime string `schema:"start_time,omitempty"`
	EndTime   string `schema:"end_time,omitempty"`
	PageNum   int    `schema:"page_num,omitempty"`
	PageSize  int    `schema:"page_size,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *SMSGetRecordV1Request) Verify() error {
	if len(req.StartTime) == 0 {
		return errors.New("Miss param: start_time")
	}
	if len(req.EndTime) == 0 {
		return errors.New("Miss param: end_time")
	}
	if req.PageNum == 0 {
		return errors.New("Miss param: page_num")
	}
	if req.PageSize == 0 {
		return errors.New("Miss param: page_size")
	}
	return nil
}

// SMSGetRecordV1Response - 查短信发送记录响应(V1)
type SMSGetRecordV1Response struct {
	V1Response
	SMS []GetRecordResponse `json:"sms"`
}

// GetRecordV1 - 查短信发送记录接口(V1)
func (sms *SMS) GetRecordV1(input *SMSGetRecordV1Request) (*SMSGetRecordV1Response, error) {
	if input == nil {
		input = &SMSGetRecordV1Request{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/get_record.json")
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

	var result SMSGetRecordV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SMSCountV1Request - 统计短信条数请求参数(V1)
type SMSCountV1Request struct {
	Mobile    string `schema:"mobile,omitempty"`
	StartTime string `schema:"start_time,omitempty"`
	EndTime   string `schema:"end_time,omitempty"`
	PageNum   int    `schema:"page_num,omitempty"`
	PageSize  int    `schema:"page_size,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *SMSCountV1Request) Verify() error {
	if len(req.StartTime) == 0 {
		return errors.New("Miss param: start_time")
	}
	if len(req.EndTime) == 0 {
		return errors.New("Miss param: end_time")
	}
	return nil
}

// SMSCountV1Response - 统计短信条数响应(V1)
type SMSCountV1Response struct {
	V1Response
	Total int `json:"total"`
}

// CountV1 - 统计短信条数接口(V1)
func (sms *SMS) CountV1(input *SMSCountV1Request) (*SMSCountV1Response, error) {
	if input == nil {
		input = &SMSCountV1Request{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := sms.c.newRequest("POST", sms.c.config.smsHost, "/v1/sms/count.json")
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

	var result SMSCountV1Response
	if err = sms.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
