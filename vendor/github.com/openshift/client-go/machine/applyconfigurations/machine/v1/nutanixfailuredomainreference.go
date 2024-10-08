// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NutanixFailureDomainReferenceApplyConfiguration represents a declarative configuration of the NutanixFailureDomainReference type for use
// with apply.
type NutanixFailureDomainReferenceApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// NutanixFailureDomainReferenceApplyConfiguration constructs a declarative configuration of the NutanixFailureDomainReference type for use with
// apply.
func NutanixFailureDomainReference() *NutanixFailureDomainReferenceApplyConfiguration {
	return &NutanixFailureDomainReferenceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NutanixFailureDomainReferenceApplyConfiguration) WithName(value string) *NutanixFailureDomainReferenceApplyConfiguration {
	b.Name = &value
	return b
}
