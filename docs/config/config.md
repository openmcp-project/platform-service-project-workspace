# Platform Service Configuration

The platform service requires some configuration, which is expected in the form of a `ProjectWorkspaceConfig` resource with the same name as the `PlatformService` resource.

```yaml
apiVersion: core.openmcp.cloud/v1alpha1
kind: ProjectWorkspaceConfig
metadata:
  name: project-workspace # must match name of the PlatformService
spec:
  project:
    resourcesBlockingDeletion:
    - group: mygroup.example.org
      version: v1alpha1
      kind: MyProjectScopedResource
    additionalPermissions:
      admin:
      - apiGroups:
        - mygroup.example.org
        resources:
        - myprojectscopedresources
        verbs:
        - '*'
      view:
      - apiGroups:
        - mygroup.example.org
        resources:
        - myprojectscopedresources
        verbs:
        - get
        - list
        - watch
    propagateLabelsToProjectNamespace:
    - core.openmcp.cloud/mylabel
    propagateLabelsToWorkspaceNamespaces:
    - core.openmcp.cloud/mylabel
  workspace:
    resourcesBlockingDeletion:
    - group: mygroup.example.org
      version: v1alpha1
      kind: MyWorkspaceScopedResource
    additionalPermissions: <...>
    propagateLabelsToWorkspaceNamespace:
    - core.openmcp.cloud/mylabel
  memberOverrides:
  - kind: User
    name: kubernetes-admin
    roles:
    - admin
  webhook:
    disabled: false
```

All fields directly under `spec` are optional. They will be explained in the section below.

## Configuration Options

### Project Configuration

The project configuration is under `spec.project` and contains the following sub-configurations:

#### Resources Blocking Deletion

The optional field `spec.project.resourcesBlockingDeletion` allows to list `GroupVersionKind`s of kubernetes resources. The project controller will only remove its finalizer from the `Project` resource, allowing it to be deleted, when none of the listed resources exist within the project's namespace.

Note that workspaces (api group `core.openmcp.cloud`, version `v1alpha1`, kind `Workspace`) are by default part of this list and don't have to be added via the config.

#### Additional Permissions

Via the optional `spec.project.additionalPermissions` field, end-users can be granted additional permissions within their project namespaces. The field expects a mapping from project roles (`admin`, `view`) to standard k8s RBAC definitions. Users with the corresponding role within the project will have the specified permissions within the project's namespace, in addition to the default ones.

By default, users have permissions for workspaces and serviceaccounts, with the `view` role having only read access and the `admin` role having full access for these resources. Both roles can also list pods (there are usually no pods on the onboarding cluster, this is mainly to prevent k9s from crashing) and read resourcequotas. Admins can also create tokens for serviceaccounts and manage secrets.

#### Label Propagation

The optional `propagateLabelsToProjectNamespace` field allows to list keys of metadata labels which will be propagated from the `Project` resource to its namespace, if they exist.
Similarly, `propagateLabelsToWorkspaceNamespaces` can be used to propagate labels from the `Project` to the namespaces of all `Workspace` resources within the project's namespace. If the workspace configuration's `propagateLabelsToWorkspaceNamespace` contains one or more labely keys which overlap with the ones from this list, the labels from the `Workspace` will take precedence over the ones from the `Project`, if the label key exists on both resources.

Note that this feature will only ever add or update labels on the respective namespaces. Neither removing a label from a `Project`, nor removing its key from this configuration, will cause the label to be removed from the namespace, once it has been propagated there.

### Workspace configuration

The workspace configuration under `spec.workspace` is pretty much identical to the project one, only that they affect workspace namespaces instead of project ones. Therefore, the sections below will just list the different defaults.

In addition to the defaults listed below, both the resources that block workspace deletion and the permissions for end-users are dynamically adapted to include service resources. This is explained in more detail in the [config controller documentation](../controllers/config.md).

#### Resources Blocking Deletion

By default, only `ControlPlane` resources block workspace deletion. If the platform service is running in [v1 support mode](./v1.md), `ManagedControlPlane` and `ClusterAdmin` resources will also block workspace deletion.

#### Additional Permissions

Both roles can manage (read for `view`, read and write for `admin`) `ControlPlane` resources, as well as secrets, configmaps, and serviceaccounts. In [v1 support mode](./v1.md), `ManagedControlPlane` and `ClusterAdmin` resources are covered as well. Similar to projects, both roles can list pods and read resourcequotas, with the `admin` additionally being able to create tokens for serviceaccounts.

#### Label Propagation

The optional `propagateLabelsToWorkspaceNamespace` field can be used to propagate labels from the `Workspace` resource onto its corresponding namespace.

As mentioned for the propagation from `Project` resources described above, this feature will only add or update labels on the respective namespaces, never delete them.

### Member Overrides

This configuration has its own [documentation](member_overrides.md).

### Webhook

This optional section just allows to disable the webhooks by setting `spec.webhook.disabled` to `true`.
