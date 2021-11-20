package store

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/s3-management/message/npool"
	testinit "github.com/NpoolPlatform/s3-management/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestS3(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()
	imgType := "test"
	imgBase64 := "iVBORw0KGgoAAAANSUhEUgAAAB4AAAAZCAYAAAAmNZ4aAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAySURBVEhL7c2hAQAgDMTAp/vv3CI6QxDkTGROX3mgtjjHGMcYxxjHGMcYxxjHmN/GyQBA0AQuiLmS2gAAAABJRU5ErkJggg=="
	imgID := imgType + userID

	resp, err := UploadKycImg(context.Background(), &npool.UploadKycImgRequest{
		UserID:    userID,
		ImgType:   imgType,
		ImgBase64: imgBase64,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
		assert.Equal(t, "kyc/"+imgID, resp.Info)
	}

	resp1, err := GetKycImg(context.Background(), &npool.GetKycImgRequest{
		ImgID: resp.Info,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
		assert.Equal(t, imgBase64, resp1.Info)
	}
}
