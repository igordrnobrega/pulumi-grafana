// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Grafana
{
    [GrafanaResourceType("grafana:index/organization:Organization")]
    public partial class Organization : Pulumi.CustomResource
    {
        /// <summary>
        /// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
        /// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
        /// this parameter keeps Terraform from removing it from organizations.
        /// </summary>
        [Output("adminUser")]
        public Output<string?> AdminUser { get; private set; } = null!;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given admin access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        [Output("admins")]
        public Output<ImmutableArray<string>> Admins { get; private set; } = null!;

        /// <summary>
        /// Whether or not to create Grafana users specified in the organization's membership if they don't already exist in
        /// Grafana. If unspecified, this parameter defaults to true, creating placeholder users with the name, login, and email set
        /// to the email of the user, and a random password. Setting this option to false will cause an error to be thrown for any
        /// users that do not already exist in Grafana.
        /// </summary>
        [Output("createUsers")]
        public Output<bool?> CreateUsers { get; private set; } = null!;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given editor access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        [Output("editors")]
        public Output<ImmutableArray<string>> Editors { get; private set; } = null!;

        /// <summary>
        /// The display name for the Grafana organization created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization id assigned to this organization by Grafana.
        /// </summary>
        [Output("orgId")]
        public Output<int> OrgId { get; private set; } = null!;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given viewer access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        [Output("viewers")]
        public Output<ImmutableArray<string>> Viewers { get; private set; } = null!;


        /// <summary>
        /// Create a Organization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Organization(string name, OrganizationArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/organization:Organization", name, args ?? new OrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Organization(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/organization:Organization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Organization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Organization Get(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
        {
            return new Organization(name, id, state, options);
        }
    }

    public sealed class OrganizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
        /// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
        /// this parameter keeps Terraform from removing it from organizations.
        /// </summary>
        [Input("adminUser")]
        public Input<string>? AdminUser { get; set; }

        [Input("admins")]
        private InputList<string>? _admins;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given admin access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        public InputList<string> Admins
        {
            get => _admins ?? (_admins = new InputList<string>());
            set => _admins = value;
        }

        /// <summary>
        /// Whether or not to create Grafana users specified in the organization's membership if they don't already exist in
        /// Grafana. If unspecified, this parameter defaults to true, creating placeholder users with the name, login, and email set
        /// to the email of the user, and a random password. Setting this option to false will cause an error to be thrown for any
        /// users that do not already exist in Grafana.
        /// </summary>
        [Input("createUsers")]
        public Input<bool>? CreateUsers { get; set; }

        [Input("editors")]
        private InputList<string>? _editors;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given editor access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        public InputList<string> Editors
        {
            get => _editors ?? (_editors = new InputList<string>());
            set => _editors = value;
        }

        /// <summary>
        /// The display name for the Grafana organization created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("viewers")]
        private InputList<string>? _viewers;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given viewer access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        public InputList<string> Viewers
        {
            get => _viewers ?? (_viewers = new InputList<string>());
            set => _viewers = value;
        }

        public OrganizationArgs()
        {
        }
    }

    public sealed class OrganizationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
        /// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
        /// this parameter keeps Terraform from removing it from organizations.
        /// </summary>
        [Input("adminUser")]
        public Input<string>? AdminUser { get; set; }

        [Input("admins")]
        private InputList<string>? _admins;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given admin access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        public InputList<string> Admins
        {
            get => _admins ?? (_admins = new InputList<string>());
            set => _admins = value;
        }

        /// <summary>
        /// Whether or not to create Grafana users specified in the organization's membership if they don't already exist in
        /// Grafana. If unspecified, this parameter defaults to true, creating placeholder users with the name, login, and email set
        /// to the email of the user, and a random password. Setting this option to false will cause an error to be thrown for any
        /// users that do not already exist in Grafana.
        /// </summary>
        [Input("createUsers")]
        public Input<bool>? CreateUsers { get; set; }

        [Input("editors")]
        private InputList<string>? _editors;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given editor access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        public InputList<string> Editors
        {
            get => _editors ?? (_editors = new InputList<string>());
            set => _editors = value;
        }

        /// <summary>
        /// The display name for the Grafana organization created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization id assigned to this organization by Grafana.
        /// </summary>
        [Input("orgId")]
        public Input<int>? OrgId { get; set; }

        [Input("viewers")]
        private InputList<string>? _viewers;

        /// <summary>
        /// A list of email addresses corresponding to users who should be given viewer access to the organization. Note: users
        /// specified here must already exist in Grafana unless 'create_users' is set to true.
        /// </summary>
        public InputList<string> Viewers
        {
            get => _viewers ?? (_viewers = new InputList<string>());
            set => _viewers = value;
        }

        public OrganizationState()
        {
        }
    }
}