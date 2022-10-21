/*
Copyright 2022 The KServe Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/kserve/kserve/pkg/clientv1alpha1/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterServingRuntimes returns a ClusterServingRuntimeInformer.
	ClusterServingRuntimes() ClusterServingRuntimeInformer
	// InferenceGraphs returns a InferenceGraphInformer.
	InferenceGraphs() InferenceGraphInformer
	// ServingRuntimes returns a ServingRuntimeInformer.
	ServingRuntimes() ServingRuntimeInformer
	// TrainedModels returns a TrainedModelInformer.
	TrainedModels() TrainedModelInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterServingRuntimes returns a ClusterServingRuntimeInformer.
func (v *version) ClusterServingRuntimes() ClusterServingRuntimeInformer {
	return &clusterServingRuntimeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InferenceGraphs returns a InferenceGraphInformer.
func (v *version) InferenceGraphs() InferenceGraphInformer {
	return &inferenceGraphInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServingRuntimes returns a ServingRuntimeInformer.
func (v *version) ServingRuntimes() ServingRuntimeInformer {
	return &servingRuntimeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TrainedModels returns a TrainedModelInformer.
func (v *version) TrainedModels() TrainedModelInformer {
	return &trainedModelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}