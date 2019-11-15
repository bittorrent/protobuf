package typedeclimport

import subpkg "github.com/tron-us/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
