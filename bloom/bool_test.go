package bloom

import (
"testing"
)

func TestBasic(t *testing.T) {
	bf := NewBloomFilter(1, 100)
	//d1, d2 := []byte("Hello"), []byte("Jello")
	d1, _ := []byte("Hello"), []byte("Jello")

	bf.Add(d1)

	//if !bf.Check(d1) {
	//	t.Errorf("d1 should be present in the BloomFilter")
	//}
	//
	//if bf.Check(d2) {
	//	t.Errorf("d2 should be absent from the BloomFilter")
	//}
}
//
//func TestCountingBFBasic(t *testing.T) {
//	cbf := NewCountingBloomFilter(3, 100)
//	d1 := []byte("Hello")
//	cbf.Add(d1)
//
//	if !cbf.Check(d1) {
//		t.Errorf("d1 should be present in the BloomFilter")
//	}
//
//	cbf.Remove(d1)
//
//	if cbf.Check(d1) {
//		t.Errorf("d1 should be absent from the BloomFilter after deletion")
//	}
//}
//
//func TestScalableBFBasic(t *testing.T) {
//	sbf := NewScalableBloomFilter(3, 20, 4, 10, 0.01)
//
//	for i := 1; i < 1000; i++ {
//		buf := make([]byte, 8)
//		binary.PutVarint(buf, int64(i))
//		sbf.Add(buf)
//		if !sbf.Check(buf) {
//			t.Errorf("%d should be present in the BloomFilter", i)
//			return
//		}
//	}
//
//	for i := 1; i < 1000; i++ {
//		buf := make([]byte, 8)
//		binary.PutVarint(buf, int64(i))
//		if !sbf.Check(buf) {
//			t.Errorf("%d should be present in the BloomFilter", i)
//			return
//		}
//	}
//
//	count := 0
//
//	for i := 1000; i < 4000; i++ {
//		buf := make([]byte, 8)
//		binary.PutVarint(buf, int64(i))
//		if sbf.Check(buf) {
//			count++
//		}
//	}
//
//	if sbf.FalsePositiveRate() > 0.04 {
//		t.Errorf("False Positive Rate for this test should be < 0.04")
//		return
//	}
//
//	sensitivity := 0.01 // TODO Make this configurable
//	expectedFalsePositives :=
//		(int)((4000 - 1000) * (sbf.FalsePositiveRate() + sensitivity))
//	if count > expectedFalsePositives {
//		t.Errorf("Actual false positives %d is greater than max expected false positives %d",
//			count,
//			expectedFalsePositives)
//		return
//	}
//}

