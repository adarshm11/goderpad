#!/bin/bash

# Script to kill the goderpad tmux session

SESSION_NAME="goderpad"

# Check if the session exists
tmux has-session -t $SESSION_NAME 2>/dev/null

if [ $? == 0 ]; then
    echo "Killing tmux session: $SESSION_NAME"
    tmux kill-session -t $SESSION_NAME
    echo "Session killed successfully."
else
    echo "Session $SESSION_NAME does not exist."
fi
