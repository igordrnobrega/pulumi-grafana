# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DataSourceArgs', 'DataSource']

@pulumi.input_type
class DataSourceArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 access_mode: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password: Optional[pulumi.Input[str]] = None,
                 basic_auth_username: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 json_datas: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 secure_json_datas: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceSecureJsonDataArgs']]]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataSource resource.
        :param pulumi.Input[str] type: The data source type. Must be one of the supported data source keywords.
        :param pulumi.Input[str] access_mode: The method by which Grafana will access the data source: `proxy` or `direct`.
        :param pulumi.Input[bool] basic_auth_enabled: Whether to enable basic auth for the data source.
        :param pulumi.Input[str] basic_auth_password: Basic auth password.
        :param pulumi.Input[str] basic_auth_username: Basic auth username.
        :param pulumi.Input[str] database_name: (Required by some data source types) The name of the database to use on the selected data source server.
        :param pulumi.Input[bool] is_default: Whether to set the data source as default. This should only be `true` to a single data source.
        :param pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]] json_datas: (Required by some data source types)
        :param pulumi.Input[str] name: A unique name for the data source.
        :param pulumi.Input[str] password: (Required by some data source types) The password to use to authenticate to the data source.
        :param pulumi.Input[str] url: The URL for the data source. The type of URL required varies depending on the chosen data source type.
        :param pulumi.Input[str] username: (Required by some data source types) The username to use to authenticate to the data source.
        """
        pulumi.set(__self__, "type", type)
        if access_mode is not None:
            pulumi.set(__self__, "access_mode", access_mode)
        if basic_auth_enabled is not None:
            pulumi.set(__self__, "basic_auth_enabled", basic_auth_enabled)
        if basic_auth_password is not None:
            pulumi.set(__self__, "basic_auth_password", basic_auth_password)
        if basic_auth_username is not None:
            pulumi.set(__self__, "basic_auth_username", basic_auth_username)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if json_datas is not None:
            pulumi.set(__self__, "json_datas", json_datas)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if secure_json_datas is not None:
            pulumi.set(__self__, "secure_json_datas", secure_json_datas)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The data source type. Must be one of the supported data source keywords.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="accessMode")
    def access_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The method by which Grafana will access the data source: `proxy` or `direct`.
        """
        return pulumi.get(self, "access_mode")

    @access_mode.setter
    def access_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_mode", value)

    @property
    @pulumi.getter(name="basicAuthEnabled")
    def basic_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable basic auth for the data source.
        """
        return pulumi.get(self, "basic_auth_enabled")

    @basic_auth_enabled.setter
    def basic_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "basic_auth_enabled", value)

    @property
    @pulumi.getter(name="basicAuthPassword")
    def basic_auth_password(self) -> Optional[pulumi.Input[str]]:
        """
        Basic auth password.
        """
        return pulumi.get(self, "basic_auth_password")

    @basic_auth_password.setter
    def basic_auth_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_password", value)

    @property
    @pulumi.getter(name="basicAuthUsername")
    def basic_auth_username(self) -> Optional[pulumi.Input[str]]:
        """
        Basic auth username.
        """
        return pulumi.get(self, "basic_auth_username")

    @basic_auth_username.setter
    def basic_auth_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_username", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        (Required by some data source types) The name of the database to use on the selected data source server.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to set the data source as default. This should only be `true` to a single data source.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter(name="jsonDatas")
    def json_datas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]]]:
        """
        (Required by some data source types)
        """
        return pulumi.get(self, "json_datas")

    @json_datas.setter
    def json_datas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]]]):
        pulumi.set(self, "json_datas", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        (Required by some data source types) The password to use to authenticate to the data source.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="secureJsonDatas")
    def secure_json_datas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceSecureJsonDataArgs']]]]:
        return pulumi.get(self, "secure_json_datas")

    @secure_json_datas.setter
    def secure_json_datas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceSecureJsonDataArgs']]]]):
        pulumi.set(self, "secure_json_datas", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for the data source. The type of URL required varies depending on the chosen data source type.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        (Required by some data source types) The username to use to authenticate to the data source.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _DataSourceState:
    def __init__(__self__, *,
                 access_mode: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password: Optional[pulumi.Input[str]] = None,
                 basic_auth_username: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 json_datas: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 secure_json_datas: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceSecureJsonDataArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataSource resources.
        :param pulumi.Input[str] access_mode: The method by which Grafana will access the data source: `proxy` or `direct`.
        :param pulumi.Input[bool] basic_auth_enabled: Whether to enable basic auth for the data source.
        :param pulumi.Input[str] basic_auth_password: Basic auth password.
        :param pulumi.Input[str] basic_auth_username: Basic auth username.
        :param pulumi.Input[str] database_name: (Required by some data source types) The name of the database to use on the selected data source server.
        :param pulumi.Input[bool] is_default: Whether to set the data source as default. This should only be `true` to a single data source.
        :param pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]] json_datas: (Required by some data source types)
        :param pulumi.Input[str] name: A unique name for the data source.
        :param pulumi.Input[str] password: (Required by some data source types) The password to use to authenticate to the data source.
        :param pulumi.Input[str] type: The data source type. Must be one of the supported data source keywords.
        :param pulumi.Input[str] url: The URL for the data source. The type of URL required varies depending on the chosen data source type.
        :param pulumi.Input[str] username: (Required by some data source types) The username to use to authenticate to the data source.
        """
        if access_mode is not None:
            pulumi.set(__self__, "access_mode", access_mode)
        if basic_auth_enabled is not None:
            pulumi.set(__self__, "basic_auth_enabled", basic_auth_enabled)
        if basic_auth_password is not None:
            pulumi.set(__self__, "basic_auth_password", basic_auth_password)
        if basic_auth_username is not None:
            pulumi.set(__self__, "basic_auth_username", basic_auth_username)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if json_datas is not None:
            pulumi.set(__self__, "json_datas", json_datas)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if secure_json_datas is not None:
            pulumi.set(__self__, "secure_json_datas", secure_json_datas)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="accessMode")
    def access_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The method by which Grafana will access the data source: `proxy` or `direct`.
        """
        return pulumi.get(self, "access_mode")

    @access_mode.setter
    def access_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_mode", value)

    @property
    @pulumi.getter(name="basicAuthEnabled")
    def basic_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable basic auth for the data source.
        """
        return pulumi.get(self, "basic_auth_enabled")

    @basic_auth_enabled.setter
    def basic_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "basic_auth_enabled", value)

    @property
    @pulumi.getter(name="basicAuthPassword")
    def basic_auth_password(self) -> Optional[pulumi.Input[str]]:
        """
        Basic auth password.
        """
        return pulumi.get(self, "basic_auth_password")

    @basic_auth_password.setter
    def basic_auth_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_password", value)

    @property
    @pulumi.getter(name="basicAuthUsername")
    def basic_auth_username(self) -> Optional[pulumi.Input[str]]:
        """
        Basic auth username.
        """
        return pulumi.get(self, "basic_auth_username")

    @basic_auth_username.setter
    def basic_auth_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_username", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        (Required by some data source types) The name of the database to use on the selected data source server.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to set the data source as default. This should only be `true` to a single data source.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter(name="jsonDatas")
    def json_datas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]]]:
        """
        (Required by some data source types)
        """
        return pulumi.get(self, "json_datas")

    @json_datas.setter
    def json_datas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceJsonDataArgs']]]]):
        pulumi.set(self, "json_datas", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        (Required by some data source types) The password to use to authenticate to the data source.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="secureJsonDatas")
    def secure_json_datas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceSecureJsonDataArgs']]]]:
        return pulumi.get(self, "secure_json_datas")

    @secure_json_datas.setter
    def secure_json_datas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceSecureJsonDataArgs']]]]):
        pulumi.set(self, "secure_json_datas", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The data source type. Must be one of the supported data source keywords.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for the data source. The type of URL required varies depending on the chosen data source type.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        (Required by some data source types) The username to use to authenticate to the data source.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class DataSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_mode: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password: Optional[pulumi.Input[str]] = None,
                 basic_auth_username: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 json_datas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceJsonDataArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 secure_json_datas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceSecureJsonDataArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DataSource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_mode: The method by which Grafana will access the data source: `proxy` or `direct`.
        :param pulumi.Input[bool] basic_auth_enabled: Whether to enable basic auth for the data source.
        :param pulumi.Input[str] basic_auth_password: Basic auth password.
        :param pulumi.Input[str] basic_auth_username: Basic auth username.
        :param pulumi.Input[str] database_name: (Required by some data source types) The name of the database to use on the selected data source server.
        :param pulumi.Input[bool] is_default: Whether to set the data source as default. This should only be `true` to a single data source.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceJsonDataArgs']]]] json_datas: (Required by some data source types)
        :param pulumi.Input[str] name: A unique name for the data source.
        :param pulumi.Input[str] password: (Required by some data source types) The password to use to authenticate to the data source.
        :param pulumi.Input[str] type: The data source type. Must be one of the supported data source keywords.
        :param pulumi.Input[str] url: The URL for the data source. The type of URL required varies depending on the chosen data source type.
        :param pulumi.Input[str] username: (Required by some data source types) The username to use to authenticate to the data source.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DataSource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DataSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_mode: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password: Optional[pulumi.Input[str]] = None,
                 basic_auth_username: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 json_datas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceJsonDataArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 secure_json_datas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceSecureJsonDataArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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
            __props__ = DataSourceArgs.__new__(DataSourceArgs)

            __props__.__dict__["access_mode"] = access_mode
            __props__.__dict__["basic_auth_enabled"] = basic_auth_enabled
            __props__.__dict__["basic_auth_password"] = basic_auth_password
            __props__.__dict__["basic_auth_username"] = basic_auth_username
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["is_default"] = is_default
            __props__.__dict__["json_datas"] = json_datas
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = password
            __props__.__dict__["secure_json_datas"] = secure_json_datas
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["url"] = url
            __props__.__dict__["username"] = username
        super(DataSource, __self__).__init__(
            'grafana:index/dataSource:DataSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_mode: Optional[pulumi.Input[str]] = None,
            basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
            basic_auth_password: Optional[pulumi.Input[str]] = None,
            basic_auth_username: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            json_datas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceJsonDataArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            secure_json_datas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceSecureJsonDataArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'DataSource':
        """
        Get an existing DataSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_mode: The method by which Grafana will access the data source: `proxy` or `direct`.
        :param pulumi.Input[bool] basic_auth_enabled: Whether to enable basic auth for the data source.
        :param pulumi.Input[str] basic_auth_password: Basic auth password.
        :param pulumi.Input[str] basic_auth_username: Basic auth username.
        :param pulumi.Input[str] database_name: (Required by some data source types) The name of the database to use on the selected data source server.
        :param pulumi.Input[bool] is_default: Whether to set the data source as default. This should only be `true` to a single data source.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceJsonDataArgs']]]] json_datas: (Required by some data source types)
        :param pulumi.Input[str] name: A unique name for the data source.
        :param pulumi.Input[str] password: (Required by some data source types) The password to use to authenticate to the data source.
        :param pulumi.Input[str] type: The data source type. Must be one of the supported data source keywords.
        :param pulumi.Input[str] url: The URL for the data source. The type of URL required varies depending on the chosen data source type.
        :param pulumi.Input[str] username: (Required by some data source types) The username to use to authenticate to the data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSourceState.__new__(_DataSourceState)

        __props__.__dict__["access_mode"] = access_mode
        __props__.__dict__["basic_auth_enabled"] = basic_auth_enabled
        __props__.__dict__["basic_auth_password"] = basic_auth_password
        __props__.__dict__["basic_auth_username"] = basic_auth_username
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["is_default"] = is_default
        __props__.__dict__["json_datas"] = json_datas
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["secure_json_datas"] = secure_json_datas
        __props__.__dict__["type"] = type
        __props__.__dict__["url"] = url
        __props__.__dict__["username"] = username
        return DataSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessMode")
    def access_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The method by which Grafana will access the data source: `proxy` or `direct`.
        """
        return pulumi.get(self, "access_mode")

    @property
    @pulumi.getter(name="basicAuthEnabled")
    def basic_auth_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable basic auth for the data source.
        """
        return pulumi.get(self, "basic_auth_enabled")

    @property
    @pulumi.getter(name="basicAuthPassword")
    def basic_auth_password(self) -> pulumi.Output[Optional[str]]:
        """
        Basic auth password.
        """
        return pulumi.get(self, "basic_auth_password")

    @property
    @pulumi.getter(name="basicAuthUsername")
    def basic_auth_username(self) -> pulumi.Output[Optional[str]]:
        """
        Basic auth username.
        """
        return pulumi.get(self, "basic_auth_username")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[Optional[str]]:
        """
        (Required by some data source types) The name of the database to use on the selected data source server.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to set the data source as default. This should only be `true` to a single data source.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="jsonDatas")
    def json_datas(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourceJsonData']]]:
        """
        (Required by some data source types)
        """
        return pulumi.get(self, "json_datas")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the data source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        (Required by some data source types) The password to use to authenticate to the data source.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="secureJsonDatas")
    def secure_json_datas(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourceSecureJsonData']]]:
        return pulumi.get(self, "secure_json_datas")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The data source type. Must be one of the supported data source keywords.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL for the data source. The type of URL required varies depending on the chosen data source type.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        (Required by some data source types) The username to use to authenticate to the data source.
        """
        return pulumi.get(self, "username")
