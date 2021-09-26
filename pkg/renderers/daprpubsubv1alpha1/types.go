// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package daprpubsubv1alpha1

import (
	"github.com/Azure/radius/pkg/azure/azresources"
)

const Kind = "dapr.io/PubSubTopic@v1alpha1"

var TopicResourceType = azresources.KnownType{
	Types: []azresources.ResourceType{
		{
			Type: azresources.ServiceBusNamespaces,
			Name: "*",
		},
		{
			Type: azresources.ServiceBusNamespacesTopics,
			Name: "*",
		},
	},
}

// DaprPubSubTopicComponent is the definition of the container component
type DaprPubSubTopicComponent struct {
	Name     string                   `json:"name"`
	Kind     string                   `json:"kind"`
	Config   DaprPubSubConfig         `json:"config,omitempty"`
	Run      map[string]interface{}   `json:"run,omitempty"`
	Uses     []map[string]interface{} `json:"uses,omitempty"`
	Bindings []map[string]interface{} `json:"bindings,omitempty"`
	Traits   []map[string]interface{} `json:"traits,omitempty"`
}

// DaprPubSubConfig is the defintion of the config section
type DaprPubSubConfig struct {
	Kind    string `json:"kind"`
	Managed bool   `json:"managed"`
	// The name of the Dapr pubsub Component
	Name     string `json:"name"`
	Resource string `json:"resource"`
	Topic    string `json:"topic"`
}
