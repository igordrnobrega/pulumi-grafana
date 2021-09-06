# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DashboardArgs', 'Dashboard']

@pulumi.input_type
class DashboardArgs:
    def __init__(__self__, *,
                 config_json: pulumi.Input[str],
                 folder: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Dashboard resource.
        :param pulumi.Input[str] config_json: The complete dashboard model JSON.
        :param pulumi.Input[int] folder: The id of the folder to save the dashboard in.
        """
        pulumi.set(__self__, "config_json", config_json)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)

    @property
    @pulumi.getter(name="configJson")
    def config_json(self) -> pulumi.Input[str]:
        """
        The complete dashboard model JSON.
        """
        return pulumi.get(self, "config_json")

    @config_json.setter
    def config_json(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_json", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the folder to save the dashboard in.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "folder", value)


@pulumi.input_type
class _DashboardState:
    def __init__(__self__, *,
                 config_json: Optional[pulumi.Input[str]] = None,
                 dashboard_id: Optional[pulumi.Input[int]] = None,
                 folder: Optional[pulumi.Input[int]] = None,
                 slug: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Dashboard resources.
        :param pulumi.Input[str] config_json: The complete dashboard model JSON.
        :param pulumi.Input[int] dashboard_id: The numeric ID of the dashboard computed by Grafana.
        :param pulumi.Input[int] folder: The id of the folder to save the dashboard in.
        :param pulumi.Input[str] slug: URL friendly version of the dashboard title.
        """
        if config_json is not None:
            pulumi.set(__self__, "config_json", config_json)
        if dashboard_id is not None:
            pulumi.set(__self__, "dashboard_id", dashboard_id)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)

    @property
    @pulumi.getter(name="configJson")
    def config_json(self) -> Optional[pulumi.Input[str]]:
        """
        The complete dashboard model JSON.
        """
        return pulumi.get(self, "config_json")

    @config_json.setter
    def config_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_json", value)

    @property
    @pulumi.getter(name="dashboardId")
    def dashboard_id(self) -> Optional[pulumi.Input[int]]:
        """
        The numeric ID of the dashboard computed by Grafana.
        """
        return pulumi.get(self, "dashboard_id")

    @dashboard_id.setter
    def dashboard_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dashboard_id", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the folder to save the dashboard in.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        """
        URL friendly version of the dashboard title.
        """
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)


class Dashboard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_json: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Dashboard resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_json: The complete dashboard model JSON.
        :param pulumi.Input[int] folder: The id of the folder to save the dashboard in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DashboardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Dashboard resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DashboardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DashboardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_json: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[int]] = None,
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
            __props__ = DashboardArgs.__new__(DashboardArgs)

            if config_json is None and not opts.urn:
                raise TypeError("Missing required property 'config_json'")
            __props__.__dict__["config_json"] = config_json
            __props__.__dict__["folder"] = folder
            __props__.__dict__["dashboard_id"] = None
            __props__.__dict__["slug"] = None
        super(Dashboard, __self__).__init__(
            'grafana:index/dashboard:Dashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_json: Optional[pulumi.Input[str]] = None,
            dashboard_id: Optional[pulumi.Input[int]] = None,
            folder: Optional[pulumi.Input[int]] = None,
            slug: Optional[pulumi.Input[str]] = None) -> 'Dashboard':
        """
        Get an existing Dashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_json: The complete dashboard model JSON.
        :param pulumi.Input[int] dashboard_id: The numeric ID of the dashboard computed by Grafana.
        :param pulumi.Input[int] folder: The id of the folder to save the dashboard in.
        :param pulumi.Input[str] slug: URL friendly version of the dashboard title.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DashboardState.__new__(_DashboardState)

        __props__.__dict__["config_json"] = config_json
        __props__.__dict__["dashboard_id"] = dashboard_id
        __props__.__dict__["folder"] = folder
        __props__.__dict__["slug"] = slug
        return Dashboard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configJson")
    def config_json(self) -> pulumi.Output[str]:
        """
        The complete dashboard model JSON.
        """
        return pulumi.get(self, "config_json")

    @property
    @pulumi.getter(name="dashboardId")
    def dashboard_id(self) -> pulumi.Output[int]:
        """
        The numeric ID of the dashboard computed by Grafana.
        """
        return pulumi.get(self, "dashboard_id")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[Optional[int]]:
        """
        The id of the folder to save the dashboard in.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        """
        URL friendly version of the dashboard title.
        """
        return pulumi.get(self, "slug")
