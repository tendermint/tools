FROM alpine:3.7
MAINTAINER Greg Szabo <greg@tendermint.com>

#Default home for tendermint. The node command will look for $TMHOME/config/genesis.json at initialization.
ENV TMHOME /tendermint11
ENV ID 1
ENV SIZE 1
RUN apk update && \
    apk upgrade && \
    apk --no-cache add curl jq file && \
    addgroup tmuser && \
    adduser -S -G tmuser tmuser -h "$TMHOME"

USER tmuser
VOLUME [ $TMHOME ]
WORKDIR $TMHOME
EXPOSE 46656 46657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["node", "--moniker=`hostname`"]
STOPSIGNAL SIGTERM

ARG BINARY=tendermint
COPY $BINARY /usr/bin/tendermint
COPY wrapper.sh /usr/bin/wrapper.sh
COPY tendermin11 /tendermint11
COPY tendermin14 /tendermint14
COPY tendermin24 /tendermint24
COPY tendermin34 /tendermint34
COPY tendermin44 /tendermint44

