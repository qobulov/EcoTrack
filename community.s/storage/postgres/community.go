package postgres

import (
	"database/sql"

	pb "EcoTrack/community/genproto/protos"

	_ "github.com/lib/pq"
)

type Server struct {
	pb.UnimplementedCommunityServiceServer
	db *sql.DB
}

func NewCommunityRepo(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) CreateGroup(req *pb.CreateGroupRequest) (*pb.GroupResponse, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO groups (name, description, created_by) VALUES ($1, $2, $3) RETURNING id", req.Name, req.Description, req.CreatedBy).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &pb.GroupResponse{Group: &pb.Group{Id: int32(id), Name: req.Name, Description: req.Description, CreatedBy: req.CreatedBy}}, nil
}

func (s *Server) GetGroup(req *pb.GetGroupRequest) (*pb.GroupResponse, error) {
	var name, description string
	var createdBy int32
	err := s.db.QueryRow("SELECT name, description, created_by FROM groups WHERE id = $1", req.Id).Scan(&name, &description, &createdBy)
	if err != nil {
		return nil, err
	}
	return &pb.GroupResponse{Group: &pb.Group{Id: req.Id, Name: name, Description: description, CreatedBy: createdBy}}, nil
}

func (s *Server) UpdateGroup(req *pb.UpdateGroupRequest) (*pb.GroupResponse, error) {
	_, err := s.db.Exec("UPDATE groups SET name = $1, description = $2 WHERE id = $3", req.Name, req.Description, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GroupResponse{Group: &pb.Group{Id: req.Id, Name: req.Name, Description: req.Description}}, nil
}

func (s *Server) DeleteGroup(req *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	_, err := s.db.Exec("DELETE FROM groups WHERE id = $1", req.Id)
	if err != nil {
		return &pb.DeleteGroupResponse{Message: err.Error()}, nil
	}
	return &pb.DeleteGroupResponse{Message: "Group deleted successfully"}, nil
}

func (s *Server) ListGroups(req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	rows, err := s.db.Query("SELECT id, name, description, created_by FROM groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*pb.Group
	for rows.Next() {
		var id int32
		var name, description string
		var createdBy int32
		if err := rows.Scan(&id, &name, &description, &createdBy); err != nil {
			return nil, err
		}
		groups = append(groups, &pb.Group{Id: id, Name: name, Description: description, CreatedBy: createdBy})
	}

	return &pb.ListGroupsResponse{Groups: groups}, nil
}

func (s *Server) JoinGroup(req *pb.JoinGroupRequest) (*pb.GroupMemberResponse, error) {
	_, err := s.db.Exec("INSERT INTO group_members (group_id, user_id) VALUES ($1, $2)", req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GroupMemberResponse{Member: &pb.GroupMember{GroupId: req.GroupId, UserId: req.UserId, Role: "member"}}, nil
}

func (s *Server) LeaveGroup(req *pb.LeaveGroupRequest) (*pb.GroupMemberResponse, error) {
	_, err := s.db.Exec("DELETE FROM group_members WHERE group_id = $1 AND user_id = $2", req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GroupMemberResponse{Member: &pb.GroupMember{GroupId: req.GroupId, UserId: req.UserId}}, nil
}

func (s *Server) UpdateGroupMemberRole(req *pb.UpdateGroupMemberRoleRequest) (*pb.GroupMemberResponse, error) {
	_, err := s.db.Exec("UPDATE group_members SET role = $1 WHERE group_id = $2 AND user_id = $3", req.Role, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GroupMemberResponse{Member: &pb.GroupMember{GroupId: req.GroupId, UserId: req.UserId, Role: req.Role}}, nil
}

func (s *Server) CreatePost(req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO posts (group_id, user_id, content) VALUES ($1, $2, $3) RETURNING id", req.GroupId, req.UserId, req.Content).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{Post: &pb.Post{Id: int32(id), GroupId: req.GroupId, UserId: req.UserId, Content: req.Content}}, nil
}

func (s *Server) GetPost(req *pb.GetPostRequest) (*pb.PostResponse, error) {
	var groupId, userId int32
	var content string
	err := s.db.QueryRow("SELECT group_id, user_id, content FROM posts WHERE id = $1", req.Id).Scan(&groupId, &userId, &content)
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{Post: &pb.Post{Id: req.Id, GroupId: groupId, UserId: userId, Content: content}}, nil
}

func (s *Server) UpdatePost(req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	_, err := s.db.Exec("UPDATE posts SET content = $1 WHERE id = $2", req.Content, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{Post: &pb.Post{Id: req.Id, Content: req.Content}}, nil
}

func (s *Server) DeletePost(req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	_, err := s.db.Exec("DELETE FROM posts WHERE id = $1", req.Id)
	if err != nil {
		return &pb.DeletePostResponse{Message: err.Error()}, nil
	}
	return &pb.DeletePostResponse{Message: "Post deleted successfully"}, nil
}

func (s *Server) ListGroupPosts(req *pb.ListGroupPostsRequest) (*pb.ListPostsResponse, error) {
	rows, err := s.db.Query("SELECT id, group_id, user_id, content, created_at FROM posts WHERE group_id = $1 ORDER BY created_at DESC LIMIT 20", req.GroupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*pb.Post
	for rows.Next() {
		var id, groupId, userId int32
		var content, createdAt string
		if err := rows.Scan(&id, &groupId, &userId, &content, &createdAt); err != nil {
			return nil, err
		}
		posts = append(posts, &pb.Post{Id: id, GroupId: groupId, UserId: userId, Content: content, CreatedAt: createdAt})
	}

	return &pb.ListPostsResponse{Posts: posts}, nil
}

func (s *Server) CreateComment(req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO comments (post_id, user_id, content) VALUES ($1, $2, $3) RETURNING id", req.PostId, req.UserId, req.Content).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &pb.CommentResponse{Comment: &pb.Comment{Id: int32(id), PostId: req.PostId, UserId: req.UserId, Content: req.Content}}, nil
}

func (s *Server) GetPostComments(req *pb.GetPostCommentsRequest) (*pb.ListCommentsResponse, error) {
	rows, err := s.db.Query("SELECT id, post_id, user_id, content, created_at FROM comments WHERE post_id = $1 ORDER BY created_at DESC", req.PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*pb.Comment
	for rows.Next() {
		var id, postId, userId int32
		var content, createdAt string
		if err := rows.Scan(&id, &postId, &userId, &content, &createdAt); err != nil {
			return nil, err
		}
		comments = append(comments, &pb.Comment{Id: id, PostId: postId, UserId: userId, Content: content, CreatedAt: createdAt})
	}

	return &pb.ListCommentsResponse{Comments: comments}, nil
}
