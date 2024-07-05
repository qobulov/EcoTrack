package service

import (
	storage "EcoTrack/community/storage/postgres"
	pb "EcoTrack/community/genproto/protos"
	"context"
	"fmt"
	"log"
)

type CommunityService struct {
	pb.UnimplementedCommunityServiceServer
	Storage storage.Server
}

func (s *CommunityService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.GroupResponse, error) {
	res, err := s.Storage.CreateGroup(req)
	if err != nil {
		log.Println("error while creating group: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GroupResponse, error) {
	res, err := s.Storage.GetGroup(req)
	if err != nil {
		log.Println("error while getting group: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.GroupResponse, error) {
	res, err := s.Storage.UpdateGroup(req)
	if err != nil {
		log.Println("error while updating group: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	res, err := s.Storage.DeleteGroup(req)
	if err != nil {
		log.Println("error while deleting group: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	res, err := s.Storage.ListGroups(req)
	if err != nil {
		log.Println("error while listing groups: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) JoinGroup(ctx context.Context, req *pb.JoinGroupRequest) (*pb.GroupMemberResponse, error) {
	res, err := s.Storage.JoinGroup(req)
	if err != nil {
		log.Println("error while joining group: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) LeaveGroup(ctx context.Context, req *pb.LeaveGroupRequest) (*pb.GroupMemberResponse, error) {
	res, err := s.Storage.LeaveGroup(req)
	if err != nil {
		log.Println("error while leaving group: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) UpdateGroupMemberRole(ctx context.Context, req *pb.UpdateGroupMemberRoleRequest) (*pb.GroupMemberResponse, error) {
	res, err := s.Storage.UpdateGroupMemberRole(req)
	if err != nil {
		log.Println("error while updating group member role: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	res, err := s.Storage.CreatePost(req)
	if err != nil {
		log.Println("error while creating post: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	res, err := s.Storage.GetPost(req)
	if err != nil {
		log.Println("error while getting post: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	res, err := s.Storage.UpdatePost(req)
	if err != nil {
		log.Println("error while updating post: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	res, err := s.Storage.DeletePost(req)
	if err != nil {
		log.Println("error while deleting post: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) ListGroupPosts(ctx context.Context, req *pb.ListGroupPostsRequest) (*pb.ListPostsResponse, error) {
	res, err := s.Storage.ListGroupPosts(req)
	if err != nil {
		log.Println("error while listing group posts: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	res, err := s.Storage.CreateComment(req)
	if err != nil {
		log.Println("error while creating comment: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

func (s *CommunityService) GetPostComments(ctx context.Context, req *pb.GetPostCommentsRequest) (*pb.ListCommentsResponse, error) {
	res, err := s.Storage.GetPostComments(req)
	if err != nil {
		log.Println("error while getting post comments: ", err)
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}
