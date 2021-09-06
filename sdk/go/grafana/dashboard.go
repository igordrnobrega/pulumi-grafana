// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Dashboard struct {
	pulumi.CustomResourceState

	// The complete dashboard model JSON.
	ConfigJson pulumi.StringOutput `pulumi:"configJson"`
	// The numeric ID of the dashboard computed by Grafana.
	DashboardId pulumi.IntOutput `pulumi:"dashboardId"`
	// The id of the folder to save the dashboard in.
	Folder pulumi.IntPtrOutput `pulumi:"folder"`
	// URL friendly version of the dashboard title.
	Slug pulumi.StringOutput `pulumi:"slug"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigJson == nil {
		return nil, errors.New("invalid value for required argument 'ConfigJson'")
	}
	var resource Dashboard
	err := ctx.RegisterResource("grafana:index/dashboard:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("grafana:index/dashboard:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
	// The complete dashboard model JSON.
	ConfigJson *string `pulumi:"configJson"`
	// The numeric ID of the dashboard computed by Grafana.
	DashboardId *int `pulumi:"dashboardId"`
	// The id of the folder to save the dashboard in.
	Folder *int `pulumi:"folder"`
	// URL friendly version of the dashboard title.
	Slug *string `pulumi:"slug"`
}

type DashboardState struct {
	// The complete dashboard model JSON.
	ConfigJson pulumi.StringPtrInput
	// The numeric ID of the dashboard computed by Grafana.
	DashboardId pulumi.IntPtrInput
	// The id of the folder to save the dashboard in.
	Folder pulumi.IntPtrInput
	// URL friendly version of the dashboard title.
	Slug pulumi.StringPtrInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// The complete dashboard model JSON.
	ConfigJson string `pulumi:"configJson"`
	// The id of the folder to save the dashboard in.
	Folder *int `pulumi:"folder"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// The complete dashboard model JSON.
	ConfigJson pulumi.StringInput
	// The id of the folder to save the dashboard in.
	Folder pulumi.IntPtrInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((*Dashboard)(nil))
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

func (i *Dashboard) ToDashboardPtrOutput() DashboardPtrOutput {
	return i.ToDashboardPtrOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPtrOutput)
}

type DashboardPtrInput interface {
	pulumi.Input

	ToDashboardPtrOutput() DashboardPtrOutput
	ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput
}

type dashboardPtrType DashboardArgs

func (*dashboardPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil))
}

func (i *dashboardPtrType) ToDashboardPtrOutput() DashboardPtrOutput {
	return i.ToDashboardPtrOutputWithContext(context.Background())
}

func (i *dashboardPtrType) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPtrOutput)
}

// DashboardArrayInput is an input type that accepts DashboardArray and DashboardArrayOutput values.
// You can construct a concrete instance of `DashboardArrayInput` via:
//
//          DashboardArray{ DashboardArgs{...} }
type DashboardArrayInput interface {
	pulumi.Input

	ToDashboardArrayOutput() DashboardArrayOutput
	ToDashboardArrayOutputWithContext(context.Context) DashboardArrayOutput
}

type DashboardArray []DashboardInput

func (DashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Dashboard)(nil))
}

func (i DashboardArray) ToDashboardArrayOutput() DashboardArrayOutput {
	return i.ToDashboardArrayOutputWithContext(context.Background())
}

func (i DashboardArray) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardArrayOutput)
}

// DashboardMapInput is an input type that accepts DashboardMap and DashboardMapOutput values.
// You can construct a concrete instance of `DashboardMapInput` via:
//
//          DashboardMap{ "key": DashboardArgs{...} }
type DashboardMapInput interface {
	pulumi.Input

	ToDashboardMapOutput() DashboardMapOutput
	ToDashboardMapOutputWithContext(context.Context) DashboardMapOutput
}

type DashboardMap map[string]DashboardInput

func (DashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Dashboard)(nil))
}

func (i DashboardMap) ToDashboardMapOutput() DashboardMapOutput {
	return i.ToDashboardMapOutputWithContext(context.Background())
}

func (i DashboardMap) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardMapOutput)
}

type DashboardOutput struct {
	*pulumi.OutputState
}

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dashboard)(nil))
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardPtrOutput() DashboardPtrOutput {
	return o.ToDashboardPtrOutputWithContext(context.Background())
}

func (o DashboardOutput) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return o.ApplyT(func(v Dashboard) *Dashboard {
		return &v
	}).(DashboardPtrOutput)
}

type DashboardPtrOutput struct {
	*pulumi.OutputState
}

func (DashboardPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil))
}

func (o DashboardPtrOutput) ToDashboardPtrOutput() DashboardPtrOutput {
	return o
}

func (o DashboardPtrOutput) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return o
}

type DashboardArrayOutput struct{ *pulumi.OutputState }

func (DashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dashboard)(nil))
}

func (o DashboardArrayOutput) ToDashboardArrayOutput() DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) Index(i pulumi.IntInput) DashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Dashboard {
		return vs[0].([]Dashboard)[vs[1].(int)]
	}).(DashboardOutput)
}

type DashboardMapOutput struct{ *pulumi.OutputState }

func (DashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Dashboard)(nil))
}

func (o DashboardMapOutput) ToDashboardMapOutput() DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) MapIndex(k pulumi.StringInput) DashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Dashboard {
		return vs[0].(map[string]Dashboard)[vs[1].(string)]
	}).(DashboardOutput)
}

func init() {
	pulumi.RegisterOutputType(DashboardOutput{})
	pulumi.RegisterOutputType(DashboardPtrOutput{})
	pulumi.RegisterOutputType(DashboardArrayOutput{})
	pulumi.RegisterOutputType(DashboardMapOutput{})
}