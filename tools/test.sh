#!/usr/bin/env bash

set -e
echo "mode: atomic" > coverage.txt

packages="github.com/digitalrebar/provision,\
github.com/digitalrebar/provision/backend,\
github.com/digitalrebar/provision/midlayer,\
github.com/digitalrebar/provision/frontend,\
github.com/digitalrebar/provision/embedded,\
github.com/digitalrebar/provision/server,\
github.com/digitalrebar/provision/cli\
"

for d in $(go list ./... 2>/dev/null | grep -v cmds | grep -v vendor | grep -v github.com/digitalrebar/provision/client  | grep -v github.com/digitalrebar/provision/models) ; do
    dir=${d//github.com\/digitalrebar\/provision}
    rm -f test.bin
    go test -o test.bin -c -race -covermode=atomic -coverpkg=$packages "$d"
    if [ -e test.bin ] ; then
        cd .$dir
        time ../test.bin -test.coverprofile=../profile.out
        cd ..
        rm -f test.bin
    fi

    if [ -f profile.out ]; then
        grep -h -v "^mode:" profile.out >> coverage.txt
        rm profile.out
    fi
done

