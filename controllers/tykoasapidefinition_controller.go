/*


Licensed under the Mozilla Public License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.mozilla.org/en-US/MPL/2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	util "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	tykv1alpha1 "github.com/TykTechnologies/tyk-operator/api/v1alpha1"
	"github.com/TykTechnologies/tyk-operator/pkg/client/klient"
	"github.com/TykTechnologies/tyk-operator/pkg/environmet"
	v1 "k8s.io/api/core/v1"
)

// TykOASApiDefinitionReconciler reconciles a TykOASApiDefinition object
type TykOASApiDefinitionReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Env    environmet.Env
}

//+kubebuilder:rbac:groups=tyk.tyk.io,resources=tykoasapidefinitions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tyk.tyk.io,resources=tykoasapidefinitions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tyk.tyk.io,resources=tykoasapidefinitions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TykOASApiDefinition object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *TykOASApiDefinitionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("tykoasapidefinition", req.NamespacedName)

	log.Info("Reconciling TykOASApiDefinition")

	tykOASDef := &tykv1alpha1.TykOASApiDefinition{}

	err := r.Client.Get(ctx, req.NamespacedName, tykOASDef)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	_, ctx, err = HttpContext(ctx, r.Client, r.Env, tykOASDef, log)
	if err != nil {
		return ctrl.Result{}, err
	}

	_, err = util.CreateOrUpdate(ctx, r.Client, tykOASDef, func() error {
		// create OAS API if apiID is empty
		if tykOASDef.Status.ApiID == "" {
			id, err := r.createTykOAS(ctx, *tykOASDef)
			if err != nil {
				return err
			}

			// when to update the status
			tykOASDef.Status.ApiID = id
		}

		return nil
	})
	if err != nil {
		return ctrl.Result{RequeueAfter: queueAfter}, err
	}

	err = r.Client.Status().Update(ctx, tykOASDef)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *TykOASApiDefinitionReconciler) createTykOAS(ctx context.Context, tykOASDef tykv1alpha1.TykOASApiDefinition) (string, error) {
	r.Log.Info("Creating OAS API on Tyk")

	var cm v1.ConfigMap

	cm_name := types.NamespacedName{Name: tykOASDef.Spec.OASRef.Name, Namespace: tykOASDef.Spec.OASRef.Namespace}

	err := r.Client.Get(ctx, cm_name, &cm)
	if err != nil {
		return "", err
	}

	data := cm.Data[tykOASDef.Spec.OASRef.KeyName]
	if data == "" {
		err = errors.New("OAS Spec is empty")

		r.Log.Error(err, "Failed to create OAS API Definition")

		return "", err
	}

	result, err := klient.Universal.OAS().Create(ctx, data)
	if err != nil {
		r.Log.Error(err, "Failed to create OAS API Definition")
		return "", err
	}

	r.Log.Info("Successfully created OAS API on Tyk", "id", result.Meta)

	return result.Meta, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TykOASApiDefinitionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tykv1alpha1.TykOASApiDefinition{}).
		Complete(r)
}
