package controllers

import (
	"context"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/TykTechnologies/tyk-operator/api/model"
	"github.com/TykTechnologies/tyk-operator/api/v1alpha1"
	tykClient "github.com/TykTechnologies/tyk-operator/pkg/client"
	"github.com/TykTechnologies/tyk-operator/pkg/environmet"
	"github.com/go-logr/logr"
	"github.com/mitchellh/hashstructure/v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// hashOptions is a common hashing options to be used.
var hashOptions = hashstructure.HashOptions{ZeroNil: true}

// calculateHashes calculates hashes of given two interfaces. Returns empty string for hash
// if any error occurs during calculation.
func calculateHashes(i1, i2 interface{}) (hash1, hash2 string) {
	h1, err1 := hashstructure.Hash(i1, hashstructure.FormatV2, &hashOptions)
	if err1 == nil {
		hash1 = strconv.FormatUint(h1, 10)

		h2, err2 := hashstructure.Hash(i2, hashstructure.FormatV2, &hashOptions)
		if err2 == nil {
			return hash1, strconv.FormatUint(h2, 10)
		}
	}

	return
}

// changed returns if the given hash equals to hash of given interface.
func changed(latestHash string, i1 interface{}) bool {
	api1Hash, err := hashstructure.Hash(i1, hashstructure.FormatV2, &hashOptions)
	if err != nil {
		return true
	}

	return latestHash != strconv.FormatUint(api1Hash, 10)
}

// isSameApiDefinition returns if given ApiDefinition CR equals ApiDefinition resource stored in Tyk, by
// comparing hash of latest observed hash of CR and Policy resource. If there is no change detected,
// subsequent update API calls to Tyk Gateway or Dashboard will be omitted.
func isSameApiDefinition(desired *v1alpha1.ApiDefinition, apiDefOnTyk *model.APIDefinitionSpec) bool {
	if desired == nil && apiDefOnTyk == nil {
		return true
	} else if desired == nil || apiDefOnTyk == nil {
		return false
	}

	k8sChanged := changed(desired.Status.LatestCRDHash, desired.Spec.APIDefinitionSpec)
	tykChanged := changed(desired.Status.LatestTykHash, apiDefOnTyk)

	return !k8sChanged && !tykChanged
}

const (
	proTagName = "pro"
	ossTagName = "oss"
)

func hasherTagName(ctx context.Context) string {
	r := tykClient.GetContext(ctx)
	if r.Env.Mode == "pro" {
		return proTagName
	}

	return ossTagName
}

// isSameSecurityPolicy calculates the hash of two SecurityPolicySpec structs and returns true
// if hashes are equal. It is used to prevent calling extra Update requests to Tyk Gateway
// if there is no changes on resources.
func isSameSecurityPolicy(ctx context.Context, policy1, policy2 *v1alpha1.SecurityPolicySpec) bool {
	myHashOptions := hashstructure.HashOptions{ZeroNil: true, TagName: hasherTagName(ctx)}

	pol1Hash, err1 := hashstructure.Hash(policy1, hashstructure.FormatV2, &myHashOptions)
	if err1 == nil {
		pol2Hash, err2 := hashstructure.Hash(policy2, hashstructure.FormatV2, &myHashOptions)
		if err2 == nil {
			return pol1Hash == pol2Hash
		}
	}

	return false
}

// containsString is a helper function to check string exists in a slice of strings.
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}

	return false
}

// addTarget adds given target to given slice if the slice does not contain the target.
func addTarget(slice []model.Target, target model.Target) (result []model.Target) {
	for _, item := range slice {
		if item == target {
			return slice
		}
	}

	return append(slice, target)
}

func removeTarget(slice []model.Target, s model.Target) (result []model.Target) {
	for _, item := range slice {
		if item.Equal(s) {
			continue
		}

		result = append(result, item)
	}

	return
}

// EncodeNS encodes given decoded string based on base64.
func EncodeNS(decoded string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(decoded))
}

// HttpContext creates a context.Context for Tyk API Client.
func HttpContext(
	ctx context.Context,
	rClient client.Client,
	e environmet.Env,
	object client.Object,
	log logr.Logger,
) (environmet.Env, context.Context, error) {
	get := func(opCtxRef *model.Target) error {
		if opCtxRef == nil {
			// To handle the case where operator context was used previously
			// but was removed in update operation
			if err := updateOperatorContextStatus(ctx, rClient, object, log, opCtxRef); err != nil {
				log.Error(err, "Failed to update status of operator contexts")
				return err
			}

			return nil
		}

		// If namespace is not specified in contextDef, use default namespace
		if opCtxRef.Namespace == "" {
			log.Info("Context namespace is not specified, using default")

			opCtxRef.Namespace = "default"
		}

		log.Info("Detected context for resource")

		env, err := GetContext(
			ctx, object.GetNamespace(), rClient, opCtxRef, log,
		)
		if err != nil {
			log.Error(err, "Failed to get context", "contextRef", opCtxRef.String())

			return err
		}

		log.Info("Successful acquired context", "contextRef", opCtxRef.String())

		e.Environment = *env.Spec.Env

		if err := updateOperatorContextStatus(ctx, rClient, object, log, opCtxRef); err != nil {
			log.Error(err, "Failed to update status of operator contexts")
			return err
		}

		return nil
	}

	var err error
	switch o := object.(type) {
	case *v1alpha1.ApiDefinition:
		err = get(o.Spec.Context)
	case *v1alpha1.SecurityPolicy:
		err = get(o.Spec.Context)
	case *v1alpha1.PortalAPICatalogue:
		err = get(o.Spec.Context)
	case *v1alpha1.APIDescription:
		err = get(o.Spec.Context)
	case *v1alpha1.PortalConfig:
		err = get(o.Spec.Context)
	}

	if err != nil {
		return environmet.Env{}, nil, err
	}

	return e, tykClient.SetContext(ctx, tykClient.Context{
		Env: e,
		Log: log,
	}), nil
}

