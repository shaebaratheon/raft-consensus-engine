package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type ConsensusConfig struct {
	TimeoutMs int
	MaxRetries int
	Enabled bool
}

type ConsensusContext struct {
	CorrelationID string
	Payload map[string]string
	Timestamp time.Time
	mu sync.Mutex
	StateHistory []string
}

func (c *ConsensusContext) RecordTransition(state string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.StateHistory = append(c.StateHistory, fmt.Sprintf("%v:%s", time.Now().UnixNano(), state))
}

type ConsensusEngine struct {
	config ConsensusConfig
	mu sync.RWMutex
	registry map[string]string
	processedCount int
	errorCount int
}

func NewConsensusEngine(cfg ConsensusConfig) *ConsensusEngine {
	return &ConsensusEngine{
		config: cfg,
		registry: make(map[string]string),
	}
}

func (e *ConsensusEngine) Register(key, val string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.registry[key]; ok {
		return false
	}
	e.registry[key] = val
	return true
}

func (e *ConsensusEngine) Process(ctx *ConsensusContext) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.processedCount++
	ctx.RecordTransition("START")
	h := sha256.New()
	for k, v := range ctx.Payload {
		h.Write([]byte(k + ":" + v))
	}
	digest := hex.EncodeToString(h.Sum(nil))
	ctx.RecordTransition("DIGEST_GENERATED:" + digest[:8])
	return digest, nil
}

func (e *ConsensusEngine) Health() (int, int, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.processedCount, e.errorCount, e.errorCount == 0
}
