// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TeamPreferences struct {
	pulumi.CustomResourceState

	// The numeric ID of the dashboard to display when a team member logs in.
	HomeDashboardId pulumi.IntPtrOutput `pulumi:"homeDashboardId"`
	// The numeric team ID.
	TeamId pulumi.IntOutput `pulumi:"teamId"`
	// The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
	Theme pulumi.StringPtrOutput `pulumi:"theme"`
	// The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewTeamPreferences registers a new resource with the given unique name, arguments, and options.
func NewTeamPreferences(ctx *pulumi.Context,
	name string, args *TeamPreferencesArgs, opts ...pulumi.ResourceOption) (*TeamPreferences, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	var resource TeamPreferences
	err := ctx.RegisterResource("grafana:index/teamPreferences:TeamPreferences", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamPreferences gets an existing TeamPreferences resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamPreferences(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamPreferencesState, opts ...pulumi.ResourceOption) (*TeamPreferences, error) {
	var resource TeamPreferences
	err := ctx.ReadResource("grafana:index/teamPreferences:TeamPreferences", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamPreferences resources.
type teamPreferencesState struct {
	// The numeric ID of the dashboard to display when a team member logs in.
	HomeDashboardId *int `pulumi:"homeDashboardId"`
	// The numeric team ID.
	TeamId *int `pulumi:"teamId"`
	// The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
	Theme *string `pulumi:"theme"`
	// The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
	Timezone *string `pulumi:"timezone"`
}

type TeamPreferencesState struct {
	// The numeric ID of the dashboard to display when a team member logs in.
	HomeDashboardId pulumi.IntPtrInput
	// The numeric team ID.
	TeamId pulumi.IntPtrInput
	// The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
	Theme pulumi.StringPtrInput
	// The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
	Timezone pulumi.StringPtrInput
}

func (TeamPreferencesState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamPreferencesState)(nil)).Elem()
}

type teamPreferencesArgs struct {
	// The numeric ID of the dashboard to display when a team member logs in.
	HomeDashboardId *int `pulumi:"homeDashboardId"`
	// The numeric team ID.
	TeamId int `pulumi:"teamId"`
	// The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
	Theme *string `pulumi:"theme"`
	// The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a TeamPreferences resource.
type TeamPreferencesArgs struct {
	// The numeric ID of the dashboard to display when a team member logs in.
	HomeDashboardId pulumi.IntPtrInput
	// The numeric team ID.
	TeamId pulumi.IntInput
	// The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
	Theme pulumi.StringPtrInput
	// The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
	Timezone pulumi.StringPtrInput
}

func (TeamPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamPreferencesArgs)(nil)).Elem()
}

type TeamPreferencesInput interface {
	pulumi.Input

	ToTeamPreferencesOutput() TeamPreferencesOutput
	ToTeamPreferencesOutputWithContext(ctx context.Context) TeamPreferencesOutput
}

func (*TeamPreferences) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamPreferences)(nil))
}

func (i *TeamPreferences) ToTeamPreferencesOutput() TeamPreferencesOutput {
	return i.ToTeamPreferencesOutputWithContext(context.Background())
}

func (i *TeamPreferences) ToTeamPreferencesOutputWithContext(ctx context.Context) TeamPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamPreferencesOutput)
}

func (i *TeamPreferences) ToTeamPreferencesPtrOutput() TeamPreferencesPtrOutput {
	return i.ToTeamPreferencesPtrOutputWithContext(context.Background())
}

func (i *TeamPreferences) ToTeamPreferencesPtrOutputWithContext(ctx context.Context) TeamPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamPreferencesPtrOutput)
}

type TeamPreferencesPtrInput interface {
	pulumi.Input

	ToTeamPreferencesPtrOutput() TeamPreferencesPtrOutput
	ToTeamPreferencesPtrOutputWithContext(ctx context.Context) TeamPreferencesPtrOutput
}

type teamPreferencesPtrType TeamPreferencesArgs

func (*teamPreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamPreferences)(nil))
}

func (i *teamPreferencesPtrType) ToTeamPreferencesPtrOutput() TeamPreferencesPtrOutput {
	return i.ToTeamPreferencesPtrOutputWithContext(context.Background())
}

func (i *teamPreferencesPtrType) ToTeamPreferencesPtrOutputWithContext(ctx context.Context) TeamPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamPreferencesPtrOutput)
}

// TeamPreferencesArrayInput is an input type that accepts TeamPreferencesArray and TeamPreferencesArrayOutput values.
// You can construct a concrete instance of `TeamPreferencesArrayInput` via:
//
//          TeamPreferencesArray{ TeamPreferencesArgs{...} }
type TeamPreferencesArrayInput interface {
	pulumi.Input

	ToTeamPreferencesArrayOutput() TeamPreferencesArrayOutput
	ToTeamPreferencesArrayOutputWithContext(context.Context) TeamPreferencesArrayOutput
}

