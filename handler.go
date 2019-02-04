package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/raker22/k8s-blob-replica-set-resource/pkg/apis/blobreplicaset/v1alpha1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	log.Info("TestHandler.ObjectCreated")
	// assert the type to a BlobReplicaSet object to pull out relevant data
	blobReplicaSet := obj.(*v1alpha1.BlobReplicaSet)
	log.Infof("    ResourceVersion: %s", blobReplicaSet.ObjectMeta.ResourceVersion)
	log.Infof("    Name: %s", blobReplicaSet.ObjectMeta.Name)
	log.Infof("    Replicas: %d", *blobReplicaSet.Spec.Replicas)
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}