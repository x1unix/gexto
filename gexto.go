package gexto

import (
	"fmt"
	"io"
	"os"

	"github.com/lunixbochs/struc"
)

// Superblock0Offset is ext superblock offset.
const Superblock0Offset = 1024

type File struct {
	extFile
}

type FileSystem interface {
	Open(name string) (*File, error)
	Create(name string) (*File, error)
	Remove(name string) error
	Mkdir(name string, perm os.FileMode) error
	Close() error
}

func NewFileSystem(f io.ReadSeeker) (FileSystem, error) {
	ret := fs{}

	if _, err := f.Seek(Superblock0Offset, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek at superblock: %w", err)
	}

	ret.dev = f
	ret.sb = &Superblock{
		address: Superblock0Offset,
		fs:      &ret,
	}

	if err := struc.Unpack(f, ret.sb); err != nil {
		return nil, err
	}

	numBlockGroups := (ret.sb.GetBlockCount() + int64(ret.sb.BlockPer_group) - 1) / int64(ret.sb.BlockPer_group)
	numBlockGroups2 := (ret.sb.InodeCount + ret.sb.InodePer_group - 1) / ret.sb.InodePer_group
	if numBlockGroups != int64(numBlockGroups2) {
		return nil, fmt.Errorf("block/inode mismatch: %d %d %d", ret.sb.GetBlockCount(), numBlockGroups, numBlockGroups2)
	}

	ret.sb.numBlockGroups = numBlockGroups

	return &ret, nil
}
