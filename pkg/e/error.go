package e

import "errors"

var (
	ErrorUserExist       = errors.New(CodeUserExist.Msg())
	ErrorUserNotExist    = errors.New(CodeUserNotExist.Msg())
	ErrorInvalidPassword = errors.New(CodeInvalidPassword.Msg())
	ErrorInvalidToken    = errors.New(CodeInvalidToken.Msg())
	ErrorNeedLogin       = errors.New(CodeNeedLogin.Msg())
	ErrorInvalidID       = errors.New(CodeInvalidID.Msg())
	ErrorVoteTimeExpire  = errors.New(CodeVoteTimeExpire.Msg())
	ErrorVoteRepeat      = errors.New(CodeVoteRepeat.Msg())
)
