package web

import "net/http"

// RedirectResponse 页面跳转
type RedirectResponse struct {
	response *Response
	request  *Request
	location string
	code     int
}

// NewRedirectResponse 创建RedirectResponse对象
func NewRedirectResponse(response *Response, request *Request, location string, code int) *RedirectResponse {
	return &RedirectResponse{
		response: response,
		request:  request,
		location: location,
		code:     code,
	}
}

// WithCode set response code and return itself
func (resp *RedirectResponse) WithCode(code int) *RedirectResponse {
	resp.code = code
	return resp
}

// CreateResponse 创建响应内容
func (resp *RedirectResponse) CreateResponse() error {
	http.Redirect(resp.response.w, resp.request.r, resp.location, resp.code)
	return nil
}
