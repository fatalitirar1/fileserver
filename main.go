package main

import (
	fileserver "Fileserver/eternal"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	fmt.Println("List by ReadDir")
	var path string
	ctx, err := context.WithTimeout(context.Background(), time.Minute)
	if err != nil {
		log.Print(err)
	}

	fmt.Scan(&path)
	pb := fileserver.NewFileBrunch([]byte(path), nil)
	listDirByReadDir(ctx, path, pb)
	pb.ListFullBrunch("")
	//for _, branch := range pb.ChildBrunchs {
	//	fmt.Println("--", branch)
	//}

}

func listDirByReadDir(ctx context.Context, path string, pb *fileserver.FileBrunch) {
	lst, err := ioutil.ReadDir(path)
	var size int64

	if err != nil {
		panic(err)
	}

	for _, val := range lst {
		nb := fileserver.NewFileBrunch([]byte(path+"/"+val.Name()), pb)
		if val.IsDir() {
			listDirByReadDir(ctx, path+"/"+val.Name(), nb)
		} else {
			nb.Curentfilesize += val.Size()

		}
		pb.ChildBrunchs = append(pb.ChildBrunchs, nb)
		size += val.Size()
	}
	pb.Curentfilesize = pb.GetTotalSize()
}
