// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteAssetRequest struct {
	DeleteAssetPath *DeleteAssetPath `json:"deleteAssetPath,omitempty"`
}

func (s DeleteAssetRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteAssetRequest) GoString() string {
	return s.String()
}

func (s DeleteAssetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteAssetRequest) SetDeleteAssetPath(v *DeleteAssetPath) *DeleteAssetRequest {
	s.DeleteAssetPath = v
	return s
}

type DeleteAssetRequestBuilder struct {
	s *DeleteAssetRequest
}

func NewDeleteAssetRequestBuilder() *DeleteAssetRequestBuilder {
	s := &DeleteAssetRequest{}
	b := &DeleteAssetRequestBuilder{s: s}
	return b
}

func (b *DeleteAssetRequestBuilder) DeleteAssetPath(v *DeleteAssetPath) *DeleteAssetRequestBuilder {
	b.s.DeleteAssetPath = v
	return b
}

func (b *DeleteAssetRequestBuilder) Build() *DeleteAssetRequest {
	return b.s
}
