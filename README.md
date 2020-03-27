# NEGO
[![Build Status](https://github.com/InsideCI/nego/workflows/NEGO/badge.svg)](https://github.com/InsideCI/nego/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/InsideCI/nego)](https://goreportcard.com/report/github.com/InsideCI/nego)
[![codecov](https://codecov.io/gh/InsideCI/nego/branch/master/graph/badge.svg)](https://codecov.io/gh/InsideCI/nego)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/InsideCI/nego)](https://www.tickgit.com/browse?repo=github.com/InsideCI/nego)
[![Chat](https://badgen.net/badge/discord/chat?icon=discord)](https://discord.gg/jVdUJ75)

NEGO is a UFPB SIGAA Restful API created with Golang for study purposes only, but it may fit your needs.

Feel free to use and contribute.

### MAIN FEATURES:

- Readable and simple code;
- Uses [GORM](github.com/jinzhu/gorm) as the default database management tool for a more dynamic and generic approach to databases;
- Router specification created with [CHI](github.com/go-chi/chi) reliable router structure;
- Generic repository and middlewares;
- Multi database support for each resource;
- Multi ARCH: created for both REST and gRPC support in mind;
- Complete fetching capabilities: pagination, sorting, filtering.

### AVAILABLE RESOURCES:

`/centers` or `/centers/1856`
```json
{
  "id": 1856,
  "nome": "CENTRO DE INFORMÁTICA (CI)"
}
```

`/departments` or `/departments/2151`
```json
{
  "id": 2151,
  "nome": "DEPARTAMENTO DE INFORMÁTICA", "idCentro": "1856"
}
```

`/teachers` or `/teachers/1743917`
```json
{
  "id": 1743917,
  "nome": "THAIS GAUDENCIO DO REGO",
  "grau": "DOUTOR",
  "idDepartamento": 2151
}
```

`/courses` or `/courses/1626865`
```json
{
  "id": 1626865,
  "nome": "ENGENHARIA DE COMPUTAÇÃO/CI",
  "local": "João Pessoa",
  "tipo": "Presencial",
  "coordenador": "CHRISTIAN AZAMBUJA PAGOT",
  "idCentro": 1856
}
```

`/students` or `/students/11409558`
```json
{
  "matricula": 11409558,
  "nome": "CLEANDERSON LINS COUTINHO",
  "idCurso": 1626865
}
```

`/classes` or `/classes/GDSCO0081`
```json
{
  "id": "GDSCO0081",
  "nome": "SISTEMAS EMBARCADOS I",
  "turma": "01",
  "professor": "ALISSON VASCONCELOS DE BRITO",
  "horario": "24T23",
  "idCurso": 1626865
}
```

### FETCHING FEATURES:

|Type|Example| 
|----|----| 
|Sorting|`students?sort=nome`, `students?sort=idCurso`|
|Filtering|`students?nome=cleanderson`, `students?idCurso=1626865`|
|Pagination|`students?page=3`|

And of course, you can use all of them at the same time:

`/teachers?nome=mardson&sort=nome&page=0`

It would give us:

```json
{
  "totalElements": 1,
  "totalPages": 1,
  "limit": 10,
  "page": 0,
  "payloadSize": 1,
  "payload": [
    {
      "id": 1122252,
      "nome": "MARDSON FREITAS DE AMORIM",
      "grau": "DOUTOR",
      "idDepartamento": 1858
    }
  ]
}
```

#### OBSERVATIONS:
- The `page` parameter starts by index zero, until {`totalPages` - 1};
- `totalElements` will change dinamically based on your filtering, if you have one, and not by the resource amount or the payload.
- Any field of the models described above can be used as a filter as long as it is written in lower camel case: `idCurso`;
- If more than one value to the parameter was provived, only the last will be considered, except the `sort` parameter.

...and more to come.

### MODELS RELATION LOGIC:

```sql
    CREATE TABLE CENTROS (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR (100) NOT NULL
    );
    
    CREATE TABLE DEPARTAMENTOS (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR (100) NOT NULL,
            id_centro INT NOT NULL,
            FOREIGN KEY (id_centro) REFERENCES centros(id)
    );
    
    CREATE TABLE CURSOS (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR (100) NOT NULL,
            cidade VARCHAR (100),
            tipo VARCHAR (100),
            coordenador VARCHAR (100) NOT NULL,
            id_centro INT NOT NULL,
            FOREIGN KEY (id_centro) REFERENCES centros(id)
    );
    
    CREATE TABLE TURMAS (
            codigo VARCHAR(10) NOT NULL PRIMARY KEY,
            disciplina VARCHAR (100) NOT NULL,
            turma INT NOT NULL,
            professor VARCHAR (50) NOT NULL,
            horario VARCHAR (10) NOT NULL,
            --sala VARCHAR(10),
            id_curso INT NOT NULL,
            FOREIGN KEY (id_curso) REFERENCES cursos(id)
    );
    
    CREATE TABLE ALUNOS (
            matricula BIGINT NOT NULL PRIMARY KEY,
            nome VARCHAR(100) NOT NULL,
            id_curso INT NOT NULL,
            FOREIGN KEY (id_curso) REFERENCES cursos(id)
    );
    
    CREATE TABLE PROFESSORES (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR(100) NOT NULL,
            grau VARCHAR(20),
            id_departamento INT NOT NULL,
            FOREIGN KEY (id_departamento) REFERENCES departamentos(id)
    );
```

### USAGE DETAILS

Before everything:

* You need at least an empty PostgreSQL database created with the basic model logic above.
* There's also a custom scrapper created just for this project, called [SUS](github.com/InsideCI/sus), so you can populate your database with real data.

Your must create a `app.env` file on root folder and fill those parameters:

```.env
db_name=yourDatabaseName
db_pass=yourDatabasePassword
db_user=yourDatabaseLogin
db_host=yourDatabaseIpaddress
db_port=yourDatabasePort

api_port=yourApiPort
```

And then, you can run this project by:

`go build .`

`./nego`

You now can make requests with any REST client as [Insomnia](https://github.com/getinsomnia) or web any application at the address:

`http://localhost:yourApiPort`
