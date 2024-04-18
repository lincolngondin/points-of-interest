# Pontos de Interesse por GPS


# Sobre o projeto

+ Para cadastrar um novo POI:

POST /pois
{
    "nome_do_poi": "Lanchonete",
    "x_coordenada": 27,
    "y_coordenada": 12
}

+ Listar todos os POIs cadastrados:

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

+ Listar todos os POIs por proximidade, a uma distanca menor ou igual a d_max do ponto (x,y):

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

