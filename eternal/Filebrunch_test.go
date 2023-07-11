package fileserver

import (
	"fmt"
	"testing"
)

func TestFileBrunch_GetTotalSize(t *testing.T) {

	tests := []struct {
		name string
		fb   *FileBrunch
		want int64
	}{
		{"1", GetTotalSizeGetfb(), 600},
		{"2", GetTotalSizeGetfb0(), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.GetTotalSize(); got != tt.want {
				t.Errorf("FileBrunch.GetTotalSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func GetTotalSizeGetfb() *FileBrunch {
	fb1 := NewFileBrunch([]byte("name1"), nil)

	for i := 1; i <= 3; i++ {
		cb := NewFileBrunch([]byte(fmt.Sprint("file:", i)), fb1)
		cb.Curentfilesize = int64(i * 100)
		fb1.ChildBrunchs = append(fb1.ChildBrunchs, cb)
	}
	return fb1
}

func GetTotalSizeGetfb0() *FileBrunch {
	fb1 := NewFileBrunch([]byte("name1"), nil)
	return fb1
}

func TestFileBrunch_GetDepthAndSize(t *testing.T) {
	x := NewFileBrunch([]byte("name1"), nil)
	x.Curentfilesize = 16384
	size, debth := x.GetDepthAndSize(x.Curentfilesize)
	if size != 16 {
		t.Errorf("FileBrunch.GetDepthAndSize() got = %v, want %v", size, 16)
	}
	if debth != "kbyte" {
		t.Errorf("FileBrunch.GetDepthAndSize() got = %v, want %v", debth, "kbyte")
	}

	x = NewFileBrunch([]byte("name1"), nil)
	x.Curentfilesize = 8
	size, debth = x.GetDepthAndSize(x.Curentfilesize)
	if size != 8 {
		t.Errorf("FileBrunch.GetDepthAndSize() got = %v, want %v", size, 8)
	}
	if debth != "byte" {
		t.Errorf("FileBrunch.GetDepthAndSize() got = %v, want %v", debth, "byte")
	}

	x = NewFileBrunch([]byte("name1"), nil)
	x.Curentfilesize = 25351718795
	size, debth = x.GetDepthAndSize(x.Curentfilesize)
	if size != 23.610628019087017 {
		t.Errorf("FileBrunch.GetDepthAndSize() got = %v, want %v", size, 23.610628019087017)
	}
	if debth != "gbyte" {
		t.Errorf("FileBrunch.GetDepthAndSize() got = %v, want %v", debth, "gbyte")
	}

}
