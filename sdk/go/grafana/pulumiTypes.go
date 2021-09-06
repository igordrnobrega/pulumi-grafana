// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BuiltinRoleAssignmentRole struct {
	Global *bool  `pulumi:"global"`
	Uid    string `pulumi:"uid"`
}

// BuiltinRoleAssignmentRoleInput is an input type that accepts BuiltinRoleAssignmentRoleArgs and BuiltinRoleAssignmentRoleOutput values.
// You can construct a concrete instance of `BuiltinRoleAssignmentRoleInput` via:
//
//          BuiltinRoleAssignmentRoleArgs{...}
type BuiltinRoleAssignmentRoleInput interface {
	pulumi.Input

	ToBuiltinRoleAssignmentRoleOutput() BuiltinRoleAssignmentRoleOutput
	ToBuiltinRoleAssignmentRoleOutputWithContext(context.Context) BuiltinRoleAssignmentRoleOutput
}

type BuiltinRoleAssignmentRoleArgs struct {
	Global pulumi.BoolPtrInput `pulumi:"global"`
	Uid    pulumi.StringInput  `pulumi:"uid"`
}

func (BuiltinRoleAssignmentRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltinRoleAssignmentRole)(nil)).Elem()
}

func (i BuiltinRoleAssignmentRoleArgs) ToBuiltinRoleAssignmentRoleOutput() BuiltinRoleAssignmentRoleOutput {
	return i.ToBuiltinRoleAssignmentRoleOutputWithContext(context.Background())
}

func (i BuiltinRoleAssignmentRoleArgs) ToBuiltinRoleAssignmentRoleOutputWithContext(ctx context.Context) BuiltinRoleAssignmentRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuiltinRoleAssignmentRoleOutput)
}

// BuiltinRoleAssignmentRoleArrayInput is an input type that accepts BuiltinRoleAssignmentRoleArray and BuiltinRoleAssignmentRoleArrayOutput values.
// You can construct a concrete instance of `BuiltinRoleAssignmentRoleArrayInput` via:
//
//          BuiltinRoleAssignmentRoleArray{ BuiltinRoleAssignmentRoleArgs{...} }
type BuiltinRoleAssignmentRoleArrayInput interface {
	pulumi.Input

	ToBuiltinRoleAssignmentRoleArrayOutput() BuiltinRoleAssignmentRoleArrayOutput
	ToBuiltinRoleAssignmentRoleArrayOutputWithContext(context.Context) BuiltinRoleAssignmentRoleArrayOutput
}

type BuiltinRoleAssignmentRoleArray []BuiltinRoleAssignmentRoleInput

func (BuiltinRoleAssignmentRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuiltinRoleAssignmentRole)(nil)).Elem()
}

func (i BuiltinRoleAssignmentRoleArray) ToBuiltinRoleAssignmentRoleArrayOutput() BuiltinRoleAssignmentRoleArrayOutput {
	return i.ToBuiltinRoleAssignmentRoleArrayOutputWithContext(context.Background())
}

func (i BuiltinRoleAssignmentRoleArray) ToBuiltinRoleAssignmentRoleArrayOutputWithContext(ctx context.Context) BuiltinRoleAssignmentRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuiltinRoleAssignmentRoleArrayOutput)
}

type BuiltinRoleAssignmentRoleOutput struct{ *pulumi.OutputState }

func (BuiltinRoleAssignmentRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltinRoleAssignmentRole)(nil)).Elem()
}

func (o BuiltinRoleAssignmentRoleOutput) ToBuiltinRoleAssignmentRoleOutput() BuiltinRoleAssignmentRoleOutput {
	return o
}

func (o BuiltinRoleAssignmentRoleOutput) ToBuiltinRoleAssignmentRoleOutputWithContext(ctx context.Context) BuiltinRoleAssignmentRoleOutput {
	return o
}

func (o BuiltinRoleAssignmentRoleOutput) Global() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BuiltinRoleAssignmentRole) *bool { return v.Global }).(pulumi.BoolPtrOutput)
}

func (o BuiltinRoleAssignmentRoleOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v BuiltinRoleAssignmentRole) string { return v.Uid }).(pulumi.StringOutput)
}

type BuiltinRoleAssignmentRoleArrayOutput struct{ *pulumi.OutputState }

func (BuiltinRoleAssignmentRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuiltinRoleAssignmentRole)(nil)).Elem()
}

