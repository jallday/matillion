CREATE TABLE IF NOT EXISTS public.films (
    id character varying(36) NOT NULL, 
    title character varying(128) NOT NULL,
    episode_id integer DEFAULT NULL,
    director character varying(128) DEFAULT NULL,
    producer character varying(128) DEFAULT NULL,
    release_date timestamp DEFAULT NULL,
    created_at bigint,
    updated_at bigint,
    deleted_at bigint
);

ALTER TABLE public.films OWNER TO mtuser;

CREATE TABLE IF NOT EXISTS public.ratings (
    id character varying(36) NOT NULL, 
    author character varying(128) NOT NULL,
    film_id character varying(36) NOT NULL,
    score integer DEFAULT NULL,
    created_at bigint,
    updated_at bigint,
    deleted_at bigint
);

ALTER TABLE public.ratings OWNER TO mtuser;

ALTER TABLE ONLY public.films DROP CONSTRAINT IF EXISTS films_pkey;

ALTER TABLE ONLY public.films
    ADD CONSTRAINT films_pkey UNIQUE (id);

ALTER TABLE ONLY public.films DROP CONSTRAINT IF EXISTS films_title_key;

ALTER TABLE ONLY public.films
    ADD CONSTRAINT films_title_key UNIQUE (title);

ALTER TABLE ONLY public.films DROP CONSTRAINT IF EXISTS episode_id_key;

ALTER TABLE ONLY public.films
    ADD CONSTRAINT episode_id_key UNIQUE (episode_id);


ALTER TABLE ONLY public.ratings DROP CONSTRAINT IF EXISTS ratings_pkey;

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_pkey UNIQUE (id);

ALTER TABLE ONLY public.ratings DROP CONSTRAINT IF EXISTS films_author_filmid_key;

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT films_author_filmid_key UNIQUE (author, film_id);
