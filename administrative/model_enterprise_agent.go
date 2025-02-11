/*
Administrative API

Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API. This API provides the following operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These operations can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.    For more information about the administrative models, see [Account Management](https://docs.thousandeyes.com/product-documentation/user-management).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package administrative

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/core"
	"time"
	"bytes"
	"fmt"
)

// checks if the EnterpriseAgent type satisfies the MappedNullable interface at compile time
var _ core.MappedNullable = &EnterpriseAgent{}

// EnterpriseAgent struct for EnterpriseAgent
type EnterpriseAgent struct {
	AgentType CloudEnterpriseAgentType `json:"agentType"`
	// Array of private IP addresses.
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// Array of public IP addresses.
	PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`
	// Network (including ASN) of agent’s public IP.
	Network *string `json:"network,omitempty"`
	// Unique ID of the agent.
	AgentId *string `json:"agentId,omitempty"`
	// Name of the agent.
	AgentName *string `json:"agentName,omitempty"`
	// Location of the agent.
	Location *string `json:"location,omitempty"`
	// 2-digit ISO country code
	CountryId *string `json:"countryId,omitempty"`
	// Flag indicating if the agent is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix containing agents public IP address.
	Prefix *string `json:"prefix,omitempty"`
	// Flag indicating if has normal SSL operations or  if instead it's set to ignore SSL errors on browserbot-based tests.
	VerifySslCertificates *bool `json:"verifySslCertificates,omitempty"`
	// If an enterprise agent is clustered, detailed information about each cluster member will be shown as array entries in the clusterMembers field. This field is not shown for Enterprise Agents in standalone mode, or for Cloud Agents.
	ClusterMembers []ClusterMember `json:"clusterMembers,omitempty"`
	// Shows overall utilization percentage (online Enterprise Agents and Enterprise Clusters only).
	Utilization *int32 `json:"utilization,omitempty"`
	// List of account groups. See /accounts-groups to pull a list of account IDs
	AccountGroups []AccountGroup `json:"accountGroups,omitempty"`
	Ipv6Policy *EnterpriseAgentIpv6Policy `json:"ipv6Policy,omitempty"`
	// If an enterprise agent or a cluster member presents at least one error, the errors will be shown as an array of entries in the errorDetails field (Enterprise Agents and Enterprise Cluster members only)
	ErrorDetails []ErrorDetail `json:"errorDetails,omitempty"`
	// Fully qualified domain name of the agent (Enterprise Agents only)
	Hostname *string `json:"hostname,omitempty"`
	// UTC last seen date (ISO date-time format).
	LastSeen *time.Time `json:"lastSeen,omitempty"`
	AgentState *EnterpriseAgentState `json:"agentState,omitempty"`
	// Flag indicating if the agent retains cache.
	KeepBrowserCache *bool `json:"keepBrowserCache,omitempty"`
	// UTC Agent creation date (ISO date-time format).
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Test target IP address.
	TargetForTests *string `json:"targetForTests,omitempty"`
	// To perform rDNS lookups for public IP ranges, this field represents the public IP ranges. The range must be in CIDR notation; for example, 10.1.1.0/24. Maximum of 5 prefixes allowed (Enterprise Agents and Enterprise Agent clusters only).
	LocalResolutionPrefixes []string `json:"localResolutionPrefixes,omitempty"`
	InterfaceIpMappings []InterfaceIpMapping `json:"interfaceIpMappings,omitempty"`
}

type _EnterpriseAgent EnterpriseAgent

// NewEnterpriseAgent instantiates a new EnterpriseAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAgent(agentType CloudEnterpriseAgentType) *EnterpriseAgent {
	this := EnterpriseAgent{}
	this.AgentType = agentType
	return &this
}

// NewEnterpriseAgentWithDefaults instantiates a new EnterpriseAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAgentWithDefaults() *EnterpriseAgent {
	this := EnterpriseAgent{}
	return &this
}

// GetAgentType returns the AgentType field value
func (o *EnterpriseAgent) GetAgentType() CloudEnterpriseAgentType {
	if o == nil {
		var ret CloudEnterpriseAgentType
		return ret
	}

	return o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetAgentTypeOk() (*CloudEnterpriseAgentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentType, true
}

// SetAgentType sets field value
func (o *EnterpriseAgent) SetAgentType(v CloudEnterpriseAgentType) {
	o.AgentType = v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetIpAddresses() []string {
	if o == nil || core.IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetIpAddressesOk() ([]string, bool) {
	if o == nil || core.IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasIpAddresses() bool {
	if o != nil && !core.IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *EnterpriseAgent) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetPublicIpAddresses returns the PublicIpAddresses field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetPublicIpAddresses() []string {
	if o == nil || core.IsNil(o.PublicIpAddresses) {
		var ret []string
		return ret
	}
	return o.PublicIpAddresses
}

// GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetPublicIpAddressesOk() ([]string, bool) {
	if o == nil || core.IsNil(o.PublicIpAddresses) {
		return nil, false
	}
	return o.PublicIpAddresses, true
}

// HasPublicIpAddresses returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasPublicIpAddresses() bool {
	if o != nil && !core.IsNil(o.PublicIpAddresses) {
		return true
	}

	return false
}

// SetPublicIpAddresses gets a reference to the given []string and assigns it to the PublicIpAddresses field.
func (o *EnterpriseAgent) SetPublicIpAddresses(v []string) {
	o.PublicIpAddresses = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetNetwork() string {
	if o == nil || core.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetNetworkOk() (*string, bool) {
	if o == nil || core.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasNetwork() bool {
	if o != nil && !core.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *EnterpriseAgent) SetNetwork(v string) {
	o.Network = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetAgentId() string {
	if o == nil || core.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetAgentIdOk() (*string, bool) {
	if o == nil || core.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasAgentId() bool {
	if o != nil && !core.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *EnterpriseAgent) SetAgentId(v string) {
	o.AgentId = &v
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetAgentName() string {
	if o == nil || core.IsNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetAgentNameOk() (*string, bool) {
	if o == nil || core.IsNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasAgentName() bool {
	if o != nil && !core.IsNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *EnterpriseAgent) SetAgentName(v string) {
	o.AgentName = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetLocation() string {
	if o == nil || core.IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetLocationOk() (*string, bool) {
	if o == nil || core.IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasLocation() bool {
	if o != nil && !core.IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *EnterpriseAgent) SetLocation(v string) {
	o.Location = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetCountryId() string {
	if o == nil || core.IsNil(o.CountryId) {
		var ret string
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetCountryIdOk() (*string, bool) {
	if o == nil || core.IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasCountryId() bool {
	if o != nil && !core.IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given string and assigns it to the CountryId field.
func (o *EnterpriseAgent) SetCountryId(v string) {
	o.CountryId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetEnabled() bool {
	if o == nil || core.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetEnabledOk() (*bool, bool) {
	if o == nil || core.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasEnabled() bool {
	if o != nil && !core.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EnterpriseAgent) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetPrefix() string {
	if o == nil || core.IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetPrefixOk() (*string, bool) {
	if o == nil || core.IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasPrefix() bool {
	if o != nil && !core.IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *EnterpriseAgent) SetPrefix(v string) {
	o.Prefix = &v
}

// GetVerifySslCertificates returns the VerifySslCertificates field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetVerifySslCertificates() bool {
	if o == nil || core.IsNil(o.VerifySslCertificates) {
		var ret bool
		return ret
	}
	return *o.VerifySslCertificates
}

// GetVerifySslCertificatesOk returns a tuple with the VerifySslCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetVerifySslCertificatesOk() (*bool, bool) {
	if o == nil || core.IsNil(o.VerifySslCertificates) {
		return nil, false
	}
	return o.VerifySslCertificates, true
}

// HasVerifySslCertificates returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasVerifySslCertificates() bool {
	if o != nil && !core.IsNil(o.VerifySslCertificates) {
		return true
	}

	return false
}

// SetVerifySslCertificates gets a reference to the given bool and assigns it to the VerifySslCertificates field.
func (o *EnterpriseAgent) SetVerifySslCertificates(v bool) {
	o.VerifySslCertificates = &v
}

// GetClusterMembers returns the ClusterMembers field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetClusterMembers() []ClusterMember {
	if o == nil || core.IsNil(o.ClusterMembers) {
		var ret []ClusterMember
		return ret
	}
	return o.ClusterMembers
}

// GetClusterMembersOk returns a tuple with the ClusterMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetClusterMembersOk() ([]ClusterMember, bool) {
	if o == nil || core.IsNil(o.ClusterMembers) {
		return nil, false
	}
	return o.ClusterMembers, true
}

// HasClusterMembers returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasClusterMembers() bool {
	if o != nil && !core.IsNil(o.ClusterMembers) {
		return true
	}

	return false
}

// SetClusterMembers gets a reference to the given []ClusterMember and assigns it to the ClusterMembers field.
func (o *EnterpriseAgent) SetClusterMembers(v []ClusterMember) {
	o.ClusterMembers = v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetUtilization() int32 {
	if o == nil || core.IsNil(o.Utilization) {
		var ret int32
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetUtilizationOk() (*int32, bool) {
	if o == nil || core.IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasUtilization() bool {
	if o != nil && !core.IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given int32 and assigns it to the Utilization field.
func (o *EnterpriseAgent) SetUtilization(v int32) {
	o.Utilization = &v
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetAccountGroups() []AccountGroup {
	if o == nil || core.IsNil(o.AccountGroups) {
		var ret []AccountGroup
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetAccountGroupsOk() ([]AccountGroup, bool) {
	if o == nil || core.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasAccountGroups() bool {
	if o != nil && !core.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []AccountGroup and assigns it to the AccountGroups field.
func (o *EnterpriseAgent) SetAccountGroups(v []AccountGroup) {
	o.AccountGroups = v
}

// GetIpv6Policy returns the Ipv6Policy field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetIpv6Policy() EnterpriseAgentIpv6Policy {
	if o == nil || core.IsNil(o.Ipv6Policy) {
		var ret EnterpriseAgentIpv6Policy
		return ret
	}
	return *o.Ipv6Policy
}

// GetIpv6PolicyOk returns a tuple with the Ipv6Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetIpv6PolicyOk() (*EnterpriseAgentIpv6Policy, bool) {
	if o == nil || core.IsNil(o.Ipv6Policy) {
		return nil, false
	}
	return o.Ipv6Policy, true
}

// HasIpv6Policy returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasIpv6Policy() bool {
	if o != nil && !core.IsNil(o.Ipv6Policy) {
		return true
	}

	return false
}

// SetIpv6Policy gets a reference to the given EnterpriseAgentIpv6Policy and assigns it to the Ipv6Policy field.
func (o *EnterpriseAgent) SetIpv6Policy(v EnterpriseAgentIpv6Policy) {
	o.Ipv6Policy = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetErrorDetails() []ErrorDetail {
	if o == nil || core.IsNil(o.ErrorDetails) {
		var ret []ErrorDetail
		return ret
	}
	return o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetErrorDetailsOk() ([]ErrorDetail, bool) {
	if o == nil || core.IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasErrorDetails() bool {
	if o != nil && !core.IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given []ErrorDetail and assigns it to the ErrorDetails field.
func (o *EnterpriseAgent) SetErrorDetails(v []ErrorDetail) {
	o.ErrorDetails = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetHostname() string {
	if o == nil || core.IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetHostnameOk() (*string, bool) {
	if o == nil || core.IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasHostname() bool {
	if o != nil && !core.IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *EnterpriseAgent) SetHostname(v string) {
	o.Hostname = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetLastSeen() time.Time {
	if o == nil || core.IsNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || core.IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasLastSeen() bool {
	if o != nil && !core.IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *EnterpriseAgent) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetAgentState returns the AgentState field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetAgentState() EnterpriseAgentState {
	if o == nil || core.IsNil(o.AgentState) {
		var ret EnterpriseAgentState
		return ret
	}
	return *o.AgentState
}

// GetAgentStateOk returns a tuple with the AgentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetAgentStateOk() (*EnterpriseAgentState, bool) {
	if o == nil || core.IsNil(o.AgentState) {
		return nil, false
	}
	return o.AgentState, true
}

// HasAgentState returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasAgentState() bool {
	if o != nil && !core.IsNil(o.AgentState) {
		return true
	}

	return false
}

// SetAgentState gets a reference to the given EnterpriseAgentState and assigns it to the AgentState field.
func (o *EnterpriseAgent) SetAgentState(v EnterpriseAgentState) {
	o.AgentState = &v
}

// GetKeepBrowserCache returns the KeepBrowserCache field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetKeepBrowserCache() bool {
	if o == nil || core.IsNil(o.KeepBrowserCache) {
		var ret bool
		return ret
	}
	return *o.KeepBrowserCache
}

// GetKeepBrowserCacheOk returns a tuple with the KeepBrowserCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetKeepBrowserCacheOk() (*bool, bool) {
	if o == nil || core.IsNil(o.KeepBrowserCache) {
		return nil, false
	}
	return o.KeepBrowserCache, true
}

// HasKeepBrowserCache returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasKeepBrowserCache() bool {
	if o != nil && !core.IsNil(o.KeepBrowserCache) {
		return true
	}

	return false
}

// SetKeepBrowserCache gets a reference to the given bool and assigns it to the KeepBrowserCache field.
func (o *EnterpriseAgent) SetKeepBrowserCache(v bool) {
	o.KeepBrowserCache = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetCreatedDate() time.Time {
	if o == nil || core.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || core.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasCreatedDate() bool {
	if o != nil && !core.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *EnterpriseAgent) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetTargetForTests returns the TargetForTests field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetTargetForTests() string {
	if o == nil || core.IsNil(o.TargetForTests) {
		var ret string
		return ret
	}
	return *o.TargetForTests
}

// GetTargetForTestsOk returns a tuple with the TargetForTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetTargetForTestsOk() (*string, bool) {
	if o == nil || core.IsNil(o.TargetForTests) {
		return nil, false
	}
	return o.TargetForTests, true
}

// HasTargetForTests returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasTargetForTests() bool {
	if o != nil && !core.IsNil(o.TargetForTests) {
		return true
	}

	return false
}

// SetTargetForTests gets a reference to the given string and assigns it to the TargetForTests field.
func (o *EnterpriseAgent) SetTargetForTests(v string) {
	o.TargetForTests = &v
}

// GetLocalResolutionPrefixes returns the LocalResolutionPrefixes field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetLocalResolutionPrefixes() []string {
	if o == nil || core.IsNil(o.LocalResolutionPrefixes) {
		var ret []string
		return ret
	}
	return o.LocalResolutionPrefixes
}

// GetLocalResolutionPrefixesOk returns a tuple with the LocalResolutionPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetLocalResolutionPrefixesOk() ([]string, bool) {
	if o == nil || core.IsNil(o.LocalResolutionPrefixes) {
		return nil, false
	}
	return o.LocalResolutionPrefixes, true
}

// HasLocalResolutionPrefixes returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasLocalResolutionPrefixes() bool {
	if o != nil && !core.IsNil(o.LocalResolutionPrefixes) {
		return true
	}

	return false
}

// SetLocalResolutionPrefixes gets a reference to the given []string and assigns it to the LocalResolutionPrefixes field.
func (o *EnterpriseAgent) SetLocalResolutionPrefixes(v []string) {
	o.LocalResolutionPrefixes = v
}

// GetInterfaceIpMappings returns the InterfaceIpMappings field value if set, zero value otherwise.
func (o *EnterpriseAgent) GetInterfaceIpMappings() []InterfaceIpMapping {
	if o == nil || core.IsNil(o.InterfaceIpMappings) {
		var ret []InterfaceIpMapping
		return ret
	}
	return o.InterfaceIpMappings
}

// GetInterfaceIpMappingsOk returns a tuple with the InterfaceIpMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgent) GetInterfaceIpMappingsOk() ([]InterfaceIpMapping, bool) {
	if o == nil || core.IsNil(o.InterfaceIpMappings) {
		return nil, false
	}
	return o.InterfaceIpMappings, true
}

// HasInterfaceIpMappings returns a boolean if a field has been set.
func (o *EnterpriseAgent) HasInterfaceIpMappings() bool {
	if o != nil && !core.IsNil(o.InterfaceIpMappings) {
		return true
	}

	return false
}

// SetInterfaceIpMappings gets a reference to the given []InterfaceIpMapping and assigns it to the InterfaceIpMappings field.
func (o *EnterpriseAgent) SetInterfaceIpMappings(v []InterfaceIpMapping) {
	o.InterfaceIpMappings = v
}

func (o EnterpriseAgent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseAgent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agentType"] = o.AgentType
	if !core.IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !core.IsNil(o.PublicIpAddresses) {
		toSerialize["publicIpAddresses"] = o.PublicIpAddresses
	}
	if !core.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !core.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !core.IsNil(o.AgentName) {
		toSerialize["agentName"] = o.AgentName
	}
	if !core.IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !core.IsNil(o.CountryId) {
		toSerialize["countryId"] = o.CountryId
	}
	if !core.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !core.IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !core.IsNil(o.VerifySslCertificates) {
		toSerialize["verifySslCertificates"] = o.VerifySslCertificates
	}
	if !core.IsNil(o.ClusterMembers) {
		toSerialize["clusterMembers"] = o.ClusterMembers
	}
	if !core.IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	if !core.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if !core.IsNil(o.Ipv6Policy) {
		toSerialize["ipv6Policy"] = o.Ipv6Policy
	}
	if !core.IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !core.IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !core.IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !core.IsNil(o.AgentState) {
		toSerialize["agentState"] = o.AgentState
	}
	if !core.IsNil(o.KeepBrowserCache) {
		toSerialize["keepBrowserCache"] = o.KeepBrowserCache
	}
	if !core.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !core.IsNil(o.TargetForTests) {
		toSerialize["targetForTests"] = o.TargetForTests
	}
	if !core.IsNil(o.LocalResolutionPrefixes) {
		toSerialize["localResolutionPrefixes"] = o.LocalResolutionPrefixes
	}
	if !core.IsNil(o.InterfaceIpMappings) {
		toSerialize["interfaceIpMappings"] = o.InterfaceIpMappings
	}
	return toSerialize, nil
}

func (o *EnterpriseAgent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agentType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEnterpriseAgent := _EnterpriseAgent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnterpriseAgent)

	if err != nil {
		return err
	}

	*o = EnterpriseAgent(varEnterpriseAgent)

	return err
}

type NullableEnterpriseAgent struct {
	value *EnterpriseAgent
	isSet bool
}

func (v NullableEnterpriseAgent) Get() *EnterpriseAgent {
	return v.value
}

func (v *NullableEnterpriseAgent) Set(val *EnterpriseAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAgent(val *EnterpriseAgent) *NullableEnterpriseAgent {
	return &NullableEnterpriseAgent{value: val, isSet: true}
}

func (v NullableEnterpriseAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


