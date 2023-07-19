INSERT INTO roles (role_name)
VALUES ('teacher');
INSERT INTO roles (role_name)
VALUES ('student');

INSERT INTO users (name, surname, username, email, role_id)
VALUES ('Hilmi', 'SAFAK', 'hilmi511', 'hilmisafak1@gmail.com', 1),
       ('Ahmet', 'SAFAK', 'ahmet511', 'ahmetsafak1@gmail.com', 2),
       ('Sude', 'SAFAK', 'sude511', 'sudesafak1@gmail.com', 1),
       ('Sena', 'SAFAK', 'sena511', 'senasafak1@gmail.com', 2),
       ('Mustafa', 'SAFAK', 'mustafa511', 'mustafasafak1@gmail.com', 2);