type TeamPreferencesArray []TeamPreferencesInput

func (TeamPreferencesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TeamPreferences)(nil))
}

func (i TeamPreferencesArray) ToTeamPreferencesArrayOutput() TeamPreferencesArrayOutput {
	return i.ToTeamPreferencesArrayOutputWithContext(context.Background())
}

func (i TeamPreferencesArray) ToTeamPreferencesArrayOutputWithContext(ctx context.Context) TeamPreferencesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamPreferencesArrayOutput)
}

// TeamPreferencesMapInput is an input type that accepts TeamPreferencesMap and TeamPreferencesMapOutput values.
// You can construct a concrete instance of `TeamPreferencesMapInput` via:
//
//          TeamPreferencesMap{ "key": TeamPreferencesArgs{...} }
type TeamPreferencesMapInput interface {
	pulumi.Input

	ToTeamPreferencesMapOutput() TeamPreferencesMapOutput
	ToTeamPreferencesMapOutputWithContext(context.Context) TeamPreferencesMapOutput
}

type TeamPreferencesMap map[string]TeamPreferencesInput

func (TeamPreferencesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TeamPreferences)(nil))
}

func (i TeamPreferencesMap) ToTeamPreferencesMapOutput() TeamPreferencesMapOutput {
	return i.ToTeamPreferencesMapOutputWithContext(context.Background())
}

func (i TeamPreferencesMap) ToTeamPreferencesMapOutputWithContext(ctx context.Context) TeamPreferencesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamPreferencesMapOutput)
}

type TeamPreferencesOutput struct {
	*pulumi.OutputState
}

func (TeamPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamPreferences)(nil))
}

func (o TeamPreferencesOutput) ToTeamPreferencesOutput() TeamPreferencesOutput {
	return o
}

func (o TeamPreferencesOutput) ToTeamPreferencesOutputWithContext(ctx context.Context) TeamPreferencesOutput {
	return o
}

func (o TeamPreferencesOutput) ToTeamPreferencesPtrOutput() TeamPreferencesPtrOutput {
	return o.ToTeamPreferencesPtrOutputWithContext(context.Background())
}

func (o TeamPreferencesOutput) ToTeamPreferencesPtrOutputWithContext(ctx context.Context) TeamPreferencesPtrOutput {
	return o.ApplyT(func(v TeamPreferences) *TeamPreferences {
		return &v
	}).(TeamPreferencesPtrOutput)
}

type TeamPreferencesPtrOutput struct {
	*pulumi.OutputState
}

func (TeamPreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamPreferences)(nil))
}

func (o TeamPreferencesPtrOutput) ToTeamPreferencesPtrOutput() TeamPreferencesPtrOutput {
	return o
}

func (o TeamPreferencesPtrOutput) ToTeamPreferencesPtrOutputWithContext(ctx context.Context) TeamPreferencesPtrOutput {
	return o
}

type TeamPreferencesArrayOutput struct{ *pulumi.OutputState }

func (TeamPreferencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TeamPreferences)(nil))
}

func (o TeamPreferencesArrayOutput) ToTeamPreferencesArrayOutput() TeamPreferencesArrayOutput {
	return o
}

func (o TeamPreferencesArrayOutput) ToTeamPreferencesArrayOutputWithContext(ctx context.Context) TeamPreferencesArrayOutput {
	return o
}

func (o TeamPreferencesArrayOutput) Index(i pulumi.IntInput) TeamPreferencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TeamPreferences {
		return vs[0].([]TeamPreferences)[vs[1].(int)]
	}).(TeamPreferencesOutput)
}

type TeamPreferencesMapOutput struct{ *pulumi.OutputState }

func (TeamPreferencesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TeamPreferences)(nil))
}

func (o TeamPreferencesMapOutput) ToTeamPreferencesMapOutput() TeamPreferencesMapOutput {
	return o
}

func (o TeamPreferencesMapOutput) ToTeamPreferencesMapOutputWithContext(ctx context.Context) TeamPreferencesMapOutput {
	return o
}

func (o TeamPreferencesMapOutput) MapIndex(k pulumi.StringInput) TeamPreferencesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TeamPreferences {
		return vs[0].(map[string]TeamPreferences)[vs[1].(string)]
	}).(TeamPreferencesOutput)
}

func init() {
	pulumi.RegisterOutputType(TeamPreferencesOutput{})
	pulumi.RegisterOutputType(TeamPreferencesPtrOutput{})
	pulumi.RegisterOutputType(TeamPreferencesArrayOutput{})
	pulumi.RegisterOutputType(TeamPreferencesMapOutput{})
}
