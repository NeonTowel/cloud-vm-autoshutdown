#!/bin/bash

threshold=${1:-0.01}
intervals=${2:-15}
sleep_time=${3:-30}

function require() {
  if ! which $1 >/dev/null; then
    echo "This script requires $1, aborting ..." >&2
    exit 1
  fi
}
require curl
require python3

if ! curl -s -i metadata.google.internal | grep "Metadata-Flavor: Google" >/dev/null; then
  echo "This script only works on GCE VMs, aborting ..." >&2
  exit 1
fi

COMPUTE_METADATA_URL="http://metadata.google.internal/computeMetadata/v1"
VM_PROJECT=$(curl -s "${COMPUTE_METADATA_URL}/project/project-id" -H "Metadata-Flavor: Google" || true)
VM_NAME=$(curl -s "${COMPUTE_METADATA_URL}/instance/hostname" -H "Metadata-Flavor: Google" | cut -d '.' -f 1)
VM_ZONE=$(curl -s "${COMPUTE_METADATA_URL}/instance/zone" -H "Metadata-Flavor: Google" | sed 's/.*zones\///')

ssh_users=$(sudo ss | grep -i ssh | wc -l)

count=0
while true; do
  load=$(uptime | sed -e 's/.*load average: //g' | awk '{ print $3 }')
  if python3 -c "exit(0) if $load >= $threshold else exit(1)"; then
    echo "Resetting count ..." >&2
    count=0
  elif (($ssh_users > 0)); then
    echo "Someone is logged in using ssh ($ssh_users), resetting count ..." >&2
    count=0
  else
    ((count+=1))
    echo "Idle #${count} at $load ..." >&2
  fi
  if ((count>intervals)); then
    if who | grep -v tmux 1>&2; then
      echo "Someone is logged in, won't shut down, resetting count ..." >&2
    elif (($ssh_users > 0)); then
      echo "Someone is logged in using ssh ($ssh_users), won't shut down, resetting count ..." >&2
    else
      echo Shutting down
      # wait a little bit more before actually pulling the plug
      sleep 300
      sudo poweroff
    fi
    count=0
  fi
  sleep $sleep_time
done
