export RUNNER := 'podman'
export IMAGE := 'w-and-w'

_default:
  @just --list

# Build the podman/docker image
build:
  @{{ RUNNER }} build -t {{ IMAGE }} -f Containerfile .

# Play the game WEREWOLVES AND WANDERER
play:
  @{{ RUNNER }} run -it --rm {{ IMAGE }}

# Remove the image created
clean:
  @{{ RUNNER }} rmi {{ IMAGE }}
