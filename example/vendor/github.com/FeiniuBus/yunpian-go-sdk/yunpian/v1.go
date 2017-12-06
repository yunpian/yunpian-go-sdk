package yunpian

import (
	"errors"
)

// V1Response - v1接口通用响应
type V1Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// APIV1Response - API v1 接口通用响应
type APIV1Response struct {
	V1Response
	Detail string `json:"detail"`
}

// UserGetV1Response - v1用户获取接口响应
type UserGetV1Response struct {
	V1Response
	User UserResponse `json:"user"`
}

// GetV1 - 获取用户接口v1
func (u *User) GetV1() (*UserGetV1Response, error) {
	r := u.c.newRequest("POST", u.c.config.userHost, "/v1/user/get.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := u.c.encodeFormBody(struct{}{})
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := u.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result UserGetV1Response
	if err = u.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SetV1 - 设置用户接口V1
func (u *User) SetV1(input *UserSetRequest) (*APIV1Response, error) {
	if input == nil {
		input = &UserSetRequest{}
	}

	r := u.c.newRequest("POST", u.c.config.userHost, "/v1/user/set.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := u.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := u.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result APIV1Response
	if err = u.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TPLAddV1Request - 添加模版V1请求参数
type TPLAddV1Request struct {
	Content    string `schema:"tpl_content,omitempty"`
	NotifyType int    `schema:"notify_type,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLAddV1Request) Verify() error {
	if len(req.Content) == 0 {
		return errors.New("Miss param: tpl_content")
	}
	return nil
}

// Template - 模版
type Template struct {
	ID          int64  `json:"tpl_id"`
	Content     string `json:"tpl_content"`
	CheckStatus string `json:"check_status"`
	Reason      string `json:"reason"`
}

// TPLV1Response - 添加模版v1响应
type TPLV1Response struct {
	V1Response
	Template Template `json:"template"`
}

// AddV1 - 添加模版v1接口
func (tpl *TPL) AddV1(input *TPLAddV1Request) (*TPLV1Response, error) {
	if input == nil {
		input = &TPLAddV1Request{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v1/tpl/add.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := tpl.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := tpl.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TPLV1Response
	if err = tpl.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetV1 - 获取模版v1接口
func (tpl *TPL) GetV1(id int64) (*TPLV1Response, error) {
	if id == 0 {
		return nil, errors.New("Miss param: tpl_id")
	}

	input := &TPLGetRequest{ID: id}
	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v1/tpl/get.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := tpl.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := tpl.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TPLV1Response
	if err = tpl.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TPLListV1Response - 获取模版列表v1响应
type TPLListV1Response struct {
	V1Response
	Template []Template `json:"template"`
}

// ListV1 - 获取模版列表v1接口
func (tpl *TPL) ListV1() (*TPLListV1Response, error) {
	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v1/tpl/get.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := tpl.c.encodeFormBody(struct{}{})
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := tpl.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TPLListV1Response
	if err = tpl.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TPLUpdateV1Request - 修改模版v1请求参数
type TPLUpdateV1Request struct {
	ID      int64  `schema:"tpl_id,omitempty"`
	Content string `schema:"tpl_content,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLUpdateV1Request) Verify() error {
	if req.ID == 0 {
		return errors.New("Miss param: tpl_id")
	}
	if len(req.Content) == 0 {
		return errors.New("Miss param: tpl_content")
	}
	return nil
}

// UpdateV1 - 模版更新接口v1
func (tpl *TPL) UpdateV1(input *TPLUpdateV1Request) (*TPLV1Response, error) {
	if input == nil {
		input = &TPLUpdateV1Request{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v1/tpl/update.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := tpl.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := tpl.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TPLV1Response
	if err = tpl.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DelV1 - 模版删除接口v1
func (tpl *TPL) DelV1(id int64) (*APIV1Response, error) {
	if id == 0 {
		return nil, errors.New("Miss param: tpl_id")
	}

	input := &TPLDelRequest{ID: id}
	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v1/tpl/del.json")
	r.header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	reader, err := tpl.c.encodeFormBody(input)
	if err != nil {
		return nil, err
	}
	r.body = reader

	resp, err := tpl.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result APIV1Response
	if err = tpl.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// VoiceSendV1Response - 语音验证码发送响应v1
type VoiceSendV1Response struct {
	V1Response
	Result VoiceSendResponse `json:"result"`
}

// SendV1 - 发送语音验证码v1
func (v *Voice) SendV1(input *VoiceSendRequest) (*VoiceSendV1Response, error) {
	if input == nil {
		input = &VoiceSendRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := v.c.newRequest("POST", v.c.config.voiceHost, "/v1/voice/send.json")
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

	var result VoiceSendV1Response
	if err = v.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// VoicePullStatusV1Response - 获取语音状态报告响应v1
type VoicePullStatusV1Response struct {
	V1Response
	Status []VoicePullStatusResponse `json:"voice_status"`
}

// PullStatusV1 - 获取语音状态报告接口v1
func (v *Voice) PullStatusV1(pageSize int) (*VoicePullStatusV1Response, error) {
	input := &VoicePullStatusRequest{PageSize: pageSize}

	r := v.c.newRequest("POST", v.c.config.voiceHost, "/v1/voice/pull_status.json")
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

	var result VoicePullStatusV1Response
	if err = v.c.decodeJSONBody(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
