// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDocPath struct {
	position.Path
	// 文档 Id
	DocId *string `json:"docId,omitempty"`
}

func (s DeleteDocPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteDocPath) GoString() string {
	return s.String()
}

func (s DeleteDocPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDocPath) SetDocId(v string) *DeleteDocPath {
	s.DocId = &v
	return s
}

type DeleteDocPathBuilder struct {
	s *DeleteDocPath
}

func NewDeleteDocPathBuilder() *DeleteDocPathBuilder {
	s := &DeleteDocPath{}
	b := &DeleteDocPathBuilder{s: s}
	return b
}

func (b *DeleteDocPathBuilder) DocId(v string) *DeleteDocPathBuilder {
	b.s.DocId = &v
	return b
}

func (b *DeleteDocPathBuilder) Build() *DeleteDocPath {
	return b.s
}
