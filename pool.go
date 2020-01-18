package santricity

import (
	"encoding/json"
	"fmt"
)

type Pool struct {
	SequenceNum                       int    `json:"sequenceNum"`
	Offline                           bool   `json:"offline"`
	RaidLevel                         string `json:"raidLevel"`
	WorldWideName                     string `json:"worldWideName"`
	VolumeGroupRef                    string `json:"volumeGroupRef"`
	Reserved1                         string `json:"reserved1"`
	Reserved2                         string `json:"reserved2"`
	TrayLossProtection                bool   `json:"trayLossProtection"`
	Label                             string `json:"label"`
	State                             string `json:"state"`
	SpindleSpeedMatch                 bool   `json:"spindleSpeedMatch"`
	SpindleSpeed                      int    `json:"spindleSpeed"`
	IsInaccessible                    bool   `json:"isInaccessible"`
	SecurityType                      string `json:"securityType"`
	DrawerLossProtection              bool   `json:"drawerLossProtection"`
	ProtectionInformationCapable      bool   `json:"protectionInformationCapable"`
	ProtectionInformationCapabilities struct {
		ProtectionInformationCapable bool   `json:"protectionInformationCapable"`
		ProtectionType               string `json:"protectionType"`
	} `json:"protectionInformationCapabilities"`
	VolumeGroupData struct {
		Type         string `json:"type"`
		DiskPoolData struct {
			ReconstructionReservedDriveCount        int    `json:"reconstructionReservedDriveCount"`
			ReconstructionReservedAmt               string `json:"reconstructionReservedAmt"`
			ReconstructionReservedDriveCountCurrent int    `json:"reconstructionReservedDriveCountCurrent"`
			PoolUtilizationWarningThreshold         int    `json:"poolUtilizationWarningThreshold"`
			PoolUtilizationCriticalThreshold        int    `json:"poolUtilizationCriticalThreshold"`
			PoolUtilizationState                    string `json:"poolUtilizationState"`
			UnusableCapacity                        string `json:"unusableCapacity"`
			DegradedReconstructPriority             string `json:"degradedReconstructPriority"`
			CriticalReconstructPriority             string `json:"criticalReconstructPriority"`
			BackgroundOperationPriority             string `json:"backgroundOperationPriority"`
			AllocGranularity                        string `json:"allocGranularity"`
			MinimumDriveCount                       int    `json:"minimumDriveCount"`
		} `json:"diskPoolData"`
	} `json:"volumeGroupData"`
	Usage                  string `json:"usage"`
	DriveBlockFormat       string `json:"driveBlockFormat"`
	ReservedSpaceAllocated bool   `json:"reservedSpaceAllocated"`
	SecurityLevel          string `json:"securityLevel"`
	UsedSpace              string `json:"usedSpace"`
	TotalRaidedSpace       string `json:"totalRaidedSpace"`
	Extents                []struct {
		SectorOffset   string `json:"sectorOffset"`
		RawCapacity    string `json:"rawCapacity"`
		RaidLevel      string `json:"raidLevel"`
		VolumeGroupRef string `json:"volumeGroupRef"`
		FreeExtentRef  string `json:"freeExtentRef"`
		Reserved1      string `json:"reserved1"`
		Reserved2      string `json:"reserved2"`
	} `json:"extents"`
	LargestFreeExtentSize  string `json:"largestFreeExtentSize"`
	RaidStatus             string `json:"raidStatus"`
	FreeSpace              string `json:"freeSpace"`
	DrivePhysicalType      string `json:"drivePhysicalType"`
	DriveMediaType         string `json:"driveMediaType"`
	NormalizedSpindleSpeed string `json:"normalizedSpindleSpeed"`
	DiskPool               bool   `json:"diskPool"`
	ID                     string `json:"id"`
	Name                   string `json:"name"`
}

func (c *Client) GetPools(sysid string) (vol []Pool, err error) {
	uri := "/devmgr/v2/storage-systems/" + sysid + "/storage-pools"

	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return vol, err
	}

	err = json.Unmarshal(data, &vol)
	if err != nil {
		return vol, err
	}

	return vol, err

}
