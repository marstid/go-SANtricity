package santricity

import (
	"encoding/json"
	"fmt"
)

type Volume struct {
	Offline           bool   `json:"offline"`
	ExtremeProtection bool   `json:"extremeProtection"`
	VolumeHandle      int    `json:"volumeHandle"`
	RaidLevel         string `json:"raidLevel"`
	SectorOffset      string `json:"sectorOffset"`
	WorldWideName     string `json:"worldWideName"`
	Label             string `json:"label"`
	BlkSize           int    `json:"blkSize"`
	Capacity          string `json:"capacity"`
	ReconPriority     int    `json:"reconPriority"`
	SegmentSize       int    `json:"segmentSize"`
	Action            string `json:"action"`
	Cache             struct {
		Cwob                bool   `json:"cwob"`
		EnterpriseCacheDump bool   `json:"enterpriseCacheDump"`
		MirrorActive        bool   `json:"mirrorActive"`
		MirrorEnable        bool   `json:"mirrorEnable"`
		ReadCacheActive     bool   `json:"readCacheActive"`
		ReadCacheEnable     bool   `json:"readCacheEnable"`
		WriteCacheActive    bool   `json:"writeCacheActive"`
		WriteCacheEnable    bool   `json:"writeCacheEnable"`
		CacheFlushModifier  string `json:"cacheFlushModifier"`
		ReadAheadMultiplier int    `json:"readAheadMultiplier"`
	} `json:"cache"`
	MediaScan struct {
		Enable                 bool `json:"enable"`
		ParityValidationEnable bool `json:"parityValidationEnable"`
	} `json:"mediaScan"`
	VolumeRef        string `json:"volumeRef"`
	Status           string `json:"status"`
	VolumeGroupRef   string `json:"volumeGroupRef"`
	CurrentManager   string `json:"currentManager"`
	PreferredManager string `json:"preferredManager"`
	Perms            struct {
		MapToLUN                        bool `json:"mapToLUN"`
		SnapShot                        bool `json:"snapShot"`
		Format                          bool `json:"format"`
		Reconfigure                     bool `json:"reconfigure"`
		MirrorPrimary                   bool `json:"mirrorPrimary"`
		MirrorSecondary                 bool `json:"mirrorSecondary"`
		CopySource                      bool `json:"copySource"`
		CopyTarget                      bool `json:"copyTarget"`
		Readable                        bool `json:"readable"`
		Writable                        bool `json:"writable"`
		Rollback                        bool `json:"rollback"`
		MirrorSync                      bool `json:"mirrorSync"`
		NewImage                        bool `json:"newImage"`
		AllowDVE                        bool `json:"allowDVE"`
		AllowDSS                        bool `json:"allowDSS"`
		ConcatVolumeMember              bool `json:"concatVolumeMember"`
		FlashReadCache                  bool `json:"flashReadCache"`
		AsyncMirrorPrimary              bool `json:"asyncMirrorPrimary"`
		AsyncMirrorSecondary            bool `json:"asyncMirrorSecondary"`
		PitGroup                        bool `json:"pitGroup"`
		CacheParametersChangeable       bool `json:"cacheParametersChangeable"`
		AllowThinManualExpansion        bool `json:"allowThinManualExpansion"`
		AllowThinGrowthParametersChange bool `json:"allowThinGrowthParametersChange"`
	} `json:"perms"`
	MgmtClientAttribute                         int    `json:"mgmtClientAttribute"`
	DssPreallocEnabled                          bool   `json:"dssPreallocEnabled"`
	DssMaxSegmentSize                           int    `json:"dssMaxSegmentSize"`
	PreReadRedundancyCheckEnabled               bool   `json:"preReadRedundancyCheckEnabled"`
	ProtectionInformationCapable                bool   `json:"protectionInformationCapable"`
	ProtectionType                              string `json:"protectionType"`
	ApplicationTagOwned                         bool   `json:"applicationTagOwned"`
	RepairedBlockCount                          int    `json:"repairedBlockCount"`
	ExtendedUniqueIdentifier                    string `json:"extendedUniqueIdentifier"`
	CacheMirroringValidateProtectionInformation bool   `json:"cacheMirroringValidateProtectionInformation"`
	ExpectedProtectionInformationAppTag         int    `json:"expectedProtectionInformationAppTag"`
	VolumeUse                                   string `json:"volumeUse"`
	VolumeFull                                  bool   `json:"volumeFull"`
	VolumeCopyTarget                            bool   `json:"volumeCopyTarget"`
	VolumeCopySource                            bool   `json:"volumeCopySource"`
	PitBaseVolume                               bool   `json:"pitBaseVolume"`
	AsyncMirrorTarget                           bool   `json:"asyncMirrorTarget"`
	AsyncMirrorSource                           bool   `json:"asyncMirrorSource"`
	RemoteMirrorSource                          bool   `json:"remoteMirrorSource"`
	RemoteMirrorTarget                          bool   `json:"remoteMirrorTarget"`
	DiskPool                                    bool   `json:"diskPool"`
	FlashCached                                 bool   `json:"flashCached"`
	IncreasingBy                                string `json:"increasingBy"`
	Metadata                                    []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"metadata"`
	DataAssurance         bool          `json:"dataAssurance"`
	ObjectType            string        `json:"objectType"`
	Wwn                   string        `json:"wwn"`
	PreferredControllerID string        `json:"preferredControllerId"`
	TotalSizeInBytes      string        `json:"totalSizeInBytes"`
	OnlineVolumeCopy      bool          `json:"onlineVolumeCopy"`
	ListOfMappings        []interface{} `json:"listOfMappings"`
	Mapped                bool          `json:"mapped"`
	CurrentControllerID   string        `json:"currentControllerId"`
	CacheSettings         struct {
		Cwob                bool   `json:"cwob"`
		EnterpriseCacheDump bool   `json:"enterpriseCacheDump"`
		MirrorActive        bool   `json:"mirrorActive"`
		MirrorEnable        bool   `json:"mirrorEnable"`
		ReadCacheActive     bool   `json:"readCacheActive"`
		ReadCacheEnable     bool   `json:"readCacheEnable"`
		WriteCacheActive    bool   `json:"writeCacheActive"`
		WriteCacheEnable    bool   `json:"writeCacheEnable"`
		CacheFlushModifier  string `json:"cacheFlushModifier"`
		ReadAheadMultiplier int    `json:"readAheadMultiplier"`
	} `json:"cacheSettings"`
	ThinProvisioned bool   `json:"thinProvisioned"`
	Name            string `json:"name"`
	ID              string `json:"id"`
}

type VolumeStats struct {
	ObservedTime                       string  `json:"observedTime"`
	ObservedTimeInMS                   string  `json:"observedTimeInMS"`
	LastResetTime                      string  `json:"lastResetTime"`
	LastResetTimeInMS                  string  `json:"lastResetTimeInMS"`
	SourceController                   string  `json:"sourceController"`
	VolumeGroupID                      string  `json:"volumeGroupId"`
	ControllerID                       string  `json:"controllerId"`
	VolumeID                           string  `json:"volumeId"`
	ArrayID                            string  `json:"arrayId"`
	ArrayWwn                           string  `json:"arrayWwn"`
	VolumeGroupWwn                     string  `json:"volumeGroupWwn"`
	VolumeName                         string  `json:"volumeName"`
	VolumeWwn                          string  `json:"volumeWwn"`
	WorkloadID                         string  `json:"workloadId"`
	ReadOps                            float64 `json:"readOps"`
	ReadHitOps                         float64 `json:"readHitOps"`
	ReadHitBytes                       float64 `json:"readHitBytes"`
	ReadTimeTotal                      float64 `json:"readTimeTotal"`
	ReadHitTimeTotal                   float64 `json:"readHitTimeTotal"`
	WriteOps                           float64 `json:"writeOps"`
	WriteCacheHitOps                   float64 `json:"writeCacheHitOps"`
	WriteTimeTotal                     float64 `json:"writeTimeTotal"`
	WriteHitTimeTotal                  float64 `json:"writeHitTimeTotal"`
	ErrRedundancyChkIndeterminateReads float64 `json:"errRedundancyChkIndeterminateReads"`
	ErrRedundancyChkRecoveredReads     float64 `json:"errRedundancyChkRecoveredReads"`
	ErrRedundancyChkUnrecoveredReads   float64 `json:"errRedundancyChkUnrecoveredReads"`
	IdleTime                           float64 `json:"idleTime"`
	OtherOps                           float64 `json:"otherOps"`
	OtherTimeMax                       float64 `json:"otherTimeMax"`
	OtherTimeTotal                     float64 `json:"otherTimeTotal"`
	OtherTimeTotalSq                   float64 `json:"otherTimeTotalSq"`
	ReadBytes                          float64 `json:"readBytes"`
	ReadHitTimeMax                     float64 `json:"readHitTimeMax"`
	ReadHitTimeTotalSq                 float64 `json:"readHitTimeTotalSq"`
	ReadTimeMax                        float64 `json:"readTimeMax"`
	ReadTimeTotalSq                    float64 `json:"readTimeTotalSq"`
	WriteBytes                         float64 `json:"writeBytes"`
	WriteHitBytes                      float64 `json:"writeHitBytes"`
	WriteHitOps                        float64 `json:"writeHitOps"`
	WriteHitTimeMax                    float64 `json:"writeHitTimeMax"`
	WriteHitTimeTotalSq                float64 `json:"writeHitTimeTotalSq"`
	WriteTimeMax                       float64 `json:"writeTimeMax"`
	WriteTimeTotalSq                   float64 `json:"writeTimeTotalSq"`
	QueueDepthTotal                    float64 `json:"queueDepthTotal"`
	QueueDepthMax                      float64 `json:"queueDepthMax"`
	RandomIosTotal                     float64 `json:"randomIosTotal"`
	RandomBytesTotal                   float64 `json:"randomBytesTotal"`
	CacheWriteWaitHitIops              float64 `json:"cacheWriteWaitHitIops"`
	CacheWriteWaitHitBytes             float64 `json:"cacheWriteWaitHitBytes"`
	FullStripeWriteBytes               float64 `json:"fullStripeWriteBytes"`
	TotalIosShipped                    float64 `json:"totalIosShipped"`
	TotalBlksEvicted                   float64 `json:"totalBlksEvicted"`
	CacheBlksInUse                     float64 `json:"cacheBlksInUse"`
	PrefetchHitBytes                   float64 `json:"prefetchHitBytes"`
	PrefetchMissBytes                  string  `json:"prefetchMissBytes"`
	FlashCacheReadHitOps               float64 `json:"flashCacheReadHitOps"`
	FlashCacheReadHitBytes             float64 `json:"flashCacheReadHitBytes"`
	FlashCacheReadHitTimeTotal         float64 `json:"flashCacheReadHitTimeTotal"`
	FlashCacheReadHitTimeMax           float64 `json:"flashCacheReadHitTimeMax"`
	FlashCacheReadHitTimeTotalSq       float64 `json:"flashCacheReadHitTimeTotalSq"`
}

func (c *Client) GetVolumes(sysid string) (vol []Volume, err error) {
	uri := "/devmgr/v2/storage-systems/" + sysid + "/volumes"

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

func (c *Client) GetVolume(sysid string, volid string) (vol Volume, err error) {
	uri := "/devmgr/v2/storage-systems/" + sysid + "/volumes/" + volid

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

func (c *Client) GetVolumeStats(sysid string) (volstat []VolumeStats, err error) {
	// Make sure not to get the cached stats
	uri := "/devmgr/v2/storage-systems/1/volume-statistics?usecache=false"

	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return volstat, err
	}

	err = json.Unmarshal(data, &volstat)
	if err != nil {
		return volstat, err
	}

	return volstat, err

}
