package logic

import (
	"context"

	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/svc"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/types"
	"github.com/go-hao/url-shortener/service/urlshortener/model"

	"github.com/go-hao/zero/xerrors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowReq) (resp *types.ShowResp, err error) {
	u, err := l.svcCtx.UrlMapModel.FindOneByShortUrl(l.ctx, req.ShortUrl)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("UrlMapModel.FindOneByShortUrl: %s", err.Error())
		return nil, xerrors.New(1002, "internal error")
	}

	if err == model.ErrNotFound {
		return nil, xerrors.New(1001, "invalid url")
	}

	resp = &types.ShowResp{
		LongUrl: u.LongUrl,
	}
	return resp, nil
}
