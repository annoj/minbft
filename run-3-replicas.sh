#!/bin/zsh

dir=$(dirname $0)

${dir}/sample/bin/peer run 0 --keys ${dir}/sample/keys.yaml &
${dir}/sample/bin/peer run 1 --keys ${dir}/sample/keys.yaml &
${dir}/sample/bin/peer run 2 --keys ${dir}/sample/keys.yaml &
