package handler

import (
	pb "api-getaway/genproto/protos"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (k *Handler) CreateGroup(c *gin.Context){
	rep := &pb.CreateGroupRequest{}

	if err := c.ShouldBindJSON(rep); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := k.Community.CreateGroup(c, rep)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (k *Handler) GetGroup(c *gin.Context){
	id := c.Param("id")

	rep := &pb.GetGroupRequest{
		Id: id,
	}
	resp, err := k.Community.GetGroup(c,rep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateGroup(c *gin.Context) {
	req := &pb.UpdateGroupRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Group cannot be empty",
		})
		return
	}

	resp, err := h.Community.UpdateGroup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteGroup(c *gin.Context) {
	req := &pb.DeleteGroupRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Id = c.Param("id")
	resp, err := h.Community.DeleteGroup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ListGroups(c *gin.Context) {
	req := &pb.ListGroupsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Community.ListGroups(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinGroup(c *gin.Context) {
	req := &pb.JoinGroupRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Community.JoinGroup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateGroupMemberRole(c *gin.Context) {
	req := &pb.UpdateGroupMemberRoleRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Community.UpdateGroupMemberRole(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}


func (h *Handler) LeaveGroup(c *gin.Context) {
	req := &pb.LeaveGroupRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Community.LeaveGroup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreatePost(c *gin.Context) {
	req := &pb.CreatePostRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Community.CreatePost(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetPost(c *gin.Context) {
	req := &pb.GetPostRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Id = c.Param("id")
	resp, err := h.Community.GetPost(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	req := &pb.UpdatePostRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Id = c.Param("id")
	resp, err := h.Community.UpdatePost(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeletePost(c *gin.Context) {
	req := &pb.DeletePostRequest{}
	req.Id = c.Param("id")
	resp, err := h.Community.DeletePost(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ListGroupPosts(c *gin.Context) {
	req := &pb.ListGroupPostsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Community.ListGroupPosts(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateComment(c *gin.Context) {
	req := &pb.CreateCommentRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.Community.CreateComment(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetPostComments(c *gin.Context) {
	req := &pb.GetPostCommentsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.PostId = c.Param("post_id")
	resp, err := h.Community.GetPostComments(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

