package repository

import (
	db "bareksa-test/database"
	fx "bareksa-test/function"
	md "bareksa-test/model"
	st "bareksa-test/struck"
	"context"
)

type NewsRepository struct {
	DB db.Database
}

func InitiateNewsRepository(db db.Database) md.NewsRepository {
	return &NewsRepository{
		DB: db,
	}
}

func (r *NewsRepository) Add(ctx context.Context, news md.News, logs *[]st.Log) (response *md.Response, err error) {
	// add news
	sAdd, _, _ := fx.DoRawQuery(
		r.DB.HerokuDB,
		"insert into news (id, name, topic_id, status, create_by) values (concat('N', nextval('news_seq')), @Name, @TopicID, @Status, @CreateBy);",
		map[string]interface{}{"Name": news.Name, "TopicID": news.TopicID, "Status": "D", "CreateBy": "Admin"},
		nil,
		logs)

	// err
	if sAdd == 1 {
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

func (r *NewsRepository) Update(ctx context.Context, news md.News, logs *[]st.Log) (response *md.Response, err error) {
	// update news
	sUpdate, _, _ := fx.DoRawQuery(
		r.DB.HerokuDB,
		"update news set name = @Name, topic_id = @TopicID, status = @Status, update_date = now(), update_by = @UpdateBy where id = @ID;",
		map[string]interface{}{"ID": news.ID, "Name": news.Name, "TopicID": news.TopicID, "Status": news.Status, "UpdateBy": "Admin"},
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

func (r *NewsRepository) Remove(ctx context.Context, news md.News, logs *[]st.Log) (response *md.Response, err error) {
	// update news
	sUpdate, _, _ := fx.DoRawQuery(
		r.DB.HerokuDB,
		"update news set status = @Status, update_date = now(), update_by = @UpdateBy where id = @ID;",
		map[string]interface{}{"ID": news.ID, "Status": "X", "UpdateBy": "Admin"},
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

func (r *NewsRepository) List(ctx context.Context, filter md.NewsFilter, logs *[]st.Log) (response *md.ListResponse, err error) {
	// var
	var news []md.News = []md.News{}

	// query
	query := "select id, name, topic_id, status, create_by, update_by, create_date, update_date from news where true "

	if len(filter.Status) > 0 {
		query += "and status in (@Status) "
	}

	if len(filter.Topics) > 0 {
		query += "and topic_id in (@Topic) "
	}
	// get news
	sGet, _, _ := fx.DoRawQuery(
		r.DB.HerokuDB,
		query,
		map[string]interface{}{"Status": filter.Status, "Topic": filter.Topics},
		&news,
		logs)

	// err
	if sGet == 1 {
		return &md.ListResponse{
			Response: md.Response{
				Status: "500",
				Msg:    "Terjadi Kesalahan",
			},
			News: news,
		}, nil
	}

	// empty
	if len(news) == 0 {
		return &md.ListResponse{
			Response: md.Response{
				Status: "404",
				Msg:    "Data kosong",
			},
			News: news,
		}, nil
	}

	return &md.ListResponse{
		Response: md.Response{
			Status: "200",
			Msg:    "Success",
		},
		News: news,
	}, nil
}