func (o BuiltinRoleAssignmentRoleArrayOutput) ToBuiltinRoleAssignmentRoleArrayOutput() BuiltinRoleAssignmentRoleArrayOutput {
	return o
}

func (o BuiltinRoleAssignmentRoleArrayOutput) ToBuiltinRoleAssignmentRoleArrayOutputWithContext(ctx context.Context) BuiltinRoleAssignmentRoleArrayOutput {
	return o
}

func (o BuiltinRoleAssignmentRoleArrayOutput) Index(i pulumi.IntInput) BuiltinRoleAssignmentRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuiltinRoleAssignmentRole {
		return vs[0].([]BuiltinRoleAssignmentRole)[vs[1].(int)]
	}).(BuiltinRoleAssignmentRoleOutput)
}

type DashboardPermissionPermission struct {
	Permission string  `pulumi:"permission"`
	Role       *string `pulumi:"role"`
	TeamId     *int    `pulumi:"teamId"`
	UserId     *int    `pulumi:"userId"`
}

// DashboardPermissionPermissionInput is an input type that accepts DashboardPermissionPermissionArgs and DashboardPermissionPermissionOutput values.
// You can construct a concrete instance of `DashboardPermissionPermissionInput` via:
//
//          DashboardPermissionPermissionArgs{...}
type DashboardPermissionPermissionInput interface {
	pulumi.Input

	ToDashboardPermissionPermissionOutput() DashboardPermissionPermissionOutput
	ToDashboardPermissionPermissionOutputWithContext(context.Context) DashboardPermissionPermissionOutput
}

type DashboardPermissionPermissionArgs struct {
	Permission pulumi.StringInput    `pulumi:"permission"`
	Role       pulumi.StringPtrInput `pulumi:"role"`
	TeamId     pulumi.IntPtrInput    `pulumi:"teamId"`
	UserId     pulumi.IntPtrInput    `pulumi:"userId"`
}

func (DashboardPermissionPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPermissionPermission)(nil)).Elem()
}

func (i DashboardPermissionPermissionArgs) ToDashboardPermissionPermissionOutput() DashboardPermissionPermissionOutput {
	return i.ToDashboardPermissionPermissionOutputWithContext(context.Background())
}

func (i DashboardPermissionPermissionArgs) ToDashboardPermissionPermissionOutputWithContext(ctx context.Context) DashboardPermissionPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPermissionPermissionOutput)
}

// DashboardPermissionPermissionArrayInput is an input type that accepts DashboardPermissionPermissionArray and DashboardPermissionPermissionArrayOutput values.
// You can construct a concrete instance of `DashboardPermissionPermissionArrayInput` via:
//
//          DashboardPermissionPermissionArray{ DashboardPermissionPermissionArgs{...} }
type DashboardPermissionPermissionArrayInput interface {
	pulumi.Input

	ToDashboardPermissionPermissionArrayOutput() DashboardPermissionPermissionArrayOutput
	ToDashboardPermissionPermissionArrayOutputWithContext(context.Context) DashboardPermissionPermissionArrayOutput
}

type DashboardPermissionPermissionArray []DashboardPermissionPermissionInput

func (DashboardPermissionPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardPermissionPermission)(nil)).Elem()
}

func (i DashboardPermissionPermissionArray) ToDashboardPermissionPermissionArrayOutput() DashboardPermissionPermissionArrayOutput {
	return i.ToDashboardPermissionPermissionArrayOutputWithContext(context.Background())
}

func (i DashboardPermissionPermissionArray) ToDashboardPermissionPermissionArrayOutputWithContext(ctx context.Context) DashboardPermissionPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPermissionPermissionArrayOutput)
}

type DashboardPermissionPermissionOutput struct{ *pulumi.OutputState }

func (DashboardPermissionPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPermissionPermission)(nil)).Elem()
}

func (o DashboardPermissionPermissionOutput) ToDashboardPermissionPermissionOutput() DashboardPermissionPermissionOutput {
	return o
}

func (o DashboardPermissionPermissionOutput) ToDashboardPermissionPermissionOutputWithContext(ctx context.Context) DashboardPermissionPermissionOutput {
	return o
}

func (o DashboardPermissionPermissionOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v DashboardPermissionPermission) string { return v.Permission }).(pulumi.StringOutput)
}

func (o DashboardPermissionPermissionOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DashboardPermissionPermission) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o DashboardPermissionPermissionOutput) TeamId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DashboardPermissionPermission) *int { return v.TeamId }).(pulumi.IntPtrOutput)
}

func (o DashboardPermissionPermissionOutput) UserId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DashboardPermissionPermission) *int { return v.UserId }).(pulumi.IntPtrOutput)
}

type DashboardPermissionPermissionArrayOutput struct{ *pulumi.OutputState }

func (DashboardPermissionPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardPermissionPermission)(nil)).Elem()
}

func (o DashboardPermissionPermissionArrayOutput) ToDashboardPermissionPermissionArrayOutput() DashboardPermissionPermissionArrayOutput {
	return o
}

func (o DashboardPermissionPermissionArrayOutput) ToDashboardPermissionPermissionArrayOutputWithContext(ctx context.Context) DashboardPermissionPermissionArrayOutput {
	return o
}

func (o DashboardPermissionPermissionArrayOutput) Index(i pulumi.IntInput) DashboardPermissionPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DashboardPermissionPermission {
		return vs[0].([]DashboardPermissionPermission)[vs[1].(int)]
	}).(DashboardPermissionPermissionOutput)
}

type DataSourceJsonData struct {
	AssumeRoleArn           *string `pulumi:"assumeRoleArn"`
	AuthType                *string `pulumi:"authType"`
	AuthenticationType      *string `pulumi:"authenticationType"`
	ClientEmail             *string `pulumi:"clientEmail"`
	ConnMaxLifetime         *int    `pulumi:"connMaxLifetime"`
	CustomMetricsNamespaces *string `pulumi:"customMetricsNamespaces"`
	DefaultProject          *string `pulumi:"defaultProject"`
	DefaultRegion           *string `pulumi:"defaultRegion"`
	Encrypt                 *string `pulumi:"encrypt"`
	EsVersion               *int    `pulumi:"esVersion"`
	GraphiteVersion         *string `pulumi:"graphiteVersion"`
	HttpMethod              *string `pulumi:"httpMethod"`
	Interval                *string `pulumi:"interval"`
	LogLevelField           *string `pulumi:"logLevelField"`
	LogMessageField         *string `pulumi:"logMessageField"`
	MaxIdleConns            *int    `pulumi:"maxIdleConns"`
	MaxOpenConns            *int    `pulumi:"maxOpenConns"`
	PostgresVersion         *int    `pulumi:"postgresVersion"`
	Profile                 *string `pulumi:"profile"`
	QueryTimeout            *string `pulumi:"queryTimeout"`
	SslMode                 *string `pulumi:"sslMode"`
	TimeField               *string `pulumi:"timeField"`
	TimeInterval            *string `pulumi:"timeInterval"`
	Timescaledb             *bool   `pulumi:"timescaledb"`
	TlsAuth                 *bool   `pulumi:"tlsAuth"`
	TlsAuthWithCaCert       *bool   `pulumi:"tlsAuthWithCaCert"`
	TlsSkipVerify           *bool   `pulumi:"tlsSkipVerify"`
	TokenUri                *string `pulumi:"tokenUri"`
	TsdbResolution          *string `pulumi:"tsdbResolution"`
	TsdbVersion             *string `pulumi:"tsdbVersion"`
}

// DataSourceJsonDataInput is an input type that accepts DataSourceJsonDataArgs and DataSourceJsonDataOutput values.
// You can construct a concrete instance of `DataSourceJsonDataInput` via:
//
//          DataSourceJsonDataArgs{...}
type DataSourceJsonDataInput interface {
	pulumi.Input

	ToDataSourceJsonDataOutput() DataSourceJsonDataOutput
	ToDataSourceJsonDataOutputWithContext(context.Context) DataSourceJsonDataOutput
}

