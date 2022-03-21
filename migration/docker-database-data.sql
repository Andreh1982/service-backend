create table buyers(id serial primary key, nome varchar, cpf varchar, contatocel varchar);
create table products(id serial primary key, nome varchar, descricao varchar, quantidade varchar);
create table sellers(id serial primary key, nome varchar, cpf varchar, contatocel varchar, contatomail varchar, websitelink varchar);

INSERT INTO buyers(nome, cpf, contatocel) VALUES('Segundo Cliente das Flores', '23434712978', '61997744631');
INSERT INTO products(nome, descricao, quantidade) VALUES('Produto Extravagante', 'Com um material de extrema extravagancia, este produto é extravagante.', '23');
INSERT INTO sellers(nome, cpf, contatocel, contatomail, websitelink) VALUES('Vendedor Primordial', '45628263628', '61998827389', 'vendas@vendedor.com', 'www.vendasextravagantes.com');
