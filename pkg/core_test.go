package pkg

import (
	"testing"
	"time"
)

func TestConsensusLifecycle(t *testing.T) {
	cfg := ConsensusConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewConsensusEngine(cfg)
	ctx := &ConsensusContext{
		CorrelationID: "cid-consensus-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestLogStoreLifecycle(t *testing.T) {
	cfg := LogStoreConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewLogStoreEngine(cfg)
	ctx := &LogStoreContext{
		CorrelationID: "cid-log_store-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestRpcTransportLifecycle(t *testing.T) {
	cfg := RpcTransportConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewRpcTransportEngine(cfg)
	ctx := &RpcTransportContext{
		CorrelationID: "cid-rpc_transport-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestStateMachineLifecycle(t *testing.T) {
	cfg := StateMachineConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewStateMachineEngine(cfg)
	ctx := &StateMachineContext{
		CorrelationID: "cid-state_machine-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestSnapshotterLifecycle(t *testing.T) {
	cfg := SnapshotterConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewSnapshotterEngine(cfg)
	ctx := &SnapshotterContext{
		CorrelationID: "cid-snapshotter-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestMembershipLifecycle(t *testing.T) {
	cfg := MembershipConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewMembershipEngine(cfg)
	ctx := &MembershipContext{
		CorrelationID: "cid-membership-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

