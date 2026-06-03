package basic

import (
	"testing"

	"github.com/maurodesouza/learning-go/src/utils"
)

func TestClass1(t *testing.T) {
	utils.AssertEqual(t, globalStringVar, "")
	utils.AssertEqual(t, globalIntVar, 0)
	utils.AssertEqual(t, globalBoolVar, false)
	utils.AssertEqual(t, globalFloatVar, 0.0)
}
