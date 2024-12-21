Обновление Go компилятора на Ubuntu 20.04 LTS (серверной версии) можно выполнить несколькими способами. Вот пошаговая инструкция для обновления Go до последней версии.

---

### Способ 1: Установка через официальный PPA

#### 1. **Удалите старую версию Go (если она установлена)**
Если Go уже установлен, удалите его:
```bash
sudo apt remove --purge golang
sudo rm -rf /usr/local/go
```

#### 2. **Добавьте официальный PPA для Go**
Добавьте PPA (Personal Package Archive) для Go:
```bash
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
```

#### 3. **Установите Go**
Установите последнюю версию Go:
```bash
sudo apt install golang-go
```

#### 4. **Проверьте версию Go**
Убедитесь, что Go установлен корректно:
```bash
go version
```

---

### Способ 2: Установка через официальный архив Go

#### 1. **Скачайте последнюю версию Go**
Перейдите на [официальный сайт Go](https://go.dev/dl/) и скачайте архив с последней версией для Linux:
```bash
wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
```

#### 2. **Удалите старую версию Go (если она установлена)**
Если Go уже установлен, удалите его:
```bash
sudo rm -rf /usr/local/go
```

#### 3. **Распакуйте архив**
Распакуйте архив в папку `/usr/local`:
```bash
sudo tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
```

#### 4. **Настройте переменные окружения**
Добавьте Go в переменную `PATH`. Откройте файл `~/.bashrc` или `~/.zshrc` и добавьте строки:
```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Примените изменения:
```bash
source ~/.bashrc
```

#### 5. **Проверьте версию Go**
Убедитесь, что Go установлен корректно:
```bash
go version
```

---

### Способ 3: Использование Snap (если Snap установлен)

#### 1. **Установите Go через Snap**
Если у вас установлен Snap, вы можете установить Go следующей командой:
```bash
sudo snap install go --classic
```

#### 2. **Проверьте версию Go**
Убедитесь, что Go установлен корректно:
```bash
go version
```

---

### Способ 4: Автоматическое обновление через скрипт

Если вы хотите автоматизировать процесс обновления Go, вы можете использовать следующий скрипт:

```bash
#!/bin/bash

# Удаляем старую версию Go
sudo rm -rf /usr/local/go

# Скачиваем последнюю версию Go
LATEST_GO=$(curl -s https://go.dev/dl/ | grep -oP 'go[0-9\.]+(?=\.linux-amd64\.tar\.gz)')
wget https://go.dev/dl/${LATEST_GO}.linux-amd64.tar.gz

# Распаковываем архив
sudo tar -C /usr/local -xzf ${LATEST_GO}.linux-amd64.tar.gz

# Добавляем Go в PATH
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export GOPATH=\$HOME/go" >> ~/.bashrc
echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.bashrc

# Применяем изменения
source ~/.bashrc

# Проверяем версию Go
go version
```

Сохраните этот скрипт в файл, например, `update_go.sh`, сделайте его исполняемым и запустите:
```bash
chmod +x update_go.sh
./update_go.sh
```

---

### Итог
Теперь у вас установлена последняя версия Go компилятора. Вы можете выбрать любой из способов, который вам удобен.

Если у вас есть дополнительные вопросы или нужна помощь, дайте знать! 😊