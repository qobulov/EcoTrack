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
	assert.Equal(t, "1", group.Group.Id)
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
	req := &pb.GetGroupRequest{
		Id: "1",
	}
	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_by"}).
		AddRow("1", "Test Group", "Testing group creation", 1)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, created_by FROM groups WHERE id = $1`)).
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

func TestUpdateGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	req := &pb.UpdateGroupRequest{
		Id:          "1",
		Name:        "Updated Test Group",
		Description: "Updated group description",
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE groups SET name = $1, description = $2 WHERE id = $3`)).
		WithArgs(req.Name, req.Description, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := server.UpdateGroup(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.Group.Id)
	assert.Equal(t, req.Name, resp.Group.Name)
	assert.Equal(t, req.Description, resp.Group.Description)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	req := &pb.DeleteGroupRequest{
		Id: "1",
	}

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM groups WHERE id = $1`)).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := server.DeleteGroup(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Group deleted successfully", resp.Message)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestListGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	req := &pb.ListGroupsRequest{}

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_by"}).
		AddRow("1", "Test Group 1", "Description 1", 1).
		AddRow("2", "Test Group 2", "Description 2", 2)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, created_by FROM groups`)).
		WillReturnRows(rows)

	resp, err := server.ListGroups(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Groups, 2)
	assert.Equal(t, "1", resp.Groups[0].Id)
	assert.Equal(t, "Test Group 1", resp.Groups[0].Name)
	assert.Equal(t, "Description 1", resp.Groups[0].Description)
	assert.Equal(t, int32(1), resp.Groups[0].CreatedBy)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestJoinGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	req := &pb.JoinGroupRequest{
		GroupId: "1",
		UserId:  "user1",
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO group_members (group_id, user_id) VALUES ($1, $2)`)).
		WithArgs(req.GroupId, req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := server.JoinGroup(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.GroupId, resp.Member.GroupId)
	assert.Equal(t, req.UserId, resp.Member.UserId)
	assert.Equal(t, "member", resp.Member.Role)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLeaveGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	req := &pb.LeaveGroupRequest{
		GroupId: "1",
		UserId:  "user1",
	}

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM group_members WHERE group_id = $1 AND user_id = $2`)).
		WithArgs(req.GroupId, req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := server.LeaveGroup(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.GroupId, resp.Member.GroupId)
	assert.Equal(t, req.UserId, resp.Member.UserId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateGroupMemberRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	req := &pb.UpdateGroupMemberRoleRequest{
		GroupId: "1",
		UserId:  "user1",
		Role:    "admin",
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE group_members SET role = $1 WHERE group_id = $2 AND user_id = $3`)).
		WithArgs(req.Role, req.GroupId, req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := server.UpdateGroupMemberRole(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.GroupId, resp.Member.GroupId)
	assert.Equal(t, req.UserId, resp.Member.UserId)
	assert.Equal(t, req.Role, resp.Member.Role)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	server := NewCommunityRepo(db)
	req := &pb.CreatePostRequest{
		GroupId: "1",
		UserId:  "user1",
		Content: "This is a test post",
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO posts (group_id, user_id, content) VALUES ($1, $2, $3) RETURNING id`)).
		WithArgs(req.GroupId, req.UserId, req.Content).
		WillReturnResult(sqlmock.NewResult(1, 1))

	post, err := server.CreatePost(req)
	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "1", post.Post.Id)
	assert.Equal(t, req.GroupId, post.Post.GroupId)
	assert.Equal(t, req.UserId, post.Post.UserId)
	assert.Equal(t, req.Content, post.Post.Content)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
