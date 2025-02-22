package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/metastudio/v1/model"
)

type MetaStudioClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewMetaStudioClient(hcClient *httpclient.HcHttpClient) *MetaStudioClient {
	return &MetaStudioClient{HcClient: hcClient}
}

func MetaStudioClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateDialogUrl 创建对话链接
//
// 该接口用于创建对话链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDialogUrl(request *model.CreateDialogUrlRequest) (*model.CreateDialogUrlResponse, error) {
	requestDef := GenReqDefForCreateDialogUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDialogUrlResponse), nil
	}
}

// CreateDialogUrlInvoker 创建对话链接
func (c *MetaStudioClient) CreateDialogUrlInvoker(request *model.CreateDialogUrlRequest) *CreateDialogUrlInvoker {
	requestDef := GenReqDefForCreateDialogUrl()
	return &CreateDialogUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartChatJob 查询数字人智能交互任务
//
// 该接口用于查询数字人智能交互任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartChatJob(request *model.ShowSmartChatJobRequest) (*model.ShowSmartChatJobResponse, error) {
	requestDef := GenReqDefForShowSmartChatJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartChatJobResponse), nil
	}
}

// ShowSmartChatJobInvoker 查询数字人智能交互任务
func (c *MetaStudioClient) ShowSmartChatJobInvoker(request *model.ShowSmartChatJobRequest) *ShowSmartChatJobInvoker {
	requestDef := GenReqDefForShowSmartChatJob()
	return &ShowSmartChatJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartSmartChatJob 启动数字人智能交互任务
//
// 该接口用于启动数字人智能交互任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StartSmartChatJob(request *model.StartSmartChatJobRequest) (*model.StartSmartChatJobResponse, error) {
	requestDef := GenReqDefForStartSmartChatJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartSmartChatJobResponse), nil
	}
}

// StartSmartChatJobInvoker 启动数字人智能交互任务
func (c *MetaStudioClient) StartSmartChatJobInvoker(request *model.StartSmartChatJobRequest) *StartSmartChatJobInvoker {
	requestDef := GenReqDefForStartSmartChatJob()
	return &StartSmartChatJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopSmartChatJob 结束数字人智能交互任务
//
// 该接口用于结束数字人智能交互任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StopSmartChatJob(request *model.StopSmartChatJobRequest) (*model.StopSmartChatJobResponse, error) {
	requestDef := GenReqDefForStopSmartChatJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSmartChatJobResponse), nil
	}
}

// StopSmartChatJobInvoker 结束数字人智能交互任务
func (c *MetaStudioClient) StopSmartChatJobInvoker(request *model.StopSmartChatJobRequest) *StopSmartChatJobInvoker {
	requestDef := GenReqDefForStopSmartChatJob()
	return &StopSmartChatJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDigitalAsset 创建资产
//
// 该接口用于在资产库中添加上传新的媒体资产。可上传的资产类型包括：分身数字人模型、背景图片、素材图片、素材视频、PPT等。
// * &gt; 资产类型是IMAGE时，通过system_properties来区分背景图片（BACKGROUND_IMG）、素材图片（MATERIAL_IMG）。
// * &gt; 资产类型是VIDEO时，通过system_properties来区分素材视频（MATERIAL_VIDEO）、名片视频（BUSSINESS_CARD_VIDEO）。
// * &gt; MetaStudio平台生成的视频，system_properties带CREATED_BY_PLATFORM。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDigitalAsset(request *model.CreateDigitalAssetRequest) (*model.CreateDigitalAssetResponse, error) {
	requestDef := GenReqDefForCreateDigitalAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDigitalAssetResponse), nil
	}
}

// CreateDigitalAssetInvoker 创建资产
func (c *MetaStudioClient) CreateDigitalAssetInvoker(request *model.CreateDigitalAssetRequest) *CreateDigitalAssetInvoker {
	requestDef := GenReqDefForCreateDigitalAsset()
	return &CreateDigitalAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAsset 删除资产
//
// 该接口用于删除资产库中的媒体资产。调用该接口删除媒体资产时，媒体资产会放入回收站中，不会彻底删除。如需彻底删除资产，需增加“mode&#x3D;force”参数配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteAsset(request *model.DeleteAssetRequest) (*model.DeleteAssetResponse, error) {
	requestDef := GenReqDefForDeleteAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetResponse), nil
	}
}

// DeleteAssetInvoker 删除资产
func (c *MetaStudioClient) DeleteAssetInvoker(request *model.DeleteAssetRequest) *DeleteAssetInvoker {
	requestDef := GenReqDefForDeleteAsset()
	return &DeleteAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssetSummary 查询资产概要
//
// 该接口用于查询媒体资产库中指定的多个资产的概要信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListAssetSummary(request *model.ListAssetSummaryRequest) (*model.ListAssetSummaryResponse, error) {
	requestDef := GenReqDefForListAssetSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetSummaryResponse), nil
	}
}

// ListAssetSummaryInvoker 查询资产概要
func (c *MetaStudioClient) ListAssetSummaryInvoker(request *model.ListAssetSummaryRequest) *ListAssetSummaryInvoker {
	requestDef := GenReqDefForListAssetSummary()
	return &ListAssetSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssets 查询资产列表
//
// 该接口用于查询资产库中的媒体资产列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListAssets(request *model.ListAssetsRequest) (*model.ListAssetsResponse, error) {
	requestDef := GenReqDefForListAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetsResponse), nil
	}
}

// ListAssetsInvoker 查询资产列表
func (c *MetaStudioClient) ListAssetsInvoker(request *model.ListAssetsRequest) *ListAssetsInvoker {
	requestDef := GenReqDefForListAssets()
	return &ListAssetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreAsset 恢复被删除的资产
//
// 该接口用于恢复被删除至回收站的媒体资产。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) RestoreAsset(request *model.RestoreAssetRequest) (*model.RestoreAssetResponse, error) {
	requestDef := GenReqDefForRestoreAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreAssetResponse), nil
	}
}

// RestoreAssetInvoker 恢复被删除的资产
func (c *MetaStudioClient) RestoreAssetInvoker(request *model.RestoreAssetRequest) *RestoreAssetInvoker {
	requestDef := GenReqDefForRestoreAsset()
	return &RestoreAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAsset 查询资产详情
//
// 该接口用于查询资产库中指定媒体资产的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowAsset(request *model.ShowAssetRequest) (*model.ShowAssetResponse, error) {
	requestDef := GenReqDefForShowAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetResponse), nil
	}
}

// ShowAssetInvoker 查询资产详情
func (c *MetaStudioClient) ShowAssetInvoker(request *model.ShowAssetRequest) *ShowAssetInvoker {
	requestDef := GenReqDefForShowAsset()
	return &ShowAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDigitalAsset 更新资产
//
// 该接口用于更新资产库中的媒体资产信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDigitalAsset(request *model.UpdateDigitalAssetRequest) (*model.UpdateDigitalAssetResponse, error) {
	requestDef := GenReqDefForUpdateDigitalAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDigitalAssetResponse), nil
	}
}

// UpdateDigitalAssetInvoker 更新资产
func (c *MetaStudioClient) UpdateDigitalAssetInvoker(request *model.UpdateDigitalAssetRequest) *UpdateDigitalAssetInvoker {
	requestDef := GenReqDefForUpdateDigitalAsset()
	return &UpdateDigitalAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDigitalHumanBusinessCard 创建数字人名片制作
//
// 该接口用于数字人名片制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDigitalHumanBusinessCard(request *model.CreateDigitalHumanBusinessCardRequest) (*model.CreateDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForCreateDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDigitalHumanBusinessCardResponse), nil
	}
}

// CreateDigitalHumanBusinessCardInvoker 创建数字人名片制作
func (c *MetaStudioClient) CreateDigitalHumanBusinessCardInvoker(request *model.CreateDigitalHumanBusinessCardRequest) *CreateDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForCreateDigitalHumanBusinessCard()
	return &CreateDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDigitalHumanBusinessCard 删除数字人名片制作任务
