#!/bin/zsh

dir=$(dirname $0)

${dir}/sample/bin/peer run 0 --keys ${dir}/sample/keys.yaml | tee peer0.txt 2>@1 &
${dir}/sample/bin/peer run 1 --keys ${dir}/sample/keys.yaml | tee peer1.txt 2>@1 &
${dir}/sample/bin/peer run 2 --keys ${dir}/sample/keys.yaml | tee peer2.txt 2>@1 &
