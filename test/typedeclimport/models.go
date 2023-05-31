package typedeclimport

import subpkg "github.com/bittorrent/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
