// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Children struct {
	AreaCode         *string     `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	CategoryId       *int64      `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Childrens        []*Children `json:"Childrens,omitempty" xml:"Childrens,omitempty" type:"Repeated"`
	Name             *string     `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCategoryId *int64      `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s Children) String() string {
	return tea.Prettify(s)
}

func (s Children) GoString() string {
	return s.String()
}

func (s *Children) SetAreaCode(v string) *Children {
	s.AreaCode = &v
	return s
}

func (s *Children) SetCategoryId(v int64) *Children {
	s.CategoryId = &v
	return s
}

func (s *Children) SetChildrens(v []*Children) *Children {
	s.Childrens = v
	return s
}

func (s *Children) SetName(v string) *Children {
	s.Name = &v
	return s
}

func (s *Children) SetParentCategoryId(v int64) *Children {
	s.ParentCategoryId = &v
	return s
}

type IntentCreateDTO struct {
	IntentId  *int64               `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	Name      *string              `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCheck []*RuleMtopDTO       `json:"RuleCheck,omitempty" xml:"RuleCheck,omitempty" type:"Repeated"`
	Slot      []*SlotrecordMtopDTO `json:"Slot,omitempty" xml:"Slot,omitempty" type:"Repeated"`
	UserSay   []*UsersayMtopDTO    `json:"UserSay,omitempty" xml:"UserSay,omitempty" type:"Repeated"`
}

func (s IntentCreateDTO) String() string {
	return tea.Prettify(s)
}

func (s IntentCreateDTO) GoString() string {
	return s.String()
}

func (s *IntentCreateDTO) SetIntentId(v int64) *IntentCreateDTO {
	s.IntentId = &v
	return s
}

func (s *IntentCreateDTO) SetName(v string) *IntentCreateDTO {
	s.Name = &v
	return s
}

func (s *IntentCreateDTO) SetRuleCheck(v []*RuleMtopDTO) *IntentCreateDTO {
	s.RuleCheck = v
	return s
}

func (s *IntentCreateDTO) SetSlot(v []*SlotrecordMtopDTO) *IntentCreateDTO {
	s.Slot = v
	return s
}

func (s *IntentCreateDTO) SetUserSay(v []*UsersayMtopDTO) *IntentCreateDTO {
	s.UserSay = v
	return s
}

type PaasButtonDTO struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PaasButtonDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasButtonDTO) GoString() string {
	return s.String()
}

func (s *PaasButtonDTO) SetName(v string) *PaasButtonDTO {
	s.Name = &v
	return s
}

func (s *PaasButtonDTO) SetText(v string) *PaasButtonDTO {
	s.Text = &v
	return s
}

func (s *PaasButtonDTO) SetType(v string) *PaasButtonDTO {
	s.Type = &v
	return s
}

type PaasButtonListDTO struct {
	Button []*PaasButtonDTO `json:"Button,omitempty" xml:"Button,omitempty" type:"Repeated"`
	Intro  *string          `json:"Intro,omitempty" xml:"Intro,omitempty"`
}

func (s PaasButtonListDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasButtonListDTO) GoString() string {
	return s.String()
}

func (s *PaasButtonListDTO) SetButton(v []*PaasButtonDTO) *PaasButtonListDTO {
	s.Button = v
	return s
}

func (s *PaasButtonListDTO) SetIntro(v string) *PaasButtonListDTO {
	s.Intro = &v
	return s
}

type PaasConditionEntryDTO struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Term  *string `json:"Term,omitempty" xml:"Term,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PaasConditionEntryDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasConditionEntryDTO) GoString() string {
	return s.String()
}

func (s *PaasConditionEntryDTO) SetId(v string) *PaasConditionEntryDTO {
	s.Id = &v
	return s
}

func (s *PaasConditionEntryDTO) SetName(v string) *PaasConditionEntryDTO {
	s.Name = &v
	return s
}

func (s *PaasConditionEntryDTO) SetTerm(v string) *PaasConditionEntryDTO {
	s.Term = &v
	return s
}

func (s *PaasConditionEntryDTO) SetType(v string) *PaasConditionEntryDTO {
	s.Type = &v
	return s
}

func (s *PaasConditionEntryDTO) SetValue(v string) *PaasConditionEntryDTO {
	s.Value = &v
	return s
}

type PaasConditionSetDTO struct {
	ConditionEntries []*PaasConditionEntryDTO `json:"ConditionEntries,omitempty" xml:"ConditionEntries,omitempty" type:"Repeated"`
}

func (s PaasConditionSetDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasConditionSetDTO) GoString() string {
	return s.String()
}

func (s *PaasConditionSetDTO) SetConditionEntries(v []*PaasConditionEntryDTO) *PaasConditionSetDTO {
	s.ConditionEntries = v
	return s
}

type PaasEdgeDTO struct {
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Label  *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s PaasEdgeDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasEdgeDTO) GoString() string {
	return s.String()
}

func (s *PaasEdgeDTO) SetId(v string) *PaasEdgeDTO {
	s.Id = &v
	return s
}

func (s *PaasEdgeDTO) SetLabel(v string) *PaasEdgeDTO {
	s.Label = &v
	return s
}

func (s *PaasEdgeDTO) SetSource(v string) *PaasEdgeDTO {
	s.Source = &v
	return s
}

func (s *PaasEdgeDTO) SetTarget(v string) *PaasEdgeDTO {
	s.Target = &v
	return s
}

type PaasEntryDTO struct {
	PluginFieldDataEntry *PaasEntryPluginFieldDataDTO `json:"PluginFieldDataEntry,omitempty" xml:"PluginFieldDataEntry,omitempty"`
}

func (s PaasEntryDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasEntryDTO) GoString() string {
	return s.String()
}

func (s *PaasEntryDTO) SetPluginFieldDataEntry(v *PaasEntryPluginFieldDataDTO) *PaasEntryDTO {
	s.PluginFieldDataEntry = v
	return s
}

type PaasEntryPluginFieldDataDTO struct {
	ContentEntry []*PaasConditionSetDTO `json:"ContentEntry,omitempty" xml:"ContentEntry,omitempty" type:"Repeated"`
	LifeSpan     *int64                 `json:"LifeSpan,omitempty" xml:"LifeSpan,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PaasEntryPluginFieldDataDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasEntryPluginFieldDataDTO) GoString() string {
	return s.String()
}

func (s *PaasEntryPluginFieldDataDTO) SetContentEntry(v []*PaasConditionSetDTO) *PaasEntryPluginFieldDataDTO {
	s.ContentEntry = v
	return s
}

func (s *PaasEntryPluginFieldDataDTO) SetLifeSpan(v int64) *PaasEntryPluginFieldDataDTO {
	s.LifeSpan = &v
	return s
}

func (s *PaasEntryPluginFieldDataDTO) SetName(v string) *PaasEntryPluginFieldDataDTO {
	s.Name = &v
	return s
}

type PaasFunctionDTO struct {
	PluginFieldDataFunction *PaasFunctionPluginFieldDataDTO `json:"PluginFieldDataFunction,omitempty" xml:"PluginFieldDataFunction,omitempty"`
}

func (s PaasFunctionDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasFunctionDTO) GoString() string {
	return s.String()
}

func (s *PaasFunctionDTO) SetPluginFieldDataFunction(v *PaasFunctionPluginFieldDataDTO) *PaasFunctionDTO {
	s.PluginFieldDataFunction = v
	return s
}

type PaasFunctionPluginFieldDataDTO struct {
	AliyunFunction *string                `json:"AliyunFunction,omitempty" xml:"AliyunFunction,omitempty"`
	AliyunService  *string                `json:"AliyunService,omitempty" xml:"AliyunService,omitempty"`
	Code           *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Description    *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	EndPoint       *string                `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	Function       *string                `json:"Function,omitempty" xml:"Function,omitempty"`
	Name           *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Params         map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	Switch         []*PaasSwitchCaseDTO   `json:"Switch,omitempty" xml:"Switch,omitempty" type:"Repeated"`
	Type           *string                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PaasFunctionPluginFieldDataDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasFunctionPluginFieldDataDTO) GoString() string {
	return s.String()
}

func (s *PaasFunctionPluginFieldDataDTO) SetAliyunFunction(v string) *PaasFunctionPluginFieldDataDTO {
	s.AliyunFunction = &v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetAliyunService(v string) *PaasFunctionPluginFieldDataDTO {
	s.AliyunService = &v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetCode(v string) *PaasFunctionPluginFieldDataDTO {
	s.Code = &v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetDescription(v string) *PaasFunctionPluginFieldDataDTO {
	s.Description = &v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetEndPoint(v string) *PaasFunctionPluginFieldDataDTO {
	s.EndPoint = &v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetFunction(v string) *PaasFunctionPluginFieldDataDTO {
	s.Function = &v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetName(v string) *PaasFunctionPluginFieldDataDTO {
	s.Name = &v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetParams(v map[string]interface{}) *PaasFunctionPluginFieldDataDTO {
	s.Params = v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetSwitch(v []*PaasSwitchCaseDTO) *PaasFunctionPluginFieldDataDTO {
	s.Switch = v
	return s
}

func (s *PaasFunctionPluginFieldDataDTO) SetType(v string) *PaasFunctionPluginFieldDataDTO {
	s.Type = &v
	return s
}

type PaasFunctionPluginParams struct {
	Body   *string            `json:"Body,omitempty" xml:"Body,omitempty"`
	Header map[string]*string `json:"Header,omitempty" xml:"Header,omitempty"`
	Method *string            `json:"Method,omitempty" xml:"Method,omitempty"`
	Query  map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
	Url    *string            `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PaasFunctionPluginParams) String() string {
	return tea.Prettify(s)
}

func (s PaasFunctionPluginParams) GoString() string {
	return s.String()
}

func (s *PaasFunctionPluginParams) SetBody(v string) *PaasFunctionPluginParams {
	s.Body = &v
	return s
}

func (s *PaasFunctionPluginParams) SetHeader(v map[string]*string) *PaasFunctionPluginParams {
	s.Header = v
	return s
}

func (s *PaasFunctionPluginParams) SetMethod(v string) *PaasFunctionPluginParams {
	s.Method = &v
	return s
}

func (s *PaasFunctionPluginParams) SetQuery(v map[string]*string) *PaasFunctionPluginParams {
	s.Query = v
	return s
}

func (s *PaasFunctionPluginParams) SetUrl(v string) *PaasFunctionPluginParams {
	s.Url = &v
	return s
}

type PaasNodeDTO struct {
	Code       *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	Id         *string            `json:"Id,omitempty" xml:"Id,omitempty"`
	Label      *string            `json:"Label,omitempty" xml:"Label,omitempty"`
	PluginData *PaasPluginDataDTO `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	Xx         *float64           `json:"Xx,omitempty" xml:"Xx,omitempty"`
	Yy         *float64           `json:"Yy,omitempty" xml:"Yy,omitempty"`
}

func (s PaasNodeDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasNodeDTO) GoString() string {
	return s.String()
}

func (s *PaasNodeDTO) SetCode(v string) *PaasNodeDTO {
	s.Code = &v
	return s
}

func (s *PaasNodeDTO) SetId(v string) *PaasNodeDTO {
	s.Id = &v
	return s
}

func (s *PaasNodeDTO) SetLabel(v string) *PaasNodeDTO {
	s.Label = &v
	return s
}

func (s *PaasNodeDTO) SetPluginData(v *PaasPluginDataDTO) *PaasNodeDTO {
	s.PluginData = v
	return s
}

func (s *PaasNodeDTO) SetXx(v float64) *PaasNodeDTO {
	s.Xx = &v
	return s
}

func (s *PaasNodeDTO) SetYy(v float64) *PaasNodeDTO {
	s.Yy = &v
	return s
}

type PaasPluginDataDTO struct {
	Entry    *PaasEntryDTO    `json:"Entry,omitempty" xml:"Entry,omitempty"`
	Function *PaasFunctionDTO `json:"Function,omitempty" xml:"Function,omitempty"`
	Response *PaasResponseDTO `json:"Response,omitempty" xml:"Response,omitempty"`
	Slot     *PaasSlotDTO     `json:"Slot,omitempty" xml:"Slot,omitempty"`
}

func (s PaasPluginDataDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasPluginDataDTO) GoString() string {
	return s.String()
}

func (s *PaasPluginDataDTO) SetEntry(v *PaasEntryDTO) *PaasPluginDataDTO {
	s.Entry = v
	return s
}

func (s *PaasPluginDataDTO) SetFunction(v *PaasFunctionDTO) *PaasPluginDataDTO {
	s.Function = v
	return s
}

func (s *PaasPluginDataDTO) SetResponse(v *PaasResponseDTO) *PaasPluginDataDTO {
	s.Response = v
	return s
}

func (s *PaasPluginDataDTO) SetSlot(v *PaasSlotDTO) *PaasPluginDataDTO {
	s.Slot = v
	return s
}

type PaasProcessData struct {
	Edges []*PaasEdgeDTO `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	Nodes []*PaasNodeDTO `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s PaasProcessData) String() string {
	return tea.Prettify(s)
}

func (s PaasProcessData) GoString() string {
	return s.String()
}

func (s *PaasProcessData) SetEdges(v []*PaasEdgeDTO) *PaasProcessData {
	s.Edges = v
	return s
}

func (s *PaasProcessData) SetNodes(v []*PaasNodeDTO) *PaasProcessData {
	s.Nodes = v
	return s
}

type PaasResponseDTO struct {
	PluginFieldDataResponse *PaasResponsePluginFieldDataDTO `json:"PluginFieldDataResponse,omitempty" xml:"PluginFieldDataResponse,omitempty"`
}

func (s PaasResponseDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasResponseDTO) GoString() string {
	return s.String()
}

func (s *PaasResponseDTO) SetPluginFieldDataResponse(v *PaasResponsePluginFieldDataDTO) *PaasResponseDTO {
	s.PluginFieldDataResponse = v
	return s
}

type PaasResponseNodeContentDTO struct {
	ButtonList *PaasButtonListDTO `json:"ButtonList,omitempty" xml:"ButtonList,omitempty"`
	Image      *string            `json:"Image,omitempty" xml:"Image,omitempty"`
	Text       *string            `json:"Text,omitempty" xml:"Text,omitempty"`
	Type       *string            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PaasResponseNodeContentDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasResponseNodeContentDTO) GoString() string {
	return s.String()
}

func (s *PaasResponseNodeContentDTO) SetButtonList(v *PaasButtonListDTO) *PaasResponseNodeContentDTO {
	s.ButtonList = v
	return s
}

func (s *PaasResponseNodeContentDTO) SetImage(v string) *PaasResponseNodeContentDTO {
	s.Image = &v
	return s
}

func (s *PaasResponseNodeContentDTO) SetText(v string) *PaasResponseNodeContentDTO {
	s.Text = &v
	return s
}

func (s *PaasResponseNodeContentDTO) SetType(v string) *PaasResponseNodeContentDTO {
	s.Type = &v
	return s
}

type PaasResponsePluginFieldDataDTO struct {
	ContentResponse *PaasResponseNodeContentDTO `json:"ContentResponse,omitempty" xml:"ContentResponse,omitempty"`
	Name            *string                     `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PaasResponsePluginFieldDataDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasResponsePluginFieldDataDTO) GoString() string {
	return s.String()
}

func (s *PaasResponsePluginFieldDataDTO) SetContentResponse(v *PaasResponseNodeContentDTO) *PaasResponsePluginFieldDataDTO {
	s.ContentResponse = v
	return s
}

func (s *PaasResponsePluginFieldDataDTO) SetName(v string) *PaasResponsePluginFieldDataDTO {
	s.Name = &v
	return s
}

type PaasSlotConfigDTO struct {
	IsArray     *bool     `json:"IsArray,omitempty" xml:"IsArray,omitempty"`
	IsNecessary *bool     `json:"IsNecessary,omitempty" xml:"IsNecessary,omitempty"`
	LifeSpan    *int32    `json:"LifeSpan,omitempty" xml:"LifeSpan,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Question    []*string `json:"Question,omitempty" xml:"Question,omitempty" type:"Repeated"`
	Value       *string   `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PaasSlotConfigDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasSlotConfigDTO) GoString() string {
	return s.String()
}

func (s *PaasSlotConfigDTO) SetIsArray(v bool) *PaasSlotConfigDTO {
	s.IsArray = &v
	return s
}

func (s *PaasSlotConfigDTO) SetIsNecessary(v bool) *PaasSlotConfigDTO {
	s.IsNecessary = &v
	return s
}

func (s *PaasSlotConfigDTO) SetLifeSpan(v int32) *PaasSlotConfigDTO {
	s.LifeSpan = &v
	return s
}

func (s *PaasSlotConfigDTO) SetName(v string) *PaasSlotConfigDTO {
	s.Name = &v
	return s
}

func (s *PaasSlotConfigDTO) SetQuestion(v []*string) *PaasSlotConfigDTO {
	s.Question = v
	return s
}

func (s *PaasSlotConfigDTO) SetValue(v string) *PaasSlotConfigDTO {
	s.Value = &v
	return s
}

type PaasSlotDTO struct {
	PluginFieldDataSlot *PaasSlotPluginFieldDataDTO `json:"PluginFieldDataSlot,omitempty" xml:"PluginFieldDataSlot,omitempty"`
}

func (s PaasSlotDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasSlotDTO) GoString() string {
	return s.String()
}

func (s *PaasSlotDTO) SetPluginFieldDataSlot(v *PaasSlotPluginFieldDataDTO) *PaasSlotDTO {
	s.PluginFieldDataSlot = v
	return s
}

type PaasSlotPluginFieldDataDTO struct {
	ContentSlot []*PaasSlotConfigDTO `json:"ContentSlot,omitempty" xml:"ContentSlot,omitempty" type:"Repeated"`
	IntentId    *string              `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	IntentName  *string              `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	IsSysIntent *bool                `json:"IsSysIntent,omitempty" xml:"IsSysIntent,omitempty"`
	Name        *string              `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PaasSlotPluginFieldDataDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasSlotPluginFieldDataDTO) GoString() string {
	return s.String()
}

func (s *PaasSlotPluginFieldDataDTO) SetContentSlot(v []*PaasSlotConfigDTO) *PaasSlotPluginFieldDataDTO {
	s.ContentSlot = v
	return s
}

func (s *PaasSlotPluginFieldDataDTO) SetIntentId(v string) *PaasSlotPluginFieldDataDTO {
	s.IntentId = &v
	return s
}

func (s *PaasSlotPluginFieldDataDTO) SetIntentName(v string) *PaasSlotPluginFieldDataDTO {
	s.IntentName = &v
	return s
}

func (s *PaasSlotPluginFieldDataDTO) SetIsSysIntent(v bool) *PaasSlotPluginFieldDataDTO {
	s.IsSysIntent = &v
	return s
}

func (s *PaasSlotPluginFieldDataDTO) SetName(v string) *PaasSlotPluginFieldDataDTO {
	s.Name = &v
	return s
}

type PaasSwitchCaseDTO struct {
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Label        *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
}

func (s PaasSwitchCaseDTO) String() string {
	return tea.Prettify(s)
}

func (s PaasSwitchCaseDTO) GoString() string {
	return s.String()
}

func (s *PaasSwitchCaseDTO) SetId(v string) *PaasSwitchCaseDTO {
	s.Id = &v
	return s
}

func (s *PaasSwitchCaseDTO) SetLabel(v string) *PaasSwitchCaseDTO {
	s.Label = &v
	return s
}

func (s *PaasSwitchCaseDTO) SetType(v string) *PaasSwitchCaseDTO {
	s.Type = &v
	return s
}

func (s *PaasSwitchCaseDTO) SetValue(v string) *PaasSwitchCaseDTO {
	s.Value = &v
	return s
}

func (s *PaasSwitchCaseDTO) SetVariableName(v string) *PaasSwitchCaseDTO {
	s.VariableName = &v
	return s
}

type RuleMtopDTO struct {
	Error   []*string `json:"Error,omitempty" xml:"Error,omitempty" type:"Repeated"`
	Strict  *bool     `json:"Strict,omitempty" xml:"Strict,omitempty"`
	Text    *string   `json:"Text,omitempty" xml:"Text,omitempty"`
	Warning []*string `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s RuleMtopDTO) String() string {
	return tea.Prettify(s)
}

func (s RuleMtopDTO) GoString() string {
	return s.String()
}

func (s *RuleMtopDTO) SetError(v []*string) *RuleMtopDTO {
	s.Error = v
	return s
}

func (s *RuleMtopDTO) SetStrict(v bool) *RuleMtopDTO {
	s.Strict = &v
	return s
}

func (s *RuleMtopDTO) SetText(v string) *RuleMtopDTO {
	s.Text = &v
	return s
}

func (s *RuleMtopDTO) SetWarning(v []*string) *RuleMtopDTO {
	s.Warning = v
	return s
}

type SectionMtopDTO struct {
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Text   *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SectionMtopDTO) String() string {
	return tea.Prettify(s)
}

func (s SectionMtopDTO) GoString() string {
	return s.String()
}

func (s *SectionMtopDTO) SetSlotId(v string) *SectionMtopDTO {
	s.SlotId = &v
	return s
}

func (s *SectionMtopDTO) SetText(v string) *SectionMtopDTO {
	s.Text = &v
	return s
}

type SlotrecordMtopDTO struct {
	Id          *string       `json:"Id,omitempty" xml:"Id,omitempty"`
	IsArray     *bool         `json:"IsArray,omitempty" xml:"IsArray,omitempty"`
	IsNecessary *bool         `json:"IsNecessary,omitempty" xml:"IsNecessary,omitempty"`
	LifeSpan    *int32        `json:"LifeSpan,omitempty" xml:"LifeSpan,omitempty"`
	Name        *string       `json:"Name,omitempty" xml:"Name,omitempty"`
	Question    []*string     `json:"Question,omitempty" xml:"Question,omitempty" type:"Repeated"`
	Tags        []*TagMtopDTO `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Value       *string       `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SlotrecordMtopDTO) String() string {
	return tea.Prettify(s)
}

func (s SlotrecordMtopDTO) GoString() string {
	return s.String()
}

func (s *SlotrecordMtopDTO) SetId(v string) *SlotrecordMtopDTO {
	s.Id = &v
	return s
}

func (s *SlotrecordMtopDTO) SetIsArray(v bool) *SlotrecordMtopDTO {
	s.IsArray = &v
	return s
}

func (s *SlotrecordMtopDTO) SetIsNecessary(v bool) *SlotrecordMtopDTO {
	s.IsNecessary = &v
	return s
}

func (s *SlotrecordMtopDTO) SetLifeSpan(v int32) *SlotrecordMtopDTO {
	s.LifeSpan = &v
	return s
}

func (s *SlotrecordMtopDTO) SetName(v string) *SlotrecordMtopDTO {
	s.Name = &v
	return s
}

func (s *SlotrecordMtopDTO) SetQuestion(v []*string) *SlotrecordMtopDTO {
	s.Question = v
	return s
}

func (s *SlotrecordMtopDTO) SetTags(v []*TagMtopDTO) *SlotrecordMtopDTO {
	s.Tags = v
	return s
}

func (s *SlotrecordMtopDTO) SetValue(v string) *SlotrecordMtopDTO {
	s.Value = &v
	return s
}

type TagMtopDTO struct {
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagMtopDTO) String() string {
	return tea.Prettify(s)
}

func (s TagMtopDTO) GoString() string {
	return s.String()
}

func (s *TagMtopDTO) SetUserSayId(v string) *TagMtopDTO {
	s.UserSayId = &v
	return s
}

func (s *TagMtopDTO) SetValue(v string) *TagMtopDTO {
	s.Value = &v
	return s
}

type UpdateDialogFlowModuleDefinition struct {
	Edges []*PaasEdgeDTO `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	Nodes []*PaasNodeDTO `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s UpdateDialogFlowModuleDefinition) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowModuleDefinition) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowModuleDefinition) SetEdges(v []*PaasEdgeDTO) *UpdateDialogFlowModuleDefinition {
	s.Edges = v
	return s
}

func (s *UpdateDialogFlowModuleDefinition) SetNodes(v []*PaasNodeDTO) *UpdateDialogFlowModuleDefinition {
	s.Nodes = v
	return s
}

