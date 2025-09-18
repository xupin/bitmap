package bitmap_test

import (
	"testing"

	"github.com/xupin/bitmap/bitmap"
)

func TestBitMap(t *testing.T) {
	bm := bitmap.New(8)

	t.Logf("初始位图: %08b", bm.GetBits())

	// 测试 Add
	bm.Add(8)
	t.Logf("Add(8) 后位图: %08b", bm.GetBits())
	if !bm.Lookup(8) {
		t.Errorf("Add(8) 后 Lookup(8) 应该为 true")
	}

	// 测试 Del
	bm.Del(8)
	t.Logf("Del(8) 后位图: %08b", bm.GetBits())
	if bm.Lookup(8) {
		t.Errorf("Del(8) 后 Lookup(8) 应该为 false")
	}
}
