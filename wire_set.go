package coraws

import (
	"github.com/google/wire"
)

// WireSet ...
var WireSet = wire.NewSet(NewAwsSession)