//
// 该接口用于删除数字人名片制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteDigitalHumanBusinessCard(request *model.DeleteDigitalHumanBusinessCardRequest) (*model.DeleteDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForDeleteDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDigitalHumanBusinessCardResponse), nil
	}
}

// DeleteDigitalHumanBusinessCardInvoker 删除数字人名片制作任务
func (c *MetaStudioClient) DeleteDigitalHumanBusinessCardInvoker(request *model.DeleteDigitalHumanBusinessCardRequest) *DeleteDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForDeleteDigitalHumanBusinessCard()
	return &DeleteDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDigitalHumanBusinessCard 查询数字人名片制作任务列表
//
// 该接口用于查询数字人名片制作任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListDigitalHumanBusinessCard(request *model.ListDigitalHumanBusinessCardRequest) (*model.ListDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForListDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDigitalHumanBusinessCardResponse), nil
	}
}

// ListDigitalHumanBusinessCardInvoker 查询数字人名片制作任务列表
func (c *MetaStudioClient) ListDigitalHumanBusinessCardInvoker(request *model.ListDigitalHumanBusinessCardRequest) *ListDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForListDigitalHumanBusinessCard()
	return &ListDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDigitalHumanBusinessCard 查询数字人名片制作任务详情
//
// 该接口用于查询数字人名片制作任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowDigitalHumanBusinessCard(request *model.ShowDigitalHumanBusinessCardRequest) (*model.ShowDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForShowDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDigitalHumanBusinessCardResponse), nil
	}
}

// ShowDigitalHumanBusinessCardInvoker 查询数字人名片制作任务详情
func (c *MetaStudioClient) ShowDigitalHumanBusinessCardInvoker(request *model.ShowDigitalHumanBusinessCardRequest) *ShowDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForShowDigitalHumanBusinessCard()
	return &ShowDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDigitalHumanBusinessCard 更新数字人名片制作
//
// 该接口用于更新数字人名片制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDigitalHumanBusinessCard(request *model.UpdateDigitalHumanBusinessCardRequest) (*model.UpdateDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForUpdateDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDigitalHumanBusinessCardResponse), nil
	}
}

// UpdateDigitalHumanBusinessCardInvoker 更新数字人名片制作
func (c *MetaStudioClient) UpdateDigitalHumanBusinessCardInvoker(request *model.UpdateDigitalHumanBusinessCardRequest) *UpdateDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForUpdateDigitalHumanBusinessCard()
	return &UpdateDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDigitalHumanVideo 查询视频制作任务列表
//
// 该接口用于查询视频制作任务列表。可查询分身数字人视频制作列表，照片数字人视频制作列表等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListDigitalHumanVideo(request *model.ListDigitalHumanVideoRequest) (*model.ListDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForListDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDigitalHumanVideoResponse), nil
	}
}

// ListDigitalHumanVideoInvoker 查询视频制作任务列表
func (c *MetaStudioClient) ListDigitalHumanVideoInvoker(request *model.ListDigitalHumanVideoRequest) *ListDigitalHumanVideoInvoker {
	requestDef := GenReqDefForListDigitalHumanVideo()
	return &ListDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Cancel2DDigitalHumanVideo 取消等待中的分身数字人视频制作任务
//
// 该接口用于取消等待中的分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Cancel2DDigitalHumanVideo(request *model.Cancel2DDigitalHumanVideoRequest) (*model.Cancel2DDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCancel2DDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Cancel2DDigitalHumanVideoResponse), nil
	}
}

// Cancel2DDigitalHumanVideoInvoker 取消等待中的分身数字人视频制作任务
func (c *MetaStudioClient) Cancel2DDigitalHumanVideoInvoker(request *model.Cancel2DDigitalHumanVideoRequest) *Cancel2DDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCancel2DDigitalHumanVideo()
	return &Cancel2DDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Create2DDigitalHumanVideo 创建分身数字人视频制作任务
//
// 该接口用于创建分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Create2DDigitalHumanVideo(request *model.Create2DDigitalHumanVideoRequest) (*model.Create2DDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCreate2DDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Create2DDigitalHumanVideoResponse), nil
	}
}

// Create2DDigitalHumanVideoInvoker 创建分身数字人视频制作任务
func (c *MetaStudioClient) Create2DDigitalHumanVideoInvoker(request *model.Create2DDigitalHumanVideoRequest) *Create2DDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCreate2DDigitalHumanVideo()
	return &Create2DDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Show2DDigitalHumanVideo 查询分身数字人视频制作任务详情
//
// 该接口用于查询分身数字人视频制作任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Show2DDigitalHumanVideo(request *model.Show2DDigitalHumanVideoRequest) (*model.Show2DDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForShow2DDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Show2DDigitalHumanVideoResponse), nil
	}
}

// Show2DDigitalHumanVideoInvoker 查询分身数字人视频制作任务详情
func (c *MetaStudioClient) Show2DDigitalHumanVideoInvoker(request *model.Show2DDigitalHumanVideoRequest) *Show2DDigitalHumanVideoInvoker {
	requestDef := GenReqDefForShow2DDigitalHumanVideo()
	return &Show2DDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelPhotoDigitalHumanVideo 取消等待中的照片分身数字人视频制作任务
//
// 该接口用于取消等待中的照片分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CancelPhotoDigitalHumanVideo(request *model.CancelPhotoDigitalHumanVideoRequest) (*model.CancelPhotoDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCancelPhotoDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelPhotoDigitalHumanVideoResponse), nil
	}
}

// CancelPhotoDigitalHumanVideoInvoker 取消等待中的照片分身数字人视频制作任务
func (c *MetaStudioClient) CancelPhotoDigitalHumanVideoInvoker(request *model.CancelPhotoDigitalHumanVideoRequest) *CancelPhotoDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCancelPhotoDigitalHumanVideo()
	return &CancelPhotoDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePhotoDetection 创建照片检测任务
//
// 该接口用于创建照片检测任务，检测照片是否满足制作照片数字人的要求。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePhotoDetection(request *model.CreatePhotoDetectionRequest) (*model.CreatePhotoDetectionResponse, error) {
	requestDef := GenReqDefForCreatePhotoDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePhotoDetectionResponse), nil
	}
}

// CreatePhotoDetectionInvoker 创建照片检测任务
func (c *MetaStudioClient) CreatePhotoDetectionInvoker(request *model.CreatePhotoDetectionRequest) *CreatePhotoDetectionInvoker {
	requestDef := GenReqDefForCreatePhotoDetection()
	return &CreatePhotoDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePhotoDigitalHumanVideo 创建照片分身数字人视频制作任务
//
// 该接口用于创建照片分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePhotoDigitalHumanVideo(request *model.CreatePhotoDigitalHumanVideoRequest) (*model.CreatePhotoDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCreatePhotoDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePhotoDigitalHumanVideoResponse), nil
	}
}

// CreatePhotoDigitalHumanVideoInvoker 创建照片分身数字人视频制作任务
func (c *MetaStudioClient) CreatePhotoDigitalHumanVideoInvoker(request *model.CreatePhotoDigitalHumanVideoRequest) *CreatePhotoDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCreatePhotoDigitalHumanVideo()
	return &CreatePhotoDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPhotoDetection 查询照片检测任务详情
//
// 该接口用于查询照片检测任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPhotoDetection(request *model.ShowPhotoDetectionRequest) (*model.ShowPhotoDetectionResponse, error) {
	requestDef := GenReqDefForShowPhotoDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPhotoDetectionResponse), nil
	}
}

// ShowPhotoDetectionInvoker 查询照片检测任务详情
func (c *MetaStudioClient) ShowPhotoDetectionInvoker(request *model.ShowPhotoDetectionRequest) *ShowPhotoDetectionInvoker {
	requestDef := GenReqDefForShowPhotoDetection()
	return &ShowPhotoDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPhotoDigitalHumanVideo 查询照片分身数字人视频制作任务详情
//
// 该接口用于查询照片分身数字人视频制作任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPhotoDigitalHumanVideo(request *model.ShowPhotoDigitalHumanVideoRequest) (*model.ShowPhotoDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForShowPhotoDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPhotoDigitalHumanVideoResponse), nil
	}
}

