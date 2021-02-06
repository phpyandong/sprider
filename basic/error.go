package basic

import (
	"errors"
	"reflect"
)

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

//===========关于 判断 error 行为 todo  ================
type Error interface {
	error
	NotFound() bool
}
type notFound interface {
	NotFound() bool
}
type ApError struct {
	// Err is the error that occurred during the operation.
	Err error
}
func (c *ApError) Error() string{
	return "ap:error "
}
func(c *ApError) NotFound() bool{
	t, ok := c.Err.(notFound)

	return ok && t.NotFound()
}
func testAppErr(){
	err := ApError{errors.New("err")}
	err.NotFound()
}
//===========================

func Is(err, target error) bool {
	if target == nil {
		return err == target
	}
	// 通过反射判读 target 是否可以被比较
	isComparable := reflect.TypeOf(target).Comparable()
	for {
		// 循环判断是否相等
		if isComparable && err == target {
			return true
		}
		// 判断是否实现了 is 接口，如果有实现就直接判断
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}

		// 去判断是否实现了 unwrap 的接口，如果实现了就进行 unwrap
		if err = errors.Unwrap(err); err == nil {
			return false
		}
	}
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
