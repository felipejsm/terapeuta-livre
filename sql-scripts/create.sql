-- Remover tabelas existentes, incluindo dependências
DROP TABLE IF EXISTS tb_file;
DROP TABLE IF EXISTS tb_file_metadata;
DROP TABLE IF EXISTS tb_patient;
DROP TABLE IF EXISTS tb_therapist;

-- Criação da tabela de terapeutas (precisa vir antes de tb_patient)
CREATE TABLE tb_therapist (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    login VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) UNIQUE,
    phone VARCHAR(20),     
    crp VARCHAR(20),       
    specialization VARCHAR(100),
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    profile_picture_key VARCHAR(255),
    timezone VARCHAR(50) DEFAULT 'America/Sao_Paulo',
    receive_notifications BOOLEAN DEFAULT 
);

-- Criação da tabela de pacientes
CREATE TABLE tb_patient (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    therapist_id INT NOT NULL,
    birth_date DATE,
    gender VARCHAR(20),       -- Pode ser enum('Masculino','Feminino','Outro','Prefiro não informar')
    phone VARCHAR(20),
    cpf VARCHAR(14) UNIQUE,
    rg VARCHAR(20),
    address TEXT,
    emergency_contact_name VARCHAR(100),
    emergency_contact_phone VARCHAR(20),
    health_insurance VARCHAR(100),
    health_insurance_number VARCHAR(50),
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT,              -- Observações gerais
    profile_picture_key VARCHAR(255),
    marital_status VARCHAR(30),
    profession VARCHAR(100),
    CONSTRAINT fk_therapist_patient FOREIGN KEY (therapist_id) REFERENCES tb_therapist(id) ON DELETE CASCADE
);

-- Criação da tabela de metadados de arquivos
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

-- Criação da tabela de arquivos (referência a tb_file_metadata)
CREATE TABLE tb_file (
    id SERIAL PRIMARY KEY,
    metadata_id INTEGER REFERENCES tb_file_metadata(id) ON DELETE CASCADE,
    file_data BYTEA NOT NULL
);

-- Inserir terapeutas
INSERT INTO tb_therapist (name, email, login, password) VALUES
    ('Dr. João Silva', 'joao.silva@clinic.com', 'joaosilva', 'hashed_password_1'),
    ('Dra. Maria Oliveira', 'maria.oliveira@clinic.com', 'mariaoliveira', 'hashed_password_2'),
    ('Dr. Pedro Santos', 'pedro.santos@clinic.com', 'pedrosantos', 'hashed_password_3'),
    ('Dra. Ana Costa', 'ana.costa@clinic.com', 'anacosta', 'hashed_password_4'),
    ('Dr. Lucas Almeida', 'lucas.almeida@clinic.com', 'lucasalmeida', 'hashed_password_5');

-- Inserir pacientes
INSERT INTO tb_patient (name, email, therapist_id) VALUES
    ('Carlos Silva', 'carlos.silva@patient.com', 1),
    ('Fernanda Lima', 'fernanda.lima@patient.com', 2),
    ('Bruno Rocha', 'bruno.rocha@patient.com', 3),
    ('Juliana Mendes', 'juliana.mendes@patient.com', 4),
    ('Rafael Pereira', 'rafael.pereira@patient.com', 5);

-- Inserir metadados de arquivos
INSERT INTO tb_file_metadata (file_name, object_key, extension, owner_id, file_size, owner_type) VALUES
    ('relatorio_carlos.pdf', 'files/carlos_silva/relatorio_carlos.pdf', '.pdf', 1, 2048, 'patient'),
    ('exame_fernanda.pdf', 'files/fernanda_lima/exame_fernanda.pdf', '.pdf', 2, 4096, 'patient'),
    ('receita_bruno.docx', 'files/bruno_rocha/receita_bruno.docx', '.docx', 3, 1024, 'patient'),
    ('relatorio_juliana.pdf', 'files/juliana_mendes/relatorio_juliana.pdf', '.pdf', 4, 2048, 'patient'),
    ('exame_rafael.pdf', 'files/rafael_pereira/exame_rafael.pdf', '.pdf', 5, 5120, 'patient');

-- Inserir arquivos (conteúdo fictício em formato hexadecimal)
INSERT INTO tb_file (metadata_id, file_data) VALUES
    (1, decode('255044462d312e350a25e2e3cf...', 'hex')), -- PDF fictício
    (2, decode('255044462d312e350a25e2e3cf...', 'hex')), -- PDF fictício
    (3, decode('504b0304140006000800000021...', 'hex')), -- DOCX fictício
    (4, decode('255044462d312e350a25e2e3cf...', 'hex')), -- PDF fictício
    (5, decode('255044462d312e350a25e2e3cf...', 'hex')); -- PDF fictício