// ShowPhotoDigitalHumanVideoInvoker 查询照片分身数字人视频制作任务详情
func (c *MetaStudioClient) ShowPhotoDigitalHumanVideoInvoker(request *model.ShowPhotoDigitalHumanVideoRequest) *ShowPhotoDigitalHumanVideoInvoker {
	requestDef := GenReqDefForShowPhotoDigitalHumanVideo()
	return &ShowPhotoDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmFileUpload 确认文件已上传
//
// 资产文件上传完毕后，通过该接口确认上传完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ConfirmFileUpload(request *model.ConfirmFileUploadRequest) (*model.ConfirmFileUploadResponse, error) {
	requestDef := GenReqDefForConfirmFileUpload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmFileUploadResponse), nil
	}
}

// ConfirmFileUploadInvoker 确认文件已上传
func (c *MetaStudioClient) ConfirmFileUploadInvoker(request *model.ConfirmFileUploadRequest) *ConfirmFileUploadInvoker {
	requestDef := GenReqDefForConfirmFileUpload()
	return &ConfirmFileUploadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFile 创建文件并获取上传URL
//
// 该接口用于创建文件并获取上传URL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateFile(request *model.CreateFileRequest) (*model.CreateFileResponse, error) {
	requestDef := GenReqDefForCreateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFileResponse), nil
	}
}

// CreateFileInvoker 创建文件并获取上传URL
func (c *MetaStudioClient) CreateFileInvoker(request *model.CreateFileRequest) *CreateFileInvoker {
	requestDef := GenReqDefForCreateFile()
	return &CreateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFile 删除文件
//
// 该接口用于删除媒体资产库中指定的文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteFile(request *model.DeleteFileRequest) (*model.DeleteFileResponse, error) {
	requestDef := GenReqDefForDeleteFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFileResponse), nil
	}
}

// DeleteFileInvoker 删除文件
func (c *MetaStudioClient) DeleteFileInvoker(request *model.DeleteFileRequest) *DeleteFileInvoker {
	requestDef := GenReqDefForDeleteFile()
	return &DeleteFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOnceCode 创建一次性鉴权码
//
// 该接口用于创建一次性鉴权码，有效期5分钟，鉴权码只能使用一次，每次使用后需要重新获取。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateOnceCode(request *model.CreateOnceCodeRequest) (*model.CreateOnceCodeResponse, error) {
	requestDef := GenReqDefForCreateOnceCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOnceCodeResponse), nil
	}
}

// CreateOnceCodeInvoker 创建一次性鉴权码
func (c *MetaStudioClient) CreateOnceCodeInvoker(request *model.CreateOnceCodeRequest) *CreateOnceCodeInvoker {
	requestDef := GenReqDefForCreateOnceCode()
	return &CreateOnceCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePictureModelingByUrlJob 基于图片URL创建照片建模任务
//
// 该接口用于从URL中获取图片进行照片建模任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePictureModelingByUrlJob(request *model.CreatePictureModelingByUrlJobRequest) (*model.CreatePictureModelingByUrlJobResponse, error) {
	requestDef := GenReqDefForCreatePictureModelingByUrlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePictureModelingByUrlJobResponse), nil
	}
}

// CreatePictureModelingByUrlJobInvoker 基于图片URL创建照片建模任务
func (c *MetaStudioClient) CreatePictureModelingByUrlJobInvoker(request *model.CreatePictureModelingByUrlJobRequest) *CreatePictureModelingByUrlJobInvoker {
	requestDef := GenReqDefForCreatePictureModelingByUrlJob()
	return &CreatePictureModelingByUrlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePictureModelingJob 创建照片建模任务
//
// 该接口用于创建风格化照片建模任务。通过上传照片，生成风格化数字人模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePictureModelingJob(request *model.CreatePictureModelingJobRequest) (*model.CreatePictureModelingJobResponse, error) {
	requestDef := GenReqDefForCreatePictureModelingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePictureModelingJobResponse), nil
	}
}

// CreatePictureModelingJobInvoker 创建照片建模任务
func (c *MetaStudioClient) CreatePictureModelingJobInvoker(request *model.CreatePictureModelingJobRequest) *CreatePictureModelingJobInvoker {
	requestDef := GenReqDefForCreatePictureModelingJob()
	return &CreatePictureModelingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPictureModelingJobs 照片建模任务列表查询
//
// 该接口用于查询风格化照片建模任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListPictureModelingJobs(request *model.ListPictureModelingJobsRequest) (*model.ListPictureModelingJobsResponse, error) {
	requestDef := GenReqDefForListPictureModelingJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPictureModelingJobsResponse), nil
	}
}

// ListPictureModelingJobsInvoker 照片建模任务列表查询
func (c *MetaStudioClient) ListPictureModelingJobsInvoker(request *model.ListPictureModelingJobsRequest) *ListPictureModelingJobsInvoker {
	requestDef := GenReqDefForListPictureModelingJobs()
	return &ListPictureModelingJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPictureModelingJob 照片建模任务详情查询
//
// 该接口用于风格化查询照片建模任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPictureModelingJob(request *model.ShowPictureModelingJobRequest) (*model.ShowPictureModelingJobResponse, error) {
	requestDef := GenReqDefForShowPictureModelingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPictureModelingJobResponse), nil
	}
}

// ShowPictureModelingJobInvoker 照片建模任务详情查询
func (c *MetaStudioClient) ShowPictureModelingJobInvoker(request *model.ShowPictureModelingJobRequest) *ShowPictureModelingJobInvoker {
	requestDef := GenReqDefForShowPictureModelingJob()
	return &ShowPictureModelingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRobot 创建应用
//
// 该接口用于创建应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateRobot(request *model.CreateRobotRequest) (*model.CreateRobotResponse, error) {
	requestDef := GenReqDefForCreateRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRobotResponse), nil
	}
}

// CreateRobotInvoker 创建应用
func (c *MetaStudioClient) CreateRobotInvoker(request *model.CreateRobotRequest) *CreateRobotInvoker {
	requestDef := GenReqDefForCreateRobot()
	return &CreateRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRobot 删除应用
//
// 该接口用于删除应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteRobot(request *model.DeleteRobotRequest) (*model.DeleteRobotResponse, error) {
	requestDef := GenReqDefForDeleteRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRobotResponse), nil
	}
}

// DeleteRobotInvoker 删除应用
func (c *MetaStudioClient) DeleteRobotInvoker(request *model.DeleteRobotRequest) *DeleteRobotInvoker {
	requestDef := GenReqDefForDeleteRobot()
	return &DeleteRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRobot 查询应用列表
//
// 该接口用于查询应用列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListRobot(request *model.ListRobotRequest) (*model.ListRobotResponse, error) {
	requestDef := GenReqDefForListRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRobotResponse), nil
	}
}

// ListRobotInvoker 查询应用列表
func (c *MetaStudioClient) ListRobotInvoker(request *model.ListRobotRequest) *ListRobotInvoker {
	requestDef := GenReqDefForListRobot()
	return &ListRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRobot 查询应用详情
//
// 该接口用于查询应用详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowRobot(request *model.ShowRobotRequest) (*model.ShowRobotResponse, error) {
	requestDef := GenReqDefForShowRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRobotResponse), nil
	}
}

// ShowRobotInvoker 查询应用详情
func (c *MetaStudioClient) ShowRobotInvoker(request *model.ShowRobotRequest) *ShowRobotInvoker {
	requestDef := GenReqDefForShowRobot()
	return &ShowRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRobot 修改应用
//
// 该接口用于修改应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateRobot(request *model.UpdateRobotRequest) (*model.UpdateRobotResponse, error) {
	requestDef := GenReqDefForUpdateRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRobotResponse), nil
	}
}

// UpdateRobotInvoker 修改应用
func (c *MetaStudioClient) UpdateRobotInvoker(request *model.UpdateRobotRequest) *UpdateRobotInvoker {
	requestDef := GenReqDefForUpdateRobot()
	return &UpdateRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSmartChatRoom 创建智能交互对话直播间
//
// 该接口用于创建智能交互对话直播间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateSmartChatRoom(request *model.CreateSmartChatRoomRequest) (*model.CreateSmartChatRoomResponse, error) {
	requestDef := GenReqDefForCreateSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSmartChatRoomResponse), nil
	}
}

