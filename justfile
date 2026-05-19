# Build docker image locally from source
[group('Docker')]
build:
  podman build -t justfile:latest .
