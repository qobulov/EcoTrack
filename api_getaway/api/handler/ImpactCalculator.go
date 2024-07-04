package handler

import (
	pb "api-getaway/genproto/protos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CalculateCarbonFootprint(c *gin.Context) {
	req := &pb.CalculateCarbonFootprintRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Impact.CalculateCarbonFootprint(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserImpact(c *gin.Context) {
	req := &pb.GetUserImpactRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID := c.Param("id")
	req.UserId = userID
	resp, err := h.Impact.GetUserImpact(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetGroupImpact(c *gin.Context) {
	req := &pb.GetGroupImpactRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	groupID := c.Param("id")
	req.GroupId = groupID
	resp, err := h.Impact.GetGroupImpact(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetLeaderboardUsers(c *gin.Context) {
	req := &pb.GetLeaderboardRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Impact.GetUserLeaderboard(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetLeaderboardGroups(c *gin.Context) {
	req := &pb.GetLeaderboardRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Impact.GetGroupLeaderboard(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateDonation(c *gin.Context) {
	req := &pb.DonateToCauseRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	resp, err := h.Impact.DonateToCause(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCauses(c *gin.Context) {
	req := &pb.GetDonationCausesRequest{}
	resp, err := h.Impact.GetDonationCauses(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}