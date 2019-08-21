package controller

import (
	"bx.com/project-service/models"
	"context"
	"fmt"
	"bx.com/project-service/proto"
	"bx.com/user-service/utils"
	"time"
	"errors"
)

type ProjectController struct{}

//Project API - Base


func (ctrl *ProjectController) QueryProjectById(ctx context.Context, in *proto.IdReq) (*proto.ProjectInfoReply, error) {
	projectInfo, err := models.QueryProjectById(in.Id)
	if err != nil || projectInfo.AdminId == 0 {
		return &proto.ProjectInfoReply{}, err
	}else {
		return &proto.ProjectInfoReply{
			Id:					projectInfo.Id,
			AdminId:			projectInfo.AdminId,
			Summary:			projectInfo.Summary,
			TargetValue:		projectInfo.TargetValue,
			IssueCoina:			projectInfo.IssueCoina,
			IssueCoinb:			projectInfo.IssueCoinb,
			StageCount:			int32(projectInfo.StageCount),
			Classify:			projectInfo.Classify,
			WhitePaper:			projectInfo.WhitePaper,
			OfficeSite:			projectInfo.OfficialSite,
			CommunityAddress:	projectInfo.CommunityAddress,
			Status:				proto.Status(projectInfo.Status),
			PrioritySort:		int32(projectInfo.PrioritySort),
			CreateTime:			utils.Time2String(projectInfo.CreateTime),
			BeginTime:			utils.Time2String(projectInfo.BeginTime),
			EndTime:			utils.Time2String(projectInfo.EndTime),
			UpdateTime:			utils.Time2String(projectInfo.UpdateTime),
		}, nil
	}
}

func (ctrl *ProjectController) QueryUserParticipationProject(ctx context.Context, in *proto.IdReq) (*proto.ProjectListReply, error) {
	projectList, err := models.QueryProjectByUser(in.Id)
	if err != nil || len(projectList) == 0 {
		return &proto.ProjectListReply{}, err
	}else{
		projectResults := proto.ProjectListReply{}
		for _, projectInfo := range projectList {
			projectResults.ProjectList = append(projectResults.ProjectList, &proto.ProjectInfoReply{
				Id:					projectInfo.Id,
				AdminId:			projectInfo.AdminId,
				Summary:			projectInfo.Summary,
				TargetValue:		projectInfo.TargetValue,
				IssueCoina:			projectInfo.IssueCoina,
				IssueCoinb:			projectInfo.IssueCoinb,
				StageCount:			int32(projectInfo.StageCount),
				Classify:			projectInfo.Classify,
				WhitePaper:			projectInfo.WhitePaper,
				OfficeSite:			projectInfo.OfficialSite,
				CommunityAddress:	projectInfo.CommunityAddress,
				Status:				proto.Status(projectInfo.Status),
				PrioritySort:		int32(projectInfo.PrioritySort),
				CreateTime:			utils.Time2String(projectInfo.CreateTime),
				BeginTime:			utils.Time2String(projectInfo.BeginTime),
				EndTime:			utils.Time2String(projectInfo.EndTime),
				UpdateTime:			utils.Time2String(projectInfo.UpdateTime),
			})
		}

		return &projectResults, nil
	}
}

func (ctrl *ProjectController) QueryProjectPhotosInfo(ctx context.Context, in *proto.IdReq) (*proto.MediaListReply, error) {
	photoList, err := models.QueryProjectMediaList(in.Id, models.ImageType)
	if err != nil || len(photoList) == 0 {
		return &proto.MediaListReply{}, err
	}else{
		photoResults := proto.MediaListReply{}
		for _, media := range photoList {
			photoResults.MediaList = append(photoResults.MediaList, &proto.MediaInfoReply{
				ProjectId:		media.ProjectId,
				Title:			media.Title,
				Address:		media.Address,
				Type:			proto.MediaType(media.Type),
				Enable:			proto.BoolValue(media.Enable),
				PrioritySort:	int32(media.PrioritySort),
				CreateTime:		utils.Time2String(media.CreateTime),
			})
		}

		return &photoResults, nil
	}
}

func (ctrl *ProjectController) QueryProjectVideosInfo(ctx context.Context, in *proto.IdReq) (*proto.MediaListReply, error) {
	videoList, err := models.QueryProjectMediaList(in.Id, models.VideoType)
	if err != nil || len(videoList) == 0 {
		return &proto.MediaListReply{}, err
	}else{
		videoResults := proto.MediaListReply{}
		for _, media := range videoList {
			videoResults.MediaList = append(videoResults.MediaList, &proto.MediaInfoReply{
				ProjectId:		media.ProjectId,
				Title:			media.Title,
				Address:		media.Address,
				Type:			proto.MediaType(media.Type),
				Enable:			proto.BoolValue(media.Enable),
				PrioritySort:	int32(media.PrioritySort),
				CreateTime:		utils.Time2String(media.CreateTime),
			})
		}

		return &videoResults, nil
	}
}

