package models

import "time"

type VirtualMachine struct {
	Status struct {
		Description      string `json:"description"`
		State            string `json:"state"`
		ExecutionContext struct {
			TaskUuids []string `json:"task_uuids"`
		} `json:"execution_context"`
		ClusterReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			Uuid string `json:"uuid"`
		} `json:"cluster_reference"`
		Resources struct {
			NumThreadsPerCore int  `json:"num_threads_per_core"`
			IsAgentVm         bool `json:"is_agent_vm"`
			MemorySizeMib     int  `json:"memory_size_mib"`
			BootConfig        struct {
				BootDevice struct {
					DiskAddress struct {
						DeviceIndex int    `json:"device_index"`
						AdapterType string `json:"adapter_type"`
					} `json:"disk_address"`
				} `json:"boot_device"`
				BootType string `json:"boot_type"`
			} `json:"boot_config"`
			DiskList []struct {
				Uuid                  string `json:"uuid"`
				IsMigrationInProgress bool   `json:"is_migration_in_progress"`
				DiskSizeBytes         int64  `json:"disk_size_bytes,omitempty"`
				StorageConfig         struct {
					StorageContainerReference struct {
						Kind string `json:"kind"`
						Uuid string `json:"uuid"`
						Name string `json:"name"`
					} `json:"storage_container_reference"`
				} `json:"storage_config,omitempty"`
				DeviceProperties struct {
					DeviceType  string `json:"device_type"`
					DiskAddress struct {
						DeviceIndex int    `json:"device_index"`
						AdapterType string `json:"adapter_type"`
					} `json:"disk_address"`
				} `json:"device_properties"`
				DiskSizeMib int `json:"disk_size_mib,omitempty"`
			} `json:"disk_list"`
			SerialPortList        []interface{} `json:"serial_port_list"`
			PowerState            string        `json:"power_state"`
			NumVcpusPerSocket     int           `json:"num_vcpus_per_socket"`
			NumSockets            int           `json:"num_sockets"`
			ProtectionType        string        `json:"protection_type"`
			GpuList               []interface{} `json:"gpu_list"`
			MachineType           string        `json:"machine_type"`
			HardwareClockTimezone string        `json:"hardware_clock_timezone"`
			PowerStateMechanism   struct {
				GuestTransitionConfig struct {
				} `json:"guest_transition_config"`
			} `json:"power_state_mechanism"`
			VgaConsoleEnabled bool `json:"vga_console_enabled"`
			VnumaConfig       struct {
				NumVnumaNodes int `json:"num_vnuma_nodes"`
			} `json:"vnuma_config"`
			NicList []struct {
				NicType        string `json:"nic_type"`
				Uuid           string `json:"uuid"`
				IpEndpointList []struct {
					Ip   string `json:"ip"`
					Type string `json:"type"`
				} `json:"ip_endpoint_list"`
				VlanMode        string `json:"vlan_mode"`
				MacAddress      string `json:"mac_address"`
				SubnetReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
					Uuid string `json:"uuid"`
				} `json:"subnet_reference"`
				IsConnected     bool          `json:"is_connected"`
				TrunkedVlanList []interface{} `json:"trunked_vlan_list"`
			} `json:"nic_list"`
			HostReference struct {
				Kind string `json:"kind"`
				Uuid string `json:"uuid"`
				Name string `json:"name"`
			} `json:"host_reference"`
			GuestTools struct {
				NutanixGuestTools struct {
					AvailableVersion      string   `json:"available_version"`
					NgtState              string   `json:"ngt_state"`
					IsoMountState         string   `json:"iso_mount_state"`
					State                 string   `json:"state"`
					EnabledCapabilityList []string `json:"enabled_capability_list"`
					IsReachable           bool     `json:"is_reachable"`
				} `json:"nutanix_guest_tools"`
			} `json:"guest_tools"`
			HypervisorType       string `json:"hypervisor_type"`
			EnableCpuPassthrough bool   `json:"enable_cpu_passthrough"`
			DisableBranding      bool   `json:"disable_branding"`
		} `json:"resources"`
		Name string `json:"name"`
	} `json:"status"`
	Spec struct {
		ClusterReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			Uuid string `json:"uuid"`
		} `json:"cluster_reference"`
		Name      string `json:"name"`
		Resources struct {
			NumThreadsPerCore int `json:"num_threads_per_core"`
			VnumaConfig       struct {
				NumVnumaNodes int `json:"num_vnuma_nodes"`
			} `json:"vnuma_config"`
			SerialPortList []interface{} `json:"serial_port_list"`
			NicList        []struct {
				NicType         string        `json:"nic_type"`
				Uuid            string        `json:"uuid"`
				IpEndpointList  []interface{} `json:"ip_endpoint_list"`
				VlanMode        string        `json:"vlan_mode"`
				MacAddress      string        `json:"mac_address"`
				SubnetReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
					Uuid string `json:"uuid"`
				} `json:"subnet_reference"`
				IsConnected     bool          `json:"is_connected"`
				TrunkedVlanList []interface{} `json:"trunked_vlan_list"`
			} `json:"nic_list"`
			NumVcpusPerSocket int `json:"num_vcpus_per_socket"`
			GuestTools        struct {
				NutanixGuestTools struct {
					IsoMountState         string   `json:"iso_mount_state"`
					NgtState              string   `json:"ngt_state"`
					State                 string   `json:"state"`
					EnabledCapabilityList []string `json:"enabled_capability_list"`
				} `json:"nutanix_guest_tools"`
			} `json:"guest_tools"`
			NumSockets           int           `json:"num_sockets"`
			EnableCpuPassthrough bool          `json:"enable_cpu_passthrough"`
			GpuList              []interface{} `json:"gpu_list"`
			IsAgentVm            bool          `json:"is_agent_vm"`
			MemorySizeMib        int           `json:"memory_size_mib"`
			BootConfig           struct {
				BootDevice struct {
					DiskAddress struct {
						DeviceIndex int    `json:"device_index"`
						AdapterType string `json:"adapter_type"`
					} `json:"disk_address"`
				} `json:"boot_device"`
				BootType string `json:"boot_type"`
			} `json:"boot_config"`
			HardwareClockTimezone string `json:"hardware_clock_timezone"`
			DisableBranding       bool   `json:"disable_branding"`
			PowerState            string `json:"power_state"`
			MachineType           string `json:"machine_type"`
			PowerStateMechanism   struct {
			} `json:"power_state_mechanism"`
			VgaConsoleEnabled bool `json:"vga_console_enabled"`
			DiskList          []struct {
				StorageConfig struct {
					StorageContainerReference struct {
						Kind string `json:"kind"`
						Uuid string `json:"uuid"`
						Name string `json:"name"`
					} `json:"storage_container_reference"`
				} `json:"storage_config,omitempty"`
				DeviceProperties struct {
					DiskAddress struct {
						DeviceIndex int    `json:"device_index"`
						AdapterType string `json:"adapter_type"`
					} `json:"disk_address"`
					DeviceType string `json:"device_type"`
				} `json:"device_properties"`
				Uuid          string `json:"uuid"`
				DiskSizeBytes int64  `json:"disk_size_bytes,omitempty"`
				DiskSizeMib   int    `json:"disk_size_mib,omitempty"`
			} `json:"disk_list"`
		} `json:"resources"`
		Description string `json:"description"`
	} `json:"spec"`
	Metadata struct {
		LastUpdateTime   time.Time `json:"last_update_time"`
		Kind             string    `json:"kind"`
		Uuid             string    `json:"uuid"`
		ProjectReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			Uuid string `json:"uuid"`
		} `json:"project_reference"`
		CreationTime      time.Time `json:"creation_time"`
		SpecVersion       int       `json:"spec_version"`
		CategoriesMapping struct {
			BluePrism []string `json:"BluePrism"`
		} `json:"categories_mapping"`
		EntityVersion  string `json:"entity_version"`
		OwnerReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			Uuid string `json:"uuid"`
		} `json:"owner_reference"`
		Categories struct {
			BluePrism string `json:"BluePrism"`
		} `json:"categories"`
	} `json:"metadata"`
}
