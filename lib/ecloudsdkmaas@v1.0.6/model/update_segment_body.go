// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSegmentBody struct {
	position.Body
	// QA问答对的回答
	Answer *string `json:"answer,omitempty"`
	// 分段文档内容，或者QA问答对的提问
	Content *string `json:"content,omitempty"`
}

func (s UpdateSegmentBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateSegmentBody) GoString() string {
	return s.String()
}

func (s UpdateSegmentBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSegmentBody) SetAnswer(v string) *UpdateSegmentBody {
	s.Answer = &v
	return s
}

func (s *UpdateSegmentBody) SetContent(v string) *UpdateSegmentBody {
	s.Content = &v
	return s
}

type UpdateSegmentBodyBuilder struct {
	s *UpdateSegmentBody
}

func NewUpdateSegmentBodyBuilder() *UpdateSegmentBodyBuilder {
	s := &UpdateSegmentBody{}
	b := &UpdateSegmentBodyBuilder{s: s}
	return b
}

func (b *UpdateSegmentBodyBuilder) Answer(v string) *UpdateSegmentBodyBuilder {
	b.s.Answer = &v
	return b
}

func (b *UpdateSegmentBodyBuilder) Content(v string) *UpdateSegmentBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *UpdateSegmentBodyBuilder) Build() *UpdateSegmentBody {
	return b.s
}
