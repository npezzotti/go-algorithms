package time_based_key_value_store

import (
	"testing"
)

func TestTimeMap(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1) // store the key "foo" and value "bar" along with timestamp = 1.
	bar := timeMap.Get("foo", 1) // return "bar"
	if bar != "bar" {
		t.Fatalf("expected \"bar\", got %q", bar)
	}

	bar = timeMap.Get("foo", 3) // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".

	if bar != "bar" {
		t.Fatalf("expected \"bar\", got %q", bar)
	}

	timeMap.Set("foo", "bar2", 4) // store the key "foo" and value "bar2" along with timestamp = 4.
	bar2 := timeMap.Get("foo", 4) // return "bar2"
	if bar2 != "bar2" {
		t.Fatalf("expected \"bar2\", got %q", bar2)
	}

	bar2 = timeMap.Get("foo", 5) // return "bar2"
	if bar2 != "bar2" {
		t.Fatalf("expected \"bar2\", got %q", bar2)
	}
}
