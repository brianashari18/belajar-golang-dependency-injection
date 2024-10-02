package simple

type Foo struct {
}

type Bar struct {
}

type FooBar struct {
	*Foo
	*Bar
}
