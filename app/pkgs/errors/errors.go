package errors

func errorWrapper(defaultMessage string, code int) func() *BusinessError {
	return func() *BusinessError {
		return NewBusinessError(defaultMessage, code)
	}
}

var (
	OK = errorWrapper("Success.", 200)

	//50000 ~ 50100 通用异常码
	InvalidRequest   = errorWrapper("Invalid Request", 50000)
	InvalidParameter = errorWrapper("Invalid Parameter", 50001)
	MissParameter    = errorWrapper("Miss Parameter", 50002)
	JsonFormatFailed = errorWrapper("Json format failed", 50003)

	//50100 ~ 业务异常

)
