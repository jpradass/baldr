package log

import "go.uber.org/zap"

func intoZapField(input any) zap.Field {
	switch v := input.(type) {
	case string:
		return zap.String(v, "")
	case int:
		return zap.Int("", v)
	default:
		return zap.String(input.(string), "")
	}
}
