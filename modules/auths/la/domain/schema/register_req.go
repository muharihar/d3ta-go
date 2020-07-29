package schema

// RegisterRequest type
type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	NickName  string `json:"nickName"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captchaID"`
}
