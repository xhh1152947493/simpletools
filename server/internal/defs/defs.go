package defs

type ExitMode int32

const (
	ExitModeSaveAndExit ExitMode = 1
	ExitModeKillNoWait  ExitMode = 2
	ExitModePanicNoWait ExitMode = 3
)

type EEventSinkSource = int32

const (
	EssRedis    EEventSinkSource = 1
	EssMysql    EEventSinkSource = 2
	EssGSignalQ EEventSinkSource = 3
)

type LanguageType int

const (
	LanguageTypeChinese LanguageType = 1
	LanguageTypeEnglish LanguageType = 2
)
