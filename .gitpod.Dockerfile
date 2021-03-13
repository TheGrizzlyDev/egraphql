FROM gitpod/workspace-full

WORKDIR /home/gitpod
RUN wget https://www.antlr.org/download/antlr-4.9-complete.jar

RUN echo 'export CLASSPATH=".:/home/gitpod/antlr-4.9-complete.jar:$CLASSPATH"' >> .bashrc
RUN echo 'antlr4 () { java -Xmx500M -cp "/home/gitpod/antlr-4.9-complete.jar:$CLASSPATH" org.antlr.v4.Tool }' >> .bashrc