// updateOperatorContextStatus updates links defined in the status of OperatorContext identified by ctxRef
// and Custom Resource identified by object.
func updateOperatorContextStatus(
	ctx context.Context,
	client client.Client,
	object client.Object,
	log logr.Logger,
	ctxRef *model.Target,
) error {
	objectTarget := model.Target{
		Name:      object.GetName(),
		Namespace: object.GetNamespace(),
	}

	// Remove link from other OperatorContext, if any,
	// as we do not know if object was referencing to different context previously.
	var opCtxList v1alpha1.OperatorContextList
	if err := client.List(ctx, &opCtxList); err != nil {
		return err
	}

	switch object.(type) {
	case *v1alpha1.ApiDefinition:
		for i := 0; i < len(opCtxList.Items); i++ {
			// do not remove link if ApiDefinition is still referring to same context and is not marked for deletion.
			if ctxRef != nil &&
				opCtxList.Items[i].Name == ctxRef.Name && opCtxList.Items[i].Namespace == ctxRef.Namespace &&
				object.GetDeletionTimestamp().IsZero() {
				continue
			}

			opCtxList.Items[i].Status.RemoveLinkedAPIDefinition(objectTarget)

			err := client.Status().Update(ctx, &opCtxList.Items[i])
			if err != nil {
				log.Error(
					err,
					"Failed to remove link of APIDefinition from operator context",
					"operatorContext", opCtxList.Items[i].Name, "ApiDefinition", objectTarget.String(),
				)

				return err
			}
		}
	case *v1alpha1.SecurityPolicy:
		for i := 0; i < len(opCtxList.Items); i++ {
			// do not remove link if SecurityPolicy is still referring to context and is not marked for deletion.
			if ctxRef != nil && opCtxList.Items[i].Name == ctxRef.Name &&
				opCtxList.Items[i].Namespace == ctxRef.Namespace && object.GetDeletionTimestamp().IsZero() {
				continue
			}

			opCtxList.Items[i].Status.RemoveLinkedSecurityPolicies(objectTarget)

			err := client.Status().Update(ctx, &opCtxList.Items[i])
			if err != nil {
				return err
			}
		}
	case *v1alpha1.PortalAPICatalogue:
		for i := 0; i < len(opCtxList.Items); i++ {
			// do not remove link if PortalAPICatalogue is still referring to context and is not marked for deletion.
			if ctxRef != nil && opCtxList.Items[i].Name == ctxRef.Name && opCtxList.Items[i].Namespace == ctxRef.Namespace &&
				object.GetDeletionTimestamp().IsZero() {
				continue
			}

			opCtxList.Items[i].Status.RemoveLinkedPortalAPICatalogues(objectTarget)

			err := client.Status().Update(ctx, &opCtxList.Items[i])
			if err != nil {
				return err
			}
		}
	case *v1alpha1.APIDescription:
		for i := 0; i < len(opCtxList.Items); i++ {
			// do not remove link if APIDescription is still referring to context and is not marked for deletion.
			if ctxRef != nil && opCtxList.Items[i].Name == ctxRef.Name &&
				opCtxList.Items[i].Namespace == ctxRef.Namespace && object.GetDeletionTimestamp().IsZero() {
				continue
			}

			opCtxList.Items[i].Status.RemoveLinkedApiDescriptions(objectTarget)

			err := client.Status().Update(ctx, &opCtxList.Items[i])
			if err != nil {
				return err
			}
		}
	case *v1alpha1.PortalConfig:
		for i := 0; i < len(opCtxList.Items); i++ {
			// do not remove link if PortalConfig is still referring to context and is not marked for deletion.
			if ctxRef != nil && opCtxList.Items[i].Name == ctxRef.Name &&
				opCtxList.Items[i].Namespace == ctxRef.Namespace && object.GetDeletionTimestamp().IsZero() {
				continue
			}

			opCtxList.Items[i].Status.RemoveLinkedPortalConfig(objectTarget)

			err := client.Status().Update(ctx, &opCtxList.Items[i])
			if err != nil {
				return err
			}
		}
	}

	// Add reference to the referred operator context
	// only if object is not marked for deletion
	if object.GetDeletionTimestamp().IsZero() && ctxRef != nil {
		// add reference to operator context
		log.Info("Adding link to apiContext", "key", ctxRef.String())

		var operatorContext v1alpha1.OperatorContext

		key := types.NamespacedName{Name: ctxRef.Name, Namespace: ctxRef.Namespace}

		if err := client.Get(ctx, key, &operatorContext); err != nil {
			log.Error(err, "failed to get operator context")
			return err
		}

		switch object.(type) {
		case *v1alpha1.ApiDefinition:
			operatorContext.Status.AddLinkedAPIDefinition(objectTarget)
		case *v1alpha1.SecurityPolicy:
			operatorContext.Status.AddLinkedSecurityPolicies(objectTarget)
		case *v1alpha1.PortalAPICatalogue:
			operatorContext.Status.AddLinkedPortalAPICatalogues(objectTarget)
		case *v1alpha1.APIDescription:
			operatorContext.Status.AddLinkedApiDescriptions(objectTarget)
		case *v1alpha1.PortalConfig:
			operatorContext.Status.AddLinkedPortalConfig(objectTarget)
		}

		return client.Status().Update(ctx, &operatorContext)
	}

	return nil
}

