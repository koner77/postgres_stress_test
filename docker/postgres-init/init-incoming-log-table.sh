#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    -- Table: public.incoming_logs

    -- DROP TABLE public.incoming_logs;

    CREATE TABLE public.incoming_logs
    (
        request_id uuid NOT NULL,
        request jsonb,
        response jsonb,
        ip character varying(15) COLLATE pg_catalog."default",
        provider text COLLATE pg_catalog."default",
        general_provider text COLLATE pg_catalog."default",
        request_at timestamp with time zone,
        response_at timestamp with time zone,
        response_time numeric(125,0),
        added_at timestamp with time zone,
        events jsonb,
        fallback_reason text COLLATE pg_catalog."default",
        CONSTRAINT incoming_logs_pkey PRIMARY KEY (request_id)
    )

    TABLESPACE pg_default;

    ALTER TABLE public.incoming_logs
        OWNER to postgres;
EOSQL
