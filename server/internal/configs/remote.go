package configs

import (
	"simpletools/internal/defs"
)

type CommonResp struct {
	Code defs.ErrCode `json:"code"`
	Data string       `json:"data"`
	Sign string       `json:"sign"`
}

type MysqlCommonRespEncrypt struct {
	Count int64  // executed sql count
	Ret   string // 自己每个接口单独处理
	OK    bool   // 验证接口，默认返回false
	Err   string // 如果有错误信息，就返回错误信息
}

type MysqlCommonResp struct {
	Encrypt string `json:"encrypt"` // 密文，解密后得到MysqlCommonRespEncrypt类型的json字符串
}
