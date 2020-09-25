package v1alpha1

import (
	"testing"

	"github.com/hashicorp/consul-k8s/api/common"
	capi "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestToConsul(t *testing.T) {
	cases := map[string]struct {
		input    *ServiceDefaults
		expected *capi.ServiceConfigEntry
	}{
		"kind:service-defaults": {
			&ServiceDefaults{},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
			},
		},
		"name:resource-name": {
			&ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "resource-name",
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Name: "resource-name",
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
			},
		},
		"protocol:http": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					Protocol: "http",
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				Protocol: "http",
			},
		},
		"protocol:https": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					Protocol: "https",
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				Protocol: "https",
			},
		},
		"protocol:''": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					Protocol: "",
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				Protocol: "",
			},
		},
		"mode:unsupported": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "unsupported",
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				MeshGateway: capi.MeshGatewayConfig{
					Mode: capi.MeshGatewayModeDefault,
				},
			},
		},
		"mode:local": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "local",
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				MeshGateway: capi.MeshGatewayConfig{
					Mode: capi.MeshGatewayModeLocal,
				},
			},
		},
		"mode:remote": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "remote",
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				MeshGateway: capi.MeshGatewayConfig{
					Mode: capi.MeshGatewayModeRemote,
				},
			},
		},
		"mode:none": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "none",
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				MeshGateway: capi.MeshGatewayConfig{
					Mode: capi.MeshGatewayModeNone,
				},
			},
		},
		"mode:default": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "default",
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				MeshGateway: capi.MeshGatewayConfig{
					Mode: capi.MeshGatewayModeDefault,
				},
			},
		},
		"mode:''": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "",
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				MeshGateway: capi.MeshGatewayConfig{
					Mode: capi.MeshGatewayModeDefault,
				},
			},
		},
		"externalSNI:test-external-sni": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					ExternalSNI: "test-external-sni",
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				ExternalSNI: "test-external-sni",
			},
		},
		"externalSNI:''": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					ExternalSNI: "",
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				ExternalSNI: "",
			},
		},
		"expose.checks:false": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					Expose: ExposeConfig{
						Checks: false,
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				Expose: capi.ExposeConfig{
					Checks: false,
				},
			},
		},
		"expose.checks:true": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					Expose: ExposeConfig{
						Checks: true,
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				Expose: capi.ExposeConfig{
					Checks: true,
				},
			},
		},
		"expose.paths:single": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					Expose: ExposeConfig{
						Paths: []ExposePath{
							{
								ListenerPort:  80,
								Path:          "/test/path",
								LocalPathPort: 42,
								Protocol:      "tcp",
							},
						},
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				Expose: capi.ExposeConfig{
					Paths: []capi.ExposePath{
						{
							ListenerPort:  80,
							Path:          "/test/path",
							LocalPathPort: 42,
							Protocol:      "tcp",
						},
					},
				},
			},
		},
		"expose.paths:multiple": {
			&ServiceDefaults{
				Spec: ServiceDefaultsSpec{
					Expose: ExposeConfig{
						Paths: []ExposePath{
							{
								ListenerPort:  80,
								Path:          "/test/path",
								LocalPathPort: 42,
								Protocol:      "tcp",
							},
							{
								ListenerPort:  8080,
								Path:          "/root/test/path",
								LocalPathPort: 4201,
								Protocol:      "https",
							},
						},
					},
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "test-dc",
				},
				Expose: capi.ExposeConfig{
					Paths: []capi.ExposePath{
						{
							ListenerPort:  80,
							Path:          "/test/path",
							LocalPathPort: 42,
							Protocol:      "tcp",
						},
						{
							ListenerPort:  8080,
							Path:          "/root/test/path",
							LocalPathPort: 4201,
							Protocol:      "https",
						},
					},
				},
			},
		},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			output := testCase.input.ToConsul("test-dc")
			require.Equal(t, testCase.expected, output)
		})
	}
}

