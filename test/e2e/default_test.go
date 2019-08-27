package e2e

import (
	goctx "context"
	"testing"
	"time"

	test "github.com/redhat-cop/quay-operator/test"

	framework "github.com/operator-framework/operator-sdk/pkg/test"

	redhatcopv1alpha1 "github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1"
	"github.com/redhat-cop/quay-operator/pkg/controller/quayecosystem/constants"

	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var name = "quay-operator"
var retryInterval = time.Second * 10
var timeout = time.Second * 80
var cleanupRetryInterval = time.Second * 1
var cleanupTimeout = time.Second * 5

func TestDefaultConfiguration(t *testing.T) {
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()
	test.AddToFrameworkSchemeForTests(t, ctx)
	defaultConfigDeploy(t, framework.Global, ctx)
}

func defaultConfigDeploy(t *testing.T, f *framework.Framework, ctx *framework.TestCtx) {
	namespace, err := ctx.GetNamespace()
	assert.NoError(t, err)

	// Stub out object placeholders for test
	quayEcosystem := &redhatcopv1alpha1.QuayEcosystem{
		TypeMeta: metav1.TypeMeta{
			Kind:       "QuayEcosystem",
			APIVersion: "redhatcop.redhat.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: redhatcopv1alpha1.QuayEcosystemSpec{
			Quay: &redhatcopv1alpha1.Quay{
				ImagePullSecretName: "redhat-pull-secret",
			},
		},
		Status: redhatcopv1alpha1.QuayEcosystemStatus{},
	}

	err = f.Client.Create(goctx.TODO(), quayEcosystem, &framework.CleanupOptions{TestContext: ctx, Timeout: timeout, RetryInterval: retryInterval})
	assert.NoError(t, err)

	// Check if the CRD has been created
	crd := &redhatcopv1alpha1.QuayEcosystem{}
	err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, crd)

	assert.NoError(t, err)

	//Wait for the redis pod deployment
	success := WaitForPodWithImage(t, f, ctx, namespace, "quay-operator-redis", "registry.access.redhat.com/rhscl/redis-32-rhel7:latest", retryInterval, timeout)
	assert.NoError(t, success)

	// Verify the crd has been given default values
	crd = &redhatcopv1alpha1.QuayEcosystem{}
	err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, crd)
	assert.NoError(t, err)
	assert.Equal(t, crd.Spec.Quay.Image, constants.QuayImage)

	// Wait for the postgresql deployment
	success = WaitForPodWithImage(t, f, ctx, namespace, "quay-operator-quay-postgresql", "registry.access.redhat.com/rhscl/postgresql-96-rhel7:1", retryInterval, timeout)
	assert.NoError(t, success)

	// Wait for the quay deployment
	success = WaitForPodWithImage(t, f, ctx, namespace, "quay-operator-quay-config", "quay.io/redhat/quay:v3.0.4", retryInterval, timeout)
	assert.NoError(t, success)
}
