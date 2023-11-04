FROM dockcenter/velocity:stable

EXPOSE 25577

RUN wget https://api.papermc.io/v2/projects/velocity/versions/3.2.0-SNAPSHOT/builds/294/downloads/velocity-3.2.0-SNAPSHOT-294.jar
RUN mv velocity-*.jar /opt/velocity/velocity.jar
COPY velocity.toml /data