func TestMatchesConsul(t *testing.T) {
	cases := map[string]struct {
		internal *ServiceDefaults
		consul   *capi.ServiceConfigEntry
	}{
		"no fields set": {
			&ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "defaults",
				},
				Spec: ServiceDefaultsSpec{},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Name: "defaults",
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "",
				},
			},
		},
		"all fields set": {
			&ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "defaults",
				},
				Spec: ServiceDefaultsSpec{
					Protocol: "tcp",
					MeshGateway: MeshGatewayConfig{
						Mode: "remote",
					},
					Expose: ExposeConfig{
						Checks: true,
						Paths: []ExposePath{
							{
								ListenerPort:  80,
								Path:          "/default",
								LocalPathPort: 9091,
								Protocol:      "http2",
							},
							{
								ListenerPort:  8080,
								Path:          "/remote",
								LocalPathPort: 9000,
								Protocol:      "tcp",
							},
						},
					},
					ExternalSNI: "valid-external-sni",
				},
			},
			&capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Name: "defaults",
				Meta: map[string]string{
					common.SourceKey:     common.SourceValue,
					common.DatacenterKey: "",
				},
				Protocol: "tcp",
				MeshGateway: capi.MeshGatewayConfig{
					Mode: capi.MeshGatewayModeRemote,
				},
				Expose: capi.ExposeConfig{
					Checks: true,
					Paths: []capi.ExposePath{
						{
							ListenerPort:  80,
							Path:          "/default",
							LocalPathPort: 9091,
							Protocol:      "http2",
						},
						{
							ListenerPort:  8080,
							Path:          "/remote",
							LocalPathPort: 9000,
							Protocol:      "tcp",
						},
					},
				},
				ExternalSNI: "valid-external-sni",
			},
		},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			result := testCase.internal.MatchesConsul(testCase.consul, "")
			require.True(t, result)
		})
	}
}

func TestServiceDefaults_MatchesDatacenter(t *testing.T) {
	cases := map[string]struct {
		ConfigEntry    *capi.ServiceConfigEntry
		DatacenterName string
		Matches        bool
	}{
		"Datacenter empty": {
			ConfigEntry: &capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Name: "svc-default",
				Meta: map[string]string{
					common.DatacenterKey: "this-datacenter",
				},
			},
			DatacenterName: "",
			Matches:        false,
		},
		"Metadata empty": {
			ConfigEntry: &capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Name: "svc-default",
				Meta: map[string]string{
					common.DatacenterKey: "",
				},
			},
			DatacenterName: "this-datacenter",
			Matches:        false,
		},
		"Different values": {
			ConfigEntry: &capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Name: "svc-default",
				Meta: map[string]string{
					common.DatacenterKey: "other-datacenter",
				},
			},
			DatacenterName: "this-datacenter",
			Matches:        false,
		},
		"Matches": {
			ConfigEntry: &capi.ServiceConfigEntry{
				Kind: capi.ServiceDefaults,
				Name: "svc-default",
				Meta: map[string]string{
					common.DatacenterKey: "this-datacenter",
				},
			},
			DatacenterName: "this-datacenter",
			Matches:        true,
		},
	}

	for name, test := range cases {
		serviceDefault := &ServiceDefaults{}
		t.Run(name, func(t *testing.T) {
			require.Equal(t, serviceDefault.MatchesDatacenter(test.ConfigEntry, test.DatacenterName), test.Matches)
		})
	}
}

