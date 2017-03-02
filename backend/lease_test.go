package backend

import (
	"net"
	"testing"
	"time"

	"github.com/digitalrebar/digitalrebar/go/common/store"
)

func TestLeaseCrud(t *testing.T) {
	bs := store.NewSimpleMemoryStore()
	dt := mkDT(bs)
	tests := []crudTest{
		{"Test Invalid Lease Create", dt.create, &Lease{p: dt}, false},
		{"Test Incorrect IP Address Create", dt.create, &Lease{p: dt, Addr: net.ParseIP("127.0.0.1"), Token: "token", ExpireTime: time.Now(), Strategy: "token"}, false},
		{"Test EmptyToken Create", dt.create, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "", ExpireTime: time.Now(), Strategy: "token"}, false},
		{"Test EmptyStrategy Create", dt.create, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "token", ExpireTime: time.Now(), Strategy: ""}, false},
		{"Test Missing Subnet Create", dt.create, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "token", Strategy: "token", ExpireTime: time.Now().Add(10 * time.Second), Valid: true}, false},
		{"Create subnet for creating leases", dt.create, &Subnet{p: dt, Name: "test", Subnet: "192.168.124.0/24", ActiveStart: net.ParseIP("192.168.124.80"), ActiveEnd: net.ParseIP("192.168.124.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "noop"}, true},
		{"Test Valid Create", dt.create, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "token", Strategy: "token", ExpireTime: time.Now().Add(10 * time.Second), Valid: true}, true},
		{"Test Duplicate IP Create", dt.create, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "token", Strategy: "token", ExpireTime: time.Now().Add(10 * time.Second), Valid: true}, false},
		{"Test Duplicate Token Create", dt.create, &Lease{p: dt, Addr: net.ParseIP("192.168.124.11"), Token: "token", Strategy: "token", ExpireTime: time.Now().Add(10 * time.Second), Valid: true}, false},
		{"Test Token Update", dt.update, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "token2", Strategy: "token", ExpireTime: time.Now().Add(10 * time.Second), Valid: true}, false},
		{"Test Strategy Update", dt.update, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "token", Strategy: "token2", ExpireTime: time.Now().Add(10 * time.Second), Valid: true}, false},
		{"Test Expire Update", dt.update, &Lease{p: dt, Addr: net.ParseIP("192.168.124.10"), Token: "token", Strategy: "token", ExpireTime: time.Now().Add(10 * time.Minute), Valid: true}, true},
	}
	for _, test := range tests {
		test.Test(t)
	}
}