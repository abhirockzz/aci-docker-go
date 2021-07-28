# Deploy multi-container app to Azure Container Instances using Docker Compose

A producer and consumer app pair sending and receiving messages from to/from Redis. You will deploy these to the same container group in Azure Container Instances using Docker CLI.

## You need

Azure CLI - You must have the Azure CLI installed on your local computer. Version 2.10.1 or later is recommended. Run az --version to find the version. If you need to install or upgrade, see Install the Azure CLI.

Docker Desktop - You must use Docker Desktop version 2.3.0.5 or later, available for Windows or macOS. Or install the Docker ACI Integration CLI for Linux.

> Before you proceed, clone this repo.

## Deploy to Azure

Replace Redis host, post and password details in `docker-compose.yaml` file.

Run locally:

```bash
docker compose up
docker compose down
```

Create context:

```bash
docker login azure
docker context create aci acictx
```

Switch to new context:

```bash
docker context ls
docker context use acictx
```

Deploy to Azure Container Instances (using Docker CLI)

```bash
docker compose up
```

Check logs:

```bash
docker logs -f aci-docker_producer
docker logs -f aci-docker_consumer
```

Navigate to Azure portal and check ACI container group

TODO img

To delete:

```bash
docker compose down
```