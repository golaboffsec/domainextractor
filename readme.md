# domainextractor

O **Domain Extractor** é uma ferramenta simples em Go que permite extrair domínios únicos de um arquivo de URLs. Ele filtra os domínios que contêm  entre pontos e salva o resultado em um novo arquivo.

## Instalação

Para instalar a ferramenta, utilize o seguinte comando:

```bash
go install github.com/golaboffsec/domainextractor@latest
```

## Exemplo
Suponha que o conteúdo de dominios.txt seja:

```bash
login.bol.com/authorize?client_id=...
www.facebook.com/bolpuntcom
careers.bol.com/kaliberjs/33.43bb319962561ec4f7c6.js
```

Após a execução do programa, filtrando por 'bol' o arquivo dominios_unicos.txt será o output e conterá:

```bash
login.bol.com
careers.bol.com
```
