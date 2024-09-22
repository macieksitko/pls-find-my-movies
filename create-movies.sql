-- SQLite
create table movies (
    id integer primary key autoincrement,
    title text not null,
    description text
);

insert into movies (title, description)
values ('The Matrix', 'A computer hacker learns about the true nature of his reality and his role in the war against its controllers.')

insert into movies (title, description)
values ('Asteroid City', 'In the aftermath of World War II, an American astronomer leads a team of scientists on a mission to find evidence of extraterrestrial life on the moon.')

insert into movies (title, description)
values ('The Batman', 'When a sadistic serial killer begins murdering key political figures in Gotham, Batman is forced to confront his dark past.')

insert into movies (title, description)
values ('The Godfather', 'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.')

insert into movies (title, description)
values ('The Dark Knight', 'When the menace known as the Joker emerges from his mysterious past, he wreaks havoc and chaos on the people of Gotham.')

select *
from movies

CREATE TABLE IF NOT EXISTS streamings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    service TEXT NOT NULL,
    link TEXT NOT NULL
);
INSERT INTO streamings (service, link) VALUES
    ('Netflix', 'https://www.netflix.com'),
    ('Amazon Prime Video', 'https://www.amazon.com/Prime-Video'),
    ('HBO Max', 'https://www.hbomax.com');

select * from streamings;

DROP TABLE IF EXISTS movie_streamings;

CREATE TABLE IF NOT EXISTS movie_streamings (
    movie_id INTEGER,
    streaming_id INTEGER,
    movie_url TEXT,
    PRIMARY KEY (movie_id, streaming_id),
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (streaming_id) REFERENCES streamings(id)
);

-- Example insertions (adjust according to your actual movie and streaming IDs)
INSERT INTO movie_streamings (movie_id, streaming_id, movie_url) VALUES
    (1, 1, 'https://www.netflix.com/watch/12345'),  -- Movie 1 on Netflix, free with subscription
    (1, 2, 'https://www.amazon.com/gp/video/detail/B00123ABCD'),  -- Movie 1 on Amazon Prime Video, $3.99
    (2, 3, 'https://play.hbomax.com/movie/urn:hbo:movie:ABCDEF123456'),  -- Movie 2 on HBO Max, free with subscription
    (3, 1, 'https://www.netflix.com/watch/67890'),  -- Movie 3 on Netflix, free with subscription
    (3, 2, 'https://www.amazon.com/gp/video/detail/B00456WXYZ');  -- Movie 3 on Amazon Prime Video, $2.99

-- Query to check the inserted data
SELECT * FROM movie_streamings;
