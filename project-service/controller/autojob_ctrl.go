package controller

import (
	"bx.com/project-service/models"
	"bx.com/project-service/proto"
	"context"
)
type AutoJobController struct{}

//自动更新项目状态开始或结束
func (ctrl *AutoJobController) UpdateProjectsStatus(ctx context.Context, in *proto.Empty) (*proto.Empty, error){
	err := models.UpdateProjectsStatusInfo()

	return &proto.Empty{}, err
}

//自动更新项目阶段状态开始或结束
func (ctrl *AutoJobController) UpdateStagesStatus(ctx context.Context, in *proto.Empty) (*proto.Empty, error){
	err := models.UpdateStageStatusInfo()

	return &proto.Empty{}, err
}

//自动更新订单为未付款过期
func (ctrl *AutoJobController) UpdateOtcDetailsStatus(ctx context.Context, in *proto.Empty) (*proto.Empty, error){
	err := models.UpdateOrderDetailsStatus()

	return &proto.Empty{}, err
}