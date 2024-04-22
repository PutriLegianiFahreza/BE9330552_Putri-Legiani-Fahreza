package Testing

import (
	"testing"
)

func TestTriangle_Area(t *testing.T) {
	segitiga := Triangle{
		Base:   10.0,
		Height: 5.0,
	}
	expected := 25.0
	if segitiga.Area() != expected {
		t.Errorf("salah, seharusnya %f, tapi hasilnya %f", expected, segitiga.Area())
	}
}

func SetupTest(t *testing.T) {
	t.Log("Setup test")
}
