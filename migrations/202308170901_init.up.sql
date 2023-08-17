BEGIN;

CREATE TABLE public.uogcal_user (
    uid uuid NOT NULL,
    display_name character varying NOT NULL,
    password character varying NOT NULL,

    CONSTRAINT uogcal_user_pk PRIMARY KEY (uid),
    CONSTRAINT uogcal_user_uq UNIQUE (display_name)
);

CREATE TABLE public.course_section (
    code character varying NOT NULL,
    name character varying NOT NULL,

    CONSTRAINT course_section_pk PRIMARY KEY (code),
    CONSTRAINT course_section_uq UNIQUE (code)
);

CREATE TABLE public.section_meeting (
    code character varying NOT NULL,
    type character varying NOT NULL,
    created timestamp without time zone NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    start_time time without time zone NOT NULL,
    end_time time without time zone NOT NULL,
    meeting_days smallint[] NOT NULL,
    location character varying NOT NULL,
    last_modified timestamp without time zone NOT NULL,
    update_count integer NOT NULL,

    CONSTRAINT section_meeting_pk PRIMARY KEY (code, type, created)
);

CREATE TABLE public.uogcal_user_course_section (
    uogcal_user_uid uuid NOT NULL,
    course_section_code character varying NOT NULL,

    CONSTRAINT uogcal_user_course_section_uq UNIQUE (uogcal_user_uid, course_section_code)
);

ALTER TABLE public.section_meeting
    ADD CONSTRAINT course_section_section_meeting_fk FOREIGN KEY (code)
    REFERENCES public.course_section (code)
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE public.uogcal_user_course_section
    ADD CONSTRAINT uogcal_user_uogcal_user_course_section_fk FOREIGN KEY (uogcal_user_uid)
    REFERENCES public.uogcal_user (uid)
    ON UPDATE CASCADE
    ON DELETE CASCADE;


ALTER TABLE public.uogcal_user_course_section
    ADD CONSTRAINT course_section_uogcal_user_course_section_fk FOREIGN KEY (course_section_code)
    REFERENCES public.course_section (code)
    ON UPDATE CASCADE
    ON DELETE CASCADE;

END;