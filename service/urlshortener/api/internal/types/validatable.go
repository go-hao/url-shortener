package types

import "github.com/go-hao/zero/xvalidator"

func (r *ConvertReq) GetErrors() xvalidator.Errors {
	return xvalidator.Errors{
		"LongUrl.required": "long_url cannot be empty",
	}
}

func (r *ShowReq) GetErrors() xvalidator.Errors {
	return xvalidator.Errors{
		"ShortUrl.required": "short_url cannot be empty",
	}
}
