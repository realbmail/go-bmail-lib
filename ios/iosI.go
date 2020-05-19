package bmailLib

import (
	"github.com/BASChain/go-bmail-lib/utils"
	"github.com/BASChain/go-bmail-resolver"
)

type UICallBack interface {
	Notification(typ int, msg string)
	Error(typ int, msg string)
}

const (
	BMErrNone = iota
	BMErrClientInvalid
	BMErrWalletInvalid
	BMErrNoSuchBas
	BMErrCryptFailed
	BMErrInvalidJson
	BMErrSendFailed
)

var uiCallback UICallBack
var basResolver resolver.NameResolver

func InitSystem(cb UICallBack, debug bool) {
	uiCallback = cb
	basResolver = resolver.NewEthResolver(debug)
}

func CalculateHash(mailName string) string {
	return resolver.BMailNameHash(mailName)
}

func MailBcaByHash(mailHash string) string {
	bca, _ := basResolver.BMailBCA(mailHash)
	return string(bca)
}

func CName(mailName string) string {
	mailHash := resolver.BMailNameHash(mailName)
	_, cname := basResolver.BMailBCA(mailHash)
	return cname
}

func MailIcon(mailName string) []byte {
	if mailName == "" {
		return nil
	}
	return utils.GenIDIcon(mailName)
}

func MailID() string {
	return utils.UUID()
}
