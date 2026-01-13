# syntax=docker/dockerfile:1.4
FROM docker.io/library/ubuntu:22.04

RUN <<EOF
echo "Hello"
echo "World"
mkdir -p /hello/world
# not detected by checkov:
apt update
EOF

# detected by checkov:
RUN apt update
# ID-1768294475-3064da8a
