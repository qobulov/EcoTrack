package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "api-getaway/genproto"

	"api-getaway/api/handler"
)

func CreateRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	Habits := pb.NewHabitTrackerServiceClient(conn)
	Impact := pb.NewImpactCalculatorClient(conn)
	handler := handler.NewHandler(Habits, Impact)

	router.POST("/api/habits", handler.CreateHabit)
	router.GET("/api/habits/:id", handler.GetHabits)
	router.PUT("/api/habits/:id", handler.UpdateHabit)
	router.DELETE("/api/habits/:id", handler.DeleteHabit)
	router.GET("/api/users/:id/habits", handler.GetUserHabits)
	router.POST("/api/habits/:id/log", handler.LogHabit)
	router.GET("/api/habits/:id/logs", handler.GetHabitLogs)
	router.GET("/api/habits/suggestions", handler.GetHabitSuggestions)
	return router
}