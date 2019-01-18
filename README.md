# Kurogo
This is a simple template processor by golang.

# How to install
```bash
go get -u github.com/5hyn3/kurogo/kurogo
```


# How to use

1. Write config file with yaml.

2. Write Template file.
    - You can use text/template`s syntax.There are also functions expanded with this application.
Please refer to item "Template Syntax" below

3. Run Command.
```bash
kurogo *.yml
```

Please refer to sample/Simple.

## Template Syntax
- text/template`s syntax
  - This application is thin wrapper that text/template. So you can use text/template`s syntax.

- Println "any string"
  - Show string in cli.

- RequestParameter "parameter description"
  - Prompt the user for parameters.

- SetGlobalParameter "parameter name"
  - Set GlobalParameter.GlobalParameter must be set in the Config file beforehand.

# Origin of name
This name is from [Kurogo(黒子、くろご、くろこ)](https://en.wikipedia.org/wiki/Kurogo).
