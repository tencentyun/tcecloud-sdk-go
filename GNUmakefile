PKG_NAME=tcecloud

tce:
	find ./$(PKG_NAME) -type f|xargs sed -i 's#github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud#github.com/tencentyun/tcecloud-sdk-go/tcecloud#g'
	find ./$(PKG_NAME) -type f|xargs sed -i 's#TencentCloudSDKError#TceCloudSDKError#g'
