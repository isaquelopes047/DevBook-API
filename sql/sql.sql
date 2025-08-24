DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(255) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id) 
    REFERENCES usuarios(id) 
    ON DELETE CASCADE,

    seguidor_id INT NOT NULL,
    FOREIGN KEY (seguidor_id) 
    REFERENCES usuarios(id) 
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id INT auto_increment PRIMARY KEY,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,
    autor_id INT NO NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas INT default 0
    criadaEm TIMESTAMP default current_timestamp
) ENGINE=INNODB;