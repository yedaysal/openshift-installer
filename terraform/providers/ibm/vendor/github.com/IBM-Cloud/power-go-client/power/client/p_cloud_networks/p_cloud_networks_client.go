// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new p cloud networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PcloudNetworksDelete deletes a network
*/
func (a *Client) PcloudNetworksDelete(params *PcloudNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksDeleteOK), nil

}

/*
PcloudNetworksGet gets a network s current state information
*/
func (a *Client) PcloudNetworksGet(params *PcloudNetworksGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksGetOK), nil

}

/*
PcloudNetworksGetall gets all networks in this cloud instance
*/
func (a *Client) PcloudNetworksGetall(params *PcloudNetworksGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksGetallOK), nil

}

/*
PcloudNetworksPortsDelete deletes a network port
*/
func (a *Client) PcloudNetworksPortsDelete(params *PcloudNetworksPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPortsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksPortsDeleteOK), nil

}

/*
PcloudNetworksPortsGet gets a port s information
*/
func (a *Client) PcloudNetworksPortsGet(params *PcloudNetworksPortsGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPortsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksPortsGetOK), nil

}

/*
PcloudNetworksPortsGetall gets all ports for this network
*/
func (a *Client) PcloudNetworksPortsGetall(params *PcloudNetworksPortsGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPortsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksPortsGetallOK), nil

}

/*
PcloudNetworksPortsPost performs port addition deletion and listing
*/
func (a *Client) PcloudNetworksPortsPost(params *PcloudNetworksPortsPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPortsPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksPortsPostCreated), nil

}

/*
PcloudNetworksPortsPut updates a port s information
*/
func (a *Client) PcloudNetworksPortsPut(params *PcloudNetworksPortsPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPortsPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksPortsPutOK), nil

}

/*
PcloudNetworksPost creates a new network
*/
func (a *Client) PcloudNetworksPost(params *PcloudNetworksPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPostOK, *PcloudNetworksPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudNetworksPostOK:
		return value, nil, nil
	case *PcloudNetworksPostCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PcloudNetworksPut updates a network
*/
func (a *Client) PcloudNetworksPut(params *PcloudNetworksPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudNetworksPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.networks.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudNetworksPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}