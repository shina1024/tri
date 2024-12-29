package tri

import (
	"errors"
	"testing"
)

func TestIf(t *testing.T) {
	result := If(true, "true value", "false value")
	if result != "true value" {
		t.Errorf("Expected 'true value', got %s", result)
	}

	result = If(false, "true value", "false value")
	if result != "false value" {
		t.Errorf("Expected 'false value', got %s", result)
	}
}

func TestIfExec(t *testing.T) {
	result := IfExec(true, func() string { return "true func" }, func() string { return "false func" })
	if result != "true func" {
		t.Errorf("Expected 'true func', got %s", result)
	}

	result = IfExec(false, func() string { return "true func" }, func() string { return "false func" })
	if result != "false func" {
		t.Errorf("Expected 'false func', got %s", result)
	}
}

func TestIfExecWithErr(t *testing.T) {
	result, err := IfExecWithErr(true, func() (string, error) { return "true func", nil }, func() (string, error) { return "false func", nil })
	if err != nil || result != "true func" {
		t.Errorf("Expected 'true func', got %s", result)
	}

	result, err = IfExecWithErr(false, func() (string, error) { return "true func", nil }, func() (string, error) { return "false func", nil })
	if err != nil || result != "false func" {
		t.Errorf("Expected 'false func', got %s", result)
	}
}

func TestIfExecWithErrReturnsErr(t *testing.T) {
	expectedErr := errors.New("error")
	_, err := IfExecWithErr(true, func() (string, error) { return "", expectedErr }, func() (string, error) { return "", nil })
	if err != expectedErr {
		t.Errorf("Expected error, got %v", err)
	}

	_, err = IfExecWithErr(false, func() (string, error) { return "", nil }, func() (string, error) { return "", expectedErr })
	if err != expectedErr {
		t.Errorf("Expected error, got %v", err)
	}
}

func TestIfExecOrVal(t *testing.T) {
	result := IfExecOrVal(true, func() string { return "true func" }, "false value")
	if result != "true func" {
		t.Errorf("Expected 'true func', got %s", result)
	}

	result = IfExecOrVal(false, func() string { return "true func" }, "false value")
	if result != "false value" {
		t.Errorf("Expected 'false value', got %s", result)
	}
}

func TestIfExecOrValWithErr(t *testing.T) {
	result, err := IfExecOrValWithErr(true, func() (string, error) { return "true func", nil }, "false value")
	if err != nil || result != "true func" {
		t.Errorf("Expected 'true func', got %s", result)
	}

	result, err = IfExecOrValWithErr(false, func() (string, error) { return "true func", nil }, "false value")
	if err != nil || result != "false value" {
		t.Errorf("Expected 'false value', got %s", result)
	}
}

func TestIfExecOrValWithErrReturnsErr(t *testing.T) {
	expectedErr := errors.New("error")
	_, err := IfExecOrValWithErr(true, func() (string, error) { return "", expectedErr }, "false value")
	if err != expectedErr {
		t.Errorf("Expected error, got %v", err)
	}

	_, err = IfExecOrValWithErr(false, func() (string, error) { return "", nil }, "false value")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestIfValOrExec(t *testing.T) {
	result := IfValOrExec(true, "true value", func() string { return "false func" })
	if result != "true value" {
		t.Errorf("Expected 'true value', got %s", result)
	}

	result = IfValOrExec(false, "true value", func() string { return "false func" })
	if result != "false func" {
		t.Errorf("Expected 'false func', got %s", result)
	}
}

func TestIfValOrExecWithErr(t *testing.T) {
	result, err := IfValOrExecWithErr(true, "true value", func() (string, error) { return "false func", nil })
	if err != nil || result != "true value" {
		t.Errorf("Expected 'true value', got %s", result)
	}

	result, err = IfValOrExecWithErr(false, "true value", func() (string, error) { return "false func", nil })
	if err != nil || result != "false func" {
		t.Errorf("Expected 'false func', got %s", result)
	}
}

func TestIfValOrExecWithErrReturnsErr(t *testing.T) {
	expectedErr := errors.New("error")
	_, err := IfValOrExecWithErr(true, "true value", func() (string, error) { return "", expectedErr })
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = IfValOrExecWithErr(false, "true value", func() (string, error) { return "", expectedErr })
	if err != expectedErr {
		t.Errorf("Expected error, got %v", err)
	}
}