type DataSourceJsonDataArgs struct {
	AssumeRoleArn           pulumi.StringPtrInput `pulumi:"assumeRoleArn"`
	AuthType                pulumi.StringPtrInput `pulumi:"authType"`
	AuthenticationType      pulumi.StringPtrInput `pulumi:"authenticationType"`
	ClientEmail             pulumi.StringPtrInput `pulumi:"clientEmail"`
	ConnMaxLifetime         pulumi.IntPtrInput    `pulumi:"connMaxLifetime"`
	CustomMetricsNamespaces pulumi.StringPtrInput `pulumi:"customMetricsNamespaces"`
	DefaultProject          pulumi.StringPtrInput `pulumi:"defaultProject"`
	DefaultRegion           pulumi.StringPtrInput `pulumi:"defaultRegion"`
	Encrypt                 pulumi.StringPtrInput `pulumi:"encrypt"`
	EsVersion               pulumi.IntPtrInput    `pulumi:"esVersion"`
	GraphiteVersion         pulumi.StringPtrInput `pulumi:"graphiteVersion"`
	HttpMethod              pulumi.StringPtrInput `pulumi:"httpMethod"`
	Interval                pulumi.StringPtrInput `pulumi:"interval"`
	LogLevelField           pulumi.StringPtrInput `pulumi:"logLevelField"`
	LogMessageField         pulumi.StringPtrInput `pulumi:"logMessageField"`
	MaxIdleConns            pulumi.IntPtrInput    `pulumi:"maxIdleConns"`
	MaxOpenConns            pulumi.IntPtrInput    `pulumi:"maxOpenConns"`
	PostgresVersion         pulumi.IntPtrInput    `pulumi:"postgresVersion"`
	Profile                 pulumi.StringPtrInput `pulumi:"profile"`
	QueryTimeout            pulumi.StringPtrInput `pulumi:"queryTimeout"`
	SslMode                 pulumi.StringPtrInput `pulumi:"sslMode"`
	TimeField               pulumi.StringPtrInput `pulumi:"timeField"`
	TimeInterval            pulumi.StringPtrInput `pulumi:"timeInterval"`
	Timescaledb             pulumi.BoolPtrInput   `pulumi:"timescaledb"`
	TlsAuth                 pulumi.BoolPtrInput   `pulumi:"tlsAuth"`
	TlsAuthWithCaCert       pulumi.BoolPtrInput   `pulumi:"tlsAuthWithCaCert"`
	TlsSkipVerify           pulumi.BoolPtrInput   `pulumi:"tlsSkipVerify"`
	TokenUri                pulumi.StringPtrInput `pulumi:"tokenUri"`
	TsdbResolution          pulumi.StringPtrInput `pulumi:"tsdbResolution"`
	TsdbVersion             pulumi.StringPtrInput `pulumi:"tsdbVersion"`
}

func (DataSourceJsonDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceJsonData)(nil)).Elem()
}

func (i DataSourceJsonDataArgs) ToDataSourceJsonDataOutput() DataSourceJsonDataOutput {
	return i.ToDataSourceJsonDataOutputWithContext(context.Background())
}

func (i DataSourceJsonDataArgs) ToDataSourceJsonDataOutputWithContext(ctx context.Context) DataSourceJsonDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceJsonDataOutput)
}

// DataSourceJsonDataArrayInput is an input type that accepts DataSourceJsonDataArray and DataSourceJsonDataArrayOutput values.
// You can construct a concrete instance of `DataSourceJsonDataArrayInput` via:
//
//          DataSourceJsonDataArray{ DataSourceJsonDataArgs{...} }
type DataSourceJsonDataArrayInput interface {
	pulumi.Input

	ToDataSourceJsonDataArrayOutput() DataSourceJsonDataArrayOutput
	ToDataSourceJsonDataArrayOutputWithContext(context.Context) DataSourceJsonDataArrayOutput
}

type DataSourceJsonDataArray []DataSourceJsonDataInput

func (DataSourceJsonDataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourceJsonData)(nil)).Elem()
}

func (i DataSourceJsonDataArray) ToDataSourceJsonDataArrayOutput() DataSourceJsonDataArrayOutput {
	return i.ToDataSourceJsonDataArrayOutputWithContext(context.Background())
}

func (i DataSourceJsonDataArray) ToDataSourceJsonDataArrayOutputWithContext(ctx context.Context) DataSourceJsonDataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceJsonDataArrayOutput)
}

type DataSourceJsonDataOutput struct{ *pulumi.OutputState }

func (DataSourceJsonDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceJsonData)(nil)).Elem()
}

func (o DataSourceJsonDataOutput) ToDataSourceJsonDataOutput() DataSourceJsonDataOutput {
	return o
}

func (o DataSourceJsonDataOutput) ToDataSourceJsonDataOutputWithContext(ctx context.Context) DataSourceJsonDataOutput {
	return o
}

func (o DataSourceJsonDataOutput) AssumeRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.AssumeRoleArn }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) ClientEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.ClientEmail }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) ConnMaxLifetime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *int { return v.ConnMaxLifetime }).(pulumi.IntPtrOutput)
}