type UsersayMtopDTO struct {
	Data   []*SectionMtopDTO `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Id     *string           `json:"Id,omitempty" xml:"Id,omitempty"`
	Strict *bool             `json:"Strict,omitempty" xml:"Strict,omitempty"`
}

func (s UsersayMtopDTO) String() string {
	return tea.Prettify(s)
}

func (s UsersayMtopDTO) GoString() string {
	return s.String()
}

func (s *UsersayMtopDTO) SetData(v []*SectionMtopDTO) *UsersayMtopDTO {
	s.Data = v
	return s
}

func (s *UsersayMtopDTO) SetId(v string) *UsersayMtopDTO {
	s.Id = &v
	return s
}

func (s *UsersayMtopDTO) SetStrict(v bool) *UsersayMtopDTO {
	s.Strict = &v
	return s
}

type ActivatePerspectiveRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s ActivatePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *ActivatePerspectiveRequest) SetAgentKey(v string) *ActivatePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *ActivatePerspectiveRequest) SetPerspectiveId(v string) *ActivatePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type ActivatePerspectiveResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivatePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *ActivatePerspectiveResponseBody) SetRequestId(v string) *ActivatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

type ActivatePerspectiveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivatePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *ActivatePerspectiveResponse) SetHeaders(v map[string]*string) *ActivatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *ActivatePerspectiveResponse) SetStatusCode(v int32) *ActivatePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivatePerspectiveResponse) SetBody(v *ActivatePerspectiveResponseBody) *ActivatePerspectiveResponse {
	s.Body = v
	return s
}

type AddSynonymRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	Synonym      *string `json:"Synonym,omitempty" xml:"Synonym,omitempty"`
}

func (s AddSynonymRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSynonymRequest) GoString() string {
	return s.String()
}

func (s *AddSynonymRequest) SetAgentKey(v string) *AddSynonymRequest {
	s.AgentKey = &v
	return s
}

func (s *AddSynonymRequest) SetCoreWordName(v string) *AddSynonymRequest {
	s.CoreWordName = &v
	return s
}

func (s *AddSynonymRequest) SetSynonym(v string) *AddSynonymRequest {
	s.Synonym = &v
	return s
}

type AddSynonymResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSynonymResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSynonymResponseBody) GoString() string {
	return s.String()
}

func (s *AddSynonymResponseBody) SetRequestId(v string) *AddSynonymResponseBody {
	s.RequestId = &v
	return s
}

type AddSynonymResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSynonymResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSynonymResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSynonymResponse) GoString() string {
	return s.String()
}

func (s *AddSynonymResponse) SetHeaders(v map[string]*string) *AddSynonymResponse {
	s.Headers = v
	return s
}

func (s *AddSynonymResponse) SetStatusCode(v int32) *AddSynonymResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSynonymResponse) SetBody(v *AddSynonymResponseBody) *AddSynonymResponse {
	s.Body = v
	return s
}

type AppendEntityMemberRequest struct {
	AgentKey  *string                          `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ApplyType *string                          `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	EntityId  *int64                           `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Member    *AppendEntityMemberRequestMember `json:"Member,omitempty" xml:"Member,omitempty" type:"Struct"`
}

func (s AppendEntityMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberRequest) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberRequest) SetAgentKey(v string) *AppendEntityMemberRequest {
	s.AgentKey = &v
	return s
}

func (s *AppendEntityMemberRequest) SetApplyType(v string) *AppendEntityMemberRequest {
	s.ApplyType = &v
	return s
}

func (s *AppendEntityMemberRequest) SetEntityId(v int64) *AppendEntityMemberRequest {
	s.EntityId = &v
	return s
}

func (s *AppendEntityMemberRequest) SetMember(v *AppendEntityMemberRequestMember) *AppendEntityMemberRequest {
	s.Member = v
	return s
}

type AppendEntityMemberRequestMember struct {
	MemberName *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s AppendEntityMemberRequestMember) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberRequestMember) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberRequestMember) SetMemberName(v string) *AppendEntityMemberRequestMember {
	s.MemberName = &v
	return s
}

func (s *AppendEntityMemberRequestMember) SetSynonyms(v []*string) *AppendEntityMemberRequestMember {
	s.Synonyms = v
	return s
}

type AppendEntityMemberShrinkRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ApplyType    *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	EntityId     *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	MemberShrink *string `json:"Member,omitempty" xml:"Member,omitempty"`
}

func (s AppendEntityMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberShrinkRequest) SetAgentKey(v string) *AppendEntityMemberShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *AppendEntityMemberShrinkRequest) SetApplyType(v string) *AppendEntityMemberShrinkRequest {
	s.ApplyType = &v
	return s
}

func (s *AppendEntityMemberShrinkRequest) SetEntityId(v int64) *AppendEntityMemberShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *AppendEntityMemberShrinkRequest) SetMemberShrink(v string) *AppendEntityMemberShrinkRequest {
	s.MemberShrink = &v
	return s
}

type AppendEntityMemberResponseBody struct {
	EntityId  *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppendEntityMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberResponseBody) SetEntityId(v int64) *AppendEntityMemberResponseBody {
	s.EntityId = &v
	return s
}

func (s *AppendEntityMemberResponseBody) SetRequestId(v string) *AppendEntityMemberResponseBody {
	s.RequestId = &v
	return s
}

type AppendEntityMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendEntityMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendEntityMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberResponse) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberResponse) SetHeaders(v map[string]*string) *AppendEntityMemberResponse {
	s.Headers = v
	return s
}

func (s *AppendEntityMemberResponse) SetStatusCode(v int32) *AppendEntityMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendEntityMemberResponse) SetBody(v *AppendEntityMemberResponseBody) *AppendEntityMemberResponse {
	s.Body = v
	return s
}

type AssociateRequest struct {
	AgentKey     *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Perspective  []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	RecommendNum *int32    `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	SessionId    *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Utterance    *string   `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s AssociateRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateRequest) GoString() string {
	return s.String()
}

func (s *AssociateRequest) SetAgentKey(v string) *AssociateRequest {
	s.AgentKey = &v
	return s
}

func (s *AssociateRequest) SetInstanceId(v string) *AssociateRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateRequest) SetPerspective(v []*string) *AssociateRequest {
	s.Perspective = v
	return s
}

func (s *AssociateRequest) SetRecommendNum(v int32) *AssociateRequest {
	s.RecommendNum = &v
	return s
}

func (s *AssociateRequest) SetSessionId(v string) *AssociateRequest {
	s.SessionId = &v
	return s
}

func (s *AssociateRequest) SetUtterance(v string) *AssociateRequest {
	s.Utterance = &v
	return s
}

type AssociateResponseBody struct {
	Associate []*AssociateResponseBodyAssociate `json:"Associate,omitempty" xml:"Associate,omitempty" type:"Repeated"`
	MessageId *string                           `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionId *string                           `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AssociateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResponseBody) SetAssociate(v []*AssociateResponseBodyAssociate) *AssociateResponseBody {
	s.Associate = v
	return s
}

func (s *AssociateResponseBody) SetMessageId(v string) *AssociateResponseBody {
	s.MessageId = &v
	return s
}

func (s *AssociateResponseBody) SetRequestId(v string) *AssociateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResponseBody) SetSessionId(v string) *AssociateResponseBody {
	s.SessionId = &v
	return s
}

type AssociateResponseBodyAssociate struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AssociateResponseBodyAssociate) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponseBodyAssociate) GoString() string {
	return s.String()
}

func (s *AssociateResponseBodyAssociate) SetTitle(v string) *AssociateResponseBodyAssociate {
	s.Title = &v
	return s
}

type AssociateResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponse) GoString() string {
	return s.String()
}

func (s *AssociateResponse) SetHeaders(v map[string]*string) *AssociateResponse {
	s.Headers = v
	return s
}

func (s *AssociateResponse) SetStatusCode(v int32) *AssociateResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResponse) SetBody(v *AssociateResponseBody) *AssociateResponse {
	s.Body = v
	return s
}

type ChatRequest struct {
	AgentKey    *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InstanceId  *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentName  *string   `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	KnowledgeId *string   `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	SenderId    *string   `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SenderNick  *string   `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	SessionId   *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Tag         *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Utterance   *string   `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	VendorParam *string   `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
}

func (s ChatRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) SetAgentKey(v string) *ChatRequest {
	s.AgentKey = &v
	return s
}

func (s *ChatRequest) SetInstanceId(v string) *ChatRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatRequest) SetIntentName(v string) *ChatRequest {
	s.IntentName = &v
	return s
}

func (s *ChatRequest) SetKnowledgeId(v string) *ChatRequest {
	s.KnowledgeId = &v
	return s
}

func (s *ChatRequest) SetPerspective(v []*string) *ChatRequest {
	s.Perspective = v
	return s
}

func (s *ChatRequest) SetSenderId(v string) *ChatRequest {
	s.SenderId = &v
	return s
}

func (s *ChatRequest) SetSenderNick(v string) *ChatRequest {
	s.SenderNick = &v
	return s
}

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

func (s *ChatRequest) SetTag(v string) *ChatRequest {
	s.Tag = &v
	return s
}

func (s *ChatRequest) SetUtterance(v string) *ChatRequest {
	s.Utterance = &v
	return s
}

func (s *ChatRequest) SetVendorParam(v string) *ChatRequest {
	s.VendorParam = &v
	return s
}

type ChatResponseBody struct {
	MessageId *string                     `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Messages  []*ChatResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionId *string                     `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Tag       *string                     `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) SetMessageId(v string) *ChatResponseBody {
	s.MessageId = &v
	return s
}

func (s *ChatResponseBody) SetMessages(v []*ChatResponseBodyMessages) *ChatResponseBody {
	s.Messages = v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) SetSessionId(v string) *ChatResponseBody {
	s.SessionId = &v
	return s
}

func (s *ChatResponseBody) SetTag(v string) *ChatResponseBody {
	s.Tag = &v
	return s
}

type ChatResponseBodyMessages struct {
	AnswerSource  *string                                `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	AnswerType    *string                                `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Knowledge     *ChatResponseBodyMessagesKnowledge     `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
	Recommends    []*ChatResponseBodyMessagesRecommends  `json:"Recommends,omitempty" xml:"Recommends,omitempty" type:"Repeated"`
	Text          *ChatResponseBodyMessagesText          `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
	Title         *string                                `json:"Title,omitempty" xml:"Title,omitempty"`
	Type          *string                                `json:"Type,omitempty" xml:"Type,omitempty"`
	VoiceStrategy *ChatResponseBodyMessagesVoiceStrategy `json:"VoiceStrategy,omitempty" xml:"VoiceStrategy,omitempty" type:"Struct"`
	VoiceTitle    *string                                `json:"VoiceTitle,omitempty" xml:"VoiceTitle,omitempty"`
}

func (s ChatResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessages) SetAnswerSource(v string) *ChatResponseBodyMessages {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessages) SetAnswerType(v string) *ChatResponseBodyMessages {
	s.AnswerType = &v
	return s
}

func (s *ChatResponseBodyMessages) SetKnowledge(v *ChatResponseBodyMessagesKnowledge) *ChatResponseBodyMessages {
	s.Knowledge = v
	return s
}

func (s *ChatResponseBodyMessages) SetRecommends(v []*ChatResponseBodyMessagesRecommends) *ChatResponseBodyMessages {
	s.Recommends = v
	return s
}

func (s *ChatResponseBodyMessages) SetText(v *ChatResponseBodyMessagesText) *ChatResponseBodyMessages {
	s.Text = v
	return s
}

func (s *ChatResponseBodyMessages) SetTitle(v string) *ChatResponseBodyMessages {
	s.Title = &v
	return s
}

func (s *ChatResponseBodyMessages) SetType(v string) *ChatResponseBodyMessages {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyMessages) SetVoiceStrategy(v *ChatResponseBodyMessagesVoiceStrategy) *ChatResponseBodyMessages {
	s.VoiceStrategy = v
	return s
}

func (s *ChatResponseBodyMessages) SetVoiceTitle(v string) *ChatResponseBodyMessages {
	s.VoiceTitle = &v
	return s
}

type ChatResponseBodyMessagesKnowledge struct {
	AnswerSource      *string                                               `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	Category          *string                                               `json:"Category,omitempty" xml:"Category,omitempty"`
	Content           *string                                               `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentType       *string                                               `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	HitStatement      *string                                               `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	Id                *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	RelatedKnowledges []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges `json:"RelatedKnowledges,omitempty" xml:"RelatedKnowledges,omitempty" type:"Repeated"`
	Score             *float64                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Summary           *string                                               `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title             *string                                               `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ChatResponseBodyMessagesKnowledge) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesKnowledge) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesKnowledge) SetAnswerSource(v string) *ChatResponseBodyMessagesKnowledge {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetCategory(v string) *ChatResponseBodyMessagesKnowledge {
	s.Category = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetContent(v string) *ChatResponseBodyMessagesKnowledge {
	s.Content = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetContentType(v string) *ChatResponseBodyMessagesKnowledge {
	s.ContentType = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetHitStatement(v string) *ChatResponseBodyMessagesKnowledge {
	s.HitStatement = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetId(v string) *ChatResponseBodyMessagesKnowledge {
	s.Id = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetRelatedKnowledges(v []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges) *ChatResponseBodyMessagesKnowledge {
	s.RelatedKnowledges = v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetScore(v float64) *ChatResponseBodyMessagesKnowledge {
	s.Score = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetSummary(v string) *ChatResponseBodyMessagesKnowledge {
	s.Summary = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetTitle(v string) *ChatResponseBodyMessagesKnowledge {
	s.Title = &v
	return s
}

type ChatResponseBodyMessagesKnowledgeRelatedKnowledges struct {
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ChatResponseBodyMessagesKnowledgeRelatedKnowledges) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesKnowledgeRelatedKnowledges) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) SetKnowledgeId(v string) *ChatResponseBodyMessagesKnowledgeRelatedKnowledges {
	s.KnowledgeId = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) SetTitle(v string) *ChatResponseBodyMessagesKnowledgeRelatedKnowledges {
	s.Title = &v
	return s
}

type ChatResponseBodyMessagesRecommends struct {
	AnswerSource *string  `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	KnowledgeId  *string  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	Score        *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	Title        *string  `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ChatResponseBodyMessagesRecommends) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesRecommends) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesRecommends) SetAnswerSource(v string) *ChatResponseBodyMessagesRecommends {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetKnowledgeId(v string) *ChatResponseBodyMessagesRecommends {
	s.KnowledgeId = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetScore(v float64) *ChatResponseBodyMessagesRecommends {
	s.Score = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetTitle(v string) *ChatResponseBodyMessagesRecommends {
	s.Title = &v
	return s
}

type ChatResponseBodyMessagesText struct {
	AnswerSource         *string                              `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	ArticleTitle         *string                              `json:"ArticleTitle,omitempty" xml:"ArticleTitle,omitempty"`
	Commands             map[string]interface{}               `json:"Commands,omitempty" xml:"Commands,omitempty"`
	Content              *string                              `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentType          *string                              `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DialogName           *string                              `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	Ext                  map[string]interface{}               `json:"Ext,omitempty" xml:"Ext,omitempty"`
	ExternalFlags        map[string]interface{}               `json:"ExternalFlags,omitempty" xml:"ExternalFlags,omitempty"`
	HitStatement         *string                              `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	IntentName           *string                              `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	MetaData             *string                              `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	NodeId               *string                              `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName             *string                              `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Score                *float64                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Slots                []*ChatResponseBodyMessagesTextSlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	UserDefinedChatTitle *string                              `json:"UserDefinedChatTitle,omitempty" xml:"UserDefinedChatTitle,omitempty"`
}

func (s ChatResponseBodyMessagesText) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesText) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesText) SetAnswerSource(v string) *ChatResponseBodyMessagesText {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetArticleTitle(v string) *ChatResponseBodyMessagesText {
	s.ArticleTitle = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetCommands(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.Commands = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetContent(v string) *ChatResponseBodyMessagesText {
	s.Content = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetContentType(v string) *ChatResponseBodyMessagesText {
	s.ContentType = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetDialogName(v string) *ChatResponseBodyMessagesText {
	s.DialogName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetExt(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.Ext = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetExternalFlags(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.ExternalFlags = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetHitStatement(v string) *ChatResponseBodyMessagesText {
	s.HitStatement = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetIntentName(v string) *ChatResponseBodyMessagesText {
	s.IntentName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetMetaData(v string) *ChatResponseBodyMessagesText {
	s.MetaData = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetNodeId(v string) *ChatResponseBodyMessagesText {
	s.NodeId = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetNodeName(v string) *ChatResponseBodyMessagesText {
	s.NodeName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetScore(v float64) *ChatResponseBodyMessagesText {
	s.Score = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetSlots(v []*ChatResponseBodyMessagesTextSlots) *ChatResponseBodyMessagesText {
	s.Slots = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetUserDefinedChatTitle(v string) *ChatResponseBodyMessagesText {
	s.UserDefinedChatTitle = &v
	return s
}

type ChatResponseBodyMessagesTextSlots struct {
	IsHit  *bool   `json:"IsHit,omitempty" xml:"IsHit,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ChatResponseBodyMessagesTextSlots) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesTextSlots) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesTextSlots) SetIsHit(v bool) *ChatResponseBodyMessagesTextSlots {
	s.IsHit = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetName(v string) *ChatResponseBodyMessagesTextSlots {
	s.Name = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetOrigin(v string) *ChatResponseBodyMessagesTextSlots {
	s.Origin = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetValue(v string) *ChatResponseBodyMessagesTextSlots {
	s.Value = &v
	return s
}

type ChatResponseBodyMessagesVoiceStrategy struct {
	AsrMaxEndSilence *string `json:"AsrMaxEndSilence,omitempty" xml:"AsrMaxEndSilence,omitempty"`
	CollectNumber    *bool   `json:"CollectNumber,omitempty" xml:"CollectNumber,omitempty"`
	HangUp           *bool   `json:"HangUp,omitempty" xml:"HangUp,omitempty"`
	Interruptible    *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	MaxDigits        *int64  `json:"MaxDigits,omitempty" xml:"MaxDigits,omitempty"`
	ReplyTimeout     *int64  `json:"ReplyTimeout,omitempty" xml:"ReplyTimeout,omitempty"`
	Terminator       *string `json:"Terminator,omitempty" xml:"Terminator,omitempty"`
}

func (s ChatResponseBodyMessagesVoiceStrategy) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesVoiceStrategy) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesVoiceStrategy) SetAsrMaxEndSilence(v string) *ChatResponseBodyMessagesVoiceStrategy {
	s.AsrMaxEndSilence = &v
	return s
}

func (s *ChatResponseBodyMessagesVoiceStrategy) SetCollectNumber(v bool) *ChatResponseBodyMessagesVoiceStrategy {
	s.CollectNumber = &v
	return s
}

func (s *ChatResponseBodyMessagesVoiceStrategy) SetHangUp(v bool) *ChatResponseBodyMessagesVoiceStrategy {
	s.HangUp = &v
	return s
}

func (s *ChatResponseBodyMessagesVoiceStrategy) SetInterruptible(v bool) *ChatResponseBodyMessagesVoiceStrategy {
	s.Interruptible = &v
	return s
}

func (s *ChatResponseBodyMessagesVoiceStrategy) SetMaxDigits(v int64) *ChatResponseBodyMessagesVoiceStrategy {
	s.MaxDigits = &v
	return s
}

func (s *ChatResponseBodyMessagesVoiceStrategy) SetReplyTimeout(v int64) *ChatResponseBodyMessagesVoiceStrategy {
	s.ReplyTimeout = &v
	return s
}

func (s *ChatResponseBodyMessagesVoiceStrategy) SetTerminator(v string) *ChatResponseBodyMessagesVoiceStrategy {
	s.Terminator = &v
	return s
}

type ChatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatResponse) GoString() string {
	return s.String()
}

func (s *ChatResponse) SetHeaders(v map[string]*string) *ChatResponse {
	s.Headers = v
	return s
}

func (s *ChatResponse) SetStatusCode(v int32) *ChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatResponse) SetBody(v *ChatResponseBody) *ChatResponse {
	s.Body = v
	return s
}

type CreateBotRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Avatar       *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RobotType    *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s CreateBotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBotRequest) GoString() string {
	return s.String()
}

func (s *CreateBotRequest) SetAgentKey(v string) *CreateBotRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateBotRequest) SetAvatar(v string) *CreateBotRequest {
	s.Avatar = &v
	return s
}

func (s *CreateBotRequest) SetIntroduction(v string) *CreateBotRequest {
	s.Introduction = &v
	return s
}

func (s *CreateBotRequest) SetLanguageCode(v string) *CreateBotRequest {
	s.LanguageCode = &v
	return s
}

func (s *CreateBotRequest) SetName(v string) *CreateBotRequest {
	s.Name = &v
	return s
}

func (s *CreateBotRequest) SetRobotType(v string) *CreateBotRequest {
	s.RobotType = &v
	return s
}

type CreateBotResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBotResponseBody) SetInstanceId(v string) *CreateBotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateBotResponseBody) SetRequestId(v string) *CreateBotResponseBody {
	s.RequestId = &v
	return s
}

type CreateBotResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBotResponse) GoString() string {
	return s.String()
}

func (s *CreateBotResponse) SetHeaders(v map[string]*string) *CreateBotResponse {
	s.Headers = v
	return s
}

func (s *CreateBotResponse) SetStatusCode(v int32) *CreateBotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBotResponse) SetBody(v *CreateBotResponseBody) *CreateBotResponse {
	s.Body = v
	return s
}

type CreateCategoryRequest struct {
	AgentKey         *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s CreateCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCategoryRequest) SetAgentKey(v string) *CreateCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateCategoryRequest) SetName(v string) *CreateCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateCategoryRequest) SetParentCategoryId(v int64) *CreateCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

type CreateCategoryResponseBody struct {
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponseBody) SetCategoryId(v int64) *CreateCategoryResponseBody {
	s.CategoryId = &v
	return s
}

func (s *CreateCategoryResponseBody) SetRequestId(v string) *CreateCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCategoryResponseBody) SetSuccess(v bool) *CreateCategoryResponseBody {
	s.Success = &v
	return s
}

type CreateCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponse) SetHeaders(v map[string]*string) *CreateCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateCategoryResponse) SetStatusCode(v int32) *CreateCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCategoryResponse) SetBody(v *CreateCategoryResponseBody) *CreateCategoryResponse {
	s.Body = v
	return s
}

type CreateCoreWordRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
}

func (s CreateCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCoreWordRequest) GoString() string {
	return s.String()
}

func (s *CreateCoreWordRequest) SetAgentKey(v string) *CreateCoreWordRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateCoreWordRequest) SetCoreWordName(v string) *CreateCoreWordRequest {
	s.CoreWordName = &v
	return s
}

type CreateCoreWordResponseBody struct {
	CoreWordCode *string `json:"CoreWordCode,omitempty" xml:"CoreWordCode,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCoreWordResponseBody) SetCoreWordCode(v string) *CreateCoreWordResponseBody {
	s.CoreWordCode = &v
	return s
}

func (s *CreateCoreWordResponseBody) SetRequestId(v string) *CreateCoreWordResponseBody {
	s.RequestId = &v
	return s
}

type CreateCoreWordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCoreWordResponse) GoString() string {
	return s.String()
}

func (s *CreateCoreWordResponse) SetHeaders(v map[string]*string) *CreateCoreWordResponse {
	s.Headers = v
	return s
}

func (s *CreateCoreWordResponse) SetStatusCode(v int32) *CreateCoreWordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCoreWordResponse) SetBody(v *CreateCoreWordResponseBody) *CreateCoreWordResponse {
	s.Body = v
	return s
}

type CreateDialogRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DialogName  *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogRequest) GoString() string {
	return s.String()
}

func (s *CreateDialogRequest) SetAgentKey(v string) *CreateDialogRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDialogRequest) SetDescription(v string) *CreateDialogRequest {
	s.Description = &v
	return s
}

func (s *CreateDialogRequest) SetDialogName(v string) *CreateDialogRequest {
	s.DialogName = &v
	return s
}

func (s *CreateDialogRequest) SetInstanceId(v string) *CreateDialogRequest {
	s.InstanceId = &v
	return s
}

type CreateDialogResponseBody struct {
	DialogId  *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDialogResponseBody) SetDialogId(v int64) *CreateDialogResponseBody {
	s.DialogId = &v
	return s
}

func (s *CreateDialogResponseBody) SetRequestId(v string) *CreateDialogResponseBody {
	s.RequestId = &v
	return s
}

type CreateDialogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogResponse) GoString() string {
	return s.String()
}

func (s *CreateDialogResponse) SetHeaders(v map[string]*string) *CreateDialogResponse {
	s.Headers = v
	return s
}

func (s *CreateDialogResponse) SetStatusCode(v int32) *CreateDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDialogResponse) SetBody(v *CreateDialogResponseBody) *CreateDialogResponse {
	s.Body = v
	return s
}

type CreateEntityRequest struct {
	AgentKey   *string                       `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId   *int64                        `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	EntityName *string                       `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType *string                       `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Members    []*CreateEntityRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Regex      *string                       `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s CreateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityRequest) SetAgentKey(v string) *CreateEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateEntityRequest) SetDialogId(v int64) *CreateEntityRequest {
	s.DialogId = &v
	return s
}

func (s *CreateEntityRequest) SetEntityName(v string) *CreateEntityRequest {
	s.EntityName = &v
	return s
}

func (s *CreateEntityRequest) SetEntityType(v string) *CreateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *CreateEntityRequest) SetMembers(v []*CreateEntityRequestMembers) *CreateEntityRequest {
	s.Members = v
	return s
}

func (s *CreateEntityRequest) SetRegex(v string) *CreateEntityRequest {
	s.Regex = &v
	return s
}

type CreateEntityRequestMembers struct {
	MemberName *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s CreateEntityRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateEntityRequestMembers) SetMemberName(v string) *CreateEntityRequestMembers {
	s.MemberName = &v
	return s
}

func (s *CreateEntityRequestMembers) SetSynonyms(v []*string) *CreateEntityRequestMembers {
	s.Synonyms = v
	return s
}

type CreateEntityShrinkRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId      *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	EntityName    *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	Regex         *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s CreateEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityShrinkRequest) SetAgentKey(v string) *CreateEntityShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetDialogId(v int64) *CreateEntityShrinkRequest {
	s.DialogId = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetEntityName(v string) *CreateEntityShrinkRequest {
	s.EntityName = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetEntityType(v string) *CreateEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetMembersShrink(v string) *CreateEntityShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetRegex(v string) *CreateEntityShrinkRequest {
	s.Regex = &v
	return s
}

type CreateEntityResponseBody struct {
	EntityId  *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityResponseBody) SetEntityId(v int64) *CreateEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *CreateEntityResponseBody) SetRequestId(v string) *CreateEntityResponseBody {
	s.RequestId = &v
	return s
}

type CreateEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityResponse) SetHeaders(v map[string]*string) *CreateEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityResponse) SetStatusCode(v int32) *CreateEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEntityResponse) SetBody(v *CreateEntityResponseBody) *CreateEntityResponse {
	s.Body = v
	return s
}

type CreateIntentRequest struct {
	AgentKey         *string          `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId         *int64           `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	IntentDefinition *IntentCreateDTO `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
}

func (s CreateIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentRequest) SetAgentKey(v string) *CreateIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateIntentRequest) SetDialogId(v int64) *CreateIntentRequest {
	s.DialogId = &v
	return s
}

func (s *CreateIntentRequest) SetIntentDefinition(v *IntentCreateDTO) *CreateIntentRequest {
	s.IntentDefinition = v
	return s
}

type CreateIntentShrinkRequest struct {
	AgentKey               *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId               *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
}

func (s CreateIntentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentShrinkRequest) SetAgentKey(v string) *CreateIntentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateIntentShrinkRequest) SetDialogId(v int64) *CreateIntentShrinkRequest {
	s.DialogId = &v
	return s
}

func (s *CreateIntentShrinkRequest) SetIntentDefinitionShrink(v string) *CreateIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

type CreateIntentResponseBody struct {
	IntentId  *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntentResponseBody) SetIntentId(v int64) *CreateIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *CreateIntentResponseBody) SetRequestId(v string) *CreateIntentResponseBody {
	s.RequestId = &v
	return s
}

type CreateIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentResponse) GoString() string {
	return s.String()
}

func (s *CreateIntentResponse) SetHeaders(v map[string]*string) *CreateIntentResponse {
	s.Headers = v
	return s
}

func (s *CreateIntentResponse) SetStatusCode(v int32) *CreateIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntentResponse) SetBody(v *CreateIntentResponseBody) *CreateIntentResponse {
	s.Body = v
	return s
}

type CreateKnowledgeRequest struct {
	AgentKey  *string                          `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Knowledge *CreateKnowledgeRequestKnowledge `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
}

func (s CreateKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeRequest) SetAgentKey(v string) *CreateKnowledgeRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateKnowledgeRequest) SetKnowledge(v *CreateKnowledgeRequestKnowledge) *CreateKnowledgeRequest {
	s.Knowledge = v
	return s
}

type CreateKnowledgeRequestKnowledge struct {
	CategoryId     *int64                                         `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	EndDate        *string                                        `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	KnowledgeTitle *string                                        `json:"KnowledgeTitle,omitempty" xml:"KnowledgeTitle,omitempty"`
	KnowledgeType  *int32                                         `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	Outlines       []*CreateKnowledgeRequestKnowledgeOutlines     `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	SimQuestions   []*CreateKnowledgeRequestKnowledgeSimQuestions `json:"SimQuestions,omitempty" xml:"SimQuestions,omitempty" type:"Repeated"`
	Solutions      []*CreateKnowledgeRequestKnowledgeSolutions    `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	StartDate      *string                                        `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s CreateKnowledgeRequestKnowledge) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeRequestKnowledge) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeRequestKnowledge) SetCategoryId(v int64) *CreateKnowledgeRequestKnowledge {
	s.CategoryId = &v
	return s
}

func (s *CreateKnowledgeRequestKnowledge) SetEndDate(v string) *CreateKnowledgeRequestKnowledge {
	s.EndDate = &v
	return s
}

func (s *CreateKnowledgeRequestKnowledge) SetKnowledgeTitle(v string) *CreateKnowledgeRequestKnowledge {
	s.KnowledgeTitle = &v
	return s
}

func (s *CreateKnowledgeRequestKnowledge) SetKnowledgeType(v int32) *CreateKnowledgeRequestKnowledge {
	s.KnowledgeType = &v
	return s
}

func (s *CreateKnowledgeRequestKnowledge) SetOutlines(v []*CreateKnowledgeRequestKnowledgeOutlines) *CreateKnowledgeRequestKnowledge {
	s.Outlines = v
	return s
}

func (s *CreateKnowledgeRequestKnowledge) SetSimQuestions(v []*CreateKnowledgeRequestKnowledgeSimQuestions) *CreateKnowledgeRequestKnowledge {
	s.SimQuestions = v
	return s
}

func (s *CreateKnowledgeRequestKnowledge) SetSolutions(v []*CreateKnowledgeRequestKnowledgeSolutions) *CreateKnowledgeRequestKnowledge {
	s.Solutions = v
	return s
}

func (s *CreateKnowledgeRequestKnowledge) SetStartDate(v string) *CreateKnowledgeRequestKnowledge {
	s.StartDate = &v
	return s
}

type CreateKnowledgeRequestKnowledgeOutlines struct {
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	OutlineId   *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateKnowledgeRequestKnowledgeOutlines) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeRequestKnowledgeOutlines) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeRequestKnowledgeOutlines) SetKnowledgeId(v int64) *CreateKnowledgeRequestKnowledgeOutlines {
	s.KnowledgeId = &v
	return s
}

