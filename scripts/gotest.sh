#!/usr/bin/env bash

echo "==> Running go test and creating a coverage profile..."
i=0
for testingpkg in $(go list ./selvpcclient/.../testing); do
  coverpkg=${testingpkg::-8}
  go test -v -covermode count -coverprofile "./${i}.coverprofile" -coverpkg $coverpkg $testingpkg
  ((i++))
done
gocovmerge $(ls ./*.coverprofile) > coverage.out
rm *.coverprofile