package yunpian

import "errors"
import "net/http"
import "encoding/json"
import "io/ioutil"
import "fmt"

// TPL is used to manipulate the tpl API
type TPL struct {
	c *Client
}

// TPL is used to return a handle to the TPL APIs
func (c *Client) TPL() *TPL {
	return &TPL{c}
}

// TPLAddRequest - 添加模版请求参数
type TPLAddRequest struct {
	Content    string `schema:"tpl_content,omitempty"`
	NotifyType int    `schema:"notify_type,omitempty"`
	Language   string `schema:"lang,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLAddRequest) Verify() error {
	if len(req.Content) == 0 {
		return errors.New("Miss param: tpl_content")
	}
	return nil
}

// TPLResponse - 模版响应
type TPLResponse struct {
	ID          int64  `json:"tpl_id"`
	Content     string `json:"tpl_content"`
	Status      string `json:"check_status"`
	Reason      string `json:"reason"`
	Language    string `json:"lang"`
	CountryCode string `json:"country_code"`
}

// Add - 添加模版接口
func (tpl *TPL) Add(input *TPLAddRequest) (*TPLResponse, error) {
	if input == nil {
		input = &TPLAddRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v2/tpl/add.json")
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

	var result TPLResponse
	if err = tpl.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TPLGetRequest - 获取模板请求参数
type TPLGetRequest struct {
	ID int64 `schema:"tpl_id,omitempty"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLGetRequest) Verify() error {
	if req.ID == 0 {
		return errors.New("Miss param: tpl_id")
	}
	return nil
}

// Get - 获取模板接口
func (tpl *TPL) Get(input *TPLGetRequest) (*TPLResponse, error) {
	if input == nil {
		input = &TPLGetRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v2/tpl/get.json")
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

	var result TPLResponse
	if err = tpl.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDefault - 取默认模板接口
func (tpl *TPL) GetDefault(input *TPLGetRequest) (*TPLResponse, error) {
	if input == nil {
		input = &TPLGetRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v2/tpl/get_default.json")
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

	var result TPLResponse
	if err = tpl.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// List - 获取模版列表接口
func (tpl *TPL) List() ([]*TPLResponse, error) {
	input := &TPLGetRequest{}
	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v2/tpl/get.json")
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

	var result []*TPLResponse
	if err = tpl.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListDefault - 获取默认模版列表接口
func (tpl *TPL) ListDefault() ([]*TPLResponse, error) {
	input := &TPLGetRequest{}
	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v2/tpl/get_default.json")
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

	var result []*TPLResponse
	if err = tpl.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TPLUpdateRequest - 修改模版请求参数
type TPLUpdateRequest struct {
	ID       int64  `schema:"tpl_id,omitempty"`
	Content  string `schema:"tpl_content,omitempty"`
	Language string `json:"lang,omitempty"`
}

// TPLUpdateResponse - 修改模版响应
type TPLUpdateResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"msg"`
	Template TPLResponse `json:"template"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLUpdateRequest) Verify() error {
	if req.ID == 0 {
		return errors.New("Miss param: tpl_id")
	}
	if len(req.Content) == 0 {
		return errors.New("Miss param: tpl_content")
	}
	return nil
}

// Update - 修改模版接口
func (tpl *TPL) Update(input *TPLUpdateRequest) (*TPLUpdateResponse, error) {
	if input == nil {
		input = &TPLUpdateRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v2/tpl/update.json")
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		var e ErrorResponse
		if err := json.Unmarshal(body, &e); err != nil {
			return nil, err
		}
		return nil, e
	}

	var result TPLResponse
	if err = json.Unmarshal(body, &result); err == nil {
		return &TPLUpdateResponse{
			Code:     0,
			Message:  "ok",
			Template: result,
		}, nil
	}

	var updateResult TPLUpdateResponse
	if err = json.Unmarshal(body, &updateResult); err == nil {
		return &updateResult, nil
	}

	return nil, fmt.Errorf("解析响应数据失败：%s", string(body))
}

// TPLDelRequest - 删除模板请求参数
type TPLDelRequest struct {
	ID int64 `schema:"tpl_id"`
}

// Verify used to check the correctness of the request parameters
func (req *TPLDelRequest) Verify() error {
	if req.ID == 0 {
		return errors.New("Miss param: tpl_id")
	}
	return nil
}

// Del - 删除模板接口
func (tpl *TPL) Del(input *TPLDelRequest) (*TPLResponse, error) {
	if input == nil {
		input = &TPLDelRequest{}
	}

	err := input.Verify()
	if err != nil {
		return nil, err
	}

	r := tpl.c.newRequest("POST", tpl.c.config.tplHost, "/v2/tpl/del.json")
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

	var result TPLResponse
	if err = tpl.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
