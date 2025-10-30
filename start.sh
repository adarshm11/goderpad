#!/bin/bash

# Script to start frontend and backend in tmux

SESSION_NAME="goderpad"

# Check if the session already exists
tmux has-session -t $SESSION_NAME 2>/dev/null

if [ $? != 0 ]; then
    echo "Creating new tmux session: $SESSION_NAME"

    # Create new session with backend in first pane
    tmux new-session -d -s $SESSION_NAME -n "dev"

    # Start backend in the first pane
    tmux send-keys -t $SESSION_NAME:0.0 "cd backend && go run main.go" C-m

    # Split window vertically and start frontend
    tmux split-window -h -t $SESSION_NAME:0
    tmux send-keys -t $SESSION_NAME:0.1 "cd frontend && npm run dev" C-m

    # Select the first pane
    tmux select-pane -t $SESSION_NAME:0.0

    echo "Session created. Attaching..."
else
    echo "Session $SESSION_NAME already exists. Attaching..."
fi

# Attach to the session
tmux attach-session -t $SESSION_NAME