func (s *CreateKnowledgeRequestKnowledgeOutlines) SetOutlineId(v int64) *CreateKnowledgeRequestKnowledgeOutlines {
	s.OutlineId = &v
	return s
}

func (s *CreateKnowledgeRequestKnowledgeOutlines) SetTitle(v string) *CreateKnowledgeRequestKnowledgeOutlines {
	s.Title = &v
	return s
}

type CreateKnowledgeRequestKnowledgeSimQuestions struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateKnowledgeRequestKnowledgeSimQuestions) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeRequestKnowledgeSimQuestions) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeRequestKnowledgeSimQuestions) SetTitle(v string) *CreateKnowledgeRequestKnowledgeSimQuestions {
	s.Title = &v
	return s
}

type CreateKnowledgeRequestKnowledgeSolutions struct {
	Content        *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	PerspectiveIds []*string `json:"PerspectiveIds,omitempty" xml:"PerspectiveIds,omitempty" type:"Repeated"`
	PlainText      *string   `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
}

func (s CreateKnowledgeRequestKnowledgeSolutions) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeRequestKnowledgeSolutions) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeRequestKnowledgeSolutions) SetContent(v string) *CreateKnowledgeRequestKnowledgeSolutions {
	s.Content = &v
	return s
}

func (s *CreateKnowledgeRequestKnowledgeSolutions) SetPerspectiveIds(v []*string) *CreateKnowledgeRequestKnowledgeSolutions {
	s.PerspectiveIds = v
	return s
}

func (s *CreateKnowledgeRequestKnowledgeSolutions) SetPlainText(v string) *CreateKnowledgeRequestKnowledgeSolutions {
	s.PlainText = &v
	return s
}

type CreateKnowledgeShrinkRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	KnowledgeShrink *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
}

func (s CreateKnowledgeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeShrinkRequest) SetAgentKey(v string) *CreateKnowledgeShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateKnowledgeShrinkRequest) SetKnowledgeShrink(v string) *CreateKnowledgeShrinkRequest {
	s.KnowledgeShrink = &v
	return s
}

type CreateKnowledgeResponseBody struct {
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeResponseBody) SetKnowledgeId(v int64) *CreateKnowledgeResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *CreateKnowledgeResponseBody) SetRequestId(v string) *CreateKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type CreateKnowledgeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeResponse) SetHeaders(v map[string]*string) *CreateKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeResponse) SetStatusCode(v int32) *CreateKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeResponse) SetBody(v *CreateKnowledgeResponseBody) *CreateKnowledgeResponse {
	s.Body = v
	return s
}

type CreatePerspectiveRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveRequest) SetAgentKey(v string) *CreatePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *CreatePerspectiveRequest) SetName(v string) *CreatePerspectiveRequest {
	s.Name = &v
	return s
}

type CreatePerspectiveResponseBody struct {
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveResponseBody) SetPerspectiveId(v string) *CreatePerspectiveResponseBody {
	s.PerspectiveId = &v
	return s
}

func (s *CreatePerspectiveResponseBody) SetRequestId(v string) *CreatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

type CreatePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveResponse) SetHeaders(v map[string]*string) *CreatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *CreatePerspectiveResponse) SetStatusCode(v int32) *CreatePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePerspectiveResponse) SetBody(v *CreatePerspectiveResponseBody) *CreatePerspectiveResponse {
	s.Body = v
	return s
}

type DeleteBotRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteBotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBotRequest) GoString() string {
	return s.String()
}

func (s *DeleteBotRequest) SetAgentKey(v string) *DeleteBotRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteBotRequest) SetInstanceId(v string) *DeleteBotRequest {
	s.InstanceId = &v
	return s
}

type DeleteBotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBotResponseBody) SetRequestId(v string) *DeleteBotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBotResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBotResponse) GoString() string {
	return s.String()
}

func (s *DeleteBotResponse) SetHeaders(v map[string]*string) *DeleteBotResponse {
	s.Headers = v
	return s
}

func (s *DeleteBotResponse) SetStatusCode(v int32) *DeleteBotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBotResponse) SetBody(v *DeleteBotResponseBody) *DeleteBotResponse {
	s.Body = v
	return s
}

type DeleteCategoryRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DeleteCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategoryRequest) SetAgentKey(v string) *DeleteCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCategoryRequest) SetCategoryId(v int64) *DeleteCategoryRequest {
	s.CategoryId = &v
	return s
}

type DeleteCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBody) SetRequestId(v string) *DeleteCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetSuccess(v bool) *DeleteCategoryResponseBody {
	s.Success = &v
	return s
}

type DeleteCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponse) SetHeaders(v map[string]*string) *DeleteCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteCategoryResponse) SetStatusCode(v int32) *DeleteCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCategoryResponse) SetBody(v *DeleteCategoryResponseBody) *DeleteCategoryResponse {
	s.Body = v
	return s
}

type DeleteCoreWordRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
}

func (s DeleteCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCoreWordRequest) GoString() string {
	return s.String()
}

func (s *DeleteCoreWordRequest) SetAgentKey(v string) *DeleteCoreWordRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCoreWordRequest) SetCoreWordName(v string) *DeleteCoreWordRequest {
	s.CoreWordName = &v
	return s
}

type DeleteCoreWordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCoreWordResponseBody) SetRequestId(v string) *DeleteCoreWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCoreWordResponseBody) SetSuccess(v bool) *DeleteCoreWordResponseBody {
	s.Success = &v
	return s
}

type DeleteCoreWordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCoreWordResponse) GoString() string {
	return s.String()
}

func (s *DeleteCoreWordResponse) SetHeaders(v map[string]*string) *DeleteCoreWordResponse {
	s.Headers = v
	return s
}

func (s *DeleteCoreWordResponse) SetStatusCode(v int32) *DeleteCoreWordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCoreWordResponse) SetBody(v *DeleteCoreWordResponseBody) *DeleteCoreWordResponse {
	s.Body = v
	return s
}

type DeleteDialogRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s DeleteDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDialogRequest) GoString() string {
	return s.String()
}

func (s *DeleteDialogRequest) SetAgentKey(v string) *DeleteDialogRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDialogRequest) SetDialogId(v int64) *DeleteDialogRequest {
	s.DialogId = &v
	return s
}

type DeleteDialogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDialogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDialogResponseBody) SetRequestId(v string) *DeleteDialogResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDialogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDialogResponse) GoString() string {
	return s.String()
}

func (s *DeleteDialogResponse) SetHeaders(v map[string]*string) *DeleteDialogResponse {
	s.Headers = v
	return s
}

func (s *DeleteDialogResponse) SetStatusCode(v int32) *DeleteDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDialogResponse) SetBody(v *DeleteDialogResponseBody) *DeleteDialogResponse {
	s.Body = v
	return s
}

type DeleteEntityRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EntityId *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DeleteEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityRequest) SetAgentKey(v string) *DeleteEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteEntityRequest) SetEntityId(v int64) *DeleteEntityRequest {
	s.EntityId = &v
	return s
}

type DeleteEntityResponseBody struct {
	EntityId  *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityResponseBody) SetEntityId(v int64) *DeleteEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *DeleteEntityResponseBody) SetRequestId(v string) *DeleteEntityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityResponse) SetHeaders(v map[string]*string) *DeleteEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityResponse) SetStatusCode(v int32) *DeleteEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEntityResponse) SetBody(v *DeleteEntityResponseBody) *DeleteEntityResponse {
	s.Body = v
	return s
}

type DeleteIntentRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	IntentId *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DeleteIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntentRequest) SetAgentKey(v string) *DeleteIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteIntentRequest) SetIntentId(v int64) *DeleteIntentRequest {
	s.IntentId = &v
	return s
}

type DeleteIntentResponseBody struct {
	IntentId  *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponseBody) SetIntentId(v int64) *DeleteIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DeleteIntentResponseBody) SetRequestId(v string) *DeleteIntentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponse) SetHeaders(v map[string]*string) *DeleteIntentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntentResponse) SetStatusCode(v int32) *DeleteIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntentResponse) SetBody(v *DeleteIntentResponseBody) *DeleteIntentResponse {
	s.Body = v
	return s
}

type DeleteKnowledgeRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DeleteKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeRequest) SetAgentKey(v string) *DeleteKnowledgeRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetKnowledgeId(v int64) *DeleteKnowledgeRequest {
	s.KnowledgeId = &v
	return s
}

type DeleteKnowledgeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponseBody) SetRequestId(v string) *DeleteKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKnowledgeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeResponse) SetStatusCode(v int32) *DeleteKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeResponse) SetBody(v *DeleteKnowledgeResponseBody) *DeleteKnowledgeResponse {
	s.Body = v
	return s
}

type DeletePerspectiveRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s DeletePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveRequest) SetAgentKey(v string) *DeletePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *DeletePerspectiveRequest) SetPerspectiveId(v string) *DeletePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type DeletePerspectiveResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeletePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveResponseBody) SetRequestId(v string) *DeletePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePerspectiveResponseBody) SetResult(v bool) *DeletePerspectiveResponseBody {
	s.Result = &v
	return s
}

type DeletePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveResponse) SetHeaders(v map[string]*string) *DeletePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *DeletePerspectiveResponse) SetStatusCode(v int32) *DeletePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePerspectiveResponse) SetBody(v *DeletePerspectiveResponseBody) *DeletePerspectiveResponse {
	s.Body = v
	return s
}

type DescribeBotRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeBotRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotRequest) GoString() string {
	return s.String()
}

func (s *DescribeBotRequest) SetAgentKey(v string) *DescribeBotRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeBotRequest) SetInstanceId(v string) *DescribeBotRequest {
	s.InstanceId = &v
	return s
}

type DescribeBotResponseBody struct {
	Avatar       *string                              `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Categories   []*DescribeBotResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreateTime   *string                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId   *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Introduction *string                              `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	LanguageCode *string                              `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	Logo         *string                              `json:"Logo,omitempty" xml:"Logo,omitempty"`
	Name         *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeZone     *string                              `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeBotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBotResponseBody) SetAvatar(v string) *DescribeBotResponseBody {
	s.Avatar = &v
	return s
}

func (s *DescribeBotResponseBody) SetCategories(v []*DescribeBotResponseBodyCategories) *DescribeBotResponseBody {
	s.Categories = v
	return s
}

func (s *DescribeBotResponseBody) SetCreateTime(v string) *DescribeBotResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeBotResponseBody) SetInstanceId(v string) *DescribeBotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeBotResponseBody) SetIntroduction(v string) *DescribeBotResponseBody {
	s.Introduction = &v
	return s
}

func (s *DescribeBotResponseBody) SetLanguageCode(v string) *DescribeBotResponseBody {
	s.LanguageCode = &v
	return s
}

func (s *DescribeBotResponseBody) SetLogo(v string) *DescribeBotResponseBody {
	s.Logo = &v
	return s
}

func (s *DescribeBotResponseBody) SetName(v string) *DescribeBotResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeBotResponseBody) SetRequestId(v string) *DescribeBotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBotResponseBody) SetTimeZone(v string) *DescribeBotResponseBody {
	s.TimeZone = &v
	return s
}

type DescribeBotResponseBodyCategories struct {
	CategoryId       *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s DescribeBotResponseBodyCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *DescribeBotResponseBodyCategories) SetCategoryId(v int64) *DescribeBotResponseBodyCategories {
	s.CategoryId = &v
	return s
}

func (s *DescribeBotResponseBodyCategories) SetName(v string) *DescribeBotResponseBodyCategories {
	s.Name = &v
	return s
}

func (s *DescribeBotResponseBodyCategories) SetParentCategoryId(v int64) *DescribeBotResponseBodyCategories {
	s.ParentCategoryId = &v
	return s
}

type DescribeBotResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBotResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotResponse) GoString() string {
	return s.String()
}

func (s *DescribeBotResponse) SetHeaders(v map[string]*string) *DescribeBotResponse {
	s.Headers = v
	return s
}

func (s *DescribeBotResponse) SetStatusCode(v int32) *DescribeBotResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBotResponse) SetBody(v *DescribeBotResponseBody) *DescribeBotResponse {
	s.Body = v
	return s
}

type DescribeCategoryRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DescribeCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryRequest) SetAgentKey(v string) *DescribeCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeCategoryRequest) SetCategoryId(v int64) *DescribeCategoryRequest {
	s.CategoryId = &v
	return s
}

type DescribeCategoryResponseBody struct {
	CategoryId       *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponseBody) SetCategoryId(v int64) *DescribeCategoryResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBody) SetName(v string) *DescribeCategoryResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCategoryResponseBody) SetParentCategoryId(v int64) *DescribeCategoryResponseBody {
	s.ParentCategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBody) SetRequestId(v string) *DescribeCategoryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCategoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponse) SetHeaders(v map[string]*string) *DescribeCategoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoryResponse) SetStatusCode(v int32) *DescribeCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCategoryResponse) SetBody(v *DescribeCategoryResponseBody) *DescribeCategoryResponse {
	s.Body = v
	return s
}

type DescribeCoreWordRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
}

func (s DescribeCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCoreWordRequest) GoString() string {
	return s.String()
}

func (s *DescribeCoreWordRequest) SetAgentKey(v string) *DescribeCoreWordRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeCoreWordRequest) SetCoreWordName(v string) *DescribeCoreWordRequest {
	s.CoreWordName = &v
	return s
}

type DescribeCoreWordResponseBody struct {
	CoreWordCode *string   `json:"CoreWordCode,omitempty" xml:"CoreWordCode,omitempty"`
	CoreWordName *string   `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime   *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Synonyms     []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s DescribeCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCoreWordResponseBody) SetCoreWordCode(v string) *DescribeCoreWordResponseBody {
	s.CoreWordCode = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetCoreWordName(v string) *DescribeCoreWordResponseBody {
	s.CoreWordName = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetCreateTime(v string) *DescribeCoreWordResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetModifyTime(v string) *DescribeCoreWordResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetRequestId(v string) *DescribeCoreWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetSynonyms(v []*string) *DescribeCoreWordResponseBody {
	s.Synonyms = v
	return s
}

type DescribeCoreWordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCoreWordResponse) GoString() string {
	return s.String()
}

func (s *DescribeCoreWordResponse) SetHeaders(v map[string]*string) *DescribeCoreWordResponse {
	s.Headers = v
	return s
}

func (s *DescribeCoreWordResponse) SetStatusCode(v int32) *DescribeCoreWordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCoreWordResponse) SetBody(v *DescribeCoreWordResponseBody) *DescribeCoreWordResponse {
	s.Body = v
	return s
}

type DescribeDialogRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s DescribeDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDialogRequest) SetAgentKey(v string) *DescribeDialogRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDialogRequest) SetDialogId(v int64) *DescribeDialogRequest {
	s.DialogId = &v
	return s
}

type DescribeDialogResponseBody struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DialogId       *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	DialogName     *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	IsOnline       *bool   `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	IsSampleDialog *bool   `json:"IsSampleDialog,omitempty" xml:"IsSampleDialog,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserId   *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDialogResponseBody) SetCreateTime(v string) *DescribeDialogResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDialogResponseBody) SetCreateUserId(v string) *DescribeDialogResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetCreateUserName(v string) *DescribeDialogResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeDialogResponseBody) SetDescription(v string) *DescribeDialogResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDialogResponseBody) SetDialogId(v int64) *DescribeDialogResponseBody {
	s.DialogId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetDialogName(v string) *DescribeDialogResponseBody {
	s.DialogName = &v
	return s
}

func (s *DescribeDialogResponseBody) SetIsOnline(v bool) *DescribeDialogResponseBody {
	s.IsOnline = &v
	return s
}

func (s *DescribeDialogResponseBody) SetIsSampleDialog(v bool) *DescribeDialogResponseBody {
	s.IsSampleDialog = &v
	return s
}

func (s *DescribeDialogResponseBody) SetModifyTime(v string) *DescribeDialogResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDialogResponseBody) SetModifyUserId(v string) *DescribeDialogResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetModifyUserName(v string) *DescribeDialogResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeDialogResponseBody) SetRequestId(v string) *DescribeDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetStatus(v int32) *DescribeDialogResponseBody {
	s.Status = &v
	return s
}

type DescribeDialogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDialogResponse) SetHeaders(v map[string]*string) *DescribeDialogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDialogResponse) SetStatusCode(v int32) *DescribeDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDialogResponse) SetBody(v *DescribeDialogResponseBody) *DescribeDialogResponse {
	s.Body = v
	return s
}

type DescribeDialogFlowRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s DescribeDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeDialogFlowRequest) SetAgentKey(v string) *DescribeDialogFlowRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDialogFlowRequest) SetDialogId(v int64) *DescribeDialogFlowRequest {
	s.DialogId = &v
	return s
}

