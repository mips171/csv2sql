#!/bin/bash

# Directory where your images are located
IMAGE_DIR="catalog/products"
OUTPUT_DIR="webp"

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Counter for the number of concurrent jobs
CONCURRENT_JOBS=0
MAX_CONCURRENT_JOBS=12

# Loop through all images in the directory
for image in "$IMAGE_DIR"/*; do
  # Get the filename without the extension
  filename=$(basename -- "$image")
  extension="${filename##*.}"
  filename="${filename%.*}"

  # Convert the image to WebP and send it to the background
  convert "$image" "${OUTPUT_DIR}/${filename}.webp" &

  # Increment the jobs counter
  ((CONCURRENT_JOBS++))

  # If we reach the max concurrent jobs, wait for all to complete before continuing
  if (( CONCURRENT_JOBS == MAX_CONCURRENT_JOBS )); then
    wait # Wait for all background jobs to finish
    CONCURRENT_JOBS=0 # Reset the jobs counter
  fi
done

# Wait for the last set of jobs to finish
wait

echo "Conversion to WebP completed."
