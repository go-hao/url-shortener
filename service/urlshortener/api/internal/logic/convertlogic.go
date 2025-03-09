package logic

import (
	"context"
	"net/url"
	"path"

	"github.com/go-hao/url-shortener/pkg/connect"
	"github.com/go-hao/url-shortener/pkg/md5"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/svc"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/types"
	"github.com/go-hao/url-shortener/service/urlshortener/model"
	"github.com/go-hao/url-shortener/service/urlshortener/modelc"

	"github.com/go-hao/zero/xerrors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertReq) (resp *types.ConvertResp, err error) {
	if ok := connect.Ping(req.LongUrl); !ok {
		return nil, xerrors.New(1001, "invalid url")
	}
	longUrlMd5 := md5.Sum(req.LongUrl)

	_, err = l.svcCtx.UrlMapModel.FindOneByLongUrlMd5(l.ctx, longUrlMd5)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("UrlMapModel.FindOneByLongUrlMd5: %s", err.Error())
		return nil, xerrors.New(1002, "internal error")
	}
	if err == nil {
		return nil, xerrors.New(1003, "already exists")
	}

	goodUrl, err := url.Parse(req.LongUrl)
	if err != nil {
		return nil, xerrors.New(1001, "invalid url")
	}

	baseUrl := path.Base(goodUrl.Path)
	_, err = l.svcCtx.UrlMapModel.FindOneByShortUrl(l.ctx, baseUrl)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("UrlMapModel.FindOneByShortUrl: %s", err.Error())
		return nil, xerrors.New(1002, "internal error")
	}
	if err == nil {
		return nil, xerrors.New(1003, "already exists")
	}

	var shortUrl string
	for {
		id, err := l.svcCtx.SequenceModel.Next(l.ctx)
		if err != nil {
			l.Logger.Errorf("SequenceModel.Next: %s", err.Error())
			return nil, xerrors.New(1002, "internal error")
		}

		shortUrl = l.svcCtx.Base62.Encode(id)

		if _, ok := l.svcCtx.ShortUrlBlackList[shortUrl]; !ok {
			break
		}
	}

	_, err = l.svcCtx.UrlMapModel.Insert(l.ctx, &modelc.UrlMap{
		LongUrl:    req.LongUrl,
		LongUrlMd5: longUrlMd5,
		ShortUrl:   shortUrl,
	})
	if err != nil {
		l.Logger.Errorf("UrlMapModel.Insert: %s", err.Error())
		return nil, xerrors.New(1002, "internal error")
	}

	resp = &types.ConvertResp{
		ShortUrl: l.svcCtx.Config.ShortUrlDomain + "/" + shortUrl,
	}

	return resp, nil
}
