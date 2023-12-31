#!/bin/bash

# folder names for the images
IMAGES=("identity-service")

for IMAGE in $IMAGES;
do
    # Set the image name
    IMAGE_NAME="$IMAGE-test"

    # Build the Docker image
    docker build -t $IMAGE_NAME -f ./$IMAGE/Dockerfile.test ./$IMAGE

    # Run the Docker container
    docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock --env-file=./.env $IMAGE_NAME sh -c "go test ./..."
done
