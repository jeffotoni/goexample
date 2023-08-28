## Cutelyst web api 

#### Install Ubuntu

```bash
$ sudo apt update
$ sudo apt upgrade

$ sudo apt install build-essential qt5-default libgrantlee5-dev libboost-dev libqt5sql5-psql libqt5sql5-mysql


# Instale as dependências necessárias para compilar o Cutelyst
sudo apt install git cmake

# Clone o repositório do Cutelyst
git clone https://github.com/cutelyst/cutelyst.git

# Entre no diretório clonado
cd cutelyst

# Crie e entre no diretório de compilação
mkdir build && cd build

# Configure e compile o projeto
cmake ..
make

# Instale o Cutelyst
sudo make install

# Atualize o cache do linker dinâmico
sudo ldconfig

```
