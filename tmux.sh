#! /usr/bin/env bash

die() {
  echo "$*" 1>&2
  exit 1
}

toggle_status() {
  local status
  status="$1"

  if [ "$status" = on ]; then
    tmux set -g status off || exit
  else
    tmux set -g status on || exit
  fi
}

if [ -z "$TMUX" ]; then
  die not running inside tmux
fi

current_status=$(tmux show-option -gv status) || exit
trap "tmux set -g status \"$current_status\"" EXIT
toggle_status "$current_status"

session_name=$(tmux list-sessions -F '#{session_last_attached}@#{session_name}' | sort -r | awk -F@ '{printf "%s@%s\n", $2, $1}' | samurai) || exit

echo "$session_name"
