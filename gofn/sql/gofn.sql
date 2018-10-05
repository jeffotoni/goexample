--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.9
-- Dumped by pg_dump version 9.6.9

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: login; Type: TABLE; Schema: public; Owner: gofn
--

CREATE TABLE public.login (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(200) NOT NULL,
    sucess integer NOT NULL,
    error integer NOT NULL
);


ALTER TABLE public.login OWNER TO gofn;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: gofn
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO gofn;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gofn
--

ALTER SEQUENCE public.user_id_seq OWNED BY public.login.id;


--
-- Name: login id; Type: DEFAULT; Schema: public; Owner: gofn
--

ALTER TABLE ONLY public.login ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: login; Type: TABLE DATA; Schema: public; Owner: gofn
--

INSERT INTO public.login VALUES (1, 'jefferson otoni lima', 'jeff.otoni@s3wf.com.br', 0, 0);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gofn
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- Name: login user_email_key; Type: CONSTRAINT; Schema: public; Owner: gofn
--

ALTER TABLE ONLY public.login
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: login user_pkey; Type: CONSTRAINT; Schema: public; Owner: gofn
--

ALTER TABLE ONLY public.login
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

