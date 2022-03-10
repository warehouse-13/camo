FROM alpine

RUN apk add --update \
	musl-dev \
	gcc \
	libffi-dev \
	python3 \
	python3-dev \
	py3-pip

RUN pip install bcrypt

COPY ./bcrypt.sh /root/bcrypt.sh

CMD /root/bcrypt.sh