func (ctrl *ProjectController) QueryProjectDescriptionInfo(ctx context.Context, in *proto.IdReq) (*proto.ProjectDescriptionListReply, error) {
	descriptionList, err := models.QueryProjectsTextByProject(in.Id)
	if err != nil || len(descriptionList) == 0 {
		return &proto.ProjectDescriptionListReply{}, err
	} else {
		descriptionResults := proto.ProjectDescriptionListReply{}
		for _, description := range descriptionList {
			descriptionResults.ProjectDescriptionList = append(descriptionResults.ProjectDescriptionList, &proto.ProjectDescriptionInfoReply{
				ProjectId:		description.ProjectId,
				Title:			description.Title,
				ProjectText:	description.ProjectText,
				Enable:			proto.BoolValue(description.Enable),
				PrioritySort:   int32(description.PrioritySort),
				CreateTime:     utils.Time2String(description.CreateTime),
				UpdateTime:		utils.Time2String(description.UpdateTime),
			})
		}

		return &descriptionResults, nil
	}
}

func (ctrl *ProjectController) QueryProjectCertificationInfo(ctx context.Context, in *proto.IdReq) (*proto.CertificationListReply, error) {
	certificationList, err := models.QueryProjectCertificationList(in.Id)
	if err != nil || len(certificationList) == 0 {
		return &proto.CertificationListReply{}, nil
	} else {
		certiResults := proto.CertificationListReply{}
		for _, certi := range certificationList {
			certiResults.CertificationList = append(certiResults.CertificationList, &proto.CertificationInfoReply{
				Id:				certi.Id,
				Name:			certi.Name,
				Description:	certi.Description,
				ImageUrl:		certi.Image,
				PrioritySort:	int32(certi.PrioritySort),
			})
		}

		return &certiResults, nil
	}
}

func (ctrl *ProjectController) QueryProjectMembersInfo(ctx context.Context, in *proto.IdReq) (*proto.MemberListReply, error) {
	membersList, err := models.QueryProjectMemberList(in.Id, 0)
	if err != nil || len(membersList) == 0 {
		return &proto.MemberListReply{}, err
	} else {
		memberResults := proto.MemberListReply{}
		for _, member := range membersList {
			memberResults.MemberList = append(memberResults.MemberList, &proto.MemeberReply{
				Id:					member.Id,
				ProjectId:			member.ProjectId,
				Name:				member.Name,
				Position:			member.Position,
				Description:		member.Description,
				MemberType:			proto.MemberType(member.MemberType),
				ImageUrl:			member.Image,
				PrioritySort:		int32(member.PrioritySort),
				JoinTime:			utils.Time2String(member.JoinTime),
			})
		}

		return &memberResults, nil
	}
}

func (ctrl *ProjectController) QueryProjectStagesInfo(ctx context.Context, in *proto.IdReq) (*proto.StageListReply, error) {
	stageList, err := models.QueryProjectStageList(in.Id)
	if err != nil || len(stageList) == 0 {
		return &proto.StageListReply{}, err
	} else {
		stageResults := proto.StageListReply{}
		for _, stage := range stageList {
			stageResults.StageList = append(stageResults.StageList, &proto.StageInfoReply{
				Id:				stage.Id,
				ProjectId:		stage.ProjectId,
				StageNumber:	int32(stage.StageNumber),
				StageName:		stage.StageName,
				CoinCount:		int32(stage.CoinCount),
				TargetValue:	stage.TargetValue,
				SoftValue:		stage.SoftValue,
				Discount:		stage.Discount,
				CompleteValue:	stage.CompleteValue,
				StageStatus:	proto.Status(stage.StageStatus),
				BeginTime:		utils.Time2String(stage.BeginTime),
				EndTime:		utils.Time2String(stage.EndTime),
				CreateTime:		utils.Time2String(stage.CreateTime),
				UpdateTime:		utils.Time2String(stage.UpdateTime),
			})
		}

		return &stageResults, nil
	}

}

