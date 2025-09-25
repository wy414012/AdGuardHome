package dhcpd

import (
	"github.com/AdguardTeam/golibs/logutil/slogutil"
)

// testLogger is a logger used in tests.
var testLogger = slogutil.NewDiscardLogger()
