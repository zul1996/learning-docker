# learning-docker

## Hasil Instalasi docker

![Docker Login](./run-existing-docker-image/LoginDocker.png)

## Version Docker Installed

![Docker Version](./run-existing-docker-image/VersionDocker.png)

## Running Image Docker docker/welcome-to-docker

Install dengan perintah dibawah ini

1. -d untuk running background terminal
2. --name untuk pemberian nama Image
3. 8080:80 port

```bash
  docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker
```
