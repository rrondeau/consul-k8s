package v1alpha1

import (
	"context"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/hashicorp/consul-k8s/api/common"
	capi "github.com/hashicorp/consul/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func NewServiceResolverValidator(client client.Client, consulClient *capi.Client, logger logr.Logger) *serviceResolverValidator {
	return &serviceResolverValidator{
		Client:       client,
		ConsulClient: consulClient,
		Logger:       logger,
	}
}

type serviceResolverValidator struct {
	client.Client
	ConsulClient *capi.Client
	Logger       logr.Logger
	decoder      *admission.Decoder
}

// NOTE: The path value in the below line is the path to the webhook.
// If it is updated, run code-gen, update subcommand/controller/command.go
// and the consul-helm value for the path to the webhook.
//
// NOTE: The below line cannot be combined with any other comment. If it is
// it will break the code generation.
//
// +kubebuilder:webhook:verbs=create;update,path=/mutate-v1alpha1-serviceresolver,mutating=true,failurePolicy=fail,groups=consul.hashicorp.com,resources=serviceresolvers,versions=v1alpha1,name=mutate-serviceresolver.consul.hashicorp.com

func (v *serviceResolverValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	var svcResolver ServiceResolver
	err := v.decoder.Decode(req, &svcResolver)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	return common.ValidateConfigEntry(ctx,
		req,
		v.Logger,
		v,
		&svcResolver,
		"ServiceResolver")
}

func (v *serviceResolverValidator) List(ctx context.Context) ([]common.ConfigEntryResource, error) {
	var svcResolverList ServiceResolverList
	if err := v.Client.List(ctx, &svcResolverList); err != nil {
		return nil, err
	}
	var entries []common.ConfigEntryResource
	for _, item := range svcResolverList.Items {
		entries = append(entries, common.ConfigEntryResource(&item))
	}
	return entries, nil
}

func (v *serviceResolverValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}