type DescribeDialogFlowResponseBody struct {
	AccountId        *string                `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateTime       *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId     *string                `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName   *string                `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DialogId         *int64                 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	DialogName       *string                `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	GlobalVars       map[string]interface{} `json:"GlobalVars,omitempty" xml:"GlobalVars,omitempty"`
	InstanceId       *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModifyTime       *string                `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserId     *string                `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName   *string                `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	ModuleDefinition *PaasProcessData       `json:"ModuleDefinition,omitempty" xml:"ModuleDefinition,omitempty"`
	ModuleId         *int64                 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName       *string                `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	RequestId        *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 0 for deleted, 1 for inSandbox, 2 for draft，3 for online
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// null
	Tags      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Templates *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s DescribeDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDialogFlowResponseBody) SetAccountId(v string) *DescribeDialogFlowResponseBody {
	s.AccountId = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetCreateTime(v string) *DescribeDialogFlowResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetCreateUserId(v string) *DescribeDialogFlowResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetCreateUserName(v string) *DescribeDialogFlowResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetDialogId(v int64) *DescribeDialogFlowResponseBody {
	s.DialogId = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetDialogName(v string) *DescribeDialogFlowResponseBody {
	s.DialogName = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetGlobalVars(v map[string]interface{}) *DescribeDialogFlowResponseBody {
	s.GlobalVars = v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetInstanceId(v string) *DescribeDialogFlowResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetModifyTime(v string) *DescribeDialogFlowResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetModifyUserId(v string) *DescribeDialogFlowResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetModifyUserName(v string) *DescribeDialogFlowResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetModuleDefinition(v *PaasProcessData) *DescribeDialogFlowResponseBody {
	s.ModuleDefinition = v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetModuleId(v int64) *DescribeDialogFlowResponseBody {
	s.ModuleId = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetModuleName(v string) *DescribeDialogFlowResponseBody {
	s.ModuleName = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetRequestId(v string) *DescribeDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetStatus(v int32) *DescribeDialogFlowResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetTags(v string) *DescribeDialogFlowResponseBody {
	s.Tags = &v
	return s
}

func (s *DescribeDialogFlowResponseBody) SetTemplates(v string) *DescribeDialogFlowResponseBody {
	s.Templates = &v
	return s
}

type DescribeDialogFlowResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeDialogFlowResponse) SetHeaders(v map[string]*string) *DescribeDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeDialogFlowResponse) SetStatusCode(v int32) *DescribeDialogFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDialogFlowResponse) SetBody(v *DescribeDialogFlowResponseBody) *DescribeDialogFlowResponse {
	s.Body = v
	return s
}

type DescribeEntitiesRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EntityId *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DescribeEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEntitiesRequest) SetAgentKey(v string) *DescribeEntitiesRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeEntitiesRequest) SetEntityId(v int64) *DescribeEntitiesRequest {
	s.EntityId = &v
	return s
}

type DescribeEntitiesResponseBody struct {
	CreateTime     *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string                                `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string                                `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	EntityId       *int64                                 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName     *string                                `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType     *string                                `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Members        []*DescribeEntitiesResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	ModifyTime     *string                                `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserId   *string                                `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string                                `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Regex          *string                                `json:"Regex,omitempty" xml:"Regex,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEntitiesResponseBody) SetCreateTime(v string) *DescribeEntitiesResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetCreateUserId(v string) *DescribeEntitiesResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetCreateUserName(v string) *DescribeEntitiesResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetEntityId(v int64) *DescribeEntitiesResponseBody {
	s.EntityId = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetEntityName(v string) *DescribeEntitiesResponseBody {
	s.EntityName = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetEntityType(v string) *DescribeEntitiesResponseBody {
	s.EntityType = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetMembers(v []*DescribeEntitiesResponseBodyMembers) *DescribeEntitiesResponseBody {
	s.Members = v
	return s
}

func (s *DescribeEntitiesResponseBody) SetModifyTime(v string) *DescribeEntitiesResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetModifyUserId(v string) *DescribeEntitiesResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetModifyUserName(v string) *DescribeEntitiesResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetRegex(v string) *DescribeEntitiesResponseBody {
	s.Regex = &v
	return s
}

func (s *DescribeEntitiesResponseBody) SetRequestId(v string) *DescribeEntitiesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEntitiesResponseBodyMembers struct {
	MemberName *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s DescribeEntitiesResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntitiesResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *DescribeEntitiesResponseBodyMembers) SetMemberName(v string) *DescribeEntitiesResponseBodyMembers {
	s.MemberName = &v
	return s
}

func (s *DescribeEntitiesResponseBodyMembers) SetSynonyms(v []*string) *DescribeEntitiesResponseBodyMembers {
	s.Synonyms = v
	return s
}

type DescribeEntitiesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEntitiesResponse) SetHeaders(v map[string]*string) *DescribeEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEntitiesResponse) SetStatusCode(v int32) *DescribeEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEntitiesResponse) SetBody(v *DescribeEntitiesResponseBody) *DescribeEntitiesResponse {
	s.Body = v
	return s
}

type DescribeIntentRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	IntentId *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DescribeIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntentRequest) SetAgentKey(v string) *DescribeIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeIntentRequest) SetIntentId(v int64) *DescribeIntentRequest {
	s.IntentId = &v
	return s
}

type DescribeIntentResponseBody struct {
	CreateTime     *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string                                `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string                                `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DialogId       *int64                                 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	IntentId       *int64                                 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	ModifyTime     *string                                `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserId   *string                                `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string                                `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name           *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleCheck      []*DescribeIntentResponseBodyRuleCheck `json:"RuleCheck,omitempty" xml:"RuleCheck,omitempty" type:"Repeated"`
	Slot           []*DescribeIntentResponseBodySlot      `json:"Slot,omitempty" xml:"Slot,omitempty" type:"Repeated"`
	Type           *string                                `json:"Type,omitempty" xml:"Type,omitempty"`
	UserSay        []*DescribeIntentResponseBodyUserSay   `json:"UserSay,omitempty" xml:"UserSay,omitempty" type:"Repeated"`
}

func (s DescribeIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBody) SetCreateTime(v string) *DescribeIntentResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateUserId(v string) *DescribeIntentResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateUserName(v string) *DescribeIntentResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetDialogId(v int64) *DescribeIntentResponseBody {
	s.DialogId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetIntentId(v int64) *DescribeIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyTime(v string) *DescribeIntentResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyUserId(v string) *DescribeIntentResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyUserName(v string) *DescribeIntentResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetName(v string) *DescribeIntentResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeIntentResponseBody) SetRequestId(v string) *DescribeIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetRuleCheck(v []*DescribeIntentResponseBodyRuleCheck) *DescribeIntentResponseBody {
	s.RuleCheck = v
	return s
}

func (s *DescribeIntentResponseBody) SetSlot(v []*DescribeIntentResponseBodySlot) *DescribeIntentResponseBody {
	s.Slot = v
	return s
}

func (s *DescribeIntentResponseBody) SetType(v string) *DescribeIntentResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeIntentResponseBody) SetUserSay(v []*DescribeIntentResponseBodyUserSay) *DescribeIntentResponseBody {
	s.UserSay = v
	return s
}

type DescribeIntentResponseBodyRuleCheck struct {
	Error   []*string `json:"Error,omitempty" xml:"Error,omitempty" type:"Repeated"`
	Strict  *bool     `json:"Strict,omitempty" xml:"Strict,omitempty"`
	Text    *string   `json:"Text,omitempty" xml:"Text,omitempty"`
	Warning []*string `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s DescribeIntentResponseBodyRuleCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBodyRuleCheck) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodyRuleCheck) SetError(v []*string) *DescribeIntentResponseBodyRuleCheck {
	s.Error = v
	return s
}

func (s *DescribeIntentResponseBodyRuleCheck) SetStrict(v bool) *DescribeIntentResponseBodyRuleCheck {
	s.Strict = &v
	return s
}

func (s *DescribeIntentResponseBodyRuleCheck) SetText(v string) *DescribeIntentResponseBodyRuleCheck {
	s.Text = &v
	return s
}

func (s *DescribeIntentResponseBodyRuleCheck) SetWarning(v []*string) *DescribeIntentResponseBodyRuleCheck {
	s.Warning = v
	return s
}

type DescribeIntentResponseBodySlot struct {
	IsArray     *bool                                 `json:"IsArray,omitempty" xml:"IsArray,omitempty"`
	IsNecessary *bool                                 `json:"IsNecessary,omitempty" xml:"IsNecessary,omitempty"`
	LifeSpan    *int32                                `json:"LifeSpan,omitempty" xml:"LifeSpan,omitempty"`
	Name        *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Question    []*string                             `json:"Question,omitempty" xml:"Question,omitempty" type:"Repeated"`
	SlotId      *string                               `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Tags        []*DescribeIntentResponseBodySlotTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Value       *string                               `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIntentResponseBodySlot) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBodySlot) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodySlot) SetIsArray(v bool) *DescribeIntentResponseBodySlot {
	s.IsArray = &v
	return s
}

func (s *DescribeIntentResponseBodySlot) SetIsNecessary(v bool) *DescribeIntentResponseBodySlot {
	s.IsNecessary = &v
	return s
}

func (s *DescribeIntentResponseBodySlot) SetLifeSpan(v int32) *DescribeIntentResponseBodySlot {
	s.LifeSpan = &v
	return s
}

func (s *DescribeIntentResponseBodySlot) SetName(v string) *DescribeIntentResponseBodySlot {
	s.Name = &v
	return s
}

func (s *DescribeIntentResponseBodySlot) SetQuestion(v []*string) *DescribeIntentResponseBodySlot {
	s.Question = v
	return s
}

func (s *DescribeIntentResponseBodySlot) SetSlotId(v string) *DescribeIntentResponseBodySlot {
	s.SlotId = &v
	return s
}

func (s *DescribeIntentResponseBodySlot) SetTags(v []*DescribeIntentResponseBodySlotTags) *DescribeIntentResponseBodySlot {
	s.Tags = v
	return s
}

func (s *DescribeIntentResponseBodySlot) SetValue(v string) *DescribeIntentResponseBodySlot {
	s.Value = &v
	return s
}

type DescribeIntentResponseBodySlotTags struct {
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIntentResponseBodySlotTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBodySlotTags) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodySlotTags) SetUserSayId(v string) *DescribeIntentResponseBodySlotTags {
	s.UserSayId = &v
	return s
}

func (s *DescribeIntentResponseBodySlotTags) SetValue(v string) *DescribeIntentResponseBodySlotTags {
	s.Value = &v
	return s
}

type DescribeIntentResponseBodyUserSay struct {
	Data      []*DescribeIntentResponseBodyUserSayData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Strict    *bool                                    `json:"Strict,omitempty" xml:"Strict,omitempty"`
	UserSayId *string                                  `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s DescribeIntentResponseBodyUserSay) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBodyUserSay) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodyUserSay) SetData(v []*DescribeIntentResponseBodyUserSayData) *DescribeIntentResponseBodyUserSay {
	s.Data = v
	return s
}

func (s *DescribeIntentResponseBodyUserSay) SetStrict(v bool) *DescribeIntentResponseBodyUserSay {
	s.Strict = &v
	return s
}

func (s *DescribeIntentResponseBodyUserSay) SetUserSayId(v string) *DescribeIntentResponseBodyUserSay {
	s.UserSayId = &v
	return s
}

type DescribeIntentResponseBodyUserSayData struct {
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Text   *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DescribeIntentResponseBodyUserSayData) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBodyUserSayData) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodyUserSayData) SetSlotId(v string) *DescribeIntentResponseBodyUserSayData {
	s.SlotId = &v
	return s
}

func (s *DescribeIntentResponseBodyUserSayData) SetText(v string) *DescribeIntentResponseBodyUserSayData {
	s.Text = &v
	return s
}

type DescribeIntentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponse) SetHeaders(v map[string]*string) *DescribeIntentResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntentResponse) SetStatusCode(v int32) *DescribeIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIntentResponse) SetBody(v *DescribeIntentResponseBody) *DescribeIntentResponse {
	s.Body = v
	return s
}

type DescribeKnowledgeRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// FAQ ID
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DescribeKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeRequest) SetAgentKey(v string) *DescribeKnowledgeRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeKnowledgeRequest) SetKnowledgeId(v int64) *DescribeKnowledgeRequest {
	s.KnowledgeId = &v
	return s
}

type DescribeKnowledgeResponseBody struct {
	CategoryId     *int64    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CoreWords      []*string `json:"CoreWords,omitempty" xml:"CoreWords,omitempty" type:"Repeated"`
	CreateTime     *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserName *string   `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	EndDate        *string   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	KeyWords       []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// FAQ ID
	KnowledgeId     *int64                                       `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	KnowledgeStatus *int32                                       `json:"KnowledgeStatus,omitempty" xml:"KnowledgeStatus,omitempty"`
	KnowledgeTitle  *string                                      `json:"KnowledgeTitle,omitempty" xml:"KnowledgeTitle,omitempty"`
	KnowledgeType   *int32                                       `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	ModifyTime      *string                                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserName  *string                                      `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Outlines        []*DescribeKnowledgeResponseBodyOutlines     `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SimQuestions    []*DescribeKnowledgeResponseBodySimQuestions `json:"SimQuestions,omitempty" xml:"SimQuestions,omitempty" type:"Repeated"`
	Solutions       []*DescribeKnowledgeResponseBodySolutions    `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	StartDate       *string                                      `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Version         *int32                                       `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeResponseBody) SetCategoryId(v int64) *DescribeKnowledgeResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetCoreWords(v []*string) *DescribeKnowledgeResponseBody {
	s.CoreWords = v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetCreateTime(v string) *DescribeKnowledgeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetCreateUserName(v string) *DescribeKnowledgeResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetEndDate(v string) *DescribeKnowledgeResponseBody {
	s.EndDate = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetKeyWords(v []*string) *DescribeKnowledgeResponseBody {
	s.KeyWords = v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetKnowledgeId(v int64) *DescribeKnowledgeResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetKnowledgeStatus(v int32) *DescribeKnowledgeResponseBody {
	s.KnowledgeStatus = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetKnowledgeTitle(v string) *DescribeKnowledgeResponseBody {
	s.KnowledgeTitle = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetKnowledgeType(v int32) *DescribeKnowledgeResponseBody {
	s.KnowledgeType = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetModifyTime(v string) *DescribeKnowledgeResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetModifyUserName(v string) *DescribeKnowledgeResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetOutlines(v []*DescribeKnowledgeResponseBodyOutlines) *DescribeKnowledgeResponseBody {
	s.Outlines = v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetRequestId(v string) *DescribeKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetSimQuestions(v []*DescribeKnowledgeResponseBodySimQuestions) *DescribeKnowledgeResponseBody {
	s.SimQuestions = v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetSolutions(v []*DescribeKnowledgeResponseBodySolutions) *DescribeKnowledgeResponseBody {
	s.Solutions = v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetStartDate(v string) *DescribeKnowledgeResponseBody {
	s.StartDate = &v
	return s
}

func (s *DescribeKnowledgeResponseBody) SetVersion(v int32) *DescribeKnowledgeResponseBody {
	s.Version = &v
	return s
}

type DescribeKnowledgeResponseBodyOutlines struct {
	// FAQ ID
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	OutlineId   *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeKnowledgeResponseBodyOutlines) String() string {
	return tea.Prettify(s)
}

func (s DescribeKnowledgeResponseBodyOutlines) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeResponseBodyOutlines) SetKnowledgeId(v int64) *DescribeKnowledgeResponseBodyOutlines {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeKnowledgeResponseBodyOutlines) SetOutlineId(v int64) *DescribeKnowledgeResponseBodyOutlines {
	s.OutlineId = &v
	return s
}

func (s *DescribeKnowledgeResponseBodyOutlines) SetTitle(v string) *DescribeKnowledgeResponseBodyOutlines {
	s.Title = &v
	return s
}

type DescribeKnowledgeResponseBodySimQuestions struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime    *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SimQuestionId *int64  `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeKnowledgeResponseBodySimQuestions) String() string {
	return tea.Prettify(s)
}

func (s DescribeKnowledgeResponseBodySimQuestions) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeResponseBodySimQuestions) SetCreateTime(v string) *DescribeKnowledgeResponseBodySimQuestions {
	s.CreateTime = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySimQuestions) SetModifyTime(v string) *DescribeKnowledgeResponseBodySimQuestions {
	s.ModifyTime = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySimQuestions) SetSimQuestionId(v int64) *DescribeKnowledgeResponseBodySimQuestions {
	s.SimQuestionId = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySimQuestions) SetTitle(v string) *DescribeKnowledgeResponseBodySimQuestions {
	s.Title = &v
	return s
}

type DescribeKnowledgeResponseBodySolutions struct {
	Content        *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime     *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime     *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PerspectiveIds []*string `json:"PerspectiveIds,omitempty" xml:"PerspectiveIds,omitempty" type:"Repeated"`
	PlainText      *string   `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
	SolutionId     *int64    `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	Summary        *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s DescribeKnowledgeResponseBodySolutions) String() string {
	return tea.Prettify(s)
}

func (s DescribeKnowledgeResponseBodySolutions) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeResponseBodySolutions) SetContent(v string) *DescribeKnowledgeResponseBodySolutions {
	s.Content = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySolutions) SetCreateTime(v string) *DescribeKnowledgeResponseBodySolutions {
	s.CreateTime = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySolutions) SetModifyTime(v string) *DescribeKnowledgeResponseBodySolutions {
	s.ModifyTime = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySolutions) SetPerspectiveIds(v []*string) *DescribeKnowledgeResponseBodySolutions {
	s.PerspectiveIds = v
	return s
}

func (s *DescribeKnowledgeResponseBodySolutions) SetPlainText(v string) *DescribeKnowledgeResponseBodySolutions {
	s.PlainText = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySolutions) SetSolutionId(v int64) *DescribeKnowledgeResponseBodySolutions {
	s.SolutionId = &v
	return s
}

func (s *DescribeKnowledgeResponseBodySolutions) SetSummary(v string) *DescribeKnowledgeResponseBodySolutions {
	s.Summary = &v
	return s
}

type DescribeKnowledgeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeResponse) SetStatusCode(v int32) *DescribeKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeResponse) SetBody(v *DescribeKnowledgeResponseBody) *DescribeKnowledgeResponse {
	s.Body = v
	return s
}

type DescribePerspectiveRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s DescribePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveRequest) SetAgentKey(v string) *DescribePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribePerspectiveRequest) SetPerspectiveId(v string) *DescribePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type DescribePerspectiveResponseBody struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserName  *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	ModifyTime      *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserName  *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	PerspectiveId   *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SelfDefine      *bool   `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveResponseBody) SetCreateTime(v string) *DescribePerspectiveResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetCreateUserName(v string) *DescribePerspectiveResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetModifyTime(v string) *DescribePerspectiveResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetModifyUserName(v string) *DescribePerspectiveResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetName(v string) *DescribePerspectiveResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveCode(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveCode = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveId(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveId = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetRequestId(v string) *DescribePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetSelfDefine(v bool) *DescribePerspectiveResponseBody {
	s.SelfDefine = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetStatus(v int32) *DescribePerspectiveResponseBody {
	s.Status = &v
	return s
}

type DescribePerspectiveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveResponse) SetHeaders(v map[string]*string) *DescribePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *DescribePerspectiveResponse) SetStatusCode(v int32) *DescribePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePerspectiveResponse) SetBody(v *DescribePerspectiveResponseBody) *DescribePerspectiveResponse {
	s.Body = v
	return s
}

type DisableDialogFlowRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s DisableDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *DisableDialogFlowRequest) SetAgentKey(v string) *DisableDialogFlowRequest {
	s.AgentKey = &v
	return s
}

func (s *DisableDialogFlowRequest) SetDialogId(v int64) *DisableDialogFlowRequest {
	s.DialogId = &v
	return s
}

type DisableDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDialogFlowResponseBody) SetRequestId(v string) *DisableDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type DisableDialogFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *DisableDialogFlowResponse) SetHeaders(v map[string]*string) *DisableDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *DisableDialogFlowResponse) SetStatusCode(v int32) *DisableDialogFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDialogFlowResponse) SetBody(v *DisableDialogFlowResponseBody) *DisableDialogFlowResponse {
	s.Body = v
	return s
}

type DisableKnowledgeRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// FAQ ID
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DisableKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DisableKnowledgeRequest) SetAgentKey(v string) *DisableKnowledgeRequest {
	s.AgentKey = &v
	return s
}

func (s *DisableKnowledgeRequest) SetKnowledgeId(v int64) *DisableKnowledgeRequest {
	s.KnowledgeId = &v
	return s
}

type DisableKnowledgeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DisableKnowledgeResponseBody) SetRequestId(v string) *DisableKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type DisableKnowledgeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DisableKnowledgeResponse) SetHeaders(v map[string]*string) *DisableKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DisableKnowledgeResponse) SetStatusCode(v int32) *DisableKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableKnowledgeResponse) SetBody(v *DisableKnowledgeResponseBody) *DisableKnowledgeResponse {
	s.Body = v
	return s
}

type FeedbackRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Feedback   *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageId  *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackRequest) GoString() string {
	return s.String()
}

func (s *FeedbackRequest) SetAgentKey(v string) *FeedbackRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackRequest) SetFeedback(v string) *FeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *FeedbackRequest) SetInstanceId(v string) *FeedbackRequest {
	s.InstanceId = &v
	return s
}

func (s *FeedbackRequest) SetMessageId(v string) *FeedbackRequest {
	s.MessageId = &v
	return s
}

func (s *FeedbackRequest) SetSessionId(v string) *FeedbackRequest {
	s.SessionId = &v
	return s
}

type FeedbackResponseBody struct {
	Feedback   *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	HttpStatus *int64  `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	MessageId  *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackResponseBody) SetFeedback(v string) *FeedbackResponseBody {
	s.Feedback = &v
	return s
}

func (s *FeedbackResponseBody) SetHttpStatus(v int64) *FeedbackResponseBody {
	s.HttpStatus = &v
	return s
}

func (s *FeedbackResponseBody) SetMessageId(v string) *FeedbackResponseBody {
	s.MessageId = &v
	return s
}

func (s *FeedbackResponseBody) SetRequestId(v string) *FeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *FeedbackResponseBody) SetSuccess(v bool) *FeedbackResponseBody {
	s.Success = &v
	return s
}

type FeedbackResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackResponse) GoString() string {
	return s.String()
}

func (s *FeedbackResponse) SetHeaders(v map[string]*string) *FeedbackResponse {
	s.Headers = v
	return s
}

func (s *FeedbackResponse) SetStatusCode(v int32) *FeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackResponse) SetBody(v *FeedbackResponseBody) *FeedbackResponse {
	s.Body = v
	return s
}

type GetAsyncResultRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAsyncResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncResultRequest) SetAgentKey(v string) *GetAsyncResultRequest {
	s.AgentKey = &v
	return s
}

func (s *GetAsyncResultRequest) SetTaskId(v string) *GetAsyncResultRequest {
	s.TaskId = &v
	return s
}

type GetAsyncResultResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsyncResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponseBody) SetCode(v int32) *GetAsyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetData(v string) *GetAsyncResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetMessage(v string) *GetAsyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetRequestId(v string) *GetAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetSuccess(v bool) *GetAsyncResultResponseBody {
	s.Success = &v
	return s
}

type GetAsyncResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponse) SetHeaders(v map[string]*string) *GetAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncResultResponse) SetStatusCode(v int32) *GetAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncResultResponse) SetBody(v *GetAsyncResultResponseBody) *GetAsyncResultResponse {
	s.Body = v
	return s
}

type GetBotChatDataRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBotChatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotChatDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotChatDataRequest) SetAgentKey(v string) *GetBotChatDataRequest {
	s.AgentKey = &v
	return s
}

func (s *GetBotChatDataRequest) SetEndTime(v string) *GetBotChatDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotChatDataRequest) SetRobotInstanceId(v string) *GetBotChatDataRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *GetBotChatDataRequest) SetStartTime(v string) *GetBotChatDataRequest {
	s.StartTime = &v
	return s
}

type GetBotChatDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBotChatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotChatDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotChatDataResponseBody) SetCostTime(v string) *GetBotChatDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotChatDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotChatDataResponseBody {
	s.Datas = v
	return s
}

func (s *GetBotChatDataResponseBody) SetRequestId(v string) *GetBotChatDataResponseBody {
	s.RequestId = &v
	return s
}

type GetBotChatDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBotChatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBotChatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotChatDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotChatDataResponse) SetHeaders(v map[string]*string) *GetBotChatDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotChatDataResponse) SetStatusCode(v int32) *GetBotChatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBotChatDataResponse) SetBody(v *GetBotChatDataResponseBody) *GetBotChatDataResponse {
	s.Body = v
	return s
}

type GetBotDsStatDataRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBotDsStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotDsStatDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotDsStatDataRequest) SetAgentKey(v string) *GetBotDsStatDataRequest {
	s.AgentKey = &v
	return s
}

func (s *GetBotDsStatDataRequest) SetEndTime(v string) *GetBotDsStatDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotDsStatDataRequest) SetRobotInstanceId(v string) *GetBotDsStatDataRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *GetBotDsStatDataRequest) SetStartTime(v string) *GetBotDsStatDataRequest {
	s.StartTime = &v
	return s
}

type GetBotDsStatDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBotDsStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotDsStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotDsStatDataResponseBody) SetCostTime(v string) *GetBotDsStatDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotDsStatDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotDsStatDataResponseBody {
	s.Datas = v
	return s
}

