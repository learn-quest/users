--
-- PostgreSQL database dump
--

-- Dumped from database version 16.6 (Debian 16.6-1.pgdg120+1)
-- Dumped by pg_dump version 16.6 (Debian 16.6-1.pgdg120+1)
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;
--
-- Name: citus; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citus WITH SCHEMA pg_catalog;
--
-- Name: EXTENSION citus; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citus IS 'Citus distributed database';
--
-- Name: citus_columnar; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citus_columnar WITH SCHEMA pg_catalog;
--
-- Name: EXTENSION citus_columnar; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citus_columnar IS 'Citus Columnar extension';
--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;
--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';
--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
SET default_tablespace = '';
SET default_table_access_method = heap;
--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    _id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    username character varying(100) NOT NULL,
    is_banned boolean DEFAULT false,
    profile_pic text,
    country character varying(100) NOT NULL,
    last_logged_in timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE public.users OWNER TO postgres;
--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
ADD CONSTRAINT users_pkey PRIMARY KEY (_id);
-- Distribute the table by hashing on 'shard_key_column'
SELECT create_distributed_table('users', '_id');
--
-- Name: users_email_hash_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX users_email_hash_index ON public.users USING hash (email);
--
-- Name: users_username_hash_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX users_username_hash_index ON public.users USING hash (username);
--
-- PostgreSQL database dump complete
--