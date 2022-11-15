
CREATE TABLE public.users (
                              id integer NOT NULL,
                              first_name character varying(255) DEFAULT ''::character varying NOT NULL,
                              last_name character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.users OWNER TO markopoikkimaki;


CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


CREATE TABLE public.note (
                             id integer NOT NULL,
                             message text not null,
                             user_id int not null
);


ALTER TABLE public.note OWNER TO markopoikkimaki;


CREATE SEQUENCE public.note_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;