package usecase

import (
	md "bareksa-test/model"
	st "bareksa-test/struck"
	"context"
)

type TagsUsecase struct {
	tagsRepository md.TagsUsecase
}

func InitiateTagsUsecase(tagsRepository md.TagsRepository) md.TagsUsecase {
	return &TagsUsecase{
		tagsRepository: tagsRepository,
	}
}

func (u *TagsUsecase) Add(ctx context.Context, tag md.Tag, logs *[]st.Log) (response *md.Response, err error) {
	return u.tagsRepository.Add(ctx, tag, logs)
}

func (u *TagsUsecase) Remove(ctx context.Context, tag md.Tag, logs *[]st.Log) (response *md.Response, err error) {
	return u.tagsRepository.Remove(ctx, tag, logs)
}

func (u *TagsUsecase) List(ctx context.Context, tag md.Tag, logs *[]st.Log) (response *md.TagListResponse, err error) {
	return u.tagsRepository.List(ctx, tag, logs)
}
