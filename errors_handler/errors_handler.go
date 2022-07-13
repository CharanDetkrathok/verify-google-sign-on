package errors_handler

import "net/http"

type ErrorHandler struct {
	Code       int    `json:"code"`
	MessageTh  string `json:"message_th"`
	MessageEng string `json:"message_eng"`
}

func (e ErrorHandler) Error() string {
	return e.MessageTh
}

// oauth2 verify id token google
func NewVerifyGoogleUnauthorizedError() error {
	return ErrorHandler{
		Code:       http.StatusUnauthorized,
		MessageTh:  "ไม่ได้รับอนุญาตจาก google",
		MessageEng: "Not authorized by Google",
	}
}

// STATUS 401 Unauthorized — client ยัง ไม่ได้ระบุตัวตน หรือไม่มี header (สำหรับยังไม่ได้ Login)
// oauth2 verify id token google
func NewNotRumailVerifyGoogleError() error {
	return ErrorHandler{
		Code:       http.StatusUnauthorized,
		MessageTh:  "The email must be @rumail.ru.ac.th only!",
		MessageEng: "email ต้องเป็น @rumail.ru.ac.th เท่านั้น!",
	}
}

// STATUS 401 Unauthorized — client ยัง ไม่ได้ระบุตัวตน หรือไม่มี header (สำหรับยังไม่ได้ Login)
// oauth2 verify id token google
func NewNotStdCodeVerifyGoogleError() error {
	return ErrorHandler{
		Code:       http.StatusUnauthorized,
		MessageTh:  "email ต้องเริ่มต้นด้วย รหัสนักศึกษา เท่านั้น! เช่น 6555555555@rumail.ru.ac.th",
		MessageEng: "Email must start with student ID only! For example 65555@rumail.ru.ac.th",
	}
}