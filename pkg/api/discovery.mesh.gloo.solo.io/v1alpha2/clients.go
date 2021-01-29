// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha2

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the discovery.mesh.gloo.solo.io/v1alpha2 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the discovery.mesh.gloo.solo.io/v1alpha2 APIs
type Clientset interface {
	// clienset for the discovery.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
	TrafficTargets() TrafficTargetClient
	// clienset for the discovery.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
	Workloads() WorkloadClient
	// clienset for the discovery.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
	Meshes() MeshClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the discovery.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
func (c *clientSet) TrafficTargets() TrafficTargetClient {
	return NewTrafficTargetClient(c.client)
}

// clienset for the discovery.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
func (c *clientSet) Workloads() WorkloadClient {
	return NewWorkloadClient(c.client)
}

// clienset for the discovery.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
func (c *clientSet) Meshes() MeshClient {
	return NewMeshClient(c.client)
}

// Reader knows how to read and list TrafficTargets.
type TrafficTargetReader interface {
	// Get retrieves a TrafficTarget for the given object key
	GetTrafficTarget(ctx context.Context, key client.ObjectKey) (*TrafficTarget, error)

	// List retrieves list of TrafficTargets for a given namespace and list options.
	ListTrafficTarget(ctx context.Context, opts ...client.ListOption) (*TrafficTargetList, error)
}

// TrafficTargetTransitionFunction instructs the TrafficTargetWriter how to transition between an existing
// TrafficTarget object and a desired on an Upsert
type TrafficTargetTransitionFunction func(existing, desired *TrafficTarget) error

// Writer knows how to create, delete, and update TrafficTargets.
type TrafficTargetWriter interface {
	// Create saves the TrafficTarget object.
	CreateTrafficTarget(ctx context.Context, obj *TrafficTarget, opts ...client.CreateOption) error

	// Delete deletes the TrafficTarget object.
	DeleteTrafficTarget(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given TrafficTarget object.
	UpdateTrafficTarget(ctx context.Context, obj *TrafficTarget, opts ...client.UpdateOption) error

	// Patch patches the given TrafficTarget object.
	PatchTrafficTarget(ctx context.Context, obj *TrafficTarget, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all TrafficTarget objects matching the given options.
	DeleteAllOfTrafficTarget(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the TrafficTarget object.
	UpsertTrafficTarget(ctx context.Context, obj *TrafficTarget, transitionFuncs ...TrafficTargetTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a TrafficTarget object.
type TrafficTargetStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given TrafficTarget object.
	UpdateTrafficTargetStatus(ctx context.Context, obj *TrafficTarget, opts ...client.UpdateOption) error

	// Patch patches the given TrafficTarget object's subresource.
	PatchTrafficTargetStatus(ctx context.Context, obj *TrafficTarget, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on TrafficTargets.
type TrafficTargetClient interface {
	TrafficTargetReader
	TrafficTargetWriter
	TrafficTargetStatusWriter
}

type trafficTargetClient struct {
	client client.Client
}

func NewTrafficTargetClient(client client.Client) *trafficTargetClient {
	return &trafficTargetClient{client: client}
}

func (c *trafficTargetClient) GetTrafficTarget(ctx context.Context, key client.ObjectKey) (*TrafficTarget, error) {
	obj := &TrafficTarget{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *trafficTargetClient) ListTrafficTarget(ctx context.Context, opts ...client.ListOption) (*TrafficTargetList, error) {
	list := &TrafficTargetList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *trafficTargetClient) CreateTrafficTarget(ctx context.Context, obj *TrafficTarget, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *trafficTargetClient) DeleteTrafficTarget(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &TrafficTarget{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *trafficTargetClient) UpdateTrafficTarget(ctx context.Context, obj *TrafficTarget, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *trafficTargetClient) PatchTrafficTarget(ctx context.Context, obj *TrafficTarget, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *trafficTargetClient) DeleteAllOfTrafficTarget(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &TrafficTarget{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *trafficTargetClient) UpsertTrafficTarget(ctx context.Context, obj *TrafficTarget, transitionFuncs ...TrafficTargetTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*TrafficTarget), desired.(*TrafficTarget)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *trafficTargetClient) UpdateTrafficTargetStatus(ctx context.Context, obj *TrafficTarget, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *trafficTargetClient) PatchTrafficTargetStatus(ctx context.Context, obj *TrafficTarget, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides TrafficTargetClients for multiple clusters.
type MulticlusterTrafficTargetClient interface {
	// Cluster returns a TrafficTargetClient for the given cluster
	Cluster(cluster string) (TrafficTargetClient, error)
}

type multiclusterTrafficTargetClient struct {
	client multicluster.Client
}

func NewMulticlusterTrafficTargetClient(client multicluster.Client) MulticlusterTrafficTargetClient {
	return &multiclusterTrafficTargetClient{client: client}
}

func (m *multiclusterTrafficTargetClient) Cluster(cluster string) (TrafficTargetClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewTrafficTargetClient(client), nil
}

// Reader knows how to read and list Workloads.
type WorkloadReader interface {
	// Get retrieves a Workload for the given object key
	GetWorkload(ctx context.Context, key client.ObjectKey) (*Workload, error)

	// List retrieves list of Workloads for a given namespace and list options.
	ListWorkload(ctx context.Context, opts ...client.ListOption) (*WorkloadList, error)
}

// WorkloadTransitionFunction instructs the WorkloadWriter how to transition between an existing
// Workload object and a desired on an Upsert
type WorkloadTransitionFunction func(existing, desired *Workload) error

// Writer knows how to create, delete, and update Workloads.
type WorkloadWriter interface {
	// Create saves the Workload object.
	CreateWorkload(ctx context.Context, obj *Workload, opts ...client.CreateOption) error

	// Delete deletes the Workload object.
	DeleteWorkload(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Workload object.
	UpdateWorkload(ctx context.Context, obj *Workload, opts ...client.UpdateOption) error

	// Patch patches the given Workload object.
	PatchWorkload(ctx context.Context, obj *Workload, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Workload objects matching the given options.
	DeleteAllOfWorkload(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Workload object.
	UpsertWorkload(ctx context.Context, obj *Workload, transitionFuncs ...WorkloadTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Workload object.
type WorkloadStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Workload object.
	UpdateWorkloadStatus(ctx context.Context, obj *Workload, opts ...client.UpdateOption) error

	// Patch patches the given Workload object's subresource.
	PatchWorkloadStatus(ctx context.Context, obj *Workload, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Workloads.
type WorkloadClient interface {
	WorkloadReader
	WorkloadWriter
	WorkloadStatusWriter
}

type workloadClient struct {
	client client.Client
}

func NewWorkloadClient(client client.Client) *workloadClient {
	return &workloadClient{client: client}
}

func (c *workloadClient) GetWorkload(ctx context.Context, key client.ObjectKey) (*Workload, error) {
	obj := &Workload{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *workloadClient) ListWorkload(ctx context.Context, opts ...client.ListOption) (*WorkloadList, error) {
	list := &WorkloadList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *workloadClient) CreateWorkload(ctx context.Context, obj *Workload, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *workloadClient) DeleteWorkload(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Workload{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *workloadClient) UpdateWorkload(ctx context.Context, obj *Workload, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *workloadClient) PatchWorkload(ctx context.Context, obj *Workload, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *workloadClient) DeleteAllOfWorkload(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Workload{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *workloadClient) UpsertWorkload(ctx context.Context, obj *Workload, transitionFuncs ...WorkloadTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Workload), desired.(*Workload)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *workloadClient) UpdateWorkloadStatus(ctx context.Context, obj *Workload, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *workloadClient) PatchWorkloadStatus(ctx context.Context, obj *Workload, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides WorkloadClients for multiple clusters.
type MulticlusterWorkloadClient interface {
	// Cluster returns a WorkloadClient for the given cluster
	Cluster(cluster string) (WorkloadClient, error)
}

type multiclusterWorkloadClient struct {
	client multicluster.Client
}

func NewMulticlusterWorkloadClient(client multicluster.Client) MulticlusterWorkloadClient {
	return &multiclusterWorkloadClient{client: client}
}

func (m *multiclusterWorkloadClient) Cluster(cluster string) (WorkloadClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewWorkloadClient(client), nil
}

// Reader knows how to read and list Meshs.
type MeshReader interface {
	// Get retrieves a Mesh for the given object key
	GetMesh(ctx context.Context, key client.ObjectKey) (*Mesh, error)

	// List retrieves list of Meshs for a given namespace and list options.
	ListMesh(ctx context.Context, opts ...client.ListOption) (*MeshList, error)
}

// MeshTransitionFunction instructs the MeshWriter how to transition between an existing
// Mesh object and a desired on an Upsert
type MeshTransitionFunction func(existing, desired *Mesh) error

// Writer knows how to create, delete, and update Meshs.
type MeshWriter interface {
	// Create saves the Mesh object.
	CreateMesh(ctx context.Context, obj *Mesh, opts ...client.CreateOption) error

	// Delete deletes the Mesh object.
	DeleteMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Mesh object.
	UpdateMesh(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error

	// Patch patches the given Mesh object.
	PatchMesh(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Mesh objects matching the given options.
	DeleteAllOfMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Mesh object.
	UpsertMesh(ctx context.Context, obj *Mesh, transitionFuncs ...MeshTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Mesh object.
type MeshStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Mesh object.
	UpdateMeshStatus(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error

	// Patch patches the given Mesh object's subresource.
	PatchMeshStatus(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Meshs.
type MeshClient interface {
	MeshReader
	MeshWriter
	MeshStatusWriter
}

type meshClient struct {
	client client.Client
}

func NewMeshClient(client client.Client) *meshClient {
	return &meshClient{client: client}
}

func (c *meshClient) GetMesh(ctx context.Context, key client.ObjectKey) (*Mesh, error) {
	obj := &Mesh{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *meshClient) ListMesh(ctx context.Context, opts ...client.ListOption) (*MeshList, error) {
	list := &MeshList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *meshClient) CreateMesh(ctx context.Context, obj *Mesh, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *meshClient) DeleteMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Mesh{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *meshClient) UpdateMesh(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *meshClient) PatchMesh(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *meshClient) DeleteAllOfMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Mesh{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *meshClient) UpsertMesh(ctx context.Context, obj *Mesh, transitionFuncs ...MeshTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Mesh), desired.(*Mesh)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *meshClient) UpdateMeshStatus(ctx context.Context, obj *Mesh, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *meshClient) PatchMeshStatus(ctx context.Context, obj *Mesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides MeshClients for multiple clusters.
type MulticlusterMeshClient interface {
	// Cluster returns a MeshClient for the given cluster
	Cluster(cluster string) (MeshClient, error)
}

type multiclusterMeshClient struct {
	client multicluster.Client
}

func NewMulticlusterMeshClient(client multicluster.Client) MulticlusterMeshClient {
	return &multiclusterMeshClient{client: client}
}

func (m *multiclusterMeshClient) Cluster(cluster string) (MeshClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewMeshClient(client), nil
}