// CreateSmartChatRoomInvoker 创建智能交互对话直播间
func (c *MetaStudioClient) CreateSmartChatRoomInvoker(request *model.CreateSmartChatRoomRequest) *CreateSmartChatRoomInvoker {
	requestDef := GenReqDefForCreateSmartChatRoom()
	return &CreateSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSmartChatRoom 删除智能交互对话直播间
//
// 该接口用于删除智能交互对话直播间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteSmartChatRoom(request *model.DeleteSmartChatRoomRequest) (*model.DeleteSmartChatRoomResponse, error) {
	requestDef := GenReqDefForDeleteSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSmartChatRoomResponse), nil
	}
}

// DeleteSmartChatRoomInvoker 删除智能交互对话直播间
func (c *MetaStudioClient) DeleteSmartChatRoomInvoker(request *model.DeleteSmartChatRoomRequest) *DeleteSmartChatRoomInvoker {
	requestDef := GenReqDefForDeleteSmartChatRoom()
	return &DeleteSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartChatRooms 查询智能交互对话直播间列表
//
// 该接口用于智能交互对话直播间列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartChatRooms(request *model.ListSmartChatRoomsRequest) (*model.ListSmartChatRoomsResponse, error) {
	requestDef := GenReqDefForListSmartChatRooms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartChatRoomsResponse), nil
	}
}

// ListSmartChatRoomsInvoker 查询智能交互对话直播间列表
func (c *MetaStudioClient) ListSmartChatRoomsInvoker(request *model.ListSmartChatRoomsRequest) *ListSmartChatRoomsInvoker {
	requestDef := GenReqDefForListSmartChatRooms()
	return &ListSmartChatRoomsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartChatRoom 查询智能交互对话直播间详情
//
// 该接口用于查询智能交互对话直播间详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartChatRoom(request *model.ShowSmartChatRoomRequest) (*model.ShowSmartChatRoomResponse, error) {
	requestDef := GenReqDefForShowSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartChatRoomResponse), nil
	}
}

// ShowSmartChatRoomInvoker 查询智能交互对话直播间详情
func (c *MetaStudioClient) ShowSmartChatRoomInvoker(request *model.ShowSmartChatRoomRequest) *ShowSmartChatRoomInvoker {
	requestDef := GenReqDefForShowSmartChatRoom()
	return &ShowSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSmartChatRoom 更新智能交互对话直播间信息
//
// 该接口用于智能交互对话直播间信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateSmartChatRoom(request *model.UpdateSmartChatRoomRequest) (*model.UpdateSmartChatRoomResponse, error) {
	requestDef := GenReqDefForUpdateSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSmartChatRoomResponse), nil
	}
}

// UpdateSmartChatRoomInvoker 更新智能交互对话直播间信息
func (c *MetaStudioClient) UpdateSmartChatRoomInvoker(request *model.UpdateSmartChatRoomRequest) *UpdateSmartChatRoomInvoker {
	requestDef := GenReqDefForUpdateSmartChatRoom()
	return &UpdateSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteSmartLiveCommand 控制数字人直播过程
//
// 该接口用于控制数字人直播过程。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ExecuteSmartLiveCommand(request *model.ExecuteSmartLiveCommandRequest) (*model.ExecuteSmartLiveCommandResponse, error) {
	requestDef := GenReqDefForExecuteSmartLiveCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSmartLiveCommandResponse), nil
	}
}

// ExecuteSmartLiveCommandInvoker 控制数字人直播过程
func (c *MetaStudioClient) ExecuteSmartLiveCommandInvoker(request *model.ExecuteSmartLiveCommandRequest) *ExecuteSmartLiveCommandInvoker {
	requestDef := GenReqDefForExecuteSmartLiveCommand()
	return &ExecuteSmartLiveCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLive 查询数字人智能直播任务列表
//
// 该接口用于查询数字人智能直播任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLive(request *model.ListSmartLiveRequest) (*model.ListSmartLiveResponse, error) {
	requestDef := GenReqDefForListSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveResponse), nil
	}
}

// ListSmartLiveInvoker 查询数字人智能直播任务列表
func (c *MetaStudioClient) ListSmartLiveInvoker(request *model.ListSmartLiveRequest) *ListSmartLiveInvoker {
	requestDef := GenReqDefForListSmartLive()
	return &ListSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLiveJobs 查询数字人智能直播任务列表
//
// 该接口用于查询数字人智能直播任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLiveJobs(request *model.ListSmartLiveJobsRequest) (*model.ListSmartLiveJobsResponse, error) {
	requestDef := GenReqDefForListSmartLiveJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveJobsResponse), nil
	}
}

// ListSmartLiveJobsInvoker 查询数字人智能直播任务列表
func (c *MetaStudioClient) ListSmartLiveJobsInvoker(request *model.ListSmartLiveJobsRequest) *ListSmartLiveJobsInvoker {
	requestDef := GenReqDefForListSmartLiveJobs()
	return &ListSmartLiveJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LiveEventReport 上报直播间事件
//
// 该接口用于上报直播间事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) LiveEventReport(request *model.LiveEventReportRequest) (*model.LiveEventReportResponse, error) {
	requestDef := GenReqDefForLiveEventReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LiveEventReportResponse), nil
	}
}

// LiveEventReportInvoker 上报直播间事件
func (c *MetaStudioClient) LiveEventReportInvoker(request *model.LiveEventReportRequest) *LiveEventReportInvoker {
	requestDef := GenReqDefForLiveEventReport()
	return &LiveEventReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartLive 查询数字人智能直播任务详情
//
// 该接口用于查询数字人智能直播任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartLive(request *model.ShowSmartLiveRequest) (*model.ShowSmartLiveResponse, error) {
	requestDef := GenReqDefForShowSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartLiveResponse), nil
	}
}

// ShowSmartLiveInvoker 查询数字人智能直播任务详情
func (c *MetaStudioClient) ShowSmartLiveInvoker(request *model.ShowSmartLiveRequest) *ShowSmartLiveInvoker {
	requestDef := GenReqDefForShowSmartLive()
	return &ShowSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartSmartLive 启动数字人智能直播任务
//
// 该接口用于启动数字人智能直播任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StartSmartLive(request *model.StartSmartLiveRequest) (*model.StartSmartLiveResponse, error) {
	requestDef := GenReqDefForStartSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartSmartLiveResponse), nil
	}
}

// StartSmartLiveInvoker 启动数字人智能直播任务
func (c *MetaStudioClient) StartSmartLiveInvoker(request *model.StartSmartLiveRequest) *StartSmartLiveInvoker {
	requestDef := GenReqDefForStartSmartLive()
	return &StartSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopSmartLive 结束数字人智能直播任务
//
// 该接口用于结束数字人智能直播任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StopSmartLive(request *model.StopSmartLiveRequest) (*model.StopSmartLiveResponse, error) {
	requestDef := GenReqDefForStopSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSmartLiveResponse), nil
	}
}

// StopSmartLiveInvoker 结束数字人智能直播任务
func (c *MetaStudioClient) StopSmartLiveInvoker(request *model.StopSmartLiveRequest) *StopSmartLiveInvoker {
	requestDef := GenReqDefForStopSmartLive()
	return &StopSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckTextLanguage 检测音色与文本的语言一致性
//
// 检测音色与文本的语言一致性，音色与文本不一致会导致报错
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CheckTextLanguage(request *model.CheckTextLanguageRequest) (*model.CheckTextLanguageResponse, error) {
	requestDef := GenReqDefForCheckTextLanguage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckTextLanguageResponse), nil
	}
}

// CheckTextLanguageInvoker 检测音色与文本的语言一致性
func (c *MetaStudioClient) CheckTextLanguageInvoker(request *model.CheckTextLanguageRequest) *CheckTextLanguageInvoker {
	requestDef := GenReqDefForCheckTextLanguage()
	return &CheckTextLanguageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInteractionRuleGroup 创建智能直播间互动规则库
//
// 该接口用于创建智能直播间互动规则库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateInteractionRuleGroup(request *model.CreateInteractionRuleGroupRequest) (*model.CreateInteractionRuleGroupResponse, error) {
	requestDef := GenReqDefForCreateInteractionRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInteractionRuleGroupResponse), nil
	}
}

