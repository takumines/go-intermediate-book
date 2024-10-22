insert into articles (title, contents, username, nice, created_at) values
   ('first post', 'This is my first blog', 'takumi', 2, now());
insert into articles (title, contents, username, nice, created_at) values
   ('second post', 'This is my second blog', 'takumi', 3, now());

insert into comments (article_id, message) values
   (1, 'This is a comment');
insert into comments (article_id, message) values
   (2, 'This is a comment');