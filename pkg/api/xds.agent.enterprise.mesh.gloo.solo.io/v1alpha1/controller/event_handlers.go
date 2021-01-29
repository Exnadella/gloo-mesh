// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the XdsConfig Resource
// DEPRECATED: Prefer reconciler pattern.
type XdsConfigEventHandler interface {
	CreateXdsConfig(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
	UpdateXdsConfig(old, new *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
	DeleteXdsConfig(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
	GenericXdsConfig(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
}

type XdsConfigEventHandlerFuncs struct {
	OnCreate  func(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
	OnUpdate  func(old, new *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
	OnDelete  func(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
	OnGeneric func(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error
}

func (f *XdsConfigEventHandlerFuncs) CreateXdsConfig(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *XdsConfigEventHandlerFuncs) DeleteXdsConfig(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *XdsConfigEventHandlerFuncs) UpdateXdsConfig(objOld, objNew *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *XdsConfigEventHandlerFuncs) GenericXdsConfig(obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type XdsConfigEventWatcher interface {
	AddEventHandler(ctx context.Context, h XdsConfigEventHandler, predicates ...predicate.Predicate) error
}

type xdsConfigEventWatcher struct {
	watcher events.EventWatcher
}

func NewXdsConfigEventWatcher(name string, mgr manager.Manager) XdsConfigEventWatcher {
	return &xdsConfigEventWatcher{
		watcher: events.NewWatcher(name, mgr, &xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig{}),
	}
}

func (c *xdsConfigEventWatcher) AddEventHandler(ctx context.Context, h XdsConfigEventHandler, predicates ...predicate.Predicate) error {
	handler := genericXdsConfigHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericXdsConfigHandler implements a generic events.EventHandler
type genericXdsConfigHandler struct {
	handler XdsConfigEventHandler
}

func (h genericXdsConfigHandler) Create(object client.Object) error {
	obj, ok := object.(*xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig)
	if !ok {
		return errors.Errorf("internal error: XdsConfig handler received event for %T", object)
	}
	return h.handler.CreateXdsConfig(obj)
}

func (h genericXdsConfigHandler) Delete(object client.Object) error {
	obj, ok := object.(*xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig)
	if !ok {
		return errors.Errorf("internal error: XdsConfig handler received event for %T", object)
	}
	return h.handler.DeleteXdsConfig(obj)
}

func (h genericXdsConfigHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig)
	if !ok {
		return errors.Errorf("internal error: XdsConfig handler received event for %T", old)
	}
	objNew, ok := new.(*xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig)
	if !ok {
		return errors.Errorf("internal error: XdsConfig handler received event for %T", new)
	}
	return h.handler.UpdateXdsConfig(objOld, objNew)
}

func (h genericXdsConfigHandler) Generic(object client.Object) error {
	obj, ok := object.(*xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig)
	if !ok {
		return errors.Errorf("internal error: XdsConfig handler received event for %T", object)
	}
	return h.handler.GenericXdsConfig(obj)
}
