package flowout

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlowOutHandler struct {
	service *FlowOutService
}

func NewFlowOutHandler(service *FlowOutService) *FlowOutHandler {
	return &FlowOutHandler{service: service}
}

// Get
func (h *FlowOutHandler) GetFlowOut(c *gin.Context) {
	flow_out, err := h.service.FetchFlowOut()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"flow_out": flow_out})
}

func (h *FlowOutHandler) GetFlowOutByID(c *gin.Context) {
	id := c.Param("id")

	flow_out, err := h.service.FetchFlowOutByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flow_out": flow_out})
}

// Create
func (h *FlowOutHandler) CreateFlowOut(c *gin.Context) {
	var flow_out FlowOut
	if err := c.ShouldBindJSON(&flow_out); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	createdFlowOut, err := h.service.CreateFlowOut(flow_out)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"flow_out": createdFlowOut})
}

// Update
func (h *FlowOutHandler) UpdateFlowOut(c *gin.Context) {
	id := c.Param("id")
	var flow_out FlowOut
	if err := c.ShouldBindJSON(&flow_out); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedFlowOut, err := h.service.UpdateFlowOut(id, flow_out)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"flow_out": updatedFlowOut})
}

// Delete
func (h *FlowOutHandler) DeleteFlowOut(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteFlowOut(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
