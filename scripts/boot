#!/bin/bash
base_path=$(realpath --relative-to="$(pwd)" ~/GolandProjects/)

function process() {
  local cmd="cd $base_path/${1} && make server; exec bash"
  gnome-terminal --tab --title="${1}" -- bash -c "${cmd}"
}

process "auth-svc"
process "docker-svc"
process "api-gateway"


