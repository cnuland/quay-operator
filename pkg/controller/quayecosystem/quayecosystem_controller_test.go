package quayecosystem

import (
	"context"
	"testing"

	routev1 "github.com/openshift/api/route/v1"
	ossecurityv1 "github.com/openshift/api/security/v1"
	"github.com/redhat-cop/operator-utils/pkg/util"
	redhatcopv1alpha1 "github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1"
	"github.com/redhat-cop/quay-operator/pkg/controller/quayecosystem/constants"
	"github.com/redhat-cop/quay-operator/pkg/controller/quayecosystem/setup"
	test "github.com/redhat-cop/quay-operator/test"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var name = "quay-operator"
var namespace = "quay-enterprise"

func TestDefaultConfiguration(t *testing.T) {
	test.SetupLogging()

	quayEcosystem := &redhatcopv1alpha1.QuayEcosystem{
		TypeMeta: metav1.TypeMeta{
			Kind:       "QuayEcosystem",
			APIVersion: "redhatcop.redhat.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	// Objects to track in the fake client.
	objs := []runtime.Object{
		quayEcosystem,
	}
	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(redhatcopv1alpha1.SchemeGroupVersion, quayEcosystem)
	if err := ossecurityv1.AddToScheme(s); err != nil {
		t.Fatal(err, "")
	}
	if err := routev1.AddToScheme(s); err != nil {
		t.Fatal(err, "")
	}
	// Initialize fake client
	cl := fake.NewFakeClient(objs...)

	reconcilerBase := util.NewReconcilerBase(cl, s, nil, nil)
	r := &ReconcileQuayEcosystem{reconcilerBase: reconcilerBase, k8sclient: nil, quaySetupManager: setup.NewQuaySetupManager(reconcilerBase, nil)}

	// Create the Quay Service Account first
	err := cl.Create(context.TODO(), test.ServiceAccount)
	assert.NoError(t, err)
	err = cl.Create(context.TODO(), test.SCCAnyUID)
	assert.NoError(t, err)

	// Initialize the reconicer request
	nsn := types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}

	req := reconcile.Request{
		NamespacedName: nsn,
	}

	r.Reconcile(req)

	// Check if the CR has been created
	crd := &redhatcopv1alpha1.QuayEcosystem{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, crd)
	assert.NoError(t, err)
	// Make sure one of the default values was assigned
	assert.Equal(t, crd.Spec.Quay.Image, constants.QuayImage)

}
