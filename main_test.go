package main

import (
	fileserver "Fileserver/eternal"
	"context"
	"fmt"
	"testing"
)

func Test_listDirByReadDir(t *testing.T) {
	type args struct {
		ctx  context.Context
		path string
		pb   *fileserver.FileBrunch
	}
	tests := []struct {
		name string
		args args
		wont int64
	}{
		{"DLL", args{context.Background(), "D:/AnyDLL", fileserver.NewFileBrunch([]byte("D:/AnyDLL"), nil)}, 849360},
		{"Санитары Подземелий", args{context.Background(), "D:/Games/Санитары Подземелий", fileserver.NewFileBrunch([]byte("D:/Games/Санитары Подземелий"), nil)}, 2526520466},
		{"testPackage", args{context.Background(), "D:/testPackage", fileserver.NewFileBrunch([]byte("D:/testPackage"), nil)}, 2520000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listDirByReadDir(tt.args.ctx, tt.args.path, tt.args.pb)
			if tt.args.pb.Curentfilesize != tt.wont {
				t.Errorf("listDirByReadDir = %v, want %v want", tt.args.pb.Curentfilesize, tt.wont)
				tt.args.pb.ListFullBrunch("")
			} else {
				fmt.Printf("ok: size of %s : %v \n", tt.name, tt.args.pb.Curentfilesize)
			}
		})
	}
}
