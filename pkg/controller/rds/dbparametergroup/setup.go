package dbparametergroup

import (
	"context"
	"fmt"

	svcsdk "github.com/aws/aws-sdk-go/service/rds"
	svcsdkapi "github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"regexp"

	svcapitypes "github.com/crossplane/provider-aws/apis/rds/v1alpha1"
	"github.com/crossplane/provider-aws/apis/v1alpha1"
	awsclients "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/crossplane/provider-aws/pkg/features"
)

// SetupDBParameterGroup adds a controller that reconciles DBParametergroup.
func SetupDBParameterGroup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(svcapitypes.DBParameterGroupGroupKind)
	opts := []option{
		func(e *external) {
			e.preCreate = preCreate
			e.preObserve = preObserve
			e.preUpdate = preUpdate
			e.preDelete = preDelete
			e.postObserve = postObserve
			c := &custom{client: e.client, kube: e.kube}
			e.isUpToDate = c.isUpToDate
		},
	}

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), v1alpha1.StoreConfigGroupVersionKind))
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&svcapitypes.DBParameterGroup{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.DBParameterGroupGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithPollInterval(o.PollInterval),
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
			managed.WithConnectionPublishers(cps...)))
}

type custom struct {
	kube   client.Client
	client svcsdkapi.RDSAPI
}

func preObserve(_ context.Context, cr *svcapitypes.DBParameterGroup, obj *svcsdk.DescribeDBParameterGroupsInput) error {
	obj.DBParameterGroupName = awsclients.String(meta.GetExternalName(cr))
	return nil
}

func postObserve(_ context.Context, cr *svcapitypes.DBParameterGroup, obj *svcsdk.DescribeDBParameterGroupsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	cr.SetConditions(xpv1.Available())
	return obs, err
}

func preCreate(_ context.Context, cr *svcapitypes.DBParameterGroup, obj *svcsdk.CreateDBParameterGroupInput) error {
	obj.DBParameterGroupName = awsclients.String(meta.GetExternalName(cr))
	return nil
}

func preUpdate(_ context.Context, cr *svcapitypes.DBParameterGroup, obj *svcsdk.ModifyDBParameterGroupInput) error {
	obj.DBParameterGroupName = awsclients.String(meta.GetExternalName(cr))
	obj.Parameters = make([]*svcsdk.Parameter, len(cr.Spec.ForProvider.Parameters))

	for i, v := range cr.Spec.ForProvider.Parameters {
		// check if mandatory parameters are set (ApplyMethod, ParameterName, ParameterValue)
		if (v.ApplyMethod == nil) || (v.ParameterName == nil) || (v.ParameterValue == nil) {
			return errors.New("ApplyMethod, ParameterName and ParameterValue are mandatory fields and can not be nil")
		}
		obj.Parameters[i] = &svcsdk.Parameter{
			AllowedValues:        v.AllowedValues,
			ApplyMethod:          v.ApplyMethod,
			ApplyType:            v.ApplyType,
			DataType:             v.DataType,
			Description:          v.Description,
			IsModifiable:         v.IsModifiable,
			MinimumEngineVersion: v.MinimumEngineVersion,
			ParameterName:        v.ParameterName,
			ParameterValue:       v.ParameterValue,
			Source:               v.Source,
			SupportedEngineModes: v.SupportedEngineModes,
		}
	}
	return nil
}

func preDelete(_ context.Context, cr *svcapitypes.DBParameterGroup, obj *svcsdk.DeleteDBParameterGroupInput) (bool, error) {
	obj.DBParameterGroupName = awsclients.String(meta.GetExternalName(cr))
	return false, nil
}

func (e *custom) isUpToDate(cr *svcapitypes.DBParameterGroup, obj *svcsdk.DescribeDBParameterGroupsOutput) (bool, error) {
	// TODO(Dkaykay): We need isUpToDate to have context.
	ctx := context.TODO()
	results, err := e.getCurrentDBParameters(ctx, cr)
	if err != nil {
		return false, err
	}
	observed := make(map[string]svcsdk.Parameter, len(results))
	for _, p := range results {
		observed[awsclients.StringValue(p.ParameterName)] = *p
	}
	// compare CR with currently set Parameters
	for _, v := range cr.Spec.ForProvider.Parameters {
		existing, ok := observed[awsclients.StringValue(v.ParameterName)]
		if !ok {
			fmt.Println("missing ", awsclients.StringValue(v.ParameterName))
			return false, nil
		}
		normalize := func(s string) string { return s }
		if awsclients.StringValue(existing.DataType) == "list" {
			normalize = func(s string) string {
				return regexp.MustCompile(", +").ReplaceAllString(s, ",")
			}
		}
		switch {
		case normalize(awsclients.StringValue(existing.ParameterValue)) != normalize(awsclients.StringValue(v.ParameterValue)):
			fmt.Printf("diff DBParameterValue %v %v %v", awsclients.StringValue(v.ParameterName), awsclients.StringValue(existing.ParameterValue), awsclients.StringValue(v.ParameterValue))
			return false, nil
		case awsclients.StringValue(existing.ApplyMethod) != awsclients.StringValue(v.ApplyMethod):
			fmt.Printf("diff DBParameter Apply Method %v, %v, %v", awsclients.StringValue(v.ParameterName), awsclients.StringValue(existing.ApplyMethod), awsclients.StringValue(v.ApplyMethod))
			return false, nil
		}
	}
	return true, err
}

func (e *custom) getCurrentDBParameters(ctx context.Context, cr *svcapitypes.DBParameterGroup) ([]*svcsdk.Parameter, error) {
	input := &svcsdk.DescribeDBParametersInput{
		DBParameterGroupName: awsclients.String(meta.GetExternalName(cr)),
		MaxRecords:           awsclients.Int64(100),
	}
	var results []*svcsdk.Parameter
	err := e.client.DescribeDBParametersPagesWithContext(ctx, input, func(page *svcsdk.DescribeDBParametersOutput, lastPage bool) bool {
		results = append(results, page.Parameters...)
		return !lastPage
	})
	if err != nil {
		return results, err
	}
	return results, nil
}
