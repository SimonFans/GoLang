/*
Copyright 2024.

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

package controller

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	envv1alpha1 "myoperator/api/v1alpha1"
)

// EnvironmentReconciler reconciles a Environment object
type EnvironmentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=env.code.operator.com,resources=environments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=env.code.operator.com,resources=environments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=env.code.operator.com,resources=environments/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Environment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *EnvironmentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = log.FromContext(ctx)
	l := log.FromContext(ctx)
	l.Info("Reconcile", "req", req)
	// l.Info("Enter Reconcile ", req)

	// TODO(user): your logic here
	env := &envv1alpha1.Environment{}
	err := r.Get(ctx, req.NamespacedName, env)
	if err != nil {
		if errors.IsNotFound(err) {
			// Environment CR is not found, may have been deleted
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if the corresponding namespace exists
	namespaceName := env.Name
	namespace := &corev1.Namespace{}
	err = r.Get(ctx, types.NamespacedName{Name: namespaceName}, namespace)
	if err != nil && errors.IsNotFound(err) {
		// Namespace does not exist, create it
		namespace = &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespaceName,
			},
		}
		if err := controllerutil.SetControllerReference(env, namespace, r.Scheme); err != nil {
			return reconcile.Result{}, err
		}

		if err := r.Create(ctx, namespace); err != nil {
			return reconcile.Result{}, err
		}

		// Namespace created successfully, update status
		env.Status.NamespaceCreated = true
		if err := r.Status().Update(ctx, env); err != nil {
			return reconcile.Result{}, err
		}

		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		// Error fetching the namespace - requeue the request.
		return reconcile.Result{}, err
	}

	// Namespace already exists, nothing to do
	return reconcile.Result{}, nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *EnvironmentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&envv1alpha1.Environment{}).
		Owns(&corev1.Namespace{}).
		// Watches(
		// 	&source.Kind{Type: &envv1alpha1.Environment{}},
		// 	&handler.EnqueueRequestForOwner{OwnerType: &envv1alpha1.Environment{}},
		// ).
		Complete(r)
}
