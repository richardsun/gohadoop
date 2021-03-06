// Code generated by protoc-gen-go.
// source: yarn_service_protos.proto
// DO NOT EDIT!

package hadoop_yarn

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"
import hadoop_common "github.com/hortonworks/gohadoop/hadoop_common"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// ////////////////////////////////////////////////////
// ///// AM_RM_Protocol ///////////////////////////////
// ////////////////////////////////////////////////////
type RegisterApplicationMasterRequestProto struct {
	Host                 *string                    `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	RpcPort              *int32                     `protobuf:"varint,2,opt,name=rpc_port" json:"rpc_port,omitempty"`
	TrackingUrl          *string                    `protobuf:"bytes,3,opt,name=tracking_url" json:"tracking_url,omitempty"`
	ApplicationAttemptId *ApplicationAttemptIdProto `protobuf:"bytes,4,opt,name=application_attempt_id" json:"application_attempt_id,omitempty"`
	XXX_unrecognized     []byte                     `json:"-"`
}

func (m *RegisterApplicationMasterRequestProto) Reset()         { *m = RegisterApplicationMasterRequestProto{} }
func (m *RegisterApplicationMasterRequestProto) String() string { return proto.CompactTextString(m) }
func (*RegisterApplicationMasterRequestProto) ProtoMessage()    {}

func (m *RegisterApplicationMasterRequestProto) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *RegisterApplicationMasterRequestProto) GetRpcPort() int32 {
	if m != nil && m.RpcPort != nil {
		return *m.RpcPort
	}
	return 0
}

func (m *RegisterApplicationMasterRequestProto) GetTrackingUrl() string {
	if m != nil && m.TrackingUrl != nil {
		return *m.TrackingUrl
	}
	return ""
}

func (m *RegisterApplicationMasterRequestProto) GetApplicationAttemptId() *ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

type RegisterApplicationMasterResponseProto struct {
	MaximumCapability        *ResourceProto            `protobuf:"bytes,1,opt,name=maximumCapability" json:"maximumCapability,omitempty"`
	ClientToAmTokenMasterKey []byte                    `protobuf:"bytes,2,opt,name=client_to_am_token_master_key" json:"client_to_am_token_master_key,omitempty"`
	Application_ACLs         []*ApplicationACLMapProto `protobuf:"bytes,3,rep,name=application_ACLs" json:"application_ACLs,omitempty"`
	XXX_unrecognized         []byte                    `json:"-"`
}

func (m *RegisterApplicationMasterResponseProto) Reset() {
	*m = RegisterApplicationMasterResponseProto{}
}
func (m *RegisterApplicationMasterResponseProto) String() string { return proto.CompactTextString(m) }
func (*RegisterApplicationMasterResponseProto) ProtoMessage()    {}

func (m *RegisterApplicationMasterResponseProto) GetMaximumCapability() *ResourceProto {
	if m != nil {
		return m.MaximumCapability
	}
	return nil
}

func (m *RegisterApplicationMasterResponseProto) GetClientToAmTokenMasterKey() []byte {
	if m != nil {
		return m.ClientToAmTokenMasterKey
	}
	return nil
}

func (m *RegisterApplicationMasterResponseProto) GetApplication_ACLs() []*ApplicationACLMapProto {
	if m != nil {
		return m.Application_ACLs
	}
	return nil
}

type FinishApplicationMasterRequestProto struct {
	Diagnostics            *string                      `protobuf:"bytes,1,opt,name=diagnostics" json:"diagnostics,omitempty"`
	TrackingUrl            *string                      `protobuf:"bytes,2,opt,name=tracking_url" json:"tracking_url,omitempty"`
	FinalApplicationStatus *FinalApplicationStatusProto `protobuf:"varint,3,opt,name=final_application_status,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	ApplicationAttemptId   *ApplicationAttemptIdProto   `protobuf:"bytes,4,opt,name=application_attempt_id" json:"application_attempt_id,omitempty"`
	XXX_unrecognized       []byte                       `json:"-"`
}

func (m *FinishApplicationMasterRequestProto) Reset()         { *m = FinishApplicationMasterRequestProto{} }
func (m *FinishApplicationMasterRequestProto) String() string { return proto.CompactTextString(m) }
func (*FinishApplicationMasterRequestProto) ProtoMessage()    {}

