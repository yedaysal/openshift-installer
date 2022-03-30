package storagesync

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ChangeDetectionMode enumerates the values for change detection mode.
type ChangeDetectionMode string

const (
	// Default ...
	Default ChangeDetectionMode = "Default"
	// Recursive ...
	Recursive ChangeDetectionMode = "Recursive"
)

// PossibleChangeDetectionModeValues returns an array of possible values for the ChangeDetectionMode const type.
func PossibleChangeDetectionModeValues() []ChangeDetectionMode {
	return []ChangeDetectionMode{Default, Recursive}
}

// FeatureStatus enumerates the values for feature status.
type FeatureStatus string

const (
	// Off ...
	Off FeatureStatus = "off"
	// On ...
	On FeatureStatus = "on"
)

// PossibleFeatureStatusValues returns an array of possible values for the FeatureStatus const type.
func PossibleFeatureStatusValues() []FeatureStatus {
	return []FeatureStatus{Off, On}
}

// IncomingTrafficPolicy enumerates the values for incoming traffic policy.
type IncomingTrafficPolicy string

const (
	// AllowAllTraffic ...
	AllowAllTraffic IncomingTrafficPolicy = "AllowAllTraffic"
	// AllowVirtualNetworksOnly ...
	AllowVirtualNetworksOnly IncomingTrafficPolicy = "AllowVirtualNetworksOnly"
)

// PossibleIncomingTrafficPolicyValues returns an array of possible values for the IncomingTrafficPolicy const type.
func PossibleIncomingTrafficPolicyValues() []IncomingTrafficPolicy {
	return []IncomingTrafficPolicy{AllowAllTraffic, AllowVirtualNetworksOnly}
}

// InitialDownloadPolicy enumerates the values for initial download policy.
type InitialDownloadPolicy string

const (
	// AvoidTieredFiles ...
	AvoidTieredFiles InitialDownloadPolicy = "AvoidTieredFiles"
	// NamespaceOnly ...
	NamespaceOnly InitialDownloadPolicy = "NamespaceOnly"
	// NamespaceThenModifiedFiles ...
	NamespaceThenModifiedFiles InitialDownloadPolicy = "NamespaceThenModifiedFiles"
)

// PossibleInitialDownloadPolicyValues returns an array of possible values for the InitialDownloadPolicy const type.
func PossibleInitialDownloadPolicyValues() []InitialDownloadPolicy {
	return []InitialDownloadPolicy{AvoidTieredFiles, NamespaceOnly, NamespaceThenModifiedFiles}
}

// LocalCacheMode enumerates the values for local cache mode.
type LocalCacheMode string

const (
	// DownloadNewAndModifiedFiles ...
	DownloadNewAndModifiedFiles LocalCacheMode = "DownloadNewAndModifiedFiles"
	// UpdateLocallyCachedFiles ...
	UpdateLocallyCachedFiles LocalCacheMode = "UpdateLocallyCachedFiles"
)

// PossibleLocalCacheModeValues returns an array of possible values for the LocalCacheMode const type.
func PossibleLocalCacheModeValues() []LocalCacheMode {
	return []LocalCacheMode{DownloadNewAndModifiedFiles, UpdateLocallyCachedFiles}
}

// NameAvailabilityReason enumerates the values for name availability reason.
type NameAvailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists NameAvailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid NameAvailabilityReason = "Invalid"
)

// PossibleNameAvailabilityReasonValues returns an array of possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{AlreadyExists, Invalid}
}

// OperationDirection enumerates the values for operation direction.
type OperationDirection string

const (
	// Cancel ...
	Cancel OperationDirection = "cancel"
	// Do ...
	Do OperationDirection = "do"
	// Undo ...
	Undo OperationDirection = "undo"
)

// PossibleOperationDirectionValues returns an array of possible values for the OperationDirection const type.
func PossibleOperationDirectionValues() []OperationDirection {
	return []OperationDirection{Cancel, Do, Undo}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// Creating ...
	Creating PrivateEndpointConnectionProvisioningState = "Creating"
	// Deleting ...
	Deleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// Failed ...
	Failed PrivateEndpointConnectionProvisioningState = "Failed"
	// Succeeded ...
	Succeeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{Creating, Deleting, Failed, Succeeded}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// Approved ...
	Approved PrivateEndpointServiceConnectionStatus = "Approved"
	// Pending ...
	Pending PrivateEndpointServiceConnectionStatus = "Pending"
	// Rejected ...
	Rejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{Approved, Pending, Rejected}
}

// ProgressType enumerates the values for progress type.
type ProgressType string

const (
	// Download ...
	Download ProgressType = "download"
	// Initialize ...
	Initialize ProgressType = "initialize"
	// None ...
	None ProgressType = "none"
	// Recall ...
	Recall ProgressType = "recall"
	// Upload ...
	Upload ProgressType = "upload"
)

// PossibleProgressTypeValues returns an array of possible values for the ProgressType const type.
func PossibleProgressTypeValues() []ProgressType {
	return []ProgressType{Download, Initialize, None, Recall, Upload}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// Deleted ...
	Deleted Reason = "Deleted"
	// Registered ...
	Registered Reason = "Registered"
	// Suspended ...
	Suspended Reason = "Suspended"
	// Unregistered ...
	Unregistered Reason = "Unregistered"
	// Warned ...
	Warned Reason = "Warned"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{Deleted, Registered, Suspended, Unregistered, Warned}
}

// ServerEndpointCloudTieringHealthState enumerates the values for server endpoint cloud tiering health state.
type ServerEndpointCloudTieringHealthState string

