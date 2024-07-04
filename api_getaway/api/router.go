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
	handler := handler.NewHandler(Habits, Impact)

	router.POST("/api/habits", handler.CreateHabit)
	router.GET("/api/habits/:id", handler.GetHabits)
	router.PUT("/api/habits/:id", handler.UpdateHabit)
	router.DELETE("/api/habits/:id", handler.DeleteHabit)
	router.GET("/api/users/:id/habits", handler.GetUserHabits)
	router.POST("/api/habits/log", handler.LogHabit)
	router.GET("/api/habits/:id/logs", handler.GetHabitLogs)
	router.GET("/api/habits/suggestions", handler.GetHabitSuggestions)

	router.POST("/api/impact/carbon-footprint", handler.CalculateCarbonFootprint)
	router.GET("/api/users/:id/impact", handler.GetUserImpact)
	router.GET("/api/groups/:id/impact", handler.GetGroupImpact)
	router.GET("/api/leaderboard/users", handler.GetLeaderboardUsers)
	router.GET("/api/leaderboard/groups", handler.GetLeaderboardGroups)
	router.POST("/api/donations", handler.CreateDonation)
	router.GET("/api/donations/causes", handler.GetCauses)

	return router
}