func (s *GetBotDsStatDataResponseBody) SetRequestId(v string) *GetBotDsStatDataResponseBody {
	s.RequestId = &v
	return s
}

type GetBotDsStatDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBotDsStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBotDsStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotDsStatDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotDsStatDataResponse) SetHeaders(v map[string]*string) *GetBotDsStatDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotDsStatDataResponse) SetStatusCode(v int32) *GetBotDsStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBotDsStatDataResponse) SetBody(v *GetBotDsStatDataResponseBody) *GetBotDsStatDataResponse {
	s.Body = v
	return s
}

type GetBotKnowledgeStatDataRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBotKnowledgeStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotKnowledgeStatDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotKnowledgeStatDataRequest) SetAgentKey(v string) *GetBotKnowledgeStatDataRequest {
	s.AgentKey = &v
	return s
}

func (s *GetBotKnowledgeStatDataRequest) SetEndTime(v string) *GetBotKnowledgeStatDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotKnowledgeStatDataRequest) SetRobotInstanceId(v string) *GetBotKnowledgeStatDataRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *GetBotKnowledgeStatDataRequest) SetStartTime(v string) *GetBotKnowledgeStatDataRequest {
	s.StartTime = &v
	return s
}

type GetBotKnowledgeStatDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBotKnowledgeStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotKnowledgeStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotKnowledgeStatDataResponseBody) SetCostTime(v string) *GetBotKnowledgeStatDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotKnowledgeStatDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotKnowledgeStatDataResponseBody {
	s.Datas = v
	return s
}

func (s *GetBotKnowledgeStatDataResponseBody) SetRequestId(v string) *GetBotKnowledgeStatDataResponseBody {
	s.RequestId = &v
	return s
}

type GetBotKnowledgeStatDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBotKnowledgeStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBotKnowledgeStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotKnowledgeStatDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotKnowledgeStatDataResponse) SetHeaders(v map[string]*string) *GetBotKnowledgeStatDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotKnowledgeStatDataResponse) SetStatusCode(v int32) *GetBotKnowledgeStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBotKnowledgeStatDataResponse) SetBody(v *GetBotKnowledgeStatDataResponseBody) *GetBotKnowledgeStatDataResponse {
	s.Body = v
	return s
}

type GetBotSessionDataRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBotSessionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotSessionDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataRequest) SetAgentKey(v string) *GetBotSessionDataRequest {
	s.AgentKey = &v
	return s
}

func (s *GetBotSessionDataRequest) SetEndTime(v string) *GetBotSessionDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotSessionDataRequest) SetRobotInstanceId(v string) *GetBotSessionDataRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *GetBotSessionDataRequest) SetStartTime(v string) *GetBotSessionDataRequest {
	s.StartTime = &v
	return s
}

type GetBotSessionDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBotSessionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotSessionDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataResponseBody) SetCostTime(v string) *GetBotSessionDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotSessionDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotSessionDataResponseBody {
	s.Datas = v
	return s
}

func (s *GetBotSessionDataResponseBody) SetRequestId(v string) *GetBotSessionDataResponseBody {
	s.RequestId = &v
	return s
}

type GetBotSessionDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBotSessionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBotSessionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotSessionDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataResponse) SetHeaders(v map[string]*string) *GetBotSessionDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotSessionDataResponse) SetStatusCode(v int32) *GetBotSessionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBotSessionDataResponse) SetBody(v *GetBotSessionDataResponseBody) *GetBotSessionDataResponse {
	s.Body = v
	return s
}

type GetConversationListRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SenderId   *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetConversationListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationListRequest) GoString() string {
	return s.String()
}

func (s *GetConversationListRequest) SetAgentKey(v string) *GetConversationListRequest {
	s.AgentKey = &v
	return s
}

func (s *GetConversationListRequest) SetEndDate(v string) *GetConversationListRequest {
	s.EndDate = &v
	return s
}

func (s *GetConversationListRequest) SetInstanceId(v string) *GetConversationListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConversationListRequest) SetPageNumber(v string) *GetConversationListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetConversationListRequest) SetPageSize(v string) *GetConversationListRequest {
	s.PageSize = &v
	return s
}

func (s *GetConversationListRequest) SetSenderId(v string) *GetConversationListRequest {
	s.SenderId = &v
	return s
}

func (s *GetConversationListRequest) SetSessionId(v string) *GetConversationListRequest {
	s.SessionId = &v
	return s
}

func (s *GetConversationListRequest) SetStartDate(v string) *GetConversationListRequest {
	s.StartDate = &v
	return s
}

type GetConversationListResponseBody struct {
	Messages    []map[string]*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	PageNumber  *int64               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCounts *int64               `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s GetConversationListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationListResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationListResponseBody) SetMessages(v []map[string]*string) *GetConversationListResponseBody {
	s.Messages = v
	return s
}

func (s *GetConversationListResponseBody) SetPageNumber(v int64) *GetConversationListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetConversationListResponseBody) SetPageSize(v int64) *GetConversationListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetConversationListResponseBody) SetRequestId(v string) *GetConversationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConversationListResponseBody) SetTotalCounts(v int64) *GetConversationListResponseBody {
	s.TotalCounts = &v
	return s
}

type GetConversationListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationListResponse) GoString() string {
	return s.String()
}

func (s *GetConversationListResponse) SetHeaders(v map[string]*string) *GetConversationListResponse {
	s.Headers = v
	return s
}

func (s *GetConversationListResponse) SetStatusCode(v int32) *GetConversationListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationListResponse) SetBody(v *GetConversationListResponseBody) *GetConversationListResponse {
	s.Body = v
	return s
}

type ListBotChatHistorysRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotChatHistorysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotChatHistorysRequest) GoString() string {
	return s.String()
}

func (s *ListBotChatHistorysRequest) SetAgentKey(v string) *ListBotChatHistorysRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetEndTime(v string) *ListBotChatHistorysRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetLimit(v int32) *ListBotChatHistorysRequest {
	s.Limit = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetRobotInstanceId(v string) *ListBotChatHistorysRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetStartTime(v string) *ListBotChatHistorysRequest {
	s.StartTime = &v
	return s
}

type ListBotChatHistorysResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotChatHistorysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotChatHistorysResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotChatHistorysResponseBody) SetCostTime(v string) *ListBotChatHistorysResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotChatHistorysResponseBody) SetDatas(v []map[string]interface{}) *ListBotChatHistorysResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotChatHistorysResponseBody) SetRequestId(v string) *ListBotChatHistorysResponseBody {
	s.RequestId = &v
	return s
}

type ListBotChatHistorysResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotChatHistorysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotChatHistorysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotChatHistorysResponse) GoString() string {
	return s.String()
}

func (s *ListBotChatHistorysResponse) SetHeaders(v map[string]*string) *ListBotChatHistorysResponse {
	s.Headers = v
	return s
}

func (s *ListBotChatHistorysResponse) SetStatusCode(v int32) *ListBotChatHistorysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotChatHistorysResponse) SetBody(v *ListBotChatHistorysResponseBody) *ListBotChatHistorysResponse {
	s.Body = v
	return s
}

type ListBotColdDsDatasRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotColdDsDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdDsDatasRequest) GoString() string {
	return s.String()
}

func (s *ListBotColdDsDatasRequest) SetAgentKey(v string) *ListBotColdDsDatasRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetEndTime(v string) *ListBotColdDsDatasRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetLimit(v int32) *ListBotColdDsDatasRequest {
	s.Limit = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetRobotInstanceId(v string) *ListBotColdDsDatasRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetStartTime(v string) *ListBotColdDsDatasRequest {
	s.StartTime = &v
	return s
}

type ListBotColdDsDatasResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotColdDsDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdDsDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotColdDsDatasResponseBody) SetCostTime(v string) *ListBotColdDsDatasResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotColdDsDatasResponseBody) SetDatas(v []map[string]interface{}) *ListBotColdDsDatasResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotColdDsDatasResponseBody) SetRequestId(v string) *ListBotColdDsDatasResponseBody {
	s.RequestId = &v
	return s
}

type ListBotColdDsDatasResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotColdDsDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotColdDsDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdDsDatasResponse) GoString() string {
	return s.String()
}

func (s *ListBotColdDsDatasResponse) SetHeaders(v map[string]*string) *ListBotColdDsDatasResponse {
	s.Headers = v
	return s
}

func (s *ListBotColdDsDatasResponse) SetStatusCode(v int32) *ListBotColdDsDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotColdDsDatasResponse) SetBody(v *ListBotColdDsDatasResponseBody) *ListBotColdDsDatasResponse {
	s.Body = v
	return s
}

type ListBotColdKnowledgesRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotColdKnowledgesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdKnowledgesRequest) GoString() string {
	return s.String()
}

func (s *ListBotColdKnowledgesRequest) SetAgentKey(v string) *ListBotColdKnowledgesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetEndTime(v string) *ListBotColdKnowledgesRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetLimit(v int32) *ListBotColdKnowledgesRequest {
	s.Limit = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetRobotInstanceId(v string) *ListBotColdKnowledgesRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetStartTime(v string) *ListBotColdKnowledgesRequest {
	s.StartTime = &v
	return s
}

type ListBotColdKnowledgesResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotColdKnowledgesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdKnowledgesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotColdKnowledgesResponseBody) SetCostTime(v string) *ListBotColdKnowledgesResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotColdKnowledgesResponseBody) SetDatas(v []map[string]interface{}) *ListBotColdKnowledgesResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotColdKnowledgesResponseBody) SetRequestId(v string) *ListBotColdKnowledgesResponseBody {
	s.RequestId = &v
	return s
}

type ListBotColdKnowledgesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotColdKnowledgesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotColdKnowledgesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdKnowledgesResponse) GoString() string {
	return s.String()
}

func (s *ListBotColdKnowledgesResponse) SetHeaders(v map[string]*string) *ListBotColdKnowledgesResponse {
	s.Headers = v
	return s
}

func (s *ListBotColdKnowledgesResponse) SetStatusCode(v int32) *ListBotColdKnowledgesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotColdKnowledgesResponse) SetBody(v *ListBotColdKnowledgesResponseBody) *ListBotColdKnowledgesResponse {
	s.Body = v
	return s
}

type ListBotDsDetailsRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotDsDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotDsDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListBotDsDetailsRequest) SetAgentKey(v string) *ListBotDsDetailsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetEndTime(v string) *ListBotDsDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetLimit(v int32) *ListBotDsDetailsRequest {
	s.Limit = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetRobotInstanceId(v string) *ListBotDsDetailsRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetStartTime(v string) *ListBotDsDetailsRequest {
	s.StartTime = &v
	return s
}

type ListBotDsDetailsResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotDsDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotDsDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotDsDetailsResponseBody) SetCostTime(v string) *ListBotDsDetailsResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotDsDetailsResponseBody) SetDatas(v []map[string]interface{}) *ListBotDsDetailsResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotDsDetailsResponseBody) SetRequestId(v string) *ListBotDsDetailsResponseBody {
	s.RequestId = &v
	return s
}

type ListBotDsDetailsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotDsDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotDsDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotDsDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListBotDsDetailsResponse) SetHeaders(v map[string]*string) *ListBotDsDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListBotDsDetailsResponse) SetStatusCode(v int32) *ListBotDsDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotDsDetailsResponse) SetBody(v *ListBotDsDetailsResponseBody) *ListBotDsDetailsResponse {
	s.Body = v
	return s
}

type ListBotHotDsDatasRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotHotDsDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotDsDatasRequest) GoString() string {
	return s.String()
}

func (s *ListBotHotDsDatasRequest) SetAgentKey(v string) *ListBotHotDsDatasRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetEndTime(v string) *ListBotHotDsDatasRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetLimit(v int32) *ListBotHotDsDatasRequest {
	s.Limit = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetRobotInstanceId(v string) *ListBotHotDsDatasRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetStartTime(v string) *ListBotHotDsDatasRequest {
	s.StartTime = &v
	return s
}

type ListBotHotDsDatasResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotHotDsDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotDsDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotHotDsDatasResponseBody) SetCostTime(v string) *ListBotHotDsDatasResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotHotDsDatasResponseBody) SetDatas(v []map[string]interface{}) *ListBotHotDsDatasResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotHotDsDatasResponseBody) SetRequestId(v string) *ListBotHotDsDatasResponseBody {
	s.RequestId = &v
	return s
}

type ListBotHotDsDatasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotHotDsDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotHotDsDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotDsDatasResponse) GoString() string {
	return s.String()
}

func (s *ListBotHotDsDatasResponse) SetHeaders(v map[string]*string) *ListBotHotDsDatasResponse {
	s.Headers = v
	return s
}

func (s *ListBotHotDsDatasResponse) SetStatusCode(v int32) *ListBotHotDsDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotHotDsDatasResponse) SetBody(v *ListBotHotDsDatasResponseBody) *ListBotHotDsDatasResponse {
	s.Body = v
	return s
}

type ListBotHotKnowledgesRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotHotKnowledgesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotKnowledgesRequest) GoString() string {
	return s.String()
}

func (s *ListBotHotKnowledgesRequest) SetAgentKey(v string) *ListBotHotKnowledgesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetEndTime(v string) *ListBotHotKnowledgesRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetLimit(v int32) *ListBotHotKnowledgesRequest {
	s.Limit = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetRobotInstanceId(v string) *ListBotHotKnowledgesRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetStartTime(v string) *ListBotHotKnowledgesRequest {
	s.StartTime = &v
	return s
}

type ListBotHotKnowledgesResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotHotKnowledgesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotKnowledgesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotHotKnowledgesResponseBody) SetCostTime(v string) *ListBotHotKnowledgesResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotHotKnowledgesResponseBody) SetDatas(v []map[string]interface{}) *ListBotHotKnowledgesResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotHotKnowledgesResponseBody) SetRequestId(v string) *ListBotHotKnowledgesResponseBody {
	s.RequestId = &v
	return s
}

type ListBotHotKnowledgesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotHotKnowledgesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotHotKnowledgesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotKnowledgesResponse) GoString() string {
	return s.String()
}

func (s *ListBotHotKnowledgesResponse) SetHeaders(v map[string]*string) *ListBotHotKnowledgesResponse {
	s.Headers = v
	return s
}

func (s *ListBotHotKnowledgesResponse) SetStatusCode(v int32) *ListBotHotKnowledgesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotHotKnowledgesResponse) SetBody(v *ListBotHotKnowledgesResponseBody) *ListBotHotKnowledgesResponse {
	s.Body = v
	return s
}

type ListBotKnowledgeDetailsRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit           *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotKnowledgeDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotKnowledgeDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListBotKnowledgeDetailsRequest) SetAgentKey(v string) *ListBotKnowledgeDetailsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetEndTime(v string) *ListBotKnowledgeDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetLimit(v string) *ListBotKnowledgeDetailsRequest {
	s.Limit = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetRobotInstanceId(v string) *ListBotKnowledgeDetailsRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetStartTime(v string) *ListBotKnowledgeDetailsRequest {
	s.StartTime = &v
	return s
}

type ListBotKnowledgeDetailsResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotKnowledgeDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotKnowledgeDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotKnowledgeDetailsResponseBody) SetCostTime(v string) *ListBotKnowledgeDetailsResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotKnowledgeDetailsResponseBody) SetDatas(v []map[string]interface{}) *ListBotKnowledgeDetailsResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotKnowledgeDetailsResponseBody) SetRequestId(v string) *ListBotKnowledgeDetailsResponseBody {
	s.RequestId = &v
	return s
}

type ListBotKnowledgeDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotKnowledgeDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotKnowledgeDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotKnowledgeDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListBotKnowledgeDetailsResponse) SetHeaders(v map[string]*string) *ListBotKnowledgeDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListBotKnowledgeDetailsResponse) SetStatusCode(v int32) *ListBotKnowledgeDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotKnowledgeDetailsResponse) SetBody(v *ListBotKnowledgeDetailsResponseBody) *ListBotKnowledgeDetailsResponse {
	s.Body = v
	return s
}

type ListBotReceptionDetailDatasRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBotReceptionDetailDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotReceptionDetailDatasRequest) GoString() string {
	return s.String()
}

func (s *ListBotReceptionDetailDatasRequest) SetAgentKey(v string) *ListBotReceptionDetailDatasRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetEndTime(v string) *ListBotReceptionDetailDatasRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetRobotInstanceId(v string) *ListBotReceptionDetailDatasRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetStartTime(v string) *ListBotReceptionDetailDatasRequest {
	s.StartTime = &v
	return s
}

type ListBotReceptionDetailDatasResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBotReceptionDetailDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotReceptionDetailDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotReceptionDetailDatasResponseBody) SetCostTime(v string) *ListBotReceptionDetailDatasResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotReceptionDetailDatasResponseBody) SetDatas(v []map[string]interface{}) *ListBotReceptionDetailDatasResponseBody {
	s.Datas = v
	return s
}

func (s *ListBotReceptionDetailDatasResponseBody) SetRequestId(v string) *ListBotReceptionDetailDatasResponseBody {
	s.RequestId = &v
	return s
}

type ListBotReceptionDetailDatasResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotReceptionDetailDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotReceptionDetailDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotReceptionDetailDatasResponse) GoString() string {
	return s.String()
}

func (s *ListBotReceptionDetailDatasResponse) SetHeaders(v map[string]*string) *ListBotReceptionDetailDatasResponse {
	s.Headers = v
	return s
}

func (s *ListBotReceptionDetailDatasResponse) SetStatusCode(v int32) *ListBotReceptionDetailDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotReceptionDetailDatasResponse) SetBody(v *ListBotReceptionDetailDatasResponseBody) *ListBotReceptionDetailDatasResponse {
	s.Body = v
	return s
}

type ListConversationLogsRequest struct {
	AgentKey  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ListConversationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConversationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationLogsRequest) SetAgentKey(v string) *ListConversationLogsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListConversationLogsRequest) SetSessionId(v string) *ListConversationLogsRequest {
	s.SessionId = &v
	return s
}

type ListConversationLogsResponseBody struct {
	ChatLogs  []map[string]interface{} `json:"ChatLogs,omitempty" xml:"ChatLogs,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rounds    *int64                   `json:"Rounds,omitempty" xml:"Rounds,omitempty"`
}

func (s ListConversationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConversationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationLogsResponseBody) SetChatLogs(v []map[string]interface{}) *ListConversationLogsResponseBody {
	s.ChatLogs = v
	return s
}

func (s *ListConversationLogsResponseBody) SetRequestId(v string) *ListConversationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConversationLogsResponseBody) SetRounds(v int64) *ListConversationLogsResponseBody {
	s.Rounds = &v
	return s
}

type ListConversationLogsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConversationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConversationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConversationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListConversationLogsResponse) SetHeaders(v map[string]*string) *ListConversationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListConversationLogsResponse) SetStatusCode(v int32) *ListConversationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConversationLogsResponse) SetBody(v *ListConversationLogsResponseBody) *ListConversationLogsResponse {
	s.Body = v
	return s
}

type MoveKnowledgeCategoryRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryId  *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s MoveKnowledgeCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveKnowledgeCategoryRequest) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeCategoryRequest) SetAgentKey(v string) *MoveKnowledgeCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *MoveKnowledgeCategoryRequest) SetCategoryId(v int64) *MoveKnowledgeCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *MoveKnowledgeCategoryRequest) SetKnowledgeId(v int64) *MoveKnowledgeCategoryRequest {
	s.KnowledgeId = &v
	return s
}

type MoveKnowledgeCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveKnowledgeCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveKnowledgeCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeCategoryResponseBody) SetRequestId(v string) *MoveKnowledgeCategoryResponseBody {
	s.RequestId = &v
	return s
}

type MoveKnowledgeCategoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveKnowledgeCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveKnowledgeCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveKnowledgeCategoryResponse) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeCategoryResponse) SetHeaders(v map[string]*string) *MoveKnowledgeCategoryResponse {
	s.Headers = v
	return s
}

func (s *MoveKnowledgeCategoryResponse) SetStatusCode(v int32) *MoveKnowledgeCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveKnowledgeCategoryResponse) SetBody(v *MoveKnowledgeCategoryResponseBody) *MoveKnowledgeCategoryResponse {
	s.Body = v
	return s
}

type PublishDialogFlowRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s PublishDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *PublishDialogFlowRequest) SetAgentKey(v string) *PublishDialogFlowRequest {
	s.AgentKey = &v
	return s
}

func (s *PublishDialogFlowRequest) SetDialogId(v int64) *PublishDialogFlowRequest {
	s.DialogId = &v
	return s
}

type PublishDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDialogFlowResponseBody) SetRequestId(v string) *PublishDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type PublishDialogFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *PublishDialogFlowResponse) SetHeaders(v map[string]*string) *PublishDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *PublishDialogFlowResponse) SetStatusCode(v int32) *PublishDialogFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishDialogFlowResponse) SetBody(v *PublishDialogFlowResponseBody) *PublishDialogFlowResponse {
	s.Body = v
	return s
}

type PublishKnowledgeRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Async       *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s PublishKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *PublishKnowledgeRequest) SetAgentKey(v string) *PublishKnowledgeRequest {
	s.AgentKey = &v
	return s
}

func (s *PublishKnowledgeRequest) SetAsync(v bool) *PublishKnowledgeRequest {
	s.Async = &v
	return s
}

func (s *PublishKnowledgeRequest) SetKnowledgeId(v int64) *PublishKnowledgeRequest {
	s.KnowledgeId = &v
	return s
}

type PublishKnowledgeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *PublishKnowledgeResponseBody) SetRequestId(v string) *PublishKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type PublishKnowledgeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *PublishKnowledgeResponse) SetHeaders(v map[string]*string) *PublishKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *PublishKnowledgeResponse) SetStatusCode(v int32) *PublishKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishKnowledgeResponse) SetBody(v *PublishKnowledgeResponseBody) *PublishKnowledgeResponse {
	s.Body = v
	return s
}

type QueryBotsRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryBotsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsRequest) GoString() string {
	return s.String()
}

func (s *QueryBotsRequest) SetAgentKey(v string) *QueryBotsRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryBotsRequest) SetPageNumber(v int32) *QueryBotsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryBotsRequest) SetPageSize(v int32) *QueryBotsRequest {
	s.PageSize = &v
	return s
}

type QueryBotsResponseBody struct {
	Bots       []*QueryBotsResponseBodyBots `json:"Bots,omitempty" xml:"Bots,omitempty" type:"Repeated"`
	PageNumber *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryBotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBotsResponseBody) SetBots(v []*QueryBotsResponseBodyBots) *QueryBotsResponseBody {
	s.Bots = v
	return s
}

func (s *QueryBotsResponseBody) SetPageNumber(v int32) *QueryBotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryBotsResponseBody) SetPageSize(v int32) *QueryBotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBotsResponseBody) SetRequestId(v string) *QueryBotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBotsResponseBody) SetTotalCount(v int32) *QueryBotsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryBotsResponseBodyBots struct {
	Avatar       *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TimeZone     *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s QueryBotsResponseBodyBots) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsResponseBodyBots) GoString() string {
	return s.String()
}

