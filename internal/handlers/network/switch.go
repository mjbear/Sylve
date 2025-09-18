package networkHandlers

import (
	"net/http"

	"github.com/alchemillahq/sylve/internal"
	networkModels "github.com/alchemillahq/sylve/internal/db/models/network"
	"github.com/alchemillahq/sylve/internal/services/network"
	"github.com/gin-gonic/gin"
)

type ListSwitchResponse struct {
	Standard []networkModels.StandardSwitch `json:"standard"`
	Manual   []networkModels.ManualSwitch   `json:"manual"`
}

// @Summary List Network Switches
// @Description List all network switches on the system
// @Tags Network
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} internal.APIResponse[ListSwitchResponse] "Success"
// @Failure 400 {object} internal.APIResponse[any] "Bad Request"
// @Failure 500 {object} internal.APIResponse[any] "Internal Server Error"
// @Router /network/switch [get]
func ListSwitches(networkService *network.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var response ListSwitchResponse
		standardSwitches, err := networkService.GetStandardSwitches()

		if err != nil {
			c.JSON(http.StatusInternalServerError, internal.APIResponse[any]{
				Status:  "error",
				Message: "failed_to_get_switches",
				Error:   err.Error(),
				Data:    nil,
			})
			return
		}

		manualSwitches, err := networkService.GetManualSwitches()
		if err != nil {
			c.JSON(http.StatusInternalServerError, internal.APIResponse[any]{
				Status:  "error",
				Message: "failed_to_get_switches",
				Error:   err.Error(),
				Data:    nil,
			})
			return
		}

		response.Standard = standardSwitches
		response.Manual = manualSwitches

		c.JSON(http.StatusOK, internal.APIResponse[ListSwitchResponse]{
			Status:  "success",
			Message: "switches_list",
			Error:   "",
			Data:    response,
		})
	}
}
