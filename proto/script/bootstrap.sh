#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/a.b.c"
exec "$CURDIR/bin/a.b.c"
