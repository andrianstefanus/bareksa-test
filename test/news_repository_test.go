package test

import (
	db "bareksa-test/database"
	md "bareksa-test/model"
	r "bareksa-test/repository"
	st "bareksa-test/struck"
	"context"

	"github.com/stretchr/testify/suite"
)

type NewsRepositorySuite struct {
	suite.Suite
	newsRepository md.NewsRepository
}

func (suite *NewsRepositorySuite) SetupSuite() {
	repository := r.InitiateNewsRepository(db.Databases)
	suite.newsRepository = repository
}

func (suite *NewsRepositorySuite) TestAdd() {
	// var
	var logs []st.Log
	ctx := context.Background()
	input := md.News{
		Name:    "Testing News - Unit Test",
		TopicID: "TO2",
	}

	_, err := suite.newsRepository.Add(ctx, input, &logs)

	suite.NoError(err, "no error")
}
