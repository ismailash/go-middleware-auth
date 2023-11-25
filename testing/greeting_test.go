package testing

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

//func TestSayHello_Success(t *testing.T) {
//	// expectation
//	expected := "Hello Maiing 11"
//
//	// execution
//	actual, err := SayHello("Maiing 11")
//
//	// assertion
//	if err != nil {
//		t.Fatal("SayHello_Success() failed, name empty")
//	}
//
//	if expected != actual {
//		t.Fatalf(`SayHello() failed, actual %v, expected %s, %v`, actual, expected, err)
//	}
//}
//
//func TestSayHello_Fail(t *testing.T) {
//	// expected
//	expected := ""
//
//	// actual
//	actual, err := SayHello("")
//
//	if err == nil {
//		t.Fatal("SayHello_Fail() failed, name empty")
//	}
//
//	if expected != actual {
//		t.Fatalf(`SayHello() failed, actual %v, expected %s, %v`, actual, expected, err)
//	}
//}

// Mock struct
type mockGreetingService struct {
	mock.Mock
}

// Mock function
func (m *mockGreetingService) SayHello(name string) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

type GreetingServiceTestSuite struct {
	suite.Suite
	greetingService GreetingService
	greetingMock    *mockGreetingService
}

func (suite *GreetingServiceTestSuite) SetupTest() {
	suite.greetingService = NewGreetingService()
	suite.greetingMock = new(mockGreetingService)
}

// TEST CASES
// Positive
func (suite *GreetingServiceTestSuite) TestSayHello_Success() {
	// ekspektasi
	name := "Maiing"
	expected := "Hello " + name
	suite.greetingMock.On("SayHello", "Maiing").Return(expected, nil)

	// eksekusi / aktual
	actual, err := suite.greetingService.SayHello(name)

	// assertion
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

// suite run
func TestGreetingServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GreetingServiceTestSuite))
}