func (m *FinishApplicationMasterRequestProto) GetDiagnostics() string {
	if m != nil && m.Diagnostics != nil {
		return *m.Diagnostics
	}
	return ""
}

func (m *FinishApplicationMasterRequestProto) GetTrackingUrl() string {
	if m != nil && m.TrackingUrl != nil {
		return *m.TrackingUrl
	}
	return ""
}

func (m *FinishApplicationMasterRequestProto) GetFinalApplicationStatus() FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return 0
}

func (m *FinishApplicationMasterRequestProto) GetApplicationAttemptId() *ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

type FinishApplicationMasterResponseProto struct {
	IsUnregistered   *bool  `protobuf:"varint,1,opt,name=isUnregistered,def=0" json:"isUnregistered,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FinishApplicationMasterResponseProto) Reset()         { *m = FinishApplicationMasterResponseProto{} }
func (m *FinishApplicationMasterResponseProto) String() string { return proto.CompactTextString(m) }
func (*FinishApplicationMasterResponseProto) ProtoMessage()    {}

const Default_FinishApplicationMasterResponseProto_IsUnregistered bool = false

func (m *FinishApplicationMasterResponseProto) GetIsUnregistered() bool {
	if m != nil && m.IsUnregistered != nil {
		return *m.IsUnregistered
	}
	return Default_FinishApplicationMasterResponseProto_IsUnregistered
}

type AllocateRequestProto struct {
	Ask                  []*ResourceRequestProto        `protobuf:"bytes,1,rep,name=ask" json:"ask,omitempty"`
	Release              []*ContainerIdProto            `protobuf:"bytes,2,rep,name=release" json:"release,omitempty"`
	BlacklistRequest     *ResourceBlacklistRequestProto `protobuf:"bytes,3,opt,name=blacklist_request" json:"blacklist_request,omitempty"`
	ResponseId           *int32                         `protobuf:"varint,4,opt,name=response_id" json:"response_id,omitempty"`
	Progress             *float32                       `protobuf:"fixed32,5,opt,name=progress" json:"progress,omitempty"`
	ApplicationAttemptId *ApplicationAttemptIdProto     `protobuf:"bytes,6,opt,name=application_attempt_id" json:"application_attempt_id,omitempty"`
	XXX_unrecognized     []byte                         `json:"-"`
}

func (m *AllocateRequestProto) Reset()         { *m = AllocateRequestProto{} }
func (m *AllocateRequestProto) String() string { return proto.CompactTextString(m) }
func (*AllocateRequestProto) ProtoMessage()    {}

func (m *AllocateRequestProto) GetAsk() []*ResourceRequestProto {
	if m != nil {
		return m.Ask
	}
	return nil
}

func (m *AllocateRequestProto) GetRelease() []*ContainerIdProto {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *AllocateRequestProto) GetBlacklistRequest() *ResourceBlacklistRequestProto {
	if m != nil {
		return m.BlacklistRequest
	}
	return nil
}

func (m *AllocateRequestProto) GetResponseId() int32 {
	if m != nil && m.ResponseId != nil {
		return *m.ResponseId
	}
	return 0
}

func (m *AllocateRequestProto) GetProgress() float32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

func (m *AllocateRequestProto) GetApplicationAttemptId() *ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

type NMTokenProto struct {
	NodeId           *NodeIdProto              `protobuf:"bytes,1,opt,name=nodeId" json:"nodeId,omitempty"`
	Token            *hadoop_common.TokenProto `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *NMTokenProto) Reset()         { *m = NMTokenProto{} }
func (m *NMTokenProto) String() string { return proto.CompactTextString(m) }
func (*NMTokenProto) ProtoMessage()    {}

func (m *NMTokenProto) GetNodeId() *NodeIdProto {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NMTokenProto) GetToken() *hadoop_common.TokenProto {
	if m != nil {
		return m.Token
	}
	return nil
}

