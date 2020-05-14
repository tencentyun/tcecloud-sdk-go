package errors

import (
	"fmt"
)

type TceCloudSDKError struct {
	Code      string
	Message   string
	RequestId string
}

func (e *TceCloudSDKError) Error() string {
	return fmt.Sprintf("[TceCloudSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewTceCloudSDKError(code, message, requestId string) error {
	return &TceCloudSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *TceCloudSDKError) GetCode() string {
	return e.Code
}

func (e *TceCloudSDKError) GetMessage() string {
	return e.Message
}

func (e *TceCloudSDKError) GetRequestId() string {
	return e.RequestId
}
