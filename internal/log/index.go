package log

import (
	"io"
	"os"
	"github.com/tysonmote/gommap"
)
var (
	offWidth uint64 = 4
	posWidth uint64 = 8
	endWidth = offWidth + posWidth
)
type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}


func newIndx(f *os.File, c Config) (*index, error) {
	idx := &index {
		file: f,
	}
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fi.Size())
	if err = os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes),); err != nil{
			return nil, err
		}
	if idx.mmap, err = gommap.Map(idx.file.Fd(), gommap.PROT_READ|gommap.PRO_WRITE, gommap.MAP_SHARED); err != nil {
		return nil, err
	}
	return idx, nil

}

func (i *index) Close() error {
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	if err := i.file.Sync(); err != nil {
		return err
	}
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}
	return i.file.Close()
}

func (i *index) Write(off uint32, pos uint64) error{
	if uint64(len(i.map)) < i.size+endWidth {
		return io.EOF
	}
	enc.PutUint32(i.mmap[i.size:i.size+offWidth],pos)
	i.size+= uint64(endWidth)
	return nil
}
func (i *index) Name() string {
	return i.file.Name()
}