package model

import (
	st "bareksa-test/struck"
	"context"
)

type News struct {
	ID         string
	Name       string
	TopicID    string
	Status     string
	CreateBy   string
	UpdateBy   string
	CreateDate string
	UpdateDate string
}

type NewsFilter struct {
	Status []string
	Topics []string
}

type ListResponse struct {
	Response Response
	News     []News
}

type NewsUsecase interface {
	Add(ctx context.Context, news News, logs *[]st.Log) (*Response, error)
	Update(ctx context.Context, news News, logs *[]st.Log) (*Response, error)
	Remove(ctx context.Context, news News, logs *[]st.Log) (*Response, error)
	List(ctx context.Context, filter NewsFilter, logs *[]st.Log) (*ListResponse, error)
}

type NewsRepository interface {
	Add(ctx context.Context, news News, logs *[]st.Log) (*Response, error)
	Update(ctx context.Context, news News, logs *[]st.Log) (*Response, error)
	Remove(ctx context.Context, news News, logs *[]st.Log) (*Response, error)
	List(ctx context.Context, filter NewsFilter, logs *[]st.Log) (*ListResponse, error)
}
