-- Remover tabelas existentes, incluindo dependências
DROP TABLE IF EXISTS tb_file_metadata CASCADE;
DROP TABLE IF EXISTS tb_file CASCADE;
DROP TABLE IF EXISTS tb_patient;
DROP TABLE IF EXISTS tb_therapist;

-- Criação da tabela de arquivos
CREATE TABLE tb_file_metadata (
    id SERIAL PRIMARY KEY,
    file_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    object_key VARCHAR(255) NOT NULL,
    extension VARCHAR(10) NOT NULL,
    owner_id INT NOT NULL,
    file_size INT NOT NULL,
    owner_type VARCHAR(10) NOT NULL
);

CREATE TABLE tb_file (
    id SERIAL PRIMARY KEY,
    metadata_id INTEGER REFERENCES tb_file_metadata(id) ON DELETE CASCADE,
    file_data BYTEA NOT NULL
)

CREATE TABLE tb_patient (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    therapist_id INT NOT NULL,
    CONSTRAINT fk_therapist_patient FOREIGN KEY (therapist_id) REFERENCES tb_therapist(id) ON DELETE CASCADE
);

CREATE TABLE tb_therapist (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    login VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);
