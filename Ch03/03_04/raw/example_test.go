package elevator

var doorEvent = []byte(`
{
  "type": "door",
  "payload": {
    "action": "open",
    "floor": 3
  }
}
`)

var buttonEvent = []byte(`
{
  "type": "button",
  "payload": {
    "button": 2
  }
}
`)

func ExampleHandleEvent() {
	HandleEvent(doorEvent)
	HandleEvent(buttonEvent)

	// Output:
	// door event: {open 3}
	// button event: {2}
}
