// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package libvirtHandlers

import (
	"strconv"

	"github.com/alchemillahq/sylve/internal"
	"github.com/alchemillahq/sylve/internal/services/libvirt"
	"github.com/gin-gonic/gin"
)

type ModifyWakeOnLanRequest struct {
	Enabled *bool `json:"enabled"`
}

type ModifyBootOrderRequest struct {
	StartAtBoot *bool `json:"startAtBoot"`
	BootOrder   *int  `json:"bootOrder"`
}

// @Summary Modify Wake-on-LAN of a Virtual Machine
// @Description Modify the Wake-on-LAN configuration of a virtual machine
// @Tags VM
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body ModifyWakeOnLanRequest true "Modify Wake-on-LAN Request"
// @Success 200 {object} internal.APIResponse[any] "Success"
// @Failure 400 {object} internal.APIResponse[any] "Bad Request"
// @Failure 500 {object} internal.APIResponse[any] "Internal Server Error"
// @Router /options/wol/:vmid [put]
func ModifyWakeOnLan(libvirtService *libvirt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		vmId := c.Param("vmid")
		if vmId == "" {
			c.JSON(400, internal.APIResponse[any]{
				Status:  "error",
				Message: "invalid_request",
				Data:    nil,
				Error:   "vmid_not_provided",
			})
			return
		}

		vmIdInt, err := strconv.Atoi(vmId)
		if err != nil {
			c.JSON(400, internal.APIResponse[any]{
				Status:  "error",
				Message: "invalid_request",
				Data:    nil,
				Error:   "invalid_vmid_format",
			})
			return
		}

		var req ModifyWakeOnLanRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, internal.APIResponse[any]{
				Status:  "error",
				Message: "invalid_request",
				Data:    nil,
				Error:   "invalid_request: " + err.Error(),
			})
			return
		}

		enabled := false
		if req.Enabled != nil {
			enabled = *req.Enabled
		}

		if err := libvirtService.ModifyWakeOnLan(vmIdInt, enabled); err != nil {
			c.JSON(500, internal.APIResponse[any]{
				Status:  "error",
				Message: "internal_server_error",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		c.JSON(200, internal.APIResponse[any]{
			Status:  "success",
			Message: "wol_modified",
			Data:    nil,
			Error:   "",
		})
	}
}

// @Summary Modify Boot Order of a Virtual Machine
// @Description Modify the Boot Order configuration of a virtual machine
// @Tags VM
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body ModifyBootOrderRequest true "Modify Boot Order Request"
// @Success 200 {object} internal.APIResponse[any] "Success"
// @Failure 400 {object} internal.APIResponse[any] "Bad Request"
// @Failure 500 {object} internal.APIResponse[any] "Internal Server Error"
// @Router /options/boot-order/:vmid [put]
func ModifyBootOrder(libvirtService *libvirt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		vmId := c.Param("vmid")
		if vmId == "" {
			c.JSON(400, internal.APIResponse[any]{
				Status:  "error",
				Message: "invalid_request",
				Data:    nil,
				Error:   "vmid_not_provided",
			})
			return
		}

		vmIdInt, err := strconv.Atoi(vmId)
		if err != nil {
			c.JSON(400, internal.APIResponse[any]{
				Status:  "error",
				Message: "invalid_request",
				Data:    nil,
				Error:   "invalid_vmid_format",
			})
			return
		}

		var req ModifyBootOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, internal.APIResponse[any]{
				Status:  "error",
				Message: "invalid_request",
				Data:    nil,
				Error:   "invalid_request: " + err.Error(),
			})
			return
		}

		startAtBoot := false
		if req.StartAtBoot != nil {
			startAtBoot = *req.StartAtBoot
		}

		bootOrder := 0
		if req.BootOrder != nil {
			bootOrder = *req.BootOrder
		}

		if err := libvirtService.ModifyBootOrder(vmIdInt, startAtBoot, bootOrder); err != nil {
			c.JSON(500, internal.APIResponse[any]{
				Status:  "error",
				Message: "internal_server_error",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		c.JSON(200, internal.APIResponse[any]{
			Status:  "success",
			Message: "boot_order_modified",
			Data:    nil,
			Error:   "",
		})
	}
}