type AllocateResponseProto struct {
	AMCommand                  *AMCommandProto         `protobuf:"varint,1,opt,name=a_m_command,enum=hadoop.yarn.AMCommandProto" json:"a_m_command,omitempty"`
	ResponseId                 *int32                  `protobuf:"varint,2,opt,name=response_id" json:"response_id,omitempty"`
	AllocatedContainers        []*ContainerProto       `protobuf:"bytes,3,rep,name=allocated_containers" json:"allocated_containers,omitempty"`
	CompletedContainerStatuses []*ContainerStatusProto `protobuf:"bytes,4,rep,name=completed_container_statuses" json:"completed_container_statuses,omitempty"`
	Limit                      *ResourceProto          `protobuf:"bytes,5,opt,name=limit" json:"limit,omitempty"`
	UpdatedNodes               []*NodeReportProto      `protobuf:"bytes,6,rep,name=updated_nodes" json:"updated_nodes,omitempty"`
	NumClusterNodes            *int32                  `protobuf:"varint,7,opt,name=num_cluster_nodes" json:"num_cluster_nodes,omitempty"`
	Preempt                    *PreemptionMessageProto `protobuf:"bytes,8,opt,name=preempt" json:"preempt,omitempty"`
	NmTokens                   []*NMTokenProto         `protobuf:"bytes,9,rep,name=nm_tokens" json:"nm_tokens,omitempty"`
	XXX_unrecognized           []byte                  `json:"-"`
}

func (m *AllocateResponseProto) Reset()         { *m = AllocateResponseProto{} }
func (m *AllocateResponseProto) String() string { return proto.CompactTextString(m) }
func (*AllocateResponseProto) ProtoMessage()    {}

func (m *AllocateResponseProto) GetAMCommand() AMCommandProto {
	if m != nil && m.AMCommand != nil {
		return *m.AMCommand
	}
	return 0
}

func (m *AllocateResponseProto) GetResponseId() int32 {
	if m != nil && m.ResponseId != nil {
		return *m.ResponseId
	}
	return 0
}

func (m *AllocateResponseProto) GetAllocatedContainers() []*ContainerProto {
	if m != nil {
		return m.AllocatedContainers
	}
	return nil
}

func (m *AllocateResponseProto) GetCompletedContainerStatuses() []*ContainerStatusProto {
	if m != nil {
		return m.CompletedContainerStatuses
	}
	return nil
}

func (m *AllocateResponseProto) GetLimit() *ResourceProto {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *AllocateResponseProto) GetUpdatedNodes() []*NodeReportProto {
	if m != nil {
		return m.UpdatedNodes
	}
	return nil
}

func (m *AllocateResponseProto) GetNumClusterNodes() int32 {
	if m != nil && m.NumClusterNodes != nil {
		return *m.NumClusterNodes
	}
	return 0
}

func (m *AllocateResponseProto) GetPreempt() *PreemptionMessageProto {
	if m != nil {
		return m.Preempt
	}
	return nil
}

func (m *AllocateResponseProto) GetNmTokens() []*NMTokenProto {
	if m != nil {
		return m.NmTokens
	}
	return nil
}

type GetNewApplicationRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetNewApplicationRequestProto) Reset()         { *m = GetNewApplicationRequestProto{} }
func (m *GetNewApplicationRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetNewApplicationRequestProto) ProtoMessage()    {}

type GetNewApplicationResponseProto struct {
	ApplicationId     *ApplicationIdProto `protobuf:"bytes,1,opt,name=application_id" json:"application_id,omitempty"`
	MaximumCapability *ResourceProto      `protobuf:"bytes,2,opt,name=maximumCapability" json:"maximumCapability,omitempty"`
	XXX_unrecognized  []byte              `json:"-"`
}

func (m *GetNewApplicationResponseProto) Reset()         { *m = GetNewApplicationResponseProto{} }
func (m *GetNewApplicationResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetNewApplicationResponseProto) ProtoMessage()    {}

func (m *GetNewApplicationResponseProto) GetApplicationId() *ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *GetNewApplicationResponseProto) GetMaximumCapability() *ResourceProto {
	if m != nil {
		return m.MaximumCapability
	}
	return nil
}

