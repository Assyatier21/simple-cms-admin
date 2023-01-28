--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Homebrew)
-- Dumped by pg_dump version 15.1

-- Started on 2022-12-30 18:46:26 WIB

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

DROP DATABASE "cms-repository";
--
-- TOC entry 3605 (class 1262 OID 16385)
-- Name: cms-repository; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "cms-repository" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';


ALTER DATABASE "cms-repository" OWNER TO postgres;

\connect -reuse-previous=on "dbname='cms-repository'"

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
-- TOC entry 5 (class 2615 OID 16386)
-- Name: cms-repository; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA "cms-repository";


ALTER SCHEMA "cms-repository" OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 16388)
-- Name: cms_article; Type: TABLE; Schema: cms-repository; Owner: postgres
--

CREATE TABLE "cms-repository".cms_article (
    id integer NOT NULL,
    title character varying NOT NULL,
    slug character varying NOT NULL,
    htmlcontent character varying NOT NULL,
    categoryid integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    metadata json
);


ALTER TABLE "cms-repository".cms_article OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16387)
-- Name: cms_article_id_seq; Type: SEQUENCE; Schema: cms-repository; Owner: postgres
--

ALTER TABLE "cms-repository".cms_article ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "cms-repository".cms_article_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 212 (class 1259 OID 16452)
-- Name: cms_category; Type: TABLE; Schema: cms-repository; Owner: postgres
--

CREATE TABLE "cms-repository".cms_category (
    id integer NOT NULL,
    title text NOT NULL,
    slug text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE "cms-repository".cms_category OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16451)
-- Name: cms_categoryid_seq; Type: SEQUENCE; Schema: cms-repository; Owner: postgres
--

ALTER TABLE "cms-repository".cms_category ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "cms-repository".cms_categoryid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3597 (class 0 OID 16388)
-- Dependencies: 210
-- Data for Name: cms_article; Type: TABLE DATA; Schema: cms-repository; Owner: postgres
--

INSERT INTO "cms-repository".cms_article OVERRIDING SYSTEM VALUE VALUES (2, 'Newest Article Updated', 'newest-article-updated', '<h1> This is new article </h1>', 3, '2022-12-19 23:20:15', '2022-12-28 15:22:34', '{"meta_title":"Title 3 Updated","meta_description":"This is Updated Description 3","meta_author":"Muhammad Sholeh","meta_keywords":["description updated","testing updated"],"meta_robots":["following","no-index"]}');
INSERT INTO "cms-repository".cms_article OVERRIDING SYSTEM VALUE VALUES (3, 'My Article Title 2', 'my-article-slug-v2', '<p>This is the HTML content of my article 2.</p>', 3, '2022-12-19 23:20:35', '2022-12-19 23:20:15', '{
  "meta_title": "Title 2",
  "meta_description": "This is Description 2",
  "meta_author": "Muhammad Sholeh",
  "meta_keywords": [
    "description",
    "testing2"
  ],
  "meta_robots": [
    "following",
    "no-index"
  ]
}');
INSERT INTO "cms-repository".cms_article OVERRIDING SYSTEM VALUE VALUES (4, 'My Article Title 3', 'my-article-slug-v3', '<p>This is the HTML content of my article 3.</p>', 1, '2022-12-19 23:20:48', '2022-12-19 23:20:15', '{
  "meta_title": "Title 3",
  "meta_description": "This is Description 3",
  "meta_author": "Muhammad Sholeh",
  "meta_keywords": [
    "description",
    "testing3"
  ],
  "meta_robots": [
    "following",
    "no-index"
  ]
}');
INSERT INTO "cms-repository".cms_article OVERRIDING SYSTEM VALUE VALUES (21, 'Newest Article Updated v2', 'newest-article-updated-2', '<p> This is newest updated article </p>', 3, '2022-12-28 15:55:27', '2022-12-28 15:59:57', '{"meta_title":"Title 3 Updated","meta_description":"This is Updated Description 3","meta_author":"Muhammad Sholeh","meta_keywords":["description updated","testing updated"],"meta_robots":["following","no-index"]}');
INSERT INTO "cms-repository".cms_article OVERRIDING SYSTEM VALUE VALUES (22, 'Newest Article', 'newest-article', '<p> This is newest article </p>', 1, '2022-12-29 14:30:12', '2022-12-29 14:30:12', '{"meta_title":"Title 3","meta_description":"This is Description 3","meta_author":"Muhammad Sholeh","meta_keywords":["description","testing3"],"meta_robots":["following","no-index"]}');


--
-- TOC entry 3599 (class 0 OID 16452)
-- Dependencies: 212
-- Data for Name: cms_category; Type: TABLE DATA; Schema: cms-repository; Owner: postgres
--

INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (1, 'Category 1', 'category-1', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (2, 'Category 2', 'category-2', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (3, 'Category 3', 'category-3', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (4, 'Category 4', 'category-4', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (5, 'Category 5', 'category-5', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (6, 'Category 6', 'category-6', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (7, 'Category 7', 'category-7', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (8, 'Category 8', 'category-8', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (9, 'Category 9', 'category-9', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (10, 'Category 10', 'category-10', '2022-12-19 23:46:18.330326', '2022-12-19 23:46:18.330326');
INSERT INTO "cms-repository".cms_category OVERRIDING SYSTEM VALUE VALUES (12, 'Newest Category', 'newest-category', '2022-12-29 14:31:00', '2022-12-29 14:31:00');


--
-- TOC entry 3606 (class 0 OID 0)
-- Dependencies: 209
-- Name: cms_article_id_seq; Type: SEQUENCE SET; Schema: cms-repository; Owner: postgres
--

SELECT pg_catalog.setval('"cms-repository".cms_article_id_seq', 23, true);


--
-- TOC entry 3607 (class 0 OID 0)
-- Dependencies: 211
-- Name: cms_categoryid_seq; Type: SEQUENCE SET; Schema: cms-repository; Owner: postgres
--

SELECT pg_catalog.setval('"cms-repository".cms_categoryid_seq', 12, true);


--
-- TOC entry 3452 (class 2606 OID 16400)
-- Name: cms_article cms_article_pkey; Type: CONSTRAINT; Schema: cms-repository; Owner: postgres
--

ALTER TABLE ONLY "cms-repository".cms_article
    ADD CONSTRAINT cms_article_pkey PRIMARY KEY (id);


--
-- TOC entry 3456 (class 2606 OID 16458)
-- Name: cms_category cms_category_pkey; Type: CONSTRAINT; Schema: cms-repository; Owner: postgres
--

ALTER TABLE ONLY "cms-repository".cms_category
    ADD CONSTRAINT cms_category_pkey PRIMARY KEY (id);


--
-- TOC entry 3454 (class 2606 OID 16408)
-- Name: cms_article slug; Type: CONSTRAINT; Schema: cms-repository; Owner: postgres
--

ALTER TABLE ONLY "cms-repository".cms_article
    ADD CONSTRAINT slug UNIQUE (slug);


-- Completed on 2022-12-30 18:46:26 WIB

--
-- PostgreSQL database dump complete
--