// CreateInteractionRuleGroupInvoker 创建智能直播间互动规则库
func (c *MetaStudioClient) CreateInteractionRuleGroupInvoker(request *model.CreateInteractionRuleGroupRequest) *CreateInteractionRuleGroupInvoker {
	requestDef := GenReqDefForCreateInteractionRuleGroup()
	return &CreateInteractionRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSmartLiveRoom 创建智能直播间
//
// 该接口用于创建智能直播间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateSmartLiveRoom(request *model.CreateSmartLiveRoomRequest) (*model.CreateSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForCreateSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSmartLiveRoomResponse), nil
	}
}

// CreateSmartLiveRoomInvoker 创建智能直播间
func (c *MetaStudioClient) CreateSmartLiveRoomInvoker(request *model.CreateSmartLiveRoomRequest) *CreateSmartLiveRoomInvoker {
	requestDef := GenReqDefForCreateSmartLiveRoom()
	return &CreateSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInteractionRuleGroup 删除智能直播间互动规则库
//
// 该接口用于删除智能直播间互动规则库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteInteractionRuleGroup(request *model.DeleteInteractionRuleGroupRequest) (*model.DeleteInteractionRuleGroupResponse, error) {
	requestDef := GenReqDefForDeleteInteractionRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInteractionRuleGroupResponse), nil
	}
}

// DeleteInteractionRuleGroupInvoker 删除智能直播间互动规则库
func (c *MetaStudioClient) DeleteInteractionRuleGroupInvoker(request *model.DeleteInteractionRuleGroupRequest) *DeleteInteractionRuleGroupInvoker {
	requestDef := GenReqDefForDeleteInteractionRuleGroup()
	return &DeleteInteractionRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSmartLiveRoom 删除智能直播间
//
// 该接口用于删除智能直播间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteSmartLiveRoom(request *model.DeleteSmartLiveRoomRequest) (*model.DeleteSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForDeleteSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSmartLiveRoomResponse), nil
	}
}

// DeleteSmartLiveRoomInvoker 删除智能直播间
func (c *MetaStudioClient) DeleteSmartLiveRoomInvoker(request *model.DeleteSmartLiveRoomRequest) *DeleteSmartLiveRoomInvoker {
	requestDef := GenReqDefForDeleteSmartLiveRoom()
	return &DeleteSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInteractionRuleGroups 查询智能直播间互动规则库列表
//
// 该接口用于智能直播间互动规则库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListInteractionRuleGroups(request *model.ListInteractionRuleGroupsRequest) (*model.ListInteractionRuleGroupsResponse, error) {
	requestDef := GenReqDefForListInteractionRuleGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInteractionRuleGroupsResponse), nil
	}
}

// ListInteractionRuleGroupsInvoker 查询智能直播间互动规则库列表
func (c *MetaStudioClient) ListInteractionRuleGroupsInvoker(request *model.ListInteractionRuleGroupsRequest) *ListInteractionRuleGroupsInvoker {
	requestDef := GenReqDefForListInteractionRuleGroups()
	return &ListInteractionRuleGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLiveRooms 查询智能直播间列表
//
// 该接口用于智能直播间列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLiveRooms(request *model.ListSmartLiveRoomsRequest) (*model.ListSmartLiveRoomsResponse, error) {
	requestDef := GenReqDefForListSmartLiveRooms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveRoomsResponse), nil
	}
}

// ListSmartLiveRoomsInvoker 查询智能直播间列表
func (c *MetaStudioClient) ListSmartLiveRoomsInvoker(request *model.ListSmartLiveRoomsRequest) *ListSmartLiveRoomsInvoker {
	requestDef := GenReqDefForListSmartLiveRooms()
	return &ListSmartLiveRoomsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartLiveRoom 查询智能直播剧本详情
//
// 该接口用于查询智能直播剧本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartLiveRoom(request *model.ShowSmartLiveRoomRequest) (*model.ShowSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForShowSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartLiveRoomResponse), nil
	}
}

// ShowSmartLiveRoomInvoker 查询智能直播剧本详情
func (c *MetaStudioClient) ShowSmartLiveRoomInvoker(request *model.ShowSmartLiveRoomRequest) *ShowSmartLiveRoomInvoker {
	requestDef := GenReqDefForShowSmartLiveRoom()
	return &ShowSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInteractionRuleGroup 更新智能直播间互动规则库
//
// 该接口用于更新智能直播间互动规则库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateInteractionRuleGroup(request *model.UpdateInteractionRuleGroupRequest) (*model.UpdateInteractionRuleGroupResponse, error) {
	requestDef := GenReqDefForUpdateInteractionRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInteractionRuleGroupResponse), nil
	}
}

// UpdateInteractionRuleGroupInvoker 更新智能直播间互动规则库
func (c *MetaStudioClient) UpdateInteractionRuleGroupInvoker(request *model.UpdateInteractionRuleGroupRequest) *UpdateInteractionRuleGroupInvoker {
	requestDef := GenReqDefForUpdateInteractionRuleGroup()
	return &UpdateInteractionRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSmartLiveRoom 更新智能直播间信息
//
// 该接口用于智能直播间信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateSmartLiveRoom(request *model.UpdateSmartLiveRoomRequest) (*model.UpdateSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForUpdateSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSmartLiveRoomResponse), nil
	}
}

// UpdateSmartLiveRoomInvoker 更新智能直播间信息
func (c *MetaStudioClient) UpdateSmartLiveRoomInvoker(request *model.UpdateSmartLiveRoomRequest) *UpdateSmartLiveRoomInvoker {
	requestDef := GenReqDefForUpdateSmartLiveRoom()
	return &UpdateSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStyles 查询数字人风格列表
//
// 查询数字人风格列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListStyles(request *model.ListStylesRequest) (*model.ListStylesResponse, error) {
	requestDef := GenReqDefForListStyles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStylesResponse), nil
	}
}

// ListStylesInvoker 查询数字人风格列表
func (c *MetaStudioClient) ListStylesInvoker(request *model.ListStylesRequest) *ListStylesInvoker {
	requestDef := GenReqDefForListStyles()
	return &ListStylesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CommitVoiceTrainingJob 提交语音训练任务
//
// 提交训练任务,执行该接口后,任务会进入审核状态,审核完成后会等待训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CommitVoiceTrainingJob(request *model.CommitVoiceTrainingJobRequest) (*model.CommitVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForCommitVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CommitVoiceTrainingJobResponse), nil
	}
}

// CommitVoiceTrainingJobInvoker 提交语音训练任务
func (c *MetaStudioClient) CommitVoiceTrainingJobInvoker(request *model.CommitVoiceTrainingJobRequest) *CommitVoiceTrainingJobInvoker {
	requestDef := GenReqDefForCommitVoiceTrainingJob()
	return &CommitVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmTrainingSegment 确认在线录音结果
//
// 确认在线录音结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ConfirmTrainingSegment(request *model.ConfirmTrainingSegmentRequest) (*model.ConfirmTrainingSegmentResponse, error) {
	requestDef := GenReqDefForConfirmTrainingSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmTrainingSegmentResponse), nil
	}
}

// ConfirmTrainingSegmentInvoker 确认在线录音结果
func (c *MetaStudioClient) ConfirmTrainingSegmentInvoker(request *model.ConfirmTrainingSegmentRequest) *ConfirmTrainingSegmentInvoker {
	requestDef := GenReqDefForConfirmTrainingSegment()
	return &ConfirmTrainingSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingAdvanceJob 创建高级版语音训练任务
//
// 用户创建语音训练基础版任务,该接口会返回一个obs上传地址，用于上传语音文件。
// 仅支持zip包方式上传语音文件：
// * 语音文件打包成zip上传：上传的训练数据为一个zip格式压缩文件,其中包含一段wav格式的长音频文件。
//
// &gt; * 文件上传后，调用“提交语音训练任务”接口，启动审核和训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTrainingAdvanceJob(request *model.CreateTrainingAdvanceJobRequest) (*model.CreateTrainingAdvanceJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingAdvanceJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingAdvanceJobResponse), nil
	}
}

