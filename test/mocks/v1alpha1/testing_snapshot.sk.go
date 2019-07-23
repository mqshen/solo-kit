// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"fmt"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type TestingSnapshot struct {
	Mocks MockResourceList
}

func (s TestingSnapshot) Clone() TestingSnapshot {
	return TestingSnapshot{
		Mocks: s.Mocks.Clone(),
	}
}

func (s TestingSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMocks(),
	)
}

func (s TestingSnapshot) hashMocks() uint64 {
	return hashutils.HashAll(s.Mocks.AsInterfaces()...)
}

func (s TestingSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("mocks", s.hashMocks()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type TestingSnapshotStringer struct {
	Version uint64
	Mocks   []string
}

func (ss TestingSnapshotStringer) String() string {
	s := fmt.Sprintf("TestingSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Mocks %v\n", len(ss.Mocks))
	for _, name := range ss.Mocks {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s TestingSnapshot) Stringer() TestingSnapshotStringer {
	return TestingSnapshotStringer{
		Version: s.Hash(),
		Mocks:   s.Mocks.NamespacesDotNames(),
	}
}