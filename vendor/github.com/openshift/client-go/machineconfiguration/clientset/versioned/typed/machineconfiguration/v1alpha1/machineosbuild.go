// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/openshift/api/machineconfiguration/v1alpha1"
	machineconfigurationv1alpha1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1alpha1"
	scheme "github.com/openshift/client-go/machineconfiguration/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// MachineOSBuildsGetter has a method to return a MachineOSBuildInterface.
// A group's client should implement this interface.
type MachineOSBuildsGetter interface {
	MachineOSBuilds() MachineOSBuildInterface
}

// MachineOSBuildInterface has methods to work with MachineOSBuild resources.
type MachineOSBuildInterface interface {
	Create(ctx context.Context, machineOSBuild *v1alpha1.MachineOSBuild, opts v1.CreateOptions) (*v1alpha1.MachineOSBuild, error)
	Update(ctx context.Context, machineOSBuild *v1alpha1.MachineOSBuild, opts v1.UpdateOptions) (*v1alpha1.MachineOSBuild, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, machineOSBuild *v1alpha1.MachineOSBuild, opts v1.UpdateOptions) (*v1alpha1.MachineOSBuild, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MachineOSBuild, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MachineOSBuildList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineOSBuild, err error)
	Apply(ctx context.Context, machineOSBuild *machineconfigurationv1alpha1.MachineOSBuildApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MachineOSBuild, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, machineOSBuild *machineconfigurationv1alpha1.MachineOSBuildApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MachineOSBuild, err error)
	MachineOSBuildExpansion
}

// machineOSBuilds implements MachineOSBuildInterface
type machineOSBuilds struct {
	*gentype.ClientWithListAndApply[*v1alpha1.MachineOSBuild, *v1alpha1.MachineOSBuildList, *machineconfigurationv1alpha1.MachineOSBuildApplyConfiguration]
}

// newMachineOSBuilds returns a MachineOSBuilds
func newMachineOSBuilds(c *MachineconfigurationV1alpha1Client) *machineOSBuilds {
	return &machineOSBuilds{
		gentype.NewClientWithListAndApply[*v1alpha1.MachineOSBuild, *v1alpha1.MachineOSBuildList, *machineconfigurationv1alpha1.MachineOSBuildApplyConfiguration](
			"machineosbuilds",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.MachineOSBuild { return &v1alpha1.MachineOSBuild{} },
			func() *v1alpha1.MachineOSBuildList { return &v1alpha1.MachineOSBuildList{} }),
	}
}