func (s *QueryBotsResponseBodyBots) SetAvatar(v string) *QueryBotsResponseBodyBots {
	s.Avatar = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetCreateTime(v string) *QueryBotsResponseBodyBots {
	s.CreateTime = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetInstanceId(v string) *QueryBotsResponseBodyBots {
	s.InstanceId = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetIntroduction(v string) *QueryBotsResponseBodyBots {
	s.Introduction = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetLanguageCode(v string) *QueryBotsResponseBodyBots {
	s.LanguageCode = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetName(v string) *QueryBotsResponseBodyBots {
	s.Name = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetTimeZone(v string) *QueryBotsResponseBodyBots {
	s.TimeZone = &v
	return s
}

type QueryBotsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBotsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsResponse) GoString() string {
	return s.String()
}

func (s *QueryBotsResponse) SetHeaders(v map[string]*string) *QueryBotsResponse {
	s.Headers = v
	return s
}

func (s *QueryBotsResponse) SetStatusCode(v int32) *QueryBotsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBotsResponse) SetBody(v *QueryBotsResponseBody) *QueryBotsResponse {
	s.Body = v
	return s
}

type QueryCategoriesRequest struct {
	AgentKey         *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	KnowledgeType    *int64  `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	ShowChildrens    *bool   `json:"ShowChildrens,omitempty" xml:"ShowChildrens,omitempty"`
}

func (s QueryCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoriesRequest) GoString() string {
	return s.String()
}

func (s *QueryCategoriesRequest) SetAgentKey(v string) *QueryCategoriesRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryCategoriesRequest) SetKnowledgeType(v int64) *QueryCategoriesRequest {
	s.KnowledgeType = &v
	return s
}

func (s *QueryCategoriesRequest) SetParentCategoryId(v int64) *QueryCategoriesRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *QueryCategoriesRequest) SetShowChildrens(v bool) *QueryCategoriesRequest {
	s.ShowChildrens = &v
	return s
}

type QueryCategoriesResponseBody struct {
	Categories []*Children `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	RequestId  *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCategoriesResponseBody) SetCategories(v []*Children) *QueryCategoriesResponseBody {
	s.Categories = v
	return s
}

func (s *QueryCategoriesResponseBody) SetRequestId(v string) *QueryCategoriesResponseBody {
	s.RequestId = &v
	return s
}

type QueryCategoriesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoriesResponse) GoString() string {
	return s.String()
}

func (s *QueryCategoriesResponse) SetHeaders(v map[string]*string) *QueryCategoriesResponse {
	s.Headers = v
	return s
}

func (s *QueryCategoriesResponse) SetStatusCode(v int32) *QueryCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCategoriesResponse) SetBody(v *QueryCategoriesResponseBody) *QueryCategoriesResponse {
	s.Body = v
	return s
}

type QueryCoreWordsRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Synonym      *string `json:"Synonym,omitempty" xml:"Synonym,omitempty"`
}

func (s QueryCoreWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCoreWordsRequest) GoString() string {
	return s.String()
}

func (s *QueryCoreWordsRequest) SetAgentKey(v string) *QueryCoreWordsRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryCoreWordsRequest) SetCoreWordName(v string) *QueryCoreWordsRequest {
	s.CoreWordName = &v
	return s
}

func (s *QueryCoreWordsRequest) SetPageNumber(v int32) *QueryCoreWordsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCoreWordsRequest) SetPageSize(v int32) *QueryCoreWordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCoreWordsRequest) SetSynonym(v string) *QueryCoreWordsRequest {
	s.Synonym = &v
	return s
}

type QueryCoreWordsResponseBody struct {
	CoreWords  []*QueryCoreWordsResponseBodyCoreWords `json:"CoreWords,omitempty" xml:"CoreWords,omitempty" type:"Repeated"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCoreWordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCoreWordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCoreWordsResponseBody) SetCoreWords(v []*QueryCoreWordsResponseBodyCoreWords) *QueryCoreWordsResponseBody {
	s.CoreWords = v
	return s
}

func (s *QueryCoreWordsResponseBody) SetPageNumber(v int32) *QueryCoreWordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryCoreWordsResponseBody) SetPageSize(v int32) *QueryCoreWordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryCoreWordsResponseBody) SetRequestId(v string) *QueryCoreWordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCoreWordsResponseBody) SetTotalCount(v int32) *QueryCoreWordsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCoreWordsResponseBodyCoreWords struct {
	CoreWordCode *string   `json:"CoreWordCode,omitempty" xml:"CoreWordCode,omitempty"`
	CoreWordName *string   `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime   *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Synonyms     []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s QueryCoreWordsResponseBodyCoreWords) String() string {
	return tea.Prettify(s)
}

func (s QueryCoreWordsResponseBodyCoreWords) GoString() string {
	return s.String()
}

func (s *QueryCoreWordsResponseBodyCoreWords) SetCoreWordCode(v string) *QueryCoreWordsResponseBodyCoreWords {
	s.CoreWordCode = &v
	return s
}

func (s *QueryCoreWordsResponseBodyCoreWords) SetCoreWordName(v string) *QueryCoreWordsResponseBodyCoreWords {
	s.CoreWordName = &v
	return s
}

func (s *QueryCoreWordsResponseBodyCoreWords) SetCreateTime(v string) *QueryCoreWordsResponseBodyCoreWords {
	s.CreateTime = &v
	return s
}

func (s *QueryCoreWordsResponseBodyCoreWords) SetModifyTime(v string) *QueryCoreWordsResponseBodyCoreWords {
	s.ModifyTime = &v
	return s
}

func (s *QueryCoreWordsResponseBodyCoreWords) SetSynonyms(v []*string) *QueryCoreWordsResponseBodyCoreWords {
	s.Synonyms = v
	return s
}

type QueryCoreWordsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCoreWordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCoreWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCoreWordsResponse) GoString() string {
	return s.String()
}

func (s *QueryCoreWordsResponse) SetHeaders(v map[string]*string) *QueryCoreWordsResponse {
	s.Headers = v
	return s
}

func (s *QueryCoreWordsResponse) SetStatusCode(v int32) *QueryCoreWordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCoreWordsResponse) SetBody(v *QueryCoreWordsResponseBody) *QueryCoreWordsResponse {
	s.Body = v
	return s
}

type QueryDialogsRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogName *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryDialogsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsRequest) GoString() string {
	return s.String()
}

func (s *QueryDialogsRequest) SetAgentKey(v string) *QueryDialogsRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryDialogsRequest) SetDialogName(v string) *QueryDialogsRequest {
	s.DialogName = &v
	return s
}

func (s *QueryDialogsRequest) SetInstanceId(v string) *QueryDialogsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDialogsRequest) SetPageNumber(v int32) *QueryDialogsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryDialogsRequest) SetPageSize(v int32) *QueryDialogsRequest {
	s.PageSize = &v
	return s
}

type QueryDialogsResponseBody struct {
	Dialogs    []*QueryDialogsResponseBodyDialogs `json:"Dialogs,omitempty" xml:"Dialogs,omitempty" type:"Repeated"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryDialogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDialogsResponseBody) SetDialogs(v []*QueryDialogsResponseBodyDialogs) *QueryDialogsResponseBody {
	s.Dialogs = v
	return s
}

func (s *QueryDialogsResponseBody) SetPageNumber(v int32) *QueryDialogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDialogsResponseBody) SetPageSize(v int32) *QueryDialogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDialogsResponseBody) SetRequestId(v string) *QueryDialogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDialogsResponseBody) SetTotalCount(v int32) *QueryDialogsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryDialogsResponseBodyDialogs struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DialogId       *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	DialogName     *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	IsOnline       *bool   `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	IsSampleDialog *bool   `json:"IsSampleDialog,omitempty" xml:"IsSampleDialog,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserId   *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDialogsResponseBodyDialogs) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsResponseBodyDialogs) GoString() string {
	return s.String()
}

func (s *QueryDialogsResponseBodyDialogs) SetCreateTime(v string) *QueryDialogsResponseBodyDialogs {
	s.CreateTime = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetCreateUserId(v string) *QueryDialogsResponseBodyDialogs {
	s.CreateUserId = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetCreateUserName(v string) *QueryDialogsResponseBodyDialogs {
	s.CreateUserName = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetDescription(v string) *QueryDialogsResponseBodyDialogs {
	s.Description = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetDialogId(v int64) *QueryDialogsResponseBodyDialogs {
	s.DialogId = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetDialogName(v string) *QueryDialogsResponseBodyDialogs {
	s.DialogName = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetIsOnline(v bool) *QueryDialogsResponseBodyDialogs {
	s.IsOnline = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetIsSampleDialog(v bool) *QueryDialogsResponseBodyDialogs {
	s.IsSampleDialog = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetModifyTime(v string) *QueryDialogsResponseBodyDialogs {
	s.ModifyTime = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetModifyUserId(v string) *QueryDialogsResponseBodyDialogs {
	s.ModifyUserId = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetModifyUserName(v string) *QueryDialogsResponseBodyDialogs {
	s.ModifyUserName = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetStatus(v int32) *QueryDialogsResponseBodyDialogs {
	s.Status = &v
	return s
}

type QueryDialogsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDialogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDialogsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsResponse) GoString() string {
	return s.String()
}

func (s *QueryDialogsResponse) SetHeaders(v map[string]*string) *QueryDialogsResponse {
	s.Headers = v
	return s
}

func (s *QueryDialogsResponse) SetStatusCode(v int32) *QueryDialogsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDialogsResponse) SetBody(v *QueryDialogsResponseBody) *QueryDialogsResponse {
	s.Body = v
	return s
}

type QueryEntitiesRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId   *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEntitiesRequest) GoString() string {
	return s.String()
}

func (s *QueryEntitiesRequest) SetAgentKey(v string) *QueryEntitiesRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryEntitiesRequest) SetDialogId(v int64) *QueryEntitiesRequest {
	s.DialogId = &v
	return s
}

func (s *QueryEntitiesRequest) SetEntityName(v string) *QueryEntitiesRequest {
	s.EntityName = &v
	return s
}

func (s *QueryEntitiesRequest) SetPageNumber(v int32) *QueryEntitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryEntitiesRequest) SetPageSize(v int32) *QueryEntitiesRequest {
	s.PageSize = &v
	return s
}

type QueryEntitiesResponseBody struct {
	Entities   []*QueryEntitiesResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEntitiesResponseBody) SetEntities(v []*QueryEntitiesResponseBodyEntities) *QueryEntitiesResponseBody {
	s.Entities = v
	return s
}

func (s *QueryEntitiesResponseBody) SetPageNumber(v int32) *QueryEntitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryEntitiesResponseBody) SetPageSize(v int32) *QueryEntitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryEntitiesResponseBody) SetRequestId(v string) *QueryEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEntitiesResponseBody) SetTotalCount(v int32) *QueryEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type QueryEntitiesResponseBodyEntities struct {
	CreateTime     *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string                                     `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string                                     `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	EntityId       *int64                                      `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName     *string                                     `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType     *string                                     `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Members        []*QueryEntitiesResponseBodyEntitiesMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	ModifyTime     *string                                     `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserId   *string                                     `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string                                     `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Regex          *string                                     `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s QueryEntitiesResponseBodyEntities) String() string {
	return tea.Prettify(s)
}

func (s QueryEntitiesResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *QueryEntitiesResponseBodyEntities) SetCreateTime(v string) *QueryEntitiesResponseBodyEntities {
	s.CreateTime = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetCreateUserId(v string) *QueryEntitiesResponseBodyEntities {
	s.CreateUserId = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetCreateUserName(v string) *QueryEntitiesResponseBodyEntities {
	s.CreateUserName = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetEntityId(v int64) *QueryEntitiesResponseBodyEntities {
	s.EntityId = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetEntityName(v string) *QueryEntitiesResponseBodyEntities {
	s.EntityName = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetEntityType(v string) *QueryEntitiesResponseBodyEntities {
	s.EntityType = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetMembers(v []*QueryEntitiesResponseBodyEntitiesMembers) *QueryEntitiesResponseBodyEntities {
	s.Members = v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetModifyTime(v string) *QueryEntitiesResponseBodyEntities {
	s.ModifyTime = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetModifyUserId(v string) *QueryEntitiesResponseBodyEntities {
	s.ModifyUserId = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetModifyUserName(v string) *QueryEntitiesResponseBodyEntities {
	s.ModifyUserName = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntities) SetRegex(v string) *QueryEntitiesResponseBodyEntities {
	s.Regex = &v
	return s
}

type QueryEntitiesResponseBodyEntitiesMembers struct {
	MemberName *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s QueryEntitiesResponseBodyEntitiesMembers) String() string {
	return tea.Prettify(s)
}

func (s QueryEntitiesResponseBodyEntitiesMembers) GoString() string {
	return s.String()
}

func (s *QueryEntitiesResponseBodyEntitiesMembers) SetMemberName(v string) *QueryEntitiesResponseBodyEntitiesMembers {
	s.MemberName = &v
	return s
}

func (s *QueryEntitiesResponseBodyEntitiesMembers) SetSynonyms(v []*string) *QueryEntitiesResponseBodyEntitiesMembers {
	s.Synonyms = v
	return s
}

type QueryEntitiesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEntitiesResponse) GoString() string {
	return s.String()
}

func (s *QueryEntitiesResponse) SetHeaders(v map[string]*string) *QueryEntitiesResponse {
	s.Headers = v
	return s
}

func (s *QueryEntitiesResponse) SetStatusCode(v int32) *QueryEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEntitiesResponse) SetBody(v *QueryEntitiesResponseBody) *QueryEntitiesResponse {
	s.Body = v
	return s
}

type QueryIntentsRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId   *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryIntentsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsRequest) GoString() string {
	return s.String()
}

func (s *QueryIntentsRequest) SetAgentKey(v string) *QueryIntentsRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryIntentsRequest) SetDialogId(v int64) *QueryIntentsRequest {
	s.DialogId = &v
	return s
}

func (s *QueryIntentsRequest) SetInstanceId(v string) *QueryIntentsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryIntentsRequest) SetIntentName(v string) *QueryIntentsRequest {
	s.IntentName = &v
	return s
}

func (s *QueryIntentsRequest) SetPageNumber(v int32) *QueryIntentsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryIntentsRequest) SetPageSize(v int32) *QueryIntentsRequest {
	s.PageSize = &v
	return s
}

type QueryIntentsResponseBody struct {
	Intents    []*QueryIntentsResponseBodyIntents `json:"Intents,omitempty" xml:"Intents,omitempty" type:"Repeated"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryIntentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponseBody) SetIntents(v []*QueryIntentsResponseBodyIntents) *QueryIntentsResponseBody {
	s.Intents = v
	return s
}

func (s *QueryIntentsResponseBody) SetPageNumber(v int32) *QueryIntentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryIntentsResponseBody) SetPageSize(v int32) *QueryIntentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryIntentsResponseBody) SetRequestId(v string) *QueryIntentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIntentsResponseBody) SetTotalCount(v int32) *QueryIntentsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryIntentsResponseBodyIntents struct {
	CreateTime     *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string                                     `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string                                     `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	IntentId       *int64                                      `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	ModifyTime     *string                                     `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserId   *string                                     `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string                                     `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name           *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCheck      []*QueryIntentsResponseBodyIntentsRuleCheck `json:"RuleCheck,omitempty" xml:"RuleCheck,omitempty" type:"Repeated"`
	Slot           []*QueryIntentsResponseBodyIntentsSlot      `json:"Slot,omitempty" xml:"Slot,omitempty" type:"Repeated"`
	UserSay        []*QueryIntentsResponseBodyIntentsUserSay   `json:"UserSay,omitempty" xml:"UserSay,omitempty" type:"Repeated"`
}

func (s QueryIntentsResponseBodyIntents) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponseBodyIntents) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponseBodyIntents) SetCreateTime(v string) *QueryIntentsResponseBodyIntents {
	s.CreateTime = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetCreateUserId(v string) *QueryIntentsResponseBodyIntents {
	s.CreateUserId = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetCreateUserName(v string) *QueryIntentsResponseBodyIntents {
	s.CreateUserName = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetIntentId(v int64) *QueryIntentsResponseBodyIntents {
	s.IntentId = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetModifyTime(v string) *QueryIntentsResponseBodyIntents {
	s.ModifyTime = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetModifyUserId(v string) *QueryIntentsResponseBodyIntents {
	s.ModifyUserId = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetModifyUserName(v string) *QueryIntentsResponseBodyIntents {
	s.ModifyUserName = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetName(v string) *QueryIntentsResponseBodyIntents {
	s.Name = &v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetRuleCheck(v []*QueryIntentsResponseBodyIntentsRuleCheck) *QueryIntentsResponseBodyIntents {
	s.RuleCheck = v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetSlot(v []*QueryIntentsResponseBodyIntentsSlot) *QueryIntentsResponseBodyIntents {
	s.Slot = v
	return s
}

func (s *QueryIntentsResponseBodyIntents) SetUserSay(v []*QueryIntentsResponseBodyIntentsUserSay) *QueryIntentsResponseBodyIntents {
	s.UserSay = v
	return s
}

type QueryIntentsResponseBodyIntentsRuleCheck struct {
	Error   []*string `json:"Error,omitempty" xml:"Error,omitempty" type:"Repeated"`
	Strict  *bool     `json:"Strict,omitempty" xml:"Strict,omitempty"`
	Text    *string   `json:"Text,omitempty" xml:"Text,omitempty"`
	Warning []*string `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s QueryIntentsResponseBodyIntentsRuleCheck) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponseBodyIntentsRuleCheck) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponseBodyIntentsRuleCheck) SetError(v []*string) *QueryIntentsResponseBodyIntentsRuleCheck {
	s.Error = v
	return s
}

func (s *QueryIntentsResponseBodyIntentsRuleCheck) SetStrict(v bool) *QueryIntentsResponseBodyIntentsRuleCheck {
	s.Strict = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsRuleCheck) SetText(v string) *QueryIntentsResponseBodyIntentsRuleCheck {
	s.Text = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsRuleCheck) SetWarning(v []*string) *QueryIntentsResponseBodyIntentsRuleCheck {
	s.Warning = v
	return s
}

type QueryIntentsResponseBodyIntentsSlot struct {
	IsArray     *bool                                      `json:"IsArray,omitempty" xml:"IsArray,omitempty"`
	IsNecessary *bool                                      `json:"IsNecessary,omitempty" xml:"IsNecessary,omitempty"`
	LifeSpan    *int32                                     `json:"LifeSpan,omitempty" xml:"LifeSpan,omitempty"`
	Name        *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Question    []*string                                  `json:"Question,omitempty" xml:"Question,omitempty" type:"Repeated"`
	SlotId      *string                                    `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Tags        []*QueryIntentsResponseBodyIntentsSlotTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Value       *string                                    `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryIntentsResponseBodyIntentsSlot) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponseBodyIntentsSlot) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetIsArray(v bool) *QueryIntentsResponseBodyIntentsSlot {
	s.IsArray = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetIsNecessary(v bool) *QueryIntentsResponseBodyIntentsSlot {
	s.IsNecessary = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetLifeSpan(v int32) *QueryIntentsResponseBodyIntentsSlot {
	s.LifeSpan = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetName(v string) *QueryIntentsResponseBodyIntentsSlot {
	s.Name = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetQuestion(v []*string) *QueryIntentsResponseBodyIntentsSlot {
	s.Question = v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetSlotId(v string) *QueryIntentsResponseBodyIntentsSlot {
	s.SlotId = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetTags(v []*QueryIntentsResponseBodyIntentsSlotTags) *QueryIntentsResponseBodyIntentsSlot {
	s.Tags = v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlot) SetValue(v string) *QueryIntentsResponseBodyIntentsSlot {
	s.Value = &v
	return s
}

type QueryIntentsResponseBodyIntentsSlotTags struct {
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryIntentsResponseBodyIntentsSlotTags) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponseBodyIntentsSlotTags) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponseBodyIntentsSlotTags) SetUserSayId(v string) *QueryIntentsResponseBodyIntentsSlotTags {
	s.UserSayId = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsSlotTags) SetValue(v string) *QueryIntentsResponseBodyIntentsSlotTags {
	s.Value = &v
	return s
}

type QueryIntentsResponseBodyIntentsUserSay struct {
	Data      []*QueryIntentsResponseBodyIntentsUserSayData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Strict    *bool                                         `json:"Strict,omitempty" xml:"Strict,omitempty"`
	UserSayId *string                                       `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s QueryIntentsResponseBodyIntentsUserSay) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponseBodyIntentsUserSay) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponseBodyIntentsUserSay) SetData(v []*QueryIntentsResponseBodyIntentsUserSayData) *QueryIntentsResponseBodyIntentsUserSay {
	s.Data = v
	return s
}

func (s *QueryIntentsResponseBodyIntentsUserSay) SetStrict(v bool) *QueryIntentsResponseBodyIntentsUserSay {
	s.Strict = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsUserSay) SetUserSayId(v string) *QueryIntentsResponseBodyIntentsUserSay {
	s.UserSayId = &v
	return s
}

type QueryIntentsResponseBodyIntentsUserSayData struct {
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Text   *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QueryIntentsResponseBodyIntentsUserSayData) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponseBodyIntentsUserSayData) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponseBodyIntentsUserSayData) SetSlotId(v string) *QueryIntentsResponseBodyIntentsUserSayData {
	s.SlotId = &v
	return s
}

func (s *QueryIntentsResponseBodyIntentsUserSayData) SetText(v string) *QueryIntentsResponseBodyIntentsUserSayData {
	s.Text = &v
	return s
}

type QueryIntentsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIntentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIntentsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentsResponse) GoString() string {
	return s.String()
}

func (s *QueryIntentsResponse) SetHeaders(v map[string]*string) *QueryIntentsResponse {
	s.Headers = v
	return s
}

func (s *QueryIntentsResponse) SetStatusCode(v int32) *QueryIntentsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIntentsResponse) SetBody(v *QueryIntentsResponseBody) *QueryIntentsResponse {
	s.Body = v
	return s
}

type QueryKnowledgesRequest struct {
	AgentKey       *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryId     *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CoreWordName   *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	KnowledgeTitle *string `json:"KnowledgeTitle,omitempty" xml:"KnowledgeTitle,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryKnowledgesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryKnowledgesRequest) GoString() string {
	return s.String()
}

func (s *QueryKnowledgesRequest) SetAgentKey(v string) *QueryKnowledgesRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryKnowledgesRequest) SetCategoryId(v int64) *QueryKnowledgesRequest {
	s.CategoryId = &v
	return s
}

func (s *QueryKnowledgesRequest) SetCoreWordName(v string) *QueryKnowledgesRequest {
	s.CoreWordName = &v
	return s
}

func (s *QueryKnowledgesRequest) SetKnowledgeTitle(v string) *QueryKnowledgesRequest {
	s.KnowledgeTitle = &v
	return s
}

func (s *QueryKnowledgesRequest) SetPageNumber(v int32) *QueryKnowledgesRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryKnowledgesRequest) SetPageSize(v int32) *QueryKnowledgesRequest {
	s.PageSize = &v
	return s
}

type QueryKnowledgesResponseBody struct {
	Knowledges []*QueryKnowledgesResponseBodyKnowledges `json:"Knowledges,omitempty" xml:"Knowledges,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryKnowledgesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryKnowledgesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryKnowledgesResponseBody) SetKnowledges(v []*QueryKnowledgesResponseBodyKnowledges) *QueryKnowledgesResponseBody {
	s.Knowledges = v
	return s
}

func (s *QueryKnowledgesResponseBody) SetPageNumber(v int32) *QueryKnowledgesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryKnowledgesResponseBody) SetPageSize(v int32) *QueryKnowledgesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryKnowledgesResponseBody) SetRequestId(v string) *QueryKnowledgesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryKnowledgesResponseBody) SetTotalCount(v int32) *QueryKnowledgesResponseBody {
	s.TotalCount = &v
	return s
}

type QueryKnowledgesResponseBodyKnowledges struct {
	CategoryId     *int64    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CoreWords      []*string `json:"CoreWords,omitempty" xml:"CoreWords,omitempty" type:"Repeated"`
	CreateTime     *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserName *string   `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	EndDate        *string   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// FAQ ID
	KnowledgeId     *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	KnowledgeStatus *int32  `json:"KnowledgeStatus,omitempty" xml:"KnowledgeStatus,omitempty"`
	KnowledgeTitle  *string `json:"KnowledgeTitle,omitempty" xml:"KnowledgeTitle,omitempty"`
	ModifyTime      *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserName  *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	StartDate       *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryKnowledgesResponseBodyKnowledges) String() string {
	return tea.Prettify(s)
}

func (s QueryKnowledgesResponseBodyKnowledges) GoString() string {
	return s.String()
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetCategoryId(v int64) *QueryKnowledgesResponseBodyKnowledges {
	s.CategoryId = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetCoreWords(v []*string) *QueryKnowledgesResponseBodyKnowledges {
	s.CoreWords = v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetCreateTime(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.CreateTime = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetCreateUserName(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.CreateUserName = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetEndDate(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.EndDate = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetKnowledgeId(v int64) *QueryKnowledgesResponseBodyKnowledges {
	s.KnowledgeId = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetKnowledgeStatus(v int32) *QueryKnowledgesResponseBodyKnowledges {
	s.KnowledgeStatus = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetKnowledgeTitle(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.KnowledgeTitle = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetModifyTime(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.ModifyTime = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetModifyUserName(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.ModifyUserName = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetStartDate(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.StartDate = &v
	return s
}

func (s *QueryKnowledgesResponseBodyKnowledges) SetVersion(v string) *QueryKnowledgesResponseBodyKnowledges {
	s.Version = &v
	return s
}

type QueryKnowledgesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryKnowledgesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryKnowledgesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryKnowledgesResponse) GoString() string {
	return s.String()
}

func (s *QueryKnowledgesResponse) SetHeaders(v map[string]*string) *QueryKnowledgesResponse {
	s.Headers = v
	return s
}

func (s *QueryKnowledgesResponse) SetStatusCode(v int32) *QueryKnowledgesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryKnowledgesResponse) SetBody(v *QueryKnowledgesResponseBody) *QueryKnowledgesResponse {
	s.Body = v
	return s
}

type QueryPerspectivesRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s QueryPerspectivesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesRequest) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesRequest) SetAgentKey(v string) *QueryPerspectivesRequest {
	s.AgentKey = &v
	return s
}

type QueryPerspectivesResponseBody struct {
	Perspectives []*QueryPerspectivesResponseBodyPerspectives `json:"Perspectives,omitempty" xml:"Perspectives,omitempty" type:"Repeated"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPerspectivesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBody) SetPerspectives(v []*QueryPerspectivesResponseBodyPerspectives) *QueryPerspectivesResponseBody {
	s.Perspectives = v
	return s
}

func (s *QueryPerspectivesResponseBody) SetRequestId(v string) *QueryPerspectivesResponseBody {
	s.RequestId = &v
	return s
}

type QueryPerspectivesResponseBodyPerspectives struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserName  *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	ModifyTime      *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserName  *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	PerspectiveId   *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	SelfDefine      *bool   `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryPerspectivesResponseBodyPerspectives) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponseBodyPerspectives) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetCreateTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.CreateTime = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetCreateUserName(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.CreateUserName = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetModifyTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.ModifyTime = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetModifyUserName(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.ModifyUserName = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetName(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.Name = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetPerspectiveCode(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.PerspectiveCode = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetPerspectiveId(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.PerspectiveId = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetSelfDefine(v bool) *QueryPerspectivesResponseBodyPerspectives {
	s.SelfDefine = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetStatus(v int32) *QueryPerspectivesResponseBodyPerspectives {
	s.Status = &v
	return s
}

type QueryPerspectivesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPerspectivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPerspectivesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponse) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponse) SetHeaders(v map[string]*string) *QueryPerspectivesResponse {
	s.Headers = v
	return s
}

func (s *QueryPerspectivesResponse) SetStatusCode(v int32) *QueryPerspectivesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPerspectivesResponse) SetBody(v *QueryPerspectivesResponseBody) *QueryPerspectivesResponse {
	s.Body = v
	return s
}

type QuerySystemEntitiesRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s QuerySystemEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesRequest) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesRequest) SetAgentKey(v string) *QuerySystemEntitiesRequest {
	s.AgentKey = &v
	return s
}

func (s *QuerySystemEntitiesRequest) SetEntityName(v string) *QuerySystemEntitiesRequest {
	s.EntityName = &v
	return s
}

type QuerySystemEntitiesResponseBody struct {
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemEntities []*QuerySystemEntitiesResponseBodySystemEntities `json:"SystemEntities,omitempty" xml:"SystemEntities,omitempty" type:"Repeated"`
}

func (s QuerySystemEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesResponseBody) SetRequestId(v string) *QuerySystemEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySystemEntitiesResponseBody) SetSystemEntities(v []*QuerySystemEntitiesResponseBodySystemEntities) *QuerySystemEntitiesResponseBody {
	s.SystemEntities = v
	return s
}

type QuerySystemEntitiesResponseBodySystemEntities struct {
	DefaultQuestion *string `json:"DefaultQuestion,omitempty" xml:"DefaultQuestion,omitempty"`
	EntityCode      *string `json:"EntityCode,omitempty" xml:"EntityCode,omitempty"`
	EntityName      *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s QuerySystemEntitiesResponseBodySystemEntities) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesResponseBodySystemEntities) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesResponseBodySystemEntities) SetDefaultQuestion(v string) *QuerySystemEntitiesResponseBodySystemEntities {
	s.DefaultQuestion = &v
	return s
}

func (s *QuerySystemEntitiesResponseBodySystemEntities) SetEntityCode(v string) *QuerySystemEntitiesResponseBodySystemEntities {
	s.EntityCode = &v
	return s
}

func (s *QuerySystemEntitiesResponseBodySystemEntities) SetEntityName(v string) *QuerySystemEntitiesResponseBodySystemEntities {
	s.EntityName = &v
	return s
}

type QuerySystemEntitiesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySystemEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySystemEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesResponse) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesResponse) SetHeaders(v map[string]*string) *QuerySystemEntitiesResponse {
	s.Headers = v
	return s
}

func (s *QuerySystemEntitiesResponse) SetStatusCode(v int32) *QuerySystemEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySystemEntitiesResponse) SetBody(v *QuerySystemEntitiesResponseBody) *QuerySystemEntitiesResponse {
	s.Body = v
	return s
}

type RemoveEntityMemberRequest struct {
	AgentKey   *string                          `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EntityId   *int64                           `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Member     *RemoveEntityMemberRequestMember `json:"Member,omitempty" xml:"Member,omitempty" type:"Struct"`
	RemoveType *string                          `json:"RemoveType,omitempty" xml:"RemoveType,omitempty"`
}

func (s RemoveEntityMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberRequest) SetAgentKey(v string) *RemoveEntityMemberRequest {
	s.AgentKey = &v
	return s
}

func (s *RemoveEntityMemberRequest) SetEntityId(v int64) *RemoveEntityMemberRequest {
	s.EntityId = &v
	return s
}

func (s *RemoveEntityMemberRequest) SetMember(v *RemoveEntityMemberRequestMember) *RemoveEntityMemberRequest {
	s.Member = v
	return s
}

func (s *RemoveEntityMemberRequest) SetRemoveType(v string) *RemoveEntityMemberRequest {
	s.RemoveType = &v
	return s
}

type RemoveEntityMemberRequestMember struct {
	MemberName *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s RemoveEntityMemberRequestMember) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberRequestMember) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberRequestMember) SetMemberName(v string) *RemoveEntityMemberRequestMember {
	s.MemberName = &v
	return s
}

func (s *RemoveEntityMemberRequestMember) SetSynonyms(v []*string) *RemoveEntityMemberRequestMember {
	s.Synonyms = v
	return s
}

type RemoveEntityMemberShrinkRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EntityId     *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	MemberShrink *string `json:"Member,omitempty" xml:"Member,omitempty"`
	RemoveType   *string `json:"RemoveType,omitempty" xml:"RemoveType,omitempty"`
}

