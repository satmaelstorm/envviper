package envviper

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type replacementTestSuite struct {
	suite.Suite
}

func TestReplacement(t *testing.T) {
	suite.Run(t, new(replacementTestSuite))
}

func (s *replacementTestSuite) Test() {
	r := NewReplacement(".", "_")
	s.Equal(".", r.InVar)
	s.Equal("_", r.InEnv)
}
