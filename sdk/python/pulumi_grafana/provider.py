# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 auth: pulumi.Input[str],
                 org_id: pulumi.Input[int],
                 url: pulumi.Input[str],
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 tls_cert: Optional[pulumi.Input[str]] = None,
                 tls_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] auth: API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        :param pulumi.Input[int] org_id: The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
               variable.
        :param pulumi.Input[str] url: The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        :param pulumi.Input[str] ca_cert: Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
               `GRAFANA_CA_CERT` environment variable.
        :param pulumi.Input[bool] insecure_skip_verify: Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        :param pulumi.Input[str] tls_cert: Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
               `GRAFANA_TLS_CERT` environment variable.
        :param pulumi.Input[str] tls_key: Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT`
               environment variable.
        """
        pulumi.set(__self__, "auth", auth)
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "url", url)
        if ca_cert is not None:
            pulumi.set(__self__, "ca_cert", ca_cert)
        if insecure_skip_verify is not None:
            pulumi.set(__self__, "insecure_skip_verify", insecure_skip_verify)
        if tls_cert is not None:
            pulumi.set(__self__, "tls_cert", tls_cert)
        if tls_key is not None:
            pulumi.set(__self__, "tls_key", tls_key)

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Input[str]:
        """
        API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[int]:
        """
        The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
        variable.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
        `GRAFANA_CA_CERT` environment variable.
        """
        return pulumi.get(self, "ca_cert")

    @ca_cert.setter
    def ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert", value)

    @property
    @pulumi.getter(name="insecureSkipVerify")
    def insecure_skip_verify(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        """
        return pulumi.get(self, "insecure_skip_verify")

    @insecure_skip_verify.setter
    def insecure_skip_verify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure_skip_verify", value)

    @property
    @pulumi.getter(name="tlsCert")
    def tls_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
        `GRAFANA_TLS_CERT` environment variable.
        """
        return pulumi.get(self, "tls_cert")

    @tls_cert.setter
    def tls_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_cert", value)

    @property
    @pulumi.getter(name="tlsKey")
    def tls_key(self) -> Optional[pulumi.Input[str]]:
        """
        Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT`
        environment variable.
        """
        return pulumi.get(self, "tls_key")

    @tls_key.setter
    def tls_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_key", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 org_id: Optional[pulumi.Input[int]] = None,
                 tls_cert: Optional[pulumi.Input[str]] = None,
                 tls_key: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the grafana package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth: API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        :param pulumi.Input[str] ca_cert: Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
               `GRAFANA_CA_CERT` environment variable.
        :param pulumi.Input[bool] insecure_skip_verify: Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        :param pulumi.Input[int] org_id: The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
               variable.
        :param pulumi.Input[str] tls_cert: Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
               `GRAFANA_TLS_CERT` environment variable.
        :param pulumi.Input[str] tls_key: Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT`
               environment variable.
        :param pulumi.Input[str] url: The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the grafana package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 org_id: Optional[pulumi.Input[int]] = None,
                 tls_cert: Optional[pulumi.Input[str]] = None,
                 tls_key: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if auth is None and not opts.urn:
                raise TypeError("Missing required property 'auth'")
            __props__.__dict__["auth"] = auth
            __props__.__dict__["ca_cert"] = ca_cert
            __props__.__dict__["insecure_skip_verify"] = pulumi.Output.from_input(insecure_skip_verify).apply(pulumi.runtime.to_json) if insecure_skip_verify is not None else None
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = pulumi.Output.from_input(org_id).apply(pulumi.runtime.to_json) if org_id is not None else None
            __props__.__dict__["tls_cert"] = tls_cert
            __props__.__dict__["tls_key"] = tls_key
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
        super(Provider, __self__).__init__(
            'grafana',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Output[str]:
        """
        API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
        `GRAFANA_CA_CERT` environment variable.
        """
        return pulumi.get(self, "ca_cert")

    @property
    @pulumi.getter(name="tlsCert")
    def tls_cert(self) -> pulumi.Output[Optional[str]]:
        """
        Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
        `GRAFANA_TLS_CERT` environment variable.
        """
        return pulumi.get(self, "tls_cert")

    @property
    @pulumi.getter(name="tlsKey")
    def tls_key(self) -> pulumi.Output[Optional[str]]:
        """
        Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT`
        environment variable.
        """
        return pulumi.get(self, "tls_key")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        return pulumi.get(self, "url")