func (s RemoveEntityMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberShrinkRequest) SetAgentKey(v string) *RemoveEntityMemberShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *RemoveEntityMemberShrinkRequest) SetEntityId(v int64) *RemoveEntityMemberShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *RemoveEntityMemberShrinkRequest) SetMemberShrink(v string) *RemoveEntityMemberShrinkRequest {
	s.MemberShrink = &v
	return s
}

func (s *RemoveEntityMemberShrinkRequest) SetRemoveType(v string) *RemoveEntityMemberShrinkRequest {
	s.RemoveType = &v
	return s
}

type RemoveEntityMemberResponseBody struct {
	EntityId  *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntityMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberResponseBody) SetEntityId(v int64) *RemoveEntityMemberResponseBody {
	s.EntityId = &v
	return s
}

func (s *RemoveEntityMemberResponseBody) SetRequestId(v string) *RemoveEntityMemberResponseBody {
	s.RequestId = &v
	return s
}

type RemoveEntityMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveEntityMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveEntityMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberResponse) SetHeaders(v map[string]*string) *RemoveEntityMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntityMemberResponse) SetStatusCode(v int32) *RemoveEntityMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEntityMemberResponse) SetBody(v *RemoveEntityMemberResponseBody) *RemoveEntityMemberResponse {
	s.Body = v
	return s
}

type RemoveSynonymRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	Synonym      *string `json:"Synonym,omitempty" xml:"Synonym,omitempty"`
}

func (s RemoveSynonymRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSynonymRequest) GoString() string {
	return s.String()
}

func (s *RemoveSynonymRequest) SetAgentKey(v string) *RemoveSynonymRequest {
	s.AgentKey = &v
	return s
}

func (s *RemoveSynonymRequest) SetCoreWordName(v string) *RemoveSynonymRequest {
	s.CoreWordName = &v
	return s
}

func (s *RemoveSynonymRequest) SetSynonym(v string) *RemoveSynonymRequest {
	s.Synonym = &v
	return s
}

type RemoveSynonymResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSynonymResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSynonymResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSynonymResponseBody) SetRequestId(v string) *RemoveSynonymResponseBody {
	s.RequestId = &v
	return s
}

type RemoveSynonymResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSynonymResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSynonymResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSynonymResponse) GoString() string {
	return s.String()
}

func (s *RemoveSynonymResponse) SetHeaders(v map[string]*string) *RemoveSynonymResponse {
	s.Headers = v
	return s
}

func (s *RemoveSynonymResponse) SetStatusCode(v int32) *RemoveSynonymResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSynonymResponse) SetBody(v *RemoveSynonymResponseBody) *RemoveSynonymResponse {
	s.Body = v
	return s
}

type TestDialogFlowRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s TestDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s TestDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *TestDialogFlowRequest) SetAgentKey(v string) *TestDialogFlowRequest {
	s.AgentKey = &v
	return s
}

func (s *TestDialogFlowRequest) SetDialogId(v int64) *TestDialogFlowRequest {
	s.DialogId = &v
	return s
}

type TestDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TestDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *TestDialogFlowResponseBody) SetRequestId(v string) *TestDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type TestDialogFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s TestDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *TestDialogFlowResponse) SetHeaders(v map[string]*string) *TestDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *TestDialogFlowResponse) SetStatusCode(v int32) *TestDialogFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *TestDialogFlowResponse) SetBody(v *TestDialogFlowResponseBody) *TestDialogFlowResponse {
	s.Body = v
	return s
}

type UpdateCategoryRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryRequest) SetAgentKey(v string) *UpdateCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateCategoryRequest) SetCategoryId(v int64) *UpdateCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateCategoryRequest) SetName(v string) *UpdateCategoryRequest {
	s.Name = &v
	return s
}

type UpdateCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponseBody) SetRequestId(v string) *UpdateCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCategoryResponseBody) SetSuccess(v bool) *UpdateCategoryResponseBody {
	s.Success = &v
	return s
}

type UpdateCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponse) SetHeaders(v map[string]*string) *UpdateCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateCategoryResponse) SetStatusCode(v int32) *UpdateCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCategoryResponse) SetBody(v *UpdateCategoryResponseBody) *UpdateCategoryResponse {
	s.Body = v
	return s
}

type UpdateCoreWordRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CoreWordCode *string `json:"CoreWordCode,omitempty" xml:"CoreWordCode,omitempty"`
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
}

func (s UpdateCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoreWordRequest) GoString() string {
	return s.String()
}

func (s *UpdateCoreWordRequest) SetAgentKey(v string) *UpdateCoreWordRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateCoreWordRequest) SetCoreWordCode(v string) *UpdateCoreWordRequest {
	s.CoreWordCode = &v
	return s
}

func (s *UpdateCoreWordRequest) SetCoreWordName(v string) *UpdateCoreWordRequest {
	s.CoreWordName = &v
	return s
}

type UpdateCoreWordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCoreWordResponseBody) SetRequestId(v string) *UpdateCoreWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCoreWordResponseBody) SetSuccess(v bool) *UpdateCoreWordResponseBody {
	s.Success = &v
	return s
}

type UpdateCoreWordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoreWordResponse) GoString() string {
	return s.String()
}

func (s *UpdateCoreWordResponse) SetHeaders(v map[string]*string) *UpdateCoreWordResponse {
	s.Headers = v
	return s
}

func (s *UpdateCoreWordResponse) SetStatusCode(v int32) *UpdateCoreWordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCoreWordResponse) SetBody(v *UpdateCoreWordResponseBody) *UpdateCoreWordResponse {
	s.Body = v
	return s
}

type UpdateDialogRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DialogId    *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	DialogName  *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
}

func (s UpdateDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogRequest) GoString() string {
	return s.String()
}

func (s *UpdateDialogRequest) SetAgentKey(v string) *UpdateDialogRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDialogRequest) SetDescription(v string) *UpdateDialogRequest {
	s.Description = &v
	return s
}

func (s *UpdateDialogRequest) SetDialogId(v int64) *UpdateDialogRequest {
	s.DialogId = &v
	return s
}

func (s *UpdateDialogRequest) SetDialogName(v string) *UpdateDialogRequest {
	s.DialogName = &v
	return s
}

type UpdateDialogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDialogResponseBody) SetRequestId(v string) *UpdateDialogResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDialogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogResponse) GoString() string {
	return s.String()
}

func (s *UpdateDialogResponse) SetHeaders(v map[string]*string) *UpdateDialogResponse {
	s.Headers = v
	return s
}

func (s *UpdateDialogResponse) SetStatusCode(v int32) *UpdateDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDialogResponse) SetBody(v *UpdateDialogResponseBody) *UpdateDialogResponse {
	s.Body = v
	return s
}

type UpdateDialogFlowRequest struct {
	AgentKey         *string                                  `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId         *int64                                   `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	ModuleDefinition *UpdateDialogFlowRequestModuleDefinition `json:"ModuleDefinition,omitempty" xml:"ModuleDefinition,omitempty" type:"Struct"`
}

func (s UpdateDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowRequest) SetAgentKey(v string) *UpdateDialogFlowRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDialogFlowRequest) SetDialogId(v int64) *UpdateDialogFlowRequest {
	s.DialogId = &v
	return s
}

func (s *UpdateDialogFlowRequest) SetModuleDefinition(v *UpdateDialogFlowRequestModuleDefinition) *UpdateDialogFlowRequest {
	s.ModuleDefinition = v
	return s
}

type UpdateDialogFlowRequestModuleDefinition struct {
	GlobalVars       map[string]interface{} `json:"GlobalVars,omitempty" xml:"GlobalVars,omitempty"`
	ModuleDefinition *PaasProcessData       `json:"ModuleDefinition,omitempty" xml:"ModuleDefinition,omitempty"`
}

func (s UpdateDialogFlowRequestModuleDefinition) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowRequestModuleDefinition) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowRequestModuleDefinition) SetGlobalVars(v map[string]interface{}) *UpdateDialogFlowRequestModuleDefinition {
	s.GlobalVars = v
	return s
}

func (s *UpdateDialogFlowRequestModuleDefinition) SetModuleDefinition(v *PaasProcessData) *UpdateDialogFlowRequestModuleDefinition {
	s.ModuleDefinition = v
	return s
}

type UpdateDialogFlowShrinkRequest struct {
	AgentKey               *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DialogId               *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	ModuleDefinitionShrink *string `json:"ModuleDefinition,omitempty" xml:"ModuleDefinition,omitempty"`
}

func (s UpdateDialogFlowShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowShrinkRequest) SetAgentKey(v string) *UpdateDialogFlowShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDialogFlowShrinkRequest) SetDialogId(v int64) *UpdateDialogFlowShrinkRequest {
	s.DialogId = &v
	return s
}

func (s *UpdateDialogFlowShrinkRequest) SetModuleDefinitionShrink(v string) *UpdateDialogFlowShrinkRequest {
	s.ModuleDefinitionShrink = &v
	return s
}

type UpdateDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowResponseBody) SetRequestId(v string) *UpdateDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDialogFlowResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowResponse) SetHeaders(v map[string]*string) *UpdateDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *UpdateDialogFlowResponse) SetStatusCode(v int32) *UpdateDialogFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDialogFlowResponse) SetBody(v *UpdateDialogFlowResponseBody) *UpdateDialogFlowResponse {
	s.Body = v
	return s
}

type UpdateEntityRequest struct {
	AgentKey   *string                       `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EntityId   *int64                        `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName *string                       `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType *string                       `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Members    []*UpdateEntityRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Regex      *string                       `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s UpdateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityRequest) SetAgentKey(v string) *UpdateEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateEntityRequest) SetEntityId(v int64) *UpdateEntityRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateEntityRequest) SetEntityName(v string) *UpdateEntityRequest {
	s.EntityName = &v
	return s
}

func (s *UpdateEntityRequest) SetEntityType(v string) *UpdateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityRequest) SetMembers(v []*UpdateEntityRequestMembers) *UpdateEntityRequest {
	s.Members = v
	return s
}

func (s *UpdateEntityRequest) SetRegex(v string) *UpdateEntityRequest {
	s.Regex = &v
	return s
}

type UpdateEntityRequestMembers struct {
	MemberName *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s UpdateEntityRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateEntityRequestMembers) SetMemberName(v string) *UpdateEntityRequestMembers {
	s.MemberName = &v
	return s
}

func (s *UpdateEntityRequestMembers) SetSynonyms(v []*string) *UpdateEntityRequestMembers {
	s.Synonyms = v
	return s
}

type UpdateEntityShrinkRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	EntityId      *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName    *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	Regex         *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s UpdateEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityShrinkRequest) SetAgentKey(v string) *UpdateEntityShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetEntityId(v int64) *UpdateEntityShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetEntityName(v string) *UpdateEntityShrinkRequest {
	s.EntityName = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetEntityType(v string) *UpdateEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetMembersShrink(v string) *UpdateEntityShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetRegex(v string) *UpdateEntityShrinkRequest {
	s.Regex = &v
	return s
}

type UpdateEntityResponseBody struct {
	EntityId  *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponseBody) SetEntityId(v int64) *UpdateEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *UpdateEntityResponseBody) SetRequestId(v string) *UpdateEntityResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponse) SetHeaders(v map[string]*string) *UpdateEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateEntityResponse) SetStatusCode(v int32) *UpdateEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEntityResponse) SetBody(v *UpdateEntityResponseBody) *UpdateEntityResponse {
	s.Body = v
	return s
}

type UpdateIntentRequest struct {
	AgentKey         *string          `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	IntentDefinition *IntentCreateDTO `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	IntentId         *int64           `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequest) SetAgentKey(v string) *UpdateIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateIntentRequest) SetIntentDefinition(v *IntentCreateDTO) *UpdateIntentRequest {
	s.IntentDefinition = v
	return s
}

func (s *UpdateIntentRequest) SetIntentId(v int64) *UpdateIntentRequest {
	s.IntentId = &v
	return s
}

type UpdateIntentShrinkRequest struct {
	AgentKey               *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	IntentId               *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntentShrinkRequest) SetAgentKey(v string) *UpdateIntentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetIntentDefinitionShrink(v string) *UpdateIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetIntentId(v int64) *UpdateIntentShrinkRequest {
	s.IntentId = &v
	return s
}

type UpdateIntentResponseBody struct {
	IntentId  *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntentResponseBody) SetIntentId(v int64) *UpdateIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *UpdateIntentResponseBody) SetRequestId(v string) *UpdateIntentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntentResponse) SetHeaders(v map[string]*string) *UpdateIntentResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntentResponse) SetStatusCode(v int32) *UpdateIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntentResponse) SetBody(v *UpdateIntentResponseBody) *UpdateIntentResponse {
	s.Body = v
	return s
}

type UpdateKnowledgeRequest struct {
	AgentKey  *string                          `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Knowledge *UpdateKnowledgeRequestKnowledge `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeRequest) SetAgentKey(v string) *UpdateKnowledgeRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateKnowledgeRequest) SetKnowledge(v *UpdateKnowledgeRequestKnowledge) *UpdateKnowledgeRequest {
	s.Knowledge = v
	return s
}

type UpdateKnowledgeRequestKnowledge struct {
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// FAQ ID
	KnowledgeId    *int64                                         `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	KnowledgeTitle *string                                        `json:"KnowledgeTitle,omitempty" xml:"KnowledgeTitle,omitempty"`
	KnowledgeType  *int32                                         `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	Outlines       []*UpdateKnowledgeRequestKnowledgeOutlines     `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	SimQuestions   []*UpdateKnowledgeRequestKnowledgeSimQuestions `json:"SimQuestions,omitempty" xml:"SimQuestions,omitempty" type:"Repeated"`
	Solutions      []*UpdateKnowledgeRequestKnowledgeSolutions    `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	StartDate      *string                                        `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s UpdateKnowledgeRequestKnowledge) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeRequestKnowledge) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeRequestKnowledge) SetCategoryId(v int64) *UpdateKnowledgeRequestKnowledge {
	s.CategoryId = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetEndDate(v string) *UpdateKnowledgeRequestKnowledge {
	s.EndDate = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetKnowledgeId(v int64) *UpdateKnowledgeRequestKnowledge {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetKnowledgeTitle(v string) *UpdateKnowledgeRequestKnowledge {
	s.KnowledgeTitle = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetKnowledgeType(v int32) *UpdateKnowledgeRequestKnowledge {
	s.KnowledgeType = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetOutlines(v []*UpdateKnowledgeRequestKnowledgeOutlines) *UpdateKnowledgeRequestKnowledge {
	s.Outlines = v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetSimQuestions(v []*UpdateKnowledgeRequestKnowledgeSimQuestions) *UpdateKnowledgeRequestKnowledge {
	s.SimQuestions = v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetSolutions(v []*UpdateKnowledgeRequestKnowledgeSolutions) *UpdateKnowledgeRequestKnowledge {
	s.Solutions = v
	return s
}

func (s *UpdateKnowledgeRequestKnowledge) SetStartDate(v string) *UpdateKnowledgeRequestKnowledge {
	s.StartDate = &v
	return s
}

type UpdateKnowledgeRequestKnowledgeOutlines struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	OutlineId   *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateKnowledgeRequestKnowledgeOutlines) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeRequestKnowledgeOutlines) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeRequestKnowledgeOutlines) SetAction(v string) *UpdateKnowledgeRequestKnowledgeOutlines {
	s.Action = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeOutlines) SetKnowledgeId(v int64) *UpdateKnowledgeRequestKnowledgeOutlines {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeOutlines) SetOutlineId(v int64) *UpdateKnowledgeRequestKnowledgeOutlines {
	s.OutlineId = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeOutlines) SetTitle(v string) *UpdateKnowledgeRequestKnowledgeOutlines {
	s.Title = &v
	return s
}

type UpdateKnowledgeRequestKnowledgeSimQuestions struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	SimQuestionId *int64  `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateKnowledgeRequestKnowledgeSimQuestions) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeRequestKnowledgeSimQuestions) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeRequestKnowledgeSimQuestions) SetAction(v string) *UpdateKnowledgeRequestKnowledgeSimQuestions {
	s.Action = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeSimQuestions) SetSimQuestionId(v int64) *UpdateKnowledgeRequestKnowledgeSimQuestions {
	s.SimQuestionId = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeSimQuestions) SetTitle(v string) *UpdateKnowledgeRequestKnowledgeSimQuestions {
	s.Title = &v
	return s
}

type UpdateKnowledgeRequestKnowledgeSolutions struct {
	Action         *string   `json:"Action,omitempty" xml:"Action,omitempty"`
	Content        *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	PerspectiveIds []*string `json:"PerspectiveIds,omitempty" xml:"PerspectiveIds,omitempty" type:"Repeated"`
	PlainText      *string   `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
	SolutionId     *int64    `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s UpdateKnowledgeRequestKnowledgeSolutions) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeRequestKnowledgeSolutions) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeRequestKnowledgeSolutions) SetAction(v string) *UpdateKnowledgeRequestKnowledgeSolutions {
	s.Action = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeSolutions) SetContent(v string) *UpdateKnowledgeRequestKnowledgeSolutions {
	s.Content = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeSolutions) SetPerspectiveIds(v []*string) *UpdateKnowledgeRequestKnowledgeSolutions {
	s.PerspectiveIds = v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeSolutions) SetPlainText(v string) *UpdateKnowledgeRequestKnowledgeSolutions {
	s.PlainText = &v
	return s
}

func (s *UpdateKnowledgeRequestKnowledgeSolutions) SetSolutionId(v int64) *UpdateKnowledgeRequestKnowledgeSolutions {
	s.SolutionId = &v
	return s
}

type UpdateKnowledgeShrinkRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	KnowledgeShrink *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
}

func (s UpdateKnowledgeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeShrinkRequest) SetAgentKey(v string) *UpdateKnowledgeShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateKnowledgeShrinkRequest) SetKnowledgeShrink(v string) *UpdateKnowledgeShrinkRequest {
	s.KnowledgeShrink = &v
	return s
}

type UpdateKnowledgeResponseBody struct {
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeResponseBody) SetKnowledgeId(v int64) *UpdateKnowledgeResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateKnowledgeResponseBody) SetRequestId(v string) *UpdateKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateKnowledgeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeResponse) SetStatusCode(v int32) *UpdateKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeResponse) SetBody(v *UpdateKnowledgeResponseBody) *UpdateKnowledgeResponse {
	s.Body = v
	return s
}

type UpdatePerspectiveRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s UpdatePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveRequest) SetAgentKey(v string) *UpdatePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdatePerspectiveRequest) SetName(v string) *UpdatePerspectiveRequest {
	s.Name = &v
	return s
}

