// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Grafana.Outputs
{

    [OutputType]
    public sealed class FolderPermissionPermission
    {
        public readonly string Permission;
        public readonly string? Role;
        public readonly int? TeamId;
        public readonly int? UserId;

        [OutputConstructor]
        private FolderPermissionPermission(
            string permission,

            string? role,

            int? teamId,

            int? userId)
        {
            Permission = permission;
            Role = role;
            TeamId = teamId;
            UserId = userId;
        }
    }
}
