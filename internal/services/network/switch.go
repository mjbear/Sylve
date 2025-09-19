// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package network

import (
	"fmt"

	networkModels "github.com/alchemillahq/sylve/internal/db/models/network"
)

func (s *Service) GetBridgeNameByIDType(id uint, swType string) (string, error) {
	if swType == "manual" {
		var manualSwitches []networkModels.ManualSwitch
		if err := s.DB.
			Find(&manualSwitches).Error; err != nil {
			return "", err
		}

		for _, sw := range manualSwitches {
			if sw.ID == id {
				return sw.Bridge, nil
			}
		}
	} else if swType == "standard" {
		var standardSwitches []networkModels.StandardSwitch
		if err := s.DB.
			Find(&standardSwitches).Error; err != nil {
			return "", err
		}

		for _, sw := range standardSwitches {
			if sw.ID == id {
				return sw.BridgeName, nil
			}
		}
	}

	return "", fmt.Errorf("switch/bridge with ID %d not found", id)
}
