package vpp_policer_test

import (
	"testing"

	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
)

func TestPolicerKey(t *testing.T) {
	key := policer.PolicerConfigKey("policer1")
	t.Log(key)
}

func TestPolicerInterfaceKey(t *testing.T) {
	tests := []struct {
		name        string
		policerName string
		ifName      string
		isOutput    bool
		expectedKey string
	}{
		{
			name:        "valid policer name & iface name & input feature",
			policerName: "policer1",
			ifName:      "if1",
			isOutput:    false,
			expectedKey: "vpp/policer/policer1/interface/if1/feature/input",
		},
		{
			name:        "valid policer policer name & iface name & output feature",
			policerName: "policer1",
			ifName:      "if1",
			isOutput:    true,
			expectedKey: "vpp/policer/policer1/interface/if1/feature/output",
		},
		{
			name:        "invalid interface",
			policerName: "policer1",
			ifName:      "",
			expectedKey: "vpp/policer/policer1/interface/<invalid>/feature/input",
		},
		{
			name:        "Gbe interface",
			policerName: "policer1",
			ifName:      "GigabitEthernet0/a/0",
			expectedKey: "vpp/policer/policer1/interface/GigabitEthernet0/a/0/feature/input",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key := policer.DerivedPolicerInterfaceKey(test.policerName, test.ifName, test.isOutput)
			if key != test.expectedKey {
				t.Errorf("failed for: policerName=%s idName=%s\n"+
					"expected key:\n\t%q\ngot key:\n\t%q",
					test.policerName, test.ifName, test.expectedKey, key)
			}
		})
	}
}

func TestParsePolicerInterfaceKey(t *testing.T) {
	tests := []struct {
		name                      string
		key                       string
		expectedPolicerName       string
		expectedIfName            string
		expectedIsOutput          bool
		expectedIsPolicerIfaceKey bool
	}{
		{
			name:                      "valid Policer & iface name & output",
			key:                       "vpp/policer/policer1/interface/if1/feature/output",
			expectedPolicerName:       "policer1",
			expectedIfName:            "if1",
			expectedIsOutput:          true,
			expectedIsPolicerIfaceKey: true,
		},
		{
			name:                      "valid Policer & iface name & input",
			key:                       "vpp/policer/policer1/interface/if1/feature/input",
			expectedPolicerName:       "policer1",
			expectedIfName:            "if1",
			expectedIsOutput:          false,
			expectedIsPolicerIfaceKey: true,
		},
		{
			name:                      "invalid feature",
			key:                       "vpp/policer/policer1/interface/if1/feature/xxx",
			expectedPolicerName:       "",
			expectedIfName:            "",
			expectedIsOutput:          false,
			expectedIsPolicerIfaceKey: false,
		},
		{
			name:                      "invalid interface",
			key:                       "vpp/policer/policer1/interface/<invalid>/feature/output",
			expectedPolicerName:       "policer1",
			expectedIfName:            "<invalid>",
			expectedIsOutput:          true,
			expectedIsPolicerIfaceKey: true,
		},
		{
			name:                      "Gbe interface",
			key:                       "vpp/policer/policer1/interface/GigabitEthernet0/a/0/feature/output",
			expectedPolicerName:       "policer1",
			expectedIfName:            "GigabitEthernet0/a/0",
			expectedIsOutput:          true,
			expectedIsPolicerIfaceKey: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			policerName, ifName, isOutput, isPolicerIfaceKey := policer.ParseDerivedPolicerInterfaceKey(test.key)
			if isPolicerIfaceKey != test.expectedIsPolicerIfaceKey {
				t.Errorf("expected isPolicerIfaceKey: %v\tgot: %v", test.expectedIsPolicerIfaceKey, isPolicerIfaceKey)
			}
			if policerName != test.expectedPolicerName {
				t.Errorf("expected policerName: %s\tgot: %s", test.expectedPolicerName, policerName)
			}
			if ifName != test.expectedIfName {
				t.Errorf("expected ifName: %s\tgot: %s", test.expectedIfName, ifName)
			}
			if isOutput != test.expectedIsOutput {
				t.Errorf("expected isOutput: %v\tgot: %v", test.expectedIsOutput, isOutput)
			}
		})
	}
}
