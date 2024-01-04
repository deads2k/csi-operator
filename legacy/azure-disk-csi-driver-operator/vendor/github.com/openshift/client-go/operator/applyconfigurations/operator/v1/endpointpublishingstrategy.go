// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// EndpointPublishingStrategyApplyConfiguration represents an declarative configuration of the EndpointPublishingStrategy type for use
// with apply.
type EndpointPublishingStrategyApplyConfiguration struct {
	Type         *v1.EndpointPublishingStrategyType      `json:"type,omitempty"`
	LoadBalancer *LoadBalancerStrategyApplyConfiguration `json:"loadBalancer,omitempty"`
	HostNetwork  *HostNetworkStrategyApplyConfiguration  `json:"hostNetwork,omitempty"`
	Private      *PrivateStrategyApplyConfiguration      `json:"private,omitempty"`
	NodePort     *NodePortStrategyApplyConfiguration     `json:"nodePort,omitempty"`
}

// EndpointPublishingStrategyApplyConfiguration constructs an declarative configuration of the EndpointPublishingStrategy type for use with
// apply.
func EndpointPublishingStrategy() *EndpointPublishingStrategyApplyConfiguration {
	return &EndpointPublishingStrategyApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *EndpointPublishingStrategyApplyConfiguration) WithType(value v1.EndpointPublishingStrategyType) *EndpointPublishingStrategyApplyConfiguration {
	b.Type = &value
	return b
}

// WithLoadBalancer sets the LoadBalancer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LoadBalancer field is set to the value of the last call.
func (b *EndpointPublishingStrategyApplyConfiguration) WithLoadBalancer(value *LoadBalancerStrategyApplyConfiguration) *EndpointPublishingStrategyApplyConfiguration {
	b.LoadBalancer = value
	return b
}

// WithHostNetwork sets the HostNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostNetwork field is set to the value of the last call.
func (b *EndpointPublishingStrategyApplyConfiguration) WithHostNetwork(value *HostNetworkStrategyApplyConfiguration) *EndpointPublishingStrategyApplyConfiguration {
	b.HostNetwork = value
	return b
}

// WithPrivate sets the Private field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Private field is set to the value of the last call.
func (b *EndpointPublishingStrategyApplyConfiguration) WithPrivate(value *PrivateStrategyApplyConfiguration) *EndpointPublishingStrategyApplyConfiguration {
	b.Private = value
	return b
}

// WithNodePort sets the NodePort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePort field is set to the value of the last call.
func (b *EndpointPublishingStrategyApplyConfiguration) WithNodePort(value *NodePortStrategyApplyConfiguration) *EndpointPublishingStrategyApplyConfiguration {
	b.NodePort = value
	return b
}
