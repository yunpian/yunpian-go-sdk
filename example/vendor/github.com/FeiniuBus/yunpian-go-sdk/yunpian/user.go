package yunpian

// User is used to manipulate the user API
type User struct {
	c *Client
}

// User is used to return a handle to the user APIs
func (c *Client) User() *User {
	return &User{c}
}

// UserResponse - 账户信息响应
type UserResponse struct {
	Nick             string  `json:"nick"`
	Created          string  `json:"gmt_created"`
	Mobile           string  `json:"mobile"`
	Email            string  `json:"email"`
	IPWhitelist      string  `json:"ip_whitelist"`
	APIVersion       string  `json:"api_version"`
	Balance          float64 `json:"balance"`
	AlarmBalance     int64   `json:"alarm_balance"`
	EmergencyContact string  `json:"emergency_contact"`
	EmergencyMobile  string  `json:"emergency_mobile"`
}

// Get - 查询账户信息接口
func (u *User) Get() (*UserResponse, error) {
	r := u.c.newRequest("POST", u.c.config.userHost, "/v2/user/get.json")
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

	var result UserResponse
	if err = u.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UserSetRequest - 修改账号信息请求参数
type UserSetRequest struct {
	EmergencyContact string `schema:"emergency_contact,omitempty"`
	EmergencyMobile  string `schema:"emergency_mobile,omitempty"`
	AlarmBalance     int64  `schema:"alarm_balance,omitempty"`
}

// Set - 修改账号信息接口
func (u *User) Set(input *UserSetRequest) (*UserResponse, error) {
	if input == nil {
		input = &UserSetRequest{}
	}

	r := u.c.newRequest("POST", u.c.config.userHost, "/v2/user/set.json")
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

	var result UserResponse
	if err = u.c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