func (s *UpdatePerspectiveRequest) SetPerspectiveId(v string) *UpdatePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type UpdatePerspectiveResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveResponseBody) SetRequestId(v string) *UpdatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveResponse) SetHeaders(v map[string]*string) *UpdatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *UpdatePerspectiveResponse) SetStatusCode(v int32) *UpdatePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePerspectiveResponse) SetBody(v *UpdatePerspectiveResponseBody) *UpdatePerspectiveResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("chatbot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivatePerspectiveWithOptions(request *ActivatePerspectiveRequest, runtime *util.RuntimeOptions) (_result *ActivatePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveId)) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivatePerspective"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivatePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivatePerspective(request *ActivatePerspectiveRequest) (_result *ActivatePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivatePerspectiveResponse{}
	_body, _err := client.ActivatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSynonymWithOptions(request *AddSynonymRequest, runtime *util.RuntimeOptions) (_result *AddSynonymResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	if !tea.BoolValue(util.IsUnset(request.Synonym)) {
		query["Synonym"] = request.Synonym
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSynonym"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSynonymResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSynonym(request *AddSynonymRequest) (_result *AddSynonymResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSynonymResponse{}
	_body, _err := client.AddSynonymWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppendEntityMemberWithOptions(tmpReq *AppendEntityMemberRequest, runtime *util.RuntimeOptions) (_result *AppendEntityMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AppendEntityMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Member)) {
		request.MemberShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Member, tea.String("Member"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyType)) {
		query["ApplyType"] = request.ApplyType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberShrink)) {
		query["Member"] = request.MemberShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AppendEntityMember"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppendEntityMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppendEntityMember(request *AppendEntityMemberRequest) (_result *AppendEntityMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppendEntityMemberResponse{}
	_body, _err := client.AppendEntityMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateWithOptions(request *AssociateRequest, runtime *util.RuntimeOptions) (_result *AssociateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Perspective)) {
		query["Perspective"] = request.Perspective
	}

	if !tea.BoolValue(util.IsUnset(request.RecommendNum)) {
		query["RecommendNum"] = request.RecommendNum
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Associate"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Associate(request *AssociateRequest) (_result *AssociateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateResponse{}
	_body, _err := client.AssociateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChatWithOptions(request *ChatRequest, runtime *util.RuntimeOptions) (_result *ChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentName)) {
		query["IntentName"] = request.IntentName
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.Perspective)) {
		query["Perspective"] = request.Perspective
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		query["SenderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderNick)) {
		query["SenderNick"] = request.SenderNick
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	if !tea.BoolValue(util.IsUnset(request.VendorParam)) {
		query["VendorParam"] = request.VendorParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Chat"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Chat(request *ChatRequest) (_result *ChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatResponse{}
	_body, _err := client.ChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBotWithOptions(request *CreateBotRequest, runtime *util.RuntimeOptions) (_result *CreateBotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Avatar)) {
		query["Avatar"] = request.Avatar
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		query["Introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		query["LanguageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RobotType)) {
		query["RobotType"] = request.RobotType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBot"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBot(request *CreateBotRequest) (_result *CreateBotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBotResponse{}
	_body, _err := client.CreateBotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCategoryWithOptions(request *CreateCategoryRequest, runtime *util.RuntimeOptions) (_result *CreateCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategoryId)) {
		query["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCategory"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCategory(request *CreateCategoryRequest) (_result *CreateCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCategoryResponse{}
	_body, _err := client.CreateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCoreWordWithOptions(request *CreateCoreWordRequest, runtime *util.RuntimeOptions) (_result *CreateCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCoreWord"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCoreWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCoreWord(request *CreateCoreWordRequest) (_result *CreateCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCoreWordResponse{}
	_body, _err := client.CreateCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDialogWithOptions(request *CreateDialogRequest, runtime *util.RuntimeOptions) (_result *CreateDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DialogName)) {
		query["DialogName"] = request.DialogName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDialog"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDialogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDialog(request *CreateDialogRequest) (_result *CreateDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDialogResponse{}
	_body, _err := client.CreateDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEntityWithOptions(tmpReq *CreateEntityRequest, runtime *util.RuntimeOptions) (_result *CreateEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		query["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		query["Members"] = request.MembersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Regex)) {
		query["Regex"] = request.Regex
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEntity"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEntity(request *CreateEntityRequest) (_result *CreateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEntityResponse{}
	_body, _err := client.CreateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIntentWithOptions(tmpReq *CreateIntentRequest, runtime *util.RuntimeOptions) (_result *CreateIntentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IntentDefinition)) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, tea.String("IntentDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentDefinitionShrink)) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIntent"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntent(request *CreateIntentRequest) (_result *CreateIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIntentResponse{}
	_body, _err := client.CreateIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKnowledgeWithOptions(tmpReq *CreateKnowledgeRequest, runtime *util.RuntimeOptions) (_result *CreateKnowledgeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateKnowledgeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Knowledge)) {
		request.KnowledgeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Knowledge, tea.String("Knowledge"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeShrink)) {
		body["Knowledge"] = request.KnowledgeShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKnowledge"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKnowledge(request *CreateKnowledgeRequest) (_result *CreateKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKnowledgeResponse{}
	_body, _err := client.CreateKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePerspectiveWithOptions(request *CreatePerspectiveRequest, runtime *util.RuntimeOptions) (_result *CreatePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePerspective"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePerspective(request *CreatePerspectiveRequest) (_result *CreatePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePerspectiveResponse{}
	_body, _err := client.CreatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBotWithOptions(request *DeleteBotRequest, runtime *util.RuntimeOptions) (_result *DeleteBotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBot"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBot(request *DeleteBotRequest) (_result *DeleteBotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBotResponse{}
	_body, _err := client.DeleteBotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCategoryWithOptions(request *DeleteCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCategory"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCategory(request *DeleteCategoryRequest) (_result *DeleteCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCoreWordWithOptions(request *DeleteCoreWordRequest, runtime *util.RuntimeOptions) (_result *DeleteCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCoreWord"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCoreWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCoreWord(request *DeleteCoreWordRequest) (_result *DeleteCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCoreWordResponse{}
	_body, _err := client.DeleteCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDialogWithOptions(request *DeleteDialogRequest, runtime *util.RuntimeOptions) (_result *DeleteDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDialog"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDialogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDialog(request *DeleteDialogRequest) (_result *DeleteDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDialogResponse{}
	_body, _err := client.DeleteDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEntityWithOptions(request *DeleteEntityRequest, runtime *util.RuntimeOptions) (_result *DeleteEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEntity"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEntity(request *DeleteEntityRequest) (_result *DeleteEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEntityResponse{}
	_body, _err := client.DeleteEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIntentWithOptions(request *DeleteIntentRequest, runtime *util.RuntimeOptions) (_result *DeleteIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIntent"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIntent(request *DeleteIntentRequest) (_result *DeleteIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIntentResponse{}
	_body, _err := client.DeleteIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKnowledgeWithOptions(request *DeleteKnowledgeRequest, runtime *util.RuntimeOptions) (_result *DeleteKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKnowledge"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKnowledge(request *DeleteKnowledgeRequest) (_result *DeleteKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.DeleteKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePerspectiveWithOptions(request *DeletePerspectiveRequest, runtime *util.RuntimeOptions) (_result *DeletePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveId)) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePerspective"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePerspective(request *DeletePerspectiveRequest) (_result *DeletePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePerspectiveResponse{}
	_body, _err := client.DeletePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBotWithOptions(request *DescribeBotRequest, runtime *util.RuntimeOptions) (_result *DescribeBotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBot"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBot(request *DescribeBotRequest) (_result *DescribeBotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBotResponse{}
	_body, _err := client.DescribeBotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCategoryWithOptions(request *DescribeCategoryRequest, runtime *util.RuntimeOptions) (_result *DescribeCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCategory"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCategory(request *DescribeCategoryRequest) (_result *DescribeCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCategoryResponse{}
	_body, _err := client.DescribeCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCoreWordWithOptions(request *DescribeCoreWordRequest, runtime *util.RuntimeOptions) (_result *DescribeCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCoreWord"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCoreWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCoreWord(request *DescribeCoreWordRequest) (_result *DescribeCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCoreWordResponse{}
	_body, _err := client.DescribeCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDialogWithOptions(request *DescribeDialogRequest, runtime *util.RuntimeOptions) (_result *DescribeDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDialog"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDialogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDialog(request *DescribeDialogRequest) (_result *DescribeDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDialogResponse{}
	_body, _err := client.DescribeDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDialogFlowWithOptions(request *DescribeDialogFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDialogFlow"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDialogFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDialogFlow(request *DescribeDialogFlowRequest) (_result *DescribeDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDialogFlowResponse{}
	_body, _err := client.DescribeDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEntitiesWithOptions(request *DescribeEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEntities"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEntities(request *DescribeEntitiesRequest) (_result *DescribeEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEntitiesResponse{}
	_body, _err := client.DescribeEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIntentWithOptions(request *DescribeIntentRequest, runtime *util.RuntimeOptions) (_result *DescribeIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIntent"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIntent(request *DescribeIntentRequest) (_result *DescribeIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIntentResponse{}
	_body, _err := client.DescribeIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKnowledgeWithOptions(request *DescribeKnowledgeRequest, runtime *util.RuntimeOptions) (_result *DescribeKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKnowledge"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKnowledge(request *DescribeKnowledgeRequest) (_result *DescribeKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKnowledgeResponse{}
	_body, _err := client.DescribeKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePerspectiveWithOptions(request *DescribePerspectiveRequest, runtime *util.RuntimeOptions) (_result *DescribePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveId)) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePerspective"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePerspective(request *DescribePerspectiveRequest) (_result *DescribePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePerspectiveResponse{}
	_body, _err := client.DescribePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDialogFlowWithOptions(request *DisableDialogFlowRequest, runtime *util.RuntimeOptions) (_result *DisableDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableDialogFlow"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableDialogFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDialogFlow(request *DisableDialogFlowRequest) (_result *DisableDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDialogFlowResponse{}
	_body, _err := client.DisableDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableKnowledgeWithOptions(request *DisableKnowledgeRequest, runtime *util.RuntimeOptions) (_result *DisableKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableKnowledge"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableKnowledge(request *DisableKnowledgeRequest) (_result *DisableKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableKnowledgeResponse{}
	_body, _err := client.DisableKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FeedbackWithOptions(request *FeedbackRequest, runtime *util.RuntimeOptions) (_result *FeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		query["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Feedback"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Feedback(request *FeedbackRequest) (_result *FeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FeedbackResponse{}
	_body, _err := client.FeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncResultWithOptions(request *GetAsyncResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncResult"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncResult(request *GetAsyncResultRequest) (_result *GetAsyncResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.GetAsyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotChatDataWithOptions(request *GetBotChatDataRequest, runtime *util.RuntimeOptions) (_result *GetBotChatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBotChatData"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBotChatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotChatData(request *GetBotChatDataRequest) (_result *GetBotChatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotChatDataResponse{}
	_body, _err := client.GetBotChatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotDsStatDataWithOptions(request *GetBotDsStatDataRequest, runtime *util.RuntimeOptions) (_result *GetBotDsStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBotDsStatData"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBotDsStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotDsStatData(request *GetBotDsStatDataRequest) (_result *GetBotDsStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotDsStatDataResponse{}
	_body, _err := client.GetBotDsStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotKnowledgeStatDataWithOptions(request *GetBotKnowledgeStatDataRequest, runtime *util.RuntimeOptions) (_result *GetBotKnowledgeStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBotKnowledgeStatData"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBotKnowledgeStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotKnowledgeStatData(request *GetBotKnowledgeStatDataRequest) (_result *GetBotKnowledgeStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotKnowledgeStatDataResponse{}
	_body, _err := client.GetBotKnowledgeStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotSessionDataWithOptions(request *GetBotSessionDataRequest, runtime *util.RuntimeOptions) (_result *GetBotSessionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBotSessionData"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBotSessionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotSessionData(request *GetBotSessionDataRequest) (_result *GetBotSessionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotSessionDataResponse{}
	_body, _err := client.GetBotSessionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConversationListWithOptions(request *GetConversationListRequest, runtime *util.RuntimeOptions) (_result *GetConversationListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		query["SenderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConversationList"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConversationListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConversationList(request *GetConversationListRequest) (_result *GetConversationListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConversationListResponse{}
	_body, _err := client.GetConversationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotChatHistorysWithOptions(request *ListBotChatHistorysRequest, runtime *util.RuntimeOptions) (_result *ListBotChatHistorysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotChatHistorys"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotChatHistorysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotChatHistorys(request *ListBotChatHistorysRequest) (_result *ListBotChatHistorysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotChatHistorysResponse{}
	_body, _err := client.ListBotChatHistorysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotColdDsDatasWithOptions(request *ListBotColdDsDatasRequest, runtime *util.RuntimeOptions) (_result *ListBotColdDsDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotColdDsDatas"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotColdDsDatasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotColdDsDatas(request *ListBotColdDsDatasRequest) (_result *ListBotColdDsDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotColdDsDatasResponse{}
	_body, _err := client.ListBotColdDsDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotColdKnowledgesWithOptions(request *ListBotColdKnowledgesRequest, runtime *util.RuntimeOptions) (_result *ListBotColdKnowledgesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotColdKnowledges"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotColdKnowledgesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotColdKnowledges(request *ListBotColdKnowledgesRequest) (_result *ListBotColdKnowledgesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotColdKnowledgesResponse{}
	_body, _err := client.ListBotColdKnowledgesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotDsDetailsWithOptions(request *ListBotDsDetailsRequest, runtime *util.RuntimeOptions) (_result *ListBotDsDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotDsDetails"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotDsDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotDsDetails(request *ListBotDsDetailsRequest) (_result *ListBotDsDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotDsDetailsResponse{}
	_body, _err := client.ListBotDsDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotHotDsDatasWithOptions(request *ListBotHotDsDatasRequest, runtime *util.RuntimeOptions) (_result *ListBotHotDsDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotHotDsDatas"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotHotDsDatasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotHotDsDatas(request *ListBotHotDsDatasRequest) (_result *ListBotHotDsDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotHotDsDatasResponse{}
	_body, _err := client.ListBotHotDsDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotHotKnowledgesWithOptions(request *ListBotHotKnowledgesRequest, runtime *util.RuntimeOptions) (_result *ListBotHotKnowledgesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotHotKnowledges"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotHotKnowledgesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotHotKnowledges(request *ListBotHotKnowledgesRequest) (_result *ListBotHotKnowledgesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotHotKnowledgesResponse{}
	_body, _err := client.ListBotHotKnowledgesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotKnowledgeDetailsWithOptions(request *ListBotKnowledgeDetailsRequest, runtime *util.RuntimeOptions) (_result *ListBotKnowledgeDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotKnowledgeDetails"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotKnowledgeDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotKnowledgeDetails(request *ListBotKnowledgeDetailsRequest) (_result *ListBotKnowledgeDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotKnowledgeDetailsResponse{}
	_body, _err := client.ListBotKnowledgeDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotReceptionDetailDatasWithOptions(request *ListBotReceptionDetailDatasRequest, runtime *util.RuntimeOptions) (_result *ListBotReceptionDetailDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBotReceptionDetailDatas"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBotReceptionDetailDatasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotReceptionDetailDatas(request *ListBotReceptionDetailDatasRequest) (_result *ListBotReceptionDetailDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotReceptionDetailDatasResponse{}
	_body, _err := client.ListBotReceptionDetailDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConversationLogsWithOptions(request *ListConversationLogsRequest, runtime *util.RuntimeOptions) (_result *ListConversationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConversationLogs"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConversationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConversationLogs(request *ListConversationLogsRequest) (_result *ListConversationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConversationLogsResponse{}
	_body, _err := client.ListConversationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveKnowledgeCategoryWithOptions(request *MoveKnowledgeCategoryRequest, runtime *util.RuntimeOptions) (_result *MoveKnowledgeCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveKnowledgeCategory"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveKnowledgeCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveKnowledgeCategory(request *MoveKnowledgeCategoryRequest) (_result *MoveKnowledgeCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveKnowledgeCategoryResponse{}
	_body, _err := client.MoveKnowledgeCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishDialogFlowWithOptions(request *PublishDialogFlowRequest, runtime *util.RuntimeOptions) (_result *PublishDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishDialogFlow"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishDialogFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishDialogFlow(request *PublishDialogFlowRequest) (_result *PublishDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishDialogFlowResponse{}
	_body, _err := client.PublishDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishKnowledgeWithOptions(request *PublishKnowledgeRequest, runtime *util.RuntimeOptions) (_result *PublishKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Async)) {
		query["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishKnowledge"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishKnowledge(request *PublishKnowledgeRequest) (_result *PublishKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishKnowledgeResponse{}
	_body, _err := client.PublishKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBotsWithOptions(request *QueryBotsRequest, runtime *util.RuntimeOptions) (_result *QueryBotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBots"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBots(request *QueryBotsRequest) (_result *QueryBotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBotsResponse{}
	_body, _err := client.QueryBotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCategoriesWithOptions(request *QueryCategoriesRequest, runtime *util.RuntimeOptions) (_result *QueryCategoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeType)) {
		query["KnowledgeType"] = request.KnowledgeType
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategoryId)) {
		query["ParentCategoryId"] = request.ParentCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowChildrens)) {
		query["ShowChildrens"] = request.ShowChildrens
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCategories"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCategories(request *QueryCategoriesRequest) (_result *QueryCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCategoriesResponse{}
	_body, _err := client.QueryCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCoreWordsWithOptions(request *QueryCoreWordsRequest, runtime *util.RuntimeOptions) (_result *QueryCoreWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Synonym)) {
		query["Synonym"] = request.Synonym
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCoreWords"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCoreWordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCoreWords(request *QueryCoreWordsRequest) (_result *QueryCoreWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCoreWordsResponse{}
	_body, _err := client.QueryCoreWordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ****
 *
 * @param request QueryDialogsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryDialogsResponse
 */
func (client *Client) QueryDialogsWithOptions(request *QueryDialogsRequest, runtime *util.RuntimeOptions) (_result *QueryDialogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogName)) {
		query["DialogName"] = request.DialogName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDialogs"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDialogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ****
 *
 * @param request QueryDialogsRequest
 * @return QueryDialogsResponse
 */
func (client *Client) QueryDialogs(request *QueryDialogsRequest) (_result *QueryDialogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDialogsResponse{}
	_body, _err := client.QueryDialogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEntitiesWithOptions(request *QueryEntitiesRequest, runtime *util.RuntimeOptions) (_result *QueryEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		query["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEntities"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEntities(request *QueryEntitiesRequest) (_result *QueryEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEntitiesResponse{}
	_body, _err := client.QueryEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryIntentsWithOptions(request *QueryIntentsRequest, runtime *util.RuntimeOptions) (_result *QueryIntentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentName)) {
		query["IntentName"] = request.IntentName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryIntents"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryIntentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryIntents(request *QueryIntentsRequest) (_result *QueryIntentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryIntentsResponse{}
	_body, _err := client.QueryIntentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryKnowledgesWithOptions(request *QueryKnowledgesRequest, runtime *util.RuntimeOptions) (_result *QueryKnowledgesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeTitle)) {
		query["KnowledgeTitle"] = request.KnowledgeTitle
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryKnowledges"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryKnowledgesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryKnowledges(request *QueryKnowledgesRequest) (_result *QueryKnowledgesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryKnowledgesResponse{}
	_body, _err := client.QueryKnowledgesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPerspectivesWithOptions(request *QueryPerspectivesRequest, runtime *util.RuntimeOptions) (_result *QueryPerspectivesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPerspectives"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPerspectives(request *QueryPerspectivesRequest) (_result *QueryPerspectivesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.QueryPerspectivesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySystemEntitiesWithOptions(request *QuerySystemEntitiesRequest, runtime *util.RuntimeOptions) (_result *QuerySystemEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		query["EntityName"] = request.EntityName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySystemEntities"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySystemEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySystemEntities(request *QuerySystemEntitiesRequest) (_result *QuerySystemEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySystemEntitiesResponse{}
	_body, _err := client.QuerySystemEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveEntityMemberWithOptions(tmpReq *RemoveEntityMemberRequest, runtime *util.RuntimeOptions) (_result *RemoveEntityMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveEntityMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Member)) {
		request.MemberShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Member, tea.String("Member"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberShrink)) {
		query["Member"] = request.MemberShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RemoveType)) {
		query["RemoveType"] = request.RemoveType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveEntityMember"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveEntityMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveEntityMember(request *RemoveEntityMemberRequest) (_result *RemoveEntityMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEntityMemberResponse{}
	_body, _err := client.RemoveEntityMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSynonymWithOptions(request *RemoveSynonymRequest, runtime *util.RuntimeOptions) (_result *RemoveSynonymResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	if !tea.BoolValue(util.IsUnset(request.Synonym)) {
		query["Synonym"] = request.Synonym
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSynonym"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSynonymResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSynonym(request *RemoveSynonymRequest) (_result *RemoveSynonymResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSynonymResponse{}
	_body, _err := client.RemoveSynonymWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestDialogFlowWithOptions(request *TestDialogFlowRequest, runtime *util.RuntimeOptions) (_result *TestDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TestDialogFlow"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TestDialogFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestDialogFlow(request *TestDialogFlowRequest) (_result *TestDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestDialogFlowResponse{}
	_body, _err := client.TestDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCategoryWithOptions(request *UpdateCategoryRequest, runtime *util.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCategory"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCategory(request *UpdateCategoryRequest) (_result *UpdateCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.UpdateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCoreWordWithOptions(request *UpdateCoreWordRequest, runtime *util.RuntimeOptions) (_result *UpdateCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordCode)) {
		query["CoreWordCode"] = request.CoreWordCode
	}

	if !tea.BoolValue(util.IsUnset(request.CoreWordName)) {
		query["CoreWordName"] = request.CoreWordName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCoreWord"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCoreWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCoreWord(request *UpdateCoreWordRequest) (_result *UpdateCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCoreWordResponse{}
	_body, _err := client.UpdateCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDialogWithOptions(request *UpdateDialogRequest, runtime *util.RuntimeOptions) (_result *UpdateDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	if !tea.BoolValue(util.IsUnset(request.DialogName)) {
		query["DialogName"] = request.DialogName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDialog"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDialogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDialog(request *UpdateDialogRequest) (_result *UpdateDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDialogResponse{}
	_body, _err := client.UpdateDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDialogFlowWithOptions(tmpReq *UpdateDialogFlowRequest, runtime *util.RuntimeOptions) (_result *UpdateDialogFlowResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDialogFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ModuleDefinition)) {
		request.ModuleDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModuleDefinition, tea.String("ModuleDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleDefinitionShrink)) {
		body["ModuleDefinition"] = request.ModuleDefinitionShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDialogFlow"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDialogFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDialogFlow(request *UpdateDialogFlowRequest) (_result *UpdateDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDialogFlowResponse{}
	_body, _err := client.UpdateDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEntityWithOptions(tmpReq *UpdateEntityRequest, runtime *util.RuntimeOptions) (_result *UpdateEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		query["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Regex)) {
		query["Regex"] = request.Regex
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		body["Members"] = request.MembersShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEntity"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEntity(request *UpdateEntityRequest) (_result *UpdateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEntityResponse{}
	_body, _err := client.UpdateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIntentWithOptions(tmpReq *UpdateIntentRequest, runtime *util.RuntimeOptions) (_result *UpdateIntentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IntentDefinition)) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, tea.String("IntentDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.IntentDefinitionShrink)) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIntent"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIntent(request *UpdateIntentRequest) (_result *UpdateIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIntentResponse{}
	_body, _err := client.UpdateIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKnowledgeWithOptions(tmpReq *UpdateKnowledgeRequest, runtime *util.RuntimeOptions) (_result *UpdateKnowledgeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateKnowledgeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Knowledge)) {
		request.KnowledgeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Knowledge, tea.String("Knowledge"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeShrink)) {
		body["Knowledge"] = request.KnowledgeShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKnowledge"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKnowledge(request *UpdateKnowledgeRequest) (_result *UpdateKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKnowledgeResponse{}
	_body, _err := client.UpdateKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePerspectiveWithOptions(request *UpdatePerspectiveRequest, runtime *util.RuntimeOptions) (_result *UpdatePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveId)) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePerspective"),
		Version:     tea.String("2017-10-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePerspective(request *UpdatePerspectiveRequest) (_result *UpdatePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePerspectiveResponse{}
	_body, _err := client.UpdatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
