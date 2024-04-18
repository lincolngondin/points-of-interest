# Pontos de Interesse por GPS
Nesse projeto foi criado um serviço REST para auxiliar na integração com receptores GPS, a API permite o cadastro de POIs (Point of interest) em um banco de dados, listar esses pontos de interesse retornando um JSON para o cliente e listar os pontos de interesse que estão na proximidade de um ponto de referencia. O projeto é a solução para o seguinte [desafio backend](https://github.com/backend-br/desafios/blob/master/points-of-interest/PROBLEM.md)

## Tecnologias usadas
+ Go
+ Postgres

## Sobre o projeto

+ Para cadastrar um novo POI:

```
POST /pois
{
    "nome_do_poi": "Lanchonete",
    "x_coordenada": 27,
    "y_coordenada": 12
}
```

+ Listar todos os POIs cadastrados:

```
GET /pois
[
    {
        "nome_do_poi": "Lanchonete",
        "x_coordenada": 27,
        "y_coordenada": 12
    },
    {
        "nome_do_poi": "Posto",
        "x_coordenada": 31,
        "y_coordenada": 18
    },
    {
        "nome_do_poi": "Joalheria",
        "x_coordenada": 15,
        "y_coordenada": 12
    },
    ...
]
```

+ Listar todos os POIs por proximidade, a uma distanca menor ou igual a d_max do ponto (x,y):

```
GET /pois?x=20&y=10&d_max=10
[
    {
        "nome_do_poi": "Lanchonete",
        "x_coordenada": 27,
        "y_coordenada": 12
    },
    {
        "nome_do_poi": "Joalheria",
        "x_coordenada": 15,
        "y_coordenada": 12
    },
    ...
]
```