// CreateTrainingAdvanceJobInvoker 创建高级版语音训练任务
func (c *MetaStudioClient) CreateTrainingAdvanceJobInvoker(request *model.CreateTrainingAdvanceJobRequest) *CreateTrainingAdvanceJobInvoker {
	requestDef := GenReqDefForCreateTrainingAdvanceJob()
	return &CreateTrainingAdvanceJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingBasicJob 创建基础版语音训练任务
//
// 用户创建语音训练基础版任务,该接口会返回一个obs上传地址，用于上传语音文件。
// 支持2种方式上传语音文件：
// * 语音文件和文本文件打包成zip上传：语音文件已经切分成20个wav文件，每个语音文件对应一个txt文本文件，所有文件打包成zip文件。语音文件命名规则：0.wav~19.wav；文本文件命名规则：0.txt~19.txt。
// * 语音文件和文本文件逐句上传：每次上传一句语料的语音文件和文本文件，再调用“确认在线录音结果”接口确认语音和文本内容是否一致。确认成功后再上传和确认下一句。
//
// &gt; * 文件上传后，调用“提交语音训练任务”接口，启动审核和训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTrainingBasicJob(request *model.CreateTrainingBasicJobRequest) (*model.CreateTrainingBasicJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingBasicJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingBasicJobResponse), nil
	}
}

// CreateTrainingBasicJobInvoker 创建基础版语音训练任务
func (c *MetaStudioClient) CreateTrainingBasicJobInvoker(request *model.CreateTrainingBasicJobRequest) *CreateTrainingBasicJobInvoker {
	requestDef := GenReqDefForCreateTrainingBasicJob()
	return &CreateTrainingBasicJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingMiddleJob 创建进阶版语音训练任务
//
// 用户创建语音训练基础版任务,该接口会返回一个obs上传地址，用于上传语音文件。
// 支持2种方式上传语音文件：
// * 语音文件和文本文件打包成zip上传：语音文件已经切分成100个wav文件，每个语音文件对应一个txt文本文件，所有文件打包成zip文件。语音文件命名规则：0.wav~99.wav；文本文件命名规则：0.txt~99.txt。
// * 语音文件和文本文件逐句上传：每次上传一句语料的语音文件和文本文件，再调用“确认在线录音结果”接口确认语音和文本内容是否一致。确认成功后再上传和确认下一句。
//
// &gt; * 文件上传后，调用“提交语音训练任务”接口，启动审核和训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTrainingMiddleJob(request *model.CreateTrainingMiddleJobRequest) (*model.CreateTrainingMiddleJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingMiddleJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingMiddleJobResponse), nil
	}
}

// CreateTrainingMiddleJobInvoker 创建进阶版语音训练任务
func (c *MetaStudioClient) CreateTrainingMiddleJobInvoker(request *model.CreateTrainingMiddleJobRequest) *CreateTrainingMiddleJobInvoker {
	requestDef := GenReqDefForCreateTrainingMiddleJob()
	return &CreateTrainingMiddleJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVoiceTrainingJob 删除语音训练任务
//
// 删除语音训练任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteVoiceTrainingJob(request *model.DeleteVoiceTrainingJobRequest) (*model.DeleteVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForDeleteVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVoiceTrainingJobResponse), nil
	}
}

// DeleteVoiceTrainingJobInvoker 删除语音训练任务
func (c *MetaStudioClient) DeleteVoiceTrainingJobInvoker(request *model.DeleteVoiceTrainingJobRequest) *DeleteVoiceTrainingJobInvoker {
	requestDef := GenReqDefForDeleteVoiceTrainingJob()
	return &DeleteVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVoiceTrainingJob 查询语音训练任务列表
//
// 查询语音训练任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListVoiceTrainingJob(request *model.ListVoiceTrainingJobRequest) (*model.ListVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForListVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVoiceTrainingJobResponse), nil
	}
}

// ListVoiceTrainingJobInvoker 查询语音训练任务列表
func (c *MetaStudioClient) ListVoiceTrainingJobInvoker(request *model.ListVoiceTrainingJobRequest) *ListVoiceTrainingJobInvoker {
	requestDef := GenReqDefForListVoiceTrainingJob()
	return &ListVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobAuditResult 获取语音训练任务审核结果
//
// 获取语音训练任务审核结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowJobAuditResult(request *model.ShowJobAuditResultRequest) (*model.ShowJobAuditResultResponse, error) {
	requestDef := GenReqDefForShowJobAuditResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobAuditResultResponse), nil
	}
}

// ShowJobAuditResultInvoker 获取语音训练任务审核结果
func (c *MetaStudioClient) ShowJobAuditResultInvoker(request *model.ShowJobAuditResultRequest) *ShowJobAuditResultInvoker {
	requestDef := GenReqDefForShowJobAuditResult()
	return &ShowJobAuditResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobUploadingAddress 获取语音文件上传地址
//
// 获取语音文件上传地址
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowJobUploadingAddress(request *model.ShowJobUploadingAddressRequest) (*model.ShowJobUploadingAddressResponse, error) {
	requestDef := GenReqDefForShowJobUploadingAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobUploadingAddressResponse), nil
	}
}

// ShowJobUploadingAddressInvoker 获取语音文件上传地址
func (c *MetaStudioClient) ShowJobUploadingAddressInvoker(request *model.ShowJobUploadingAddressRequest) *ShowJobUploadingAddressInvoker {
	requestDef := GenReqDefForShowJobUploadingAddress()
	return &ShowJobUploadingAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingSegmentInfo 获取在线录音确认结果
//
// 获取在线录音确认结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTrainingSegmentInfo(request *model.ShowTrainingSegmentInfoRequest) (*model.ShowTrainingSegmentInfoResponse, error) {
	requestDef := GenReqDefForShowTrainingSegmentInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingSegmentInfoResponse), nil
	}
}

// ShowTrainingSegmentInfoInvoker 获取在线录音确认结果
func (c *MetaStudioClient) ShowTrainingSegmentInfoInvoker(request *model.ShowTrainingSegmentInfoRequest) *ShowTrainingSegmentInfoInvoker {
	requestDef := GenReqDefForShowTrainingSegmentInfo()
	return &ShowTrainingSegmentInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVoiceTrainingJob 查询语音训练任务详情
//
// 查询语音训练任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVoiceTrainingJob(request *model.ShowVoiceTrainingJobRequest) (*model.ShowVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForShowVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVoiceTrainingJobResponse), nil
	}
}

// ShowVoiceTrainingJobInvoker 查询语音训练任务详情
func (c *MetaStudioClient) ShowVoiceTrainingJobInvoker(request *model.ShowVoiceTrainingJobRequest) *ShowVoiceTrainingJobInvoker {
	requestDef := GenReqDefForShowVoiceTrainingJob()
	return &ShowVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Create2dModelTrainingJob 创建分身数字人模型训练任务
//
// 该接口用于创建分身数字人模型训练任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Create2dModelTrainingJob(request *model.Create2dModelTrainingJobRequest) (*model.Create2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForCreate2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Create2dModelTrainingJobResponse), nil
	}
}

// Create2dModelTrainingJobInvoker 创建分身数字人模型训练任务
func (c *MetaStudioClient) Create2dModelTrainingJobInvoker(request *model.Create2dModelTrainingJobRequest) *Create2dModelTrainingJobInvoker {
	requestDef := GenReqDefForCreate2dModelTrainingJob()
	return &Create2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Delete2dModelTrainingJob 删除分身数字人模型训练任务
//
// 该接口用于删除分身数字人模型训练任务。同时需要删除训练任务相关的训练视频、身份证照片、授权文件、模型资产等。
// &gt; * 该接口应当在任务处于以下状态时调用：WAIT_FILE_UPLOAD、AUTO_VERIFY_FAILED、MANUAL_VERIFYING、MANUAL_VERIFY_FAILED、TRAINING_DATA_PREPROCESS_FAILED、TRAIN_FAILED、INFERENCE_DATA_PREPROCESS_FAILED、JOB_SUCCESS、WAIT_USER_CONFIRM、JOB_REJECT、JOB_FINISH
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Delete2dModelTrainingJob(request *model.Delete2dModelTrainingJobRequest) (*model.Delete2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForDelete2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Delete2dModelTrainingJobResponse), nil
	}
}

