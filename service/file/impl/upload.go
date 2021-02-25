package impl

import (
	"context"
	"log"
	filepb "shakebook/service/file/proto/api/v1"

	"google.golang.org/grpc/codes"
)

//Upload 文件上传
func Upload(c context.Context, req *filepb.UploadRequest) (*filepb.Response, error) {
	log.Println("文件上传.....")
	log.Printf("req file:%v", req)
	return &filepb.Response{
		Code:    int32(codes.OK),
		Message: "上传成功",
	}, nil
}
