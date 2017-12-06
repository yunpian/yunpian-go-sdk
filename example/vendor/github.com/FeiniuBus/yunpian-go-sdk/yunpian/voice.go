package yunpian

import (
	"errors"
)

// Voice is used to manipulate the voice API
type Voice struct {
	c *Client
}

// Voice is used to return a handle to the voice APIs
func (c *Client) Voice() *Voice {
	return &Voice{c}
}

// VoiceSendRequest - 语音验证码请求参数
type VoiceSendRequest struct {
	Mobile      string `schema:"mobile,omitempty"`
	Code        string `schema:"code,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
	DisplayNum  string `schema:"display_num,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *VoiceSendRequest) Verify() error {
	if len(req.Mobile) == 0 {
		return errors.New("Miss param: mobile")
	}
	if len(req.Code) == 0 {
		return errors.New("Miss param: code")
	}

	return nil
}

// VoiceSendResponse - 语音验证码响应
type VoiceSendResponse struct {
	Count int    `json:"count"`
	Fee   int    `json:"fee"`
	SID   string `json:"sid"`
}

// Send - 语音验证码发送接口
func (v *Voice) Send(input *VoiceSendRequest) (*VoiceSendResponse, error) {
	if input == nil {
		input = &VoiceSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := v.c.newRequest("POST", v.c.config.voiceHost, "/v2/voice/send.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := v.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := v.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result VoiceSendResponse
	if err = v.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// VoicePullStatusRequest - 语音状态报告获取请求参数
type VoicePullStatusRequest struct {
	PageSize int `schema:"page_size,omitempty"`
	Type     int `scheam:"type,omitempty"`
}

// VoicePullStatusResponse - 语音状态报告获取响应
type VoicePullStatusResponse struct {
	SID             string `json:"sid"`
	UID             string `json:"uid"`
	UserReceiveTime string `json:"user_receive_time"`
	Duration        int    `json:"duration"`
	ErrorMessage    string `json:"error_msg"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
}

// PullStatus - 语音状态报告获取接口
func (v *Voice) PullStatus(input *VoicePullStatusRequest) ([]*VoicePullStatusResponse, error) {
	if input == nil {
		input = &VoicePullStatusRequest{}
	}

	r := v.c.newRequest("POST", v.c.config.voiceHost, "/v2/voice/pull_status.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := v.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := v.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []*VoicePullStatusResponse
	if err = v.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// VoiceTPLNotifyRequest - 发送语音通知请求参数
type VoiceTPLNotifyRequest struct {
	Mobile      string `schema:"mobile,omitempty"`
	TPLID       int64  `schema:"tpl_id,omitempty"`
	TPLValue    string `schema:"tpl_value,omitempty"`
	CallbackURL string `schema:"callback_url,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *VoiceTPLNotifyRequest) Verify() error {
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

// VoiceTPLNotifyResponse - 发送语音通知响应
type VoiceTPLNotifyResponse struct {
	Count int    `json:"count"`
	Fee   int    `json:"fee"`
	SID   string `json:"sid"`
}

// TPLNotify - 发送语音通知接口
func (v *Voice) TPLNotify(input *VoiceTPLNotifyRequest) (*VoiceTPLNotifyResponse, error) {
	if input == nil {
		input = &VoiceTPLNotifyRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := v.c.newRequest("POST", v.c.config.voiceHost, "/v2/voice/tpl_notify.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := v.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := v.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result VoiceTPLNotifyResponse
	if err = v.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
