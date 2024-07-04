package handler

import (
	pb "api-getaway/genproto/protos"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateHabit(c *gin.Context) {
	req := &pb.CreateHabitRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Habits.CreateHabit(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetHabits(c *gin.Context) {
	userID := c.Param("id")

	req := &pb.GetHabitsRequest{
		UserId: userID,
	}

	resp, err := h.Habits.GetHabits(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) LogHabit(c *gin.Context) {
	var req pb.LogHabitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.Habits.LogHabit(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}


func (h *Handler) GetHabitLogs(c *gin.Context) {
	req := &pb.GetHabitLogsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	habitId := c.Param("id")
	req.HabitId = habitId

	resp, err := h.Habits.GetHabitLogs(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetHabitSuggestions(c *gin.Context) {
	req := &pb.GetHabitSuggestionsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Habits.GetHabitSuggestions(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserHabits(c *gin.Context) {
	req := &pb.GetUserHabitsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId := c.Param("id")
	req.UserId = userId

	resp, err := h.Habits.GetUserHabits(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateHabit(c *gin.Context) {
	req := &pb.UpdateHabitRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.HabitId = c.Param("id")
	if req.HabitId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "HabitId cannot be empty",
		})
		return
	}

	resp, err := h.Habits.UpdateHabit(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteHabit(c *gin.Context) {
	req := &pb.DeleteHabitRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.HabitId = c.Param("id")
	resp, err := h.Habits.DeleteHabit(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
