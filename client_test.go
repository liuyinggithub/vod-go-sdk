package vod

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"os"
	"strings"
	"testing"
)

const region = "ap-guangzhou"

func getClient() *VodUploadClient {
	client := &VodUploadClient{}
	client.SecretId = os.Getenv("SECRET_ID")
	client.SecretKey = os.Getenv("SECRET_KEY")
	return client
}

func TestLackMediaPath(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	_, err := client.Upload(region, req)
	if err.Error() != "[VodClientError] Message=lack media path" {
		t.Error(err.Error())
	}
}

func TestLackMediaType(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife")
	_, err := client.Upload(region, req)
	if err.Error() != "[VodClientError] Message=lack media type" {
		t.Error(err.Error())
	}
}

func TestInvalidMediaPath(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/WildlifeA")
	_, err := client.Upload(region, req)
	if err.Error() != "[VodClientError] Message=media path is invalid" {
		t.Error(err.Error())
	}
}

func TestInvalidCoverPath(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
	req.CoverFilePath = common.StringPtr("video/Wildlife-CoverA")
	_, err := client.Upload(region, req)
	if err.Error() != "[VodClientError] Message=cover path is invalid" {
		t.Error(err.Error())
	}
}

func TestLackCoverType(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
	req.CoverFilePath = common.StringPtr("video/Wildlife-cover")
	_, err := client.Upload(region, req)
	if err.Error() != "[VodClientError] Message=lack cover type" {
		t.Error(err.Error())
	}
}

func TestInvalidMediaType(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
	req.MediaType = common.StringPtr("test")
	_, err := client.Upload(region, req)
	if !strings.HasPrefix(err.Error(), "[TencentCloudSDKError] Code=InvalidParameterValue.VideoType, Message=invalid video type") {
		t.Error(err.Error())
	}
}

func TestInvalidCoverType(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
	req.CoverFilePath = common.StringPtr("video/Wildlife-cover.png")
	req.CoverType = common.StringPtr("test")
	_, err := client.Upload(region, req)
	if !strings.HasPrefix(err.Error(), "[TencentCloudSDKError] Code=InvalidParameterValue.CoverType, Message=invalid cover type") {
		t.Error(err.Error())
	}
}

func TestUploadNormal(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
	req.CoverFilePath = common.StringPtr("video/Wildlife-cover.png")
	rsp, err := client.Upload(region, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(*rsp.Response.FileId)
	t.Log(*rsp.Response.MediaUrl)
	t.Log(*rsp.Response.CoverUrl)
}

func TestUploadWithProcedure(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
	req.CoverFilePath = common.StringPtr("video/Wildlife-cover.png")
	req.Procedure = common.StringPtr("QCVB_SimpleProcessFile(1, 1)")
	rsp, err := client.Upload(region, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(*rsp.Response.FileId)
	t.Log(*rsp.Response.MediaUrl)
	t.Log(*rsp.Response.CoverUrl)
}

func TestUploadWithSubAppId(t *testing.T) {
	client := getClient()
	req := NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
	req.CoverFilePath = common.StringPtr("video/Wildlife-cover.png")
	req.SubAppId = common.Uint64Ptr(1400001888)
	rsp, err := client.Upload(region, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(*rsp.Response.FileId)
	t.Log(*rsp.Response.MediaUrl)
	t.Log(*rsp.Response.CoverUrl)
}
