// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Folder struct {
	pulumi.CustomResourceState

	// The title of the folder.
	Title pulumi.StringOutput `pulumi:"title"`
	// Unique identifier.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOption) (*Folder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource Folder
	err := ctx.RegisterResource("grafana:index/folder:Folder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderState, opts ...pulumi.ResourceOption) (*Folder, error) {
	var resource Folder
	err := ctx.ReadResource("grafana:index/folder:Folder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Folder resources.
type folderState struct {
	// The title of the folder.
	Title *string `pulumi:"title"`
	// Unique identifier.
	Uid *string `pulumi:"uid"`
}

type FolderState struct {
	// The title of the folder.
	Title pulumi.StringPtrInput
	// Unique identifier.
	Uid pulumi.StringPtrInput
}

func (FolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderState)(nil)).Elem()
}

type folderArgs struct {
	// The title of the folder.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// The title of the folder.
	Title pulumi.StringInput
}

func (FolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}

type FolderInput interface {
	pulumi.Input

	ToFolderOutput() FolderOutput
	ToFolderOutputWithContext(ctx context.Context) FolderOutput
}

func (*Folder) ElementType() reflect.Type {
	return reflect.TypeOf((*Folder)(nil))
}

func (i *Folder) ToFolderOutput() FolderOutput {
	return i.ToFolderOutputWithContext(context.Background())
}

func (i *Folder) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderOutput)
}

func (i *Folder) ToFolderPtrOutput() FolderPtrOutput {
	return i.ToFolderPtrOutputWithContext(context.Background())
}

func (i *Folder) ToFolderPtrOutputWithContext(ctx context.Context) FolderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderPtrOutput)
}

type FolderPtrInput interface {
	pulumi.Input

	ToFolderPtrOutput() FolderPtrOutput
	ToFolderPtrOutputWithContext(ctx context.Context) FolderPtrOutput
}

type folderPtrType FolderArgs

func (*folderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil))
}

func (i *folderPtrType) ToFolderPtrOutput() FolderPtrOutput {
	return i.ToFolderPtrOutputWithContext(context.Background())
}

func (i *folderPtrType) ToFolderPtrOutputWithContext(ctx context.Context) FolderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderPtrOutput)
}

// FolderArrayInput is an input type that accepts FolderArray and FolderArrayOutput values.
// You can construct a concrete instance of `FolderArrayInput` via:
//
//          FolderArray{ FolderArgs{...} }
type FolderArrayInput interface {
	pulumi.Input

	ToFolderArrayOutput() FolderArrayOutput
	ToFolderArrayOutputWithContext(context.Context) FolderArrayOutput
}

type FolderArray []FolderInput

func (FolderArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Folder)(nil))
}

func (i FolderArray) ToFolderArrayOutput() FolderArrayOutput {
	return i.ToFolderArrayOutputWithContext(context.Background())
}

func (i FolderArray) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderArrayOutput)
}

// FolderMapInput is an input type that accepts FolderMap and FolderMapOutput values.
// You can construct a concrete instance of `FolderMapInput` via:
//
//          FolderMap{ "key": FolderArgs{...} }
type FolderMapInput interface {
	pulumi.Input

	ToFolderMapOutput() FolderMapOutput
	ToFolderMapOutputWithContext(context.Context) FolderMapOutput
}

type FolderMap map[string]FolderInput

func (FolderMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Folder)(nil))
}

func (i FolderMap) ToFolderMapOutput() FolderMapOutput {
	return i.ToFolderMapOutputWithContext(context.Background())
}

func (i FolderMap) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderMapOutput)
}

type FolderOutput struct {
	*pulumi.OutputState
}

func (FolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Folder)(nil))
}

func (o FolderOutput) ToFolderOutput() FolderOutput {
	return o
}

func (o FolderOutput) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return o
}

func (o FolderOutput) ToFolderPtrOutput() FolderPtrOutput {
	return o.ToFolderPtrOutputWithContext(context.Background())
}

func (o FolderOutput) ToFolderPtrOutputWithContext(ctx context.Context) FolderPtrOutput {
	return o.ApplyT(func(v Folder) *Folder {
		return &v
	}).(FolderPtrOutput)
}

type FolderPtrOutput struct {
	*pulumi.OutputState
}

func (FolderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil))
}

func (o FolderPtrOutput) ToFolderPtrOutput() FolderPtrOutput {
	return o
}

func (o FolderPtrOutput) ToFolderPtrOutputWithContext(ctx context.Context) FolderPtrOutput {
	return o
}

type FolderArrayOutput struct{ *pulumi.OutputState }

func (FolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Folder)(nil))
}

func (o FolderArrayOutput) ToFolderArrayOutput() FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) Index(i pulumi.IntInput) FolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Folder {
		return vs[0].([]Folder)[vs[1].(int)]
	}).(FolderOutput)
}

type FolderMapOutput struct{ *pulumi.OutputState }

func (FolderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Folder)(nil))
}

func (o FolderMapOutput) ToFolderMapOutput() FolderMapOutput {
	return o
}

func (o FolderMapOutput) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return o
}

func (o FolderMapOutput) MapIndex(k pulumi.StringInput) FolderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Folder {
		return vs[0].(map[string]Folder)[vs[1].(string)]
	}).(FolderOutput)
}

func init() {
	pulumi.RegisterOutputType(FolderOutput{})
	pulumi.RegisterOutputType(FolderPtrOutput{})
	pulumi.RegisterOutputType(FolderArrayOutput{})
	pulumi.RegisterOutputType(FolderMapOutput{})
}