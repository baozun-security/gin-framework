package errors

func errorGenerator(defaultMessage string, code int) func() *BusinessError {
	return func() *BusinessError {
		return NewBusinessError(defaultMessage, code)
	}
}

var (
	OK = NewBusinessError("Success.", 200)

	//50000 ~ 50100 通用异常码
	InvalidRequest   = errorGenerator("Invalid Request", 50000)
	InvalidParameter = errorGenerator("Invalid Parameter", 50001)
	MissParameter    = errorGenerator("Miss Parameter", 50002)
	JsonFormatFailed = errorGenerator("Json format failed", 50003)

	//50100 ~ 业务异常

)