func (ctrl *ProjectController) FilterProjectStageInfo(ctx context.Context, in *proto.StageFilterReq) (*proto.StageInfoReply, error) {
	stageInfo, err := models.QueryProjectStageByFilter(in.ProjectId, int(in.StageNumber))
	if err != nil || stageInfo.Id == 0 {
		return &proto.StageInfoReply{}, err
	}else {
		return &proto.StageInfoReply{
			Id:				stageInfo.Id,
			ProjectId:		stageInfo.ProjectId,
			StageNumber:	int32(stageInfo.StageNumber),
			CoinCount:		int32(stageInfo.CoinCount),
			TargetValue:	stageInfo.TargetValue,
			SoftValue:		stageInfo.SoftValue,
			Discount:		stageInfo.Discount,
			CompleteValue:	stageInfo.CompleteValue,
			StageStatus:	proto.Status(stageInfo.StageStatus),
			BeginTime:		utils.Time2String(stageInfo.BeginTime),
			EndTime:		utils.Time2String(stageInfo.EndTime),
			CreateTime:		utils.Time2String(stageInfo.CreateTime),
			UpdateTime:		utils.Time2String(stageInfo.UpdateTime),
		}, nil
	}
}

func (ctrl *ProjectController) QueryStageCoinInfo(ctx context.Context, in *proto.IdReq) (*proto.StageCoinListReply, error) {
	coinInfoList, err := models.QueryStageCoinInfoByStageId(in.Id)
	if err != nil || len(coinInfoList) == 0{
		return &proto.StageCoinListReply{}, err
	}else {
		coinResults := proto.StageCoinListReply{}
		for _, coin := range coinInfoList {
			//查询Currency Info
			var coinName string
			currencyInfo, err := models.QueryCurrencyById(coin.CoinId)
			if err != nil {
				fmt.Errorf("Query currency info failed due to error :%s", err.Error())
				coinName = ""
			} else {
				coinName = currencyInfo.CurrencyTag
			}
			coinResults.StageCoinList = append(coinResults.StageCoinList, &proto.StageCoinReply{
				StageId:			coin.StageId,
				CoinId:				coin.CoinId,
				CoinName:			coinName,
				CompleteValue:		coin.CompleteValue,
				MinValue:			coin.MinValue,
				MaxValue:			coin.MaxValue,
				Price:				coin.Price,
				CreateTime:			utils.Time2String(coin.CreateTime),
				UpdateTime:			utils.Time2String(coin.UpdateTime),
			})
		}

		return &coinResults, nil
	}
}

func (ctrl *ProjectController) FilterStageCoinInfo(ctx context.Context, in *proto.StageCoinFilterReq) (*proto.StageCoinReply, error) {
	coin, err := models.QueryStageCoinByFilter(in.StageId, in.CoinId)
	if err != nil {
		return &proto.StageCoinReply{}, err
	}else {
		return &proto.StageCoinReply{
			StageId:			coin.StageId,
			CoinId:				coin.CoinId,
			CompleteValue:		coin.CompleteValue,
			MinValue:			coin.MinValue,
			MaxValue:			coin.MaxValue,
			Price:				coin.Price,
			CreateTime:			utils.Time2String(coin.CreateTime),
			UpdateTime:			utils.Time2String(coin.UpdateTime),
		}, nil
	}
}

func (ctrl *ProjectController) QueryProjectNoticesByFilter(ctx context.Context, in *proto.NoticeFilterReq) (*proto.NoticeInfoListReply, error) {
	noticeList, err := models.QueryNoticesByFilter(in.ProjectId, int(in.NoticeType), int(in.NoticeState))
	if err != nil || len(noticeList) == 0 {
		return &proto.NoticeInfoListReply{}, err
	} else{
		noticeResults := proto.NoticeInfoListReply{}
		for _, notice := range noticeList {
			noticeResults.NoticeList = append(noticeResults.NoticeList, &proto.NoticeInfoReply{
				NoticeId:			notice.Id,
				ProjectId:			notice.ProjectId,
				Title:				notice.Title,
				Description:		notice.Description,
				NoticeType:			proto.NoticeType(notice.NoticeType),
				SendType:			proto.NoticeSendType(notice.SendType),
				Status:				proto.NoticeStateType(notice.Status),
				Reason:				notice.Reason,
				CreateTime:			utils.Time2String(notice.CreateTime),
				ExpireTime:			utils.Time2String(notice.ExpireTime),
				UpdateTime:			utils.Time2String(notice.UpdateTime),
			})
		}

		return &noticeResults, nil
	}
}

