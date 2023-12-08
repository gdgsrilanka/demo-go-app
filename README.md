# Migration Seamless onto GKE from Cloud Run
Sample Go application: Migration Seamless onto GKE from Cloud Run

# Prerequisites 

- Google Cloud Account 
- Payment enabled OR cloud credits 
- Idenity and Access Management (IAM) API should be enabled to run Cloud Run
- Kubernetes Engine API should be enabled to run GKE
- GKE or GKE Autopilot cluster 

# Steps to migrate the application

## 1. Export Cloud Run yaml file
```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  annotations:
    run.googleapis.com/ingress: all
    run.googleapis.com/ingress-status: all
  labels:
    cloud.googleapis.com/location: us-central1
    commit-sha: 30add32cef487f1a5783d4b91d0886ba5730cd88
    gcb-build-id: 76dddc5b-8da5-400a-9454-62a8c7e765be
    gcb-trigger-id: 8f699efa-a34a-4818-9861-b8bfc6f49e4b
    gcb-trigger-region: global
    managed-by: gcp-cloud-build-deploy-cloud-run
  name: devfest-demo-test-1
  namespace: '875854415969'
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/maxScale: '100'
        run.googleapis.com/client-name: gcloud
        run.googleapis.com/client-version: 456.0.0
        run.googleapis.com/startup-cpu-boost: 'true'
      labels:
        client.knative.dev/nonce: bbtzrrkdxs
        commit-sha: 30add32cef487f1a5783d4b91d0886ba5730cd88
        gcb-build-id: 76dddc5b-8da5-400a-9454-62a8c7e765be
        gcb-trigger-id: 8f699efa-a34a-4818-9861-b8bfc6f49e4b
        gcb-trigger-region: global
        managed-by: gcp-cloud-build-deploy-cloud-run
        run.googleapis.com/startupProbeType: Default
    spec:
      containerConcurrency: 80
      containers:
      - image: us-central1-docker.pkg.dev/devfest-2023-407412/cloud-run-source-deploy/demo-go-app:30add32cef487f1a5783d4b91d0886ba5730cd88
        name: placeholder-1
        ports:
        - containerPort: 8080
          name: http1
        resources:
          limits:
            cpu: 1000m
            memory: 512Mi
        startupProbe:
          failureThreshold: 1
          periodSeconds: 240
          tcpSocket:
            port: 8080
          timeoutSeconds: 240
      serviceAccountName: chamod-compute@developer.gserviceaccount.com
      timeoutSeconds: 300
  traffic:
  - latestRevision: true
    percent: 100
```
## 2. Modify the YAML to match a Kubernetes deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: devfest-demo-test-1
  namespace: default
spec:
  template:
    metadata:
      labels:
        app: devfest-demo-test-1
    spec:
      containerConcurrency: 80
      containers:
      - image: busybox:latest
        name: placeholder-1
        ports:
        - containerPort: 8080
          name: http1
        resources:
          limits:
            cpu: 1000m
            memory: 512Mi
        startupProbe:
          failureThreshold: 1
          periodSeconds: 240
          tcpSocket:
            port: 8080
          timeoutSeconds: 240
      timeoutSeconds: 300
    replicas: 1
    selector:
      matchLabels:
        app: devfest-demo-test-1
```

## 3. Deploy the application to GKE
```bash

kubectl apply -f deployment.yaml

```

## 4. Expose the application to the internet
```bash

kubectl expose deployment devfest-demo-test-1 --type=LoadBalancer --port 80 --target-port 8080

```