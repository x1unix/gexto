package main

import (
	"os"
	"log"
	"github.com/lunixbochs/struc"
	"fmt"
	"io"
	"bytes"
	"encoding/binary"
)

type GroupDescriptor struct {
	Block_bitmap_lo      int32 `struc:"int32,little"`
	Inode_bitmap_lo      int32 `struc:"int32,little"`
	Inode_table_lo       int32 `struc:"int32,little"`
	Free_blocks_count_lo int16 `struc:"int16,little"`
	Free_inodes_count_lo int16 `struc:"int16,little"`
	Used_dirs_count_lo   int16 `struc:"int16,little"`
	Flags                int16 `struc:"int16,little"`
	Exclude_bitmap_lo    int32 `struc:"int32,little"`
	Block_bitmap_csum_lo int16 `struc:"int16,little"`
	Inode_bitmap_csum_lo int16 `struc:"int16,little"`
	Itable_unused_lo     int16 `struc:"int16,little"`
	Checksum             int16 `struc:"int16,little"`
	Block_bitmap_hi      int32 `struc:"int32,little"`
	Inode_bitmap_hi      int32 `struc:"int32,little"`
	Inode_table_hi       int32 `struc:"int32,little"`
	Free_blocks_count_hi int16 `struc:"int16,little"`
	Free_inodes_count_hi int16 `struc:"int16,little"`
	Used_dirs_count_hi   int16 `struc:"int16,little"`
	Itable_unused_hi     int16 `struc:"int16,little"`
	Exclude_bitmap_hi    int32 `struc:"int32,little"`
	Block_bitmap_csum_hi int16 `struc:"int16,little"`
	Inode_bitmap_csum_hi int16 `struc:"int16,little"`
	Reserved             int32 `struc:"int32,little"`
};

type Inode struct {
	Mode         int16     `struc:"int16,little"`
	Uid          int16     `struc:"int16,little"`
	Size_lo      int32     `struc:"int32,little"`
	Atime        int32     `struc:"int32,little"`
	Ctime        int32     `struc:"int32,little"`
	Mtime        int32     `struc:"int32,little"`
	Dtime        int32     `struc:"int32,little"`
	Gid          int16     `struc:"int16,little"`
	Links_count  int16     `struc:"int16,little"`
	Blocks_lo    int32     `struc:"int32,little"`
	Flags        int32     `struc:"int32,little"`
	Osd1         int32     `struc:"int32,little"`
	BlockOrExtents []byte `struc:"[60]byte,little"`
	Generation   int32     `struc:"int32,little"`
	File_acl_lo  int32     `struc:"int32,little"`
	Size_high    int32     `struc:"int32,little"`
	Obso_faddr   int32     `struc:"int32,little"`
	Osd2         [12]byte  `struc:"[12]byte"`
	Extra_isize  int16     `struc:"int16,little"`
	Checksum_hi  int16     `struc:"int16,little"`
	Ctime_extra  int32     `struc:"int32,little"`
	Mtime_extra  int32     `struc:"int32,little"`
	Atime_extra  int32     `struc:"int32,little"`
	Crtime       int32     `struc:"int32,little"`
	Crtime_extra int32     `struc:"int32,little"`
	Version_hi   int32     `struc:"int32,little"`
	Projid       int32     `struc:"int32,little"`
};

type MoveExtent struct {
	Reserved    int32  `struc:"int32,little"`
	Donor_fd    int32  `struc:"int32,little"`
	Orig_start  uint64 `struc:"uint64,little"`
	Donor_start uint64 `struc:"uint64,little"`
	Len         uint64 `struc:"uint64,little"`
	Moved_len   uint64 `struc:"uint64,little"`
};

type ExtentHeader struct {
	Magic      int16 `struc:"int16,little"`
	Entries    int16 `struc:"int16,little"`
	Max        int16 `struc:"int16,little"`
	Depth      int16 `struc:"int16,little"`
	Generation int32 `struc:"int32,little"`
}

type Extent struct {
	Block    int32 `struc:"int32,little"`
	Len      int16 `struc:"int16,little"`
	Start_hi int16 `struc:"int16,little"`
	Start_lo int32 `struc:"int32,little"`
}

