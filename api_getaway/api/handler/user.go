package handler

import (
	"net/http"

	pb "api-getaway/genproto/protos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	req := &pb.CreateUserRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.User.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetUser(c *gin.Context) {
	req := &pb.GetUserRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = c.Param("id")
	resp, err := h.User.GetUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	req := &pb.UpdateUserRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = c.Param("id")
	resp, err := h.User.UpdateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	req := &pb.DeleteUserRequest{}
	req.UserId = c.Param("id")
	resp, err := h.User.DeleteUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserProfile(c *gin.Context) {
	req := &pb.GetUserProfileRequest{}
	req.UserId = c.Param("id")
	resp, err := h.User.GetUserProfile(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUserProfile(c *gin.Context) {
	req := &pb.UpdateUserProfileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = c.Param("id")
	resp, err := h.User.UpdateUserProfile(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
