#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=front
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}