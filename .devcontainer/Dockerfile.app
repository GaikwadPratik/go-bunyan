FROM gaikwadpratik/ubuntu-vscode-golang AS development

COPY [".devcontainer/testenv.sh","/usr/bin/testenv.sh"]

COPY --chown=vscode:vscode . ${WORKDIR}