# O terrível mundo dos ponteiros

## Ponteiro

No arquivo ```ponteiro.go``` eu faço um exercício da forma correta de se usar ponteiros em Go.

## Estruturas

Para mostrar o uso de estruturas de dados, foi criado o arquivo ```struts_json.go```. Ele também converte os dados em memória para JSON e de JSON para a memória.

## Go routines

O arquivo ```rotina.go``` exemplifica concorrência e paralelismo.

A linguagem Go tem o jeitinho dele de trabalhar com agendamento cooperativo de execução de tarefas. Quem gerencia as threads é o próprio runtime do Go. Como não chama o sistema operacional para abrir uma thread, é muito mais rápido. E cada thread ocupa apenas 2k de memória, ao invés de 1m com a thread do sistema operacional.

Este runtime tem algumas regras cooperativas para chamar a próxima thread: acessos de io bloqueantes, sleep ,looping infinito etc.

## Canais

Os canais servem para que threads se comuniquem compartilhando memória.

## Generics

Forma para definir dados genéricos.