type DirectoryEntry2 struct {
	Inode int32 `struc:"int32,little"`
	Rec_len int16 `struc:"int16,little"`
	Name_len int8 `struc:"int8,sizeof=Name"`
	Flags int8 `struc:"int8"`
	Name string `struc:"[]byte"`
}

type Superblock struct {
	InodeCount         int32 `struc:"int32,little"`
	BlockCount_lo      int32 `struc:"int32,little"`
	R_blockCount_lo    int32 `struc:"int32,little"`
	Free_blockCount_lo int32 `struc:"int32,little"`
	Free_inodeCount    int32 `struc:"int32,little"`
	First_data_block   int32 `struc:"int32,little"`
	Log_block_size     int32 `struc:"int32,little"`
	Log_cluster_size   int32 `struc:"int32,little"`
	BlockPer_group     int32 `struc:"int32,little"`
	ClusterPer_group   int32 `struc:"int32,little"`
	InodePer_group     int32 `struc:"int32,little"`
	Mtime              int32 `struc:"int32,little"`
	Wtime              int32 `struc:"int32,little"`
	Mnt_count          int16 `struc:"int16,little"`
	Max_mnt_count      int16 `struc:"int16,little"`
	Magic              int16 `struc:"int16,little"`
	State              int16 `struc:"int16,little"`
	Errors             int16 `struc:"int16,little"`
	Minor_rev_level    int16 `struc:"int16,little"`
	Lastcheck          int32 `struc:"int32,little"`
	Checkinterval      int32 `struc:"int32,little"`
	Creator_os         int32 `struc:"int32,little"`
	Rev_level          int32 `struc:"int32,little"`
	Def_resuid         int16 `struc:"int16,little"`
	Def_resgid         int16 `struc:"int16,little"`
	// Dynamic_rev superblocks only
	First_ino              int32    `struc:"int32,little"`
	Inode_size             int16    `struc:"int16,little"`
	Block_group_nr         int16    `struc:"int16,little"`
	Feature_compat         int32    `struc:"int32,little"`
	Feature_incompat       int32    `struc:"int32,little"`
	Feature_ro_compat      int32    `struc:"int32,little"`
	Uuid                   [16]byte `struc:"byte"`
	Volume_name            [16]byte `struc:"byte"`
	Last_mounted           [64]byte `struc:"byte"`
	Algorithm_usage_bitmap int32    `struc:"int32,little"`
	// Performance hints
	Prealloc_blocks     byte  `struc:"byte"`
	Prealloc_dir_blocks byte  `struc:"byte"`
	Reserved_gdt_blocks int16 `struc:"int16,little"`
	// Journal

	Journal_Uuid       [16]byte  `struc:"byte"`
	Journal_inum       int32     `struc:"int32,little"`
	Journal_dev        int32     `struc:"int32,little"`
	Last_orphan        int32     `struc:"int32,little"`
	Hash_seed          [4]int32  `struc:"[4]int32,little"`
	Def_hash_version   byte      `struc:"byte"`
	Jnl_backup_type    byte      `struc:"byte"`
	Desc_size          int16     `struc:"int16,little"`
	Default_mount_opts int32     `struc:"int32,little"`
	First_meta_bg      int32     `struc:"int32,little"`
	MkfTime            int32     `struc:"int32,little"`
	Jnl_blocks         [17]int32 `struc:"[17]int32,little"`

	BlockCount_hi         int32     `struc:"int32,little"`
	R_blockCount_hi       int32     `struc:"int32,little"`
	Free_blockCount_hi    int32     `struc:"int32,little"`
	Min_extra_isize       int16     `struc:"int16,little"`
	Want_extra_isize      int16     `struc:"int16,little"`
	Flags                 int32     `struc:"int32,little"`
	Raid_stride           int16     `struc:"int16,little"`
	Mmp_update_interval   int16     `struc:"int16,little"`
	Mmp_block             int64     `struc:"int64,little"`
	Raid_stripe_width     int32     `struc:"int32,little"`
	Log_groupPer_flex     byte      `struc:"byte"`
	Checksum_type         byte      `struc:"byte"`
	Encryption_level      byte      `struc:"byte"`
	Reserved_pad          byte      `struc:"byte"`
	KbyteWritten          int64     `struc:"int64,little"`
	Snapshot_inum         int32     `struc:"int32,little"`
	Snapshot_id           int32     `struc:"int32,little"`
	Snapshot_r_blockCount int64     `struc:"int64,little"`
	Snapshot_list         int32     `struc:"int32,little"`
	Error_count           int32     `struc:"int32,little"`
	First_error_time      int32     `struc:"int32,little"`
	First_error_ino       int32     `struc:"int32,little"`
	First_error_block     int64     `struc:"int64,little"`
	First_error_func      [32]byte  `struc:"pad"`
	First_error_line      int32     `struc:"int32,little"`
	Last_error_time       int32     `struc:"int32,little"`
	Last_error_ino        int32     `struc:"int32,little"`
	Last_error_line       int32     `struc:"int32,little"`
	Last_error_block      int64     `struc:"int64,little"`
	Last_error_func       [32]byte  `struc:"pad"`
	Mount_opts            [64]byte  `struc:"pad"`
	Usr_quota_inum        int32     `struc:"int32,little"`
	Grp_quota_inum        int32     `struc:"int32,little"`
	Overhead_clusters     int32     `struc:"int32,little"`
	Backup_bgs            [2]int32  `struc:"[2]int32,little"`
	Encrypt_algos         [4]byte   `struc:"pad"`
	Encrypt_pw_salt       [16]byte  `struc:"pad"`
	Lpf_ino               int32     `struc:"int32,little"`
	Prj_quota_inum        int32     `struc:"int32,little"`
	Checksum_seed         int32     `struc:"int32,little"`
	Reserved              [98]int32 `struc:"[98]int32,little"`
	Checksum              int32     `struc:"int32,little"`
};

