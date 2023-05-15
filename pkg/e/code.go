package e

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidToken
	CodeNeedLogin
	CodeInvalidID
	CodeVoteTimeExpire
	CodeVoteRepeat
	CodeLoginRepeat
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "成功",
	CodeInvalidParam:    "请求参数有误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidToken:    "无效的token",
	CodeNeedLogin:       "需要登录",
	CodeInvalidID:       "无效的id",
	CodeVoteTimeExpire:  "投票时间已过",
	CodeVoteRepeat:      "重复投票",
	CodeLoginRepeat:     "用户已在别处登录",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
