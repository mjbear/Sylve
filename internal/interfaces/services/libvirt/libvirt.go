// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package libvirtServiceInterfaces

import vmModels "github.com/alchemillahq/sylve/internal/db/models/vm"

type LibvirtServiceInterface interface {
	CheckVersion() error
	StartTPM() error

	ListStoragePools() ([]StoragePool, error)
	CreateStoragePool(name string) error
	DeleteStoragePool(name string) error
	RescanStoragePools() error

	NetworkDetach(vmId int, networkId int) error
	NetworkAttach(vmId int, switchName string, emulation string, macObjId uint) error
	FindAndChangeMAC(vmId int, oldMac string, newMac string) error

	StoreVMUsage() error

	FindISOByUUID(uuid string, includeImg bool) (string, error)

	GetLvDomain(vmId int) (*LvDomain, error)
	IsDomainInactive(vmId int) (bool, error)

	FindVmByMac(mac string) (vmModels.VM, error)
	WolTasks()
}

type LvDomain struct {
	ID     int32  `json:"id"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type SimpleList struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	VMID  int    `json:"vmId"`
	State string `json:"state"`
}