func (sb *Superblock) FeatureCompatDir_prealloc() bool  { return (sb.Feature_compat&FEATURE_COMPAT_DIR_PREALLOC != 0) }
func (sb *Superblock) FeatureCompatImagic_inodes() bool { return (sb.Feature_compat&FEATURE_COMPAT_IMAGIC_INODES != 0) }
func (sb *Superblock) FeatureCompatHas_journal() bool   { return (sb.Feature_compat&FEATURE_COMPAT_HAS_JOURNAL != 0) }
func (sb *Superblock) FeatureCompatExt_attr() bool      { return (sb.Feature_compat&FEATURE_COMPAT_EXT_ATTR != 0) }
func (sb *Superblock) FeatureCompatResize_inode() bool  { return (sb.Feature_compat&FEATURE_COMPAT_RESIZE_INODE != 0) }
func (sb *Superblock) FeatureCompatDir_index() bool     { return (sb.Feature_compat&FEATURE_COMPAT_DIR_INDEX != 0) }
func (sb *Superblock) FeatureCompatSparse_super2() bool { return (sb.Feature_compat&FEATURE_COMPAT_SPARSE_SUPER2 != 0) }

func (sb *Superblock) FeatureRoCompatSparse_super() bool  { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_SPARSE_SUPER != 0) }
func (sb *Superblock) FeatureRoCompatLarge_file() bool    { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_LARGE_FILE != 0) }
func (sb *Superblock) FeatureRoCompatBtree_dir() bool     { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_BTREE_DIR != 0) }
func (sb *Superblock) FeatureRoCompatHuge_file() bool     { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_HUGE_FILE != 0) }
func (sb *Superblock) FeatureRoCompatGdt_csum() bool      { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_GDT_CSUM != 0) }
func (sb *Superblock) FeatureRoCompatDir_nlink() bool     { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_DIR_NLINK != 0) }
func (sb *Superblock) FeatureRoCompatExtra_isize() bool   { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_EXTRA_ISIZE != 0) }
func (sb *Superblock) FeatureRoCompatQuota() bool         { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_QUOTA != 0) }
func (sb *Superblock) FeatureRoCompatBigalloc() bool      { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_BIGALLOC != 0) }
func (sb *Superblock) FeatureRoCompatMetadata_csum() bool { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_METADATA_CSUM != 0) }
func (sb *Superblock) FeatureRoCompatReadonly() bool      { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_READONLY != 0) }
func (sb *Superblock) FeatureRoCompatProject() bool       { return (sb.Feature_ro_compat&FEATURE_RO_COMPAT_PROJECT != 0) }

