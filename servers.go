package steam

import (
	"math/rand"
	"time"

	"github.com/defval/go-steam/v3/netutil"
)

// CMServers contains a list of worlwide servers
var CMServers = []string{
	"155.133.248.52:27018",
	"162.254.197.40:27018",
	"162.254.197.180:27019",
	"155.133.248.50:27019",
	"162.254.197.181:27017",
	"162.254.197.42:27019",
	"162.254.197.180:27017",
	"162.254.197.181:27018",
	"162.254.197.42:27018",
	"155.133.248.50:27017",
	"155.133.248.52:27019",
	"155.133.248.51:27019",
	"155.133.248.53:27019",
	"155.133.248.51:27017",
	"155.133.248.53:27017",
	"155.133.248.52:27017",
	"155.133.248.50:27018",
	"162.254.197.180:27018",
	"162.254.197.40:27017",
	"162.254.197.40:27019",
	"162.254.197.42:27017",
	"162.254.197.181:27019",
	"155.133.248.53:27018",
	"155.133.248.51:27018",
	"146.66.152.11:27017",
	"146.66.152.10:27019",
	"146.66.152.10:27017",
	"146.66.152.10:27018",
	"146.66.152.11:27019",
	"162.254.198.133:27017",
	"162.254.198.133:27018",
	"162.254.198.130:27019",
	"162.254.198.130:27017",
	"162.254.198.132:27018",
	"162.254.198.130:27018",
	"162.254.198.132:27017",
	"162.254.198.132:27019",
	"162.254.198.131:27019",
	"162.254.198.131:27017",
	"146.66.152.11:27018",
	"162.254.198.131:27018",
	"162.254.198.133:27019",
	"185.25.182.77:27017",
	"185.25.182.76:27018",
	"185.25.182.76:27019",
	"185.25.182.77:27018",
	"185.25.182.76:27017",
	"185.25.182.77:27019",
	"162.254.196.67:27019",
	"162.254.196.67:27018",
	"162.254.196.83:27018",
	"162.254.196.84:27018",
	"162.254.196.83:27017",
	"162.254.196.84:27017",
	"162.254.196.68:27019",
	"162.254.196.68:27017",
	"162.254.196.84:27019",
	"162.254.196.67:27017",
	"162.254.196.83:27019",
	"162.254.196.68:27018",
	"146.66.155.101:27017",
	"146.66.155.101:27018",
	"146.66.155.100:27017",
	"146.66.155.100:27018",
	"146.66.155.101:27019",
	"146.66.155.100:27019",
	"155.133.230.50:27017",
	"155.133.230.34:27018",
	"155.133.230.34:27017",
	"155.133.230.50:27019",
	"155.133.230.34:27019",
	"155.133.230.50:27018",
	"162.254.192.100:27017",
	"162.254.192.108:27017",
	"155.133.246.68:27017",
	"155.133.246.68:27018",
	"155.133.246.68:27019",
	"155.133.246.69:27019",
	"155.133.246.69:27017",
	"155.133.246.69:27018",
	"162.254.192.108:27018",
	"162.254.192.101:27018",
	"162.254.192.101:27019",
	"162.254.192.109:27018",
	"162.254.192.100:27018",
	"162.254.192.109:27017",
	"162.254.192.109:27019",
	"162.254.192.108:27019",
	"162.254.192.101:27017",
	"162.254.192.100:27019",
	"162.254.193.46:27019",
	"162.254.193.6:27018",
	"162.254.193.47:27018",
	"162.254.193.6:27019",
	"162.254.193.7:27018",
	"162.254.193.7:27017",
	"162.254.193.7:27019",
	"162.254.193.47:27017",
	"162.254.193.47:27019",
	"162.254.193.46:27018",
}

// GetRandomCM returns a random server from a built-in IP list.
//
// Prefer Client.Connect(), which uses IPs from the Steam Directory,
// which is always more up-to-date.
func GetRandomCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := netutil.ParsePortAddr(CMServers[rng.Int31n(int32(len(CMServers)))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}