// +kubebuilder:rbac:groups=tyk.tyk.io,resources=operatorcontexts,verbs=get;list;watch;create;update;patch;delete

// GetContext returns a OperatorContext resource from k8s cluster with
// namespace/name derived from target. When Spec.FromSecret is provided this
// reads the secret and loads the environment from it. Values set in .Spec.Env
// takes precedence over the values from secret
func GetContext(
	ctx context.Context,
	ns string,
	client client.Client,
	target *model.Target,
	log logr.Logger,
) (*v1alpha1.OperatorContext, error) {
	newTarget := target.NS(ns)

	log.Info("Getting context", "contextRef", newTarget.String())

	var o v1alpha1.OperatorContext

	err := client.Get(ctx, newTarget, &o)
	if err != nil {
		return nil, err
	}

	if o.Spec.Env == nil {
		o.Spec.Env = &v1alpha1.Environment{}
	}

	if o.Spec.FromSecret != nil {
		var secret v1.Secret

		if err := client.Get(ctx, o.Spec.FromSecret.NS(o.Namespace), &secret); err != nil {
			return nil, err
		}

		value := func(key string, fn func(string) error) error {
			if v, ok := secret.Data[key]; ok {
				return fn(string(v))
			}

			return nil
		}
		// we are setting all values that are not set on env but present in secret.
		e := o.Spec.Env
		if e.Mode == "" {
			err := value(v1alpha1.TykMode, func(s string) error {
				e.Mode = v1alpha1.OperatorContextMode(s)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}

		if !e.InsecureSkipVerify {
			err = value(v1alpha1.SkipVerify, func(s string) (err error) {
				e.InsecureSkipVerify, err = strconv.ParseBool(s)
				return
			})
			if err != nil {
				return nil, err
			}
		}

		if e.URL == "" {
			err := value(v1alpha1.TykURL, func(s string) (err error) {
				e.URL = s
				return
			})
			if err != nil {
				return nil, err
			}
		}

		if e.Auth == "" {
			err := value(v1alpha1.TykAuth, func(s string) (err error) {
				e.Auth = s
				return
			})
			if err != nil {
				return nil, err
			}
		}

		if e.UserOwners == nil {
			err := value(v1alpha1.TykUserOwners, func(s string) (err error) {
				for _, user := range strings.Split(s, ",") {
					if o := strings.TrimSpace(user); o != "" {
						e.UserOwners = append(e.UserOwners, o)
					}
				}
				return
			})
			if err != nil {
				return nil, err
			}
		}

		if e.UserGroupOwners == nil {
			err := value(v1alpha1.TykUserGroupOwners, func(s string) (err error) {
				for _, group := range strings.Split(s, ",") {
					if o := strings.TrimSpace(group); o != "" {
						e.UserGroupOwners = append(e.UserGroupOwners, o)
					}
				}
				return
			})
			if err != nil {
				return nil, err
			}
		}

		if e.Org == "" {
			err := value(v1alpha1.TykORG, func(s string) (err error) {
				e.Org = s
				return
			})
			if err != nil {
				return nil, err
			}
		}

		if e.Ingress.HTTPPort == 0 {
			err = value(v1alpha1.IngressHTTPPort, func(s string) (err error) {
				e.Ingress.HTTPPort, err = strconv.Atoi(s)
				return
			})
			if err != nil {
				return nil, err
			}
		}

		if e.Ingress.HTTPSPort == 0 {
			err = value(v1alpha1.IngressTLSPort, func(s string) (err error) {
				e.Ingress.HTTPSPort, err = strconv.Atoi(s)
				return
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return &o, nil
}
