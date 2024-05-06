-- DROP DATABASE emple_dep;

CREATE DATABASE IF NOT EXISTS emple_dep;

USE emple_dep;

CREATE TABLE IF NOT EXISTS departament (
    departament_id 						INT AUTO_INCREMENT PRIMARY KEY,
    departament_name 					VARCHAR(50) NOT NULL,
    departament_address					VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS employees (
    file_number 						INT NOT NULL,
    employee_id 						INT AUTO_INCREMENT PRIMARY KEY,
    employee_name 						VARCHAR(50) NOT NULL,
    employee_last_name 					VARCHAR(50) NOT NULL,
    employee_birth_date 				DATE NOT NULL,
    incorporation_date 					DATE NOT NULL,
    job_role 							VARCHAR(50) NOT NULL,
    net_salary 							FLOAT NOT NULL,
    departament_id 						INT NOT NULL,
    FOREIGN KEY (departament_id) 		REFERENCES departament(departament_id)
);

INSERT INTO departament (departament_name, departament_address) VALUES
	('Vendas', 							'123 Rua Fictícia, Cidade Fictícia'),
	('Marketing', 						'456 Avenida Fictícia, Cidade Fictícia'),
	('Recursos Humanos',			 	'789 Praça Fictícia, Cidade Fictícia'),
	('Financeiro', 						'1011 Estrada Fictícia, Cidade Fictícia'),
	('Tecnologia', 						'1213 Rodovia Fictícia, Cidade Fictícia');

INSERT INTO employees (file_number, employee_name, employee_last_name, employee_birth_date, incorporation_date, job_role, net_salary, departament_id) VALUES
	(1, 	'João', 	'Silva', 		'1980-01-01', '2000-01-01', 'Vendedor', 				3000.00, 1),
	(2, 	'Maria', 	'Santos', 		'1985-02-02', '2005-02-02', 'Analista de Marketing', 	3500.00, 2),
	(3, 	'Pedro', 	'Oliveira', 	'1990-03-03', '2010-03-03', 'Recrutador', 				4000.00, 3),
	(4, 	'Ana', 		'Rodrigues', 	'1975-04-04', '1995-04-04', 'Contador', 				4500.00, 4),
	(5, 	'Carlos', 	'Ferreira', 	'1988-05-05', '2008-05-05', 'Desenvolvedor', 			5000.00, 5),
	(6, 	'Laura', 	'Almeida', 		'1993-06-06', '2013-06-06', 'Vendedor', 				3000.00, 1),
	(7, 	'Rafael', 	'Martins', 		'1982-07-07', '2002-07-07', 'Analista de Marketing', 	3500.00, 2),
	(8, 	'Fernanda', 'Pereira', 		'1995-08-08', '2015-08-08', 'Recrutador', 				4000.00, 3),
	(9, 	'Rodrigo', 	'Gomes', 		'1978-09-09', '1998-09-09', 'Contador', 				4500.00, 4),
	(10, 	'Juliana', 	'Costa', 		'1986-10-10', '2006-10-10', 'Desenvolvedor', 			5000.00, 5),
	(11, 	'Gustavo', 	'Nunes', 		'1984-11-11', '2004-11-11', 'Vendedor', 				3000.00, 1),
	(12, 	'Aline', 	'Lima', 		'1992-12-12', '2012-12-12', 'Analista de Marketing', 	3500.00, 2),
	(13, 	'Ricardo', 	'Sousa', 		'1983-11-13', '2003-11-13', 'Recrutador', 				4000.00, 3),
	(14, 	'Isabela', 	'Rocha', 		'1977-11-14', '1997-11-14', 'Contador', 				4500.00, 4),
	(15, 	'Mariana', 	'Pires', 		'1989-11-15', '2009-11-15', 'Desenvolvedor', 			5000.00, 5);

SELECT * FROM departament;
SELECT * FROM employees;