func (o DataSourceJsonDataOutput) CustomMetricsNamespaces() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.CustomMetricsNamespaces }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) DefaultProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.DefaultProject }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) DefaultRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.DefaultRegion }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) Encrypt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.Encrypt }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) EsVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *int { return v.EsVersion }).(pulumi.IntPtrOutput)
}

func (o DataSourceJsonDataOutput) GraphiteVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.GraphiteVersion }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) HttpMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.HttpMethod }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.Interval }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) LogLevelField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.LogLevelField }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) LogMessageField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.LogMessageField }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) MaxIdleConns() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *int { return v.MaxIdleConns }).(pulumi.IntPtrOutput)
}

func (o DataSourceJsonDataOutput) MaxOpenConns() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *int { return v.MaxOpenConns }).(pulumi.IntPtrOutput)
}

func (o DataSourceJsonDataOutput) PostgresVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *int { return v.PostgresVersion }).(pulumi.IntPtrOutput)
}

func (o DataSourceJsonDataOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) QueryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.QueryTimeout }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) SslMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.SslMode }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) TimeField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.TimeField }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) Timescaledb() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *bool { return v.Timescaledb }).(pulumi.BoolPtrOutput)
}

func (o DataSourceJsonDataOutput) TlsAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *bool { return v.TlsAuth }).(pulumi.BoolPtrOutput)
}

func (o DataSourceJsonDataOutput) TlsAuthWithCaCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *bool { return v.TlsAuthWithCaCert }).(pulumi.BoolPtrOutput)
}

func (o DataSourceJsonDataOutput) TlsSkipVerify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *bool { return v.TlsSkipVerify }).(pulumi.BoolPtrOutput)
}

func (o DataSourceJsonDataOutput) TokenUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.TokenUri }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) TsdbResolution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.TsdbResolution }).(pulumi.StringPtrOutput)
}

func (o DataSourceJsonDataOutput) TsdbVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceJsonData) *string { return v.TsdbVersion }).(pulumi.StringPtrOutput)
}

type DataSourceJsonDataArrayOutput struct{ *pulumi.OutputState }

func (DataSourceJsonDataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourceJsonData)(nil)).Elem()
}

func (o DataSourceJsonDataArrayOutput) ToDataSourceJsonDataArrayOutput() DataSourceJsonDataArrayOutput {
	return o
}

func (o DataSourceJsonDataArrayOutput) ToDataSourceJsonDataArrayOutputWithContext(ctx context.Context) DataSourceJsonDataArrayOutput {
	return o
}

func (o DataSourceJsonDataArrayOutput) Index(i pulumi.IntInput) DataSourceJsonDataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSourceJsonData {
		return vs[0].([]DataSourceJsonData)[vs[1].(int)]
	}).(DataSourceJsonDataOutput)
}

type DataSourceSecureJsonData struct {
	AccessKey         *string `pulumi:"accessKey"`
	BasicAuthPassword *string `pulumi:"basicAuthPassword"`
	Password          *string `pulumi:"password"`
	PrivateKey        *string `pulumi:"privateKey"`
	SecretKey         *string `pulumi:"secretKey"`
	TlsCaCert         *string `pulumi:"tlsCaCert"`
	TlsClientCert     *string `pulumi:"tlsClientCert"`
	TlsClientKey      *string `pulumi:"tlsClientKey"`
}

// DataSourceSecureJsonDataInput is an input type that accepts DataSourceSecureJsonDataArgs and DataSourceSecureJsonDataOutput values.
// You can construct a concrete instance of `DataSourceSecureJsonDataInput` via:
//
//          DataSourceSecureJsonDataArgs{...}
type DataSourceSecureJsonDataInput interface {
	pulumi.Input

	ToDataSourceSecureJsonDataOutput() DataSourceSecureJsonDataOutput
	ToDataSourceSecureJsonDataOutputWithContext(context.Context) DataSourceSecureJsonDataOutput
}

type DataSourceSecureJsonDataArgs struct {
	AccessKey         pulumi.StringPtrInput `pulumi:"accessKey"`
	BasicAuthPassword pulumi.StringPtrInput `pulumi:"basicAuthPassword"`
	Password          pulumi.StringPtrInput `pulumi:"password"`
	PrivateKey        pulumi.StringPtrInput `pulumi:"privateKey"`
	SecretKey         pulumi.StringPtrInput `pulumi:"secretKey"`
	TlsCaCert         pulumi.StringPtrInput `pulumi:"tlsCaCert"`
	TlsClientCert     pulumi.StringPtrInput `pulumi:"tlsClientCert"`
	TlsClientKey      pulumi.StringPtrInput `pulumi:"tlsClientKey"`
}

