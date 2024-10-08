// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NutanixPrismEndpointApplyConfiguration represents a declarative configuration of the NutanixPrismEndpoint type for use
// with apply.
type NutanixPrismEndpointApplyConfiguration struct {
	Address *string `json:"address,omitempty"`
	Port    *int32  `json:"port,omitempty"`
}

// NutanixPrismEndpointApplyConfiguration constructs a declarative configuration of the NutanixPrismEndpoint type for use with
// apply.
func NutanixPrismEndpoint() *NutanixPrismEndpointApplyConfiguration {
	return &NutanixPrismEndpointApplyConfiguration{}
}

// WithAddress sets the Address field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Address field is set to the value of the last call.
func (b *NutanixPrismEndpointApplyConfiguration) WithAddress(value string) *NutanixPrismEndpointApplyConfiguration {
	b.Address = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *NutanixPrismEndpointApplyConfiguration) WithPort(value int32) *NutanixPrismEndpointApplyConfiguration {
	b.Port = &value
	return b
}
