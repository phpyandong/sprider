package basic


/* ================================================================================
 * 错误模块
 * ================================================================================ */
const (
	ArgumentErrorCode         int32 = 422
	BadrequestErrCode         int32 = 400
	UserPermissionErrorCode   int32 = 401
	NotFoundErrorCode         int32 = 204
	BalanceNotEnoughCode      int32 = 423
	OverLimitCode             int32 = 424
	EndofmonthCantCashoutCode int32 = 425
	DirtyVoiceCode            int32 = 10000
	VoiceTransformCode        int32 = 10001
)

var (
	ArgumentError              *CustomError
	UserPermissionError        *CustomError
	BadrequestErr              *CustomError
	NotFoundError              *CustomError
	BalanceNotEnough           *CustomError
	OverLimit                  *CustomError
	EndofmonthCantCashoutError *CustomError
	DirtyVoiceErr              *CustomError
	VoiceTransformErr          *CustomError
)

var (
	currentLanguageCode string
	errs                map[string]map[int32]string
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 模块初始化
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func init() {
	currentLanguageCode = "zh-cn"
	errs = make(map[string]map[int32]string, 0)

	initMessage()
	initError()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化消息字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func initMessage() {
	chineseLanguageCode := "zh-cn"
	errs[chineseLanguageCode] = make(map[int32]string, 0)

	errs[chineseLanguageCode][ArgumentErrorCode] = "请求参数错误"
	errs[chineseLanguageCode][BadrequestErrCode] = "错误请求"
	errs[chineseLanguageCode][UserPermissionErrorCode] = "权限不允许"
	errs[chineseLanguageCode][NotFoundErrorCode] = "未找到"
	errs[chineseLanguageCode][BalanceNotEnoughCode] = "余额不足"
	errs[chineseLanguageCode][OverLimitCode] = "超限"
	errs[chineseLanguageCode][EndofmonthCantCashoutCode] = "月末不允许提现"
	errs[chineseLanguageCode][DirtyVoiceCode] = "包含敏感词"
	errs[chineseLanguageCode][VoiceTransformCode] = "未识别到语音"

}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化Error
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func initError() {
	ArgumentError = customError(ArgumentErrorCode)
	UserPermissionError = customError(UserPermissionErrorCode)
	BadrequestErr = customError(BadrequestErrCode)
	NotFoundError = customError(NotFoundErrorCode)
	BalanceNotEnough = customError(BalanceNotEnoughCode)
	OverLimit = customError(OverLimitCode)
	EndofmonthCantCashoutError = customError(EndofmonthCantCashoutCode)
	DirtyVoiceErr = customError(DirtyVoiceCode)
	VoiceTransformErr = customError(VoiceTransformCode)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 返回指定错误对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func customError(code int32) *CustomError {
	msg := errs[currentLanguageCode][code]
	return &CustomError{
		Code: code,
		Msg:  msg,
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 返回自定义错误对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewCustomError(code int32, msg string) *CustomError {
	return &CustomError{
		Code: code,
		Msg:  msg,
	}
}
