package sysreq

type SystemRequirementInterface interface {
	Has() bool
	GetDescription() string
}
