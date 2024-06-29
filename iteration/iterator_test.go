package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestGetIndex(t *testing.T) {
	t.Run("Getting first index of 'a' in 'cat'", func(t *testing.T) {
		got := GetIndex("Cat", "a")
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Getting first index of 'is' in 'tiramisu'", func(t *testing.T) {
		got := GetIndex("Tiramisu", "is")
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
