package randomizer

import "testing"

func TestParseVersionFromString(t *testing.T) {
	v, err := ParseVersionFromString("1.2.3-test-4.5")
	if err != nil {
		t.Fatalf("got Err = %v, expected nil", err)
	}
	if v.Major != 1 {
		t.Fatalf("expected Major = 1, got %v", v.Major)
	}
	if v.Minor != 2 {
		t.Fatalf("expected Minor = 2, got %v", v.Minor)
	}
	if v.Patch != 3 {
		t.Fatalf("expected Patch = 3, got %v", v.Patch)
	}
	if v.Branch != "test" {
		t.Fatalf("expected Branch = test, got %v", v.Branch)
	}
	if v.BranchMajor != 4 {
		t.Fatalf("expected BranchMinor = 5, got %v", v.BranchMajor)
	}
	if v.BranchMinor != 5 {
		t.Fatalf("expected BranchMajor = 4, got %v", v.BranchMinor)
	}
	t.Logf("ParseVersionFromString(%s) => %v", "1.2.3-test-4.5", v)
}
