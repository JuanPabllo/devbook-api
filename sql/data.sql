insert into users (name, nick, email, password)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- usuario3

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts(title, content, author_id)
values
("Post 1", "Conteúdo do post 1", 1),
("Post 2", "Conteúdo do post 2", 2),
("Post 3", "Conteúdo do post 3", 3);