package santricity

import (
	"encoding/json"
	"fmt"
)

// System type generated with https://mholt.github.io/json-to-go/
type System struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Wwn               string   `json:"wwn"`
	PasswordStatus    string   `json:"passwordStatus"`
	PasswordSet       bool     `json:"passwordSet"`
	Status            string   `json:"status"`
	CertificateStatus string   `json:"certificateStatus"`
	IP1               string   `json:"ip1"`
	IP2               string   `json:"ip2"`
	ManagementPaths   []string `json:"managementPaths"`
	Controllers       []struct {
		ControllerID      string      `json:"controllerId"`
		IPAddresses       []string    `json:"ipAddresses"`
		CertificateStatus interface{} `json:"certificateStatus"`
	} `json:"controllers"`
	DriveCount              int           `json:"driveCount"`
	TrayCount               int           `json:"trayCount"`
	TraceEnabled            bool          `json:"traceEnabled"`
	Types                   string        `json:"types"`
	Model                   string        `json:"model"`
	MetaTags                []interface{} `json:"metaTags"`
	HotSpareSize            string        `json:"hotSpareSize"`
	UsedPoolSpace           string        `json:"usedPoolSpace"`
	FreePoolSpace           string        `json:"freePoolSpace"`
	UnconfiguredSpace       string        `json:"unconfiguredSpace"`
	DriveTypes              []string      `json:"driveTypes"`
	HostSpareCountInStandby int           `json:"hostSpareCountInStandby"`
	HotSpareCount           int           `json:"hotSpareCount"`
	HostSparesUsed          int           `json:"hostSparesUsed"`
	BootTime                string        `json:"bootTime"`
	FwVersion               string        `json:"fwVersion"`
	AppVersion              string        `json:"appVersion"`
	BootVersion             string        `json:"bootVersion"`
	NvsramVersion           string        `json:"nvsramVersion"`
	ChassisSerialNumber     string        `json:"chassisSerialNumber"`
	AccessVolume            struct {
		Enabled               bool   `json:"enabled"`
		VolumeHandle          int    `json:"volumeHandle"`
		Capacity              string `json:"capacity"`
		AccessVolumeRef       string `json:"accessVolumeRef"`
		Reserved1             string `json:"reserved1"`
		ObjectType            string `json:"objectType"`
		Wwn                   string `json:"wwn"`
		PreferredControllerID string `json:"preferredControllerId"`
		TotalSizeInBytes      string `json:"totalSizeInBytes"`
		ListOfMappings        []struct {
			LunMappingRef string `json:"lunMappingRef"`
			Lun           int    `json:"lun"`
			Ssid          int    `json:"ssid"`
			Perms         int    `json:"perms"`
			VolumeRef     string `json:"volumeRef"`
			Type          string `json:"type"`
			MapRef        string `json:"mapRef"`
			ID            string `json:"id"`
		} `json:"listOfMappings"`
		Mapped              bool   `json:"mapped"`
		CurrentControllerID string `json:"currentControllerId"`
		Name                string `json:"name"`
		ID                  string `json:"id"`
	} `json:"accessVolume"`
	UnconfiguredSpaceByDriveType struct {
	} `json:"unconfiguredSpaceByDriveType"`
	MediaScanPeriod                  int      `json:"mediaScanPeriod"`
	DriveChannelPortDisabled         bool     `json:"driveChannelPortDisabled"`
	RecoveryModeEnabled              bool     `json:"recoveryModeEnabled"`
	AutoLoadBalancingEnabled         bool     `json:"autoLoadBalancingEnabled"`
	HostConnectivityReportingEnabled bool     `json:"hostConnectivityReportingEnabled"`
	RemoteMirroringEnabled           bool     `json:"remoteMirroringEnabled"`
	FcRemoteMirroringState           string   `json:"fcRemoteMirroringState"`
	AsupEnabled                      bool     `json:"asupEnabled"`
	SecurityKeyEnabled               bool     `json:"securityKeyEnabled"`
	ExternalKeyEnabled               bool     `json:"externalKeyEnabled"`
	LastContacted                    string   `json:"lastContacted"`
	DefinedPartitionCount            int      `json:"definedPartitionCount"`
	SimplexModeEnabled               bool     `json:"simplexModeEnabled"`
	SupportedManagementPorts         []string `json:"supportedManagementPorts"`
	UnconfiguredSpaceAsStrings       string   `json:"unconfiguredSpaceAsStrings"`
	UsedPoolSpaceAsString            string   `json:"usedPoolSpaceAsString"`
	FreePoolSpaceAsString            string   `json:"freePoolSpaceAsString"`
	HotSpareSizeAsString             string   `json:"hotSpareSizeAsString"`
}

func (c *Client) GetSystems() (sys []System, err error) {
	uri := "/devmgr/v2/storage-systems"

	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return sys, err
	}

	err = json.Unmarshal(data, &sys)
	if err != nil {
		return sys, err
	}

	return sys, err

}

func (c *Client) GetSystem(id string) (sys System, err error) {
	uri := "/devmgr/v2/storage-systems/" + id

	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return sys, err
	}

	err = json.Unmarshal(data, &sys)
	if err != nil {
		return sys, err
	}

	return sys, err

}
