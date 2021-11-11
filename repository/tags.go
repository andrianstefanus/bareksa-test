package repository

import (
	db "bareksa-test/database"
	fx "bareksa-test/function"
	md "bareksa-test/model"
	st "bareksa-test/struck"
	"context"
)

type TagsRepository struct {
	DB db.Database
}

func InitiateTagsRepository(db db.Database) md.TagsRepository {
	return &TagsRepository{
		DB: db,
	}
}

func (r *TagsRepository) Add(ctx context.Context, tag md.Tag, logs *[]st.Log) (response *md.Response, err error) {
	// add news
	sAdd, _, _ := fx.DoRawQuery(
		r.DB.HerokuDB,
		"insert into news_tag (news_id, tag_id, status, create_by) values (@NewsID, @TagID, @Status, @CreateBy) on conflict (news_id, tag_id) do update set status = @Status, update_date = now(), update_by = @UpdateBy;",
		map[string]interface{}{"NewsID": tag.NewsID, "TagID": tag.ID, "Status": "A", "CreateBy": "Admin", "UpdateBy": "Admin"},
		nil,
		logs)

	// err
	if sAdd == 1 {
		return &md.Response{
			Status: "500",
			Msg:    "Terjadi Kesalahan",
		}, nil
	}

	// // duplicate user
	// if eCreate.Status == fx.ConvStrToInt(os.Getenv("ERR_DATA_DUPLICATE_CODE"), 553) {
	// 	return ct.Register(ct.Response(os.Getenv("RESPONSE_ACCOUNT_DUPLICATED_CODE"), os.Getenv("RESPONSE_ACCOUNT_DUPLICATED")), md.User{}), nil
	// }

	return &md.Response{
		Status: "200",
		Msg:    "Success",
	}, nil
}

func (r *TagsRepository) Remove(ctx context.Context, tag md.Tag, logs *[]st.Log) (response *md.Response, err error) {
	// update news
	sUpdate, _, _ := fx.DoRawQuery(
		r.DB.HerokuDB,
		"update news_tag set status = @Status, update_date = now(), update_by = @UpdateBy where news_id = @NewsID and tag_id = @TagID;",
		map[string]interface{}{"NewsID": tag.NewsID, "TagID": tag.ID, "Status": "I", "UpdateBy": "Admin"},
		nil,
		logs)

	// err
	if sUpdate == 1 {
		return &md.Response{
			Status: "500",
			Msg:    "Terjadi Kesalahan",
		}, nil
	}

	return &md.Response{
		Status: "200",
		Msg:    "Success",
	}, nil
}

func (r *TagsRepository) List(ctx context.Context, tag md.Tag, logs *[]st.Log) (response *md.TagListResponse, err error) {
	// var
	var tags []md.Tag = []md.Tag{}

	// get news
	sGet, _, _ := fx.DoRawQuery(
		r.DB.HerokuDB,
		"select b.id, b.name, a.news_id, a.status, a.create_by, a.update_by, a.create_date, a.update_date from news_tag a inner join master_tag b on a.tag_id = b.id where news_id = @NewsID ",
		map[string]interface{}{"NewsID": tag.NewsID},
		&tags,
		logs)

	// err
	if sGet == 1 {
		return &md.TagListResponse{
			Response: md.Response{
				Status: "500",
				Msg:    "Terjadi Kesalahan",
			},
			Tags: tags,
		}, nil
	}

	// empty
	if len(tags) == 0 {
		return &md.TagListResponse{
			Response: md.Response{
				Status: "404",
				Msg:    "Data kosong",
			},
			Tags: tags,
		}, nil
	}

	return &md.TagListResponse{
		Response: md.Response{
			Status: "200",
			Msg:    "Success",
		},
		Tags: tags,
	}, nil
}
