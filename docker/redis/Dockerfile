FROM redislabs/redisjson2:latest

ADD bootsync.sh /var/lib/boot2docker

RUN apt update && \
    apt install net-tools tzdata -y

ENV TZ=Asia/Kolkata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
	echo $TZ > /etc/timezone