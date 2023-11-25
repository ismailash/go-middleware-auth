package testing

import "errors"

//func SayHello(name string) (string, error) {
//	if name == "" {
//		return "", errors.New("name can't be empty")
//	}
//
//	return "Hello " + name, nil
//}

type GreetingService interface {
	SayHello(name string) (string, error)
}

type greetingService struct {
}

func (g *greetingService) SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can't be empty")
	}
	return "Hello " + name, nil
}

func NewGreetingService() GreetingService {
	return &greetingService{}
}

// BUAT TESTNYA