const (
	// ServerEndpointCloudTieringHealthStateError ...
	ServerEndpointCloudTieringHealthStateError ServerEndpointCloudTieringHealthState = "Error"
	// ServerEndpointCloudTieringHealthStateHealthy ...
	ServerEndpointCloudTieringHealthStateHealthy ServerEndpointCloudTieringHealthState = "Healthy"
)

// PossibleServerEndpointCloudTieringHealthStateValues returns an array of possible values for the ServerEndpointCloudTieringHealthState const type.
func PossibleServerEndpointCloudTieringHealthStateValues() []ServerEndpointCloudTieringHealthState {
	return []ServerEndpointCloudTieringHealthState{ServerEndpointCloudTieringHealthStateError, ServerEndpointCloudTieringHealthStateHealthy}
}

// ServerEndpointOfflineDataTransferState enumerates the values for server endpoint offline data transfer
// state.
type ServerEndpointOfflineDataTransferState string

const (
	// Complete ...
	Complete ServerEndpointOfflineDataTransferState = "Complete"
	// InProgress ...
	InProgress ServerEndpointOfflineDataTransferState = "InProgress"
	// NotRunning ...
	NotRunning ServerEndpointOfflineDataTransferState = "NotRunning"
	// Stopping ...
	Stopping ServerEndpointOfflineDataTransferState = "Stopping"
)

// PossibleServerEndpointOfflineDataTransferStateValues returns an array of possible values for the ServerEndpointOfflineDataTransferState const type.
func PossibleServerEndpointOfflineDataTransferStateValues() []ServerEndpointOfflineDataTransferState {
	return []ServerEndpointOfflineDataTransferState{Complete, InProgress, NotRunning, Stopping}
}

// ServerEndpointSyncActivityState enumerates the values for server endpoint sync activity state.
type ServerEndpointSyncActivityState string

const (
	// ServerEndpointSyncActivityStateDownload ...
	ServerEndpointSyncActivityStateDownload ServerEndpointSyncActivityState = "Download"
	// ServerEndpointSyncActivityStateUpload ...
	ServerEndpointSyncActivityStateUpload ServerEndpointSyncActivityState = "Upload"
	// ServerEndpointSyncActivityStateUploadAndDownload ...
	ServerEndpointSyncActivityStateUploadAndDownload ServerEndpointSyncActivityState = "UploadAndDownload"
)

// PossibleServerEndpointSyncActivityStateValues returns an array of possible values for the ServerEndpointSyncActivityState const type.
func PossibleServerEndpointSyncActivityStateValues() []ServerEndpointSyncActivityState {
	return []ServerEndpointSyncActivityState{ServerEndpointSyncActivityStateDownload, ServerEndpointSyncActivityStateUpload, ServerEndpointSyncActivityStateUploadAndDownload}
}

// ServerEndpointSyncHealthState enumerates the values for server endpoint sync health state.
type ServerEndpointSyncHealthState string

const (
	// ServerEndpointSyncHealthStateError ...
	ServerEndpointSyncHealthStateError ServerEndpointSyncHealthState = "Error"
	// ServerEndpointSyncHealthStateHealthy ...
	ServerEndpointSyncHealthStateHealthy ServerEndpointSyncHealthState = "Healthy"
	// ServerEndpointSyncHealthStateNoActivity ...
	ServerEndpointSyncHealthStateNoActivity ServerEndpointSyncHealthState = "NoActivity"
	// ServerEndpointSyncHealthStateSyncBlockedForChangeDetectionPostRestore ...
	ServerEndpointSyncHealthStateSyncBlockedForChangeDetectionPostRestore ServerEndpointSyncHealthState = "SyncBlockedForChangeDetectionPostRestore"
	// ServerEndpointSyncHealthStateSyncBlockedForRestore ...
	ServerEndpointSyncHealthStateSyncBlockedForRestore ServerEndpointSyncHealthState = "SyncBlockedForRestore"
)

// PossibleServerEndpointSyncHealthStateValues returns an array of possible values for the ServerEndpointSyncHealthState const type.
func PossibleServerEndpointSyncHealthStateValues() []ServerEndpointSyncHealthState {
	return []ServerEndpointSyncHealthState{ServerEndpointSyncHealthStateError, ServerEndpointSyncHealthStateHealthy, ServerEndpointSyncHealthStateNoActivity, ServerEndpointSyncHealthStateSyncBlockedForChangeDetectionPostRestore, ServerEndpointSyncHealthStateSyncBlockedForRestore}
}

// WorkflowStatus enumerates the values for workflow status.
type WorkflowStatus string

const (
	// WorkflowStatusAborted ...
	WorkflowStatusAborted WorkflowStatus = "aborted"
	// WorkflowStatusActive ...
	WorkflowStatusActive WorkflowStatus = "active"
	// WorkflowStatusExpired ...
	WorkflowStatusExpired WorkflowStatus = "expired"
	// WorkflowStatusFailed ...
	WorkflowStatusFailed WorkflowStatus = "failed"
	// WorkflowStatusSucceeded ...
	WorkflowStatusSucceeded WorkflowStatus = "succeeded"
)

// PossibleWorkflowStatusValues returns an array of possible values for the WorkflowStatus const type.
func PossibleWorkflowStatusValues() []WorkflowStatus {
	return []WorkflowStatus{WorkflowStatusAborted, WorkflowStatusActive, WorkflowStatusExpired, WorkflowStatusFailed, WorkflowStatusSucceeded}
}