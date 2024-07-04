package postgres

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	pb "EcoTrack/community/genproto/protos"
)

func TestCreateGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	// ctx := context.Background()
	req := &pb.CreateGroupRequest{
		Name:        "Test Group",
		Description: "Testing group creation",
		CreatedBy:   1,
	}

	// Expected query and arguments matcher
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO groups (name, description, created_by) VALUES ($1, $2, $3) RETURNING id`)).
		WithArgs(req.Name, req.Description, req.CreatedBy).
		WillReturnResult(sqlmock.NewResult(1, 1))

	group, err := server.CreateGroup(req)
	assert.NoError(t, err)
	assert.NotNil(t, group)
	assert.Equal(t, int32(1), group.Group.Id)
	assert.Equal(t, req.Name, group.Group.Name)
	assert.Equal(t, req.Description, group.Group.Description)
	assert.Equal(t, req.CreatedBy, group.Group.CreatedBy)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	// ctx := context.Background()
	req := &pb.GetGroupRequest{
		Id: 1,
	}
	rows := sqlmock.NewRows([]string{"name", "description", "created_by"}).
		AddRow("Test Group", "Testing group creation", 1)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT name, description, created_by FROM groups WHERE id = $1`)).
		WithArgs(req.Id).
		WillReturnRows(rows)

	resp, err := server.GetGroup(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.Group.Id)
	assert.Equal(t, "Test Group", resp.Group.Name)
	assert.Equal(t, "Testing group creation", resp.Group.Description)
	assert.Equal(t, int32(1), resp.Group.CreatedBy)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
