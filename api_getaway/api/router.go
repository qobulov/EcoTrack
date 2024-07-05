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
	AuthService := pb.NewAuthServiceClient(conn)

	handler := handler.NewHandler(Habits, Impact, Community, User,AuthService)

	// Habits
	habitsGroup := router.Group("/habits")
	{
		habitsGroup.POST("/create", handler.CreateHabit)
		habitsGroup.GET("/user/:id", handler.GetUserHabits)
		habitsGroup.GET("/get/:id", handler.GetHabits)
		habitsGroup.PUT("/update/:id", handler.UpdateHabit)
		habitsGroup.DELETE("/delete/:id", handler.DeleteHabit)
		habitsGroup.POST("/log", handler.LogHabit)
		habitsGroup.GET("/logs/:id", handler.GetHabitLogs)
		habitsGroup.GET("/suggestions", handler.GetHabitSuggestions)
	}

	// Community
	communityGroup := router.Group("/community")
	{
		communityGroup.POST("/groups/create", handler.CreateGroup)
		communityGroup.GET("/get/:id", handler.GetGroup)
		communityGroup.PUT("/update/:id", handler.UpdateGroup)
		communityGroup.DELETE("/delete/:id/group", handler.DeleteGroup)
		communityGroup.GET("/getall", handler.ListGroups)
		communityGroup.POST("/join", handler.JoinGroup)
		communityGroup.POST("/leave", handler.LeaveGroup)
		communityGroup.PUT("/member/role", handler.UpdateGroupMemberRole)
		communityGroup.POST("/post", handler.CreatePost)
		communityGroup.GET("/post/:id", handler.GetPost)
		communityGroup.PUT("/post/update/:id", handler.UpdatePost)
		communityGroup.DELETE("/delete/:id/post", handler.DeletePost)
		communityGroup.GET("/groups/posts/:id", handler.ListGroupPosts)
		communityGroup.POST("/comments", handler.CreateComment)
		communityGroup.GET("/posts/:id/comments", handler.GetPostComments)
	}

	// User
	userGroup := router.Group("/users")
	{
		userGroup.POST("/login", handler.CreateUser)
		userGroup.GET("/:id", handler.GetUser)
		userGroup.PUT("/update/:id", handler.UpdateUser)
		userGroup.DELETE("/delete/:id", handler.DeleteUser)
		userGroup.GET("/profile/:id", handler.GetUserProfile)
		userGroup.PUT("/profile/:id/update", handler.UpdateUserProfile)
	}

	// Impact
	impactGroup := router.Group("/impact")
	{
		impactGroup.GET("/users/:id", handler.GetUserImpact)
		impactGroup.GET("/groups/:id/", handler.GetGroupImpact)
		impactGroup.GET("/leaderboard/users", handler.GetLeaderboardUsers)
		impactGroup.GET("/leaderboard/groups", handler.GetLeaderboardGroups)
		impactGroup.POST("/donations", handler.CreateDonation)
		impactGroup.GET("/donations/causes", handler.GetCauses)
	}

	auth := router.Group("/api/auth")
    {
        auth.POST("/register", handler.Register)
        auth.POST("/login", handler.Login)
        auth.POST("/logout", handler.Logout)
        auth.GET("/refresh/token", handler.RefreshToken)
    }

	return router
}
