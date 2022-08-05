// команда запуска контейнера с утилитами nats
```
docker run --rm -it --network intern_l0_l0 natsio/nats-box:latest
```
// очистить очередь nats
```
nats stream purge --server nats1:4222
```
// создание образа vegeta из папки с докерфайлом
```
docker build -t vegeta .
```
