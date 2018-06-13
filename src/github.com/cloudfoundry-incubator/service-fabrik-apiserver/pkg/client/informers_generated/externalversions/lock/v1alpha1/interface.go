//TODO copyright header

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DeploymentLocks returns a DeploymentLockInformer.
	DeploymentLocks() DeploymentLockInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// DeploymentLocks returns a DeploymentLockInformer.
func (v *version) DeploymentLocks() DeploymentLockInformer {
	return &deploymentLockInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
