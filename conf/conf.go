package conf

var JWTSecret = "User-api-secret"

var GoMail = struct {
	Username string
	Password string
	Host     string
	Port     int
}{
	Username: "2214380963@qq.com",
	Password: "uylmxoxamqsddjfc",
	Host:     "smtp.qq.com",
	Port:     587,
}

var HttpCode = struct {
	HttpStatusOK      int32
	HttpForbidden     int32
	HttpBadRequest    int32
	HttpNotFound      int32
	HttpParamsMissing int32
}{
	HttpStatusOK:      200,
	HttpForbidden:     403,
	HttpBadRequest:    400,
	HttpNotFound:      404,
	HttpParamsMissing: 422,
}
var HttpMessage = struct {
	MsgOK     string
	MsgError  string
	MsgFailed string
}{
	MsgOK:     "OK",
	MsgError:  "Error",
	MsgFailed: "Failed",
}