type GetApplicationReportRequestProto struct {
	ApplicationId    *ApplicationIdProto `protobuf:"bytes,1,opt,name=application_id" json:"application_id,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetApplicationReportRequestProto) Reset()         { *m = GetApplicationReportRequestProto{} }
func (m *GetApplicationReportRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetApplicationReportRequestProto) ProtoMessage()    {}

func (m *GetApplicationReportRequestProto) GetApplicationId() *ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

type GetApplicationReportResponseProto struct {
	ApplicationReport *ApplicationReportProto `protobuf:"bytes,1,opt,name=application_report" json:"application_report,omitempty"`
	XXX_unrecognized  []byte                  `json:"-"`
}

func (m *GetApplicationReportResponseProto) Reset()         { *m = GetApplicationReportResponseProto{} }
func (m *GetApplicationReportResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetApplicationReportResponseProto) ProtoMessage()    {}

func (m *GetApplicationReportResponseProto) GetApplicationReport() *ApplicationReportProto {
	if m != nil {
		return m.ApplicationReport
	}
	return nil
}

type SubmitApplicationRequestProto struct {
	ApplicationSubmissionContext *ApplicationSubmissionContextProto `protobuf:"bytes,1,opt,name=application_submission_context" json:"application_submission_context,omitempty"`
	XXX_unrecognized             []byte                             `json:"-"`
}

func (m *SubmitApplicationRequestProto) Reset()         { *m = SubmitApplicationRequestProto{} }
func (m *SubmitApplicationRequestProto) String() string { return proto.CompactTextString(m) }
func (*SubmitApplicationRequestProto) ProtoMessage()    {}

func (m *SubmitApplicationRequestProto) GetApplicationSubmissionContext() *ApplicationSubmissionContextProto {
	if m != nil {
		return m.ApplicationSubmissionContext
	}
	return nil
}

type SubmitApplicationResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SubmitApplicationResponseProto) Reset()         { *m = SubmitApplicationResponseProto{} }
func (m *SubmitApplicationResponseProto) String() string { return proto.CompactTextString(m) }
func (*SubmitApplicationResponseProto) ProtoMessage()    {}

type KillApplicationRequestProto struct {
	ApplicationId    *ApplicationIdProto `protobuf:"bytes,1,opt,name=application_id" json:"application_id,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *KillApplicationRequestProto) Reset()         { *m = KillApplicationRequestProto{} }
func (m *KillApplicationRequestProto) String() string { return proto.CompactTextString(m) }
func (*KillApplicationRequestProto) ProtoMessage()    {}

func (m *KillApplicationRequestProto) GetApplicationId() *ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

type KillApplicationResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *KillApplicationResponseProto) Reset()         { *m = KillApplicationResponseProto{} }
func (m *KillApplicationResponseProto) String() string { return proto.CompactTextString(m) }
func (*KillApplicationResponseProto) ProtoMessage()    {}

type GetClusterMetricsRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetClusterMetricsRequestProto) Reset()         { *m = GetClusterMetricsRequestProto{} }
func (m *GetClusterMetricsRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetClusterMetricsRequestProto) ProtoMessage()    {}

