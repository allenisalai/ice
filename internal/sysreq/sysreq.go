package sysreq

type SystemRequirementInterface interface {
	Has() bool
	Name() string
	InstallationDescription() string
}

var sysReqs []SystemRequirementInterface

func AddRequirement(sr SystemRequirementInterface) {
	sysReqs = append(sysReqs, sr)
}

func CheckSystemRequirements() []SystemRequirementInterface {
	var missingReqs []SystemRequirementInterface
	for _, sr := range sysReqs {
		if sr.Has() == false {
			missingReqs = append(missingReqs, sr)
		}
	}

	return missingReqs
}

func GetSystemRequirements() []SystemRequirementInterface {
	return sysReqs
}
