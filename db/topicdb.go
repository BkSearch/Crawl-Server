package db

import (
	"fmt"

	"github.com/BkSearch/Crawl-Server/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TopicDB struct {
	dbRead  *sqlx.DB
	dbWrite *sqlx.DB
}

func NewTopicDB(
	dbRead, dbWrite *sqlx.DB,
) *TopicDB {
	return &TopicDB{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}

func (topicdb *TopicDB) DB() *sqlx.DB {
	return topicdb.dbWrite
}

func (topicdb *TopicDB) AddNewQuestion(question *common.Question) error {
	_, err := topicdb.dbWrite.Exec(
		`INSERT INTO "Question" (id, content, amount_answer, url, post_at, vote)
    VALUES ($1, $2, $3, $4, $5);`,
		question.ID, question.Content, question.AmountAnswer, question.URL, question.PostAt, question.Vote,
	)

	return err
}

func (topicdb *TopicDB) GetQuestion(id int) (*common.Question, error) {
	var question common.Question
	row := topicdb.dbRead.QueryRowx(`SELECT * FROM "Question" WHERE id = $1 LIMIT 1;`, id)
	err := row.StructScan(&question)
	return &question, err
}

func (topicdb *TopicDB) AddNewAnswer(answer *common.Answer) error {
	_, err := topicdb.dbWrite.Exec(
		`INSERT INTO "Answer" (id, content, url , post_at, vote, question_id)
    VALUES ($1, $2, $3, $4, $5);`,
		answer.ID, answer.Content, answer.URL, answer.PostAt, answer.Vote, answer.Question_ID,
	)

	return err
}

func (topicdb *TopicDB) getAnswer(id int) (*common.Answer, error) {
  var answer common.Answer
  row := topicdb.dbRead.QueryRowx(`SELECT * FROM "Question" WHERE id = $1 LIMIT 1;`, id)
  err := row.StructScan(&answer)
  return &answer, err
}

func (topicdb *TopicDB) getListAnswerOfQuestion(id int) ([]common.Answer, error) {
  // var answer []common.Answer
  rows := topicdb.dbRead.QueryRowx(
    `SELECT * FROM "Answer" WHERE question_id = $1;`, id)
  fmt.Print(rows)
  return nil, nil
}

func (topicdb *TopicDB) addNewTopic(question *common.Question, answer []common.Answer) error {
  return nil
}
