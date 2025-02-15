package kubecontroller

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type NodeNetworkConfigFilter struct {
	predicate.Funcs
	nodeName string
}

// Returns true if request is to be processed by Reconciler
// Checks that old generation equals new generation because status changes don't change generation number
func (n NodeNetworkConfigFilter) Update(e event.UpdateEvent) bool {
	isNodeName := n.isNodeName(e.ObjectOld.GetName())
	oldGeneration := e.ObjectOld.GetGeneration()
	newGeneration := e.ObjectOld.GetGeneration()
	return (oldGeneration == newGeneration) && isNodeName
}

// Only process create events if CRD name equals this host's name
func (n NodeNetworkConfigFilter) Create(e event.CreateEvent) bool {
	return n.isNodeName(e.Object.GetName())
}

// Delete is a noop filter to ignore all delete events.
// TODO: Decide what deleting crd means with DNC
func (n NodeNetworkConfigFilter) Delete(e event.DeleteEvent) bool {
	return false
}

// Given a string, returns if that string equals the nodename running this program
func (n NodeNetworkConfigFilter) isNodeName(metaName string) bool {
	return metaName == n.nodeName
}
