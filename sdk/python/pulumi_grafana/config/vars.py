# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('grafana')


class _ExportableConfig(types.ModuleType):
    @property
    def auth(self) -> str:
        """
        API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        """
        return __config__.get('auth')

    @property
    def ca_cert(self) -> Optional[str]:
        """
        Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
        `GRAFANA_CA_CERT` environment variable.
        """
        return __config__.get('caCert')

    @property
    def insecure_skip_verify(self) -> Optional[str]:
        """
        Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        """
        return __config__.get('insecureSkipVerify')

    @property
    def org_id(self) -> str:
        """
        The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
        variable.
        """
        return __config__.get('orgId')

    @property
    def tls_cert(self) -> Optional[str]:
        """
        Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
        `GRAFANA_TLS_CERT` environment variable.
        """
        return __config__.get('tlsCert')

    @property
    def tls_key(self) -> Optional[str]:
        """
        Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT`
        environment variable.
        """
        return __config__.get('tlsKey')

    @property
    def url(self) -> str:
        """
        The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        return __config__.get('url')

