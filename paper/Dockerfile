FROM marctv/minecraft-papermc-server:latest

COPY server.properties .
RUN mkdir config
COPY paper-global.yml config/
COPY setup-and-start-server.sh /opt/
RUN chmod +x /opt/setup-and-start-server.sh

ENTRYPOINT [ "/opt/setup-and-start-server.sh" ]