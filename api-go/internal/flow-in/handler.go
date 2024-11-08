package flowin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlowInHandler struct {
	service *FlowInService
}

func NewFlowInHandler(service *FlowInService) *FlowInHandler {
	return &FlowInHandler{service: service}
}

// Get
func (h *FlowInHandler) GetFlowIn(c *gin.Context) {
	flow_in, err := h.service.FetchFlowIn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"flow_in": flow_in})
}

func (h *FlowInHandler) GetFlowInByID(c *gin.Context) {
	id := c.Param("id")

	flow_in, err := h.service.FetchFlowInByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flow_in": flow_in})
}

// Create
func (h *FlowInHandler) CreateFlowIn(c *gin.Context) {
	var flow_in FlowIn
	if err := c.ShouldBindJSON(&flow_in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	createdFlowIn, err := h.service.CreateFlowIn(flow_in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"flow_in": createdFlowIn})
}

// Update
func (h *FlowInHandler) UpdateFlowIn(c *gin.Context) {
	id := c.Param("id")
	var flow_in FlowIn
	if err := c.ShouldBindJSON(&flow_in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedFlowIn, err := h.service.UpdateFlowIn(id, flow_in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flow_in": updatedFlowIn})
}

// Delete
func (h *FlowInHandler) DeleteFlowIn(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteFlowIn(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
