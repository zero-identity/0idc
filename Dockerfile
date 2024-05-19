FROM alpine:3.9

COPY main .

EXPOSE 9096

CMD [ "./main" ]
