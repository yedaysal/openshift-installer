// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new p cloud v p n connections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud v p n connections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PcloudVpnconnectionsDelete deletes v p n connection

Delete VPN Connection (by its identifier)
*/
func (a *Client) PcloudVpnconnectionsDelete(params *PcloudVpnconnectionsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsDeleteAccepted), nil

}

/*
PcloudVpnconnectionsGet gets v p n connection

Get a VPN Connection
*/
func (a *Client) PcloudVpnconnectionsGet(params *PcloudVpnconnectionsGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsGetOK), nil

}

/*
PcloudVpnconnectionsGetall gets all v p n connections

Get all VPN Connections
*/
func (a *Client) PcloudVpnconnectionsGetall(params *PcloudVpnconnectionsGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsGetallOK), nil

}

/*
PcloudVpnconnectionsNetworksDelete detaches network

Detach network from a specific VPN Connection
*/
func (a *Client) PcloudVpnconnectionsNetworksDelete(params *PcloudVpnconnectionsNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsNetworksDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsNetworksDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.networks.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsNetworksDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsNetworksDeleteAccepted), nil

}

/*
PcloudVpnconnectionsNetworksGet gets attached networks

Get a list of network IDs attached to a VPN Connection
*/
func (a *Client) PcloudVpnconnectionsNetworksGet(params *PcloudVpnconnectionsNetworksGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsNetworksGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsNetworksGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.networks.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsNetworksGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsNetworksGetOK), nil

}

/*
PcloudVpnconnectionsNetworksPut attaches network

Attach a network to a VPN Connection
*/
func (a *Client) PcloudVpnconnectionsNetworksPut(params *PcloudVpnconnectionsNetworksPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsNetworksPutAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsNetworksPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.networks.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsNetworksPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsNetworksPutAccepted), nil

}

/*
PcloudVpnconnectionsPeersubnetsDelete detaches peer subnet

Detach peer subnet from a VPN Connection
*/
func (a *Client) PcloudVpnconnectionsPeersubnetsDelete(params *PcloudVpnconnectionsPeersubnetsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsPeersubnetsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsPeersubnetsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.peersubnets.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsPeersubnetsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsPeersubnetsDeleteOK), nil

}

/*
PcloudVpnconnectionsPeersubnetsGet gets peer subnets

Get a list of peer subnets attached to a specific VPN Connection
*/
func (a *Client) PcloudVpnconnectionsPeersubnetsGet(params *PcloudVpnconnectionsPeersubnetsGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsPeersubnetsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsPeersubnetsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.peersubnets.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsPeersubnetsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsPeersubnetsGetOK), nil

}

/*
PcloudVpnconnectionsPeersubnetsPut attaches peer subnet

Attach peer subnet to a VPN Connection
*/
func (a *Client) PcloudVpnconnectionsPeersubnetsPut(params *PcloudVpnconnectionsPeersubnetsPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsPeersubnetsPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsPeersubnetsPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.peersubnets.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsPeersubnetsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsPeersubnetsPutOK), nil

}

/*
PcloudVpnconnectionsPost creates v p n connection

Create a new VPN Connection
*/
func (a *Client) PcloudVpnconnectionsPost(params *PcloudVpnconnectionsPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsPostAccepted), nil

}

/*
PcloudVpnconnectionsPut updates v p n connection

update a VPN Connection (by its identifier)
*/
func (a *Client) PcloudVpnconnectionsPut(params *PcloudVpnconnectionsPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudVpnconnectionsPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVpnconnectionsPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.vpnconnections.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVpnconnectionsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudVpnconnectionsPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}