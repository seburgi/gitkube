/*
Copyright 2018 Hasura.io

*/

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/hasura/gitkube/pkg/apis/gitkube.sh/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RemoteLister helps list Remotes.
type RemoteLister interface {
	// List lists all Remotes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Remote, err error)
	// Remotes returns an object that can list and get Remotes.
	Remotes(namespace string) RemoteNamespaceLister
	RemoteListerExpansion
}

// remoteLister implements the RemoteLister interface.
type remoteLister struct {
	indexer cache.Indexer
}

// NewRemoteLister returns a new RemoteLister.
func NewRemoteLister(indexer cache.Indexer) RemoteLister {
	return &remoteLister{indexer: indexer}
}

// List lists all Remotes in the indexer.
func (s *remoteLister) List(selector labels.Selector) (ret []*v1alpha1.Remote, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Remote))
	})
	return ret, err
}

// Remotes returns an object that can list and get Remotes.
func (s *remoteLister) Remotes(namespace string) RemoteNamespaceLister {
	return remoteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RemoteNamespaceLister helps list and get Remotes.
type RemoteNamespaceLister interface {
	// List lists all Remotes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Remote, err error)
	// Get retrieves the Remote from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Remote, error)
	RemoteNamespaceListerExpansion
}

// remoteNamespaceLister implements the RemoteNamespaceLister
// interface.
type remoteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Remotes in the indexer for a given namespace.
func (s remoteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Remote, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Remote))
	})
	return ret, err
}

// Get retrieves the Remote from the indexer for a given namespace and name.
func (s remoteNamespaceLister) Get(name string) (*v1alpha1.Remote, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("remote"), name)
	}
	return obj.(*v1alpha1.Remote), nil
}