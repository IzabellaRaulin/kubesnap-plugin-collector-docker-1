package fs

import "time"

type DeviceInfo struct {
	Device string
	Major  uint
	Minor  uint
}

type FsType string

func (ft FsType) String() string {
	return string(ft)
}

const (
	ZFS          FsType = "zfs"
	DeviceMapper FsType = "devicemapper"
	VFS          FsType = "vfs"
)

type Fs struct {
	DeviceInfo
	Type       FsType
	Capacity   uint64
	Free       uint64
	Available  uint64
	Inodes     uint64
	InodesFree uint64
	DiskStats  DiskStats
}

type DiskStats struct {
	ReadsCompleted  uint64
	ReadsMerged     uint64
	SectorsRead     uint64
	ReadTime        uint64
	WritesCompleted uint64
	WritesMerged    uint64
	SectorsWritten  uint64
	WriteTime       uint64
	IoInProgress    uint64
	IoTime          uint64
	WeightedIoTime  uint64
}

type FsInfo interface {
	// Returns capacity and free space, in bytes, of all the ext2, ext3, ext4 filesystems on the host.
	GetGlobalFsInfo() ([]Fs, error)

	// Returns capacity and free space, in bytes, of the set of mounts passed.
	GetFsInfoForPath(mountSet map[string]struct{}) ([]Fs, error)

	// Returns number of bytes occupied by 'dir'.
	GetDirUsage(dir string, timeout time.Duration) (uint64, error)

	// Returns the block device info of the filesystem on which 'dir' resides.
	GetDirFsDevice(dir string) (*DeviceInfo, error)

	// Returns the device name associated with a particular label.
	GetDeviceForLabel(label string) (string, error)

	// Returns all labels associated with a particular device name.
	GetLabelsForDevice(device string) ([]string, error)

	// Returns the mountpoint associated with a particular device.
	GetMountpointForDevice(device string) (string, error)
}
