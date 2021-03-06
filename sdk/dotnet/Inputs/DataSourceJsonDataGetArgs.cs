// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Grafana.Inputs
{

    public sealed class DataSourceJsonDataGetArgs : Pulumi.ResourceArgs
    {
        [Input("assumeRoleArn")]
        public Input<string>? AssumeRoleArn { get; set; }

        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        [Input("connMaxLifetime")]
        public Input<int>? ConnMaxLifetime { get; set; }

        [Input("customMetricsNamespaces")]
        public Input<string>? CustomMetricsNamespaces { get; set; }

        [Input("defaultProject")]
        public Input<string>? DefaultProject { get; set; }

        [Input("defaultRegion")]
        public Input<string>? DefaultRegion { get; set; }

        [Input("encrypt")]
        public Input<string>? Encrypt { get; set; }

        [Input("esVersion")]
        public Input<int>? EsVersion { get; set; }

        [Input("graphiteVersion")]
        public Input<string>? GraphiteVersion { get; set; }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("logLevelField")]
        public Input<string>? LogLevelField { get; set; }

        [Input("logMessageField")]
        public Input<string>? LogMessageField { get; set; }

        [Input("maxIdleConns")]
        public Input<int>? MaxIdleConns { get; set; }

        [Input("maxOpenConns")]
        public Input<int>? MaxOpenConns { get; set; }

        [Input("postgresVersion")]
        public Input<int>? PostgresVersion { get; set; }

        [Input("profile")]
        public Input<string>? Profile { get; set; }

        [Input("queryTimeout")]
        public Input<string>? QueryTimeout { get; set; }

        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        [Input("timeField")]
        public Input<string>? TimeField { get; set; }

        [Input("timeInterval")]
        public Input<string>? TimeInterval { get; set; }

        [Input("timescaledb")]
        public Input<bool>? Timescaledb { get; set; }

        [Input("tlsAuth")]
        public Input<bool>? TlsAuth { get; set; }

        [Input("tlsAuthWithCaCert")]
        public Input<bool>? TlsAuthWithCaCert { get; set; }

        [Input("tlsSkipVerify")]
        public Input<bool>? TlsSkipVerify { get; set; }

        [Input("tokenUri")]
        public Input<string>? TokenUri { get; set; }

        [Input("tsdbResolution")]
        public Input<string>? TsdbResolution { get; set; }

        [Input("tsdbVersion")]
        public Input<string>? TsdbVersion { get; set; }

        public DataSourceJsonDataGetArgs()
        {
        }
    }
}
