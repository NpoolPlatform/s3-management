package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/s3-management/message/npool"
	mystore "github.com/NpoolPlatform/s3-management/pkg/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UploadImgToS3(ctx context.Context, in *npool.UploadKycImgRequest) (*npool.UploadKycImgResponse, error) {
	resp, err := mystore.UploadKycImg(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to upload img to s3: %v", err)
		return &npool.UploadKycImgResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetImgFromS3(ctx context.Context, in *npool.GetKycImgRequest) (*npool.GetKycImgResponse, error) {
	resp, err := mystore.GetKycImg(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get img from s3: %v", err)
		return &npool.GetKycImgResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
