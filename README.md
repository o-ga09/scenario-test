# Scenario Test Sample

シナリオテストのツール試す

- Artillery
- Runn
- Stepci
- Scenarigo

## Install

```bash
# Artillery
npm i -g artillery

# Runn
go install github.com/k1LoW/runn/cmd/runn@latest

# Stepci
npm i -g stepci

# Scenarigo
go install github.com/zoncoen/scenarigo/cmd/scenarigo@v0.17.1

```

## Init

### Artillery
zero config

### Runn
zero config

### Stepci
zero config

### Scenarigo

```bash
scenarigo config init
```

## Run

```scenario.yml```は、シナリオファイル

### Artillery

```bash
artillery run scenario.yml
```
### Runn

```bash
runn run scenari.yml
```

### Stepci

```bash
stepci run scenario.yml
```

### Scenarigo

```bash
scenarigo run
```

## Note

scenarigoを使用して、pluginで拡張するためにはMac or Linuxのみサポートする。