package conf

var JWTSecret = "user-api-secret"

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
