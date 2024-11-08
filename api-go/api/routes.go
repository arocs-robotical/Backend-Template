package api

import (
	flowin "api-go/internal/flow-in"
	flowout "api-go/internal/flow-out"
	product "api-go/internal/product"
	robotallocation "api-go/internal/robot-allocation"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(pocketbaseURL string) *gin.Engine {
	r := gin.Default()

	productService := product.NewProductService(pocketbaseURL)
	productHandler := product.NewProductHandler(productService)

	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.POST("/products", productHandler.CreateProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	flowInService := flowin.NewFlowInService(pocketbaseURL)
	flowInHandler := flowin.NewFlowInHandler(flowInService)

	r.GET("/flow_in", flowInHandler.GetFlowIn)
	r.GET("/flow_in/:id", flowInHandler.GetFlowInByID)
	r.POST("/flow_in", flowInHandler.CreateFlowIn)
	r.PUT("/flow_in/:id", flowInHandler.UpdateFlowIn)
	r.DELETE("/flow_in/:id", flowInHandler.DeleteFlowIn)

	FlowOutService := flowout.NewFlowOutService(pocketbaseURL)
	FlowOutHandler := flowout.NewFlowOutHandler(FlowOutService)

	r.GET("/flow_out", FlowOutHandler.GetFlowOut)
	r.GET("/flow_out/:id", FlowOutHandler.GetFlowOutByID)
	r.POST("/flow_out", FlowOutHandler.CreateFlowOut)
	r.PUT("/flow_out/:id", FlowOutHandler.UpdateFlowOut)
	r.DELETE("/flow_out/:id", FlowOutHandler.DeleteFlowOut)

	RobotAllocationService := robotallocation.NewRobotAllocationService(pocketbaseURL)
	RobotAllocationHandler := robotallocation.NewRobotAllocationHandler(RobotAllocationService)

	r.GET("/robot_allocation", RobotAllocationHandler.GetRobotAllocation)
	r.GET("/robot_allocation/:id", RobotAllocationHandler.GetRobotAllocationByID)
	r.POST("/robot_allocation", RobotAllocationHandler.CreateRobotAllocation)
	r.PUT("/robot_allocation/:id", RobotAllocationHandler.UpdateRobotAllocation)
	r.DELETE("/robot_allocation/:id", RobotAllocationHandler.DeleteRobotAllocation)

	return r
}
