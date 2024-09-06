package gemini

import "context"

type IService interface {
	GenContent(ctx context.Context, content string) (res RespContent, err error)
}
