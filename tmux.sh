#! /usr/bin/env bash

session_name=$(tmux list-sessions -F '#{session_last_attached}@#{session_name}' | sort -r | awk -F@ '{printf "%s@%s\n", $2, $1}' | samurai) || exit

[ -z "$session_name" ] && exit 0

tmux switch-client -t "$session_name"
