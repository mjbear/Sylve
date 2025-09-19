// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package libvirt

import vmModels "github.com/alchemillahq/sylve/internal/db/models/vm"

func (s *Service) ModifyWakeOnLan(vmId int, enabled bool) error {
	err := s.DB.
		Model(&vmModels.VM{}).
		Where("vm_id = ?", vmId).
		Update("wo_l", enabled).Error
	return err
}

func (s *Service) ModifyBootOrder(vmId int, startAtBoot bool, bootOrder int) error {
	err := s.DB.
		Model(&vmModels.VM{}).
		Where("vm_id = ?", vmId).
		Updates(map[string]interface{}{
			"start_order":   bootOrder,
			"start_at_boot": startAtBoot,
		}).Error
	return err
}
