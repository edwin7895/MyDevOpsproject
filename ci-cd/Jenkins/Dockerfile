# Utiliza la imagen oficial de Jenkins como base
FROM jenkins/jenkins:lts

# Usuario root para instalar paquetes adicionales
USER root

# Instala curl, unzip y otras dependencias necesarias
RUN apt-get update && apt-get install -y \
    curl \
    unzip \
    ca-certificates \
    gnupg \
    lsb-release

# Configura el repositorio de Docker manualmente
RUN mkdir -p /etc/apt/keyrings \
    && curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

# Instala Docker CLI
RUN apt-get update && apt-get install -y docker-ce-cli

# Deshabilita la verificación SSL en curl e instala Terraform
RUN curl -k -LO https://releases.hashicorp.com/terraform/1.5.3/terraform_1.5.3_linux_amd64.zip \
    && unzip terraform_1.5.3_linux_amd64.zip \
    && mv terraform /usr/local/bin/ \
    && rm terraform_1.5.3_linux_amd64.zip

# Deshabilita la verificación SSL en curl e instala AWS CLI v2
RUN curl -k "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf awscliv2.zip aws

# Crea el grupo docker y agrega el usuario jenkins al grupo docker
RUN groupadd docker && usermod -aG docker jenkins

# Cambia de nuevo al usuario Jenkins
USER jenkins

# Establece el directorio de trabajo en el home de Jenkins
WORKDIR /var/jenkins_home

# Exposición de puertos
EXPOSE 8080
EXPOSE 50000

# Comando por defecto para iniciar Jenkins
CMD ["bash", "-c", "/usr/local/bin/jenkins.sh"]
