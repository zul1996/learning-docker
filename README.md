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

```sh
docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker
```

## Running Preview

![Docker Image running preview](./run-existing-docker-image/PreviewRunning.png)git

## Cek Log Running container

1. welcome 1 nama file containernya
2. tail untuk cek logs
3. 20 jumlah log terakhir

```sh
docker logs -f --tail 20 welcome1
```

![Docker Logs container](./run-existing-docker-image/LoginDocker.png)git

## execute file in container

1. Lakukan exec di docker secara langsung

```sh
cat /usr/share/nginx/html/index.html
```

![Docker exec container](./run-existing-docker-image/docker_exec.png)git

## docker stop container service

```sh
 docker stop welcome1
```

![Docker stop service container](./run-existing-docker-image/docker_stop.png)git
