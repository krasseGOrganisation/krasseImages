FROM ubuntu

COPY server.cfg .

RUN sed -i 's/30120/30124/g' server.cfg
RUN sed -i 's/\#sv_master1/sv_master1/g' server.cfg
RUN sed -i 's/sv_maxclients 32/sv_maxclients 10/g' server.cfg
