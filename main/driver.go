package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/prashantgolash/kubecrd/v1alpha1"

	//"github.com/prashantgolash/kubecrd/v1alpha1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"time"
)

var (
	// Set during build
	version string

	proxyURL = flag.String("proxy", "",
		`If specified, it is assumed that a kubctl proxy server is running on the
		given url and creates a proxy client. In case it is not given InCluster
		kubernetes setup will be used`)
)

func main() {

	flag.Parse()
	var err error

	var config *rest.Config
	if *proxyURL != "" {
		config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{},
			&clientcmd.ConfigOverrides{
				ClusterInfo: clientcmdapi.Cluster{
					Server: *proxyURL,
				},
			}).ClientConfig()
		if err != nil {
			glog.Fatalf("error creating client configuration: %v", err)
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			glog.Fatalf("error creating client configuration: %v", err)
		}
	}

	kubeClient, err := apiextension.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Failed to create client: %v", err)
	}
	// Create the CRD
	err = v1alpha1.CreateCRD(kubeClient)
	if err != nil {
		glog.Fatalf("Failed to create crd: %v", err)
	}

	// Wait for the CRD to be created before we use it.
	time.Sleep(5 * time.Second)

	// Create a new clientset which include our CRD schema
	crdclient, err := v1alpha1.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Create a new Respool object
	// Create a new Respool object
	RespoolCRD := &v1alpha1.Respool {
		ObjectMeta: meta_v1.ObjectMeta{
			Name:   "respool-obj",
			Labels: map[string]string{"mylabel": "test"},
		},

		Spec: v1alpha1.RespoolSpec {
			RespoolID:   "1",
			RespoolConfig:    "respool-config",
			Owner: "new-owner",
			CreationTime: time.Now(),
			UpdateTime: time.Now(),
		},

		Status: v1alpha1.RespoolStatus {
			State:   "created",
			Message: "Created, not processed yet",
		},
	}
	// Create the Respool object we create above in the k8s cluster
	resp, err := crdclient.Respools("default").Create(RespoolCRD)
	if err != nil {
		fmt.Printf("error while creating object: %v\n", err)
	} else {
		fmt.Printf("object created: %v\n", resp)
	}

	obj, err := crdclient.Respools("default").Get(RespoolCRD.ObjectMeta.Name)
	if err != nil {
		glog.Infof("error while getting the object %v\n", err)
	}
	fmt.Printf("Respool Objects Found: \n%v\n", obj)
	select {}
}