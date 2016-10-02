#!/bin/bash

COVERAGE=`go test -cover -v ./... | grep -o '^coverage: [0-9]*.[0-9]*\%' | awk '{print $2}'`
NOPERCENT=`echo $COVERAGE | tr '%' ' '`
COUNT=`echo $NOPERCENT | wc -w`
TOTAL=0
for i in $NOPERCENT; do
    TOTAL=`echo "scale = 2; $TOTAL + $i" | bc`
done
printf "coverage: %.1f%%\n" `echo "scale = 2; $TOTAL / $COUNT" | bc`
