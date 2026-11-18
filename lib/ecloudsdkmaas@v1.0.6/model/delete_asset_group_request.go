// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteAssetGroupRequest struct {
	DeleteAssetGroupPath *DeleteAssetGroupPath `json:"deleteAssetGroupPath,omitempty"`
}

func (s DeleteAssetGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteAssetGroupRequest) GoString() string {
	return s.String()
}

func (s DeleteAssetGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteAssetGroupRequest) SetDeleteAssetGroupPath(v *DeleteAssetGroupPath) *DeleteAssetGroupRequest {
	s.DeleteAssetGroupPath = v
	return s
}

type DeleteAssetGroupRequestBuilder struct {
	s *DeleteAssetGroupRequest
}

func NewDeleteAssetGroupRequestBuilder() *DeleteAssetGroupRequestBuilder {
	s := &DeleteAssetGroupRequest{}
	b := &DeleteAssetGroupRequestBuilder{s: s}
	return b
}

func (b *DeleteAssetGroupRequestBuilder) DeleteAssetGroupPath(v *DeleteAssetGroupPath) *DeleteAssetGroupRequestBuilder {
	b.s.DeleteAssetGroupPath = v
	return b
}

func (b *DeleteAssetGroupRequestBuilder) Build() *DeleteAssetGroupRequest {
	return b.s
}
