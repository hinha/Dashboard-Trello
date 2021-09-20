package pb

import (
	"context"
	"github.com/hinha/PAM-Trello/app/pb/trello"
)

type Services interface {
	Analyze(ctx context.Context, req *trello.PamInput) (*trello.AnalyzeResponse, error)
	GetCard(ctx context.Context, req *trello.Request) (*trello.Response, error)
}