func (ctrl *ProjectController) QueryProjectNoticeVoteInfo(ctx context.Context, in *proto.IdReq) (*proto.NoticeVoteReply, error) {
	noticeVote, err := models.QueryNoticeVote(in.Id)
	if err != nil || noticeVote.Id == 0{
		return &proto.NoticeVoteReply{}, err
	} else {
		noticeNews, err := models.QueryNoticeNewsById(noticeVote.NoticeId)
		if err != nil {
			return &proto.NoticeVoteReply{}, err
		}

		project, err := models.QueryProjectById(noticeNews.ProjectId)
		if err != nil {
			return &proto.NoticeVoteReply{}, err
		}

		return &proto.NoticeVoteReply{
			NoticeId:			noticeVote.NoticeId,
			ApproveVote:		noticeVote.ApproveVote,
			DisapproveVote:		noticeVote.DisapproveVote,
			AbstentionVote:		noticeVote.AbstentionVote,
			PlatformVoteMax:	noticeVote.PlatformVoteMax,
			PlatformVoteVolumn:	noticeVote.PlatformVoteVolumn,
			PlatformVoteType:	proto.VoteResultType(noticeVote.PlatformVoteType),
			PlatformVoteReason:	noticeVote.PlatformVoteReason,
			PlatformVoteTime:	utils.Time2String(noticeVote.PlatformVoteTime),
			VoteResult:			proto.NoticeResultType(noticeVote.VoteResult),
			CreateTime:			utils.Time2String(noticeVote.CreateTime),
			UpdateTime:			utils.Time2String(noticeVote.UpdateTime),
			BeginTime:			utils.Time2String(noticeVote.BeginTime),
			EndTime:			utils.Time2String(noticeVote.EndTime),
			TotalVotes:			project.TargetValue,
		}, nil
	}
}

func (ctrl *ProjectController) CheckUserCanVoteNotice(ctx context.Context, in *proto.NoticeRightReq) (*proto.BoolReply, error) {
	resultInfo, err := models.CheckUserCanVoteNotice(in.UserId, in.NoticeId)
	if err != nil {
		return &proto.BoolReply{}, err
	}else{
		var result proto.BoolValue
		if resultInfo {
			result = proto.BoolValue_True
		}else {
			result = proto.BoolValue_False
		}
		return &proto.BoolReply{
			Result:	result,
		}, nil
	}
}

func (ctrl *ProjectController) UserNoticeVote(ctx context.Context, in *proto.UserVoteReq) (*proto.Empty, error) {
	err := models.UserVoteNotice(in.UserId, in.NoticeId, int(in.VoteType))

	return &proto.Empty{}, err
}

func (ctrl *ProjectController) JoinProjectIco(ctx context.Context, in *proto.IcoOrderReq) (*proto.IcoInfoReply, error) {
	icoInfoResult := proto.IcoInfoReply{}
	icoInfoResult.UserId = in.UserId

	//查询userBalance
	userBalance, err := models.QueryUserBalanceByFilter(in.UserId, in.CurrencyId)
	if err != nil {
		return &icoInfoResult, err
	}
	if userBalance.BalanceValue < in.PayCount {
		return &icoInfoResult, errors.New("User pay amount is over user balance value.")
	}

	//查询project状态
	projectInfo, err := models.QueryProjectById(in.ProjectId)
	if err != nil {
		return &icoInfoResult, err
	}
	if projectInfo.Status != models.Started {
		return &icoInfoResult, errors.New("Project is not in start status.")
	}
	icoInfoResult.ProjectId = projectInfo.Id
	icoInfoResult.LockType = int32(projectInfo.LockType)
	icoInfoResult.TargetCoin = projectInfo.IssueCoina

	//阶段信息
	stageInfo, err := models.QueryProjectStageByFilter(in.ProjectId, int(in.StageNumber))
	if err != nil {
		return &icoInfoResult, err
	}
	if stageInfo.Id == 0 {
		return &icoInfoResult, errors.New("Project stage number not correct.")
	}
	if stageInfo.StageStatus != models.Started {
		return &icoInfoResult, errors.New("Current project stage not started.")
	}
	icoInfoResult.StageId = stageInfo.Id

	//募集币种信息
	coinInfo, err := models.QueryStageCoinByFilter(stageInfo.Id, in.CurrencyId)
	if err != nil {
		return &icoInfoResult, err
	}
	if coinInfo.Id == 0 {
		return &icoInfoResult, errors.New("Current stage not raise currency you specified.")
	}
	if in.PayCount > coinInfo.MaxValue || in.PayCount < coinInfo.MinValue {
		return &icoInfoResult, errors.New("Ico order investment amount is not on limited.")
	}
	icoInfoResult.BaseCoin = coinInfo.CoinId
	icoInfoResult.Price = coinInfo.Price
	icoInfoResult.PayAmount = in.PayCount

	return &icoInfoResult, err
}

