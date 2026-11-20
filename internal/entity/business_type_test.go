package entity

import (
	"reflect"
	"testing"
)

func TestRecommendedLandmarkArrayScanStringValues(t *testing.T) {
	var landmarks RecommendedLandmarkArray
	if err := landmarks.Scan([]byte(`["office","market","residential"]`)); err != nil {
		t.Fatalf("Scan() error = %v", err)
	}

	want := RecommendedLandmarkArray{"office", "market", "residential"}
	if !reflect.DeepEqual(landmarks, want) {
		t.Fatalf("Scan() = %v, want %v", landmarks, want)
	}
}
