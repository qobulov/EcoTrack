package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "api-getaway/genproto/protos"

	"api-getaway/api/handler"
)

func CreateRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	Habits := pb.NewHabitTrackerServiceClient(conn)
	Impact := pb.NewImpactCalculatorServiceClient(conn)
	Community := pb.NewCommunityServiceClient(conn)
	User := pb.NewUserServiceClient(conn)
	handler := handler.NewHandler(Habits, Impact, Community, User)

	//Habits
	router.POST("/api/habits", handler.CreateHabit)
	router.GET("/api/habits/:id", handler.GetHabits)
	router.PUT("/api/habits/:id", handler.UpdateHabit)
	router.DELETE("/api/habits/:id", handler.DeleteHabit)
	router.GET("/api/users/:id/habits", handler.GetUserHabits)
	router.POST("/api/habits/log", handler.LogHabit)
	router.GET("/api/habits/:id/logs", handler.GetHabitLogs)
	router.GET("/api/habits/suggestions", handler.GetHabitSuggestions)

	//Community
	router.POST("/api/groups", handler.CreateGroup)
	router.GET("/api/groups/:id", handler.GetGroup)
	router.PUT("/api/groups/:id", handler.UpdateGroup)
	router.DELETE("/api/groups/:id", handler.DeleteGroup)
	router.GET("/api/groups", handler.ListGroups)
	router.POST("/api/groups/join", handler.JoinGroup)
	router.POST("/api/groups/leave", handler.LeaveGroup)
	router.PUT("/api/groups/member/role", handler.UpdateGroupMemberRole)
	router.POST("/api/posts", handler.CreatePost)
	router.GET("/api/posts/:id", handler.GetPost)
	router.PUT("/api/posts/:id", handler.UpdatePost)
	router.DELETE("/api/posts/:id", handler.DeletePost)
	router.GET("/api/groups/:group_id/posts", handler.ListGroupPosts)
	router.POST("/api/comments", handler.CreateComment)
	router.GET("/api/posts/:post_id/comments", handler.GetPostComments)

	//User
	router.GET("/api/user/:id", handler.GetUser)
	router.PUT("/api/user/:id", handler.UpdateUser)
	router.DELETE("/api/user/:id", handler.DeleteUser)
	router.GET("/api/user/:id/profile", handler.GetUserProfile)
	router.PUT("/api/user/:id/profile", handler.UpdateUserProfile)

	//Impact
	router.GET("/api/users/:id/impact", handler.GetUserImpact)
	router.GET("/api/groups/:id/impact", handler.GetGroupImpact)
	router.GET("/api/leaderboard/users", handler.GetLeaderboardUsers)
	router.GET("/api/leaderboard/groups", handler.GetLeaderboardGroups)
	router.POST("/api/donations", handler.CreateDonation)
	router.GET("/api/donations/causes", handler.GetCauses)

	return router
}