func (ctrl *ProjectController) CheckUserKycStatusInfo(ctx context.Context, in *proto.IdReq) (*proto.BoolReply, error) {
	status, err := models.CheckUserKycStatus(in.Id)
	if err != nil {
		return &proto.BoolReply{}, err
	} else{
		if status {
			return &proto.BoolReply{Result:proto.BoolValue_True}, nil
		}else{
			return &proto.BoolReply{Result:proto.BoolValue_False}, nil
		}
	}
}

//推荐项目
func (ctrl *ProjectController) QueryRecommendProjectsInfo(ctx context.Context, in *proto.Empty) (*proto.RecommendProjectListReply, error) {
	projects, err := models.QueryAllProjects()
	if err != nil || len(projects) == 0 {
		return &proto.RecommendProjectListReply{}, err
	}

	recommendProjects := proto.RecommendProjectListReply{}
	for _, project := range projects {
		recommandProject := proto.RecommendProjectReply{}
		recommandProject.ProjectId = project.Id
		recommandProject.Summary = project.Summary
		recommandProject.Description = project.Description
		recommandProject.ProjectStatus = int32(project.Status)

		//获取图片
		images, err := models.QueryProjectMediaList(project.Id, models.ImageType)
		if err != nil{
			fmt.Errorf("Get project image failed: %s", err)
		}
		if len(images) != 0 {
			recommandProject.ImageUrl = images[0].Address
		}

		recommendProjects.ProjectList = append(recommendProjects.ProjectList, &recommandProject)
	}

	return &recommendProjects, nil
}

//用户参与的项目
func (ctrl *ProjectController) QueryUserJoinProjectsInfo(ctx context.Context, in *proto.IdReq) (*proto.RecommendProjectListReply, error){
	projects, err := models.QueryProjectByUser(in.Id)
	if err != nil || len(projects) == 0 {
		return &proto.RecommendProjectListReply{}, err
	}

	recommendProjects := proto.RecommendProjectListReply{}
	for _, project := range projects {
		recommandProject := proto.RecommendProjectReply{}
		recommandProject.ProjectId = project.Id
		recommandProject.Summary = project.Summary
		recommandProject.Description = project.Description
		recommandProject.ProjectStatus = int32(project.Status)

		//获取图片
		images, err := models.QueryProjectMediaList(project.Id, models.ImageType)
		if err != nil {
			fmt.Errorf("Get project image failed: %s", err)
		}
		if len(images) != 0 {
			recommandProject.ImageUrl = images[0].Address
		}

		recommendProjects.ProjectList = append(recommendProjects.ProjectList, &recommandProject)
	}

	return &recommendProjects, nil
}