func (DataSourceSecureJsonDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceSecureJsonData)(nil)).Elem()
}

func (i DataSourceSecureJsonDataArgs) ToDataSourceSecureJsonDataOutput() DataSourceSecureJsonDataOutput {
	return i.ToDataSourceSecureJsonDataOutputWithContext(context.Background())
}

func (i DataSourceSecureJsonDataArgs) ToDataSourceSecureJsonDataOutputWithContext(ctx context.Context) DataSourceSecureJsonDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceSecureJsonDataOutput)
}

// DataSourceSecureJsonDataArrayInput is an input type that accepts DataSourceSecureJsonDataArray and DataSourceSecureJsonDataArrayOutput values.
// You can construct a concrete instance of `DataSourceSecureJsonDataArrayInput` via:
//
//          DataSourceSecureJsonDataArray{ DataSourceSecureJsonDataArgs{...} }
type DataSourceSecureJsonDataArrayInput interface {
	pulumi.Input

	ToDataSourceSecureJsonDataArrayOutput() DataSourceSecureJsonDataArrayOutput
	ToDataSourceSecureJsonDataArrayOutputWithContext(context.Context) DataSourceSecureJsonDataArrayOutput
}

type DataSourceSecureJsonDataArray []DataSourceSecureJsonDataInput

func (DataSourceSecureJsonDataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourceSecureJsonData)(nil)).Elem()
}

func (i DataSourceSecureJsonDataArray) ToDataSourceSecureJsonDataArrayOutput() DataSourceSecureJsonDataArrayOutput {
	return i.ToDataSourceSecureJsonDataArrayOutputWithContext(context.Background())
}

func (i DataSourceSecureJsonDataArray) ToDataSourceSecureJsonDataArrayOutputWithContext(ctx context.Context) DataSourceSecureJsonDataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceSecureJsonDataArrayOutput)
}

type DataSourceSecureJsonDataOutput struct{ *pulumi.OutputState }

func (DataSourceSecureJsonDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceSecureJsonData)(nil)).Elem()
}

func (o DataSourceSecureJsonDataOutput) ToDataSourceSecureJsonDataOutput() DataSourceSecureJsonDataOutput {
	return o
}

func (o DataSourceSecureJsonDataOutput) ToDataSourceSecureJsonDataOutputWithContext(ctx context.Context) DataSourceSecureJsonDataOutput {
	return o
}

func (o DataSourceSecureJsonDataOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o DataSourceSecureJsonDataOutput) BasicAuthPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.BasicAuthPassword }).(pulumi.StringPtrOutput)
}

func (o DataSourceSecureJsonDataOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o DataSourceSecureJsonDataOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o DataSourceSecureJsonDataOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.SecretKey }).(pulumi.StringPtrOutput)
}

func (o DataSourceSecureJsonDataOutput) TlsCaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.TlsCaCert }).(pulumi.StringPtrOutput)
}

func (o DataSourceSecureJsonDataOutput) TlsClientCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.TlsClientCert }).(pulumi.StringPtrOutput)
}

func (o DataSourceSecureJsonDataOutput) TlsClientKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataSourceSecureJsonData) *string { return v.TlsClientKey }).(pulumi.StringPtrOutput)
}

type DataSourceSecureJsonDataArrayOutput struct{ *pulumi.OutputState }

func (DataSourceSecureJsonDataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourceSecureJsonData)(nil)).Elem()
}

func (o DataSourceSecureJsonDataArrayOutput) ToDataSourceSecureJsonDataArrayOutput() DataSourceSecureJsonDataArrayOutput {
	return o
}

func (o DataSourceSecureJsonDataArrayOutput) ToDataSourceSecureJsonDataArrayOutputWithContext(ctx context.Context) DataSourceSecureJsonDataArrayOutput {
	return o
}

func (o DataSourceSecureJsonDataArrayOutput) Index(i pulumi.IntInput) DataSourceSecureJsonDataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSourceSecureJsonData {
		return vs[0].([]DataSourceSecureJsonData)[vs[1].(int)]
	}).(DataSourceSecureJsonDataOutput)
}

