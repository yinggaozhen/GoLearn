package suite

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MySuite struct {
	suite.Suite
}

func (s *MySuite) SetupTest() {
	fmt.Println("SetupTest")
}

func (s *MySuite) TearDownTest() {
	fmt.Println("TearDownTest")
}

func (s *MySuite) SetupSuite() {
	fmt.Println("SetupSuite")
}

func (s *MySuite) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *MySuite) TestExample1() {
	fmt.Println("TestExample1")
}

func (s *MySuite) TestExample2() {
	fmt.Println("TestExample2")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}