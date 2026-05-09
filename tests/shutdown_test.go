package tests

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/logger/writers"
	"github.com/aliworkshop/sample_project/chat/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeWS is a minimal gateway.WebSocketHandler used to drive client.Stop()
// without a real socket. Read blocks until Close is called and then returns a
// plain error (NOT a websocket.CloseError) so the read loop exits without
// recursively calling client.Stop().
type fakeWS struct {
	closed   chan struct{}
	once     sync.Once
	closeHit atomic.Int32
}

func newFakeWS() *fakeWS { return &fakeWS{closed: make(chan struct{})} }

func (f *fakeWS) Read(ctx context.Context) (int, []byte, error) {
	select {
	case <-f.closed:
		return 0, nil, errors.New("connection closed")
	case <-ctx.Done():
		return 0, nil, ctx.Err()
	}
}

func (f *fakeWS) Write(_ context.Context, _ int, _ []byte) error   { return nil }
func (f *fakeWS) WriteJson(_ context.Context, _ interface{}) error { return nil }
func (f *fakeWS) SetWriteDeadLine(_ time.Duration) error           { return nil }
func (f *fakeWS) SetReadDeadLine(_ time.Duration) error            { return nil }
func (f *fakeWS) Ping(_ context.Context) error                     { return nil }
func (f *fakeWS) Close() {
	f.closeHit.Add(1)
	f.once.Do(func() { close(f.closed) })
}

func newTestUC(t *testing.T) interface {
	Start()
	Stop()
} {
	t.Helper()
	log := logger.NewSimpleLogger(writers.WarnLevel, logger.JsonEncoding)
	return usecase.NewUseCase(log)
}

// withDeadline runs fn and fails the test if it doesn't return within d.
func withDeadline(t *testing.T, d time.Duration, name string, fn func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		fn()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(d):
		t.Fatalf("%s did not return within %s", name, d)
	}
}

func TestShutdown_EmptyUseCase(t *testing.T) {
	uc := newTestUC(t)
	uc.Start()
	withDeadline(t, time.Second, "Stop", uc.Stop)
}

func TestShutdown_DrainsClients(t *testing.T) {
	log := logger.NewSimpleLogger(writers.WarnLevel, logger.JsonEncoding)
	uc := usecase.NewUseCase(log)
	uc.Start()

	const n = 5
	wsList := make([]*fakeWS, n)
	for i := 0; i < n; i++ {
		ws := newFakeWS()
		wsList[i] = ws
		c, err := uc.Subscribe(uint64(i+1), ws)
		require.Nil(t, err)
		require.NotNil(t, c)
	}

	// Give the clients a beat to start their read/write goroutines so
	// Stop() exercises the full drain path.
	time.Sleep(20 * time.Millisecond)

	withDeadline(t, 2*time.Second, "Stop", uc.Stop)

	for i, ws := range wsList {
		assert.Greater(t, ws.closeHit.Load(), int32(0), "ws[%d] was not closed", i)
	}
}

func TestShutdown_Idempotent(t *testing.T) {
	uc := newTestUC(t)
	uc.Start()

	withDeadline(t, time.Second, "first Stop", uc.Stop)
	// Second call must be a no-op and must not panic.
	withDeadline(t, time.Second, "second Stop", uc.Stop)
}

func TestShutdown_ConcurrentStopCalls(t *testing.T) {
	uc := newTestUC(t)
	uc.Start()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			uc.Stop()
		}()
	}
	withDeadline(t, 2*time.Second, "concurrent Stop", wg.Wait)
}
