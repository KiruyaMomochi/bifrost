package tables

import (
	"testing"
	"time"
)

// TestIsExpiredAt validates the expiry boundary conditions for TableVirtualKey.
// All tests use fixed timestamps; no raw time.Now() in test logic.
func TestIsExpiredAt(t *testing.T) {
	now := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)

	t.Run("nil receiver returns false", func(t *testing.T) {
		var vk *TableVirtualKey
		if vk.IsExpiredAt(now) {
			t.Fatal("nil receiver should return false")
		}
	})

	t.Run("nil ExpiresAt returns false", func(t *testing.T) {
		vk := &TableVirtualKey{}
		if vk.IsExpiredAt(now) {
			t.Fatal("nil ExpiresAt should return false (never expires)")
		}
	})

	t.Run("future ExpiresAt returns false", func(t *testing.T) {
		future := now.Add(time.Hour)
		vk := &TableVirtualKey{ExpiresAt: &future}
		if vk.IsExpiredAt(now) {
			t.Fatal("future ExpiresAt should return false")
		}
	})

	t.Run("past ExpiresAt returns true", func(t *testing.T) {
		past := now.Add(-time.Second)
		vk := &TableVirtualKey{ExpiresAt: &past}
		if !vk.IsExpiredAt(now) {
			t.Fatal("past ExpiresAt should return true")
		}
	})

	t.Run("now == ExpiresAt is expired (boundary)", func(t *testing.T) {
		expiry := now // exactly equal
		vk := &TableVirtualKey{ExpiresAt: &expiry}
		if !vk.IsExpiredAt(now) {
			t.Fatal("ExpiresAt exactly equal to now should be expired")
		}
	})

	t.Run("ExpiresAt one nanosecond in the future is not expired", func(t *testing.T) {
		almostNow := now.Add(time.Nanosecond)
		vk := &TableVirtualKey{ExpiresAt: &almostNow}
		if vk.IsExpiredAt(now) {
			t.Fatal("ExpiresAt one nanosecond in the future should not be expired")
		}
	})
}
