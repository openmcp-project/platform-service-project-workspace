package v1alpha1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	EventReasonReconcileFailed    = "ReconcileFailed"
	EventReasonReconcileSucceeded = "ReconcileSucceeded"

	SourceBuiltin                = "Builtin"
	SourceProjectWorkspaceConfig = "ProjectWorkspaceConfig"
	SourceServiceProviderPrefix  = "ServiceProvider"
)

// ProjectWorkspaceConfigSpec defines the desired state of ProjectWorkspaceConfig
type ProjectWorkspaceConfigSpec struct {
	// +optional
	Project ProjectConfig `json:"project"`
	// +optional
	Workspace WorkspaceConfig `json:"workspace"`
	// MemberOverrides allows to specify users and groups which should have admin permissions to projects and workspaces.
	// This basically disables the 'you must be admin of a project/workspace in order to modify it' check.
	// Leave empty to disable.
	// +optional
	MemberOverrides MemberOverrides `json:"memberOverrides,omitempty"`
	// Webhook contains the configuration for the webhooks.
	// +optional
	Webhook WebhookConfig `json:"webhook"`
}

// ProjectWorkspaceConfig is the Schema for the ProjectWorkspaceConfigs API
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=pwcfg
// +kubebuilder:metadata:labels="openmcp.cloud/cluster=platform"
type ProjectWorkspaceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec ProjectWorkspaceConfigSpec `json:"spec"`
}

// ProjectConfig contains the configuration for projects.
type ProjectConfig struct {
	// +optional
	ResourcesBlockingDeletion []metav1.GroupVersionKind `json:"resourcesBlockingDeletion,omitempty"`
	// AdditionalPermissions defines additional permissions users should have in a project, depending on their role.
	// +optional
	AdditionalPermissions map[ProjectMemberRole][]rbacv1.PolicyRule `json:"additionalPermissions,omitempty"`
	// PropagateLabelsToProjectNamespace lists labels which should be propagated to the project namespace.
	// +optional
	PropagateLabelsToProjectNamespace []string `json:"propagateLabelsToProjectNamespace,omitempty"`
	// PropagateLabelsToWorkspaceNamespaces lists labels which should be propagated to all namespaces belonging to workspaces in the project.
	// If a label is configured to be propagated from Workspace and Project, but has a different value, the value from the Workspace will be used.
	// +optional
	PropagateLabelsToWorkspaceNamespaces []string `json:"propagateLabelsToWorkspaceNamespaces,omitempty"`
}

// WorkspaceConfig contains the configuration for workspaces.
type WorkspaceConfig struct {
	// +optional
	ResourcesBlockingDeletion []metav1.GroupVersionKind `json:"resourcesBlockingDeletion,omitempty"`
	// AdditionalPermissions defines additional permissions users should have in a workspace, depending on their role.
	// +optional
	AdditionalPermissions map[WorkspaceMemberRole][]rbacv1.PolicyRule `json:"additionalPermissions,omitempty"`
	// PropagateLabelsToWorkspaceNamespace lists labels which should be propagated to the workspace namespace.
	// +optional
	PropagateLabelsToWorkspaceNamespace []string `json:"propagateLabelsToWorkspaceNamespace,omitempty"`
}

type WebhookConfig struct {
	// Disabled specifies whether the webhooks should be disabled.
	// +optional
	Disabled bool `json:"disabled"`
}

// +kubebuilder:object:root=true

// ProjectWorkspaceConfigList contains a list of ProjectWorkspaceConfig
type ProjectWorkspaceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ProjectWorkspaceConfig `json:"items"`
}

func init() {
	RegisterToSchemeBuilder(&ProjectWorkspaceConfig{}, &ProjectWorkspaceConfigList{})
}

// SetDefaults sets the default values for the project workspace configuration when not set.
func (pwc *ProjectWorkspaceConfig) SetDefaults() {}

// Validate validates the project workspace configuration.
func (pwc *ProjectWorkspaceConfig) Validate() error {
	return nil
}
