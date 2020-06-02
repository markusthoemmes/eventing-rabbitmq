/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	"github.com/apache/camel-k/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CamelV1Interface interface {
	RESTClient() rest.Interface
	BuildsGetter
	IntegrationsGetter
	IntegrationKitsGetter
	IntegrationPlatformsGetter
}

// CamelV1Client is used to interact with features provided by the camel.apache.org group.
type CamelV1Client struct {
	restClient rest.Interface
}

func (c *CamelV1Client) Builds(namespace string) BuildInterface {
	return newBuilds(c, namespace)
}

func (c *CamelV1Client) Integrations(namespace string) IntegrationInterface {
	return newIntegrations(c, namespace)
}

func (c *CamelV1Client) IntegrationKits(namespace string) IntegrationKitInterface {
	return newIntegrationKits(c, namespace)
}

func (c *CamelV1Client) IntegrationPlatforms(namespace string) IntegrationPlatformInterface {
	return newIntegrationPlatforms(c, namespace)
}

// NewForConfig creates a new CamelV1Client for the given config.
func NewForConfig(c *rest.Config) (*CamelV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CamelV1Client{client}, nil
}

// NewForConfigOrDie creates a new CamelV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CamelV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CamelV1Client for the given RESTClient.
func New(c rest.Interface) *CamelV1Client {
	return &CamelV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CamelV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
