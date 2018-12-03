// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package dataproc provides access to the Cloud Dataproc API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/dataproc/apiv1 instead.
//
// See https://cloud.google.com/dataproc/
//
// Usage example:
//
//   import "google.golang.org/api/dataproc/v1"
//   ...
//   dataprocService, err := dataproc.New(oauthHttpClient)
package dataproc // import "google.golang.org/api/dataproc/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "dataproc:v1"
const apiName = "dataproc"
const apiVersion = "v1"
const basePath = "https://dataproc.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Locations = NewProjectsLocationsService(s)
	rs.Regions = NewProjectsRegionsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService

	Regions *ProjectsRegionsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	rs.WorkflowTemplates = NewProjectsLocationsWorkflowTemplatesService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	WorkflowTemplates *ProjectsLocationsWorkflowTemplatesService
}

func NewProjectsLocationsWorkflowTemplatesService(s *Service) *ProjectsLocationsWorkflowTemplatesService {
	rs := &ProjectsLocationsWorkflowTemplatesService{s: s}
	return rs
}

type ProjectsLocationsWorkflowTemplatesService struct {
	s *Service
}

func NewProjectsRegionsService(s *Service) *ProjectsRegionsService {
	rs := &ProjectsRegionsService{s: s}
	rs.Clusters = NewProjectsRegionsClustersService(s)
	rs.Jobs = NewProjectsRegionsJobsService(s)
	rs.Operations = NewProjectsRegionsOperationsService(s)
	rs.WorkflowTemplates = NewProjectsRegionsWorkflowTemplatesService(s)
	return rs
}

type ProjectsRegionsService struct {
	s *Service

	Clusters *ProjectsRegionsClustersService

	Jobs *ProjectsRegionsJobsService

	Operations *ProjectsRegionsOperationsService

	WorkflowTemplates *ProjectsRegionsWorkflowTemplatesService
}

func NewProjectsRegionsClustersService(s *Service) *ProjectsRegionsClustersService {
	rs := &ProjectsRegionsClustersService{s: s}
	return rs
}

type ProjectsRegionsClustersService struct {
	s *Service
}

func NewProjectsRegionsJobsService(s *Service) *ProjectsRegionsJobsService {
	rs := &ProjectsRegionsJobsService{s: s}
	return rs
}

type ProjectsRegionsJobsService struct {
	s *Service
}

func NewProjectsRegionsOperationsService(s *Service) *ProjectsRegionsOperationsService {
	rs := &ProjectsRegionsOperationsService{s: s}
	return rs
}

type ProjectsRegionsOperationsService struct {
	s *Service
}

func NewProjectsRegionsWorkflowTemplatesService(s *Service) *ProjectsRegionsWorkflowTemplatesService {
	rs := &ProjectsRegionsWorkflowTemplatesService{s: s}
	return rs
}

type ProjectsRegionsWorkflowTemplatesService struct {
	s *Service
}

