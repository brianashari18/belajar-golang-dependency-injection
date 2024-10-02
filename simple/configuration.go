package simple

type Configuration struct {
	name string
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	return &Application{Configuration: &Configuration{
		name: "Application",
	}}
}
