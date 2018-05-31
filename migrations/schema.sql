--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

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
-- Name: authors; Type: TABLE; Schema: public; Owner: nghiatran
--

CREATE TABLE public.authors (
    id uuid NOT NULL,
    name text NOT NULL,
    born text NOT NULL,
    died text NOT NULL,
    nationality text NOT NULL,
    wikipedia text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.authors OWNER TO nghiatran;

--
-- Name: photos; Type: TABLE; Schema: public; Owner: nghiatran
--

CREATE TABLE public.photos (
    id uuid NOT NULL,
    name text NOT NULL,
    image_url text NOT NULL,
    author_id uuid NOT NULL,
    width integer NOT NULL,
    height integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.photos OWNER TO nghiatran;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: nghiatran
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO nghiatran;

--
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: nghiatran
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- Name: photos photos_pkey; Type: CONSTRAINT; Schema: public; Owner: nghiatran
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: nghiatran
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: photos photos_authors_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: nghiatran
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_authors_id_fk FOREIGN KEY (author_id) REFERENCES public.authors(id) ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

