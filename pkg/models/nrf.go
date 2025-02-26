package models

import (
	"time"
)

// NfType represents the type of Network Function
type NfType string

const (
	// NfTypeAMF represents Access and Mobility Management Function
	NfTypeAMF NfType = "AMF"
	
	// NfTypeSMF represents Session Management Function
	NfTypeSMF NfType = "SMF"
	
	// NfTypeUPF represents User Plane Function
	NfTypeUPF NfType = "UPF"
	
	// NfTypePCF represents Policy Control Function
	NfTypePCF NfType = "PCF"
	
	// NfTypeUDM represents Unified Data Management
	NfTypeUDM NfType = "UDM"
	
	// NfTypeAUSF represents Authentication Server Function
	NfTypeAUSF NfType = "AUSF"
	
	// NfTypeNRF represents Network Repository Function
	NfTypeNRF NfType = "NRF"
	
	// NfTypeNSSF represents Network Slice Selection Function
	NfTypeNSSF NfType = "NSSF"
)

// NfStatus represents the status of a Network Function
type NfStatus string

const (
	// NfStatusRegistered means the NF is registered with the NRF
	NfStatusRegistered NfStatus = "REGISTERED"
	
	// NfStatusDeregistered means the NF is deregistered with the NRF
	NfStatusDeregistered NfStatus = "DEREGISTERED"
	
	// NfStatusSuspended means the NF is temporarily suspended
	NfStatusSuspended NfStatus = "SUSPENDED"
)

// NfProfile represents a Network Function profile registered with the NRF
type NfProfile struct {
	// Unique identifier of the NF instance
	NfInstanceID string `json:"nfInstanceId"`
	
	// Type of the Network Function
	NfType NfType `json:"nfType"`
	
	// Name of the NF instance
	NfInstanceName string `json:"nfInstanceName,omitempty"`
	
	// Status of the NF instance
	NfStatus NfStatus `json:"nfStatus"`
	
	// Heart-beat timer in seconds
	HeartbeatTimer int `json:"heartbeatTimer,omitempty"`
	
	// IPv4 addresses of the Network Function
	IPv4Addresses []string `json:"ipv4Addresses,omitempty"`
	
	// IPv6 addresses of the Network Function
	IPv6Addresses []string `json:"ipv6Addresses,omitempty"`
	
	// FQDN of the Network Function
	FQDN string `json:"fqdn,omitempty"`
	
	// Priority of the NF instance (lower value means higher priority)
	Priority int `json:"priority,omitempty"`
	
	// Capacity of the NF instance (0-100)
	Capacity int `json:"capacity,omitempty"`
	
	// Load of the NF instance (0-100)
	Load int `json:"load,omitempty"`
	
	// Locality of the NF instance (e.g., geographic location)
	Locality string `json:"locality,omitempty"`
	
	// NF services exposed by the NF instance
	NfServices []NfService `json:"nfServices,omitempty"`
	
	// Types of NF that can access this NF
	AllowedNfTypes []NfType `json:"allowedNfTypes,omitempty"`
	
	// Time when the NF instance was registered
	RegisterTime time.Time `json:"registerTime,omitempty"`
	
	// Last time a heartbeat was received
	LastHeartbeatTime time.Time `json:"lastHeartbeatTime,omitempty"`
}

// NfService represents a service exposed by a Network Function
type NfService struct {
	// Service instance ID
	ServiceInstanceID string `json:"serviceInstanceId"`
	
	// Service name
	ServiceName string `json:"serviceName"`
	
	// Versions supported by the service
	Versions []NfServiceVersion `json:"versions"`
	
	// Service scheme (e.g., "http", "https")
	Scheme string `json:"scheme"`
	
	// FQDN of the service
	FQDN string `json:"fqdn,omitempty"`
	
	// IP addresses of the service
	IPAddresses []string `json:"ipAddresses,omitempty"`
	
	// Port number of the service
	Port int `json:"port,omitempty"`
	
	// URI prefix of the service
	URIPrefix string `json:"uriPrefix,omitempty"`
	
	// Priority of the service
	Priority int `json:"priority,omitempty"`
}

// NfServiceVersion represents a version of an NF service
type NfServiceVersion struct {
	// API version
	APIVersion string `json:"apiVersion"`
	
	// API full version
	APIFullVersion string `json:"apiFullVersion"`
	
	// API URI prefix
	URIPrefix string `json:"uriPrefix,omitempty"`
}

// NfRegistrationData represents data sent to register an NF with the NRF
type NfRegistrationData struct {
	// NF Profile
	NfProfile NfProfile `json:"nfProfile"`
}

// NfRegistrationResponse represents the response to an NF registration
type NfRegistrationResponse struct {
	// NF Instance ID
	NfInstanceID string `json:"nfInstanceId"`
	
	// Heart-beat timer in seconds
	HeartbeatTimer int `json:"heartbeatTimer"`
}

// NfDiscoveryRequest represents a request to discover NFs
type NfDiscoveryRequest struct {
	// Target NF type to discover
	TargetNfType NfType `json:"targetNfType"`
	
	// Requester NF type
	RequesterNfType NfType `json:"requesterNfType"`
	
	// Preferred locality
	PreferredLocality string `json:"preferredLocality,omitempty"`
	
	// Minimum capacity required
	MinCapacity int `json:"minCapacity,omitempty"`
	
	// Maximum load allowed
	MaxLoad int `json:"maxLoad,omitempty"`
	
	// Required service names
	RequiredServiceNames []string `json:"requiredServiceNames,omitempty"`
	
	// Limit of NF instances to return
	Limit int `json:"limit,omitempty"`
}

// NfDiscoveryResponse represents the response to an NF discovery request
type NfDiscoveryResponse struct {
	// Validity period of the response in seconds
	ValidityPeriod int `json:"validityPeriod,omitempty"`
	
	// Discovered NF instances
	NfInstances []NfProfile `json:"nfInstances"`
}
