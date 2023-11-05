FROM dockcenter/velocity:stable

EXPOSE 25577

RUN wget https://api.papermc.io/v2/projects/velocity/versions/3.2.0-SNAPSHOT/builds/294/downloads/velocity-3.2.0-SNAPSHOT-294.jar
RUN mv velocity-*.jar /opt/velocity/velocity.jar
RUN wget https://github.com/Aelysium-Group/rusty-connector/releases/download/v0.7.0/rusty-connector-0.7.0.jar
RUN mkdir /data/plugins
RUN mv rusty-connector-*.jar /data/plugins/rusty-connector.jar
COPY velocity.toml /data
COPY --chown=velocity:velocity velocity_config  /data/plugins/rustyconnector-velocity/
RUN ln -s /mnt/secrets/private.key /data/plugins/rustyconnector-velocity/private.key