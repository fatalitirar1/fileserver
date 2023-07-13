package fileserver

import "fmt"

type FileBrunch struct {
	ParentBrunch   *FileBrunch
	ChildBrunchs   []*FileBrunch
	Curentfilesize int64
	Path           []byte
}

func NewFileBrunch(Path []byte, ParentBrunch *FileBrunch) *FileBrunch {
	nfb := &FileBrunch{}
	nfb.Path = Path
	if ParentBrunch != nil {
		nfb.ParentBrunch = ParentBrunch
	}
	return nfb
}

func (fb FileBrunch) GetTotalSize() int64 {
	var size int64
	for _, brunch := range fb.ChildBrunchs {
		size += brunch.Curentfilesize
	}
	return size
}

func (fb FileBrunch) String() string {
	size, debth := fb.GetDepthAndSize(fb.Curentfilesize)
	return fmt.Sprint(string(fb.Path), ": ", size, " ", debth, "Real size: ", fb.Curentfilesize)
}

func (fb FileBrunch) GetDepthAndSize(size int64) (float64, string) {
	depth := 0
	var depthR string
	rSize := float64(size)
	for rSize >= 1024 {
		rSize /= 1024
		depth++
	}
	depthR = fb.GetDepthString(depth)

	return rSize, depthR
}

func (fb FileBrunch) GetDepthString(n int) string {
	return [...]string{"byte", "kbyte", "mbyte", "gbyte", "tbyte"}[n]
}

func (fb *FileBrunch) SetSize(size int64) {
	fb.Curentfilesize = size
}

func (fb *FileBrunch) ListFullBrunch(indent string) {
	for _, brunch := range fb.ChildBrunchs {
		if len(brunch.ChildBrunchs) > 0 {
			fmt.Println(brunch)
			brunch.ListFullBrunch(indent + "--")
		} else {
			fmt.Println(indent, brunch)
		}
	}
}
