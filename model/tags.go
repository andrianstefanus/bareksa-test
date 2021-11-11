package model

import (
	st "bareksa-test/struck"
	"context"
)

type Tag struct {
	ID         string
	Name       string
	NewsID     string
	Status     string
	CreateBy   string
	UpdateBy   string
	CreateDate string
	UpdateDate string
}

type TagListResponse struct {
	Response Response
	Tags     []Tag
}

type TagsUsecase interface {
	Add(ctx context.Context, tag Tag, logs *[]st.Log) (*Response, error)
	Remove(ctx context.Context, tag Tag, logs *[]st.Log) (*Response, error)
	List(ctx context.Context, tag Tag, logs *[]st.Log) (*TagListResponse, error)
}

type TagsRepository interface {
	Add(ctx context.Context, tag Tag, logs *[]st.Log) (*Response, error)
	Remove(ctx context.Context, tag Tag, logs *[]st.Log) (*Response, error)
	List(ctx context.Context, tag Tag, logs *[]st.Log) (*TagListResponse, error)
}
