// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package cache

import (
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	informersv1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	k8scache "k8s.io/client-go/tools/cache"
	"sync"
)

type SecretsCache struct {
	sync.RWMutex

	log logr.Logger

	// ConfigMap informer
	informer   k8scache.SharedInformer
	secretsMap map[string]bool
}

func NewSecretsCache(clientset kubernetes.Interface, log logr.Logger) *SecretsCache {
	sharedInformer := informersv1.NewSecretInformer(
		clientset,
		currentNamespace,
		informerResyncPeriod,
		k8scache.Indexers{},
	)
	return &SecretsCache{
		informer:   sharedInformer,
		log:        log.WithName("secrets"),
		secretsMap: make(map[string]bool),
	}
}

func (c *SecretsCache) Run(stopCh <-chan struct{}) {
	c.informer.AddEventHandler(k8scache.ResourceEventHandlerFuncs{
		UpdateFunc: func(orig, desired interface{}) {
			cm := desired.(*corev1.Secret)

			d, ok := c.secretsMap[cm.ObjectMeta.Name]

			if ok && d {
				// We need to reque the request
			}
		},
	})
	go c.informer.Run(stopCh)
}
