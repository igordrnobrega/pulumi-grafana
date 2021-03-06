// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("grafana");

/**
 * API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
 */
export let auth: string | undefined = __config.get("auth");
/**
 * Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
 * `GRAFANA_CA_CERT` environment variable.
 */
export let caCert: string | undefined = __config.get("caCert");
/**
 * Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
 */
export let insecureSkipVerify: boolean | undefined = __config.getObject<boolean>("insecureSkipVerify");
/**
 * The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
 * variable.
 */
export let orgId: number | undefined = __config.getObject<number>("orgId");
/**
 * Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
 * `GRAFANA_TLS_CERT` environment variable.
 */
export let tlsCert: string | undefined = __config.get("tlsCert");
/**
 * Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT`
 * environment variable.
 */
export let tlsKey: string | undefined = __config.get("tlsKey");
/**
 * The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
 */
export let url: string | undefined = __config.get("url");
