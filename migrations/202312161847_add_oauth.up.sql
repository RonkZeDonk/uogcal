CREATE TABLE IF NOT EXISTS public.oauth_user
(
    oauth_id character varying COLLATE pg_catalog."default" NOT NULL,
    user_uid uuid NOT NULL,
    CONSTRAINT oauth_user_pk PRIMARY KEY (user_uid, oauth_id),
    CONSTRAINT oauth_user_oauth_id_uq UNIQUE (oauth_id),
    CONSTRAINT oauth_user_user_id_uq UNIQUE (user_uid),
    CONSTRAINT uogcal_user_oauth_user_fk FOREIGN KEY (user_uid)
        REFERENCES public.uogcal_user (uid) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.oauth_user
    OWNER to postgres;