type GetClusterMetricsResponseProto struct {
	ClusterMetrics   *YarnClusterMetricsProto `protobuf:"bytes,1,opt,name=cluster_metrics" json:"cluster_metrics,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *GetClusterMetricsResponseProto) Reset()         { *m = GetClusterMetricsResponseProto{} }
func (m *GetClusterMetricsResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetClusterMetricsResponseProto) ProtoMessage()    {}

func (m *GetClusterMetricsResponseProto) GetClusterMetrics() *YarnClusterMetricsProto {
	if m != nil {
		return m.ClusterMetrics
	}
	return nil
}

type GetApplicationsRequestProto struct {
	ApplicationTypes  []string                    `protobuf:"bytes,1,rep,name=application_types" json:"application_types,omitempty"`
	ApplicationStates []YarnApplicationStateProto `protobuf:"varint,2,rep,name=application_states,enum=hadoop.yarn.YarnApplicationStateProto" json:"application_states,omitempty"`
	XXX_unrecognized  []byte                      `json:"-"`
}

func (m *GetApplicationsRequestProto) Reset()         { *m = GetApplicationsRequestProto{} }
func (m *GetApplicationsRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetApplicationsRequestProto) ProtoMessage()    {}

func (m *GetApplicationsRequestProto) GetApplicationTypes() []string {
	if m != nil {
		return m.ApplicationTypes
	}
	return nil
}

func (m *GetApplicationsRequestProto) GetApplicationStates() []YarnApplicationStateProto {
	if m != nil {
		return m.ApplicationStates
	}
	return nil
}

type GetApplicationsResponseProto struct {
	Applications     []*ApplicationReportProto `protobuf:"bytes,1,rep,name=applications" json:"applications,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *GetApplicationsResponseProto) Reset()         { *m = GetApplicationsResponseProto{} }
func (m *GetApplicationsResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetApplicationsResponseProto) ProtoMessage()    {}

func (m *GetApplicationsResponseProto) GetApplications() []*ApplicationReportProto {
	if m != nil {
		return m.Applications
	}
	return nil
}

type GetClusterNodesRequestProto struct {
	NodeStates       []NodeStateProto `protobuf:"varint,1,rep,name=nodeStates,enum=hadoop.yarn.NodeStateProto" json:"nodeStates,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GetClusterNodesRequestProto) Reset()         { *m = GetClusterNodesRequestProto{} }
func (m *GetClusterNodesRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetClusterNodesRequestProto) ProtoMessage()    {}

func (m *GetClusterNodesRequestProto) GetNodeStates() []NodeStateProto {
	if m != nil {
		return m.NodeStates
	}
	return nil
}

type GetClusterNodesResponseProto struct {
	NodeReports      []*NodeReportProto `protobuf:"bytes,1,rep,name=nodeReports" json:"nodeReports,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *GetClusterNodesResponseProto) Reset()         { *m = GetClusterNodesResponseProto{} }
func (m *GetClusterNodesResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetClusterNodesResponseProto) ProtoMessage()    {}

func (m *GetClusterNodesResponseProto) GetNodeReports() []*NodeReportProto {
	if m != nil {
		return m.NodeReports
	}
	return nil
}

type GetQueueInfoRequestProto struct {
	QueueName           *string `protobuf:"bytes,1,opt,name=queueName" json:"queueName,omitempty"`
	IncludeApplications *bool   `protobuf:"varint,2,opt,name=includeApplications" json:"includeApplications,omitempty"`
	IncludeChildQueues  *bool   `protobuf:"varint,3,opt,name=includeChildQueues" json:"includeChildQueues,omitempty"`
	Recursive           *bool   `protobuf:"varint,4,opt,name=recursive" json:"recursive,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *GetQueueInfoRequestProto) Reset()         { *m = GetQueueInfoRequestProto{} }
func (m *GetQueueInfoRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetQueueInfoRequestProto) ProtoMessage()    {}

func (m *GetQueueInfoRequestProto) GetQueueName() string {
	if m != nil && m.QueueName != nil {
		return *m.QueueName
	}
	return ""
}

func (m *GetQueueInfoRequestProto) GetIncludeApplications() bool {
	if m != nil && m.IncludeApplications != nil {
		return *m.IncludeApplications
	}
	return false
}

func (m *GetQueueInfoRequestProto) GetIncludeChildQueues() bool {
	if m != nil && m.IncludeChildQueues != nil {
		return *m.IncludeChildQueues
	}
	return false
}

func (m *GetQueueInfoRequestProto) GetRecursive() bool {
	if m != nil && m.Recursive != nil {
		return *m.Recursive
	}
	return false
}

type GetQueueInfoResponseProto struct {
	QueueInfo        *QueueInfoProto `protobuf:"bytes,1,opt,name=queueInfo" json:"queueInfo,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetQueueInfoResponseProto) Reset()         { *m = GetQueueInfoResponseProto{} }
func (m *GetQueueInfoResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetQueueInfoResponseProto) ProtoMessage()    {}

func (m *GetQueueInfoResponseProto) GetQueueInfo() *QueueInfoProto {
	if m != nil {
		return m.QueueInfo
	}
	return nil
}

type GetQueueUserAclsInfoRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetQueueUserAclsInfoRequestProto) Reset()         { *m = GetQueueUserAclsInfoRequestProto{} }
func (m *GetQueueUserAclsInfoRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetQueueUserAclsInfoRequestProto) ProtoMessage()    {}

type GetQueueUserAclsInfoResponseProto struct {
	QueueUserAcls    []*QueueUserACLInfoProto `protobuf:"bytes,1,rep,name=queueUserAcls" json:"queueUserAcls,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *GetQueueUserAclsInfoResponseProto) Reset()         { *m = GetQueueUserAclsInfoResponseProto{} }
func (m *GetQueueUserAclsInfoResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetQueueUserAclsInfoResponseProto) ProtoMessage()    {}

func (m *GetQueueUserAclsInfoResponseProto) GetQueueUserAcls() []*QueueUserACLInfoProto {
	if m != nil {
		return m.QueueUserAcls
	}
	return nil
}

type StartContainerRequestProto struct {
	ContainerLaunchContext *ContainerLaunchContextProto `protobuf:"bytes,1,opt,name=container_launch_context" json:"container_launch_context,omitempty"`
	ContainerToken         *hadoop_common.TokenProto    `protobuf:"bytes,2,opt,name=container_token" json:"container_token,omitempty"`
	Container              *ContainerProto              `protobuf:"bytes,3,opt,name=container" json:"container,omitempty"`
	XXX_unrecognized       []byte                       `json:"-"`
}

func (m *StartContainerRequestProto) Reset()         { *m = StartContainerRequestProto{} }
func (m *StartContainerRequestProto) String() string { return proto.CompactTextString(m) }
func (*StartContainerRequestProto) ProtoMessage()    {}

func (m *StartContainerRequestProto) GetContainerLaunchContext() *ContainerLaunchContextProto {
	if m != nil {
		return m.ContainerLaunchContext
	}
	return nil
}

func (m *StartContainerRequestProto) GetContainerToken() *hadoop_common.TokenProto {
	if m != nil {
		return m.ContainerToken
	}
	return nil
}

func (m *StartContainerRequestProto) GetContainer() *ContainerProto {
	if m != nil {
		return m.Container
	}
	return nil
}

type StartContainerResponseProto struct {
	ServicesMetaData []*StringBytesMapProto `protobuf:"bytes,1,rep,name=services_meta_data" json:"services_meta_data,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *StartContainerResponseProto) Reset()         { *m = StartContainerResponseProto{} }
func (m *StartContainerResponseProto) String() string { return proto.CompactTextString(m) }
func (*StartContainerResponseProto) ProtoMessage()    {}

func (m *StartContainerResponseProto) GetServicesMetaData() []*StringBytesMapProto {
	if m != nil {
		return m.ServicesMetaData
	}
	return nil
}

type StopContainerRequestProto struct {
	ContainerId      *ContainerIdProto `protobuf:"bytes,1,opt,name=container_id" json:"container_id,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *StopContainerRequestProto) Reset()         { *m = StopContainerRequestProto{} }
func (m *StopContainerRequestProto) String() string { return proto.CompactTextString(m) }
func (*StopContainerRequestProto) ProtoMessage()    {}

func (m *StopContainerRequestProto) GetContainerId() *ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

type StopContainerResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StopContainerResponseProto) Reset()         { *m = StopContainerResponseProto{} }
func (m *StopContainerResponseProto) String() string { return proto.CompactTextString(m) }
func (*StopContainerResponseProto) ProtoMessage()    {}

type GetContainerStatusRequestProto struct {
	ContainerId      *ContainerIdProto `protobuf:"bytes,1,opt,name=container_id" json:"container_id,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *GetContainerStatusRequestProto) Reset()         { *m = GetContainerStatusRequestProto{} }
func (m *GetContainerStatusRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetContainerStatusRequestProto) ProtoMessage()    {}

func (m *GetContainerStatusRequestProto) GetContainerId() *ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

type GetContainerStatusResponseProto struct {
	Status           *ContainerStatusProto `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *GetContainerStatusResponseProto) Reset()         { *m = GetContainerStatusResponseProto{} }
func (m *GetContainerStatusResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetContainerStatusResponseProto) ProtoMessage()    {}

func (m *GetContainerStatusResponseProto) GetStatus() *ContainerStatusProto {
	if m != nil {
		return m.Status
	}
	return nil
}

// // bulk API records
type StartContainersRequestProto struct {
	StartContainerRequest []*StartContainerRequestProto `protobuf:"bytes,1,rep,name=start_container_request" json:"start_container_request,omitempty"`
	XXX_unrecognized      []byte                        `json:"-"`
}

func (m *StartContainersRequestProto) Reset()         { *m = StartContainersRequestProto{} }
func (m *StartContainersRequestProto) String() string { return proto.CompactTextString(m) }
func (*StartContainersRequestProto) ProtoMessage()    {}

func (m *StartContainersRequestProto) GetStartContainerRequest() []*StartContainerRequestProto {
	if m != nil {
		return m.StartContainerRequest
	}
	return nil
}

type ContainerExceptionMapProto struct {
	ContainerId      *ContainerIdProto         `protobuf:"bytes,1,opt,name=container_id" json:"container_id,omitempty"`
	Exception        *SerializedExceptionProto `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *ContainerExceptionMapProto) Reset()         { *m = ContainerExceptionMapProto{} }
func (m *ContainerExceptionMapProto) String() string { return proto.CompactTextString(m) }
func (*ContainerExceptionMapProto) ProtoMessage()    {}

func (m *ContainerExceptionMapProto) GetContainerId() *ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *ContainerExceptionMapProto) GetException() *SerializedExceptionProto {
	if m != nil {
		return m.Exception
	}
	return nil
}

type StartContainersResponseProto struct {
	ServicesMetaData  []*StringBytesMapProto        `protobuf:"bytes,1,rep,name=services_meta_data" json:"services_meta_data,omitempty"`
	SucceededRequests []*ContainerIdProto           `protobuf:"bytes,2,rep,name=succeeded_requests" json:"succeeded_requests,omitempty"`
	FailedRequests    []*ContainerExceptionMapProto `protobuf:"bytes,3,rep,name=failed_requests" json:"failed_requests,omitempty"`
	XXX_unrecognized  []byte                        `json:"-"`
}

func (m *StartContainersResponseProto) Reset()         { *m = StartContainersResponseProto{} }
func (m *StartContainersResponseProto) String() string { return proto.CompactTextString(m) }
func (*StartContainersResponseProto) ProtoMessage()    {}

func (m *StartContainersResponseProto) GetServicesMetaData() []*StringBytesMapProto {
	if m != nil {
		return m.ServicesMetaData
	}
	return nil
}

func (m *StartContainersResponseProto) GetSucceededRequests() []*ContainerIdProto {
	if m != nil {
		return m.SucceededRequests
	}
	return nil
}

func (m *StartContainersResponseProto) GetFailedRequests() []*ContainerExceptionMapProto {
	if m != nil {
		return m.FailedRequests
	}
	return nil
}

type StopContainersRequestProto struct {
	ContainerId      []*ContainerIdProto `protobuf:"bytes,1,rep,name=container_id" json:"container_id,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *StopContainersRequestProto) Reset()         { *m = StopContainersRequestProto{} }
func (m *StopContainersRequestProto) String() string { return proto.CompactTextString(m) }
func (*StopContainersRequestProto) ProtoMessage()    {}

func (m *StopContainersRequestProto) GetContainerId() []*ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

type StopContainersResponseProto struct {
	SucceededRequests []*ContainerIdProto           `protobuf:"bytes,1,rep,name=succeeded_requests" json:"succeeded_requests,omitempty"`
	FailedRequests    []*ContainerExceptionMapProto `protobuf:"bytes,2,rep,name=failed_requests" json:"failed_requests,omitempty"`
	XXX_unrecognized  []byte                        `json:"-"`
}

func (m *StopContainersResponseProto) Reset()         { *m = StopContainersResponseProto{} }
func (m *StopContainersResponseProto) String() string { return proto.CompactTextString(m) }
func (*StopContainersResponseProto) ProtoMessage()    {}

func (m *StopContainersResponseProto) GetSucceededRequests() []*ContainerIdProto {
	if m != nil {
		return m.SucceededRequests
	}
	return nil
}

func (m *StopContainersResponseProto) GetFailedRequests() []*ContainerExceptionMapProto {
	if m != nil {
		return m.FailedRequests
	}
	return nil
}

type GetContainerStatusesRequestProto struct {
	ContainerId      []*ContainerIdProto `protobuf:"bytes,1,rep,name=container_id" json:"container_id,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetContainerStatusesRequestProto) Reset()         { *m = GetContainerStatusesRequestProto{} }
func (m *GetContainerStatusesRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetContainerStatusesRequestProto) ProtoMessage()    {}

func (m *GetContainerStatusesRequestProto) GetContainerId() []*ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

type GetContainerStatusesResponseProto struct {
	Status           []*ContainerStatusProto       `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
	FailedRequests   []*ContainerExceptionMapProto `protobuf:"bytes,2,rep,name=failed_requests" json:"failed_requests,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *GetContainerStatusesResponseProto) Reset()         { *m = GetContainerStatusesResponseProto{} }
func (m *GetContainerStatusesResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetContainerStatusesResponseProto) ProtoMessage()    {}

func (m *GetContainerStatusesResponseProto) GetStatus() []*ContainerStatusProto {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetContainerStatusesResponseProto) GetFailedRequests() []*ContainerExceptionMapProto {
	if m != nil {
		return m.FailedRequests
	}
	return nil
}

func init() {
}