//项目详情
func (ctrl *ProjectController) QueryProjectDetailsInfo(ctx context.Context, in *proto.IdReq) (*proto.ProjectDetailsReply, error){
	project, err := models.QueryProjectById(in.Id)
	if err != nil {
		fmt.Errorf("Query project got error: %s", err.Error())
		return &proto.ProjectDetailsReply{}, err
	}
	if project.Summary == "" {
		fmt.Errorf("Query project return null.")
		return &proto.ProjectDetailsReply{}, err
	}

	projectDetails := proto.ProjectDetailsReply{}
	projectDetails.CoinId = project.IssueCoina
	projectDetails.Summary = project.Summary
	projectDetails.Description = project.Description
	projectDetails.Status = int32(project.Status)

	//判断是否可参与
	availableStage, err := models.QueryProjectAvailableStage(in.Id)
	if err != nil || availableStage == nil {
		projectDetails.CanJoin = proto.BoolValue(models.FalseValue)
	}else{
		projectDetails.CanJoin = proto.BoolValue(models.TrueValue)
	}

	//获取项目发布币名称
	coinInfo, err := models.QueryCurrencyById(project.IssueCoina)
	if err != nil {
		fmt.Errorf("Query currency info got error: %s", err.Error())
		return &proto.ProjectDetailsReply{}, err
	}
	projectDetails.CoinName = coinInfo.CurrencyTag

	//获取图片
	images, err := models.QueryProjectMediaList(project.Id, models.ImageType)
	if err != nil{
		fmt.Errorf("Get project image failed or return null: %s", err)
	}
	if len(images) != 0 {
		projectDetails.ImageUrl = images[0].Address
	}
	//获取勋章
	certifications, err := models.QueryProjectCertificationList(in.Id)
	if err != nil {
		fmt.Errorf("Get project certification failed or return null: %s", err)
	}
	if len(certifications) != 0 {
		for _, certi := range certifications {
			projectDetails.CertificationList = append(projectDetails.CertificationList, certi.Image)
		}
	}
	//预售期信息
	stageList, err := models.QueryDisplayStageInfo(in.Id)
	if err != nil {
		fmt.Errorf("Get project stage information failed: %s", err)
		return &projectDetails, err
	}
	if len(stageList) == 0 {
		fmt.Errorf("Get project stage information return null.")
		return &projectDetails, errors.New("Stage information is null.")
	}
	for _, stage := range stageList {
		sellStage := proto.SellStageReply{}
		sellStage.Name = stage.StageName
		sellStage.EndTime = utils.Time2String(stage.EndTime)
		sellStage.LeftDays = int32(models.GetDaysInterval(stage.EndTime))
		sellStage.Discount = stage.Discount

		//获取币种信息
		coinInfo, err := models.QueryStageCoinInfoByStageId(stage.Id)
		if err != nil {
			sellStage.CoinInfo = nil
			fmt.Errorf("Get stage coin info got error: %", err.Error())
		}
		if len(coinInfo) == 0 {
			sellStage.CoinInfo = nil
			fmt.Errorf("Get stage coin info return null.")
		} else{
			for _, coin := range coinInfo {
				currency, err := models.QueryCurrencyById(coin.CoinId)
				if err != nil {
					fmt.Errorf("Get currency info got error: %s", err.Error())
				}
				if currency.CurrencyTag == "" {
					fmt.Errorf("Get currency info return null.")
				}

				sellStage.CoinInfo = append(sellStage.CoinInfo, &proto.ProjectCoinReply{
					CoinName:		currency.CurrencyTag,
					TargetValue:	coin.TargetValue,
					CompleteValue:	coin.CompleteValue,
					MinValue:		coin.MinValue,
					MaxValue:		coin.MaxValue,
				})
			}
		}
		projectDetails.SellStageList = append(projectDetails.SellStageList, &sellStage)
	}

	//介绍
	descriptions, err := models.QueryProjectsTextByProject(in.Id)
	if err != nil {
		fmt.Errorf("Get project introduction failed or return null: %s", err)
		return &projectDetails, err
	}
	if len(descriptions) == 0 {
		fmt.Errorf("Get project introduction return null.")
	} else {
		for _, desc := range descriptions {
			projectDetails.ProjectContent = append(projectDetails.ProjectContent, &proto.ProjectDescription{
				Title:			desc.Title,
				Content:		desc.ProjectText,
				PrioritySort:	int32(desc.PrioritySort),
			})
		}
	}
	//团队
	membersList, err := models.QueryProjectMemberList(in.Id, models.Partner)
	if err != nil {
		fmt.Errorf("Get project member information failed: %s", err.Error())
	}
	if len(membersList) == 0 {
		fmt.Errorf("Get project member information return null.")
	} else {
		for _, member := range membersList {
			projectDetails.MemberList = append(projectDetails.MemberList, &proto.TeamMember{
				ImageUrl:		member.Image,
				Name:			member.Name,
				Position:		member.Position,
				PersonalInfo:	member.Description,
				PrioritySort:	int32(member.PrioritySort),
			})
		}
	}

	return &projectDetails, nil
}

