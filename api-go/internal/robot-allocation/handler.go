package robotallocation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RobotAllocationHandler struct {
	service *RobotAllocationService
}

func NewRobotAllocationHandler(service *RobotAllocationService) *RobotAllocationHandler {
	return &RobotAllocationHandler{service: service}
}

// Get
func (h *RobotAllocationHandler) GetRobotAllocation(c *gin.Context) {
	robot_allocation, err := h.service.FetchRobotAllocation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"robot_allocation": robot_allocation})
}

func (h *RobotAllocationHandler) GetRobotAllocationByID(c *gin.Context) {
	id := c.Param("id")

	robot_allocation, err := h.service.FetchRobotAllocationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"robot_allocation": robot_allocation})
}

// Create
func (h *RobotAllocationHandler) CreateRobotAllocation(c *gin.Context) {
	var robot_allocation RobotAllocation
	if err := c.ShouldBindJSON(&robot_allocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	createdRobotAllocation, err := h.service.CreateRobotAllocation(robot_allocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"robot_allocation": createdRobotAllocation})
}

// Update
func (h *RobotAllocationHandler) UpdateRobotAllocation(c *gin.Context) {
	id := c.Param("id")
	var robot_allocation RobotAllocation
	if err := c.ShouldBindJSON(&robot_allocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedRobotAllocation, err := h.service.UpdateRobotAllocation(id, robot_allocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"robot_allocation": updatedRobotAllocation})
}

// Delete
func (h *RobotAllocationHandler) DeleteRobotAllocation(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteRobotAllocation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
