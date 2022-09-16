// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package v20220315privatepreview

import (
	"github.com/project-radius/radius/pkg/armrpc/api/conv"
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/connectorrp/datamodel"

	"github.com/Azure/go-autorest/autorest/to"
)

// ConvertTo converts from the versioned DaprInvokeHttpRoute resource to version-agnostic datamodel.
func (src *DaprInvokeHTTPRouteResource) ConvertTo() (conv.DataModelInterface, error) {
	converted := &datamodel.DaprInvokeHttpRoute{
		TrackedResource: v1.TrackedResource{
			ID:       to.String(src.ID),
			Name:     to.String(src.Name),
			Type:     to.String(src.Type),
			Location: to.String(src.Location),
			Tags:     to.StringMap(src.Tags),
		},
		Properties: datamodel.DaprInvokeHttpRouteProperties{
			BasicResourceProperties: v1.BasicResourceProperties{
				Environment: to.String(src.Properties.Environment),
				Application: to.String(src.Properties.Application),
			},
			ProvisioningState: toProvisioningStateDataModel(src.Properties.ProvisioningState),
			AppId:             to.String(src.Properties.AppID),
		},
		InternalMetadata: v1.InternalMetadata{
			UpdatedAPIVersion: Version,
		},
	}

	if src.Properties.Recipe != nil {
		converted.Properties.Recipe.Name = to.String(src.Properties.Recipe.Name)
		if src.Properties.Recipe.Parameters != nil {
			converted.Properties.Recipe.Parameters = src.Properties.Recipe.Parameters
		}
	}

	return converted, nil
}

// ConvertFrom converts from version-agnostic datamodel to the versioned DaprInvokeHttpRoute resource.
func (dst *DaprInvokeHTTPRouteResource) ConvertFrom(src conv.DataModelInterface) error {
	daprHttpRoute, ok := src.(*datamodel.DaprInvokeHttpRoute)
	if !ok {
		return conv.ErrInvalidModelConversion
	}

	dst.ID = to.StringPtr(daprHttpRoute.ID)
	dst.Name = to.StringPtr(daprHttpRoute.Name)
	dst.Type = to.StringPtr(daprHttpRoute.Type)
	dst.SystemData = fromSystemDataModel(daprHttpRoute.SystemData)
	dst.Location = to.StringPtr(daprHttpRoute.Location)
	dst.Tags = *to.StringMapPtr(daprHttpRoute.Tags)
	dst.Properties = &DaprInvokeHTTPRouteProperties{
		Status: &ResourceStatus{
			OutputResources: v1.BuildExternalOutputResources(daprHttpRoute.Properties.Status.OutputResources),
		},
		ProvisioningState: fromProvisioningStateDataModel(daprHttpRoute.Properties.ProvisioningState),
		Environment:       to.StringPtr(daprHttpRoute.Properties.Environment),
		Application:       to.StringPtr(daprHttpRoute.Properties.Application),
		AppID:             to.StringPtr(daprHttpRoute.Properties.AppId),
	}

	if daprHttpRoute.Properties.Recipe.Name != "" {
		dst.Properties.Recipe = &Recipe{
			Name:       to.StringPtr(daprHttpRoute.Properties.Recipe.Name),
			Parameters: daprHttpRoute.Properties.Recipe.Parameters,
		}
	}

	return nil
}
