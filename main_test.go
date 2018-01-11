package main

import "testing"

func TestSamplerN1(t *testing.T) {
	sampler := NewSampler(1)
	for i := 0; i < 10; i++ {
		if !sampler.Sample() {
			t.Fatal("expected Sample to return true")
		}
	}
}

func TestSamplerN2(t *testing.T) {
	sampler := NewSampler(2)
	for i := 0; i < 10; i++ {
		if sampler.Sample() {
			t.Fatal("expected Sample to return false")
		}

		if !sampler.Sample() {
			t.Fatal("expected Sample to return true")
		}
	}
}

func TestSamplerN3(t *testing.T) {
	sampler := NewSampler(3)
	for i := 0; i < 10; i++ {
		if sampler.Sample() {
			t.Fatal("expected Sample to return false")
		}

		if sampler.Sample() {
			t.Fatal("expected Sample to return false")
		}

		if !sampler.Sample() {
			t.Fatal("expected Sample to return true")
		}
	}
}
