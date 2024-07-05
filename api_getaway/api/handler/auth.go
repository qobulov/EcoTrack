package handler

import (
	"context"
	"net/http"

	pb "api-getaway/genproto/protos"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
    req := &pb.RegisterRequest{}
    if err := c.ShouldBindJSON(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    resp, err := h.AuthService.Register(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c *gin.Context) {
    req := &pb.LoginRequest{}
    if err := c.ShouldBindJSON(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    resp, err := h.AuthService.Login(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
}

func (h *Handler) Logout(c *gin.Context) {
    req := &pb.LogoutRequest{}
    if err := c.ShouldBindJSON(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    resp, err := h.AuthService.Logout(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
}

func (h *Handler) RefreshToken(c *gin.Context) {
    req := &pb.RefreshTokenRequest{}
    if err := c.ShouldBindJSON(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    resp, err := h.AuthService.RefreshToken(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
}