type FolderPermissionPermission struct {
	Permission string  `pulumi:"permission"`
	Role       *string `pulumi:"role"`
	TeamId     *int    `pulumi:"teamId"`
	UserId     *int    `pulumi:"userId"`
}

// FolderPermissionPermissionInput is an input type that accepts FolderPermissionPermissionArgs and FolderPermissionPermissionOutput values.
// You can construct a concrete instance of `FolderPermissionPermissionInput` via:
//
//          FolderPermissionPermissionArgs{...}
type FolderPermissionPermissionInput interface {
	pulumi.Input

	ToFolderPermissionPermissionOutput() FolderPermissionPermissionOutput
	ToFolderPermissionPermissionOutputWithContext(context.Context) FolderPermissionPermissionOutput
}

type FolderPermissionPermissionArgs struct {
	Permission pulumi.StringInput    `pulumi:"permission"`
	Role       pulumi.StringPtrInput `pulumi:"role"`
	TeamId     pulumi.IntPtrInput    `pulumi:"teamId"`
	UserId     pulumi.IntPtrInput    `pulumi:"userId"`
}

func (FolderPermissionPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderPermissionPermission)(nil)).Elem()
}

func (i FolderPermissionPermissionArgs) ToFolderPermissionPermissionOutput() FolderPermissionPermissionOutput {
	return i.ToFolderPermissionPermissionOutputWithContext(context.Background())
}

func (i FolderPermissionPermissionArgs) ToFolderPermissionPermissionOutputWithContext(ctx context.Context) FolderPermissionPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderPermissionPermissionOutput)
}

// FolderPermissionPermissionArrayInput is an input type that accepts FolderPermissionPermissionArray and FolderPermissionPermissionArrayOutput values.
// You can construct a concrete instance of `FolderPermissionPermissionArrayInput` via:
//
//          FolderPermissionPermissionArray{ FolderPermissionPermissionArgs{...} }
type FolderPermissionPermissionArrayInput interface {
	pulumi.Input

	ToFolderPermissionPermissionArrayOutput() FolderPermissionPermissionArrayOutput
	ToFolderPermissionPermissionArrayOutputWithContext(context.Context) FolderPermissionPermissionArrayOutput
}

type FolderPermissionPermissionArray []FolderPermissionPermissionInput

func (FolderPermissionPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FolderPermissionPermission)(nil)).Elem()
}

func (i FolderPermissionPermissionArray) ToFolderPermissionPermissionArrayOutput() FolderPermissionPermissionArrayOutput {
	return i.ToFolderPermissionPermissionArrayOutputWithContext(context.Background())
}

func (i FolderPermissionPermissionArray) ToFolderPermissionPermissionArrayOutputWithContext(ctx context.Context) FolderPermissionPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderPermissionPermissionArrayOutput)
}

type FolderPermissionPermissionOutput struct{ *pulumi.OutputState }

func (FolderPermissionPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderPermissionPermission)(nil)).Elem()
}

func (o FolderPermissionPermissionOutput) ToFolderPermissionPermissionOutput() FolderPermissionPermissionOutput {
	return o
}

func (o FolderPermissionPermissionOutput) ToFolderPermissionPermissionOutputWithContext(ctx context.Context) FolderPermissionPermissionOutput {
	return o
}

func (o FolderPermissionPermissionOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v FolderPermissionPermission) string { return v.Permission }).(pulumi.StringOutput)
}

func (o FolderPermissionPermissionOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FolderPermissionPermission) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o FolderPermissionPermissionOutput) TeamId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FolderPermissionPermission) *int { return v.TeamId }).(pulumi.IntPtrOutput)
}

func (o FolderPermissionPermissionOutput) UserId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FolderPermissionPermission) *int { return v.UserId }).(pulumi.IntPtrOutput)
}

type FolderPermissionPermissionArrayOutput struct{ *pulumi.OutputState }

func (FolderPermissionPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FolderPermissionPermission)(nil)).Elem()
}

func (o FolderPermissionPermissionArrayOutput) ToFolderPermissionPermissionArrayOutput() FolderPermissionPermissionArrayOutput {
	return o
}

func (o FolderPermissionPermissionArrayOutput) ToFolderPermissionPermissionArrayOutputWithContext(ctx context.Context) FolderPermissionPermissionArrayOutput {
	return o
}

