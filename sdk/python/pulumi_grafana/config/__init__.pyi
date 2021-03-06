# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

auth: str
"""
API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
"""

caCert: Optional[str]
"""
Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
`GRAFANA_CA_CERT` environment variable.
"""

insecureSkipVerify: Optional[str]
"""
Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
"""

orgId: str
"""
The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
variable.
"""

tlsCert: Optional[str]
"""
Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
`GRAFANA_TLS_CERT` environment variable.
"""

tlsKey: Optional[str]
"""
Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT`
environment variable.
"""

url: str
"""
The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
"""

