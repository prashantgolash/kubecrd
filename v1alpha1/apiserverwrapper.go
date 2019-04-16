package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func (c *RespoolV1Alpha1Client) Respools(namespace string) RespoolInterface {
	return &respoolclient{
		client: c.restClient,
		ns:     namespace,
	}
}

type RespoolV1Alpha1Client struct {
	restClient rest.Interface
}

type RespoolInterface interface {
	Create(obj *Respool) (*Respool, error)
	Update(obj *Respool) (*Respool, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string) (*Respool, error)
}

type respoolclient struct {
	client rest.Interface
	ns     string
}

func (c *respoolclient) Create(obj *Respool) (*Respool, error) {
	result := &Respool{}
	err := c.client.Post().
		Namespace(c.ns).Resource("respools").
		Body(obj).Do().Into(result)
	return result, err
}

func (c *respoolclient) Update(obj *Respool) (*Respool, error) {
	result := &Respool{}
	err := c.client.Put().
		Namespace(c.ns).Resource("respools").
		Body(obj).Do().Into(result)
	return result, err
}

func (c *respoolclient) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).Resource("respools").
		Name(name).Body(options).Do().
		Error()
}

func (c *respoolclient) Get(name string) (*Respool, error) {
	result := &Respool{}
	err := c.client.Get().
		Namespace(c.ns).Resource("respools").
		Name(name).Do().Into(result)
	return result, err
}
