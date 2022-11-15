CREATE TABLE public.users (
                              id  SERIAL PRIMARY KEY,
                              first_name character varying(255) DEFAULT ''::character varying NOT NULL,
                              last_name character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE TABLE public.note (
                             id  SERIAL PRIMARY KEY,
                             message text not null,
                             user_id int not null
);