// Delete2dModelTrainingJobInvoker 删除分身数字人模型训练任务
func (c *MetaStudioClient) Delete2dModelTrainingJobInvoker(request *model.Delete2dModelTrainingJobRequest) *Delete2dModelTrainingJobInvoker {
	requestDef := GenReqDefForDelete2dModelTrainingJob()
	return &Delete2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Execute2dModelTrainingCommandByUser 租户执行分身数字人模型训练任务命令
//
// 该接口用于租户执行分身数字人模型训练任务命令，如提交训练审核等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Execute2dModelTrainingCommandByUser(request *model.Execute2dModelTrainingCommandByUserRequest) (*model.Execute2dModelTrainingCommandByUserResponse, error) {
	requestDef := GenReqDefForExecute2dModelTrainingCommandByUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Execute2dModelTrainingCommandByUserResponse), nil
	}
}

// Execute2dModelTrainingCommandByUserInvoker 租户执行分身数字人模型训练任务命令
func (c *MetaStudioClient) Execute2dModelTrainingCommandByUserInvoker(request *model.Execute2dModelTrainingCommandByUserRequest) *Execute2dModelTrainingCommandByUserInvoker {
	requestDef := GenReqDefForExecute2dModelTrainingCommandByUser()
	return &Execute2dModelTrainingCommandByUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// List2dModelTrainingJob 查询分身数字人模型训练任务列表
//
// 该接口用于查询分身数字人模型训练任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) List2dModelTrainingJob(request *model.List2dModelTrainingJobRequest) (*model.List2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForList2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.List2dModelTrainingJobResponse), nil
	}
}

// List2dModelTrainingJobInvoker 查询分身数字人模型训练任务列表
func (c *MetaStudioClient) List2dModelTrainingJobInvoker(request *model.List2dModelTrainingJobRequest) *List2dModelTrainingJobInvoker {
	requestDef := GenReqDefForList2dModelTrainingJob()
	return &List2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Show2dModelTrainingJob 查询分身数字人模型训练任务详情
//
// 该接口用于查询分身数字人模型训练任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Show2dModelTrainingJob(request *model.Show2dModelTrainingJobRequest) (*model.Show2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForShow2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Show2dModelTrainingJobResponse), nil
	}
}

// Show2dModelTrainingJobInvoker 查询分身数字人模型训练任务详情
func (c *MetaStudioClient) Show2dModelTrainingJobInvoker(request *model.Show2dModelTrainingJobRequest) *Show2dModelTrainingJobInvoker {
	requestDef := GenReqDefForShow2dModelTrainingJob()
	return &Show2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Update2dModelTrainingJob 更新分身数字人模型训练任务
//
// 该接口用于更新分身数字人模型训练任务。用于在自动审核或者人工审核不通过情况下，更新训练视频、身份证照片等。
// &gt; * 该接口只能在AUTO_VERIFY_FAILED或者MANUAL_VERIFY_FAILED状态下调用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Update2dModelTrainingJob(request *model.Update2dModelTrainingJobRequest) (*model.Update2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForUpdate2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Update2dModelTrainingJobResponse), nil
	}
}

// Update2dModelTrainingJobInvoker 更新分身数字人模型训练任务
func (c *MetaStudioClient) Update2dModelTrainingJobInvoker(request *model.Update2dModelTrainingJobRequest) *Update2dModelTrainingJobInvoker {
	requestDef := GenReqDefForUpdate2dModelTrainingJob()
	return &Update2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFacialAnimations 创建语音驱动表情动画任务
//
// 该接口用于创建驱动数字人表情的任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateFacialAnimations(request *model.CreateFacialAnimationsRequest) (*model.CreateFacialAnimationsResponse, error) {
	requestDef := GenReqDefForCreateFacialAnimations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFacialAnimationsResponse), nil
	}
}

// CreateFacialAnimationsInvoker 创建语音驱动表情动画任务
func (c *MetaStudioClient) CreateFacialAnimationsInvoker(request *model.CreateFacialAnimationsRequest) *CreateFacialAnimationsInvoker {
	requestDef := GenReqDefForCreateFacialAnimations()
	return &CreateFacialAnimationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTtsa 创建语音驱动任务
//
// 该接口用于创建驱动数字人表情、动作及语音的任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTtsa(request *model.CreateTtsaRequest) (*model.CreateTtsaResponse, error) {
	requestDef := GenReqDefForCreateTtsa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTtsaResponse), nil
	}
}

// CreateTtsaInvoker 创建语音驱动任务
func (c *MetaStudioClient) CreateTtsaInvoker(request *model.CreateTtsaRequest) *CreateTtsaInvoker {
	requestDef := GenReqDefForCreateTtsa()
	return &CreateTtsaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFacialAnimationsData 获取语音驱动表情数据
//
// 该接口用于获取生成的数字人表情驱动数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListFacialAnimationsData(request *model.ListFacialAnimationsDataRequest) (*model.ListFacialAnimationsDataResponse, error) {
	requestDef := GenReqDefForListFacialAnimationsData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFacialAnimationsDataResponse), nil
	}
}

// ListFacialAnimationsDataInvoker 获取语音驱动表情数据
func (c *MetaStudioClient) ListFacialAnimationsDataInvoker(request *model.ListFacialAnimationsDataRequest) *ListFacialAnimationsDataInvoker {
	requestDef := GenReqDefForListFacialAnimationsData()
	return &ListFacialAnimationsDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTtsaData 获取语音驱动数据
//
// 该接口用于获取生成的数字人驱动数据，包括语音、表情、动作等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTtsaData(request *model.ListTtsaDataRequest) (*model.ListTtsaDataResponse, error) {
	requestDef := GenReqDefForListTtsaData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTtsaDataResponse), nil
	}
}

// ListTtsaDataInvoker 获取语音驱动数据
func (c *MetaStudioClient) ListTtsaDataInvoker(request *model.ListTtsaDataRequest) *ListTtsaDataInvoker {
	requestDef := GenReqDefForListTtsaData()
	return &ListTtsaDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTtsaJobs 获取语音驱动任务列表
//
// 该接口用于查询驱动数字人表情、动作及语音的任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTtsaJobs(request *model.ListTtsaJobsRequest) (*model.ListTtsaJobsResponse, error) {
	requestDef := GenReqDefForListTtsaJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTtsaJobsResponse), nil
	}
}

// ListTtsaJobsInvoker 获取语音驱动任务列表
func (c *MetaStudioClient) ListTtsaJobsInvoker(request *model.ListTtsaJobsRequest) *ListTtsaJobsInvoker {
	requestDef := GenReqDefForListTtsaJobs()
	return &ListTtsaJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTtsAudition 创建TTS试听任务
//
// 该接口用于创建生成播报内容的语音试听文件任务。第三方音色试听需要收费，收费标准参考：https://marketplace.huaweicloud.com/product/OFFI919400645308506112#productid&#x3D;OFFI919400645308506112
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTtsAudition(request *model.CreateTtsAuditionRequest) (*model.CreateTtsAuditionResponse, error) {
	requestDef := GenReqDefForCreateTtsAudition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTtsAuditionResponse), nil
	}
}

// CreateTtsAuditionInvoker 创建TTS试听任务
func (c *MetaStudioClient) CreateTtsAuditionInvoker(request *model.CreateTtsAuditionRequest) *CreateTtsAuditionInvoker {
	requestDef := GenReqDefForCreateTtsAudition()
	return &CreateTtsAuditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTtsAuditionFile 获取TTS试听文件
//
// 该接口用于获取TTS试听文件下载链接，返回List中包含当前已生产的试听文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTtsAuditionFile(request *model.ShowTtsAuditionFileRequest) (*model.ShowTtsAuditionFileResponse, error) {
	requestDef := GenReqDefForShowTtsAuditionFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTtsAuditionFileResponse), nil
	}
}