func TestServiceDefaults_Validate(t *testing.T) {
	cases := map[string]struct {
		input          *ServiceDefaults
		expectedErrMsg string
	}{
		"valid": {
			input: &ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-service",
				},
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "remote",
					},
				},
			},
			expectedErrMsg: "",
		},
		"meshgateway.mode": {
			&ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-service",
				},
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "foobar",
					},
				},
			},
			`servicedefaults.consul.hashicorp.com "my-service" is invalid: spec.meshGateway.mode: Invalid value: "foobar": must be one of "remote", "local", "none", ""`,
		},
		"expose.paths[].protocol": {
			&ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-service",
				},
				Spec: ServiceDefaultsSpec{
					Expose: ExposeConfig{
						Paths: []ExposePath{
							{
								Protocol: "invalid-protocol",
								Path:     "/valid-path",
							},
						},
					},
				},
			},
			`servicedefaults.consul.hashicorp.com "my-service" is invalid: spec.expose.paths[0].protocol: Invalid value: "invalid-protocol": must be one of "http", "http2"`,
		},
		"expose.paths[].path": {
			&ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-service",
				},
				Spec: ServiceDefaultsSpec{
					Expose: ExposeConfig{
						Paths: []ExposePath{
							{
								Protocol: "http",
								Path:     "invalid-path",
							},
						},
					},
				},
			},
			`servicedefaults.consul.hashicorp.com "my-service" is invalid: spec.expose.paths[0].path: Invalid value: "invalid-path": must begin with a '/'`,
		},
		"multi-error": {
			&ServiceDefaults{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-service",
				},
				Spec: ServiceDefaultsSpec{
					MeshGateway: MeshGatewayConfig{
						Mode: "invalid-mode",
					},
					Expose: ExposeConfig{
						Paths: []ExposePath{
							{
								Protocol: "invalid-protocol",
								Path:     "invalid-path",
							},
						},
					},
				},
			},
			`servicedefaults.consul.hashicorp.com "my-service" is invalid: [spec.meshGateway.mode: Invalid value: "invalid-mode": must be one of "remote", "local", "none", "", spec.expose.paths[0].path: Invalid value: "invalid-path": must begin with a '/', spec.expose.paths[0].protocol: Invalid value: "invalid-protocol": must be one of "http", "http2"]`,
		},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			err := testCase.input.Validate()
			if testCase.expectedErrMsg != "" {
				require.EqualError(t, err, testCase.expectedErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestServiceDefaults_AddFinalizer(t *testing.T) {
	serviceDefaults := &ServiceDefaults{}
	serviceDefaults.AddFinalizer("finalizer")
	require.Equal(t, []string{"finalizer"}, serviceDefaults.ObjectMeta.Finalizers)
}

func TestServiceDefaults_RemoveFinalizer(t *testing.T) {
	serviceDefaults := &ServiceDefaults{
		ObjectMeta: metav1.ObjectMeta{
			Finalizers: []string{"f1", "f2"},
		},
	}
	serviceDefaults.RemoveFinalizer("f1")
	require.Equal(t, []string{"f2"}, serviceDefaults.ObjectMeta.Finalizers)
}

func TestServiceDefaults_SetSyncedCondition(t *testing.T) {
	serviceDefaults := &ServiceDefaults{}
	serviceDefaults.SetSyncedCondition(corev1.ConditionTrue, "reason", "message")

	require.Equal(t, corev1.ConditionTrue, serviceDefaults.Status.Conditions[0].Status)
	require.Equal(t, "reason", serviceDefaults.Status.Conditions[0].Reason)
	require.Equal(t, "message", serviceDefaults.Status.Conditions[0].Message)
	now := metav1.Now()
	require.True(t, serviceDefaults.Status.Conditions[0].LastTransitionTime.Before(&now))
}

func TestServiceDefaults_GetSyncedConditionStatus(t *testing.T) {
	cases := []corev1.ConditionStatus{
		corev1.ConditionUnknown,
		corev1.ConditionFalse,
		corev1.ConditionTrue,
	}
	for _, status := range cases {
		t.Run(string(status), func(t *testing.T) {
			serviceDefaults := &ServiceDefaults{
				Status: Status{
					Conditions: []Condition{{
						Type:   ConditionSynced,
						Status: status,
					}},
				},
			}

			require.Equal(t, status, serviceDefaults.SyncedConditionStatus())
		})
	}
}

// Test that if status is empty then GetCondition returns nil.
func TestServiceDefaults_GetConditionWhenNil(t *testing.T) {
	serviceDefaults := &ServiceDefaults{}
	require.Nil(t, serviceDefaults.GetCondition(ConditionSynced))
}
