# NEGO
[![Build Status](https://github.com/InsideCI/nego/workflows/NEGO/badge.svg)](https://github.com/InsideCI/nego/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/InsideCI/nego)](https://goreportcard.com/report/github.com/InsideCI/nego)
[![codecov](https://codecov.io/gh/InsideCI/nego/branch/master/graph/badge.svg)](https://codecov.io/gh/InsideCI/nego)

NEGO is a UFPB SIGAA Restful API created with Golang for study purposes only, but it may fit your needs.

### MAIN FEATURES:

- Uses [GORM](github.com/jinzhu/gorm) as the default database management tool for a more dynamic approach to databases.
- Router specification created with [CHI](github.com/go-chi/chi) reliable router structure.
- Readable and simple code.

### AVAILABLE RESOURCES:

`/centers`

```json
{
  "id": 1856,
  "nome": "CENTRO DE INFORMÁTICA (CI)"
}
```

`/departments`
```json
{
  "id": 2151,
  "nome": "DEPARTAMENTO DE INFORMÁTICA", "idCentro": "1856"
}
```

`/teachers`
```json
{
  "id": 1743917,
  "nome": "THAIS GAUDENCIO DO REGO",
  "grau": "DOUTOR",
  "idDepartamento": 2151
}
```

`/courses`
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

`/students`
```json
{
  "matricula": 11409558,
  "nome": "CLEANDERSON LINS COUTINHO",
  "idCurso": 1626865
}
```

`/classes`
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

...and more to come.

Today's API version is 'v1'.

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
            grau VARCHAR(20) NOT NULL,
            id_departamento INT NOT NULL,
            FOREIGN KEY (id_departamento) REFERENCES departamentos(id)
    );
```

### USAGE DETAILS

Before everything:

* You need at least an empty PostgreSQL database created with the basic model logic above.
* There's also a custom scrapper created just for this project, called [SUS](github.com/InsideCI/sus), so you can populate your database with real data.

Your must create a `app.env` file on root folder and fill those parameters:

```yaml
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

You now can make requests with any REST client as [Insomnia](https://github.com/getinsomnia) or web application at the address:

`http://localhost:yourApiPort`
