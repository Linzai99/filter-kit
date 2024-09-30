package bitmap

import (
	"math"
	"testing"
)

func TestBitMap_Contains(t *testing.T) {
	type args struct {
		key uint64
	}
	tests := []struct {
		name string
		bm   *BitMap
		args args
		want bool
	}{
		{"add 0", NewBitMap(100), args{0}, false},
		{"add 1", NewBitMap(100), args{1}, false},
		{"add 2", NewBitMap(100), args{2}, false},
		{"add 3", NewBitMap(100), args{3}, false},
		{"add 4", NewBitMap(100), args{4}, false},
		{"add max", NewBitMap(100), args{math.MaxUint64}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bm.Contains(tt.args.key); got != tt.want {
				t.Errorf("BitMap.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitMap_Add(t *testing.T) {
	type args struct {
		key uint64
	}
	tests := []struct {
		name    string
		bm      *BitMap
		args    args
		wantErr error
	}{
		{"add 0", NewBitMap(100), args{0}, nil},
		{"add 1", NewBitMap(100), args{1}, nil},
		{"add 2", NewBitMap(100), args{2}, nil},
		{"add 3", NewBitMap(100), args{3}, nil},
		{"add MAX", NewBitMap(100), args{math.MaxUint64}, ErrKeyGtFilterSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bm.Add(tt.args.key); err != nil && err != tt.wantErr {
				t.Errorf("BitMap.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
