package handlers

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	ctx "simpletools/internal/api/context"
	"simpletools/internal/data"
	"simpletools/internal/defs"
	"simpletools/internal/utils"
)

const (
	deepSeekApiUrl    = "https://api.deepseek.com/chat/completions"
	deepSeekApiKey    = "sk-a773e5fe5edc4616ba955267b5286625"
	deepSeekChatModel = "deepseek-chat"
	aiTranslatePrompt = `
你是一个翻译专家。用户可以向助手发送需要翻译的内容，在用户输入的文本中可能包含多种语言、网络用语、缩写或混合表达登，助手会回答相应的翻译结果，你可以调整语气和风格，并考虑到某些词语的文化内涵和地区差异。同时作为翻译家，需将原文翻译成具有信达雅标准的译文。"信" 即忠实于原文的内容与意图；"达" 意味着译文应通顺易懂，表达清晰；"雅" 则追求译文的文化审美和语言的优美。目标是创作出既忠于原作精神，又符合目标语言文化和读者审美的翻译。本次需要翻译的目标语言为：%s。
`
	languageCN = "中文"
	languageEN = "英文"

	aiNamedPrompt = `
你是一个经验丰富的编程专家。用户会向你发送一段函数或者变量的描述，你需要根据这段描述准确的给出5个命名结果，5个结果以：1.xxx 2.xxx 3.xxx 4.xxx 5.xxx 这样的形式返回，只需要返回这个格式的内容，不要回复任何多余的内容。本次需要命名的风格为：%s。
`
	VarStylePascalCase = "大驼峰"
	VarStyleCamelCase  = "小驼峰"
	VarStyleSnakeCase  = "下划线"
	VarStyleUpperCase  = "全大写下划线常量"
)

type deepSeekApiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type deepSeekApiReq struct {
	Model    string               `json:"model"`
	Messages []deepSeekApiMessage `json:"messages"`
	Stream   bool                 `json:"stream"`
}

type deepSeekApiPromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type deepSeekApiUsage struct {
	PromptTokens         int                            `json:"prompt_tokens"`
	CompletionTokens     int                            `json:"completion_tokens"`
	TotalTokens          int                            `json:"total_tokens"`
	PromptTokensDetails  deepSeekApiPromptTokensDetails `json:"prompt_tokens_details"`
	PromptCacheHitTokens int                            `json:"prompt_cache_hit_tokens"`
}

type deepSeekApiLogprobs struct {
	TokenLogprobs      []float64 `json:"token_logprobs"`
	TokenTextOffset    []int     `json:"token_text_offset"`
	SequenceLogprobs   []float64 `json:"sequence_logprobs"`
	CompletionLogprobs []float64 `json:"completion_logprobs"`
}

type deepSeekApiChoice struct {
	Index        int                 `json:"index"`
	Message      deepSeekApiMessage  `json:"message"`
	FinishReason string              `json:"finish_reason"`
	Logprobs     deepSeekApiLogprobs `json:"logprobs"`
}

type deepSeekApiResp struct {
	ID                string              `json:"id"`
	Object            string              `json:"object"`
	Created           int64               `json:"created"`
	Model             string              `json:"model"`
	Usage             deepSeekApiUsage    `json:"usage"`
	SystemFingerprint string              `json:"system_fingerprint"`
	Choices           []deepSeekApiChoice `json:"choices"`
}

func SendContentToDeepSeek(systemContent, userContent string) (string, error) {
	req := deepSeekApiReq{
		Model:    deepSeekChatModel,
		Messages: []deepSeekApiMessage{{Role: "system", Content: systemContent}, {Role: "user", Content: userContent}},
		Stream:   false,
	}
	bs, err := jsoniter.Marshal(req)
	if err != nil {
		return "", err
	}
	resp, err := utils.HttpPostWithHeader(deepSeekApiUrl, bs, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + deepSeekApiKey,
	})
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func GetResultFromDeepSeekResp(resp string) (string, error) {
	rlt := &deepSeekApiResp{}
	if err := jsoniter.UnmarshalFromString(resp, &rlt); err != nil {
		return "", err
	}
	result := rlt.Choices[0].Message.Content
	return result, nil
}

func GetTranslatePrompt(language int64) string {
	switch defs.LanguageType(language) {
	case defs.LanguageTypeChinese:
		return fmt.Sprintf(aiTranslatePrompt, languageCN)
	case defs.LanguageTypeEnglish:
		return fmt.Sprintf(aiTranslatePrompt, languageEN)
	}
	return fmt.Sprintf(aiTranslatePrompt, languageCN)
}

func OnAITranslateHandler(ctx *ctx.CustomContext) *defs.CustomError {
	content := ctx.GetString("content")
	language := ctx.GetInt64("language")

	resp, err := SendContentToDeepSeek(GetTranslatePrompt(language), content)
	if err != nil {
		return defs.NewCustomError(defs.ErrCodeSystemError, err)
	}

	result, err := GetResultFromDeepSeekResp(resp)
	if err != nil {
		return defs.NewCustomError(defs.ErrCodeSystemError, err)
	}

	ctx.AnswerOK(result)
	data.Log().Info().Str("content", content).Str("result", result).Msg("OnAITranslateHandler success")
	return nil
}

func GetNamedPrompt(style int64) string {
	switch defs.VariableType(style) {
	case defs.VariableTypePascalCase:
		return fmt.Sprintf(aiNamedPrompt, VarStylePascalCase)
	case defs.VariableTypeCamelCase:
		return fmt.Sprintf(aiNamedPrompt, VarStyleCamelCase)
	case defs.VariableTypeSnakeCase:
		return fmt.Sprintf(aiNamedPrompt, VarStyleSnakeCase)
	case defs.VariableTypeUpperCase:
		return fmt.Sprintf(aiNamedPrompt, VarStyleUpperCase)
	}
	return fmt.Sprintf(aiNamedPrompt, VarStylePascalCase)
}

func OnAINamedHandler(ctx *ctx.CustomContext) *defs.CustomError {
	content := ctx.GetString("content")
	style := ctx.GetInt64("style")

	resp, err := SendContentToDeepSeek(GetNamedPrompt(style), content)
	if err != nil {
		return defs.NewCustomError(defs.ErrCodeSystemError, err)
	}

	result, err := GetResultFromDeepSeekResp(resp)
	if err != nil {
		return defs.NewCustomError(defs.ErrCodeSystemError, err)
	}

	ctx.AnswerOK(result)
	data.Log().Info().Str("content", content).Str("result", result).Msg("OnAINamedHandler success")
	return nil
}
