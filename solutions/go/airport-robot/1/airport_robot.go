package airportrobot

type Greeter interface {
	LanguageName() string
	Greet(visitorsName string) string
}

func SayHello(visitorsName string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(visitorsName)
}

type Italian struct {
	name string
}

type Portuguese struct {
	name string
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitorsName string) string {
	return "Ciao " + visitorsName + "!"
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(visitorsName string) string {
	return "Olá " + visitorsName + "!"
}
