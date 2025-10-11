#! /usr/bin/env bash

tmux list-sessions -F '#{session_last_attached}@#{session_name}' | sort -r | awk -F@ '{printf "%s@%s\n", $2, $1}' | samurai