func (sb *Superblock) FeatureIncompat64bit() bool       { return (sb.Feature_incompat&FEATURE_INCOMPAT_64BIT != 0) }
func (sb *Superblock) FeatureIncompatCompression() bool { return (sb.Feature_incompat&FEATURE_INCOMPAT_COMPRESSION != 0) }
func (sb *Superblock) FeatureIncompatFiletype() bool    { return (sb.Feature_incompat&FEATURE_INCOMPAT_FILETYPE != 0) }
func (sb *Superblock) FeatureIncompatRecover() bool     { return (sb.Feature_incompat&FEATURE_INCOMPAT_RECOVER != 0) }
func (sb *Superblock) FeatureIncompatJournal_dev() bool { return (sb.Feature_incompat&FEATURE_INCOMPAT_JOURNAL_DEV != 0) }
func (sb *Superblock) FeatureIncompatMeta_bg() bool     { return (sb.Feature_incompat&FEATURE_INCOMPAT_META_BG != 0) }
func (sb *Superblock) FeatureIncompatExtents() bool     { return (sb.Feature_incompat&FEATURE_INCOMPAT_EXTENTS != 0) }
func (sb *Superblock) FeatureIncompatMmp() bool         { return (sb.Feature_incompat&FEATURE_INCOMPAT_MMP != 0) }
func (sb *Superblock) FeatureIncompatFlex_bg() bool     { return (sb.Feature_incompat&FEATURE_INCOMPAT_FLEX_BG != 0) }
func (sb *Superblock) FeatureIncompatEa_inode() bool    { return (sb.Feature_incompat&FEATURE_INCOMPAT_EA_INODE != 0) }
func (sb *Superblock) FeatureIncompatDirdata() bool     { return (sb.Feature_incompat&FEATURE_INCOMPAT_DIRDATA != 0) }
func (sb *Superblock) FeatureIncompatCsum_seed() bool   { return (sb.Feature_incompat&FEATURE_INCOMPAT_CSUM_SEED != 0) }
func (sb *Superblock) FeatureIncompatLargedir() bool    { return (sb.Feature_incompat&FEATURE_INCOMPAT_LARGEDIR != 0) }
func (sb *Superblock) FeatureIncompatInline_data() bool { return (sb.Feature_incompat&FEATURE_INCOMPAT_INLINE_DATA != 0) }
func (sb *Superblock) FeatureIncompatEncrypt() bool     { return (sb.Feature_incompat&FEATURE_INCOMPAT_ENCRYPT != 0) }

func (sb *Superblock) GetBlockCount() int64 {
	if sb.FeatureIncompat64bit() {
		return (int64(sb.BlockCount_hi) << 32) | int64(sb.BlockCount_lo)
	} else {
		return int64(sb.BlockCount_lo)
	}
}

func (sb *Superblock) GetBlockSize() int64 {
	return int64(1024 << uint(sb.Log_block_size))
}

func getBlockGroupDescriptor(blockGroupNum int64,  sb *Superblock, dev *os.File) *GroupDescriptor {
	blockSize := sb.GetBlockSize()
	bgdtLocation := 1024/blockSize + 1

	bgd := &GroupDescriptor{}
	dev.Seek((bgdtLocation+blockGroupNum)*blockSize, 0)
	if sb.FeatureIncompat64bit() {
		struc.Unpack(dev, &bgd)
	} else {
		struc.Unpack(io.LimitReader(dev, 32), &bgd)
	}
	//fmt.Printf("Read block group %d, contents:\n%+v\n", blockGroupNum, bgd)
	return bgd
}

func (bgd *GroupDescriptor) GetInodeTableLoc(sb *Superblock) int64 {
	if sb.FeatureIncompat64bit() {
		return (int64(bgd.Inode_table_hi) << 32) | int64(bgd.Inode_table_lo)
	} else {
		return int64(bgd.Inode_table_lo)
	}
}

func getInode(inodeAddress int64, sb *Superblock, dev *os.File) *Inode {
	bgd := getBlockGroupDescriptor((inodeAddress - 1) / int64(sb.InodePer_group), sb, dev)
	index := (inodeAddress - 1) % int64(sb.InodePer_group)
	pos := bgd.GetInodeTableLoc(sb) * sb.GetBlockSize() + index * int64(sb.Inode_size)
	//log.Printf("%d %d %d %d", bgd.GetInodeTableLoc(sb), sb.GetBlockSize(), index, sb.Inode_size)
	dev.Seek(pos, 0)

	inode := &Inode{}
	struc.Unpack(dev, &inode)
	//log.Printf("Read inode at offset %d, contents:\n%+v\n", pos, inode)
	return inode
}

