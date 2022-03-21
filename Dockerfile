FROM alpine

ADD ./camo /root/camo

CMD ["sh", "-c", "echo -n $CLEAR_PASSWORD | /root/camo --encode"]