//募集细则
func (ctrl *ProjectController) QueryRaiseInvestmentDetailsInfo(ctx context.Context, in *proto.IdReq) (*proto.RaiseRuleReply, error) {
	project, err := models.QueryProjectById(in.Id)
	if err != nil {
		fmt.Errorf("Get project introduction failed: %s", err)
		return &proto.RaiseRuleReply{}, err
	}
	if project.AdminId == 0 {
	fmt.Errorf("Get project introduction return null: %s", err)
	return &proto.RaiseRuleReply{}, err
	}

	raiseRule := proto.RaiseRuleReply{}

	//查询代码
	currency, err := models.QueryCurrencyById(project.IssueCoina)
	if err != nil {
		fmt.Errorf("Get currency information failed: %s", err)
	}
	if currency.CurrencyTag == ""{
		fmt.Errorf("Get currency information return null.")
	} else{
		raiseRule.CurrencyName = currency.CurrencyTag
	}

	raiseRule.SoftTarget = project.SoftValue
	raiseRule.HardTarget = project.HardValue
	raiseRule.TotalTarget = project.TargetValue

	//查询阶段信息
	stageList, err := models.QueryProjectStageList(in.Id)
	if err != nil {
		fmt.Errorf("Get stage information failed: %s", err)
	}
	if len(stageList) == 0 {
		fmt.Errorf("Get stage information return null.")
	} else {
		for _, stage := range stageList {
			stageInfo := proto.RaiseStageReply{}
			stageInfo.Name = stage.StageName
			stageInfo.BeginTime = utils.Time2String(stage.BeginTime)
			stageInfo.EndTime = utils.Time2String(stage.EndTime)
			stageInfo.TargetInfo = string(stage.TargetValue) + raiseRule.CurrencyName
			stageInfo.Discount = stage.Discount

			//Coin Info
			//获取币种信息
			coinInfo, err := models.QueryStageCoinInfoByStageId(stage.Id)
			if err != nil {
				stageInfo.CoinList = nil
				fmt.Errorf("Get stage coin info got error: %s", err.Error())
			}
			if len(coinInfo) == 0 {
				stageInfo.CoinList = nil
				fmt.Errorf("Get stage coin info return null.")
			} else{
				for _, coin := range coinInfo {
					currency, err := models.QueryCurrencyById(coin.CoinId)
					if err != nil {
						fmt.Errorf("Get currency info got error: %s", err.Error())
					}
					if currency.CurrencyTag == "" {
						fmt.Errorf("Get currency info return null.")
					}

					stageInfo.CoinList = append(stageInfo.CoinList, &proto.ProjectCoinReply{
						CoinName:		currency.CurrencyTag,
						TargetValue:	coin.TargetValue,
						CompleteValue:	coin.CompleteValue,
						MinValue:		coin.MinValue,
						MaxValue:		coin.MaxValue,
					})
				}
			}

			raiseRule.StageList = append(raiseRule.StageList, &stageInfo)
		}
	}

	return &raiseRule, nil
 }

 //查询项目公告列表
func (ctrl *ProjectController) QueryProjectNoticesInfo(ctx context.Context, in *proto.IdReq) (*proto.ProjectNoticeListReply, error) {
	noticeList, err := models.QueryNoticesByProjectId(in.Id)
	if err != nil {
		fmt.Errorf("Query project notice got error: %s", err.Error())
		return &proto.ProjectNoticeListReply{}, err
	}
	if len(noticeList) == 0 {
		fmt.Errorf("Query project notice return is null")
		return &proto.ProjectNoticeListReply{}, err
	} else {
		var projectNotices proto.ProjectNoticeListReply
		for _, notice := range noticeList {
			projectNotices.NoticeList = append(projectNotices.NoticeList, &proto.ProjectNoticeReply{
				NoticeId:			notice.Id,
				NoticeTitle:		notice.Title,
				Description:		notice.Description,
				Type:				proto.NoticeType(notice.NoticeType),
				CreateTime:			utils.Time2String(notice.CreateTime),
			})
		}

		return &projectNotices, nil
	}
}

//查询公告详情
func(ctrl *ProjectController) QueryProjectNoticeNewsDetailsInfo(ctx context.Context, in *proto.IdReq) (*proto.NoticeNewsReply, error) {
	notice, err := models.QueryNoticeNewsById(in.Id)
	if err != nil {
		fmt.Errorf("Query project notice news got error: %s", err.Error())
		return &proto.NoticeNewsReply{}, err
	}
	if notice.ProjectId == 0 {
		fmt.Errorf("Query project notice news return null.")
		return &proto.NoticeNewsReply{}, err
	}

	noticeNews := proto.NoticeNewsReply{}
	noticeNews.CreateTime = utils.Time2String(notice.CreateTime)
	noticeNews.NoticeTitle = notice.Title
	noticeNews.NoticeContent = notice.Description
	noticeNews.ImageUrl = notice.Image
	noticeNews.ImageWidth = int32(notice.ImageW)
	noticeNews.ImageHeight = int32(notice.ImageH)
	noticeNews.UrlSite = notice.UrlSite

	projectInfo, err := models.QueryProjectById(notice.ProjectId)
	if err != nil {
		fmt.Errorf("Query project information got error: %s", err.Error())
		return &noticeNews, err
	}
	if projectInfo.Summary == "" {
		fmt.Errorf("Query project information return null.")
		return &noticeNews, err
	}
	noticeNews.ProjectName = projectInfo.Summary

	return &noticeNews, nil
}

