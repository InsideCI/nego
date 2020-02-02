# NEGO
[![Build Status](https://travis-ci.com/InsideCI/nego.svg?branch=master)](https://travis-ci.com/InsideCI/nego)
[![Go Report Card](https://goreportcard.com/badge/github.com/InsideCI/nego.svg?branch=master)](https://goreportcard.com/badge/github.com/InsideCI/nego)

NEGO is a UFPB SIGAA Restful API created with Golang for study purposes only, but it may fit your needs.

### MAIN FEATURES:

- Uses [GORM](github.com/jinzhu/gorm) as the default database management tool for a more dynamic approach to databases.
- Router specification created with [CHI](github.com/go-chi/chi) reliable router structure.
- Readable and simple code.

### AVAILABLE RESOURCES:

`API_VERSION/centers`

`API_VERSION/departments`

`API_VERSION/courses`

`API_VERSION/students`

`API_VERSION/disciplines`

...and more to come.

Today's API version is 'v1'.

### MODELS STRUCTURE:

```sql
    CREATE TABLE centros (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR (100) NOT NULL
    );
    
    CREATE TABLE departamentos (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR (100) NOT NULL,
            id_centro INT NOT NULL,
            FOREIGN KEY (id_centro) REFERENCES centros(id)
    );
    
    CREATE TABLE cursos (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR (100) NOT NULL,
            cidade VARCHAR (100),
            tipo VARCHAR (100),
            coordenador VARCHAR (100) NOT NULL,
            id_centro INT NOT NULL,
            FOREIGN KEY (id_centro) REFERENCES centros(id)
    );
    
    CREATE TABLE turmas (
            codigo VARCHAR(10) NOT NULL PRIMARY KEY,
            disciplina VARCHAR (100) NOT NULL,
            turma INT NOT NULL,
            professor VARCHAR (50) NOT NULL,
            horario VARCHAR (10) NOT NULL,
            --sala VARCHAR(10),
            id_curso INT NOT NULL,
            FOREIGN KEY (id_curso) REFERENCES cursos(id)
    );
    
    CREATE TABLE alunos (
            matricula BIGINT NOT NULL PRIMARY KEY,
            nome VARCHAR(100) NOT NULL,
            id_curso INT NOT NULL,
            FOREIGN KEY (id_curso) REFERENCES cursos(id)
    );
    
    CREATE TABLE professores (
            id INT NOT NULL PRIMARY KEY,
            nome VARCHAR(100) NOT NULL,
            grau VARCHAR(20) NOT NULL,
            id_departamento INT NOT NULL,
            FOREIGN KEY (id_departamento) REFERENCES departamentos(id)
    );
```

### USAGE DETAILS

#### BEFORE EVERYTHING

Your must create a `app.env` file on root folder and fill those parameters:

```.env
db_name=yourdatabasename
db_pass=yourdatabasepassword
db_user=yourdatabaselogin
db_host=yourdatabaseipaddress
db_port=yourdatabaseport
```

And then, you can run this project by:

`go build .`

`./nego`
