package usecase

import (
	md "bareksa-test/model"
	st "bareksa-test/struck"
	"context"
)

type NewsUsecase struct {
	newsRepository md.NewsRepository
}

func InitiateNewsUsecase(newsRepository md.NewsRepository) md.NewsUsecase {
	return &NewsUsecase{
		newsRepository: newsRepository,
	}
}

func (u *NewsUsecase) Add(ctx context.Context, news md.News, logs *[]st.Log) (response *md.Response, err error) {
	return u.newsRepository.Add(ctx, news, logs)
}

func (u *NewsUsecase) Update(ctx context.Context, news md.News, logs *[]st.Log) (response *md.Response, err error) {
	return u.newsRepository.Update(ctx, news, logs)
}

func (u *NewsUsecase) Remove(ctx context.Context, news md.News, logs *[]st.Log) (response *md.Response, err error) {
	return u.newsRepository.Remove(ctx, news, logs)
}

func (u *NewsUsecase) List(ctx context.Context, filter md.NewsFilter, logs *[]st.Log) (response *md.ListResponse, err error) {
	return u.newsRepository.List(ctx, filter, logs)
}
