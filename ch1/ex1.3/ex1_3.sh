#!/bin/bash

TARGET1="./echo2/main"
TARGET2="./echo3/main"
function benchmark () {
    local BIN=$1

    for I in `seq 1 1000`;
    do
        $BIN $I > /dev/null
    done
}

time benchmark $TARGET1
time benchmark $TARGET2
