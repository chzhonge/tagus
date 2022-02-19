package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"regexp"
	mockDB "tagus/db"
	"tagus/model"
	"testing"
)

type Suite struct {
	suite.Suite
	userRepo *UserRepository
	mock     sqlmock.Sqlmock
}

func (s *Suite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gdb, _ := gorm.Open(mockDB.MockDialector{Dialector: sqlite.Dialector{Conn: db}})
	s.userRepo = &UserRepository{Repository{DB: gdb}}
	s.mock = mock
}

func (s *Suite) TestFindUser() {
	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT `id` FROM `users` WHERE `users`.`username` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err := s.userRepo.Find([]string{"id"}, model.User{UserName: "hello"})
	assert.Equal(s.T(), err, gorm.ErrRecordNotFound)
}

func (s *Suite) TestFindUserWithoutColumn() {
	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`username` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err := s.userRepo.Find([]string{}, model.User{UserName: "hello"})
	assert.Equal(s.T(), err, gorm.ErrRecordNotFound)
}

func (s *Suite) TestCreate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `users` (`username`,`password`,`displayname`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs("he", "pa", "dis", sqlmock.AnyArg(), sqlmock.AnyArg(), nil).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	u, err := s.userRepo.Create("he", "pa", "dis")
	assert.Equal(s.T(), err, nil)
	assert.Equal(s.T(), u.UserName, "he")
}

func TestUser(t *testing.T) {
	suite.Run(t, new(Suite))
}
