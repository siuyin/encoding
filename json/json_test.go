package json

func ExampleLog() {
	Log(struct {
		Op string
	}{"test op"})
}
