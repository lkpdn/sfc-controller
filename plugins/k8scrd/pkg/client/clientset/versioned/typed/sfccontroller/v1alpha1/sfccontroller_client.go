// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ligato/sfc-controller/plugins/k8scrd/pkg/apis/sfccontroller/v1alpha1"
	"github.com/ligato/sfc-controller/plugins/k8scrd/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type SfccontrollerV1alpha1Interface interface {
	RESTClient() rest.Interface
	NetworkNodesGetter
}

// SfccontrollerV1alpha1Client is used to interact with features provided by the sfccontroller.ligato.github.com group.
type SfccontrollerV1alpha1Client struct {
	restClient rest.Interface
}

func (c *SfccontrollerV1alpha1Client) NetworkNodes(namespace string) NetworkNodeInterface {
	return newNetworkNodes(c, namespace)
}

// NewForConfig creates a new SfccontrollerV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*SfccontrollerV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SfccontrollerV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new SfccontrollerV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SfccontrollerV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SfccontrollerV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SfccontrollerV1alpha1Client {
	return &SfccontrollerV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SfccontrollerV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
