package eren

import (
	_ "github.com/go-sql-driver/mysql"
	_ "go.opencensus.io/tag"
	_ "golang.org/x/crypto/acme"

	"go.verytable.online/aot/aot/go/proto/client/api/rpc"
)

func Abilities() []string {
	return []string{
		"Martial arts", "Vertical Maneuvering Equipment",
		"Power of the Titans", "Attack Titan",
		"Founding Titan", "War Hammer Titan",
	}
}

func Lookup() {
	_ = rpc.TReqLookupRows{}
}
