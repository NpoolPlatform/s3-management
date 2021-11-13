package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/s3-management/message/npool"
	"github.com/NpoolPlatform/s3-management/pkg/s3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UploadImgToS3(ctx context.Context, in *npool.UploadImgToS3Request) (*npool.UploadImgToS3Response, error) {
	resp, err := s3.UploadImgToS3(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to upload img to s3: %v", err)
		return &npool.UploadImgToS3Response{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetImgFromS3(ctx context.Context, in *npool.GetImgFromS3Request) (*npool.GetImgFromS3Response, error) {
	resp, err := s3.GetImgFromS3(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get img from s3: %v", err)
		return &npool.GetImgFromS3Response{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
