// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Grafana
{
    [GrafanaResourceType("grafana:index/builtinRoleAssignment:BuiltinRoleAssignment")]
    public partial class BuiltinRoleAssignment : Pulumi.CustomResource
    {
        [Output("builtinRole")]
        public Output<string> BuiltinRole { get; private set; } = null!;

        [Output("roles")]
        public Output<ImmutableArray<Outputs.BuiltinRoleAssignmentRole>> Roles { get; private set; } = null!;


        /// <summary>
        /// Create a BuiltinRoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BuiltinRoleAssignment(string name, BuiltinRoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/builtinRoleAssignment:BuiltinRoleAssignment", name, args ?? new BuiltinRoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BuiltinRoleAssignment(string name, Input<string> id, BuiltinRoleAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/builtinRoleAssignment:BuiltinRoleAssignment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BuiltinRoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BuiltinRoleAssignment Get(string name, Input<string> id, BuiltinRoleAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new BuiltinRoleAssignment(name, id, state, options);
        }
    }

    public sealed class BuiltinRoleAssignmentArgs : Pulumi.ResourceArgs
    {
        [Input("builtinRole", required: true)]
        public Input<string> BuiltinRole { get; set; } = null!;

        [Input("roles", required: true)]
        private InputList<Inputs.BuiltinRoleAssignmentRoleArgs>? _roles;
        public InputList<Inputs.BuiltinRoleAssignmentRoleArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.BuiltinRoleAssignmentRoleArgs>());
            set => _roles = value;
        }

        public BuiltinRoleAssignmentArgs()
        {
        }
    }

    public sealed class BuiltinRoleAssignmentState : Pulumi.ResourceArgs
    {
        [Input("builtinRole")]
        public Input<string>? BuiltinRole { get; set; }

        [Input("roles")]
        private InputList<Inputs.BuiltinRoleAssignmentRoleGetArgs>? _roles;
        public InputList<Inputs.BuiltinRoleAssignmentRoleGetArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.BuiltinRoleAssignmentRoleGetArgs>());
            set => _roles = value;
        }

        public BuiltinRoleAssignmentState()
        {
        }
    }
}