// ShowTtsAuditionFileInvoker 获取TTS试听文件
func (c *MetaStudioClient) ShowTtsAuditionFileInvoker(request *model.ShowTtsAuditionFileRequest) *ShowTtsAuditionFileInvoker {
	requestDef := GenReqDefForShowTtsAuditionFile()
	return &ShowTtsAuditionFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoMotionCaptureJob 创建视频驱动任务
//
// 该接口用于创建视频驱动任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateVideoMotionCaptureJob(request *model.CreateVideoMotionCaptureJobRequest) (*model.CreateVideoMotionCaptureJobResponse, error) {
	requestDef := GenReqDefForCreateVideoMotionCaptureJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoMotionCaptureJobResponse), nil
	}
}

// CreateVideoMotionCaptureJobInvoker 创建视频驱动任务
func (c *MetaStudioClient) CreateVideoMotionCaptureJobInvoker(request *model.CreateVideoMotionCaptureJobRequest) *CreateVideoMotionCaptureJobInvoker {
	requestDef := GenReqDefForCreateVideoMotionCaptureJob()
	return &CreateVideoMotionCaptureJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteVideoMotionCaptureCommand 控制数字人驱动
//
// 该接口用于控制数字人驱动。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ExecuteVideoMotionCaptureCommand(request *model.ExecuteVideoMotionCaptureCommandRequest) (*model.ExecuteVideoMotionCaptureCommandResponse, error) {
	requestDef := GenReqDefForExecuteVideoMotionCaptureCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteVideoMotionCaptureCommandResponse), nil
	}
}

// ExecuteVideoMotionCaptureCommandInvoker 控制数字人驱动
func (c *MetaStudioClient) ExecuteVideoMotionCaptureCommandInvoker(request *model.ExecuteVideoMotionCaptureCommandRequest) *ExecuteVideoMotionCaptureCommandInvoker {
	requestDef := GenReqDefForExecuteVideoMotionCaptureCommand()
	return &ExecuteVideoMotionCaptureCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVideoMotionCaptureJobs 查询视频驱动任务列表
//
// 该接口用于查询视频驱动任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListVideoMotionCaptureJobs(request *model.ListVideoMotionCaptureJobsRequest) (*model.ListVideoMotionCaptureJobsResponse, error) {
	requestDef := GenReqDefForListVideoMotionCaptureJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVideoMotionCaptureJobsResponse), nil
	}
}

// ListVideoMotionCaptureJobsInvoker 查询视频驱动任务列表
func (c *MetaStudioClient) ListVideoMotionCaptureJobsInvoker(request *model.ListVideoMotionCaptureJobsRequest) *ListVideoMotionCaptureJobsInvoker {
	requestDef := GenReqDefForListVideoMotionCaptureJobs()
	return &ListVideoMotionCaptureJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoMotionCaptureJob 查询视频驱动任务详情
//
// 该接口用于查询视频驱动任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVideoMotionCaptureJob(request *model.ShowVideoMotionCaptureJobRequest) (*model.ShowVideoMotionCaptureJobResponse, error) {
	requestDef := GenReqDefForShowVideoMotionCaptureJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoMotionCaptureJobResponse), nil
	}
}

// ShowVideoMotionCaptureJobInvoker 查询视频驱动任务详情
func (c *MetaStudioClient) ShowVideoMotionCaptureJobInvoker(request *model.ShowVideoMotionCaptureJobRequest) *ShowVideoMotionCaptureJobInvoker {
	requestDef := GenReqDefForShowVideoMotionCaptureJob()
	return &ShowVideoMotionCaptureJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopVideoMotionCaptureJob 停止视频驱动任务
//
// 该接口用于停止视频驱动任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StopVideoMotionCaptureJob(request *model.StopVideoMotionCaptureJobRequest) (*model.StopVideoMotionCaptureJobResponse, error) {
	requestDef := GenReqDefForStopVideoMotionCaptureJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopVideoMotionCaptureJobResponse), nil
	}
}

// StopVideoMotionCaptureJobInvoker 停止视频驱动任务
func (c *MetaStudioClient) StopVideoMotionCaptureJobInvoker(request *model.StopVideoMotionCaptureJobRequest) *StopVideoMotionCaptureJobInvoker {
	requestDef := GenReqDefForStopVideoMotionCaptureJob()
	return &StopVideoMotionCaptureJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyVideoScripts 复制视频制作剧本
//
// 该接口用于复制视频制作剧本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CopyVideoScripts(request *model.CopyVideoScriptsRequest) (*model.CopyVideoScriptsResponse, error) {
	requestDef := GenReqDefForCopyVideoScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyVideoScriptsResponse), nil
	}
}

// CopyVideoScriptsInvoker 复制视频制作剧本
func (c *MetaStudioClient) CopyVideoScriptsInvoker(request *model.CopyVideoScriptsRequest) *CopyVideoScriptsInvoker {
	requestDef := GenReqDefForCopyVideoScripts()
	return &CopyVideoScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoScripts 创建视频制作剧本
//
// 该接口用于创建视频制作剧本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateVideoScripts(request *model.CreateVideoScriptsRequest) (*model.CreateVideoScriptsResponse, error) {
	requestDef := GenReqDefForCreateVideoScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoScriptsResponse), nil
	}
}

// CreateVideoScriptsInvoker 创建视频制作剧本
func (c *MetaStudioClient) CreateVideoScriptsInvoker(request *model.CreateVideoScriptsRequest) *CreateVideoScriptsInvoker {
	requestDef := GenReqDefForCreateVideoScripts()
	return &CreateVideoScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVideoScript 删除视频制作剧本
//
// 该接口用于删除视频制作剧本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteVideoScript(request *model.DeleteVideoScriptRequest) (*model.DeleteVideoScriptResponse, error) {
	requestDef := GenReqDefForDeleteVideoScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVideoScriptResponse), nil
	}
}

// DeleteVideoScriptInvoker 删除视频制作剧本
func (c *MetaStudioClient) DeleteVideoScriptInvoker(request *model.DeleteVideoScriptRequest) *DeleteVideoScriptInvoker {
	requestDef := GenReqDefForDeleteVideoScript()
	return &DeleteVideoScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVideoScripts 查询视频制作剧本列表
//
// 该接口用于查询视频制作剧本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListVideoScripts(request *model.ListVideoScriptsRequest) (*model.ListVideoScriptsResponse, error) {
	requestDef := GenReqDefForListVideoScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVideoScriptsResponse), nil
	}
}

// ListVideoScriptsInvoker 查询视频制作剧本列表
func (c *MetaStudioClient) ListVideoScriptsInvoker(request *model.ListVideoScriptsRequest) *ListVideoScriptsInvoker {
	requestDef := GenReqDefForListVideoScripts()
	return &ListVideoScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoScript 查询视频制作剧本详情
//
// 该接口用于查询视频制作剧本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVideoScript(request *model.ShowVideoScriptRequest) (*model.ShowVideoScriptResponse, error) {
	requestDef := GenReqDefForShowVideoScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoScriptResponse), nil
	}
}

// ShowVideoScriptInvoker 查询视频制作剧本详情
func (c *MetaStudioClient) ShowVideoScriptInvoker(request *model.ShowVideoScriptRequest) *ShowVideoScriptInvoker {
	requestDef := GenReqDefForShowVideoScript()
	return &ShowVideoScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVideoScript 更新视频制作剧本
//
// 该接口用于更新视频制作剧本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateVideoScript(request *model.UpdateVideoScriptRequest) (*model.UpdateVideoScriptResponse, error) {
	requestDef := GenReqDefForUpdateVideoScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVideoScriptResponse), nil
	}
}

// UpdateVideoScriptInvoker 更新视频制作剧本
func (c *MetaStudioClient) UpdateVideoScriptInvoker(request *model.UpdateVideoScriptRequest) *UpdateVideoScriptInvoker {
	requestDef := GenReqDefForUpdateVideoScript()
	return &UpdateVideoScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
