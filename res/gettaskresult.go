package res

type GetTaskResultResponse struct {
	ErrorId          int    `json:"errorId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	Solution         struct {
		Text               string            `json:"text"`
		GRecaptchaResponse string            `json:"gRecaptchaResponse"`
		UserAgent          string            `json:"userAgent"`
		RespKey            string            `json:"respKey"`
		Token              string            `json:"token"`
		Headers            map[string]string `json:"headers"`
		Cookies            map[string]string `json:"cookies"`
		User_Agent         string            `json:"user_agent"`
		Cf                 bool              `json:"cf"`
		RequestHeaders     map[string]string `json:"request_headers"`
		Content            string            `json:"content"`
	} `json:"solution"`
	Status string `json:"status"`
}
