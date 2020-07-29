package template

import "fmt"

// UserRegisterActivationTemplate template
func UserRegisterActivationTemplate(email, user string, activationURL string) string {
	body := fmt.Sprintf(`
Dear %s,

Please click on the url bellow to complete the verification process for account %s:

%s

If you didn't attempt to verify your email address with our service, delete this email.

Cheers,

Customer Service
	`, user, email, activationURL)

	return body
}

// UserRegisterActivatedTemplate template
func UserRegisterActivatedTemplate(email, user string) string {
	body := fmt.Sprintf(`
Dear %s,

Conglatulation! your account has been activated.

If you didn't attempt to verify your email address with our service, delete this email.

Cheers,

Customer Service
	`, user)

	return body
}