func (o FolderPermissionPermissionArrayOutput) Index(i pulumi.IntInput) FolderPermissionPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FolderPermissionPermission {
		return vs[0].([]FolderPermissionPermission)[vs[1].(int)]
	}).(FolderPermissionPermissionOutput)
}

type RolePermission struct {
	Action string  `pulumi:"action"`
	Scope  *string `pulumi:"scope"`
}

// RolePermissionInput is an input type that accepts RolePermissionArgs and RolePermissionOutput values.
// You can construct a concrete instance of `RolePermissionInput` via:
//
//          RolePermissionArgs{...}
type RolePermissionInput interface {
	pulumi.Input

	ToRolePermissionOutput() RolePermissionOutput
	ToRolePermissionOutputWithContext(context.Context) RolePermissionOutput
}

type RolePermissionArgs struct {
	Action pulumi.StringInput    `pulumi:"action"`
	Scope  pulumi.StringPtrInput `pulumi:"scope"`
}

func (RolePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RolePermission)(nil)).Elem()
}

func (i RolePermissionArgs) ToRolePermissionOutput() RolePermissionOutput {
	return i.ToRolePermissionOutputWithContext(context.Background())
}

func (i RolePermissionArgs) ToRolePermissionOutputWithContext(ctx context.Context) RolePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePermissionOutput)
}

// RolePermissionArrayInput is an input type that accepts RolePermissionArray and RolePermissionArrayOutput values.
// You can construct a concrete instance of `RolePermissionArrayInput` via:
//
//          RolePermissionArray{ RolePermissionArgs{...} }
type RolePermissionArrayInput interface {
	pulumi.Input

	ToRolePermissionArrayOutput() RolePermissionArrayOutput
	ToRolePermissionArrayOutputWithContext(context.Context) RolePermissionArrayOutput
}

type RolePermissionArray []RolePermissionInput

func (RolePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RolePermission)(nil)).Elem()
}

func (i RolePermissionArray) ToRolePermissionArrayOutput() RolePermissionArrayOutput {
	return i.ToRolePermissionArrayOutputWithContext(context.Background())
}

func (i RolePermissionArray) ToRolePermissionArrayOutputWithContext(ctx context.Context) RolePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePermissionArrayOutput)
}

type RolePermissionOutput struct{ *pulumi.OutputState }

func (RolePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RolePermission)(nil)).Elem()
}

func (o RolePermissionOutput) ToRolePermissionOutput() RolePermissionOutput {
	return o
}

func (o RolePermissionOutput) ToRolePermissionOutputWithContext(ctx context.Context) RolePermissionOutput {
	return o
}

func (o RolePermissionOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v RolePermission) string { return v.Action }).(pulumi.StringOutput)
}

func (o RolePermissionOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RolePermission) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type RolePermissionArrayOutput struct{ *pulumi.OutputState }

func (RolePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RolePermission)(nil)).Elem()
}

func (o RolePermissionArrayOutput) ToRolePermissionArrayOutput() RolePermissionArrayOutput {
	return o
}

func (o RolePermissionArrayOutput) ToRolePermissionArrayOutputWithContext(ctx context.Context) RolePermissionArrayOutput {
	return o
}

func (o RolePermissionArrayOutput) Index(i pulumi.IntInput) RolePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RolePermission {
		return vs[0].([]RolePermission)[vs[1].(int)]
	}).(RolePermissionOutput)
}

func init() {
	pulumi.RegisterOutputType(BuiltinRoleAssignmentRoleOutput{})
	pulumi.RegisterOutputType(BuiltinRoleAssignmentRoleArrayOutput{})
	pulumi.RegisterOutputType(DashboardPermissionPermissionOutput{})
	pulumi.RegisterOutputType(DashboardPermissionPermissionArrayOutput{})
	pulumi.RegisterOutputType(DataSourceJsonDataOutput{})
	pulumi.RegisterOutputType(DataSourceJsonDataArrayOutput{})
	pulumi.RegisterOutputType(DataSourceSecureJsonDataOutput{})
	pulumi.RegisterOutputType(DataSourceSecureJsonDataArrayOutput{})
	pulumi.RegisterOutputType(FolderPermissionPermissionOutput{})
	pulumi.RegisterOutputType(FolderPermissionPermissionArrayOutput{})
	pulumi.RegisterOutputType(RolePermissionOutput{})
	pulumi.RegisterOutputType(RolePermissionArrayOutput{})
}
