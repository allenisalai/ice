package sysreq

import (
	"testing"
)

type mockSystemRequirement struct {
	HasValue         bool
	NameValue        string
	DescriptionValue string
}

func (m mockSystemRequirement) Has() bool {
	return m.HasValue
}

func (m mockSystemRequirement) Name() string {
	return m.NameValue
}

func (m mockSystemRequirement) InstallationDescription() string {
	return m.DescriptionValue
}

func TestCheckSystemRequirements(t *testing.T) {
	t.Run("Has All system requirements", func(t *testing.T) {
		sr1 := mockSystemRequirement{
			true,
			"sr1",
			"System Requirement #1",
		}
		sr2 := mockSystemRequirement{
			true,
			"sr2",
			"System Requirement #2",
		}

		AddRequirement(sr1)
		AddRequirement(sr2)

		err := CheckSystemRequirements()
		if err != nil {
			t.Errorf("All system checks should have passed. %v", err)
		}
	})

	t.Run("Is missing a requirement", func(t *testing.T) {
		sr1 := mockSystemRequirement{
			false,
			"sr1",
			"System Requirement #1",
		}
		sr2 := mockSystemRequirement{
			false,
			"sr2",
			"System Requirement #2",
		}
		sr3 := mockSystemRequirement{
			true,
			"sr3",
			"System Requirement #3",
		}

		AddRequirement(sr1)
		AddRequirement(sr2)
		AddRequirement(sr3)

		missingReq := CheckSystemRequirements()
		if len(missingReq) == 0 {
			t.Error("2 of the requirements should have failed.")
		}

		if len(missingReq) != 2 {
			t.Error("2 of the requirements should have failed.")
		}

		if missingReq[0].InstallationDescription() != sr1.DescriptionValue {
			t.Error("The first error should be from the first requirement")
		}

		if missingReq[1].InstallationDescription() != sr2.DescriptionValue {
			t.Error("The second error should be from the second requirement")
		}
	})
}
