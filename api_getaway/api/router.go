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
	router.Group("/habits")
	router.POST("/create", handler.CreateHabit)
	router.GET("/get/:id", handler.GetHabits)
	router.PUT("/update/:id", handler.UpdateHabit)
	router.DELETE("/delete/:id", handler.DeleteHabit)
	router.POST("/log", handler.LogHabit)
	router.GET("/logs/:id", handler.GetHabitLogs)
	router.GET("/suggestions", handler.GetHabitSuggestions)
	
	//Community
	router.Group("/community")
	router.POST("/groups/create", handler.CreateGroup)
	router.GET("/get/:id", handler.GetGroup)
	router.PUT("/update/:id", handler.UpdateGroup)
	router.DELETE("/delete/:id", handler.DeleteGroup)
	router.GET("/getall", handler.ListGroups)
	router.POST("/join", handler.JoinGroup)
	router.POST("/leave", handler.LeaveGroup)
	router.PUT("/member/role", handler.UpdateGroupMemberRole)
	router.POST("/post", handler.CreatePost)
	router.GET("/post/:id", handler.GetPost)
	router.PUT("/post/update:id", handler.UpdatePost)
	router.DELETE("/delete/:id", handler.DeletePost)
	router.GET("/groups/posts/:id", handler.ListGroupPosts)
	router.POST("/comments", handler.CreateComment)
	router.GET("/posts/:id/comments", handler.GetPostComments)
	
	//User
	router.Group("/users")
	router.GET("/:id/habits", handler.GetUserHabits)
	router.GET("/:id", handler.GetUser)
	router.PUT("/update/:id", handler.UpdateUser)
	router.DELETE("/delete/:id", handler.DeleteUser)
	router.GET("/profile/:id", handler.GetUserProfile)
	router.PUT("update/profile/:id", handler.UpdateUserProfile)
																																																																																																					
	//Impact
	router.Group("/impact")
	router.GET("/users/:id", handler.GetUserImpact)
	router.GET("/groups/:id/", handler.GetGroupImpact)
	router.GET("/leaderboard/users", handler.GetLeaderboardUsers)
	router.GET("/leaderboard/groups", handler.GetLeaderboardGroups)
	router.POST("/donations", handler.CreateDonation)
	router.GET("/donations/causes", handler.GetCauses)

	return router
}