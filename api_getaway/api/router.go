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
	Impact := pb.NewImpactCalculatorClient(conn)
	Community := pb.NewCommunityServiceClient(conn)
	handler := handler.NewHandler(Habits, Impact, Community)

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
	router.POST("/api/habits", handler.CreateGroup)
	router.GET("/api/habits/:id", handler.GetGroup)
	router.PUT("/api/habits/:id", handler.UpdateGroup)
	router.DELETE("/api/habits/:id", handler.DeleteGroup)
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
	return router
}
