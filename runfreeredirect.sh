#!/bin/sh

echo "Starting Free Redirect (log output to /var/log/freeredirect)..."
cd /var/run
nohup /usr/local/bin/freeredirect >/var/log/freeredirect 2>&1 &

