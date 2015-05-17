#!/usr/bin/env bash
gcloud compute instances create zoop \
    --image container-vm \
    --metadata-from-file google-container-manifest=containers.yaml \
    --zone us-central1-a \
    --machine-type f1-micro
