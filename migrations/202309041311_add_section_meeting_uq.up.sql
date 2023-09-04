ALTER TABLE IF EXISTS public.section_meeting
    ADD CONSTRAINT section_meeting_uq UNIQUE (code, type, start_date, start_time, meeting_days, location);
ALTER TABLE IF EXISTS public.section_meeting DROP CONSTRAINT IF EXISTS course_section_section_meeting_fk;

ALTER TABLE IF EXISTS public.section_meeting
    ADD CONSTRAINT course_section_section_meeting_fk FOREIGN KEY (code)
    REFERENCES public.course_section (code) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;