func (inode *Inode) UsesExtents() bool {
	return (inode.Flags & EXTENTS_FL) != 0
}

func (inode *Inode) UsesDirectoryHashTree() bool {
	return (inode.Flags & INDEX_FL) != 0
}

func (inode *Inode) ReadFile(sb *Superblock, dev *os.File) {
	size := int64(inode.Size_lo)
	for blockTableIndex := int64(0); blockTableIndex < (int64(inode.Size_lo)+sb.GetBlockSize()-1)/sb.GetBlockSize(); blockTableIndex++ {
		blockNum := binary.LittleEndian.Uint32(inode.BlockOrExtents[blockTableIndex * 4:])
		dev.Seek(int64(blockNum) * sb.GetBlockSize(), 0)
		sizeInBlock := sb.GetBlockSize()
		if size < sizeInBlock {
			sizeInBlock = size
		}
		data := make([]byte, sizeInBlock)
		dev.Read(data)
		log.Printf("%s", string(data))
		size -= sizeInBlock
	}

	if size > 0 {
		log.Fatalf("Oversize block")
	}
}

func (inode *Inode) ReadDirectory(sb *Superblock, dev *os.File) []DirectoryEntry2 {
	if inode.UsesDirectoryHashTree() {
		log.Fatalf("Not implemented")
	}
	if inode.UsesExtents() {
		extentHeader := &ExtentHeader{}
		struc.Unpack(bytes.NewReader([]byte(inode.BlockOrExtents)), &extentHeader)
		log.Printf("extent header: %+v", extentHeader)
		if extentHeader.Depth == 0 { // Leaf
			for i := int16(0); i < extentHeader.Entries; i++ {
				extent := &Extent{}
				struc.Unpack(bytes.NewReader([]byte(inode.BlockOrExtents)[12 + i * 12:]), &extent)
				log.Printf("extent: %+v", extent)
			}
		} else {
			log.Fatalf("Not implemented")
		}
		return nil
	} else {
		ret := []DirectoryEntry2{}
		for blockTableIndex := int64(0); blockTableIndex < (int64(inode.Size_lo) + sb.GetBlockSize() - 1) / sb.GetBlockSize(); blockTableIndex++ {
			blockNum := binary.LittleEndian.Uint32(inode.BlockOrExtents[blockTableIndex * 4:])
			blockStart := int64(blockNum) * sb.GetBlockSize()
			pos := blockStart
			for i := 0; i < 16; i++ {
				dev.Seek(pos, 0)
				dirEntry := DirectoryEntry2{}
				struc.Unpack(dev, &dirEntry)
				log.Printf("dirEntry %s: %+v", string(dirEntry.Name), dirEntry)
				pos += int64(dirEntry.Rec_len)
				ret = append(ret, dirEntry)
				if pos == blockStart+sb.GetBlockSize() {
					log.Printf("Reached end of block, next block")
					break
				} else if pos > blockStart + sb.GetBlockSize() {
					log.Fatalf("Unexpected overflow out of block when directory listing")
				}
			}
		}
		return ret
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}
	f.Seek(1024, 0)

	sb := &Superblock{}
	err = struc.Unpack(f, &sb)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Super:\n%+v\n", sb)

	numBlockGroups := (sb.GetBlockCount() + int64(sb.BlockPer_group) - 1) / int64(sb.BlockPer_group)
	numBlockGroups2 := (sb.InodeCount + sb.InodePer_group - 1) / sb.InodePer_group
	if numBlockGroups != int64(numBlockGroups2) {
		log.Fatalf("Block/inode mismatch: %d %d %d", sb.GetBlockCount(), numBlockGroups, numBlockGroups2)
	}

	inode := getInode(ROOT_INO, sb, f)
	dirContents := inode.ReadDirectory(sb, f)
	for i := 0; i < len(dirContents); i++ {
		log.Println(string(dirContents[i].Name), dirContents[i].Flags)
		if dirContents[i].Flags == 1 {
			inode := getInode(int64(dirContents[i].Inode), sb, f)
			inode.ReadFile(sb, f)
		}
	}
}