// AcceleratorConfig: Specifies the type and number of accelerator cards
// attached to the instances of an instance. See GPUs on Compute Engine.
type AcceleratorConfig struct {
	// AcceleratorCount: The number of the accelerator cards of this type
	// exposed to this instance.
	AcceleratorCount int64 `json:"acceleratorCount,omitempty"`

	// AcceleratorTypeUri: Full URL, partial URI, or short name of the
	// accelerator type resource to expose to this instance. See Compute
	// Engine
	// AcceleratorTypes.Examples:
	// https://www.googleapis.com/compute/beta/pro
	// jects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80
	//
	// projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k
	// 80
	// nvidia-tesla-k80Auto Zone Exception: If you are using the Cloud
	// Dataproc Auto Zone Placement feature, you must use the short name of
	// the accelerator type resource, for example, nvidia-tesla-k80.
	AcceleratorTypeUri string `json:"acceleratorTypeUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AcceleratorCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AcceleratorCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AcceleratorConfig) MarshalJSON() ([]byte, error) {
	type NoMethod AcceleratorConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Binding: Associates members with a role.
type Binding struct {
	// Condition: Unimplemented. The condition that is associated with this
	// binding. NOTE: an unsatisfied condition will not allow user access
	// via current binding. Different bindings, including their conditions,
	// are examined independently.
	Condition *Expr `json:"condition,omitempty"`

	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource. members can have the following values:
	// allUsers: A special identifier that represents anyone who is  on the
	// internet; with or without a Google account.
	// allAuthenticatedUsers: A special identifier that represents anyone
	// who is authenticated with a Google account or a service
	// account.
	// user:{emailid}: An email address that represents a specific Google
	// account. For example, alice@gmail.com .
	// serviceAccount:{emailid}: An email address that represents a service
	// account. For example,
	// my-other-app@appspot.gserviceaccount.com.
	// group:{emailid}: An email address that represents a Google group.
	// For example, admins@example.com.
	// domain:{domain}: A Google Apps domain name that represents all the
	// users of that domain. For example, google.com or example.com.
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to members. For example, roles/viewer,
	// roles/editor, or roles/owner.
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Condition") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Condition") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Binding) MarshalJSON() ([]byte, error) {
	type NoMethod Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CancelJobRequest: A request to cancel a job.
type CancelJobRequest struct {
}

// Cluster: Describes the identifying information, config, and status of
// a cluster of Compute Engine instances.
type Cluster struct {
	// ClusterName: Required. The cluster name. Cluster names within a
	// project must be unique. Names of deleted clusters can be reused.
	ClusterName string `json:"clusterName,omitempty"`

	// ClusterUuid: Output only. A cluster UUID (Unique Universal
	// Identifier). Cloud Dataproc generates this value when it creates the
	// cluster.
	ClusterUuid string `json:"clusterUuid,omitempty"`

	// Config: Required. The cluster config. Note that Cloud Dataproc may
	// set default values, and values may change when clusters are updated.
	Config *ClusterConfig `json:"config,omitempty"`

	// Labels: Optional. The labels to associate with this cluster. Label
	// keys must contain 1 to 63 characters, and must conform to RFC 1035
	// (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty,
	// but, if present, must contain 1 to 63 characters, and must conform to
	// RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32
	// labels can be associated with a cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// Metrics: Contains cluster daemon metrics such as HDFS and YARN
	// stats.Beta Feature: This report is available for testing purposes
	// only. It may be changed before final release.
	Metrics *ClusterMetrics `json:"metrics,omitempty"`

	// ProjectId: Required. The Google Cloud Platform project ID that the
	// cluster belongs to.
	ProjectId string `json:"projectId,omitempty"`

	// Status: Output only. Cluster status.
	Status *ClusterStatus `json:"status,omitempty"`

	// StatusHistory: Output only. The previous cluster status.
	StatusHistory []*ClusterStatus `json:"statusHistory,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClusterName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusterName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Cluster) MarshalJSON() ([]byte, error) {
	type NoMethod Cluster
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClusterConfig: The cluster config.
type ClusterConfig struct {
	// ConfigBucket: Optional. A Cloud Storage staging bucket used for
	// sharing generated SSH keys and config. If you do not specify a
	// staging bucket, Cloud Dataproc will determine an appropriate Cloud
	// Storage location (US, ASIA, or EU) for your cluster's staging bucket
	// according to the Google Compute Engine zone where your cluster is
	// deployed, and then it will create and manage this project-level,
	// per-location bucket for you.
	ConfigBucket string `json:"configBucket,omitempty"`

	// EncryptionConfig: Optional. Encryption settings for the cluster.
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// GceClusterConfig: Required. The shared Compute Engine config settings
	// for all instances in a cluster.
	GceClusterConfig *GceClusterConfig `json:"gceClusterConfig,omitempty"`

	// InitializationActions: Optional. Commands to execute on each node
	// after config is completed. By default, executables are run on master
	// and all worker nodes. You can test a node's role metadata to run an
	// executable on a master or worker node, as shown below using curl (you
	// can also use wget):
	// ROLE=$(curl -H Metadata-Flavor:Google
	// http://metadata/computeMetadata/v1/instance/attributes/dataproc-role)
	//
	// if [[ "${ROLE}" == 'Master' ]]; then
	//   ... master specific actions ...
	// else
	//   ... worker specific actions ...
	// fi
	//
	InitializationActions []*NodeInitializationAction `json:"initializationActions,omitempty"`

	// MasterConfig: Optional. The Compute Engine config settings for the
	// master instance in a cluster.
	MasterConfig *InstanceGroupConfig `json:"masterConfig,omitempty"`

	// SecondaryWorkerConfig: Optional. The Compute Engine config settings
	// for additional worker instances in a cluster.
	SecondaryWorkerConfig *InstanceGroupConfig `json:"secondaryWorkerConfig,omitempty"`

	// SoftwareConfig: Optional. The config settings for software inside the
	// cluster.
	SoftwareConfig *SoftwareConfig `json:"softwareConfig,omitempty"`

	// WorkerConfig: Optional. The Compute Engine config settings for worker
	// instances in a cluster.
	WorkerConfig *InstanceGroupConfig `json:"workerConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConfigBucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConfigBucket") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClusterConfig) MarshalJSON() ([]byte, error) {
	type NoMethod ClusterConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClusterMetrics: Contains cluster daemon metrics, such as HDFS and
// YARN stats.Beta Feature: This report is available for testing
// purposes only. It may be changed before final release.
type ClusterMetrics struct {
	// HdfsMetrics: The HDFS metrics.
	HdfsMetrics map[string]string `json:"hdfsMetrics,omitempty"`

	// YarnMetrics: The YARN metrics.
	YarnMetrics map[string]string `json:"yarnMetrics,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HdfsMetrics") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HdfsMetrics") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClusterMetrics) MarshalJSON() ([]byte, error) {
	type NoMethod ClusterMetrics
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClusterOperation: The cluster operation triggered by a workflow.
type ClusterOperation struct {
	// Done: Output only. Indicates the operation is done.
	Done bool `json:"done,omitempty"`

	// Error: Output only. Error, if operation failed.
	Error string `json:"error,omitempty"`

	// OperationId: Output only. The id of the cluster operation.
	OperationId string `json:"operationId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClusterOperation) MarshalJSON() ([]byte, error) {
	type NoMethod ClusterOperation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClusterOperationMetadata: Metadata describing the operation.
type ClusterOperationMetadata struct {
	// ClusterName: Output only. Name of the cluster for the operation.
	ClusterName string `json:"clusterName,omitempty"`

	// ClusterUuid: Output only. Cluster UUID for the operation.
	ClusterUuid string `json:"clusterUuid,omitempty"`

	// Description: Output only. Short description of operation.
	Description string `json:"description,omitempty"`

	// Labels: Output only. Labels associated with the operation
	Labels map[string]string `json:"labels,omitempty"`

	// OperationType: Output only. The operation type.
	OperationType string `json:"operationType,omitempty"`

	// Status: Output only. Current operation status.
	Status *ClusterOperationStatus `json:"status,omitempty"`

	// StatusHistory: Output only. The previous operation status.
	StatusHistory []*ClusterOperationStatus `json:"statusHistory,omitempty"`

	// Warnings: Output only. Errors encountered during operation execution.
	Warnings []string `json:"warnings,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClusterName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusterName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClusterOperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod ClusterOperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClusterOperationStatus: The status of the operation.
type ClusterOperationStatus struct {
	// Details: Output only. A message containing any operation metadata
	// details.
	Details string `json:"details,omitempty"`

	// InnerState: Output only. A message containing the detailed operation
	// state.
	InnerState string `json:"innerState,omitempty"`

	// State: Output only. A message containing the operation state.
	//
	// Possible values:
	//   "UNKNOWN" - Unused.
	//   "PENDING" - The operation has been created.
	//   "RUNNING" - The operation is running.
	//   "DONE" - The operation is done; either cancelled or completed.
	State string `json:"state,omitempty"`

	// StateStartTime: Output only. The time this state was entered.
	StateStartTime string `json:"stateStartTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Details") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Details") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClusterOperationStatus) MarshalJSON() ([]byte, error) {
	type NoMethod ClusterOperationStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClusterSelector: A selector that chooses target cluster for jobs
// based on metadata.
type ClusterSelector struct {
	// ClusterLabels: Required. The cluster labels. Cluster must have all
	// labels to match.
	ClusterLabels map[string]string `json:"clusterLabels,omitempty"`

	// Zone: Optional. The zone where workflow process executes. This
	// parameter does not affect the selection of the cluster.If
	// unspecified, the zone of the first cluster matching the selector is
	// used.
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClusterLabels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusterLabels") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClusterSelector) MarshalJSON() ([]byte, error) {
	type NoMethod ClusterSelector
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClusterStatus: The status of a cluster and its instances.
type ClusterStatus struct {
	// Detail: Output only. Optional details of cluster's state.
	Detail string `json:"detail,omitempty"`

	// State: Output only. The cluster's state.
	//
	// Possible values:
	//   "UNKNOWN" - The cluster state is unknown.
	//   "CREATING" - The cluster is being created and set up. It is not
	// ready for use.
	//   "RUNNING" - The cluster is currently running and healthy. It is
	// ready for use.
	//   "ERROR" - The cluster encountered an error. It is not ready for
	// use.
	//   "DELETING" - The cluster is being deleted. It cannot be used.
	//   "UPDATING" - The cluster is being updated. It continues to accept
	// and process jobs.
	State string `json:"state,omitempty"`

	// StateStartTime: Output only. Time when this state was entered.
	StateStartTime string `json:"stateStartTime,omitempty"`

	// Substate: Output only. Additional state information that includes
	// status reported by the agent.
	//
	// Possible values:
	//   "UNSPECIFIED" - The cluster substate is unknown.
	//   "UNHEALTHY" - The cluster is known to be in an unhealthy state (for
	// example, critical daemons are not running or HDFS capacity is
	// exhausted).Applies to RUNNING state.
	//   "STALE_STATUS" - The agent-reported status is out of date (may
	// occur if Cloud Dataproc loses communication with Agent).Applies to
	// RUNNING state.
	Substate string `json:"substate,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Detail") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Detail") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClusterStatus) MarshalJSON() ([]byte, error) {
	type NoMethod ClusterStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DiagnoseClusterRequest: A request to collect cluster diagnostic
// information.
type DiagnoseClusterRequest struct {
}

// DiagnoseClusterResults: The location of diagnostic output.
type DiagnoseClusterResults struct {
	// OutputUri: Output only. The Cloud Storage URI of the diagnostic
	// output. The output report is a plain text file with a summary of
	// collected diagnostics.
	OutputUri string `json:"outputUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OutputUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OutputUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DiagnoseClusterResults) MarshalJSON() ([]byte, error) {
	type NoMethod DiagnoseClusterResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DiskConfig: Specifies the config of disk options for a group of VM
// instances.
type DiskConfig struct {
	// BootDiskSizeGb: Optional. Size in GB of the boot disk (default is
	// 500GB).
	BootDiskSizeGb int64 `json:"bootDiskSizeGb,omitempty"`

	// BootDiskType: Optional. Type of the boot disk (default is
	// "pd-standard"). Valid values: "pd-ssd" (Persistent Disk Solid State
	// Drive) or "pd-standard" (Persistent Disk Hard Disk Drive).
	BootDiskType string `json:"bootDiskType,omitempty"`

	// NumLocalSsds: Optional. Number of attached SSDs, from 0 to 4 (default
	// is 0). If SSDs are not attached, the boot disk is used to store
	// runtime logs and HDFS
	// (https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If
	// one or more SSDs are attached, this runtime bulk data is spread
	// across them, and the boot disk contains only basic config and
	// installed binaries.
	NumLocalSsds int64 `json:"numLocalSsds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BootDiskSizeGb") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BootDiskSizeGb") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DiskConfig) MarshalJSON() ([]byte, error) {
	type NoMethod DiskConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use
// it as the request or the response type of an API method. For
// instance:
// service Foo {
//   rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
// }
// The JSON representation for Empty is empty JSON object {}.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// EncryptionConfig: Encryption settings for the cluster.
type EncryptionConfig struct {
	// GcePdKmsKeyName: Optional. The Cloud KMS key name to use for PD disk
	// encryption for all instances in the cluster.
	GcePdKmsKeyName string `json:"gcePdKmsKeyName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GcePdKmsKeyName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GcePdKmsKeyName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *EncryptionConfig) MarshalJSON() ([]byte, error) {
	type NoMethod EncryptionConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Expr: Represents an expression text. Example:
// title: "User account presence"
// description: "Determines whether the request has a user
// account"
// expression: "size(request.user) > 0"
//
type Expr struct {
	// Description: An optional description of the expression. This is a
	// longer text which describes the expression, e.g. when hovered over it
	// in a UI.
	Description string `json:"description,omitempty"`

	// Expression: Textual representation of an expression in Common
	// Expression Language syntax.The application context of the containing
	// message determines which well-known feature set of CEL is supported.
	Expression string `json:"expression,omitempty"`

	// Location: An optional string indicating the location of the
	// expression for error reporting, e.g. a file name and a position in
	// the file.
	Location string `json:"location,omitempty"`

	// Title: An optional title for the expression, i.e. a short string
	// describing its purpose. This can be used e.g. in UIs which allow to
	// enter the expression.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Expr) MarshalJSON() ([]byte, error) {
	type NoMethod Expr
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GceClusterConfig: Common config settings for resources of Compute
// Engine cluster instances, applicable to all instances in the cluster.
type GceClusterConfig struct {
	// InternalIpOnly: Optional. If true, all instances in the cluster will
	// only have internal IP addresses. By default, clusters are not
	// restricted to internal IP addresses, and will have ephemeral external
	// IP addresses assigned to each instance. This internal_ip_only
	// restriction can only be enabled for subnetwork enabled networks, and
	// all off-cluster dependencies must be configured to be accessible
	// without external IP addresses.
	InternalIpOnly bool `json:"internalIpOnly,omitempty"`

	// Metadata: The Compute Engine metadata entries to add to all instances
	// (see Project and instance metadata
	// (https://cloud.google.com/compute/docs/storing-retrieving-metadata#pro
	// ject_and_instance_metadata)).
	Metadata map[string]string `json:"metadata,omitempty"`

	// NetworkUri: Optional. The Compute Engine network to be used for
	// machine communications. Cannot be specified with subnetwork_uri. If
	// neither network_uri nor subnetwork_uri is specified, the "default"
	// network of the project is used, if it exists. Cannot be a "Custom
	// Subnet Network" (see Using Subnetworks for more information).A full
	// URL, partial URI, or short name are valid.
	// Examples:
	// https://www.googleapis.com/compute/v1/projects/[project_id]/
	// regions/global/default
	// projects/[project_id]/regions/global/default
	// de
	// fault
	NetworkUri string `json:"networkUri,omitempty"`

	// ServiceAccount: Optional. The service account of the instances.
	// Defaults to the default Compute Engine service account. Custom
	// service accounts need permissions equivalent to the following IAM
	// roles:
	// roles/logging.logWriter
	// roles/storage.objectAdmin(see
	// https://cloud.google.com/compute/docs/access/service-accounts#custom_service_accounts for more information). Example:
	// [account_id]@[project_id].iam.gserviceaccount.com
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// ServiceAccountScopes: Optional. The URIs of service account scopes to
	// be included in Compute Engine instances. The following base set of
	// scopes is always
	// included:
	// https://www.googleapis.com/auth/cloud.useraccounts.readonly
	//
	// https://www.googleapis.com/auth/devstorage.read_write
	// https://www.goog
	// leapis.com/auth/logging.writeIf no scopes are specified, the
	// following defaults are also
	// provided:
	// https://www.googleapis.com/auth/bigquery
	// https://www.googlea
	// pis.com/auth/bigtable.admin.table
	// https://www.googleapis.com/auth/bigt
	// able.data
	// https://www.googleapis.com/auth/devstorage.full_control
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	// SubnetworkUri: Optional. The Compute Engine subnetwork to be used for
	// machine communications. Cannot be specified with network_uri.A full
	// URL, partial URI, or short name are valid.
	// Examples:
	// https://www.googleapis.com/compute/v1/projects/[project_id]/
	// regions/us-east1/sub0
	// projects/[project_id]/regions/us-east1/sub0
	// sub0
	SubnetworkUri string `json:"subnetworkUri,omitempty"`

	// Tags: The Compute Engine tags to add to all instances (see Tagging
	// instances).
	Tags []string `json:"tags,omitempty"`

	// ZoneUri: Optional. The zone where the Compute Engine cluster will be
	// located. On a create request, it is required in the "global" region.
	// If omitted in a non-global Cloud Dataproc region, the service will
	// pick a zone in the corresponding Compute Engine region. On a get
	// request, zone will always be present.A full URL, partial URI, or
	// short name are valid.
	// Examples:
	// https://www.googleapis.com/compute/v1/projects/[project_id]/
	// zones/[zone]
	// projects/[project_id]/zones/[zone]
	// us-central1-f
	ZoneUri string `json:"zoneUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InternalIpOnly") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InternalIpOnly") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GceClusterConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GceClusterConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetIamPolicyRequest: Request message for GetIamPolicy method.
type GetIamPolicyRequest struct {
}

// HadoopJob: A Cloud Dataproc job for running Apache Hadoop MapReduce
// (https://hadoop.apache.org/docs/current/hadoop-mapreduce-client/hadoop
// -mapreduce-client-core/MapReduceTutorial.html) jobs on Apache Hadoop
// YARN
// (https://hadoop.apache.org/docs/r2.7.1/hadoop-yarn/hadoop-yarn-site/YA
// RN.html).
type HadoopJob struct {
	// ArchiveUris: Optional. HCFS URIs of archives to be extracted in the
	// working directory of Hadoop drivers and tasks. Supported file types:
	// .jar, .tar, .tar.gz, .tgz, or .zip.
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Args: Optional. The arguments to pass to the driver. Do not include
	// arguments, such as -libjars or -Dfoo=bar, that can be set as job
	// properties, since a collision may occur that causes an incorrect job
	// submission.
	Args []string `json:"args,omitempty"`

	// FileUris: Optional. HCFS (Hadoop Compatible Filesystem) URIs of files
	// to be copied to the working directory of Hadoop drivers and
	// distributed tasks. Useful for naively parallel tasks.
	FileUris []string `json:"fileUris,omitempty"`

	// JarFileUris: Optional. Jar file URIs to add to the CLASSPATHs of the
	// Hadoop driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfig: Optional. The runtime log config for job execution.
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// MainClass: The name of the driver's main class. The jar file
	// containing the class must be in the default CLASSPATH or specified in
	// jar_file_uris.
	MainClass string `json:"mainClass,omitempty"`

	// MainJarFileUri: The HCFS URI of the jar file containing the main
	// class. Examples:
	// 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar'
	// 'hdfs:/tmp/test-samples/custom-wordcount.jar'
	// 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'
	MainJarFileUri string `json:"mainJarFileUri,omitempty"`

	// Properties: Optional. A mapping of property names to values, used to
	// configure Hadoop. Properties that conflict with values set by the
	// Cloud Dataproc API may be overwritten. Can include properties set in
	// /etc/hadoop/conf/*-site and classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ArchiveUris") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArchiveUris") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HadoopJob) MarshalJSON() ([]byte, error) {
	type NoMethod HadoopJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HiveJob: A Cloud Dataproc job for running Apache Hive
// (https://hive.apache.org/) queries on YARN.
type HiveJob struct {
	// ContinueOnFailure: Optional. Whether to continue executing queries if
	// a query fails. The default value is false. Setting to true can be
	// useful when executing independent parallel queries.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty"`

	// JarFileUris: Optional. HCFS URIs of jar files to add to the CLASSPATH
	// of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive
	// SerDes and UDFs.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Properties: Optional. A mapping of property names and values, used to
	// configure Hive. Properties that conflict with values set by the Cloud
	// Dataproc API may be overwritten. Can include properties set in
	// /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and
	// classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// QueryFileUri: The HCFS URI of the script that contains Hive queries.
	QueryFileUri string `json:"queryFileUri,omitempty"`

	// QueryList: A list of queries.
	QueryList *QueryList `json:"queryList,omitempty"`

	// ScriptVariables: Optional. Mapping of query variable names to values
	// (equivalent to the Hive command: SET name="value";).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContinueOnFailure")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContinueOnFailure") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *HiveJob) MarshalJSON() ([]byte, error) {
	type NoMethod HiveJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InstanceGroupConfig: Optional. The config settings for Compute Engine
// resources in an instance group, such as a master or worker group.
type InstanceGroupConfig struct {
	// Accelerators: Optional. The Compute Engine accelerator configuration
	// for these instances.Beta Feature: This feature is still under
	// development. It may be changed before final release.
	Accelerators []*AcceleratorConfig `json:"accelerators,omitempty"`

	// DiskConfig: Optional. Disk option config settings.
	DiskConfig *DiskConfig `json:"diskConfig,omitempty"`

	// ImageUri: Optional. The Compute Engine image resource used for
	// cluster instances. It can be specified or may be inferred from
	// SoftwareConfig.image_version.
	ImageUri string `json:"imageUri,omitempty"`

	// InstanceNames: Output only. The list of instance names. Cloud
	// Dataproc derives the names from cluster_name, num_instances, and the
	// instance group.
	InstanceNames []string `json:"instanceNames,omitempty"`

	// IsPreemptible: Optional. Specifies that this instance group contains
	// preemptible instances.
	IsPreemptible bool `json:"isPreemptible,omitempty"`

	// MachineTypeUri: Optional. The Compute Engine machine type used for
	// cluster instances.A full URL, partial URI, or short name are valid.
	// Examples:
	// https://www.googleapis.com/compute/v1/projects/[project_id]/
	// zones/us-east1-a/machineTypes/n1-standard-2
	// projects/[project_id]/zone
	// s/us-east1-a/machineTypes/n1-standard-2
	// n1-standard-2Auto Zone Exception: If you are using the Cloud Dataproc
	// Auto Zone Placement feature, you must use the short name of the
	// machine type resource, for example, n1-standard-2.
	MachineTypeUri string `json:"machineTypeUri,omitempty"`

	// ManagedGroupConfig: Output only. The config for Compute Engine
	// Instance Group Manager that manages this group. This is only used for
	// preemptible instance groups.
	ManagedGroupConfig *ManagedGroupConfig `json:"managedGroupConfig,omitempty"`

	// NumInstances: Optional. The number of VM instances in the instance
	// group. For master instance groups, must be set to 1.
	NumInstances int64 `json:"numInstances,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Accelerators") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Accelerators") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *InstanceGroupConfig) MarshalJSON() ([]byte, error) {
	type NoMethod InstanceGroupConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InstantiateWorkflowTemplateRequest: A request to instantiate a
// workflow template.
type InstantiateWorkflowTemplateRequest struct {
	// Parameters: Optional. Map from parameter names to values that should
	// be used for those parameters. Values may not exceed 100 characters.
	Parameters map[string]string `json:"parameters,omitempty"`

	// RequestId: Optional. A tag that prevents multiple concurrent workflow
	// instances with the same tag from running. This mitigates risk of
	// concurrent instances started due to retries.It is recommended to
	// always set this value to a UUID
	// (https://en.wikipedia.org/wiki/Universally_unique_identifier).The tag
	// must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). The maximum length is 40 characters.
	RequestId string `json:"requestId,omitempty"`

	// Version: Optional. The version of workflow template to instantiate.
	// If specified, the workflow will be instantiated only if the current
	// version of the workflow template has the supplied version.This option
	// cannot be used to instantiate a previous version of workflow
	// template.
	Version int64 `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Parameters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Parameters") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *InstantiateWorkflowTemplateRequest) MarshalJSON() ([]byte, error) {
	type NoMethod InstantiateWorkflowTemplateRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Job: A Cloud Dataproc job resource.
type Job struct {
	// DriverControlFilesUri: Output only. If present, the location of
	// miscellaneous control files which may be used as part of job setup
	// and handling. If not present, control files may be placed in the same
	// location as driver_output_uri.
	DriverControlFilesUri string `json:"driverControlFilesUri,omitempty"`

	// DriverOutputResourceUri: Output only. A URI pointing to the location
	// of the stdout of the job's driver program.
	DriverOutputResourceUri string `json:"driverOutputResourceUri,omitempty"`

	// HadoopJob: Job is a Hadoop job.
	HadoopJob *HadoopJob `json:"hadoopJob,omitempty"`

	// HiveJob: Job is a Hive job.
	HiveJob *HiveJob `json:"hiveJob,omitempty"`

	// JobUuid: Output only. A UUID that uniquely identifies a job within
	// the project over time. This is in contrast to a user-settable
	// reference.job_id that may be reused over time.
	JobUuid string `json:"jobUuid,omitempty"`

	// Labels: Optional. The labels to associate with this job. Label keys
	// must contain 1 to 63 characters, and must conform to RFC 1035
	// (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty,
	// but, if present, must contain 1 to 63 characters, and must conform to
	// RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32
	// labels can be associated with a job.
	Labels map[string]string `json:"labels,omitempty"`

	// PigJob: Job is a Pig job.
	PigJob *PigJob `json:"pigJob,omitempty"`

	// Placement: Required. Job information, including how, when, and where
	// to run the job.
	Placement *JobPlacement `json:"placement,omitempty"`

	// PysparkJob: Job is a Pyspark job.
	PysparkJob *PySparkJob `json:"pysparkJob,omitempty"`

	// Reference: Optional. The fully qualified reference to the job, which
	// can be used to obtain the equivalent REST path of the job resource.
	// If this property is not specified when a job is created, the server
	// generates a <code>job_id</code>.
	Reference *JobReference `json:"reference,omitempty"`

	// Scheduling: Optional. Job scheduling configuration.
	Scheduling *JobScheduling `json:"scheduling,omitempty"`

	// SparkJob: Job is a Spark job.
	SparkJob *SparkJob `json:"sparkJob,omitempty"`

	// SparkSqlJob: Job is a SparkSql job.
	SparkSqlJob *SparkSqlJob `json:"sparkSqlJob,omitempty"`

	// Status: Output only. The job status. Additional application-specific
	// status information may be contained in the <code>type_job</code> and
	// <code>yarn_applications</code> fields.
	Status *JobStatus `json:"status,omitempty"`

	// StatusHistory: Output only. The previous job status.
	StatusHistory []*JobStatus `json:"statusHistory,omitempty"`

	// YarnApplications: Output only. The collection of YARN applications
	// spun up by this job.Beta Feature: This report is available for
	// testing purposes only. It may be changed before final release.
	YarnApplications []*YarnApplication `json:"yarnApplications,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "DriverControlFilesUri") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DriverControlFilesUri") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Job) MarshalJSON() ([]byte, error) {
	type NoMethod Job
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// JobPlacement: Cloud Dataproc job config.
type JobPlacement struct {
	// ClusterName: Required. The name of the cluster where the job will be
	// submitted.
	ClusterName string `json:"clusterName,omitempty"`

	// ClusterUuid: Output only. A cluster UUID generated by the Cloud
	// Dataproc service when the job is submitted.
	ClusterUuid string `json:"clusterUuid,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClusterName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusterName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *JobPlacement) MarshalJSON() ([]byte, error) {
	type NoMethod JobPlacement
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// JobReference: Encapsulates the full scoping used to reference a job.
type JobReference struct {
	// JobId: Optional. The job ID, which must be unique within the project.
	// The job ID is generated by the server upon job submission or provided
	// by the user as a means to perform retries without creating duplicate
	// jobs. The ID must contain only letters (a-z, A-Z), numbers (0-9),
	// underscores (_), or hyphens (-). The maximum length is 100
	// characters.
	JobId string `json:"jobId,omitempty"`

	// ProjectId: Required. The ID of the Google Cloud Platform project that
	// the job belongs to.
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "JobId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "JobId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *JobReference) MarshalJSON() ([]byte, error) {
	type NoMethod JobReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// JobScheduling: Job scheduling options.
type JobScheduling struct {
	// MaxFailuresPerHour: Optional. Maximum number of times per hour a
	// driver may be restarted as a result of driver terminating with
	// non-zero code before job is reported failed.A job may be reported as
	// thrashing if driver exits with non-zero code 4 times within 10 minute
	// window.Maximum value is 10.
	MaxFailuresPerHour int64 `json:"maxFailuresPerHour,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxFailuresPerHour")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaxFailuresPerHour") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *JobScheduling) MarshalJSON() ([]byte, error) {
	type NoMethod JobScheduling
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// JobStatus: Cloud Dataproc job status.
type JobStatus struct {
	// Details: Output only. Optional job state details, such as an error
	// description if the state is <code>ERROR</code>.
	Details string `json:"details,omitempty"`

	// State: Output only. A state message specifying the overall job state.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The job state is unknown.
	//   "PENDING" - The job is pending; it has been submitted, but is not
	// yet running.
	//   "SETUP_DONE" - Job has been received by the service and completed
	// initial setup; it will soon be submitted to the cluster.
	//   "RUNNING" - The job is running on the cluster.
	//   "CANCEL_PENDING" - A CancelJob request has been received, but is
	// pending.
	//   "CANCEL_STARTED" - Transient in-flight resources have been
	// canceled, and the request to cancel the running job has been issued
	// to the cluster.
	//   "CANCELLED" - The job cancellation was successful.
	//   "DONE" - The job has completed successfully.
	//   "ERROR" - The job has completed, but encountered an error.
	//   "ATTEMPT_FAILURE" - Job attempt has failed. The detail field
	// contains failure details for this attempt.Applies to restartable jobs
	// only.
	State string `json:"state,omitempty"`

	// StateStartTime: Output only. The time when this state was entered.
	StateStartTime string `json:"stateStartTime,omitempty"`

	// Substate: Output only. Additional state information, which includes
	// status reported by the agent.
	//
	// Possible values:
	//   "UNSPECIFIED" - The job substate is unknown.
	//   "SUBMITTED" - The Job is submitted to the agent.Applies to RUNNING
	// state.
	//   "QUEUED" - The Job has been received and is awaiting execution (it
	// may be waiting for a condition to be met). See the "details" field
	// for the reason for the delay.Applies to RUNNING state.
	//   "STALE_STATUS" - The agent-reported status is out of date, which
	// may be caused by a loss of communication between the agent and Cloud
	// Dataproc. If the agent does not send a timely update, the job will
	// fail.Applies to RUNNING state.
	Substate string `json:"substate,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Details") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Details") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *JobStatus) MarshalJSON() ([]byte, error) {
	type NoMethod JobStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListClustersResponse: The list of all clusters in a project.
type ListClustersResponse struct {
	// Clusters: Output only. The clusters in the project.
	Clusters []*Cluster `json:"clusters,omitempty"`

	// NextPageToken: Output only. This token is included in the response if
	// there are more results to fetch. To fetch additional results, provide
	// this value as the page_token in a subsequent ListClustersRequest.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Clusters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Clusters") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListClustersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListClustersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListJobsResponse: A list of jobs in a project.
type ListJobsResponse struct {
	// Jobs: Output only. Jobs list.
	Jobs []*Job `json:"jobs,omitempty"`

	// NextPageToken: Optional. This token is included in the response if
	// there are more results to fetch. To fetch additional results, provide
	// this value as the page_token in a subsequent
	// <code>ListJobsRequest</code>.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Jobs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Jobs") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListJobsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListJobsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOperationsResponse: The response message for
// Operations.ListOperations.
type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListOperationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListOperationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListWorkflowTemplatesResponse: A response to a request to list
// workflow templates in a project.
type ListWorkflowTemplatesResponse struct {
	// NextPageToken: Output only. This token is included in the response if
	// there are more results to fetch. To fetch additional results, provide
	// this value as the page_token in a subsequent
	// <code>ListWorkflowTemplatesRequest</code>.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Templates: Output only. WorkflowTemplates list.
	Templates []*WorkflowTemplate `json:"templates,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListWorkflowTemplatesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListWorkflowTemplatesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LoggingConfig: The runtime logging config of the job.
type LoggingConfig struct {
	// DriverLogLevels: The per-package log levels for the driver. This may
	// include "root" package name to configure rootLogger. Examples:
	// 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DriverLogLevels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DriverLogLevels") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LoggingConfig) MarshalJSON() ([]byte, error) {
	type NoMethod LoggingConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ManagedCluster: Cluster that is managed by the workflow.
type ManagedCluster struct {
	// ClusterName: Required. The cluster name prefix. A unique cluster name
	// will be formed by appending a random suffix.The name must contain
	// only lower-case letters (a-z), numbers (0-9), and hyphens (-). Must
	// begin with a letter. Cannot begin or end with hyphen. Must consist of
	// between 2 and 35 characters.
	ClusterName string `json:"clusterName,omitempty"`

	// Config: Required. The cluster configuration.
	Config *ClusterConfig `json:"config,omitempty"`

	// Labels: Optional. The labels to associate with this cluster.Label
	// keys must be between 1 and 63 characters long, and must conform to
	// the following PCRE regular expression: \p{Ll}\p{Lo}{0,62}Label values
	// must be between 1 and 63 characters long, and must conform to the
	// following PCRE regular expression: \p{Ll}\p{Lo}\p{N}_-{0,63}No more
	// than 32 labels can be associated with a given cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClusterName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusterName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ManagedCluster) MarshalJSON() ([]byte, error) {
	type NoMethod ManagedCluster
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ManagedGroupConfig: Specifies the resources used to actively manage
// an instance group.
type ManagedGroupConfig struct {
	// InstanceGroupManagerName: Output only. The name of the Instance Group
	// Manager for this group.
	InstanceGroupManagerName string `json:"instanceGroupManagerName,omitempty"`

	// InstanceTemplateName: Output only. The name of the Instance Template
	// used for the Managed Instance Group.
	InstanceTemplateName string `json:"instanceTemplateName,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "InstanceGroupManagerName") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InstanceGroupManagerName")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ManagedGroupConfig) MarshalJSON() ([]byte, error) {
	type NoMethod ManagedGroupConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NodeInitializationAction: Specifies an executable to run on a fully
// configured node and a timeout period for executable completion.
type NodeInitializationAction struct {
	// ExecutableFile: Required. Cloud Storage URI of executable file.
	ExecutableFile string `json:"executableFile,omitempty"`

	// ExecutionTimeout: Optional. Amount of time executable has to
	// complete. Default is 10 minutes. Cluster creation fails with an
	// explanatory error message (the name of the executable that caused the
	// error and the exceeded timeout period) if the executable is not
	// completed at end of the timeout period.
	ExecutionTimeout string `json:"executionTimeout,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExecutableFile") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExecutableFile") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *NodeInitializationAction) MarshalJSON() ([]byte, error) {
	type NoMethod NodeInitializationAction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a network API call.
type Operation struct {
	// Done: If the value is false, it means the operation is still in
	// progress. If true, the operation is completed, and either error or
	// response is available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation. It
	// typically contains progress information and common metadata such as
	// create time. Some services might not provide such metadata. Any
	// method that returns a long-running operation should document the
	// metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that originally returns it. If you use the default HTTP
	// mapping, the name should have the format of
	// operations/some/unique/name.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success. If
	// the original method returns no data on success, such as Delete, the
	// response is google.protobuf.Empty. If the original method is standard
	// Get/Create/Update, the response should be the resource. For other
	// methods, the response should have the type XxxResponse, where Xxx is
	// the original method name. For example, if the original method name is
	// TakeSnapshot(), the inferred response type is TakeSnapshotResponse.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OrderedJob: A job executed by the workflow.
type OrderedJob struct {
	// HadoopJob: Job is a Hadoop job.
	HadoopJob *HadoopJob `json:"hadoopJob,omitempty"`

	// HiveJob: Job is a Hive job.
	HiveJob *HiveJob `json:"hiveJob,omitempty"`

	// Labels: Optional. The labels to associate with this job.Label keys
	// must be between 1 and 63 characters long, and must conform to the
	// following regular expression: \p{Ll}\p{Lo}{0,62}Label values must be
	// between 1 and 63 characters long, and must conform to the following
	// regular expression: \p{Ll}\p{Lo}\p{N}_-{0,63}No more than 32 labels
	// can be associated with a given job.
	Labels map[string]string `json:"labels,omitempty"`

	// PigJob: Job is a Pig job.
	PigJob *PigJob `json:"pigJob,omitempty"`

	// PrerequisiteStepIds: Optional. The optional list of prerequisite job
	// step_ids. If not specified, the job will start at the beginning of
	// workflow.
	PrerequisiteStepIds []string `json:"prerequisiteStepIds,omitempty"`

	// PysparkJob: Job is a Pyspark job.
	PysparkJob *PySparkJob `json:"pysparkJob,omitempty"`

	// Scheduling: Optional. Job scheduling configuration.
	Scheduling *JobScheduling `json:"scheduling,omitempty"`

	// SparkJob: Job is a Spark job.
	SparkJob *SparkJob `json:"sparkJob,omitempty"`

	// SparkSqlJob: Job is a SparkSql job.
	SparkSqlJob *SparkSqlJob `json:"sparkSqlJob,omitempty"`

	// StepId: Required. The step id. The id must be unique among all jobs
	// within the template.The step id is used as prefix for job id, as job
	// goog-dataproc-workflow-step-id label, and in prerequisiteStepIds
	// field from other steps.The id must contain only letters (a-z, A-Z),
	// numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end
	// with underscore or hyphen. Must consist of between 3 and 50
	// characters.
	StepId string `json:"stepId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HadoopJob") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HadoopJob") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OrderedJob) MarshalJSON() ([]byte, error) {
	type NoMethod OrderedJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ParameterValidation: Configuration for parameter validation.
type ParameterValidation struct {
	// Regex: Validation based on regular expressions.
	Regex *RegexValidation `json:"regex,omitempty"`

	// Values: Validation based on a list of allowed values.
	Values *ValueValidation `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Regex") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Regex") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ParameterValidation) MarshalJSON() ([]byte, error) {
	type NoMethod ParameterValidation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PigJob: A Cloud Dataproc job for running Apache Pig
// (https://pig.apache.org/) queries on YARN.
type PigJob struct {
	// ContinueOnFailure: Optional. Whether to continue executing queries if
	// a query fails. The default value is false. Setting to true can be
	// useful when executing independent parallel queries.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty"`

	// JarFileUris: Optional. HCFS URIs of jar files to add to the CLASSPATH
	// of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig
	// UDFs.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfig: Optional. The runtime log config for job execution.
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Properties: Optional. A mapping of property names to values, used to
	// configure Pig. Properties that conflict with values set by the Cloud
	// Dataproc API may be overwritten. Can include properties set in
	// /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties, and
	// classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// QueryFileUri: The HCFS URI of the script that contains the Pig
	// queries.
	QueryFileUri string `json:"queryFileUri,omitempty"`

	// QueryList: A list of queries.
	QueryList *QueryList `json:"queryList,omitempty"`

	// ScriptVariables: Optional. Mapping of query variable names to values
	// (equivalent to the Pig command: name=[value]).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContinueOnFailure")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContinueOnFailure") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PigJob) MarshalJSON() ([]byte, error) {
	type NoMethod PigJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Policy: Defines an Identity and Access Management (IAM) policy. It is
// used to specify access control policies for Cloud Platform
// resources.A Policy consists of a list of bindings. A binding binds a
// list of members to a role, where the members can be user accounts,
// Google groups, Google domains, and service accounts. A role is a
// named list of permissions defined by IAM.JSON Example
// {
//   "bindings": [
//     {
//       "role": "roles/owner",
//       "members": [
//         "user:mike@example.com",
//         "group:admins@example.com",
//         "domain:google.com",
//         "serviceAccount:my-other-app@appspot.gserviceaccount.com"
//       ]
//     },
//     {
//       "role": "roles/viewer",
//       "members": ["user:sean@example.com"]
//     }
//   ]
// }
// YAML Example
// bindings:
// - members:
//   - user:mike@example.com
//   - group:admins@example.com
//   - domain:google.com
//   - serviceAccount:my-other-app@appspot.gserviceaccount.com
//   role: roles/owner
// - members:
//   - user:sean@example.com
//   role: roles/viewer
// For a description of IAM and its features, see the IAM developer's
// guide (https://cloud.google.com/iam/docs).
type Policy struct {
	// Bindings: Associates a list of members to a role. bindings with no
	// members will result in an error.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: etag is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a policy from overwriting each
	// other. It is strongly suggested that systems make use of the etag in
	// the read-modify-write cycle to perform policy updates in order to
	// avoid race conditions: An etag is returned in the response to
	// getIamPolicy, and systems are expected to put that etag in the
	// request to setIamPolicy to ensure that their change will be applied
	// to the same version of the policy.If no etag is provided in the call
	// to setIamPolicy, then the existing policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	// Version: Deprecated.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bindings") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindings") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type NoMethod Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PySparkJob: A Cloud Dataproc job for running Apache PySpark
// (https://spark.apache.org/docs/0.9.0/python-programming-guide.html)
// applications on YARN.
type PySparkJob struct {
	// ArchiveUris: Optional. HCFS URIs of archives to be extracted in the
	// working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Args: Optional. The arguments to pass to the driver. Do not include
	// arguments, such as --conf, that can be set as job properties, since a
	// collision may occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty"`

	// FileUris: Optional. HCFS URIs of files to be copied to the working
	// directory of Python drivers and distributed tasks. Useful for naively
	// parallel tasks.
	FileUris []string `json:"fileUris,omitempty"`

	// JarFileUris: Optional. HCFS URIs of jar files to add to the
	// CLASSPATHs of the Python driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfig: Optional. The runtime log config for job execution.
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// MainPythonFileUri: Required. The HCFS URI of the main Python file to
	// use as the driver. Must be a .py file.
	MainPythonFileUri string `json:"mainPythonFileUri,omitempty"`

	// Properties: Optional. A mapping of property names to values, used to
	// configure PySpark. Properties that conflict with values set by the
	// Cloud Dataproc API may be overwritten. Can include properties set in
	// /etc/spark/conf/spark-defaults.conf and classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// PythonFileUris: Optional. HCFS file URIs of Python files to pass to
	// the PySpark framework. Supported file types: .py, .egg, and .zip.
	PythonFileUris []string `json:"pythonFileUris,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ArchiveUris") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArchiveUris") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PySparkJob) MarshalJSON() ([]byte, error) {
	type NoMethod PySparkJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryList: A list of queries to run on a cluster.
type QueryList struct {
	// Queries: Required. The queries to execute. You do not need to
	// terminate a query with a semicolon. Multiple queries can be specified
	// in one string by separating each with a semicolon. Here is an example
	// of an Cloud Dataproc API snippet that uses a QueryList to specify a
	// HiveJob:
	// "hiveJob": {
	//   "queryList": {
	//     "queries": [
	//       "query1",
	//       "query2",
	//       "query3;query4",
	//     ]
	//   }
	// }
	//
	Queries []string `json:"queries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Queries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Queries") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryList) MarshalJSON() ([]byte, error) {
	type NoMethod QueryList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RegexValidation: Validation based on regular expressions.
type RegexValidation struct {
	// Regexes: Required. RE2 regular expressions used to validate the
	// parameter's value. The value must match the regex in its entirety
	// (substring matches are not sufficient).
	Regexes []string `json:"regexes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Regexes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Regexes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RegexValidation) MarshalJSON() ([]byte, error) {
	type NoMethod RegexValidation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SetIamPolicyRequest: Request message for SetIamPolicy method.
type SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the resource.
	// The size of the policy is limited to a few 10s of KB. An empty policy
	// is a valid policy but certain Cloud Platform services (such as
	// Projects) might reject them.
	Policy *Policy `json:"policy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SetIamPolicyRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SoftwareConfig: Specifies the selection and config of software inside
// the cluster.
type SoftwareConfig struct {
	// ImageVersion: Optional. The version of software inside the cluster.
	// It must be one of the supported Cloud Dataproc Versions, such as
	// "1.2" (including a subminor version, such as "1.2.29"), or the
	// "preview" version. If unspecified, it defaults to the latest version.
	ImageVersion string `json:"imageVersion,omitempty"`

	// Properties: Optional. The properties to set on daemon config
	// files.Property keys are specified in prefix:property format, such as
	// core:fs.defaultFS. The following are supported prefixes and their
	// mappings:
	// capacity-scheduler: capacity-scheduler.xml
	// core: core-site.xml
	// distcp: distcp-default.xml
	// hdfs: hdfs-site.xml
	// hive: hive-site.xml
	// mapred: mapred-site.xml
	// pig: pig.properties
	// spark: spark-defaults.conf
	// yarn: yarn-site.xmlFor more information, see Cluster properties.
	Properties map[string]string `json:"properties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageVersion") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SoftwareConfig) MarshalJSON() ([]byte, error) {
	type NoMethod SoftwareConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SparkJob: A Cloud Dataproc job for running Apache Spark
// (http://spark.apache.org/) applications on YARN.
type SparkJob struct {
	// ArchiveUris: Optional. HCFS URIs of archives to be extracted in the
	// working directory of Spark drivers and tasks. Supported file types:
	// .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Args: Optional. The arguments to pass to the driver. Do not include
	// arguments, such as --conf, that can be set as job properties, since a
	// collision may occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty"`

	// FileUris: Optional. HCFS URIs of files to be copied to the working
	// directory of Spark drivers and distributed tasks. Useful for naively
	// parallel tasks.
	FileUris []string `json:"fileUris,omitempty"`

	// JarFileUris: Optional. HCFS URIs of jar files to add to the
	// CLASSPATHs of the Spark driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfig: Optional. The runtime log config for job execution.
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// MainClass: The name of the driver's main class. The jar file that
	// contains the class must be in the default CLASSPATH or specified in
	// jar_file_uris.
	MainClass string `json:"mainClass,omitempty"`

	// MainJarFileUri: The HCFS URI of the jar file that contains the main
	// class.
	MainJarFileUri string `json:"mainJarFileUri,omitempty"`

	// Properties: Optional. A mapping of property names to values, used to
	// configure Spark. Properties that conflict with values set by the
	// Cloud Dataproc API may be overwritten. Can include properties set in
	// /etc/spark/conf/spark-defaults.conf and classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ArchiveUris") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArchiveUris") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SparkJob) MarshalJSON() ([]byte, error) {
	type NoMethod SparkJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SparkSqlJob: A Cloud Dataproc job for running Apache Spark SQL
// (http://spark.apache.org/sql/) queries.
type SparkSqlJob struct {
	// JarFileUris: Optional. HCFS URIs of jar files to be added to the
	// Spark CLASSPATH.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfig: Optional. The runtime log config for job execution.
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Properties: Optional. A mapping of property names to values, used to
	// configure Spark SQL's SparkConf. Properties that conflict with values
	// set by the Cloud Dataproc API may be overwritten.
	Properties map[string]string `json:"properties,omitempty"`

	// QueryFileUri: The HCFS URI of the script that contains SQL queries.
	QueryFileUri string `json:"queryFileUri,omitempty"`

	// QueryList: A list of queries.
	QueryList *QueryList `json:"queryList,omitempty"`

	// ScriptVariables: Optional. Mapping of query variable names to values
	// (equivalent to the Spark SQL command: SET name="value";).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// ForceSendFields is a list of field names (e.g. "JarFileUris") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "JarFileUris") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SparkSqlJob) MarshalJSON() ([]byte, error) {
	type NoMethod SparkSqlJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The Status type defines a logical error model that is
// suitable for different programming environments, including REST APIs
// and RPC APIs. It is used by gRPC (https://github.com/grpc). The error
// model is designed to be:
// Simple to use and understand for most users
// Flexible enough to meet unexpected needsOverviewThe Status message
// contains three pieces of data: error code, error message, and error
// details. The error code should be an enum value of google.rpc.Code,
// but it may accept additional error codes if needed. The error message
// should be a developer-facing English message that helps developers
// understand and resolve the error. If a localized user-facing error
// message is needed, put the localized message in the error details or
// localize it in the client. The optional error details may contain
// arbitrary information about the error. There is a predefined set of
// error detail types in the package google.rpc that can be used for
// common error conditions.Language mappingThe Status message is the
// logical representation of the error model, but it is not necessarily
// the actual wire format. When the Status message is exposed in
// different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions in Java, but more likely mapped to some error codes in
// C.Other usesThe error model and the Status message can be used in a
// variety of environments, either with or without APIs, to provide a
// consistent developer experience across different environments.Example
// uses of this error model include:
// Partial errors. If a service needs to return partial errors to the
// client, it may embed the Status in the normal response to indicate
// the partial errors.
// Workflow errors. A typical workflow has multiple steps. Each step may
// have a Status message for error reporting.
// Batch operations. If a client uses batch request and batch response,
// the Status message should be used directly inside batch response, one
// for each error sub-response.
// Asynchronous operations. If an API call embeds asynchronous operation
// results in its response, the status of those operations should be
// represented directly using the Status message.
// Logging. If some API errors are stored in logs, the message Status
// could be used directly after any stripping needed for
// security/privacy reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There is a
	// common set of message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any user-facing error message should be localized and sent
	// in the google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SubmitJobRequest: A request to submit a job.
type SubmitJobRequest struct {
	// Job: Required. The job resource.
	Job *Job `json:"job,omitempty"`

	// RequestId: Optional. A unique id used to identify the request. If the
	// server receives two SubmitJobRequest requests with the same id, then
	// the second request will be ignored and the first Job created and
	// stored in the backend is returned.It is recommended to always set
	// this value to a UUID
	// (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id
	// must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). The maximum length is 40 characters.
	RequestId string `json:"requestId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Job") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Job") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SubmitJobRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SubmitJobRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TemplateParameter: A configurable parameter that replaces one or more
// fields in the template. Parameterizable fields: - Labels - File uris
// - Job properties - Job arguments - Script variables - Main class (in
// HadoopJob and SparkJob) - Zone (in ClusterSelector)
type TemplateParameter struct {
	// Description: Optional. Brief description of the parameter. Must not
	// exceed 1024 characters.
	Description string `json:"description,omitempty"`

	// Fields: Required. Paths to all fields that the parameter replaces. A
	// field is allowed to appear in at most one parameter's list of field
	// paths.A field path is similar in syntax to a
	// google.protobuf.FieldMask. For example, a field path that references
	// the zone field of a workflow template's cluster selector would be
	// specified as placement.clusterSelector.zone.Also, field paths can
	// reference fields using the following syntax:
	// Values in maps can be referenced by
	// key:
	// labels'key'
	// placement.clusterSelector.clusterLabels'key'
	// placemen
	// t.managedCluster.labels'key'
	// placement.clusterSelector.clusterLabels'k
	// ey'
	// jobs'step-id'.labels'key'
	// Jobs in the jobs list can be referenced by
	// step-id:
	// jobs'step-id'.hadoopJob.mainJarFileUri
	// jobs'step-id'.hiveJob.
	// queryFileUri
	// jobs'step-id'.pySparkJob.mainPythonFileUri
	// jobs'step-id'.
	// hadoopJob.jarFileUris0
	// jobs'step-id'.hadoopJob.archiveUris0
	// jobs'step-
	// id'.hadoopJob.fileUris0
	// jobs'step-id'.pySparkJob.pythonFileUris0
	// Items
	//  in repeated fields can be referenced by a zero-based
	// index:
	// jobs'step-id'.sparkJob.args0
	// Other
	// examples:
	// jobs'step-id'.hadoopJob.properties'key'
	// jobs'step-id'.hadoop
	// Job.args0
	// jobs'step-id'.hiveJob.scriptVariables'key'
	// jobs'step-id'.had
	// oopJob.mainJarFileUri
	// placement.clusterSelector.zoneIt may not be possible to parameterize
	// maps and repeated fields in their entirety since only individual map
	// values and individual items in repeated fields can be referenced. For
	// example, the following field paths are
	// invalid:
	// placement.clusterSelector.clusterLabels
	// jobs'step-id'.sparkJo
	// b.args
	Fields []string `json:"fields,omitempty"`

	// Name: Required. Parameter name. The parameter name is used as the
	// key, and paired with the parameter value, which are passed to the
	// template when the template is instantiated. The name must contain
	// only capital letters (A-Z), numbers (0-9), and underscores (_), and
	// must not start with a number. The maximum length is 40 characters.
	Name string `json:"name,omitempty"`

	// Validation: Optional. Validation rules to be applied to this
	// parameter's value.
	Validation *ParameterValidation `json:"validation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TemplateParameter) MarshalJSON() ([]byte, error) {
	type NoMethod TemplateParameter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsRequest: Request message for TestIamPermissions
// method.
type TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the resource.
	// Permissions with wildcards (such as '*' or 'storage.*') are not
	// allowed. For more information see IAM Overview
	// (https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod TestIamPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsResponse: Response message for TestIamPermissions
// method.
type TestIamPermissionsResponse struct {
	// Permissions: A subset of TestPermissionsRequest.permissions that the
	// caller is allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TestIamPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ValueValidation: Validation based on a list of allowed values.
type ValueValidation struct {
	// Values: Required. List of allowed values for the parameter.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ValueValidation) MarshalJSON() ([]byte, error) {
	type NoMethod ValueValidation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkflowGraph: The workflow graph.
type WorkflowGraph struct {
	// Nodes: Output only. The workflow nodes.
	Nodes []*WorkflowNode `json:"nodes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Nodes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Nodes") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WorkflowGraph) MarshalJSON() ([]byte, error) {
	type NoMethod WorkflowGraph
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkflowMetadata: A Cloud Dataproc workflow template resource.
type WorkflowMetadata struct {
	// ClusterName: Output only. The name of the managed cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// CreateCluster: Output only. The create cluster operation metadata.
	CreateCluster *ClusterOperation `json:"createCluster,omitempty"`

	// DeleteCluster: Output only. The delete cluster operation metadata.
	DeleteCluster *ClusterOperation `json:"deleteCluster,omitempty"`

	// EndTime: Output only. Workflow end time.
	EndTime string `json:"endTime,omitempty"`

	// Graph: Output only. The workflow graph.
	Graph *WorkflowGraph `json:"graph,omitempty"`

	// Parameters: Map from parameter names to values that were used for
	// those parameters.
	Parameters map[string]string `json:"parameters,omitempty"`

	// StartTime: Output only. Workflow start time.
	StartTime string `json:"startTime,omitempty"`

	// State: Output only. The workflow state.
	//
	// Possible values:
	//   "UNKNOWN" - Unused.
	//   "PENDING" - The operation has been created.
	//   "RUNNING" - The operation is running.
	//   "DONE" - The operation is done; either cancelled or completed.
	State string `json:"state,omitempty"`

	// Template: Output only. The "resource name" of the template.
	Template string `json:"template,omitempty"`

	// Version: Output only. The version of template at the time of workflow
	// instantiation.
	Version int64 `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClusterName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusterName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WorkflowMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod WorkflowMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkflowNode: The workflow node.
type WorkflowNode struct {
	// Error: Output only. The error detail.
	Error string `json:"error,omitempty"`

	// JobId: Output only. The job id; populated after the node enters
	// RUNNING state.
	JobId string `json:"jobId,omitempty"`

	// PrerequisiteStepIds: Output only. Node's prerequisite nodes.
	PrerequisiteStepIds []string `json:"prerequisiteStepIds,omitempty"`

	// State: Output only. The node state.
	//
	// Possible values:
	//   "NODE_STATE_UNSPECIFIED" - State is unspecified.
	//   "BLOCKED" - The node is awaiting prerequisite node to finish.
	//   "RUNNABLE" - The node is runnable but not running.
	//   "RUNNING" - The node is running.
	//   "COMPLETED" - The node completed successfully.
	//   "FAILED" - The node failed. A node can be marked FAILED because its
	// ancestor or peer failed.
	State string `json:"state,omitempty"`

	// StepId: Output only. The name of the node.
	StepId string `json:"stepId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WorkflowNode) MarshalJSON() ([]byte, error) {
	type NoMethod WorkflowNode
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkflowTemplate: A Cloud Dataproc workflow template resource.
type WorkflowTemplate struct {
	// CreateTime: Output only. The time template was created.
	CreateTime string `json:"createTime,omitempty"`

	// Id: Required. The template id.The id must contain only letters (a-z,
	// A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin
	// or end with underscore or hyphen. Must consist of between 3 and 50
	// characters.
	Id string `json:"id,omitempty"`

	// Jobs: Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs []*OrderedJob `json:"jobs,omitempty"`

	// Labels: Optional. The labels to associate with this template. These
	// labels will be propagated to all jobs and clusters created by the
	// workflow instance.Label keys must contain 1 to 63 characters, and
	// must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label
	// values may be empty, but, if present, must contain 1 to 63
	// characters, and must conform to RFC 1035
	// (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be
	// associated with a template.
	Labels map[string]string `json:"labels,omitempty"`

	// Name: Output only. The "resource name" of the template, as described
	// in https://cloud.google.com/apis/design/resource_names of the form
	// projects/{project_id}/regions/{region}/workflowTemplates/{template_id}
	Name string `json:"name,omitempty"`

	// Parameters: Optional. Template parameters whose values are
	// substituted into the template. Values for parameters must be provided
	// when the template is instantiated.
	Parameters []*TemplateParameter `json:"parameters,omitempty"`

	// Placement: Required. WorkflowTemplate scheduling information.
	Placement *WorkflowTemplatePlacement `json:"placement,omitempty"`

	// UpdateTime: Output only. The time template was last updated.
	UpdateTime string `json:"updateTime,omitempty"`

	// Version: Optional. Used to perform a consistent
	// read-modify-write.This field should be left blank for a
	// CreateWorkflowTemplate request. It is required for an
	// UpdateWorkflowTemplate request, and must match the current server
	// version. A typical update template flow would fetch the current
	// template with a GetWorkflowTemplate request, which will return the
	// current template with the version field filled in with the current
	// server version. The user updates other fields in the template, then
	// returns it as part of the UpdateWorkflowTemplate request.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WorkflowTemplate) MarshalJSON() ([]byte, error) {
	type NoMethod WorkflowTemplate
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkflowTemplatePlacement: Specifies workflow execution target.Either
// managed_cluster or cluster_selector is required.
type WorkflowTemplatePlacement struct {
	// ClusterSelector: Optional. A selector that chooses target cluster for
	// jobs based on metadata.The selector is evaluated at the time each job
	// is submitted.
	ClusterSelector *ClusterSelector `json:"clusterSelector,omitempty"`

	// ManagedCluster: Optional. A cluster that is managed by the workflow.
	ManagedCluster *ManagedCluster `json:"managedCluster,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClusterSelector") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClusterSelector") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *WorkflowTemplatePlacement) MarshalJSON() ([]byte, error) {
	type NoMethod WorkflowTemplatePlacement
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// YarnApplication: A YARN application created by a job. Application
// information is a subset of
// <code>org.apache.hadoop.yarn.proto.YarnProtos.ApplicationReportProto</
// code>.Beta Feature: This report is available for testing purposes
// only. It may be changed before final release.
type YarnApplication struct {
	// Name: Required. The application name.
	Name string `json:"name,omitempty"`

	// Progress: Required. The numerical progress of the application, from 1
	// to 100.
	Progress float64 `json:"progress,omitempty"`

	// State: Required. The application state.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Status is unspecified.
	//   "NEW" - Status is NEW.
	//   "NEW_SAVING" - Status is NEW_SAVING.
	//   "SUBMITTED" - Status is SUBMITTED.
	//   "ACCEPTED" - Status is ACCEPTED.
	//   "RUNNING" - Status is RUNNING.
	//   "FINISHED" - Status is FINISHED.
	//   "FAILED" - Status is FAILED.
	//   "KILLED" - Status is KILLED.
	State string `json:"state,omitempty"`

	// TrackingUrl: Optional. The HTTP URL of the ApplicationMaster,
	// HistoryServer, or TimelineServer that provides application-specific
	// information. The URL uses the internal hostname, and requires a proxy
	// server for resolution and, possibly, access.
	TrackingUrl string `json:"trackingUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *YarnApplication) MarshalJSON() ([]byte, error) {
	type NoMethod YarnApplication
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *YarnApplication) UnmarshalJSON(data []byte) error {
	type NoMethod YarnApplication
	var s1 struct {
		Progress gensupport.JSONFloat64 `json:"progress"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Progress = float64(s1.Progress)
	return nil
}

// method id "dataproc.projects.locations.workflowTemplates.create":

type ProjectsLocationsWorkflowTemplatesCreateCall struct {
	s                *Service
	parent           string
	workflowtemplate *WorkflowTemplate
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Create: Creates new workflow template.
func (r *ProjectsLocationsWorkflowTemplatesService) Create(parent string, workflowtemplate *WorkflowTemplate) *ProjectsLocationsWorkflowTemplatesCreateCall {
	c := &ProjectsLocationsWorkflowTemplatesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.workflowtemplate = workflowtemplate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesCreateCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesCreateCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.workflowtemplate)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/workflowTemplates")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.create" call.
// Exactly one of *WorkflowTemplate or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WorkflowTemplate.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowTemplatesCreateCall) Do(opts ...googleapi.CallOption) (*WorkflowTemplate, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WorkflowTemplate{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates new workflow template.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.locations.workflowTemplates.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The \"resource name\" of the region, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/workflowTemplates",
	//   "request": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "response": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.delete":

type ProjectsLocationsWorkflowTemplatesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a workflow template. It does not cancel in-progress
// workflows.
func (r *ProjectsLocationsWorkflowTemplatesService) Delete(name string) *ProjectsLocationsWorkflowTemplatesDeleteCall {
	c := &ProjectsLocationsWorkflowTemplatesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Version sets the optional parameter "version": The version of
// workflow template to delete. If specified, will only delete the
// template if the current server version matches specified version.
func (c *ProjectsLocationsWorkflowTemplatesDeleteCall) Version(version int64) *ProjectsLocationsWorkflowTemplatesDeleteCall {
	c.urlParams_.Set("version", fmt.Sprint(version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesDeleteCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesDeleteCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsWorkflowTemplatesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a workflow template. It does not cancel in-progress workflows.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates/{workflowTemplatesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.locations.workflowTemplates.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The \"resource name\" of the workflow template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Optional. The version of workflow template to delete. If specified, will only delete the template if the current server version matches specified version.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.get":

type ProjectsLocationsWorkflowTemplatesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the latest workflow template.Can retrieve previously
// instantiated template by specifying optional version parameter.
func (r *ProjectsLocationsWorkflowTemplatesService) Get(name string) *ProjectsLocationsWorkflowTemplatesGetCall {
	c := &ProjectsLocationsWorkflowTemplatesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Version sets the optional parameter "version": The version of
// workflow template to retrieve. Only previously instatiated versions
// can be retrieved.If unspecified, retrieves the current version.
func (c *ProjectsLocationsWorkflowTemplatesGetCall) Version(version int64) *ProjectsLocationsWorkflowTemplatesGetCall {
	c.urlParams_.Set("version", fmt.Sprint(version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsWorkflowTemplatesGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsWorkflowTemplatesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesGetCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.get" call.
// Exactly one of *WorkflowTemplate or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WorkflowTemplate.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowTemplatesGetCall) Do(opts ...googleapi.CallOption) (*WorkflowTemplate, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WorkflowTemplate{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the latest workflow template.Can retrieve previously instantiated template by specifying optional version parameter.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates/{workflowTemplatesId}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.locations.workflowTemplates.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The \"resource name\" of the workflow template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Optional. The version of workflow template to retrieve. Only previously instatiated versions can be retrieved.If unspecified, retrieves the current version.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.getIamPolicy":

type ProjectsLocationsWorkflowTemplatesGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Gets the access control policy for a resource. Returns
// an empty policy if the resource exists and does not have a policy
// set.
func (r *ProjectsLocationsWorkflowTemplatesService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall {
	c := &ProjectsLocationsWorkflowTemplatesGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsWorkflowTemplatesGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource. Returns an empty policy if the resource exists and does not have a policy set.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates/{workflowTemplatesId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.locations.workflowTemplates.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.instantiate":

type ProjectsLocationsWorkflowTemplatesInstantiateCall struct {
	s                                  *Service
	name                               string
	instantiateworkflowtemplaterequest *InstantiateWorkflowTemplateRequest
	urlParams_                         gensupport.URLParams
	ctx_                               context.Context
	header_                            http.Header
}

// Instantiate: Instantiates a template and begins execution.The
// returned Operation can be used to track execution of workflow by
// polling operations.get. The Operation will complete when entire
// workflow is finished.The running workflow can be aborted via
// operations.cancel. This will cause any inflight jobs to be cancelled
// and workflow-owned clusters to be deleted.The Operation.metadata will
// be WorkflowMetadata.On successful completion, Operation.response will
// be Empty.
func (r *ProjectsLocationsWorkflowTemplatesService) Instantiate(name string, instantiateworkflowtemplaterequest *InstantiateWorkflowTemplateRequest) *ProjectsLocationsWorkflowTemplatesInstantiateCall {
	c := &ProjectsLocationsWorkflowTemplatesInstantiateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.instantiateworkflowtemplaterequest = instantiateworkflowtemplaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesInstantiateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesInstantiateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesInstantiateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.instantiateworkflowtemplaterequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:instantiate")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.instantiate" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Instantiates a template and begins execution.The returned Operation can be used to track execution of workflow by polling operations.get. The Operation will complete when entire workflow is finished.The running workflow can be aborted via operations.cancel. This will cause any inflight jobs to be cancelled and workflow-owned clusters to be deleted.The Operation.metadata will be WorkflowMetadata.On successful completion, Operation.response will be Empty.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates/{workflowTemplatesId}:instantiate",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.locations.workflowTemplates.instantiate",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The \"resource name\" of the workflow template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:instantiate",
	//   "request": {
	//     "$ref": "InstantiateWorkflowTemplateRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.instantiateInline":

type ProjectsLocationsWorkflowTemplatesInstantiateInlineCall struct {
	s                *Service
	parent           string
	workflowtemplate *WorkflowTemplate
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// InstantiateInline: Instantiates a template and begins execution.This
// method is equivalent to executing the sequence
// CreateWorkflowTemplate, InstantiateWorkflowTemplate,
// DeleteWorkflowTemplate.The returned Operation can be used to track
// execution of workflow by polling operations.get. The Operation will
// complete when entire workflow is finished.The running workflow can be
// aborted via operations.cancel. This will cause any inflight jobs to
// be cancelled and workflow-owned clusters to be deleted.The
// Operation.metadata will be WorkflowMetadata.On successful completion,
// Operation.response will be Empty.
func (r *ProjectsLocationsWorkflowTemplatesService) InstantiateInline(parent string, workflowtemplate *WorkflowTemplate) *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall {
	c := &ProjectsLocationsWorkflowTemplatesInstantiateInlineCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.workflowtemplate = workflowtemplate
	return c
}

// RequestId sets the optional parameter "requestId": A tag that
// prevents multiple concurrent workflow instances with the same tag
// from running. This mitigates risk of concurrent instances started due
// to retries.It is recommended to always set this value to a UUID
// (https://en.wikipedia.org/wiki/Universally_unique_identifier).The tag
// must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
// and hyphens (-). The maximum length is 40 characters.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall) RequestId(requestId string) *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall {
	c.urlParams_.Set("requestId", requestId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.workflowtemplate)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/workflowTemplates:instantiateInline")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.instantiateInline" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowTemplatesInstantiateInlineCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Instantiates a template and begins execution.This method is equivalent to executing the sequence CreateWorkflowTemplate, InstantiateWorkflowTemplate, DeleteWorkflowTemplate.The returned Operation can be used to track execution of workflow by polling operations.get. The Operation will complete when entire workflow is finished.The running workflow can be aborted via operations.cancel. This will cause any inflight jobs to be cancelled and workflow-owned clusters to be deleted.The Operation.metadata will be WorkflowMetadata.On successful completion, Operation.response will be Empty.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates:instantiateInline",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.locations.workflowTemplates.instantiateInline",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The \"resource name\" of the workflow template region, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "Optional. A tag that prevents multiple concurrent workflow instances with the same tag from running. This mitigates risk of concurrent instances started due to retries.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The tag must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/workflowTemplates:instantiateInline",
	//   "request": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.list":

type ProjectsLocationsWorkflowTemplatesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists workflows that match the specified filter in the request.
func (r *ProjectsLocationsWorkflowTemplatesService) List(parent string) *ProjectsLocationsWorkflowTemplatesListCall {
	c := &ProjectsLocationsWorkflowTemplatesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in each response.
func (c *ProjectsLocationsWorkflowTemplatesListCall) PageSize(pageSize int64) *ProjectsLocationsWorkflowTemplatesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The page token,
// returned by a previous call, to request the next page of results.
func (c *ProjectsLocationsWorkflowTemplatesListCall) PageToken(pageToken string) *ProjectsLocationsWorkflowTemplatesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesListCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsWorkflowTemplatesListCall) IfNoneMatch(entityTag string) *ProjectsLocationsWorkflowTemplatesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesListCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/workflowTemplates")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.list" call.
// Exactly one of *ListWorkflowTemplatesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListWorkflowTemplatesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowTemplatesListCall) Do(opts ...googleapi.CallOption) (*ListWorkflowTemplatesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListWorkflowTemplatesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists workflows that match the specified filter in the request.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.locations.workflowTemplates.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return in each response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The page token, returned by a previous call, to request the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The \"resource name\" of the region, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/workflowTemplates",
	//   "response": {
	//     "$ref": "ListWorkflowTemplatesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsWorkflowTemplatesListCall) Pages(ctx context.Context, f func(*ListWorkflowTemplatesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dataproc.projects.locations.workflowTemplates.setIamPolicy":

type ProjectsLocationsWorkflowTemplatesSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any existing policy.
func (r *ProjectsLocationsWorkflowTemplatesService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall {
	c := &ProjectsLocationsWorkflowTemplatesSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsWorkflowTemplatesSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any existing policy.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates/{workflowTemplatesId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.locations.workflowTemplates.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.testIamPermissions":

type ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource. If the resource does not exist, this will return
// an empty set of permissions, not a NOT_FOUND error.Note: This
// operation is designed to be used for building permission-aware UIs
// and command-line tools, not for authorization checking. This
// operation may "fail open" without warning.
func (r *ProjectsLocationsWorkflowTemplatesService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall {
	c := &ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowTemplatesTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource. If the resource does not exist, this will return an empty set of permissions, not a NOT_FOUND error.Note: This operation is designed to be used for building permission-aware UIs and command-line tools, not for authorization checking. This operation may \"fail open\" without warning.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates/{workflowTemplatesId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.locations.workflowTemplates.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.locations.workflowTemplates.update":

type ProjectsLocationsWorkflowTemplatesUpdateCall struct {
	s                *Service
	name             string
	workflowtemplate *WorkflowTemplate
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Update: Updates (replaces) workflow template. The updated template
// must contain version that matches the current server version.
func (r *ProjectsLocationsWorkflowTemplatesService) Update(name string, workflowtemplate *WorkflowTemplate) *ProjectsLocationsWorkflowTemplatesUpdateCall {
	c := &ProjectsLocationsWorkflowTemplatesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.workflowtemplate = workflowtemplate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsWorkflowTemplatesUpdateCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowTemplatesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsWorkflowTemplatesUpdateCall) Context(ctx context.Context) *ProjectsLocationsWorkflowTemplatesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsWorkflowTemplatesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowTemplatesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.workflowtemplate)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.locations.workflowTemplates.update" call.
// Exactly one of *WorkflowTemplate or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WorkflowTemplate.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowTemplatesUpdateCall) Do(opts ...googleapi.CallOption) (*WorkflowTemplate, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WorkflowTemplate{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates (replaces) workflow template. The updated template must contain version that matches the current server version.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/workflowTemplates/{workflowTemplatesId}",
	//   "httpMethod": "PUT",
	//   "id": "dataproc.projects.locations.workflowTemplates.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Output only. The \"resource name\" of the template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "response": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.create":

type ProjectsRegionsClustersCreateCall struct {
	s          *Service
	projectId  string
	region     string
	cluster    *Cluster
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a cluster in a project.
func (r *ProjectsRegionsClustersService) Create(projectId string, region string, cluster *Cluster) *ProjectsRegionsClustersCreateCall {
	c := &ProjectsRegionsClustersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.cluster = cluster
	return c
}

// RequestId sets the optional parameter "requestId": A unique id used
// to identify the request. If the server receives two
// CreateClusterRequest requests with the same id, then the second
// request will be ignored and the first google.longrunning.Operation
// created and stored in the backend is returned.It is recommended to
// always set this value to a UUID
// (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id
// must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
// and hyphens (-). The maximum length is 40 characters.
func (c *ProjectsRegionsClustersCreateCall) RequestId(requestId string) *ProjectsRegionsClustersCreateCall {
	c.urlParams_.Set("requestId", requestId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersCreateCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersCreateCall) Context(ctx context.Context) *ProjectsRegionsClustersCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cluster)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/clusters")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsClustersCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a cluster in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/clusters",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.clusters.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "region"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "Optional. A unique id used to identify the request. If the server receives two CreateClusterRequest requests with the same id, then the second request will be ignored and the first google.longrunning.Operation created and stored in the backend is returned.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/clusters",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.delete":

type ProjectsRegionsClustersDeleteCall struct {
	s           *Service
	projectId   string
	region      string
	clusterName string
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Delete: Deletes a cluster in a project.
func (r *ProjectsRegionsClustersService) Delete(projectId string, region string, clusterName string) *ProjectsRegionsClustersDeleteCall {
	c := &ProjectsRegionsClustersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.clusterName = clusterName
	return c
}

// ClusterUuid sets the optional parameter "clusterUuid": Specifying the
// cluster_uuid means the RPC should fail (with error NOT_FOUND) if
// cluster with specified UUID does not exist.
func (c *ProjectsRegionsClustersDeleteCall) ClusterUuid(clusterUuid string) *ProjectsRegionsClustersDeleteCall {
	c.urlParams_.Set("clusterUuid", clusterUuid)
	return c
}

// RequestId sets the optional parameter "requestId": A unique id used
// to identify the request. If the server receives two
// DeleteClusterRequest requests with the same id, then the second
// request will be ignored and the first google.longrunning.Operation
// created and stored in the backend is returned.It is recommended to
// always set this value to a UUID
// (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id
// must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
// and hyphens (-). The maximum length is 40 characters.
func (c *ProjectsRegionsClustersDeleteCall) RequestId(requestId string) *ProjectsRegionsClustersDeleteCall {
	c.urlParams_.Set("requestId", requestId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersDeleteCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersDeleteCall) Context(ctx context.Context) *ProjectsRegionsClustersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"region":      c.region,
		"clusterName": c.clusterName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsClustersDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a cluster in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.regions.clusters.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "Required. The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clusterUuid": {
	//       "description": "Optional. Specifying the cluster_uuid means the RPC should fail (with error NOT_FOUND) if cluster with specified UUID does not exist.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "Optional. A unique id used to identify the request. If the server receives two DeleteClusterRequest requests with the same id, then the second request will be ignored and the first google.longrunning.Operation created and stored in the backend is returned.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.diagnose":

type ProjectsRegionsClustersDiagnoseCall struct {
	s                      *Service
	projectId              string
	region                 string
	clusterName            string
	diagnoseclusterrequest *DiagnoseClusterRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Diagnose: Gets cluster diagnostic information. After the operation
// completes, the Operation.response field contains
// DiagnoseClusterOutputLocation.
func (r *ProjectsRegionsClustersService) Diagnose(projectId string, region string, clusterName string, diagnoseclusterrequest *DiagnoseClusterRequest) *ProjectsRegionsClustersDiagnoseCall {
	c := &ProjectsRegionsClustersDiagnoseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.clusterName = clusterName
	c.diagnoseclusterrequest = diagnoseclusterrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersDiagnoseCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersDiagnoseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersDiagnoseCall) Context(ctx context.Context) *ProjectsRegionsClustersDiagnoseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersDiagnoseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersDiagnoseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.diagnoseclusterrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}:diagnose")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"region":      c.region,
		"clusterName": c.clusterName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.diagnose" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsClustersDiagnoseCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets cluster diagnostic information. After the operation completes, the Operation.response field contains DiagnoseClusterOutputLocation.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}:diagnose",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.clusters.diagnose",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "Required. The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}:diagnose",
	//   "request": {
	//     "$ref": "DiagnoseClusterRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.get":

type ProjectsRegionsClustersGetCall struct {
	s            *Service
	projectId    string
	region       string
	clusterName  string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the resource representation for a cluster in a project.
func (r *ProjectsRegionsClustersService) Get(projectId string, region string, clusterName string) *ProjectsRegionsClustersGetCall {
	c := &ProjectsRegionsClustersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.clusterName = clusterName
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersGetCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsClustersGetCall) IfNoneMatch(entityTag string) *ProjectsRegionsClustersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersGetCall) Context(ctx context.Context) *ProjectsRegionsClustersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"region":      c.region,
		"clusterName": c.clusterName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.get" call.
// Exactly one of *Cluster or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Cluster.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsClustersGetCall) Do(opts ...googleapi.CallOption) (*Cluster, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Cluster{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the resource representation for a cluster in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.clusters.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "Required. The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}",
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.getIamPolicy":

type ProjectsRegionsClustersGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Gets the access control policy for a resource. Returns
// an empty policy if the resource exists and does not have a policy
// set.
func (r *ProjectsRegionsClustersService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsRegionsClustersGetIamPolicyCall {
	c := &ProjectsRegionsClustersGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersGetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsClustersGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsClustersGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource. Returns an empty policy if the resource exists and does not have a policy set.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/clusters/{clustersId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.clusters.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/clusters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.list":

type ProjectsRegionsClustersListCall struct {
	s            *Service
	projectId    string
	region       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all regions/{region}/clusters in a project.
func (r *ProjectsRegionsClustersService) List(projectId string, region string) *ProjectsRegionsClustersListCall {
	c := &ProjectsRegionsClustersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	return c
}

// Filter sets the optional parameter "filter": A filter constraining
// the clusters to list. Filters are case-sensitive and have the
// following syntax:field = value AND field = value ...where field is
// one of status.state, clusterName, or labels.[KEY], and [KEY] is a
// label key. value can be * to match all values. status.state can be
// one of the following: ACTIVE, INACTIVE, CREATING, RUNNING, ERROR,
// DELETING, or UPDATING. ACTIVE contains the CREATING, UPDATING, and
// RUNNING states. INACTIVE contains the DELETING and ERROR states.
// clusterName is the name of the cluster provided at creation time.
// Only the logical AND operator is supported; space-separated items are
// treated as having an implicit AND operator.Example
// filter:status.state = ACTIVE AND clusterName = mycluster AND
// labels.env = staging AND labels.starred = *
func (c *ProjectsRegionsClustersListCall) Filter(filter string) *ProjectsRegionsClustersListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard List
// page size.
func (c *ProjectsRegionsClustersListCall) PageSize(pageSize int64) *ProjectsRegionsClustersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard List
// page token.
func (c *ProjectsRegionsClustersListCall) PageToken(pageToken string) *ProjectsRegionsClustersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersListCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsClustersListCall) IfNoneMatch(entityTag string) *ProjectsRegionsClustersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersListCall) Context(ctx context.Context) *ProjectsRegionsClustersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/clusters")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.list" call.
// Exactly one of *ListClustersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListClustersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsClustersListCall) Do(opts ...googleapi.CallOption) (*ListClustersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListClustersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all regions/{region}/clusters in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/clusters",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.clusters.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "region"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. A filter constraining the clusters to list. Filters are case-sensitive and have the following syntax:field = value AND field = value ...where field is one of status.state, clusterName, or labels.[KEY], and [KEY] is a label key. value can be * to match all values. status.state can be one of the following: ACTIVE, INACTIVE, CREATING, RUNNING, ERROR, DELETING, or UPDATING. ACTIVE contains the CREATING, UPDATING, and RUNNING states. INACTIVE contains the DELETING and ERROR states. clusterName is the name of the cluster provided at creation time. Only the logical AND operator is supported; space-separated items are treated as having an implicit AND operator.Example filter:status.state = ACTIVE AND clusterName = mycluster AND labels.env = staging AND labels.starred = *",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The standard List page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The standard List page token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/clusters",
	//   "response": {
	//     "$ref": "ListClustersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsRegionsClustersListCall) Pages(ctx context.Context, f func(*ListClustersResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dataproc.projects.regions.clusters.patch":

type ProjectsRegionsClustersPatchCall struct {
	s           *Service
	projectId   string
	region      string
	clusterName string
	cluster     *Cluster
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Patch: Updates a cluster in a project.
func (r *ProjectsRegionsClustersService) Patch(projectId string, region string, clusterName string, cluster *Cluster) *ProjectsRegionsClustersPatchCall {
	c := &ProjectsRegionsClustersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.clusterName = clusterName
	c.cluster = cluster
	return c
}

// GracefulDecommissionTimeout sets the optional parameter
// "gracefulDecommissionTimeout": Timeout for graceful YARN
// decomissioning. Graceful decommissioning allows removing nodes from
// the cluster without interrupting jobs in progress. Timeout specifies
// how long to wait for jobs in progress to finish before forcefully
// removing nodes (and potentially interrupting jobs). Default timeout
// is 0 (for forceful decommission), and the maximum allowed timeout is
// 1 day.Only supported on Dataproc image versions 1.2 and higher.
func (c *ProjectsRegionsClustersPatchCall) GracefulDecommissionTimeout(gracefulDecommissionTimeout string) *ProjectsRegionsClustersPatchCall {
	c.urlParams_.Set("gracefulDecommissionTimeout", gracefulDecommissionTimeout)
	return c
}

// RequestId sets the optional parameter "requestId": A unique id used
// to identify the request. If the server receives two
// UpdateClusterRequest requests with the same id, then the second
// request will be ignored and the first google.longrunning.Operation
// created and stored in the backend is returned.It is recommended to
// always set this value to a UUID
// (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id
// must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
// and hyphens (-). The maximum length is 40 characters.
func (c *ProjectsRegionsClustersPatchCall) RequestId(requestId string) *ProjectsRegionsClustersPatchCall {
	c.urlParams_.Set("requestId", requestId)
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required.
// Specifies the path, relative to Cluster, of the field to update. For
// example, to change the number of workers in a cluster to 5, the
// update_mask parameter would be specified as
// config.worker_config.num_instances, and the PATCH request body would
// specify the new value, as follows:
// {
//   "config":{
//     "workerConfig":{
//       "numInstances":"5"
//     }
//   }
// }
// Similarly, to change the number of preemptible workers in a cluster
// to 5, the update_mask parameter would be
// config.secondary_worker_config.num_instances, and the PATCH request
// body would be set as follows:
// {
//   "config":{
//     "secondaryWorkerConfig":{
//       "numInstances":"5"
//     }
//   }
// }
// <strong>Note:</strong> Currently, only the following fields can be
// updated:<table>  <tbody>  <tr>  <td><strong>Mask</strong></td>
// <td><strong>Purpose</strong></td>  </tr>  <tr>
// <td><strong><em>labels</em></strong></td>  <td>Update labels</td>
// </tr>  <tr>
// <td><strong><em>config.worker_config.num_instances</em></strong></td>
//  <td>Resize primary worker group</td>  </tr>  <tr>
// <td><strong><em>config.secondary_worker_config.num_instances</em></str
// ong></td>  <td>Resize secondary worker group</td>  </tr>  </tbody>
// </table>
func (c *ProjectsRegionsClustersPatchCall) UpdateMask(updateMask string) *ProjectsRegionsClustersPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersPatchCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersPatchCall) Context(ctx context.Context) *ProjectsRegionsClustersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cluster)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"region":      c.region,
		"clusterName": c.clusterName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsClustersPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a cluster in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}",
	//   "httpMethod": "PATCH",
	//   "id": "dataproc.projects.regions.clusters.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "Required. The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "gracefulDecommissionTimeout": {
	//       "description": "Optional. Timeout for graceful YARN decomissioning. Graceful decommissioning allows removing nodes from the cluster without interrupting jobs in progress. Timeout specifies how long to wait for jobs in progress to finish before forcefully removing nodes (and potentially interrupting jobs). Default timeout is 0 (for forceful decommission), and the maximum allowed timeout is 1 day.Only supported on Dataproc image versions 1.2 and higher.",
	//       "format": "google-duration",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "Optional. A unique id used to identify the request. If the server receives two UpdateClusterRequest requests with the same id, then the second request will be ignored and the first google.longrunning.Operation created and stored in the backend is returned.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Specifies the path, relative to Cluster, of the field to update. For example, to change the number of workers in a cluster to 5, the update_mask parameter would be specified as config.worker_config.num_instances, and the PATCH request body would specify the new value, as follows:\n{\n  \"config\":{\n    \"workerConfig\":{\n      \"numInstances\":\"5\"\n    }\n  }\n}\nSimilarly, to change the number of preemptible workers in a cluster to 5, the update_mask parameter would be config.secondary_worker_config.num_instances, and the PATCH request body would be set as follows:\n{\n  \"config\":{\n    \"secondaryWorkerConfig\":{\n      \"numInstances\":\"5\"\n    }\n  }\n}\n\u003cstrong\u003eNote:\u003c/strong\u003e Currently, only the following fields can be updated:\u003ctable\u003e  \u003ctbody\u003e  \u003ctr\u003e  \u003ctd\u003e\u003cstrong\u003eMask\u003c/strong\u003e\u003c/td\u003e  \u003ctd\u003e\u003cstrong\u003ePurpose\u003c/strong\u003e\u003c/td\u003e  \u003c/tr\u003e  \u003ctr\u003e  \u003ctd\u003e\u003cstrong\u003e\u003cem\u003elabels\u003c/em\u003e\u003c/strong\u003e\u003c/td\u003e  \u003ctd\u003eUpdate labels\u003c/td\u003e  \u003c/tr\u003e  \u003ctr\u003e  \u003ctd\u003e\u003cstrong\u003e\u003cem\u003econfig.worker_config.num_instances\u003c/em\u003e\u003c/strong\u003e\u003c/td\u003e  \u003ctd\u003eResize primary worker group\u003c/td\u003e  \u003c/tr\u003e  \u003ctr\u003e  \u003ctd\u003e\u003cstrong\u003e\u003cem\u003econfig.secondary_worker_config.num_instances\u003c/em\u003e\u003c/strong\u003e\u003c/td\u003e  \u003ctd\u003eResize secondary worker group\u003c/td\u003e  \u003c/tr\u003e  \u003c/tbody\u003e  \u003c/table\u003e",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/clusters/{clusterName}",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.setIamPolicy":

type ProjectsRegionsClustersSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any existing policy.
func (r *ProjectsRegionsClustersService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsRegionsClustersSetIamPolicyCall {
	c := &ProjectsRegionsClustersSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersSetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsClustersSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsClustersSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any existing policy.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/clusters/{clustersId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.clusters.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/clusters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.clusters.testIamPermissions":

type ProjectsRegionsClustersTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource. If the resource does not exist, this will return
// an empty set of permissions, not a NOT_FOUND error.Note: This
// operation is designed to be used for building permission-aware UIs
// and command-line tools, not for authorization checking. This
// operation may "fail open" without warning.
func (r *ProjectsRegionsClustersService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsRegionsClustersTestIamPermissionsCall {
	c := &ProjectsRegionsClustersTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsClustersTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsRegionsClustersTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsClustersTestIamPermissionsCall) Context(ctx context.Context) *ProjectsRegionsClustersTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsClustersTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsClustersTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.clusters.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsClustersTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource. If the resource does not exist, this will return an empty set of permissions, not a NOT_FOUND error.Note: This operation is designed to be used for building permission-aware UIs and command-line tools, not for authorization checking. This operation may \"fail open\" without warning.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/clusters/{clustersId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.clusters.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/clusters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.cancel":

type ProjectsRegionsJobsCancelCall struct {
	s                *Service
	projectId        string
	region           string
	jobId            string
	canceljobrequest *CancelJobRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Cancel: Starts a job cancellation request. To access the job resource
// after cancellation, call regions/{region}/jobs.list or
// regions/{region}/jobs.get.
func (r *ProjectsRegionsJobsService) Cancel(projectId string, region string, jobId string, canceljobrequest *CancelJobRequest) *ProjectsRegionsJobsCancelCall {
	c := &ProjectsRegionsJobsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.jobId = jobId
	c.canceljobrequest = canceljobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsCancelCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsCancelCall) Context(ctx context.Context) *ProjectsRegionsJobsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.canceljobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/jobs/{jobId}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
		"jobId":     c.jobId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.cancel" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsRegionsJobsCancelCall) Do(opts ...googleapi.CallOption) (*Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts a job cancellation request. To access the job resource after cancellation, call regions/{region}/jobs.list or regions/{region}/jobs.get.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.jobs.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Required. The job ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the job belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}:cancel",
	//   "request": {
	//     "$ref": "CancelJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.delete":

type ProjectsRegionsJobsDeleteCall struct {
	s          *Service
	projectId  string
	region     string
	jobId      string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the job from the project. If the job is active, the
// delete fails, and the response returns FAILED_PRECONDITION.
func (r *ProjectsRegionsJobsService) Delete(projectId string, region string, jobId string) *ProjectsRegionsJobsDeleteCall {
	c := &ProjectsRegionsJobsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.jobId = jobId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsDeleteCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsDeleteCall) Context(ctx context.Context) *ProjectsRegionsJobsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/jobs/{jobId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
		"jobId":     c.jobId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsJobsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the job from the project. If the job is active, the delete fails, and the response returns FAILED_PRECONDITION.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.regions.jobs.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Required. The job ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the job belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.get":

type ProjectsRegionsJobsGetCall struct {
	s            *Service
	projectId    string
	region       string
	jobId        string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the resource representation for a job in a project.
func (r *ProjectsRegionsJobsService) Get(projectId string, region string, jobId string) *ProjectsRegionsJobsGetCall {
	c := &ProjectsRegionsJobsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.jobId = jobId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsGetCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsJobsGetCall) IfNoneMatch(entityTag string) *ProjectsRegionsJobsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsGetCall) Context(ctx context.Context) *ProjectsRegionsJobsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/jobs/{jobId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
		"jobId":     c.jobId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.get" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsRegionsJobsGetCall) Do(opts ...googleapi.CallOption) (*Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the resource representation for a job in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.jobs.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Required. The job ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the job belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.getIamPolicy":

type ProjectsRegionsJobsGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Gets the access control policy for a resource. Returns
// an empty policy if the resource exists and does not have a policy
// set.
func (r *ProjectsRegionsJobsService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsRegionsJobsGetIamPolicyCall {
	c := &ProjectsRegionsJobsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsGetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsJobsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsJobsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource. Returns an empty policy if the resource exists and does not have a policy set.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/jobs/{jobsId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.jobs.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.list":

type ProjectsRegionsJobsListCall struct {
	s            *Service
	projectId    string
	region       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists regions/{region}/jobs in a project.
func (r *ProjectsRegionsJobsService) List(projectId string, region string) *ProjectsRegionsJobsListCall {
	c := &ProjectsRegionsJobsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	return c
}

// ClusterName sets the optional parameter "clusterName": If set, the
// returned jobs list includes only jobs that were submitted to the
// named cluster.
func (c *ProjectsRegionsJobsListCall) ClusterName(clusterName string) *ProjectsRegionsJobsListCall {
	c.urlParams_.Set("clusterName", clusterName)
	return c
}

// Filter sets the optional parameter "filter": A filter constraining
// the jobs to list. Filters are case-sensitive and have the following
// syntax:field = value AND field = value ...where field is status.state
// or labels.[KEY], and [KEY] is a label key. value can be * to match
// all values. status.state can be either ACTIVE or NON_ACTIVE. Only the
// logical AND operator is supported; space-separated items are treated
// as having an implicit AND operator.Example filter:status.state =
// ACTIVE AND labels.env = staging AND labels.starred = *
func (c *ProjectsRegionsJobsListCall) Filter(filter string) *ProjectsRegionsJobsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// JobStateMatcher sets the optional parameter "jobStateMatcher":
// Specifies enumerated categories of jobs to list. (default = match ALL
// jobs).If filter is provided, jobStateMatcher will be ignored.
//
// Possible values:
//   "ALL"
//   "ACTIVE"
//   "NON_ACTIVE"
func (c *ProjectsRegionsJobsListCall) JobStateMatcher(jobStateMatcher string) *ProjectsRegionsJobsListCall {
	c.urlParams_.Set("jobStateMatcher", jobStateMatcher)
	return c
}

// PageSize sets the optional parameter "pageSize": The number of
// results to return in each response.
func (c *ProjectsRegionsJobsListCall) PageSize(pageSize int64) *ProjectsRegionsJobsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The page token,
// returned by a previous call, to request the next page of results.
func (c *ProjectsRegionsJobsListCall) PageToken(pageToken string) *ProjectsRegionsJobsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsListCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsJobsListCall) IfNoneMatch(entityTag string) *ProjectsRegionsJobsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsListCall) Context(ctx context.Context) *ProjectsRegionsJobsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.list" call.
// Exactly one of *ListJobsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListJobsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsJobsListCall) Do(opts ...googleapi.CallOption) (*ListJobsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListJobsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists regions/{region}/jobs in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/jobs",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.jobs.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "region"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "Optional. If set, the returned jobs list includes only jobs that were submitted to the named cluster.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Optional. A filter constraining the jobs to list. Filters are case-sensitive and have the following syntax:field = value AND field = value ...where field is status.state or labels.[KEY], and [KEY] is a label key. value can be * to match all values. status.state can be either ACTIVE or NON_ACTIVE. Only the logical AND operator is supported; space-separated items are treated as having an implicit AND operator.Example filter:status.state = ACTIVE AND labels.env = staging AND labels.starred = *",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobStateMatcher": {
	//       "description": "Optional. Specifies enumerated categories of jobs to list. (default = match ALL jobs).If filter is provided, jobStateMatcher will be ignored.",
	//       "enum": [
	//         "ALL",
	//         "ACTIVE",
	//         "NON_ACTIVE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The number of results to return in each response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The page token, returned by a previous call, to request the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the job belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/jobs",
	//   "response": {
	//     "$ref": "ListJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsRegionsJobsListCall) Pages(ctx context.Context, f func(*ListJobsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dataproc.projects.regions.jobs.patch":

type ProjectsRegionsJobsPatchCall struct {
	s          *Service
	projectId  string
	region     string
	jobId      string
	job        *Job
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates a job in a project.
func (r *ProjectsRegionsJobsService) Patch(projectId string, region string, jobId string, job *Job) *ProjectsRegionsJobsPatchCall {
	c := &ProjectsRegionsJobsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.jobId = jobId
	c.job = job
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required.
// Specifies the path, relative to <code>Job</code>, of the field to
// update. For example, to update the labels of a Job the
// <code>update_mask</code> parameter would be specified as
// <code>labels</code>, and the PATCH request body would specify the new
// value. <strong>Note:</strong> Currently, <code>labels</code> is the
// only field that can be updated.
func (c *ProjectsRegionsJobsPatchCall) UpdateMask(updateMask string) *ProjectsRegionsJobsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsPatchCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsPatchCall) Context(ctx context.Context) *ProjectsRegionsJobsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/jobs/{jobId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
		"jobId":     c.jobId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.patch" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsRegionsJobsPatchCall) Do(opts ...googleapi.CallOption) (*Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a job in a project.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}",
	//   "httpMethod": "PATCH",
	//   "id": "dataproc.projects.regions.jobs.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "region",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Required. The job ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the job belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Specifies the path, relative to \u003ccode\u003eJob\u003c/code\u003e, of the field to update. For example, to update the labels of a Job the \u003ccode\u003eupdate_mask\u003c/code\u003e parameter would be specified as \u003ccode\u003elabels\u003c/code\u003e, and the PATCH request body would specify the new value. \u003cstrong\u003eNote:\u003c/strong\u003e Currently, \u003ccode\u003elabels\u003c/code\u003e is the only field that can be updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/jobs/{jobId}",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.setIamPolicy":

type ProjectsRegionsJobsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any existing policy.
func (r *ProjectsRegionsJobsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsRegionsJobsSetIamPolicyCall {
	c := &ProjectsRegionsJobsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsSetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsJobsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsJobsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any existing policy.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/jobs/{jobsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.jobs.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.submit":

type ProjectsRegionsJobsSubmitCall struct {
	s                *Service
	projectId        string
	region           string
	submitjobrequest *SubmitJobRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Submit: Submits a job to a cluster.
func (r *ProjectsRegionsJobsService) Submit(projectId string, region string, submitjobrequest *SubmitJobRequest) *ProjectsRegionsJobsSubmitCall {
	c := &ProjectsRegionsJobsSubmitCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.region = region
	c.submitjobrequest = submitjobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsSubmitCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsSubmitCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsSubmitCall) Context(ctx context.Context) *ProjectsRegionsJobsSubmitCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsSubmitCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsSubmitCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.submitjobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/regions/{region}/jobs:submit")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"region":    c.region,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.submit" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsRegionsJobsSubmitCall) Do(opts ...googleapi.CallOption) (*Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Submits a job to a cluster.",
	//   "flatPath": "v1/projects/{projectId}/regions/{region}/jobs:submit",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.jobs.submit",
	//   "parameterOrder": [
	//     "projectId",
	//     "region"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Required. The ID of the Google Cloud Platform project that the job belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "region": {
	//       "description": "Required. The Cloud Dataproc region in which to handle the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/regions/{region}/jobs:submit",
	//   "request": {
	//     "$ref": "SubmitJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.jobs.testIamPermissions":

type ProjectsRegionsJobsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource. If the resource does not exist, this will return
// an empty set of permissions, not a NOT_FOUND error.Note: This
// operation is designed to be used for building permission-aware UIs
// and command-line tools, not for authorization checking. This
// operation may "fail open" without warning.
func (r *ProjectsRegionsJobsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsRegionsJobsTestIamPermissionsCall {
	c := &ProjectsRegionsJobsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsJobsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsRegionsJobsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsJobsTestIamPermissionsCall) Context(ctx context.Context) *ProjectsRegionsJobsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsJobsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsJobsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.jobs.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsJobsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource. If the resource does not exist, this will return an empty set of permissions, not a NOT_FOUND error.Note: This operation is designed to be used for building permission-aware UIs and command-line tools, not for authorization checking. This operation may \"fail open\" without warning.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/jobs/{jobsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.jobs.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.operations.cancel":

type ProjectsRegionsOperationsCancelCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
// The server makes a best effort to cancel the operation, but success
// is not guaranteed. If the server doesn't support this method, it
// returns google.rpc.Code.UNIMPLEMENTED. Clients can use
// Operations.GetOperation or other methods to check whether the
// cancellation succeeded or whether the operation completed despite
// cancellation. On successful cancellation, the operation is not
// deleted; instead, it becomes an operation with an Operation.error
// value with a google.rpc.Status.code of 1, corresponding to
// Code.CANCELLED.
func (r *ProjectsRegionsOperationsService) Cancel(name string) *ProjectsRegionsOperationsCancelCall {
	c := &ProjectsRegionsOperationsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsOperationsCancelCall) Fields(s ...googleapi.Field) *ProjectsRegionsOperationsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsOperationsCancelCall) Context(ctx context.Context) *ProjectsRegionsOperationsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsOperationsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsOperationsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.operations.cancel" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsOperationsCancelCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts asynchronous cancellation on a long-running operation. The server makes a best effort to cancel the operation, but success is not guaranteed. If the server doesn't support this method, it returns google.rpc.Code.UNIMPLEMENTED. Clients can use Operations.GetOperation or other methods to check whether the cancellation succeeded or whether the operation completed despite cancellation. On successful cancellation, the operation is not deleted; instead, it becomes an operation with an Operation.error value with a google.rpc.Status.code of 1, corresponding to Code.CANCELLED.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:cancel",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.operations.delete":

type ProjectsRegionsOperationsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a long-running operation. This method indicates that
// the client is no longer interested in the operation result. It does
// not cancel the operation. If the server doesn't support this method,
// it returns google.rpc.Code.UNIMPLEMENTED.
func (r *ProjectsRegionsOperationsService) Delete(name string) *ProjectsRegionsOperationsDeleteCall {
	c := &ProjectsRegionsOperationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsOperationsDeleteCall) Fields(s ...googleapi.Field) *ProjectsRegionsOperationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsOperationsDeleteCall) Context(ctx context.Context) *ProjectsRegionsOperationsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsOperationsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsOperationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.operations.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsOperationsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a long-running operation. This method indicates that the client is no longer interested in the operation result. It does not cancel the operation. If the server doesn't support this method, it returns google.rpc.Code.UNIMPLEMENTED.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.regions.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.operations.get":

type ProjectsRegionsOperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation. Clients can
// use this method to poll the operation result at intervals as
// recommended by the API service.
func (r *ProjectsRegionsOperationsService) Get(name string) *ProjectsRegionsOperationsGetCall {
	c := &ProjectsRegionsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsRegionsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsRegionsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsOperationsGetCall) Context(ctx context.Context) *ProjectsRegionsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsOperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsOperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation. Clients can use this method to poll the operation result at intervals as recommended by the API service.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.operations.getIamPolicy":

type ProjectsRegionsOperationsGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Gets the access control policy for a resource. Returns
// an empty policy if the resource exists and does not have a policy
// set.
func (r *ProjectsRegionsOperationsService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsRegionsOperationsGetIamPolicyCall {
	c := &ProjectsRegionsOperationsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsOperationsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsOperationsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsOperationsGetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsOperationsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsOperationsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsOperationsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.operations.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsOperationsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource. Returns an empty policy if the resource exists and does not have a policy set.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/operations/{operationsId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.operations.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.operations.list":

type ProjectsRegionsOperationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists operations that match the specified filter in the
// request. If the server doesn't support this method, it returns
// UNIMPLEMENTED.NOTE: the name binding allows API services to override
// the binding to use different resource name schemes, such as
// users/*/operations. To override the binding, API services can add a
// binding such as "/v1/{name=users/*}/operations" to their service
// configuration. For backwards compatibility, the default name includes
// the operations collection id, however overriding users must ensure
// the name binding is the parent resource, without the operations
// collection id.
func (r *ProjectsRegionsOperationsService) List(name string) *ProjectsRegionsOperationsListCall {
	c := &ProjectsRegionsOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsRegionsOperationsListCall) Filter(filter string) *ProjectsRegionsOperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsRegionsOperationsListCall) PageSize(pageSize int64) *ProjectsRegionsOperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsRegionsOperationsListCall) PageToken(pageToken string) *ProjectsRegionsOperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsRegionsOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsOperationsListCall) IfNoneMatch(entityTag string) *ProjectsRegionsOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsOperationsListCall) Context(ctx context.Context) *ProjectsRegionsOperationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsOperationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsOperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.operations.list" call.
// Exactly one of *ListOperationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOperationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsOperationsListCall) Do(opts ...googleapi.CallOption) (*ListOperationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOperationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request. If the server doesn't support this method, it returns UNIMPLEMENTED.NOTE: the name binding allows API services to override the binding to use different resource name schemes, such as users/*/operations. To override the binding, API services can add a binding such as \"/v1/{name=users/*}/operations\" to their service configuration. For backwards compatibility, the default name includes the operations collection id, however overriding users must ensure the name binding is the parent resource, without the operations collection id.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/operations",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the operation's parent resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/operations$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard list page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsRegionsOperationsListCall) Pages(ctx context.Context, f func(*ListOperationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dataproc.projects.regions.operations.setIamPolicy":

type ProjectsRegionsOperationsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any existing policy.
func (r *ProjectsRegionsOperationsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsRegionsOperationsSetIamPolicyCall {
	c := &ProjectsRegionsOperationsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsOperationsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsOperationsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsOperationsSetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsOperationsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsOperationsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsOperationsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.operations.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsOperationsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any existing policy.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/operations/{operationsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.operations.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.operations.testIamPermissions":

type ProjectsRegionsOperationsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource. If the resource does not exist, this will return
// an empty set of permissions, not a NOT_FOUND error.Note: This
// operation is designed to be used for building permission-aware UIs
// and command-line tools, not for authorization checking. This
// operation may "fail open" without warning.
func (r *ProjectsRegionsOperationsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsRegionsOperationsTestIamPermissionsCall {
	c := &ProjectsRegionsOperationsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsOperationsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsRegionsOperationsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsOperationsTestIamPermissionsCall) Context(ctx context.Context) *ProjectsRegionsOperationsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsOperationsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsOperationsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.operations.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsOperationsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource. If the resource does not exist, this will return an empty set of permissions, not a NOT_FOUND error.Note: This operation is designed to be used for building permission-aware UIs and command-line tools, not for authorization checking. This operation may \"fail open\" without warning.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/operations/{operationsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.operations.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.create":

type ProjectsRegionsWorkflowTemplatesCreateCall struct {
	s                *Service
	parent           string
	workflowtemplate *WorkflowTemplate
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Create: Creates new workflow template.
func (r *ProjectsRegionsWorkflowTemplatesService) Create(parent string, workflowtemplate *WorkflowTemplate) *ProjectsRegionsWorkflowTemplatesCreateCall {
	c := &ProjectsRegionsWorkflowTemplatesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.workflowtemplate = workflowtemplate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesCreateCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesCreateCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.workflowtemplate)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/workflowTemplates")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.create" call.
// Exactly one of *WorkflowTemplate or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WorkflowTemplate.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsWorkflowTemplatesCreateCall) Do(opts ...googleapi.CallOption) (*WorkflowTemplate, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WorkflowTemplate{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates new workflow template.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.workflowTemplates.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The \"resource name\" of the region, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/workflowTemplates",
	//   "request": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "response": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.delete":

type ProjectsRegionsWorkflowTemplatesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a workflow template. It does not cancel in-progress
// workflows.
func (r *ProjectsRegionsWorkflowTemplatesService) Delete(name string) *ProjectsRegionsWorkflowTemplatesDeleteCall {
	c := &ProjectsRegionsWorkflowTemplatesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Version sets the optional parameter "version": The version of
// workflow template to delete. If specified, will only delete the
// template if the current server version matches specified version.
func (c *ProjectsRegionsWorkflowTemplatesDeleteCall) Version(version int64) *ProjectsRegionsWorkflowTemplatesDeleteCall {
	c.urlParams_.Set("version", fmt.Sprint(version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesDeleteCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesDeleteCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsWorkflowTemplatesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a workflow template. It does not cancel in-progress workflows.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates/{workflowTemplatesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.regions.workflowTemplates.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The \"resource name\" of the workflow template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Optional. The version of workflow template to delete. If specified, will only delete the template if the current server version matches specified version.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.get":

type ProjectsRegionsWorkflowTemplatesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the latest workflow template.Can retrieve previously
// instantiated template by specifying optional version parameter.
func (r *ProjectsRegionsWorkflowTemplatesService) Get(name string) *ProjectsRegionsWorkflowTemplatesGetCall {
	c := &ProjectsRegionsWorkflowTemplatesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Version sets the optional parameter "version": The version of
// workflow template to retrieve. Only previously instatiated versions
// can be retrieved.If unspecified, retrieves the current version.
func (c *ProjectsRegionsWorkflowTemplatesGetCall) Version(version int64) *ProjectsRegionsWorkflowTemplatesGetCall {
	c.urlParams_.Set("version", fmt.Sprint(version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesGetCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsWorkflowTemplatesGetCall) IfNoneMatch(entityTag string) *ProjectsRegionsWorkflowTemplatesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesGetCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.get" call.
// Exactly one of *WorkflowTemplate or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WorkflowTemplate.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsWorkflowTemplatesGetCall) Do(opts ...googleapi.CallOption) (*WorkflowTemplate, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WorkflowTemplate{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the latest workflow template.Can retrieve previously instantiated template by specifying optional version parameter.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates/{workflowTemplatesId}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.workflowTemplates.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The \"resource name\" of the workflow template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Optional. The version of workflow template to retrieve. Only previously instatiated versions can be retrieved.If unspecified, retrieves the current version.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.getIamPolicy":

type ProjectsRegionsWorkflowTemplatesGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Gets the access control policy for a resource. Returns
// an empty policy if the resource exists and does not have a policy
// set.
func (r *ProjectsRegionsWorkflowTemplatesService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall {
	c := &ProjectsRegionsWorkflowTemplatesGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsWorkflowTemplatesGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource. Returns an empty policy if the resource exists and does not have a policy set.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates/{workflowTemplatesId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.workflowTemplates.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.instantiate":

type ProjectsRegionsWorkflowTemplatesInstantiateCall struct {
	s                                  *Service
	name                               string
	instantiateworkflowtemplaterequest *InstantiateWorkflowTemplateRequest
	urlParams_                         gensupport.URLParams
	ctx_                               context.Context
	header_                            http.Header
}

// Instantiate: Instantiates a template and begins execution.The
// returned Operation can be used to track execution of workflow by
// polling operations.get. The Operation will complete when entire
// workflow is finished.The running workflow can be aborted via
// operations.cancel. This will cause any inflight jobs to be cancelled
// and workflow-owned clusters to be deleted.The Operation.metadata will
// be WorkflowMetadata.On successful completion, Operation.response will
// be Empty.
func (r *ProjectsRegionsWorkflowTemplatesService) Instantiate(name string, instantiateworkflowtemplaterequest *InstantiateWorkflowTemplateRequest) *ProjectsRegionsWorkflowTemplatesInstantiateCall {
	c := &ProjectsRegionsWorkflowTemplatesInstantiateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.instantiateworkflowtemplaterequest = instantiateworkflowtemplaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesInstantiateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesInstantiateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesInstantiateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.instantiateworkflowtemplaterequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:instantiate")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.instantiate" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Instantiates a template and begins execution.The returned Operation can be used to track execution of workflow by polling operations.get. The Operation will complete when entire workflow is finished.The running workflow can be aborted via operations.cancel. This will cause any inflight jobs to be cancelled and workflow-owned clusters to be deleted.The Operation.metadata will be WorkflowMetadata.On successful completion, Operation.response will be Empty.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates/{workflowTemplatesId}:instantiate",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.workflowTemplates.instantiate",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The \"resource name\" of the workflow template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:instantiate",
	//   "request": {
	//     "$ref": "InstantiateWorkflowTemplateRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.instantiateInline":

type ProjectsRegionsWorkflowTemplatesInstantiateInlineCall struct {
	s                *Service
	parent           string
	workflowtemplate *WorkflowTemplate
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// InstantiateInline: Instantiates a template and begins execution.This
// method is equivalent to executing the sequence
// CreateWorkflowTemplate, InstantiateWorkflowTemplate,
// DeleteWorkflowTemplate.The returned Operation can be used to track
// execution of workflow by polling operations.get. The Operation will
// complete when entire workflow is finished.The running workflow can be
// aborted via operations.cancel. This will cause any inflight jobs to
// be cancelled and workflow-owned clusters to be deleted.The
// Operation.metadata will be WorkflowMetadata.On successful completion,
// Operation.response will be Empty.
func (r *ProjectsRegionsWorkflowTemplatesService) InstantiateInline(parent string, workflowtemplate *WorkflowTemplate) *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall {
	c := &ProjectsRegionsWorkflowTemplatesInstantiateInlineCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.workflowtemplate = workflowtemplate
	return c
}

// RequestId sets the optional parameter "requestId": A tag that
// prevents multiple concurrent workflow instances with the same tag
// from running. This mitigates risk of concurrent instances started due
// to retries.It is recommended to always set this value to a UUID
// (https://en.wikipedia.org/wiki/Universally_unique_identifier).The tag
// must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
// and hyphens (-). The maximum length is 40 characters.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall) RequestId(requestId string) *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall {
	c.urlParams_.Set("requestId", requestId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.workflowtemplate)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/workflowTemplates:instantiateInline")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.instantiateInline" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsWorkflowTemplatesInstantiateInlineCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Instantiates a template and begins execution.This method is equivalent to executing the sequence CreateWorkflowTemplate, InstantiateWorkflowTemplate, DeleteWorkflowTemplate.The returned Operation can be used to track execution of workflow by polling operations.get. The Operation will complete when entire workflow is finished.The running workflow can be aborted via operations.cancel. This will cause any inflight jobs to be cancelled and workflow-owned clusters to be deleted.The Operation.metadata will be WorkflowMetadata.On successful completion, Operation.response will be Empty.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates:instantiateInline",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.workflowTemplates.instantiateInline",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The \"resource name\" of the workflow template region, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "Optional. A tag that prevents multiple concurrent workflow instances with the same tag from running. This mitigates risk of concurrent instances started due to retries.It is recommended to always set this value to a UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier).The tag must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/workflowTemplates:instantiateInline",
	//   "request": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.list":

type ProjectsRegionsWorkflowTemplatesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists workflows that match the specified filter in the request.
func (r *ProjectsRegionsWorkflowTemplatesService) List(parent string) *ProjectsRegionsWorkflowTemplatesListCall {
	c := &ProjectsRegionsWorkflowTemplatesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in each response.
func (c *ProjectsRegionsWorkflowTemplatesListCall) PageSize(pageSize int64) *ProjectsRegionsWorkflowTemplatesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The page token,
// returned by a previous call, to request the next page of results.
func (c *ProjectsRegionsWorkflowTemplatesListCall) PageToken(pageToken string) *ProjectsRegionsWorkflowTemplatesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesListCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsWorkflowTemplatesListCall) IfNoneMatch(entityTag string) *ProjectsRegionsWorkflowTemplatesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesListCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/workflowTemplates")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.list" call.
// Exactly one of *ListWorkflowTemplatesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListWorkflowTemplatesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsWorkflowTemplatesListCall) Do(opts ...googleapi.CallOption) (*ListWorkflowTemplatesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListWorkflowTemplatesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists workflows that match the specified filter in the request.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.regions.workflowTemplates.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return in each response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The page token, returned by a previous call, to request the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The \"resource name\" of the region, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/workflowTemplates",
	//   "response": {
	//     "$ref": "ListWorkflowTemplatesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsRegionsWorkflowTemplatesListCall) Pages(ctx context.Context, f func(*ListWorkflowTemplatesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dataproc.projects.regions.workflowTemplates.setIamPolicy":

type ProjectsRegionsWorkflowTemplatesSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any existing policy.
func (r *ProjectsRegionsWorkflowTemplatesService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall {
	c := &ProjectsRegionsWorkflowTemplatesSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsRegionsWorkflowTemplatesSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any existing policy.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates/{workflowTemplatesId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.workflowTemplates.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.testIamPermissions":

type ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource. If the resource does not exist, this will return
// an empty set of permissions, not a NOT_FOUND error.Note: This
// operation is designed to be used for building permission-aware UIs
// and command-line tools, not for authorization checking. This
// operation may "fail open" without warning.
func (r *ProjectsRegionsWorkflowTemplatesService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall {
	c := &ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsWorkflowTemplatesTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource. If the resource does not exist, this will return an empty set of permissions, not a NOT_FOUND error.Note: This operation is designed to be used for building permission-aware UIs and command-line tools, not for authorization checking. This operation may \"fail open\" without warning.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates/{workflowTemplatesId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.regions.workflowTemplates.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested. See the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.regions.workflowTemplates.update":

type ProjectsRegionsWorkflowTemplatesUpdateCall struct {
	s                *Service
	name             string
	workflowtemplate *WorkflowTemplate
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Update: Updates (replaces) workflow template. The updated template
// must contain version that matches the current server version.
func (r *ProjectsRegionsWorkflowTemplatesService) Update(name string, workflowtemplate *WorkflowTemplate) *ProjectsRegionsWorkflowTemplatesUpdateCall {
	c := &ProjectsRegionsWorkflowTemplatesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.workflowtemplate = workflowtemplate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsWorkflowTemplatesUpdateCall) Fields(s ...googleapi.Field) *ProjectsRegionsWorkflowTemplatesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsWorkflowTemplatesUpdateCall) Context(ctx context.Context) *ProjectsRegionsWorkflowTemplatesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsRegionsWorkflowTemplatesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsRegionsWorkflowTemplatesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.workflowtemplate)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dataproc.projects.regions.workflowTemplates.update" call.
// Exactly one of *WorkflowTemplate or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WorkflowTemplate.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsWorkflowTemplatesUpdateCall) Do(opts ...googleapi.CallOption) (*WorkflowTemplate, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WorkflowTemplate{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates (replaces) workflow template. The updated template must contain version that matches the current server version.",
	//   "flatPath": "v1/projects/{projectsId}/regions/{regionsId}/workflowTemplates/{workflowTemplatesId}",
	//   "httpMethod": "PUT",
	//   "id": "dataproc.projects.regions.workflowTemplates.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Output only. The \"resource name\" of the template, as described in https://cloud.google.com/apis/design/resource_names of the form projects/{project_id}/regions/{region}/workflowTemplates/{template_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/regions/[^/]+/workflowTemplates/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "response": {
	//     "$ref": "WorkflowTemplate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