//查询投票详情
func (ctrl *ProjectController) QueryNoticeVoteDetailsInfo(ctx context.Context, in *proto.NoticeRightReq) (*proto.NoticeVoteDetailReply, error){
	noticeInfo, err := models.QueryNoticeNewsById(in.NoticeId)
	if err != nil {
		fmt.Errorf("Query project information got error or return null: %s", err.Error())
		return &proto.NoticeVoteDetailReply{}, err
	}
	if noticeInfo.ProjectId == 0 {
		fmt.Errorf("Query project information return null.")
		return &proto.NoticeVoteDetailReply{}, err
	}

	noticeVoteDetails := proto.NoticeVoteDetailReply{}
	noticeVoteDetails.NoticeTitle = noticeInfo.Title
	noticeVoteDetails.NoticeContent = noticeInfo.Description
	noticeVoteDetails.CreateTime = utils.Time2String(noticeInfo.CreateTime)
	noticeVoteDetails.ImageUrl = noticeInfo.Image
	noticeVoteDetails.ImageWidth = int32(noticeInfo.ImageW)
	noticeVoteDetails.ImageHeight = int32(noticeInfo.ImageH)
	noticeVoteDetails.UrlSite = noticeInfo.UrlSite

	//查询project
	projectInfo, err := models.QueryProjectById(noticeInfo.ProjectId)
	if err != nil {
		fmt.Errorf("Query project got error or return null: %s", err.Error())
		return &noticeVoteDetails, err
	}
	if projectInfo.TargetValue == 0 {
		fmt.Errorf("Query project return null.")
		return &noticeVoteDetails, err
	}
	noticeVoteDetails.ProjectName = projectInfo.Summary
	noticeVoteDetails.TotalVote = projectInfo.TargetValue

	voteDetails, err := models.QueryNoticeVoteDetailsByFilter(noticeInfo.Id, in.UserId)
	if err != nil {
		fmt.Errorf("Query notice vote details got error: %s", err.Error())
		return &noticeVoteDetails, err
	}
	if voteDetails.VoteVolumn != 0 {
		noticeVoteDetails.MyVoteVolumn = voteDetails.VoteVolumn
	} else{
		//查询UserBalance
		userBalance, err := models.QueryUserBalanceByFilter(in.UserId, projectInfo.IssueCoina)
		if err != nil {
			fmt.Errorf("Query user balance got error: %s", err.Error())
			return &noticeVoteDetails, err
		}
		if userBalance.BalanceId == "" {
			fmt.Errorf("Query user balance return null.")
			return &noticeVoteDetails, err
		}

		noticeVoteDetails.MyVoteVolumn = userBalance.BalanceValue
	}

	//查询Vote详细
	acc, disacc, abse, err := models.QueryVoteVolumnInfo(noticeInfo.Id)
	if err != nil {
		fmt.Errorf("Query vote volumn got error or return null: %s", err.Error())
		return &noticeVoteDetails, err
	}
	noticeVoteDetails.ApproveVolumn = acc
	noticeVoteDetails.DisapproveVolumn = disacc
	noticeVoteDetails.AbstensionVolumn = abse

	//查询结束时间
	noticeVote, err := models.QueryNoticeVote(noticeInfo.Id)
	if err != nil {
		fmt.Errorf("Query notice vote information failed: %s", err.Error())
		noticeVoteDetails.EndTime = utils.Time2String(time.Now())
	}
	 noticeVoteDetails.EndTime = utils.Time2String(noticeVote.EndTime)
	 noticeVoteDetails.NoticeResult = proto.NoticeResultType(noticeVote.VoteResult)
	 noticeVoteDetails.NoticeComments = noticeVote.VoteNote

	 //投票状态
	 if time.Now().Before(noticeVote.BeginTime) {
	 	noticeVoteDetails.VoteStatus = proto.Status(models.NotStart)
	 }
	 if time.Now().After(noticeVote.EndTime) {
		 noticeVoteDetails.VoteStatus = proto.Status(models.Completed)
	 }else {
		 noticeVoteDetails.VoteStatus = proto.Status(models.Started)
	 }

	//查询投票类型
	myVoteDetails, err := models.QueryNoticeVoteDetailsByFilter(noticeInfo.Id, in.UserId)
	if err != nil {
		fmt.Errorf("Query notice vote details failed or return null: %s", err.Error())
		return &noticeVoteDetails, err
	}
	noticeVoteDetails.MyVoteVolumn = myVoteDetails.VoteVolumn
	noticeVoteDetails.MyVoteType = proto.VoteResultType(myVoteDetails.VoteType)

	return &noticeVoteDetails, err
}